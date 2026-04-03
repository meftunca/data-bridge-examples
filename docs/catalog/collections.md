---
title: Collections
---

# Collections

**Table:** `catalog.collections`

**Base path:** `/collections`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `users` | `created_by` | `collections_created_by_fkey` | [Users](./users) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `collection_products` | `collection_id` | `collection_products_collection_id_fkey` | [CollectionProducts](./collection_products) |


## Entity Relationship Diagram

```mermaid
erDiagram
    Collections }o--|| Users : "FK"
    Collections ||--o{ CollectionProducts : "ref"
```

::::tabs

=== FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `slug` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 4 | `description` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `cover_image_url` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 6 | `is_active` | `boolean` | `bool` | `boolean` | NO | `true` | - | - |
| 7 | `start_date` | `date` | `time.Time` | `string` | YES | - | - | - |
| 8 | `end_date` | `date` | `time.Time` | `string` | YES | - | - | - |
| 9 | `created_by` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | Auto-filled from session |
| 10 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 11 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `created_by` | `users` | `collections_created_by_fkey` |

## Unique Keys

- `slug` (`text`)


## Go Generated Code

> 📂 Source: [📄 `Collections.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Collections.go) · [📄 `Collections.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Collections.go) · [📄 `Collections.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/Collections.go)

### Structs

:::tabs

== Form

#### CollectionsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Collections.go#:~:text=type%20CollectionsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `CoverImageUrl` | `string` | `coverImageUrl` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `StartDate` | `*time.Time` | `startDate` | YES |
| `EndDate` | `*time.Time` | `endDate` | YES |
| `CreatedBy` | `*uuid.UUID` | `createdBy` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### Collections [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Collections.go#:~:text=type%20Collections%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `CoverImageUrl` | `string` | `coverImageUrl` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `StartDate` | `*time.Time` | `startDate` | YES |
| `EndDate` | `*time.Time` | `endDate` | YES |
| `CreatedBy` | `*uuid.UUID` | `createdBy` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### CollectionsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Collections.go#:~:text=type%20CollectionsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `Description` | `*string` | `description` | YES |
| `CoverImageUrl` | `*string` | `coverImageUrl` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `StartDate` | `*time.Time` | `startDate` | YES |
| `EndDate` | `*time.Time` | `endDate` | YES |
| `CreatedBy` | `*uuid.UUID` | `createdBy` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### CollectionsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Collections.go#:~:text=type%20CollectionsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `Description` | `*string` | `description` | YES |
| `CoverImageUrl` | `*string` | `coverImageUrl` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `StartDate` | `*time.Time` | `startDate` | YES |
| `EndDate` | `*time.Time` | `endDate` | YES |
| `CreatedBy` | `*uuid.UUID` | `createdBy` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### CollectionsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Collections.go#:~:text=type%20CollectionsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `CoverImageUrl` | `string` | `coverImageUrl` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `StartDate` | `*time.Time` | `startDate` | YES |
| `EndDate` | `*time.Time` | `endDate` | YES |
| `CreatedBy` | `*uuid.UUID` | `createdBy` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### CollectionsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Collections.go#:~:text=type%20CollectionsBatchUpdate%20struct)

```go
type CollectionsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Collections.go#:~:text=%29%20CreateCollections%28%29) | `(CollectionsService) CreateCollections(data CollectionsForm) (CollectionsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Collections.go#:~:text=%29%20CreateCollectionsMultiple%28%29) | `(CollectionsService) CreateCollectionsMultiple(data []CollectionsForm) ([]CollectionsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Collections.go#:~:text=%29%20UpdateCollections%28%29) | `(CollectionsService) UpdateCollections(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Collections.go#:~:text=%29%20UpdateCollectionsMultiple%28%29) | `(CollectionsService) UpdateCollectionsMultiple(data []CollectionsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Collections.go#:~:text=%29%20DeleteCollections%28%29) | `(CollectionsService) DeleteCollections(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/collections/` | Search with query params |
| `GET` | `/collections/pagination` | Paginated listing |
| `POST` | `/collections/` | Create single record |
| `POST` | `/collections/bulk/` | Create multiple records |
| `PUT` | `/collections/bulk/` | Batch update |
| `GET` | `/collections/with-id/:id` | Get by ID |
| `PUT` | `/collections/with-id/:id` | Update by ID |
| `DELETE` | `/collections/with-id/:id` | Delete by ID |

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
export interface Collections {
  id: string;
  name: string;
  slug: string;
  description: string;
  coverImageUrl: string;
  isActive: boolean;
  startDate?: string;
  endDate?: string;
  createdBy?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CollectionsForm {
  name: string;
  slug: string;
  description: string;
  coverImageUrl: string;
  isActive: boolean;
  startDate?: string;
  endDate?: string;
  createdBy?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CollectionsEdit {
  id: string;
  name: string;
  slug: string;
  description: string;
  coverImageUrl: string;
  isActive: boolean;
  startDate?: string;
  endDate?: string;
  createdBy?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CollectionsPage {
  data: Collections[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type CollectionsPathQuery = {
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

const CollectionsKeys = {
  all: ["collections"] as const,
  lists: () => [...CollectionsKeys.all, "list"] as const,
  detail: (id: any) => [...CollectionsKeys.all, "detail", id] as const,
} as const;

export function useCollectionsList(query?: CollectionsPathQuery) {
  return useQuery({
    queryKey: [...CollectionsKeys.lists(), query],
    queryFn: () => fetch(`/collections/pagination`, { method: "GET" }).then(r => r.json()) as Promise<CollectionsPage>,
  });
}

export function useCollectionsDetail(id: any) {
  return useQuery({
    queryKey: CollectionsKeys.detail(id),
    queryFn: () => fetch(`/collections/with-id/:id`).then(r => r.json()) as Promise<Collections>,
  });
}

export function useCreateCollections() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: CollectionsForm) =>
      fetch("/collections/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CollectionsKeys.lists() }),
  });
}

export function useUpdateCollections() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: CollectionsEdit }) =>
      fetch(`/collections/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CollectionsKeys.all }),
  });
}

export function useDeleteCollections() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/collections/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CollectionsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const CollectionsFormSchema = z.object({
  name: z.string(),
  slug: z.string(),
  description: z.string(),
  coverImageUrl: z.string(),
  isActive: z.boolean(),
  startDate: z.string().datetime().optional(),
  endDate: z.string().datetime().optional(),
  createdBy: z.string().uuid().optional(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type CollectionsFormInput = z.infer<typeof CollectionsFormSchema>;

```

:::


=== API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './collections.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Collections

```
GET /api/v1/collections/
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
| `preloads` | `string` | No | Available: CollectionProductsList, CollectionProductsList.CollectionIdDetail, CollectionProductsList.CollectionIdDetail.CollectionProductsList, CollectionProductsList.ProductIdDetail, CollectionProductsList.ProductIdDetail.ProductVariantsList, CollectionProductsList.ProductIdDetail.ProductMediaList, CollectionProductsList.ProductIdDetail.ProductReviewsList, CollectionProductsList.ProductIdDetail.CollectionProductsList, CollectionProductsList.ProductIdDetail.ProductTagsList, CollectionProductsList.ProductIdDetail.PriceHistoryList, CollectionProductsList.ProductIdDetail.BrandIdDetail, CollectionProductsList.ProductIdDetail.CategoryIdDetail |
| `joins` | `string` | No | Available: Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |
| `description` | `string` | No | Filter by description |
| `coverImageUrl` | `string` | No | Filter by cover_image_url |
| `isActive` | `boolean` | No | Filter by is_active |
| `startDate` | `string (date)` | No | Filter by start_date |
| `endDate` | `string (date)` | No | Filter by end_date |

**Response:** `Collections[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/collections/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Collections (POST)

```
POST /api/v1/collections/search
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

**Response:** `Collections[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collections/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Collections

```
GET /api/v1/collections/pagination
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
| `preloads` | `string` | No | Available: CollectionProductsList, CollectionProductsList.CollectionIdDetail, CollectionProductsList.CollectionIdDetail.CollectionProductsList, CollectionProductsList.ProductIdDetail, CollectionProductsList.ProductIdDetail.ProductVariantsList, CollectionProductsList.ProductIdDetail.ProductMediaList, CollectionProductsList.ProductIdDetail.ProductReviewsList, CollectionProductsList.ProductIdDetail.CollectionProductsList, CollectionProductsList.ProductIdDetail.ProductTagsList, CollectionProductsList.ProductIdDetail.PriceHistoryList, CollectionProductsList.ProductIdDetail.BrandIdDetail, CollectionProductsList.ProductIdDetail.CategoryIdDetail |
| `joins` | `string` | No | Available: Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |
| `description` | `string` | No | Filter by description |
| `coverImageUrl` | `string` | No | Filter by cover_image_url |
| `isActive` | `boolean` | No | Filter by is_active |
| `startDate` | `string (date)` | No | Filter by start_date |
| `endDate` | `string (date)` | No | Filter by end_date |

**Response:** `PaginationResponse<Collections>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/collections/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Collections (POST)

```
POST /api/v1/collections/pagination
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

**Response:** `PaginationResponse<Collections>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collections/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Collections

```
POST /api/v1/collections/
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
  name: string  // e.g. example_name
  slug: string  // e.g. example_slug
  description?: string  // e.g. example_description
  coverImageUrl?: string  // e.g. example_cover_image_url
  isActive?: boolean  // e.g. true
  startDate?: string  // e.g. 2026-01-15
  endDate?: string  // e.g. 2026-01-15
}
```

**Response:** `Collections`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collections/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Collections

```
POST /api/v1/collections/bulk/
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
  name: string  // e.g. example_name
  slug: string  // e.g. example_slug
  description?: string  // e.g. example_description
  coverImageUrl?: string  // e.g. example_cover_image_url
  isActive?: boolean  // e.g. true
  startDate?: string  // e.g. 2026-01-15
  endDate?: string  // e.g. 2026-01-15
}
```

**Response:** `Collections[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collections/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Collections by ID

```
GET /api/v1/collections/with-id/:id
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

**Response:** `Collections`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/collections/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Collections

```
PUT /api/v1/collections/with-id/:id
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
  slug?: string
  description?: string
  coverImageUrl?: string
  isActive?: boolean
  startDate?: string
  endDate?: string
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
  "http://localhost:3000/api/v1/collections/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Collections

```
PUT /api/v1/collections/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: CollectionsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/collections/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Collections

```
DELETE /api/v1/collections/with-id/:id
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
  "http://localhost:3000/api/v1/collections/with-id/:id"
```

</details>

---

:::


::::
