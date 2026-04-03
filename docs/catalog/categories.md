---
title: Categories
---

# Categories

**Table:** `catalog.categories`

**Base path:** `/categories`

## Related Tables

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `products` | `category_id` | `products_category_id_fkey` | [Products](./products) |


## Entity Relationship Diagram

```mermaid
erDiagram
    Categories }o--|| Categories : "FK"
    Categories ||--o{ Categories : "ref"
    Categories ||--o{ Products : "ref"
```

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `slug` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 4 | `description` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `parent_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `categories` |
| 6 | `icon` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `sort_order` | `integer` | `int` | `number` | NO | `0` | - | - |
| 8 | `is_active` | `boolean` | `bool` | `boolean` | NO | `true` | - | - |
| 9 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 10 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `parent_id` | `categories` | `categories_parent_id_fkey` |

## Unique Keys

- `slug` (`text`)


## Go Generated Code

> 📂 Source: [📄 `Categories.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Categories.go) · [📄 `Categories.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Categories.go) · [📄 `Categories.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/Categories.go)

### Structs

::::tabs

:::tab Form

#### CategoriesForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Categories.go#:~:text=type%20CategoriesForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `ParentId` | `*uuid.UUID` | `parentId` | YES |
| `Icon` | `string` | `icon` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Model

#### Categories [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Categories.go#:~:text=type%20Categories%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `ParentId` | `*uuid.UUID` | `parentId` | YES |
| `Icon` | `string` | `icon` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Edit

#### CategoriesEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Categories.go#:~:text=type%20CategoriesEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `Description` | `*string` | `description` | YES |
| `ParentId` | `*uuid.UUID` | `parentId` | YES |
| `Icon` | `*string` | `icon` | YES |
| `SortOrder` | `*int` | `sortOrder` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Filter

#### CategoriesFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Categories.go#:~:text=type%20CategoriesFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `Description` | `*string` | `description` | YES |
| `ParentId` | `*uuid.UUID` | `parentId` | YES |
| `Icon` | `*string` | `icon` | YES |
| `SortOrder` | `*int` | `sortOrder` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Page

#### CategoriesPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Categories.go#:~:text=type%20CategoriesPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `ParentId` | `*uuid.UUID` | `parentId` | YES |
| `Icon` | `string` | `icon` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab BatchUpdate

#### CategoriesBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Categories.go#:~:text=type%20CategoriesBatchUpdate%20struct)

```go
type CategoriesBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Categories.go#:~:text=)%20CreateCategories() | `(CategoriesService) CreateCategories(data CategoriesForm) (CategoriesForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Categories.go#:~:text=)%20CreateCategoriesMultiple() | `(CategoriesService) CreateCategoriesMultiple(data []CategoriesForm) ([]CategoriesForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Categories.go#:~:text=)%20UpdateCategories() | `(CategoriesService) UpdateCategories(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Categories.go#:~:text=)%20UpdateCategoriesMultiple() | `(CategoriesService) UpdateCategoriesMultiple(data []CategoriesBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Categories.go#:~:text=)%20DeleteCategories() | `(CategoriesService) DeleteCategories(id uuid.UUID) error` |

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/categories/` | Search with query params |
| `GET` | `/categories/pagination` | Paginated listing |
| `POST` | `/categories/` | Create single record |
| `POST` | `/categories/bulk/` | Create multiple records |
| `PUT` | `/categories/bulk/` | Batch update |
| `GET` | `/categories/with-id/:id` | Get by ID |
| `PUT` | `/categories/with-id/:id` | Update by ID |
| `DELETE` | `/categories/with-id/:id` | Delete by ID |

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
export interface Categories {
  id: string;
  name: string;
  slug: string;
  description: string;
  parentId?: string;
  icon: string;
  sortOrder: number;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface CategoriesForm {
  name: string;
  slug: string;
  description: string;
  parentId?: string;
  icon: string;
  sortOrder: number;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface CategoriesEdit {
  id: string;
  name: string;
  slug: string;
  description: string;
  parentId?: string;
  icon: string;
  sortOrder: number;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface CategoriesPage {
  data: Categories[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type CategoriesPathQuery = {
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

const CategoriesKeys = {
  all: ["categories"] as const,
  lists: () => [...CategoriesKeys.all, "list"] as const,
  detail: (id: any) => [...CategoriesKeys.all, "detail", id] as const,
} as const;

export function useCategoriesList(query?: CategoriesPathQuery) {
  return useQuery({
    queryKey: [...CategoriesKeys.lists(), query],
    queryFn: () => fetch(`/categories/pagination`, { method: "GET" }).then(r => r.json()) as Promise<CategoriesPage>,
  });
}

export function useCategoriesDetail(id: any) {
  return useQuery({
    queryKey: CategoriesKeys.detail(id),
    queryFn: () => fetch(`/categories/with-id/:id`).then(r => r.json()) as Promise<Categories>,
  });
}

export function useCreateCategories() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: CategoriesForm) =>
      fetch("/categories/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CategoriesKeys.lists() }),
  });
}

export function useUpdateCategories() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: CategoriesEdit }) =>
      fetch(`/categories/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CategoriesKeys.all }),
  });
}

