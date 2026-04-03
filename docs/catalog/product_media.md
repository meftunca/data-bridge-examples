---
title: ProductMedia
---

# ProductMedia

**Table:** `catalog.product_media`

**Base path:** `/product-media`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `products` | `product_id` | `product_media_product_id_fkey` | [Products](./products) |


## Entity Relationship Diagram

```mermaid
erDiagram
    ProductMedia }o--|| Products : "FK"
```

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `products` |
| 4 | `media_type` | `USER-DEFINED` | `CatalogMediaType` | `"image" \| "video" \| "document" \| "audio" \| "3d_model"` | NO | `'image'::catalog.media_type` | - | - |
| 5 | `url` | `text` | `string` | `string` | NO | - | - | - |
| 6 | `alt_text` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `sort_order` | `integer` | `int` | `number` | NO | `0` | - | - |
| 8 | `is_primary` | `boolean` | `bool` | `boolean` | NO | `false` | - | - |
| 9 | `metadata` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 10 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 11 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `product_id` | `products` | `product_media_product_id_fkey` |

## Enum Types

### MediaType

| Value | Go Constant |
|-------|-------------|
| `image` | `CatalogMediaTypeImage` |
| `video` | `CatalogMediaTypeVideo` |
| `document` | `CatalogMediaTypeDocument` |
| `audio` | `CatalogMediaTypeAudio` |
| `3d_model` | `CatalogMediaType3DModel` |


## Go Generated Code

> 📂 Source: [📄 `ProductMedia.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductMedia.go) · [📄 `ProductMedia.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductMedia.go) · [📄 `ProductMedia.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/ProductMedia.go)

### Structs

::::tabs

:::tab Form

#### ProductMediaForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductMedia.go#:~:text=type%20ProductMediaForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `MediaType` | `CatalogMediaType` | `mediaType` | NO |
| `Url` | `string` | `url` | NO |
| `AltText` | `string` | `altText` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `IsPrimary` | `bool` | `isPrimary` | NO |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Model

#### ProductMedia [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductMedia.go#:~:text=type%20ProductMedia%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `MediaType` | `CatalogMediaType` | `mediaType` | NO |
| `Url` | `string` | `url` | NO |
| `AltText` | `string` | `altText` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `IsPrimary` | `bool` | `isPrimary` | NO |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Edit

#### ProductMediaEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductMedia.go#:~:text=type%20ProductMediaEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `MediaType` | `*CatalogMediaType` | `mediaType` | YES |
| `Url` | `*string` | `url` | YES |
| `AltText` | `*string` | `altText` | YES |
| `SortOrder` | `*int` | `sortOrder` | YES |
| `IsPrimary` | `*bool` | `isPrimary` | YES |
| `Metadata` | `*json.RawMessage` | `metadata` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Filter

#### ProductMediaFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductMedia.go#:~:text=type%20ProductMediaFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `MediaType` | `*CatalogMediaType` | `mediaType` | YES |
| `Url` | `*string` | `url` | YES |
| `AltText` | `*string` | `altText` | YES |
| `SortOrder` | `*int` | `sortOrder` | YES |
| `IsPrimary` | `*bool` | `isPrimary` | YES |
| `Metadata` | `*json.RawMessage` | `metadata` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Page

#### ProductMediaPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductMedia.go#:~:text=type%20ProductMediaPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `MediaType` | `CatalogMediaType` | `mediaType` | NO |
| `Url` | `string` | `url` | NO |
| `AltText` | `string` | `altText` | NO |
| `SortOrder` | `int` | `sortOrder` | NO |
| `IsPrimary` | `bool` | `isPrimary` | NO |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab BatchUpdate

#### ProductMediaBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductMedia.go#:~:text=type%20ProductMediaBatchUpdate%20struct)

```go
type ProductMediaBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductMedia.go#:~:text=)%20CreateProductMedia() | `(ProductMediaService) CreateProductMedia(data ProductMediaForm) (ProductMediaForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductMedia.go#:~:text=)%20CreateProductMediaMultiple() | `(ProductMediaService) CreateProductMediaMultiple(data []ProductMediaForm) ([]ProductMediaForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductMedia.go#:~:text=)%20UpdateProductMedia() | `(ProductMediaService) UpdateProductMedia(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductMedia.go#:~:text=)%20UpdateProductMediaMultiple() | `(ProductMediaService) UpdateProductMediaMultiple(data []ProductMediaBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductMedia.go#:~:text=)%20DeleteProductMedia() | `(ProductMediaService) DeleteProductMedia(id uuid.UUID) error` |

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/product-media/` | Search with query params |
| `GET` | `/product-media/pagination` | Paginated listing |
| `POST` | `/product-media/` | Create single record |
| `POST` | `/product-media/bulk/` | Create multiple records |
| `PUT` | `/product-media/bulk/` | Batch update |
| `GET` | `/product-media/with-id/:id` | Get by ID |
| `PUT` | `/product-media/with-id/:id` | Update by ID |
| `DELETE` | `/product-media/with-id/:id` | Delete by ID |

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
export type CatalogMediaType =
  | "image"
  | "video"
  | "document"
  | "audio"
  | "3d_model"

export const CatalogMediaTypeValues = ["image", "video", "document", "audio", "3d_model"] as const;

export interface ProductMedia {
  id: string;
  name: string;
  productId: string;
  mediaType: CatalogMediaType;
  url: string;
  altText: string;
  sortOrder: number;
  isPrimary: boolean;
  metadata: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface ProductMediaForm {
  name: string;
  productId: string;
  mediaType: CatalogMediaType;
  url: string;
  altText: string;
  sortOrder: number;
  isPrimary: boolean;
  metadata: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface ProductMediaEdit {
  id: string;
  name: string;
  productId: string;
  mediaType: CatalogMediaType;
  url: string;
  altText: string;
  sortOrder: number;
  isPrimary: boolean;
  metadata: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface ProductMediaPage {
  data: ProductMedia[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type ProductMediaPathQuery = {
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

const ProductMediaKeys = {
  all: ["product_media"] as const,
  lists: () => [...ProductMediaKeys.all, "list"] as const,
  detail: (id: any) => [...ProductMediaKeys.all, "detail", id] as const,
} as const;

export function useProductMediaList(query?: ProductMediaPathQuery) {
  return useQuery({
    queryKey: [...ProductMediaKeys.lists(), query],
    queryFn: () => fetch(`/product-media/pagination`, { method: "GET" }).then(r => r.json()) as Promise<ProductMediaPage>,
  });
}

export function useProductMediaDetail(id: any) {
  return useQuery({
    queryKey: ProductMediaKeys.detail(id),
    queryFn: () => fetch(`/product-media/with-id/:id`).then(r => r.json()) as Promise<ProductMedia>,
  });
}

export function useCreateProductMedia() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: ProductMediaForm) =>
      fetch("/product-media/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductMediaKeys.lists() }),
  });
}

export function useUpdateProductMedia() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: ProductMediaEdit }) =>
      fetch(`/product-media/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductMediaKeys.all }),
  });
}

