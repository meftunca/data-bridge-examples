---
title: CollectionProducts
---

# CollectionProducts

**Table:** `catalog.collection_products`

**Base path:** `/collection-products`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `collections` | `collection_id` | `collection_products_collection_id_fkey` | [Collections](./collections) |
| `products` | `product_id` | `collection_products_product_id_fkey` | [Products](./products) |


## Entity Relationship Diagram

```mermaid
erDiagram
    CollectionProducts }o--|| Collections : "FK"
    CollectionProducts }o--|| Products : "FK"
```

::::tabs

=== FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `collection_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `UQ` `FK` | → References `collections` |
| 4 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `UQ` `FK` | → References `products` |
| 5 | `sort_order` | `integer` | `int` | `number` | NO | `0` | - | - |
| 6 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `collection_id` | `collections` | `collection_products_collection_id_fkey` |
| `product_id` | `products` | `collection_products_product_id_fkey` |

## Unique Keys

- `collection_id` (`uuid`)
- `product_id` (`uuid`)


## Go Generated Code

> 📂 Source: [📄 `CollectionProducts.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/CollectionProducts.go) · [📄 `CollectionProducts.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/CollectionProducts.go) · [📄 `CollectionProducts.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/CollectionProducts.go)

### Structs

:::tabs

== Form

#### CollectionProductsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/CollectionProducts.go#:~:text=type%20CollectionProductsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `CollectionId` | `uuid.UUID` | `collectionId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### CollectionProducts [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/CollectionProducts.go#:~:text=type%20CollectionProducts%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `CollectionId` | `uuid.UUID` | `collectionId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### CollectionProductsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/CollectionProducts.go#:~:text=type%20CollectionProductsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `CollectionId` | `*uuid.UUID` | `collectionId` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `SortOrder` | `*int` | `sortOrder` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### CollectionProductsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/CollectionProducts.go#:~:text=type%20CollectionProductsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `CollectionId` | `*uuid.UUID` | `collectionId` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `SortOrder` | `*int` | `sortOrder` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### CollectionProductsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/CollectionProducts.go#:~:text=type%20CollectionProductsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `CollectionId` | `uuid.UUID` | `collectionId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### CollectionProductsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/CollectionProducts.go#:~:text=type%20CollectionProductsBatchUpdate%20struct)

```go
type CollectionProductsBatchUpdate struct {
    Data       json.RawMessage `json:"data"`
    PathParams struct {
        Id uuid.UUID
    } `json:"pathParams"`
}
```

:::

### Service & Endpoints

:::tabs

== Service Methods

| Method | Signature |
|---------|-----------|
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/CollectionProducts.go#:~:text=%29%20CreateCollectionProducts%28%29) | `(CollectionProductsService) CreateCollectionProducts(data CollectionProductsForm) (CollectionProductsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/CollectionProducts.go#:~:text=%29%20CreateCollectionProductsMultiple%28%29) | `(CollectionProductsService) CreateCollectionProductsMultiple(data []CollectionProductsForm) ([]CollectionProductsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/CollectionProducts.go#:~:text=%29%20UpdateCollectionProducts%28%29) | `(CollectionProductsService) UpdateCollectionProducts(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/CollectionProducts.go#:~:text=%29%20UpdateCollectionProductsMultiple%28%29) | `(CollectionProductsService) UpdateCollectionProductsMultiple(data []CollectionProductsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/CollectionProducts.go#:~:text=%29%20DeleteCollectionProducts%28%29) | `(CollectionProductsService) DeleteCollectionProducts(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/collection-products/` | Search with query params |
| `GET` | `/collection-products/pagination` | Paginated listing |
| `POST` | `/collection-products/` | Create single record |
| `POST` | `/collection-products/bulk/` | Create multiple records |
| `PUT` | `/collection-products/bulk/` | Batch update |
| `GET` | `/collection-products/with-id/:id` | Get by ID |
| `PUT` | `/collection-products/with-id/:id` | Update by ID |
| `DELETE` | `/collection-products/with-id/:id` | Delete by ID |

