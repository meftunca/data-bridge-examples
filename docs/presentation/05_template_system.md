
# Template System

## Go Templates with Custom FuncMap

DataBridge V2 uses Go `text/template` with 40+ custom functions
to generate idiomatic, production-ready Go code.

### Template Inventory (13 Files)

| Template | Output |
|----------|--------|
| `struct.go.tmpl` | Go structs (Form, Model, Edit, Filter, Identity, Page) |
| `service.go.tmpl` | Service interface + GORM implementation |
| `controller.go.tmpl` | Fiber HTTP handlers with Swagger |
| `endpoint.go.tmpl` | Route registration with grouping |
| `enums.go.tmpl` | Typed enum constants |
| `controller_test.go.tmpl` | Unit tests with mock service |
| `api_test.go.tmpl` | HTTP integration tests |
| `main_test.go.tmpl` | Test suite setup |
| `test_helpers.go.tmpl` | Shared test utilities |
| `event_handlers.go.tmpl` | Event handler stubs |
| `rpc_service.go.tmpl` | RPC service for stored procedures |
| `rpc_controller.go.tmpl` | RPC HTTP handlers |
| `swagger.go.tmpl` | OpenAPI annotation blocks |

---

# Template Cache Architecture

The `TemplateCache` is a performance-critical singleton that pre-loads
and caches all templates at startup.

```go
type TemplateCache struct {
    templates map[string]*template.Template
    mu        sync.RWMutex
}

var (
    globalCache *TemplateCache
    cacheOnce   sync.Once
)

func GetTemplateCache() *TemplateCache {
    cacheOnce.Do(func() {
        globalCache = &TemplateCache{
            templates: make(map[string]*template.Template),
        }
        globalCache.loadAll()
    })
    return globalCache
}
```

Key design decisions:
- **`sync.Once`** ensures thread-safe lazy initialization
- Templates are **cloned** before execution to avoid shared `define` block conflicts
- **Hot-reload** support for development: detects file changes and reloads

---

# Custom FuncMap -- Naming Conventions

Functions that transform SQL identifiers into Go conventions:

| Function | Input | Output |
|----------|-------|--------|
| `ToCamel` | `user_settings` | `UserSettings` |
| `ToLowerCamel` | `user_settings` | `userSettings` |
| `ToSnake` | `UserSettings` | `user_settings` |
| `Plural` | `category` | `categories` |
| `Singular` | `categories` | `category` |
| `SafeFieldName` | `type` | `TypeField` (avoids Go keywords) |
| `TrimSchemaPrefix` | `iam_users` | `users` (with schema context) |

These ensure consistent naming across all generated files without
manual intervention or naming configuration.

---

# Custom FuncMap -- Type System

Functions that map PostgreSQL types to Go types:

```go
// GoType returns the Go type string for a column
func GoType(col models.TableColumn) string {
    return helpers.SqlTypeToGoSupportedTypeStr(
        col.Type, col.UdtName, col.Nullable,
    )
}

// GoTypePointer returns pointer variant for nullable fields
func GoTypePointer(col models.TableColumn) string { ... }
```

**Type Mapping Examples:**

| PostgreSQL | Go Type | Nullable Go Type |
|-----------|---------|-----------------|
| `uuid` | `types.URID` | `*types.URID` |
| `integer` | `types.ID` | `*types.ID` |
| `text` | `string` | `*string` |
| `boolean` | `bool` | `*bool` |
| `jsonb` | `types.JSONB` | `*types.JSONB` |
| `timestamp` | `types.NullTime` | `*types.NullTime` |
| `geometry` | `types.Geometry` | `*types.Geometry` |
| `USER-DEFINED` | `{EnumType}` | `*{EnumType}` |

---

# Custom FuncMap -- GORM and Struct Tags

The `StructTags` function generates complete struct field tags:

```go
func StructTagsController(col TableColumn, table Table) string
```

**Output Example:**
```go
type User struct {
    ID        types.URID  `gorm:"primaryKey;type:uuid" json:"id"`
    Email     string      `gorm:"not null;uniqueIndex" json:"email"
                           example:"john@example.com"`
    Name      string      `gorm:"not null" json:"name"
                           example:"John Adams"`
    Status    UserStatus  `gorm:"type:user_status" json:"status"`
    CreatedAt types.NullTime `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt types.NullTime `gorm:"autoUpdateTime" json:"updated_at"`
}
```

Tags include:
- GORM directives: `primaryKey`, `not null`, `type:`, `autoCreateTime`
- JSON field names with `omitempty` for nullable fields
- Example values generated via `gofakeit` for Swagger documentation

---

# Custom FuncMap -- Query Helpers

Functions that generate GORM query logic from schema metadata:

| Function | Purpose |
|----------|---------|
| `PreloadFields` | Generates `db.Preload()` chains for FK relations |
| `GetKeyParams` | Extracts primary key parameters from Fiber context |
| `GetWhereClause` | Builds WHERE conditions for composite PKs |
| `ParsePathParams` | Generates `c.Params()` parsing for route params |
| `SetRequiredFields` | Lists non-nullable fields for validation |
| `Dict` | Creates template-local variable dictionaries |

**Preload Example:**
```go
// For a table with FK: order.customer_id -> customers.id
// and FK: order.product_id -> products.id

func (s *OrderService) Find(id types.URID) (*Order, error) {
    var result Order
    err := s.db.
        Preload("Customer").
        Preload("Product").
        Where("id = ?", id).
        First(&result).Error
    return &result, err
}
```

---

# Custom FuncMap -- Test Helpers

Functions that support test code generation:

| Function | Purpose |
|----------|---------|
| `MockServiceMethods` | Generates mock interface implementations |
| `TestCreatePayload` | Builds valid JSON payloads for create tests |
| `TestEditPayload` | Builds valid JSON payloads for update tests |
| `AssertFields` | Generates field-by-field assertion statements |
| `RandomValue` | Produces type-appropriate random test values |

These functions ensure generated tests are immediately runnable
without manual fixture setup.

---

# Template Composition

Templates use Go's `define` and `template` blocks for composition:

```
struct.go.tmpl
  |-- defines "struct_form"      (create input)
  |-- defines "struct_model"     (database model)
  |-- defines "struct_edit"      (update input)
  |-- defines "struct_filter"    (search filters)
  |-- defines "struct_identity"  (minimal reference)
  |-- defines "struct_page"      (paginated response)

controller.go.tmpl
  |-- defines "handler_create"
  |-- defines "handler_bulk_create"
  |-- defines "handler_update"
  |-- defines "handler_delete"
  |-- defines "handler_search"
  |-- defines "handler_paginate"
  |-- defines "handler_find"
  |-- includes "swagger.go.tmpl" (annotation blocks)
```

Each block is independently testable and can be overridden
without affecting other parts of the template.

---

# Template Execution Safety

Several mechanisms prevent template errors from producing invalid code:

1. **Clone Before Execute** -- Each template execution works on a clone,
   preventing concurrent `define` block contamination.

2. **Buffered Output** -- Templates write to `bytes.Buffer` first.
   Only after successful execution is the buffer flushed to disk.

3. **Post-Processing** -- Generated code passes through `goimports`
   to fix import ordering and remove unused imports.

4. **Type Safety in FuncMap** -- All template functions perform
   explicit type assertions with clear error messages on failure.

5. **Nil-Safe Access** -- Template functions handle nil pointers
   and empty slices gracefully, producing valid empty output
   rather than panicking.