export function useDeleteProductMedia() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/product-media/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductMediaKeys.all }),
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

const CatalogMediaTypeSchema = z.enum(["image", "video", "document", "audio", "3d_model"]);

export const ProductMediaFormSchema = z.object({
  name: z.string(),
  productId: z.string().uuid(),
  mediaType: CatalogMediaTypeSchema,
  url: z.string(),
  altText: z.string(),
  sortOrder: z.number().int(),
  isPrimary: z.boolean(),
  metadata: z.record(z.unknown()),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type ProductMediaFormInput = z.infer<typeof ProductMediaFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './product_media.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search ProductMedia

```
GET /api/v1/product-media/
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
| `joins` | `string` | No | Available: Products, Products.Brands, Products.Brands.Organizations, Products.Categories, Products.Categories.Categories, Products.Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `mediaType` | `string` | No | Filter by media_type |
| `url` | `string` | No | Filter by url |
| `altText` | `string` | No | Filter by alt_text |
| `sortOrder` | `integer` | No | Filter by sort_order |
| `isPrimary` | `boolean` | No | Filter by is_primary |
| `metadata` | `string` | No | Filter by metadata |

**Response:** `ProductMedia[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-media/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search ProductMedia (POST)

```
POST /api/v1/product-media/search
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

**Response:** `ProductMedia[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-media/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate ProductMedia

```
GET /api/v1/product-media/pagination
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
| `joins` | `string` | No | Available: Products, Products.Brands, Products.Brands.Organizations, Products.Categories, Products.Categories.Categories, Products.Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `mediaType` | `string` | No | Filter by media_type |
| `url` | `string` | No | Filter by url |
| `altText` | `string` | No | Filter by alt_text |
| `sortOrder` | `integer` | No | Filter by sort_order |
| `isPrimary` | `boolean` | No | Filter by is_primary |
| `metadata` | `string` | No | Filter by metadata |

**Response:** `PaginationResponse<ProductMedia>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-media/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate ProductMedia (POST)

```
POST /api/v1/product-media/pagination
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

**Response:** `PaginationResponse<ProductMedia>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-media/pagination"
```

</details>

---

:::tab Create

#### <Badge type="tip" text="POST" /> Create ProductMedia

```
POST /api/v1/product-media/
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
  mediaType?: "image" | "video" | "document" | "audio" | "3d_model"  // e.g. image
  url: string  // e.g. example_url
  altText?: string  // e.g. example_alt_text
  sortOrder?: number  // e.g. 1
  isPrimary?: boolean  // e.g. true
  metadata?: Record<string, unknown>  // e.g. map[]
}
```

**Response:** `ProductMedia`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-media/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create ProductMedia

```
POST /api/v1/product-media/bulk/
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
  mediaType?: "image" | "video" | "document" | "audio" | "3d_model"  // e.g. image
  url: string  // e.g. example_url
  altText?: string  // e.g. example_alt_text
  sortOrder?: number  // e.g. 1
  isPrimary?: boolean  // e.g. true
  metadata?: Record<string, unknown>  // e.g. map[]
}
```

**Response:** `ProductMedia[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-media/bulk/"
```

</details>

---

:::tab Find & Update

#### <Badge type="info" text="GET" /> Find ProductMedia by ID

```
GET /api/v1/product-media/with-id/:id
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

**Response:** `ProductMedia`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-media/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update ProductMedia

```
PUT /api/v1/product-media/with-id/:id
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
  mediaType?: "image" | "video" | "document" | "audio" | "3d_model"
  url?: string
  altText?: string
  sortOrder?: number
  isPrimary?: boolean
  metadata?: Record<string, unknown>
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
  "http://localhost:3000/api/v1/product-media/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update ProductMedia

```
PUT /api/v1/product-media/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: ProductMediaEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-media/bulk/"
```

</details>

---

:::tab Delete

#### <Badge type="danger" text="DELETE" /> Delete ProductMedia

```
DELETE /api/v1/product-media/with-id/:id
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
  "http://localhost:3000/api/v1/product-media/with-id/:id"
```

</details>

---

::::


::::