== Query & Filters

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

:::

### RPC Functions

| Function | Parameters | Return | Endpoint |
|----------|-----------|--------|----------|
| `avg_product_rating` | `p_product_id uuid` | `numeric` | `/rpc/avg_product_rating` |
| `count_active_products` | - | `integer` | `/rpc/count_active_products` |
| `products_by_category` | `p_category_id uuid` | `integer` | `/rpc/products_by_category` |


=== Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export interface CollectionProducts {
  id: string;
  name: string;
  collectionId: string;
  productId: string;
  sortOrder: number;
  createdAt: string;
}

export interface CollectionProductsForm {
  name: string;
  collectionId: string;
  productId: string;
  sortOrder: number;
  createdAt: string;
}

export interface CollectionProductsEdit {
  id: string;
  name: string;
  collectionId: string;
  productId: string;
  sortOrder: number;
  createdAt: string;
}

export interface CollectionProductsPage {
  data: CollectionProducts[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type CollectionProductsPathQuery = {
  page?: number;
  size?: number;
  sort?: string;
  fields?: string;
  preloads?: string;
  filters?: string;
};

```

== React Query

```typescript
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

const CollectionProductsKeys = {
  all: ["collection_products"] as const,
  lists: () => [...CollectionProductsKeys.all, "list"] as const,
  detail: (id: any) => [...CollectionProductsKeys.all, "detail", id] as const,
} as const;

export function useCollectionProductsList(query?: CollectionProductsPathQuery) {
  return useQuery({
    queryKey: [...CollectionProductsKeys.lists(), query],
    queryFn: () => fetch(`/collection-products/pagination`, { method: "GET" }).then(r => r.json()) as Promise<CollectionProductsPage>,
  });
}

export function useCollectionProductsDetail(id: any) {
  return useQuery({
    queryKey: CollectionProductsKeys.detail(id),
    queryFn: () => fetch(`/collection-products/with-id/:id`).then(r => r.json()) as Promise<CollectionProducts>,
  });
}

export function useCreateCollectionProducts() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: CollectionProductsForm) =>
      fetch("/collection-products/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CollectionProductsKeys.lists() }),
  });
}

