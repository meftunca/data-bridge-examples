
# Architecture Overview

## Design Philosophy

DataBridge V2 follows three core principles:

1. **Schema-First Development** -- The database schema is the single source of truth.
   All API code derives from introspecting PostgreSQL `information_schema`.

2. **Zero Runtime Dependency on Generator** -- Generated code is self-contained.
   It depends only on Fiber, GORM, and the pagination/SSE/events packages.

3. **Explicit Context Over Global State** -- All configuration flows through
   `GeneratorContext`, never through environment variables or package-level globals.

---

# System Architecture

```
+-----------------------------------------------------------------+
|                     CLI / Embedded Entry                         |
|  cmd/databridge-v2/main.go  <-->  app/app.go + app/run.go       |
|                     | reads config.yaml                         |
|             app/config.go (Config, Project, SchemaConfig)        |
+-----------------------------------------------------------------+
                      |
                      | creates GeneratorContext
                      | opens PostgreSQL via db/db.go
                      v
+-----------------------------------------------------------------+
|                   parser/creator.go                             |
|  Per Schema:                                                     |
|  1. GenerateEnumTypes()     -> enums.go                          |
|  2. GetTables()             -> SQL introspection                 |
|  3. GetRPCFunctions()       -> rpc__ function discovery          |
|  4. Per Table (parallel):                                        |
|     - StructCreator         -> structures/{table}.go             |
|     - ServiceCreator        -> services/{table}.go               |
|     - ControllerCreator     -> controllers/{table}.go            |
|     - TestCreator           -> tests/{table}_test.go             |
|  5. EndPointCreator         -> {schema}/{schema}.go              |
|  6. EventSetupCreator       -> event_handlers/                   |
|  7. RPCCreator              -> rpc services + controllers        |
+-----------------------------------------------------------------+
```

---

# Layered Architecture

The generated code follows a strict **three-layer architecture**:

```
  HTTP Request
       |
       v
+------------------+
|   Controllers    |  Fiber handlers, request parsing,
|                  |  Swagger annotations, response formatting
+------------------+
       |
       v
+------------------+
|    Services      |  Business logic, GORM queries,
|                  |  transaction management, outbox
+------------------+
       |
       v
+------------------+
|    Structures    |  Go structs (Form, Model, Edit,
|                  |  Filter, Identity, Page variants)
+------------------+
       |
       v
+------------------+
|   PostgreSQL     |  Tables, ENUMs, Views, RPC Functions
+------------------+
```

---

# Package Dependency Graph

```
                    app/
                   /    \
            config.go   run.go
                 |        |
                 v        v
           generator/   parser/
           context.go   creator.go
                          |
              +-----------+-----------+
              |           |           |
              v           v           v
          struct_     service_    controller_
          creator     creator     creator
              |           |           |
              +-----------+-----------+
                          |
              +-----------+-----------+
              |           |           |
              v           v           v
          models/     helpers/    templates/
          (data)     (utils)     (.go.tmpl)
```

All creators share the same `GeneratorContext` instance -- no cross-creator coupling.

---

# Module Boundary Design

### Generator-Time Packages (Build Dependencies)

| Package | Responsibility |
|---------|---------------|
| `app/` | Orchestration, config loading, project management |
| `parser/` | Schema introspection, template execution, file output |
| `generator/` | GeneratorContext definition |
| `models/` | Table, Column, Constraint, RPC data structures |
| `helpers/` | Type conversion, struct tags, file I/O |
| `templates/` | Go template files (.go.tmpl) |

### Runtime Packages (Used by Generated Code)

| Package | Responsibility |
|---------|---------------|
| `pagination/` | Generic query engine with filtering, sorting, caching |
| `sse/` | Server-Sent Events hub with backplane support |
| `packages/events/` | Route-scoped event manager with sync/async dispatch |

---

# Data Flow: Schema to API

```
PostgreSQL               DataBridge V2              Generated API
+-------------+         +----------------+         +------------------+
| Schema      | ------> | Introspection  | ------> | Go Structs       |
| Tables      |   SQL   | (info_schema)  |  tmpl   | (6 variants)     |
| Columns     |         +----------------+         +------------------+
| Constraints |                |                          |
| FKs         |                v                          v
| ENUMs       |         +----------------+         +------------------+
| Views       |         | Type Mapping   |         | GORM Services    |
| Functions   |         | Tag Generation |         | (CRUD + Bulk)    |
+-------------+         +----------------+         +------------------+
                               |                          |
                               v                          v
                        +----------------+         +------------------+
                        | Template       |         | Fiber Controllers|
                        | Execution      |         | (Swagger + Auth) |
                        +----------------+         +------------------+
                               |                          |
                               v                          v
                        +----------------+         +------------------+
                        | File Writer    |         | Route Tree       |
                        | (Buffered I/O) |         | (Recursive Group)|
                        +----------------+         +------------------+
```

---

# Multi-Schema Support

DataBridge V2 processes multiple PostgreSQL schemas in a single run.
Each schema produces an isolated Go package with its own routes.

```
config.yaml
  projects:
    - name: "my-project"
      schemas:
        - name: "iam"         ->  api_v2/iam/
        - name: "catalog"     ->  api_v2/catalog/
        - name: "orders"      ->  api_v2/orders/
        - name: "logistics"   ->  api_v2/logistics/
        - name: "analytics"   ->  api_v2/analytics/
```

Cross-schema FK references are resolved via a shared `shared_types.go` file
generated by querying `information_schema.referential_constraints`.

The main `api.go` aggregation file wires all schema routers into a single
Fiber application with a unified route prefix.
