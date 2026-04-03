---
title: PriceHistory
---

# PriceHistory

**Table:** `catalog.price_history`

**Base path:** `/price-history`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `products` | `product_id` | `price_history_product_id_fkey` | [Products](./products) |
| `users` | `changed_by` | `price_history_changed_by_fkey` | [Users](./users) |


## Entity Relationship Diagram

```mermaid
erDiagram
    PriceHistory }o--|| Products : "FK"
    PriceHistory }o--|| Users : "FK"
```

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `products` |
| 4 | `old_price` | `numeric` | `float64` | `number` | NO | - | - | - |
| 5 | `new_price` | `numeric` | `float64` | `number` | NO | - | - | - |
| 6 | `changed_by` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `users` |
| 7 | `reason` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `product_id` | `products` | `price_history_product_id_fkey` |
| `changed_by` | `users` | `price_history_changed_by_fkey` |


## Go Generated Code

> 📂 Source: [📄 `PriceHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/PriceHistory.go) · [📄 `PriceHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/PriceHistory.go) · [📄 `PriceHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/PriceHistory.go)

### Structs

::::tabs

:::tab Form

#### PriceHistoryForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/PriceHistory.go#:~:text=type%20PriceHistoryForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `OldPrice` | `float64` | `oldPrice` | NO |
| `NewPrice` | `float64` | `newPrice` | NO |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Reason` | `string` | `reason` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab Model

#### PriceHistory [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/PriceHistory.go#:~:text=type%20PriceHistory%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `OldPrice` | `float64` | `oldPrice` | NO |
| `NewPrice` | `float64` | `newPrice` | NO |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Reason` | `string` | `reason` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab Edit

#### PriceHistoryEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/PriceHistory.go#:~:text=type%20PriceHistoryEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `OldPrice` | `*float64` | `oldPrice` | YES |
| `NewPrice` | `*float64` | `newPrice` | YES |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Reason` | `*string` | `reason` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Filter

#### PriceHistoryFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/PriceHistory.go#:~:text=type%20PriceHistoryFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `OldPrice` | `*float64` | `oldPrice` | YES |
| `NewPrice` | `*float64` | `newPrice` | YES |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Reason` | `*string` | `reason` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Page

#### PriceHistoryPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/PriceHistory.go#:~:text=type%20PriceHistoryPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `OldPrice` | `float64` | `oldPrice` | NO |
| `NewPrice` | `float64` | `newPrice` | NO |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Reason` | `string` | `reason` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab BatchUpdate

#### PriceHistoryBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/PriceHistory.go#:~:text=type%20PriceHistoryBatchUpdate%20struct)

```go
type PriceHistoryBatchUpdate struct {
    Data       json.RawMessage `json:"data"`
    PathParams struct {
        Id uuid.UUID
    } `json:"pathParams"`
}
```

::::

### Service & Endpoints

::::tabs

:::tab Service Methods

| Method | Signature |
|---------|-----------|
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/PriceHistory.go#:~:text=)%20CreatePriceHistory() | `(PriceHistoryService) CreatePriceHistory(data PriceHistoryForm) (PriceHistoryForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/PriceHistory.go#:~:text=)%20CreatePriceHistoryMultiple() | `(PriceHistoryService) CreatePriceHistoryMultiple(data []PriceHistoryForm) ([]PriceHistoryForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/PriceHistory.go#:~:text=)%20UpdatePriceHistory() | `(PriceHistoryService) UpdatePriceHistory(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/PriceHistory.go#:~:text=)%20UpdatePriceHistoryMultiple() | `(PriceHistoryService) UpdatePriceHistoryMultiple(data []PriceHistoryBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/PriceHistory.go#:~:text=)%20DeletePriceHistory() | `(PriceHistoryService) DeletePriceHistory(id uuid.UUID) error` |

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/price-history/` | Search with query params |
| `GET` | `/price-history/pagination` | Paginated listing |
| `POST` | `/price-history/` | Create single record |
| `POST` | `/price-history/bulk/` | Create multiple records |
| `PUT` | `/price-history/bulk/` | Batch update |
| `GET` | `/price-history/with-id/:id` | Get by ID |
| `PUT` | `/price-history/with-id/:id` | Update by ID |
| `DELETE` | `/price-history/with-id/:id` | Delete by ID |

:::tab Query & Filters

| Parameter | Type | Description |
|-----------|------|-------------|
| `page` | `int` | Page number (default: 1) |
| `size` | `int` | Items per page (default: 10) |
| `sort` | `string` | Sort field. Prefix `-` for descending. Example: `-created_at` |
| `fields` | `string` | Comma-separated column list to select |
| `preloads` | `string` | Comma-separated relation names to preload |
| `filters` | `array` | Filter rules: `[[field, op, value], ...]` |
| `groupby` | `string` | Group by field name |
| `aggregations` | `json` | Aggregation specs: `[{func,field,alias}]` |

**Filter Operators:** `eq` `neq` `gt` `gte` `lt` `lte` `in` `notin` `like` `ilike` `is` `isnot` `between`

::::

### RPC Functions

| Function | Parameters | Return | Endpoint |
|----------|-----------|--------|----------|
| `avg_product_rating` | `p_product_id uuid` | `numeric` | `/rpc/avg_product_rating` |
| `count_active_products` | - | `integer` | `/rpc/count_active_products` |
| `products_by_category` | `p_category_id uuid` | `integer` | `/rpc/products_by_category` |


:::tab Frontend

## TypeScript Types & Hooks

::::tabs

:::tab Interfaces

```typescript
export interface PriceHistory {
  id: string;
  name: string;
  productId: string;
  oldPrice: number;
  newPrice: number;
  changedBy?: string;
  reason: string;
  createdAt: string;
}

export interface PriceHistoryForm {
  name: string;
  productId: string;
  oldPrice: number;
  newPrice: number;
  changedBy?: string;
  reason: string;
  createdAt: string;
}

export interface PriceHistoryEdit {
  id: string;
  name: string;
  productId: string;
  oldPrice: number;
  newPrice: number;
  changedBy?: string;
  reason: string;
  createdAt: string;
}

export interface PriceHistoryPage {
  data: PriceHistory[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type PriceHistoryPathQuery = {
  page?: number;
  size?: number;
  sort?: string;
  fields?: string;
  preloads?: string;
  filters?: string;
};

```

:::tab React Query

```typescript
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

const PriceHistoryKeys = {
  all: ["price_history"] as const,
  lists: () => [...PriceHistoryKeys.all, "list"] as const,
  detail: (id: any) => [...PriceHistoryKeys.all, "detail", id] as const,
} as const;

export function usePriceHistoryList(query?: PriceHistoryPathQuery) {
  return useQuery({
    queryKey: [...PriceHistoryKeys.lists(), query],
    queryFn: () => fetch(`/price-history/pagination`, { method: "GET" }).then(r => r.json()) as Promise<PriceHistoryPage>,
  });
}

export function usePriceHistoryDetail(id: any) {
  return useQuery({
    queryKey: PriceHistoryKeys.detail(id),
    queryFn: () => fetch(`/price-history/with-id/:id`).then(r => r.json()) as Promise<PriceHistory>,
  });
}

export function useCreatePriceHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: PriceHistoryForm) =>
      fetch("/price-history/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: PriceHistoryKeys.lists() }),
  });
}