export function useUpdateCollectionProducts() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: CollectionProductsEdit }) =>
      fetch(`/collection-products/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CollectionProductsKeys.all }),
  });
}

export function useDeleteCollectionProducts() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/collection-products/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CollectionProductsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const CollectionProductsFormSchema = z.object({
  name: z.string(),
  collectionId: z.string().uuid(),
  productId: z.string().uuid(),
  sortOrder: z.number().int(),
  createdAt: z.string().datetime(),
});

export type CollectionProductsFormInput = z.infer<typeof CollectionProductsFormSchema>;

```

:::


=== API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './collection_products.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search CollectionProducts

```
GET /api/v1/collection-products/
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
| `preloads` | `string` | No | Available: CollectionIdDetail, CollectionIdDetail.CollectionProductsList, CollectionIdDetail.CollectionProductsList.CollectionIdDetail, CollectionIdDetail.CollectionProductsList.ProductIdDetail, ProductIdDetail, ProductIdDetail.ProductVariantsList, ProductIdDetail.ProductVariantsList.ProductIdDetail, ProductIdDetail.ProductMediaList, ProductIdDetail.ProductMediaList.ProductIdDetail, ProductIdDetail.ProductReviewsList, ProductIdDetail.ProductReviewsList.ProductIdDetail, ProductIdDetail.CollectionProductsList, ProductIdDetail.CollectionProductsList.CollectionIdDetail, ProductIdDetail.CollectionProductsList.ProductIdDetail, ProductIdDetail.ProductTagsList, ProductIdDetail.ProductTagsList.ProductIdDetail, ProductIdDetail.ProductTagsList.TagIdDetail, ProductIdDetail.PriceHistoryList, ProductIdDetail.PriceHistoryList.ProductIdDetail, ProductIdDetail.BrandIdDetail, ProductIdDetail.BrandIdDetail.ProductsList, ProductIdDetail.CategoryIdDetail, ProductIdDetail.CategoryIdDetail.CategoriesList, ProductIdDetail.CategoryIdDetail.ProductsList, ProductIdDetail.CategoryIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Collections, Collections.Users, Products, Products.Brands, Products.Brands.Organizations, Products.Categories, Products.Categories.Categories, Products.Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `collectionId` | `string (uuid)` | No | Filter by collection_id |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `sortOrder` | `integer` | No | Filter by sort_order |

**Response:** `CollectionProducts[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/collection-products/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search CollectionProducts (POST)

```
POST /api/v1/collection-products/search
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

**Response:** `CollectionProducts[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collection-products/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate CollectionProducts

```
GET /api/v1/collection-products/pagination
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
| `preloads` | `string` | No | Available: CollectionIdDetail, CollectionIdDetail.CollectionProductsList, CollectionIdDetail.CollectionProductsList.CollectionIdDetail, CollectionIdDetail.CollectionProductsList.ProductIdDetail, ProductIdDetail, ProductIdDetail.ProductVariantsList, ProductIdDetail.ProductVariantsList.ProductIdDetail, ProductIdDetail.ProductMediaList, ProductIdDetail.ProductMediaList.ProductIdDetail, ProductIdDetail.ProductReviewsList, ProductIdDetail.ProductReviewsList.ProductIdDetail, ProductIdDetail.CollectionProductsList, ProductIdDetail.CollectionProductsList.CollectionIdDetail, ProductIdDetail.CollectionProductsList.ProductIdDetail, ProductIdDetail.ProductTagsList, ProductIdDetail.ProductTagsList.ProductIdDetail, ProductIdDetail.ProductTagsList.TagIdDetail, ProductIdDetail.PriceHistoryList, ProductIdDetail.PriceHistoryList.ProductIdDetail, ProductIdDetail.BrandIdDetail, ProductIdDetail.BrandIdDetail.ProductsList, ProductIdDetail.CategoryIdDetail, ProductIdDetail.CategoryIdDetail.CategoriesList, ProductIdDetail.CategoryIdDetail.ProductsList, ProductIdDetail.CategoryIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Collections, Collections.Users, Products, Products.Brands, Products.Brands.Organizations, Products.Categories, Products.Categories.Categories, Products.Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `collectionId` | `string (uuid)` | No | Filter by collection_id |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `sortOrder` | `integer` | No | Filter by sort_order |

**Response:** `PaginationResponse<CollectionProducts>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/collection-products/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate CollectionProducts (POST)

```
POST /api/v1/collection-products/pagination
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

**Response:** `PaginationResponse<CollectionProducts>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collection-products/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create CollectionProducts

```
POST /api/v1/collection-products/
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
  collectionId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  sortOrder?: number  // e.g. 1
}
```

**Response:** `CollectionProducts`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collection-products/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create CollectionProducts

```
POST /api/v1/collection-products/bulk/
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
  collectionId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  sortOrder?: number  // e.g. 1
}
```

**Response:** `CollectionProducts[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collection-products/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find CollectionProducts by ID

```
GET /api/v1/collection-products/with-id/:id
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

**Response:** `CollectionProducts`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/collection-products/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update CollectionProducts

```
PUT /api/v1/collection-products/with-id/:id
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
  collectionId?: string
  productId?: string
  sortOrder?: number
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
  "http://localhost:3000/api/v1/collection-products/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update CollectionProducts

```
PUT /api/v1/collection-products/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: CollectionProductsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collection-products/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete CollectionProducts

```
DELETE /api/v1/collection-products/with-id/:id
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
  "http://localhost:3000/api/v1/collection-products/with-id/:id"
```

</details>

---

:::


::::
