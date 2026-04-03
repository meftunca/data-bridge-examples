---
title: Brands
---

# Brands

**Table:** `catalog.brands`

**Base path:** `/brands`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `organizations` | `organization_id` | `brands_organization_id_fkey` | [Organizations](./organizations) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `products` | `brand_id` | `products_brand_id_fkey` | [Products](./products) |


## Entity Relationship Diagram

erDiagram
    Brands }o--|| Organizations : "FK"
    Brands ||--o{ Products : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `slug` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 4 | `logo_url` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `website` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 6 | `description` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `organization_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `organizations` |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 9 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `organization_id` | `organizations` | `brands_organization_id_fkey` |

## Unique Keys

- `slug` (`text`)


## Go Generated Code

> 📂 Source: [📄 `Brands.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Brands.go) · [📄 `Brands.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Brands.go) · [📄 `Brands.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/Brands.go)

### Structs

:::tabs

== Form

#### BrandsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Brands.go#:~:text=type%20BrandsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `LogoUrl` | `string` | `logoUrl` | NO |
| `Website` | `string` | `website` | NO |
| `Description` | `string` | `description` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### Brands [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Brands.go#:~:text=type%20Brands%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `LogoUrl` | `string` | `logoUrl` | NO |
| `Website` | `string` | `website` | NO |
| `Description` | `string` | `description` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### BrandsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Brands.go#:~:text=type%20BrandsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `LogoUrl` | `*string` | `logoUrl` | YES |
| `Website` | `*string` | `website` | YES |
| `Description` | `*string` | `description` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### BrandsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Brands.go#:~:text=type%20BrandsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `LogoUrl` | `*string` | `logoUrl` | YES |
| `Website` | `*string` | `website` | YES |
| `Description` | `*string` | `description` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### BrandsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Brands.go#:~:text=type%20BrandsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `LogoUrl` | `string` | `logoUrl` | NO |
| `Website` | `string` | `website` | NO |
| `Description` | `string` | `description` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### BrandsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Brands.go#:~:text=type%20BrandsBatchUpdate%20struct)

