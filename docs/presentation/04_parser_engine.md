
# Parser Engine

## Schema Introspection Pipeline

The parser is the heart of DataBridge V2. It connects to PostgreSQL,
reads the schema metadata, and drives the code generation pipeline.

### Core Components

| Component | File | Responsibility |
|-----------|------|---------------|
| Creator | `parser/creator.go` | Orchestrates per-schema generation |
| Table Loader | `parser/tables.go` | SQL introspection via information_schema |
| Struct Creator | `parser/struct_creator.go` | Go struct generation |
| Service Creator | `parser/service_creator.go` | CRUD service generation |
| Controller Creator | `parser/controller_creator.go` | HTTP handler generation |
| Endpoint Creator | `parser/endpoint_creator.go` | Route registration |
| Test Creator | `parser/test_creator.go` | Test file generation |
| RPC Creator | `parser/rpc_creator.go` | Stored procedure wrappers |
| Event Creator | `parser/event_setup_creator.go` | Event handler stubs |

---

# Schema Introspection

The `GetTables()` function executes a sophisticated multi-CTE SQL query
against PostgreSQL `information_schema`:

```sql
WITH table_columns AS (
    SELECT c.table_name, c.column_name, c.data_type,
           c.is_nullable, c.udt_name, c.column_default,
           pgd.description AS comment
    FROM information_schema.columns c
    LEFT JOIN pg_catalog.pg_description pgd ...
    WHERE c.table_schema = $1
),
table_constraints AS (
    SELECT tc.constraint_type, tc.constraint_name,
           kcu.column_name, tc.table_name
    FROM information_schema.table_constraints tc
    JOIN information_schema.key_column_usage kcu ...
),
foreign_keys AS (
    SELECT kcu.column_name, ccu.table_name AS foreign_table,
           ccu.column_name AS foreign_column
    FROM information_schema.referential_constraints rc ...
)
SELECT ... FROM table_columns
JOIN table_constraints ...
LEFT JOIN foreign_keys ...
```

---

# Data Model: Table and Column

Every introspected table is represented by rich data structures:

```go
type Table struct {
    Name         string
    Schema       string
    Type         string           // "BASE TABLE", "VIEW"
    Columns      TableColumns     // []TableColumn
    PrimaryKeys  TableConstraints // composite PK support
    ForeignKeys  TableConstraints
    UniqueKeys   TableConstraints
}

type TableColumn struct {
    Name          string
    Type          string           // SQL data type
    Nullable      string           // "YES" / "NO"
    UdtName       string           // PostgreSQL user-defined type
    Default       *string          // column default expression
    Comment       *string          // pg_description
    Constraints   TableConstraints
    RelatedTables TableRelateds    // FK targets
}
```

Slice types implement `database/sql.Scanner` for GORM JSONB compatibility.

---

# RPC Function Discovery

DataBridge V2 automatically discovers PostgreSQL stored procedures
that follow the `rpc__` naming convention:

```sql
SELECT p.proname, n.nspname,
       pg_get_function_arguments(p.oid),
       pg_get_function_result(p.oid)
FROM pg_proc p
JOIN pg_namespace n ON p.pronamespace = n.oid
WHERE n.nspname = $1
  AND p.proname LIKE 'rpc__%'
```

Each discovered function becomes:
- A typed Go struct for parameters
- A service method calling `db.Raw()`
- A Fiber controller with Swagger annotations
- A registered endpoint on the schema router

---

# Generation Pipeline

```
Schema Name
     |
     v
+------------------+
| GenerateEnumTypes|  Query pg_enum -> enums.go
+------------------+
     |
     v
+------------------+
| GetTables()      |  Query information_schema -> []Table
+------------------+
     |
     v
+------------------+
| GetRPCFunctions()|  Query pg_proc -> []RPCFunction
+------------------+
     |
     v  (for each table, optionally parallel)
+------------------+     +------------------+
| StructCreator    | --> | structures/      |
+------------------+     +------------------+
| ServiceCreator   | --> | services/        |
+------------------+     +------------------+
| ControllerCreator| --> | controllers/     |
+------------------+     +------------------+
| TestCreator      | --> | tests/           |
+------------------+     +------------------+
     |
     v
+------------------+     +------------------+
| EndPointCreator  | --> | {schema}.go      |
+------------------+     +------------------+
| EventSetupCreator| --> | event_handlers/  |
+------------------+     +------------------+
| RPCCreator       | --> | rpc_service/ctrl |
+------------------+     +------------------+
```

---

# Enum Type Generation

PostgreSQL ENUMs are introspected and mapped to typed Go constants:

**Input (PostgreSQL):**
```sql
CREATE TYPE iam.user_status AS ENUM (
    'active', 'inactive', 'suspended', 'pending'
);
```

**Output (Go):**
```go
type UserStatus string

const (
    UserStatusActive    UserStatus = "active"
    UserStatusInactive  UserStatus = "inactive"
    UserStatusSuspended UserStatus = "suspended"
    UserStatusPending   UserStatus = "pending"
)
```

All enum values are type-safe. The generator validates that column
references to enum types produce the correct Go type in struct fields.

---

# Recursive Route Tree Grouping

The `groupTablesIntoRoutes()` function builds a hierarchical URL tree
from PostgreSQL table naming conventions:

**Input Tables:**
```
users, user_settings, user_preferences, user_notifications,
products, product_variants, product_variant_prices
```

**Generated Route Tree:**
```
/users
  GET    /          -> Search
  POST   /          -> Create
  GET    /:id       -> Find
  PUT    /:id       -> Update
  DELETE /:id       -> Delete
  /settings
    GET  /          -> Search user_settings
    ...
  /preferences
    GET  /          -> Search user_preferences
    ...
  /notifications
    GET  /          -> Search user_notifications
    ...
/products
  /variants
    /prices         -> product_variant_prices
```

The `relaxGrouping` flag flattens this tree for schemas where
deep nesting is not appropriate.

---

# Creator Interface Pattern

All creators follow the same structural pattern:

```go
type StructCreator struct {
    DB      *gorm.DB
    Context *generator.GeneratorContext
    // creator-specific fields
}

func (sc *StructCreator) CreateTableStruct(
    table models.Table,
    schemaName string,
) error {
    // 1. Prepare template data
    data := map[string]interface{}{
        "Table":     table,
        "Schema":    schemaName,
        "Context":   sc.Context,
        "Imports":   sc.resolveImports(table),
    }

    // 2. Execute template
    output := templateCache.Execute("struct.go.tmpl", data)

    // 3. Write to disk
    return helpers.WriteFileToDisk(path, output)
}
```

This consistency makes the system easy to extend:
adding a new output type means adding a new creator + template.
