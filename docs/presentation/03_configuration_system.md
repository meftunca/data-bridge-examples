
# Configuration System

## YAML-Driven Code Generation

All generation behavior is controlled through a single `config.yaml` file.
No code changes required to adjust output -- only configuration.

```yaml
defaultHelperLibBasePath: "github.com/maple-tech/baseline"
defaultPaginationPackagePath: "backend-generator/apiv2/pagination"
defaultImportRepo: "github.com/maple-tech"
defaultGenerateTests: true
defaultRBACEnabled: false
defaultSSEEnabled: true
defaultOutboxEnabled: true
defaultEventManagerIntegration: true
```

---

# Project Configuration

Each project defines its output target, module identity, and schema list.

```yaml
projects:
  - name: "innovation-hub"
    projectPath: "github.com/maple-tech/backend-generator/apiv2/data-bridge-examples"
    moduleName: "data-bridge-examples"
    outputFolder: "api_v2"
    generateTests: true
    sseEnabled: true
    outboxEnabled: true
    eventManagerIntegration: true
    schemas:
      - name: "iam"
      - name: "catalog"
      - name: "orders"
      - name: "logistics"
        relaxGrouping: true
      - name: "analytics"
```

---

# Schema-Level Configuration

Fine-grained control per schema:

| Option | Type | Purpose |
|--------|------|---------|
| `name` | string | PostgreSQL schema name |
| `skipSchemaPrefixInOps` | bool | Remove schema prefix from operation names |
| `ignoreTables` | []string | Exclude specific tables from generation |
| `relaxGrouping` | bool | Flatten route tree instead of deep nesting |
| `eventHandlers.generate` | bool | Generate event handler stubs |
| `eventHandlers.outputDir` | string | Custom output directory for handlers |
| `eventHandlers.packageName` | string | Go package name for handlers |

---

# Feature Flags

The configuration supports granular feature toggles at both
the global level and per-project override level.

```
                Global Defaults
                      |
          +-----------+-----------+
          |           |           |
       Project A   Project B   Project C
       (override)  (default)   (override)
```

| Flag | Effect When Enabled |
|------|-------------------|
| `generateTests` | Produces unit tests, integration tests, test helpers |
| `rbacEnabled` | Adds permission middleware to route registration |
| `sseEnabled` | Generates SSE stream endpoints and hub initialization |
| `outboxEnabled` | Adds transactional outbox pattern to service creates/updates |
| `eventManagerIntegration` | Wires EventManager hooks into controller lifecycle |

Project-level flags use pointer types (`*bool`) to distinguish
"not set" from "explicitly false", enabling clean inheritance.

---

# GeneratorContext

All configuration resolves into a single immutable struct
passed through the entire generation pipeline.

```go
type GeneratorContext struct {
    ProjectName           string
    ProjectPath           string
    OutputFolder          string
    ProjectImportPath     string
    HelperLibBasePath     string
    PaginationPackagePath string
    ImportRepoPath        string
    SSEPackagePath        string

    GenerateTests  bool
    RBACEnabled    bool
    SSEEnabled     bool
    OutboxEnabled  bool
}
```

**No global state.** Every creator, helper, and template function
receives this context explicitly, enabling safe concurrent execution.

---

# Configuration Validation

The system validates configuration at startup before any generation begins:

1. **Project name** must be non-empty
2. **Output folder** must be specified
3. **At least one schema** must be defined per project
4. **Database connection** (DSN) must be valid and reachable
5. **Import paths** must be resolvable within the Go module system

Invalid configuration produces clear error messages with the specific
field and project that failed validation -- not cryptic template errors
downstream.

---

# Configuration to Output Mapping

```
config.yaml                      Generated Output
+--------------------------+     +-----------------------------+
| moduleName: "my-app"     | --> | go.mod: module my-app       |
| outputFolder: "api_v2"   | --> | api_v2/                     |
| schemas:                  |     |   iam/                      |
|   - name: "iam"          | --> |     structures/              |
|     ignoreTables:         |     |     services/                |
|       - "migrations"      |     |     controllers/             |
|   - name: "catalog"      | --> |   catalog/                   |
|     relaxGrouping: true   |     |     structures/              |
|     skipSchemaPrefixInOps | --> |     (flat route tree)        |
| sseEnabled: true          | --> |   sse/ stream endpoints      |
| outboxEnabled: true       | --> |   outbox/ event files        |
| generateTests: true       | --> |   tests/ + test_helpers/     |
+--------------------------+     +-----------------------------+
```

Every configuration option has a direct, traceable impact on the
generated output. No hidden behavior or implicit conventions.