```go
type BrandsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Brands.go#:~:text=)%20CreateBrands() | `(BrandsService) CreateBrands(data BrandsForm) (BrandsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Brands.go#:~:text=)%20CreateBrandsMultiple() | `(BrandsService) CreateBrandsMultiple(data []BrandsForm) ([]BrandsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Brands.go#:~:text=)%20UpdateBrands() | `(BrandsService) UpdateBrands(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Brands.go#:~:text=)%20UpdateBrandsMultiple() | `(BrandsService) UpdateBrandsMultiple(data []BrandsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Brands.go#:~:text=)%20DeleteBrands() | `(BrandsService) DeleteBrands(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/brands/` | Search with query params |
| `GET` | `/brands/pagination` | Paginated listing |
| `POST` | `/brands/` | Create single record |
| `POST` | `/brands/bulk/` | Create multiple records |
| `PUT` | `/brands/bulk/` | Batch update |
| `GET` | `/brands/with-id/:id` | Get by ID |
| `PUT` | `/brands/with-id/:id` | Update by ID |
| `DELETE` | `/brands/with-id/:id` | Delete by ID |

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


:::tab Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export interface Brands {
  id: string;
  name: string;
  slug: string;
  logoUrl: string;
  website: string;
  description: string;
  organizationId?: string;
  createdAt: string;
  updatedAt: string;
}

export interface BrandsForm {
  name: string;
  slug: string;
  logoUrl: string;
  website: string;
  description: string;
  organizationId?: string;
  createdAt: string;
  updatedAt: string;
}

export interface BrandsEdit {
  id: string;
  name: string;
  slug: string;
  logoUrl: string;
  website: string;
  description: string;
  organizationId?: string;
  createdAt: string;
  updatedAt: string;
}

export interface BrandsPage {
  data: Brands[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type BrandsPathQuery = {
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

const BrandsKeys = {
  all: ["brands"] as const,
  lists: () => [...BrandsKeys.all, "list"] as const,
  detail: (id: any) => [...BrandsKeys.all, "detail", id] as const,
} as const;

export function useBrandsList(query?: BrandsPathQuery) {
  return useQuery({
    queryKey: [...BrandsKeys.lists(), query],
    queryFn: () => fetch(`/brands/pagination`, { method: "GET" }).then(r => r.json()) as Promise<BrandsPage>,
  });
}

export function useBrandsDetail(id: any) {
  return useQuery({
    queryKey: BrandsKeys.detail(id),
    queryFn: () => fetch(`/brands/with-id/:id`).then(r => r.json()) as Promise<Brands>,
  });
}

export function useCreateBrands() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: BrandsForm) =>
      fetch("/brands/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: BrandsKeys.lists() }),
  });
}

export function useUpdateBrands() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: BrandsEdit }) =>
      fetch(`/brands/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: BrandsKeys.all }),
  });
}

export function useDeleteBrands() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/brands/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: BrandsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const BrandsFormSchema = z.object({
  name: z.string(),
  slug: z.string(),
  logoUrl: z.string(),
  website: z.string(),
  description: z.string(),
  organizationId: z.string().uuid().optional(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type BrandsFormInput = z.infer<typeof BrandsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './brands.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Brands

```
GET /api/v1/brands/
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
| `preloads` | `string` | No | Available: ProductsList, ProductsList.ProductVariantsList, ProductsList.ProductVariantsList.ProductIdDetail, ProductsList.ProductMediaList, ProductsList.ProductMediaList.ProductIdDetail, ProductsList.ProductReviewsList, ProductsList.ProductReviewsList.ProductIdDetail, ProductsList.CollectionProductsList, ProductsList.CollectionProductsList.CollectionIdDetail, ProductsList.CollectionProductsList.ProductIdDetail, ProductsList.ProductTagsList, ProductsList.ProductTagsList.ProductIdDetail, ProductsList.ProductTagsList.TagIdDetail, ProductsList.PriceHistoryList, ProductsList.PriceHistoryList.ProductIdDetail, ProductsList.BrandIdDetail, ProductsList.BrandIdDetail.ProductsList, ProductsList.CategoryIdDetail, ProductsList.CategoryIdDetail.CategoriesList, ProductsList.CategoryIdDetail.ProductsList, ProductsList.CategoryIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |
| `logoUrl` | `string` | No | Filter by logo_url |
| `website` | `string` | No | Filter by website |
| `description` | `string` | No | Filter by description |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |

**Response:** `Brands[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/brands/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Brands (POST)

```
POST /api/v1/brands/search
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

**Response:** `Brands[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/brands/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Brands

```
GET /api/v1/brands/pagination
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
| `preloads` | `string` | No | Available: ProductsList, ProductsList.ProductVariantsList, ProductsList.ProductVariantsList.ProductIdDetail, ProductsList.ProductMediaList, ProductsList.ProductMediaList.ProductIdDetail, ProductsList.ProductReviewsList, ProductsList.ProductReviewsList.ProductIdDetail, ProductsList.CollectionProductsList, ProductsList.CollectionProductsList.CollectionIdDetail, ProductsList.CollectionProductsList.ProductIdDetail, ProductsList.ProductTagsList, ProductsList.ProductTagsList.ProductIdDetail, ProductsList.ProductTagsList.TagIdDetail, ProductsList.PriceHistoryList, ProductsList.PriceHistoryList.ProductIdDetail, ProductsList.BrandIdDetail, ProductsList.BrandIdDetail.ProductsList, ProductsList.CategoryIdDetail, ProductsList.CategoryIdDetail.CategoriesList, ProductsList.CategoryIdDetail.ProductsList, ProductsList.CategoryIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |
| `logoUrl` | `string` | No | Filter by logo_url |
| `website` | `string` | No | Filter by website |
| `description` | `string` | No | Filter by description |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |

**Response:** `PaginationResponse<Brands>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/brands/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Brands (POST)

```
POST /api/v1/brands/pagination
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

**Response:** `PaginationResponse<Brands>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/brands/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Brands

```
POST /api/v1/brands/
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
  logoUrl?: string  // e.g. example_logo_url
  website?: string  // e.g. example_website
  description?: string  // e.g. example_description
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
}
```

**Response:** `Brands`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/brands/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Brands

```
POST /api/v1/brands/bulk/
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
  logoUrl?: string  // e.g. example_logo_url
  website?: string  // e.g. example_website
  description?: string  // e.g. example_description
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
}
```

**Response:** `Brands[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/brands/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Brands by ID

```
GET /api/v1/brands/with-id/:id
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

**Response:** `Brands`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/brands/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Brands

```
PUT /api/v1/brands/with-id/:id
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
  logoUrl?: string
  website?: string
  description?: string
  organizationId?: string
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
  "http://localhost:3000/api/v1/brands/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Brands

```
PUT /api/v1/brands/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: BrandsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/brands/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Brands

```
DELETE /api/v1/brands/with-id/:id
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
  "http://localhost:3000/api/v1/brands/with-id/:id"
```

</details>

---

:::


::::