export function useUpdatePriceHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: PriceHistoryEdit }) =>
      fetch(`/price-history/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: PriceHistoryKeys.all }),
  });
}

export function useDeletePriceHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/price-history/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: PriceHistoryKeys.all }),
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

export const PriceHistoryFormSchema = z.object({
  name: z.string(),
  productId: z.string().uuid(),
  oldPrice: z.number(),
  newPrice: z.number(),
  changedBy: z.string().uuid().optional(),
  reason: z.string(),
  createdAt: z.string().datetime(),
});

export type PriceHistoryFormInput = z.infer<typeof PriceHistoryFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './price_history.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search PriceHistory

```
GET /api/v1/price-history/
```

> Retrieve list filtered by query parameters.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `size` | `integer` | No | Max results (default: 10) |
| `sort` | `string` | No | Sort field. Prefix `-` for DESC. e.g. `-created_at` |
| `fields` | `string` | No | Comma-separated columns to select |
| `preloads` | `string` | No | Available: ProductIdDetail, ProductIdDetail.ProductVariantsList, ProductIdDetail.ProductVariantsList.ProductIdDetail, ProductIdDetail.ProductMediaList, ProductIdDetail.ProductMediaList.ProductIdDetail, ProductIdDetail.ProductReviewsList, ProductIdDetail.ProductReviewsList.ProductIdDetail, ProductIdDetail.CollectionProductsList, ProductIdDetail.CollectionProductsList.CollectionIdDetail, ProductIdDetail.CollectionProductsList.ProductIdDetail, ProductIdDetail.ProductTagsList, ProductIdDetail.ProductTagsList.ProductIdDetail, ProductIdDetail.ProductTagsList.TagIdDetail, ProductIdDetail.PriceHistoryList, ProductIdDetail.PriceHistoryList.ProductIdDetail, ProductIdDetail.BrandIdDetail, ProductIdDetail.BrandIdDetail.ProductsList, ProductIdDetail.CategoryIdDetail, ProductIdDetail.CategoryIdDetail.CategoriesList, ProductIdDetail.CategoryIdDetail.ProductsList, ProductIdDetail.CategoryIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Products, Products.Brands, Products.Brands.Organizations, Products.Categories, Products.Categories.Categories, Products.Users, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `oldPrice` | `number` | No | Filter by old_price |
| `newPrice` | `number` | No | Filter by new_price |
| `changedBy` | `string (uuid)` | No | Filter by changed_by |
| `reason` | `string` | No | Filter by reason |

**Response:** `PriceHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/price-history/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search PriceHistory (POST)

```
POST /api/v1/price-history/search
```

> Search with body filters. Auto-used when query string > 2KB.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:**

```typescript
{
  size?: number  // e.g. 10
  sort?: string[]  // e.g. ["-createdAt"]
  filters?: FilterRule[]  // e.g. [["name", "eq", "value"]]
  fields?: string[]
  preloads?: string[]
}
```

**Response:** `PriceHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/price-history/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate PriceHistory

```
GET /api/v1/price-history/pagination
```

> Paginated listing.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `page` | `integer` | No | Page number (default: 1) |
| `size` | `integer` | No | Max results (default: 10) |
| `sort` | `string` | No | Sort field. Prefix `-` for DESC. e.g. `-created_at` |
| `fields` | `string` | No | Comma-separated columns to select |
| `preloads` | `string` | No | Available: ProductIdDetail, ProductIdDetail.ProductVariantsList, ProductIdDetail.ProductVariantsList.ProductIdDetail, ProductIdDetail.ProductMediaList, ProductIdDetail.ProductMediaList.ProductIdDetail, ProductIdDetail.ProductReviewsList, ProductIdDetail.ProductReviewsList.ProductIdDetail, ProductIdDetail.CollectionProductsList, ProductIdDetail.CollectionProductsList.CollectionIdDetail, ProductIdDetail.CollectionProductsList.ProductIdDetail, ProductIdDetail.ProductTagsList, ProductIdDetail.ProductTagsList.ProductIdDetail, ProductIdDetail.ProductTagsList.TagIdDetail, ProductIdDetail.PriceHistoryList, ProductIdDetail.PriceHistoryList.ProductIdDetail, ProductIdDetail.BrandIdDetail, ProductIdDetail.BrandIdDetail.ProductsList, ProductIdDetail.CategoryIdDetail, ProductIdDetail.CategoryIdDetail.CategoriesList, ProductIdDetail.CategoryIdDetail.ProductsList, ProductIdDetail.CategoryIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Products, Products.Brands, Products.Brands.Organizations, Products.Categories, Products.Categories.Categories, Products.Users, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `oldPrice` | `number` | No | Filter by old_price |
| `newPrice` | `number` | No | Filter by new_price |
| `changedBy` | `string (uuid)` | No | Filter by changed_by |
| `reason` | `string` | No | Filter by reason |

**Response:** `PaginationResponse<PriceHistory>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/price-history/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate PriceHistory (POST)

```
POST /api/v1/price-history/pagination
```

> Paginated listing with body filters.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:**

```typescript
{
  page?: number  // e.g. 1
  size?: number  // e.g. 10
  sort?: string[]  // e.g. ["-createdAt"]
  filters?: FilterRule[]  // e.g. [["name", "eq", "value"]]
  fields?: string[]
  preloads?: string[]
}
```

**Response:** `PaginationResponse<PriceHistory>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/price-history/pagination"
```

</details>

---

:::tab Create

#### <Badge type="tip" text="POST" /> Create PriceHistory

```
POST /api/v1/price-history/
```

> Create a new record.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:**

```typescript
{
  name?: string  // e.g. example_name
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  oldPrice: number  // e.g. 99.99
  newPrice: number  // e.g. 99.99
  changedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  reason?: string  // e.g. example_reason
}
```

**Response:** `PriceHistory`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/price-history/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create PriceHistory

```
POST /api/v1/price-history/bulk/
```

> Create multiple records in one request.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:**

```typescript
{
  name?: string  // e.g. example_name
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  oldPrice: number  // e.g. 99.99
  newPrice: number  // e.g. 99.99
  changedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  reason?: string  // e.g. example_reason
}
```

**Response:** `PriceHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/price-history/bulk/"
```

</details>

---

:::tab Find & Update

#### <Badge type="info" text="GET" /> Find PriceHistory by ID

```
GET /api/v1/price-history/with-id/:id
```

> Retrieve a single record by primary key.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `Id` | `string (uuid)` | Yes | Primary key (uuid) |

**Response:** `PriceHistory`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/price-history/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update PriceHistory

```
PUT /api/v1/price-history/with-id/:id
```

> Partial update — send only the fields to change.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `Id` | `string (uuid)` | Yes | Primary key (uuid) |

**Request Body:**

```typescript
{
  name?: string
  productId?: string
  oldPrice?: number
  newPrice?: number
  changedBy?: string
  reason?: string
}
```

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/price-history/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update PriceHistory

```
PUT /api/v1/price-history/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: PriceHistoryEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/price-history/bulk/"
```

</details>

---

:::tab Delete

#### <Badge type="danger" text="DELETE" /> Delete PriceHistory

```
DELETE /api/v1/price-history/with-id/:id
```

> Soft-delete (sets deleted_at + deleted_by).

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `Id` | `string (uuid)` | Yes | Primary key (uuid) |

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X DELETE \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/price-history/with-id/:id"
```

</details>

---

::::


::::
