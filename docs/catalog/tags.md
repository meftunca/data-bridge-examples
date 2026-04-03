---
title: Tags
---

# Tags

**Table:** `catalog.tags`

**Base path:** `/tags`

## Related Tables

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `product_tags` | `tag_id` | `product_tags_tag_id_fkey` | [ProductTags](./product_tags) |


## Entity Relationship Diagram

erDiagram
    Tags ||--o{ ProductTags : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 3 | `slug` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 4 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Unique Keys

- `name` (`text`)
- `slug` (`text`)


## Go Generated Code

> 📂 Source: [📄 `Tags.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Tags.go) · [📄 `Tags.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Tags.go) · [📄 `Tags.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/Tags.go)

### Structs

:::tabs

== Form

#### TagsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Tags.go#:~:text=type%20TagsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### Tags [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Tags.go#:~:text=type%20Tags%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### TagsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Tags.go#:~:text=type%20TagsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### TagsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Tags.go#:~:text=type%20TagsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### TagsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Tags.go#:~:text=type%20TagsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### TagsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/Tags.go#:~:text=type%20TagsBatchUpdate%20struct)

```go
type TagsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Tags.go#:~:text=)%20CreateTags() | `(TagsService) CreateTags(data TagsForm) (TagsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Tags.go#:~:text=)%20CreateTagsMultiple() | `(TagsService) CreateTagsMultiple(data []TagsForm) ([]TagsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Tags.go#:~:text=)%20UpdateTags() | `(TagsService) UpdateTags(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Tags.go#:~:text=)%20UpdateTagsMultiple() | `(TagsService) UpdateTagsMultiple(data []TagsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/Tags.go#:~:text=)%20DeleteTags() | `(TagsService) DeleteTags(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/tags/` | Search with query params |
| `GET` | `/tags/pagination` | Paginated listing |
| `POST` | `/tags/` | Create single record |
| `POST` | `/tags/bulk/` | Create multiple records |
| `PUT` | `/tags/bulk/` | Batch update |
| `GET` | `/tags/with-id/:id` | Get by ID |
| `PUT` | `/tags/with-id/:id` | Update by ID |
| `DELETE` | `/tags/with-id/:id` | Delete by ID |

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
export interface Tags {
  id: string;
  name: string;
  slug: string;
  createdAt: string;
}

export interface TagsForm {
  name: string;
  slug: string;
  createdAt: string;
}

export interface TagsEdit {
  id: string;
  name: string;
  slug: string;
  createdAt: string;
}

export interface TagsPage {
  data: Tags[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type TagsPathQuery = {
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

const TagsKeys = {
  all: ["tags"] as const,
  lists: () => [...TagsKeys.all, "list"] as const,
  detail: (id: any) => [...TagsKeys.all, "detail", id] as const,
} as const;

export function useTagsList(query?: TagsPathQuery) {
  return useQuery({
    queryKey: [...TagsKeys.lists(), query],
    queryFn: () => fetch(`/tags/pagination`, { method: "GET" }).then(r => r.json()) as Promise<TagsPage>,
  });
}

export function useTagsDetail(id: any) {
  return useQuery({
    queryKey: TagsKeys.detail(id),
    queryFn: () => fetch(`/tags/with-id/:id`).then(r => r.json()) as Promise<Tags>,
  });
}

export function useCreateTags() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: TagsForm) =>
      fetch("/tags/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: TagsKeys.lists() }),
  });
}

export function useUpdateTags() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: TagsEdit }) =>
      fetch(`/tags/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: TagsKeys.all }),
  });
}

export function useDeleteTags() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/tags/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: TagsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const TagsFormSchema = z.object({
  name: z.string(),
  slug: z.string(),
  createdAt: z.string().datetime(),
});

export type TagsFormInput = z.infer<typeof TagsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './tags.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Tags

```
GET /api/v1/tags/
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
| `preloads` | `string` | No | Available: ProductTagsList, ProductTagsList.ProductIdDetail, ProductTagsList.ProductIdDetail.ProductVariantsList, ProductTagsList.ProductIdDetail.ProductMediaList, ProductTagsList.ProductIdDetail.ProductReviewsList, ProductTagsList.ProductIdDetail.CollectionProductsList, ProductTagsList.ProductIdDetail.ProductTagsList, ProductTagsList.ProductIdDetail.PriceHistoryList, ProductTagsList.ProductIdDetail.BrandIdDetail, ProductTagsList.ProductIdDetail.CategoryIdDetail, ProductTagsList.TagIdDetail, ProductTagsList.TagIdDetail.ProductTagsList |
| `joins` | `string` | No | Comma-separated relations to join |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |

**Response:** `Tags[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/tags/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Tags (POST)

```
POST /api/v1/tags/search
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

**Response:** `Tags[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/tags/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Tags

```
GET /api/v1/tags/pagination
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
| `preloads` | `string` | No | Available: ProductTagsList, ProductTagsList.ProductIdDetail, ProductTagsList.ProductIdDetail.ProductVariantsList, ProductTagsList.ProductIdDetail.ProductMediaList, ProductTagsList.ProductIdDetail.ProductReviewsList, ProductTagsList.ProductIdDetail.CollectionProductsList, ProductTagsList.ProductIdDetail.ProductTagsList, ProductTagsList.ProductIdDetail.PriceHistoryList, ProductTagsList.ProductIdDetail.BrandIdDetail, ProductTagsList.ProductIdDetail.CategoryIdDetail, ProductTagsList.TagIdDetail, ProductTagsList.TagIdDetail.ProductTagsList |
| `joins` | `string` | No | Comma-separated relations to join |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |

**Response:** `PaginationResponse<Tags>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/tags/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Tags (POST)

```
POST /api/v1/tags/pagination
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

**Response:** `PaginationResponse<Tags>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/tags/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Tags

```
POST /api/v1/tags/
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
}
```

**Response:** `Tags`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/tags/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Tags

```
POST /api/v1/tags/bulk/
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
}
```

**Response:** `Tags[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/tags/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Tags by ID

```
GET /api/v1/tags/with-id/:id
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

**Response:** `Tags`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/tags/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Tags

```
PUT /api/v1/tags/with-id/:id
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
  "http://localhost:3000/api/v1/tags/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Tags

```
PUT /api/v1/tags/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: TagsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/tags/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Tags

```
DELETE /api/v1/tags/with-id/:id
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
  "http://localhost:3000/api/v1/tags/with-id/:id"
```

</details>

---

:::


::::
