
# Pagination Engine

## A Full-Featured Generic Query Engine

The pagination package is a standalone, generic query engine
built on top of GORM. It is the largest single component in
the system: 30+ files, 3000+ lines of code.

### Core Design

```go
result, err := pagination.NewPagination[Product]().
    With(db).
    Request(c).       // parse Fiber request
    Response()        // execute query, return Page[Product]
```

**Fluent API** -- Method chaining for clean, readable query setup.

**Generics** -- `Pagination[T any]` works with any GORM model.

**Zero Configuration** -- Sensible defaults; all options via query params.

---

# Request Parsing

The engine parses query parameters into a structured request:

```
GET /products/search
  ?page=1
  &size=20
  &sort=price:desc,name:asc
  &fields=id,name,price,category_id
  &preloads=Category,Brand
  &filters[0][field]=price&filters[0][op]=gte&filters[0][value]=100
  &filters[1][field]=status&filters[1][op]=eq&filters[1][value]=active
  &groupBy=category_id
  &aggregations[0][function]=avg&aggregations[0][field]=price
```

**Parsed into:**
```go
type ParsedRequest struct {
    Page         int
    Size         int
    Sort         []SortOrder
    Fields       []string
    Preloads     []PreloadParam
    Filters      []PageFilter
    GroupBy      []string
    Aggregations []AggregationFunction
}
```

---

# Filter System Architecture

The filter system uses the **Strategy Pattern** with three parser modes:

```
                   FilterParserComponent
                   /         |         \
        SimpleFilter    MapFilter    RecursiveFilter
        (array fmt)    (object fmt)  (nested AND/OR)
```

**Simple Format:**
```
filters[0][field]=price&filters[0][op]=gte&filters[0][value]=100
```

**Map Format:**
```json
{"price": {"gte": 100}, "status": {"eq": "active"}}
```

**Recursive Nested Format:**
```json
{
  "AND": [
    {"field": "price", "op": "gte", "value": 100},
    {"OR": [
      {"field": "status", "op": "eq", "value": "active"},
      {"field": "status", "op": "eq", "value": "pending"}
    ]}
  ]
}
```

---

# Where Clause Builder

The `WhereClauseBuilder` maintains a registry of operator handlers.
Each operator knows how to translate a filter into a GORM `Where` clause.

### Standard SQL Operators

| Operator | SQL | Example |
|----------|-----|---------|
| `eq` | `= ?` | `price eq 100` |
| `neq` | `<> ?` | `status neq deleted` |
| `gt` / `gte` | `> ?` / `>= ?` | `price gte 50` |
| `lt` / `lte` | `< ?` / `<= ?` | `price lte 500` |
| `in` | `IN (?)` | `status in active,pending` |
| `not_in` | `NOT IN (?)` | `role not_in admin,super` |
| `is_null` | `IS NULL` | `deleted_at is_null` |
| `is_not_null` | `IS NOT NULL` | `email is_not_null` |
| `between` | `BETWEEN ? AND ?` | `price between 10,100` |
| `like` | `LIKE ?` | `name like %john%` |

---

# PostgreSQL-Specific Operators

Beyond standard SQL, the engine supports PostgreSQL-native operations:

### Text Search
| Operator | SQL | Use Case |
|----------|-----|----------|
| `ilike` | `ILIKE ?` | Case-insensitive pattern match |
| `fts` | `to_tsvector(?) @@ to_tsquery(?)` | Full-text search |
| `fts_phrase` | `to_tsvector(?) @@ phraseto_tsquery(?)` | Phrase search |

Full-text search supports 16+ languages:
`english`, `turkish`, `german`, `french`, `spanish`, `arabic`, `chinese`, ...

### JSONB Operators
| Operator | SQL | Use Case |
|----------|-----|----------|
| `jsonb_contains` | `@> ?::jsonb` | JSONB containment |
| `jsonb_has_key` | `? ?` | JSONB key existence |
| `jsonb_path` | `@? ?::jsonpath` | JSONPath query |

---

# Vector Search (pgvector)

Native support for AI/ML vector similarity search:

