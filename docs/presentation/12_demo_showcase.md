
# Demo Showcase

## Innovation Hub -- Complete Generated Project

A fully generated project demonstrating all DataBridge V2 capabilities
across five PostgreSQL schemas with 54 tables.

### Schema Breakdown

| Schema | Tables | ENUMs | RPC Functions | Purpose |
|--------|--------|-------|--------------|---------|
| IAM | 11 | 3 | 2 | Identity and Access Management |
| Catalog | 11 | 3 | 3 | Product Catalog and Inventory |
| Orders | 9 | 2 | 3 | Order Processing Pipeline |
| Logistics | 11 | 3 | 3 | Warehouse and Shipment Tracking |
| Analytics | 12 | 2 | 3 | Business Intelligence and Reporting |

**Total: 54 tables, 13 ENUMs, 14 RPC functions, 2 views, 2 materialized views**

---

# Generated Output

### By the Numbers

| Artifact | Count |
|----------|-------|
| Go source files | 312 |
| Documentation pages | 121 |
| OpenAPI paths | 214 |
| OpenAPI schemas | 359 |
| MCP graph nodes | 628 |
| MCP graph edges | 723 |
| MCP doc chunks | 235 |
| Training data entries | ~500 |

### Generated Directory Structure

```
api_v2/
  api.go              Main aggregation (wires all schemas)
  shared_types.go     Cross-schema FK identity structs
  iam/                IAM schema package
  catalog/            Catalog schema package
  orders/             Orders schema package
  logistics/          Logistics schema package
  analytics/          Analytics schema package
```

---

# Per-Schema Package Structure

Each schema follows an identical package layout:

```
iam/
  iam.go                     Route registration (Run function)
  structures/
    iam_enums.go             Typed enum constants
    users.go                 User structs (Form, Model, Edit, ...)
    roles.go                 Role structs
    ... (11 files)
  services/
    users_service.go         CRUD + bulk operations
    roles_service.go
    ... (11 files)
  controllers/
    users_controller.go      Fiber handlers + Swagger
    roles_controller.go
    ... (11 files)
  tests/
    users_test.go            Unit tests with mocks
    ...
  testhelper/
    test_helpers.go          Shared test infrastructure
  pkg/iam_events/
    event_handlers.go        Event handler stubs
  outbox/
    users_outbox.go          Outbox event definitions
```

---

# IAM Schema Highlights

**Core Tables:**

| Table | Features Demonstrated |
|-------|---------------------|
| users | UUID PK, ENUM status, FK to roles, unique email |
| roles | Hierarchical (parent_id self-ref), slug unique |
| permissions | Composite unique (resource + action) |
| role_permissions | Junction table, composite FK |
| user_sessions | TTL-based, JSON metadata |
| audit_logs | Append-only, JSONB payload |
| api_keys | Hashed secret, expiration |
| user_preferences | Key-value JSONB store |
| teams | Hierarchical with avatar URL |
| team_members | Junction with role ENUM |
| invitations | Expiration + ENUM status |

---

# Catalog Schema Highlights

**Advanced Data Types:**

| Table | Notable Features |
|-------|-----------------|
| categories | Self-referential hierarchy (parent_id) |
| products | JSONB metadata, text[] tags, vector embedding |
| product_variants | Multi-attribute variants (size, color, material) |
| product_images | S3 URL references, sort ordering |
| price_lists | Multi-currency, date-range validity |
| inventory | Warehouse-product junction, reorder tracking |
| suppliers | JSONB contact info, rating system |
| brands | Slug-based routing, logo URL |
| reviews | Star rating, verified purchase flag |
| wishlists | User-product junction, privacy setting |
| collections | Curated product groupings |

Demonstrates: JSONB queries, array fields, self-referential FKs,
multi-table joins, and soft-delete patterns.

---

# Orders Schema Highlights

**Transactional Patterns:**

| Table | Notable Features |
|-------|-----------------|
| orders | Status state machine, total calculation |
| order_items | Composite FK (order + variant), quantity/price |
| payments | Gateway integration, idempotency key |
| refunds | Parent payment reference, partial amounts |
| coupons | Usage limits, date validity, discount types |
| order_coupons | Junction with discount tracking |
| order_status_history | Append-only state transitions |
| shipping_addresses | Structured address with coordinates |
| order_notes | Internal/external visibility flag |

Demonstrates: Transactional outbox, nested creation (order + items),
state machine tracking, and monetary precision handling.

---

# Logistics and Analytics Highlights

### Logistics Schema
- **Warehouses**: PostGIS geometry for location
- **Shipments**: Multi-carrier tracking, ETA calculation
- **Routes**: Optimized delivery path with waypoints
- **Vehicles**: Fleet management with capacity constraints

Demonstrates: Geospatial queries (ST_Distance, ST_Within),
`relaxGrouping` for flat route tree, real-time SSE tracking.

### Analytics Schema
- **Dashboards**: Configurable layout with widget grid
- **Reports**: Scheduled generation with template system
- **KPI Definitions**: Formula-based metrics
- **Data Sources**: Multi-database connector config
- **SQL Views**: `daily_revenue_summary`, `product_performance`
- **Materialized Views**: Pre-aggregated analytics data

Demonstrates: Views and materialized views handling, aggregation
queries, JSONB configuration storage.

---

# RPC Function Examples

Stored procedures generated as typed API endpoints:

**PostgreSQL:**
```sql
CREATE OR REPLACE FUNCTION orders.rpc__calculate_order_total(
    p_order_id UUID
) RETURNS NUMERIC AS $$
    SELECT COALESCE(SUM(quantity * unit_price), 0)
    FROM orders.order_items
    WHERE order_id = p_order_id;
$$ LANGUAGE sql STABLE;
```

**Generated Go Service:**
```go
func (s *RPCService) CalculateOrderTotal(orderID types.URID) (float64, error) {
    var result float64
    err := s.db.Raw("SELECT orders.rpc__calculate_order_total(?)",
        orderID).Scan(&result).Error
    return result, err
}
```

**Generated Endpoint:**
```
POST /api/v1/orders/rpc/calculate-order-total
Body: {"order_id": "uuid-here"}
Response: {"result": 1299.99}
```

---

# Documentation Site

The generated VitePress documentation site includes:

1. **Landing Page** -- Project overview with schema feature cards
2. **Per-Table Pages** -- Column reference, API docs, ER diagrams
3. **Swagger UI** -- Interactive OpenAPI 3.0 explorer (vitepress-openapi)
4. **Schema READMEs** -- Architecture overview per schema
5. **API Reference** -- Aggregated endpoint listing
6. **Missing Comments Report** -- Columns lacking documentation

### Deployment

```yaml
# GitHub Actions workflow included
name: Deploy Documentation
on:
  push:
    branches: [main]

# Builds VitePress + deploys to GitHub Pages
# URL: https://{org}.github.io/{repo}/
```

---

# Main Application Entry Point

```go
func main() {
    // Database connection
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    // Event Manager
    em := events.NewEventManager()

    // Fiber HTTP server
    app := fiber.New(fiber.Config{
        JSONEncoder: sonic.Marshal,
        JSONDecoder: sonic.Unmarshal,
    })

    // Middleware
    app.Use(cors.New())
    app.Use(logger.New())
    app.Use(recover.New())

    // Health check
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })

    // Mount all generated APIs
    apiGroup := app.Group("/api/v1")
    api.Setup(apiGroup, db, em)

    // Start server
    app.Listen(":3000")
}
```