export function useDeleteCategories() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/categories/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CategoriesKeys.all }),
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

export const CategoriesFormSchema = z.object({
  name: z.string(),
  slug: z.string(),
  description: z.string(),
  parentId: z.string().uuid().optional(),
  icon: z.string(),
  sortOrder: z.number().int(),
  isActive: z.boolean(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type CategoriesFormInput = z.infer<typeof CategoriesFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './categories.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search Categories

```
GET /api/v1/categories/
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
| `preloads` | `string` | No | Available: CategoriesList, CategoriesList.CategoriesList, CategoriesList.ProductsList, CategoriesList.ProductsList.ProductVariantsList, CategoriesList.ProductsList.ProductMediaList, CategoriesList.ProductsList.ProductReviewsList, CategoriesList.ProductsList.CollectionProductsList, CategoriesList.ProductsList.ProductTagsList, CategoriesList.ProductsList.PriceHistoryList, CategoriesList.ProductsList.BrandIdDetail, CategoriesList.ProductsList.CategoryIdDetail, CategoriesList.ParentIdDetail, ProductsList, ProductsList.ProductVariantsList, ProductsList.ProductVariantsList.ProductIdDetail, ProductsList.ProductMediaList, ProductsList.ProductMediaList.ProductIdDetail, ProductsList.ProductReviewsList, ProductsList.ProductReviewsList.ProductIdDetail, ProductsList.CollectionProductsList, ProductsList.CollectionProductsList.CollectionIdDetail, ProductsList.CollectionProductsList.ProductIdDetail, ProductsList.ProductTagsList, ProductsList.ProductTagsList.ProductIdDetail, ProductsList.ProductTagsList.TagIdDetail, ProductsList.PriceHistoryList, ProductsList.PriceHistoryList.ProductIdDetail, ProductsList.BrandIdDetail, ProductsList.BrandIdDetail.ProductsList, ProductsList.CategoryIdDetail, ProductsList.CategoryIdDetail.CategoriesList, ProductsList.CategoryIdDetail.ProductsList, ProductsList.CategoryIdDetail.ParentIdDetail, ParentIdDetail, ParentIdDetail.CategoriesList, ParentIdDetail.ProductsList, ParentIdDetail.ProductsList.ProductVariantsList, ParentIdDetail.ProductsList.ProductMediaList, ParentIdDetail.ProductsList.ProductReviewsList, ParentIdDetail.ProductsList.CollectionProductsList, ParentIdDetail.ProductsList.ProductTagsList, ParentIdDetail.ProductsList.PriceHistoryList, ParentIdDetail.ProductsList.BrandIdDetail, ParentIdDetail.ProductsList.CategoryIdDetail, ParentIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Categories, Categories.Categories |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |
| `description` | `string` | No | Filter by description |
| `parentId` | `string (uuid)` | No | Filter by parent_id |
| `icon` | `string` | No | Filter by icon |
| `sortOrder` | `integer` | No | Filter by sort_order |
| `isActive` | `boolean` | No | Filter by is_active |

**Response:** `Categories[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/categories/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Categories (POST)

```
POST /api/v1/categories/search
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

**Response:** `Categories[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/categories/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate Categories

```
GET /api/v1/categories/pagination
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
| `preloads` | `string` | No | Available: CategoriesList, CategoriesList.CategoriesList, CategoriesList.ProductsList, CategoriesList.ProductsList.ProductVariantsList, CategoriesList.ProductsList.ProductMediaList, CategoriesList.ProductsList.ProductReviewsList, CategoriesList.ProductsList.CollectionProductsList, CategoriesList.ProductsList.ProductTagsList, CategoriesList.ProductsList.PriceHistoryList, CategoriesList.ProductsList.BrandIdDetail, CategoriesList.ProductsList.CategoryIdDetail, CategoriesList.ParentIdDetail, ProductsList, ProductsList.ProductVariantsList, ProductsList.ProductVariantsList.ProductIdDetail, ProductsList.ProductMediaList, ProductsList.ProductMediaList.ProductIdDetail, ProductsList.ProductReviewsList, ProductsList.ProductReviewsList.ProductIdDetail, ProductsList.CollectionProductsList, ProductsList.CollectionProductsList.CollectionIdDetail, ProductsList.CollectionProductsList.ProductIdDetail, ProductsList.ProductTagsList, ProductsList.ProductTagsList.ProductIdDetail, ProductsList.ProductTagsList.TagIdDetail, ProductsList.PriceHistoryList, ProductsList.PriceHistoryList.ProductIdDetail, ProductsList.BrandIdDetail, ProductsList.BrandIdDetail.ProductsList, ProductsList.CategoryIdDetail, ProductsList.CategoryIdDetail.CategoriesList, ProductsList.CategoryIdDetail.ProductsList, ProductsList.CategoryIdDetail.ParentIdDetail, ParentIdDetail, ParentIdDetail.CategoriesList, ParentIdDetail.ProductsList, ParentIdDetail.ProductsList.ProductVariantsList, ParentIdDetail.ProductsList.ProductMediaList, ParentIdDetail.ProductsList.ProductReviewsList, ParentIdDetail.ProductsList.CollectionProductsList, ParentIdDetail.ProductsList.ProductTagsList, ParentIdDetail.ProductsList.PriceHistoryList, ParentIdDetail.ProductsList.BrandIdDetail, ParentIdDetail.ProductsList.CategoryIdDetail, ParentIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Categories, Categories.Categories |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |
| `description` | `string` | No | Filter by description |
| `parentId` | `string (uuid)` | No | Filter by parent_id |
| `icon` | `string` | No | Filter by icon |
| `sortOrder` | `integer` | No | Filter by sort_order |
| `isActive` | `boolean` | No | Filter by is_active |

**Response:** `PaginationResponse<Categories>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/categories/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Categories (POST)

```
POST /api/v1/categories/pagination
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

**Response:** `PaginationResponse<Categories>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/categories/pagination"
```

</details>

---

:::tab Create

#### <Badge type="tip" text="POST" /> Create Categories

```
POST /api/v1/categories/
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
  parentId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  icon?: string  // e.g. example_icon
  sortOrder?: number  // e.g. 1
  isActive?: boolean  // e.g. true
}
```

**Response:** `Categories`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/categories/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Categories

```
POST /api/v1/categories/bulk/
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
  parentId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  icon?: string  // e.g. example_icon
  sortOrder?: number  // e.g. 1
  isActive?: boolean  // e.g. true
}
```

**Response:** `Categories[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/categories/bulk/"
```

</details>

---

:::tab Find & Update

#### <Badge type="info" text="GET" /> Find Categories by ID

```
GET /api/v1/categories/with-id/:id
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

**Response:** `Categories`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/categories/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Categories

```
PUT /api/v1/categories/with-id/:id
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
  parentId?: string
  icon?: string
  sortOrder?: number
  isActive?: boolean
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
  "http://localhost:3000/api/v1/categories/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Categories

```
PUT /api/v1/categories/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: CategoriesEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/categories/bulk/"
```

</details>

---

:::tab Delete

#### <Badge type="danger" text="DELETE" /> Delete Categories

```
DELETE /api/v1/categories/with-id/:id
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
  "http://localhost:3000/api/v1/categories/with-id/:id"
```

</details>

---

::::


::::