| Operator | Distance Metric | SQL |
|----------|----------------|-----|
| `cosine_similarity` | Cosine | `1 - (col <=> ?)` |
| `l2_distance` | Euclidean | `col <-> ?` |
| `inner_product` | Dot Product | `col <#> ?` |
| `knn` | K-Nearest Neighbors | `ORDER BY col <-> ? LIMIT k` |

**Usage Example:**
```
GET /products/search
  ?filters[0][field]=embedding
  &filters[0][op]=cosine_similarity
  &filters[0][value]=[0.1,0.2,0.3,...,0.768]
  &filters[0][threshold]=0.8
```

This enables semantic search, recommendation engines, and
similarity-based retrieval directly through the pagination API.

---

# Geospatial Search (PostGIS)

Built-in support for geographic queries:

| Operator | PostGIS Function | Use Case |
|----------|-----------------|----------|
| `st_distance` | `ST_Distance()` | Distance calculation |
| `st_within` | `ST_DWithin()` | Radius search |
| `st_intersects` | `ST_Intersects()` | Area overlap |
| `st_contains` | `ST_Contains()` | Containment check |

**Radius Search Example:**
```
GET /warehouses/search
  ?filters[0][field]=location
  &filters[0][op]=st_within
  &filters[0][value]=POINT(29.0 41.0)
  &filters[0][distance]=5000
```

The `geometry_formatter.go` handles WKT (Well-Known Text) parsing
for Point, LineString, Polygon, and MultiPolygon types.

---

# Aggregation Functions

22+ aggregation functions for analytical queries:

| Category | Functions |
|----------|----------|
| Basic | `COUNT`, `SUM`, `AVG`, `MIN`, `MAX` |
| Statistical | `STDDEV`, `VARIANCE`, `MEDIAN`, `MODE` |
| Percentile | `PERCENTILE_CONT`, `PERCENTILE_DISC` |
| Array | `ARRAY_AGG`, `STRING_AGG` |
| Boolean | `BOOL_AND`, `BOOL_OR`, `EVERY` |
| JSON | `JSON_AGG`, `JSONB_AGG`, `JSON_OBJECT_AGG` |
| Regression | `CORR`, `COVAR_POP`, `COVAR_SAMP`, `REGR_SLOPE` |
| Window | `ROW_NUMBER`, `RANK`, `DENSE_RANK`, `NTILE` |

**Example:**
```
GET /orders/search
  ?groupBy=status
  &aggregations[0][function]=count&aggregations[0][field]=id
  &aggregations[1][function]=sum&aggregations[1][field]=total_amount
  &aggregations[2][function]=avg&aggregations[2][field]=total_amount
```

---

# Preload System

Relations are preloaded with per-relation filtering and sorting:

```
GET /orders/search
  ?preloads=Customer,Items.Product
  &preloads[Customer][fields]=id,name,email
  &preloads[Items][sort]=quantity:desc
  &preloads[Items][size]=5
  &preloads[Items.Product][fields]=id,name,price
```

The engine reflects on GORM struct tags to discover associations
and builds the preload chain automatically.

**Schema-Based Validation** prevents invalid field names:
- `strict` mode: reject unknown fields (400 error)
- `compatible` mode: silently ignore unknown fields
- `ignore` mode: pass through all fields

---

# Response Format

```go
type Page[T any] struct {
    Items      []T   `json:"items"`
    Total      int64 `json:"total"`
    Page       int   `json:"page"`
    Size       int   `json:"size"`
    TotalPages int   `json:"total_pages"`
    HasNext    bool  `json:"has_next"`
    HasPrev    bool  `json:"has_prev"`
}
```

**Collection Mode** groups results by a field:
```json
{
  "collections": {
    "active": { "items": [...], "total": 42 },
    "pending": { "items": [...], "total": 7 },
    "inactive": { "items": [...], "total": 3 }
  }
}
```

---

# Caching Layer

Built-in LRU cache with TTL for repeated queries:

```go
type QueryCache interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{}, ttl time.Duration)
    Invalidate(prefix string)
}
```

**Deterministic Cache Keys** are generated from the full query signature:
- Model type name
- Filter conditions (sorted)
- Sort order
- Page/size
- Selected fields
- Preload configuration

**Model-Scoped Invalidation**: When a record is created/updated/deleted,
all cache entries for that model type are invalidated via prefix matching.
