---
title: ProductReviews
---

# ProductReviews

**Table:** `catalog.product_reviews`

**Base path:** `/product-reviews`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `products` | `product_id` | `product_reviews_product_id_fkey` | [Products](./products) |
| `users` | `user_id` | `product_reviews_user_id_fkey` | [Users](./users) |


## Entity Relationship Diagram

erDiagram
    ProductReviews }o--|| Products : "FK"
    ProductReviews }o--|| Users : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `products` |
| 4 | `user_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `users` |
| 5 | `rating` | `smallint` | `int` | `number` | NO | - | - | - |
| 6 | `title` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `body` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 8 | `is_verified` | `boolean` | `bool` | `boolean` | NO | `false` | - | - |
| 9 | `helpful_count` | `integer` | `int` | `number` | NO | `0` | - | - |
| 10 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 11 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `product_id` | `products` | `product_reviews_product_id_fkey` |
| `user_id` | `users` | `product_reviews_user_id_fkey` |


## Go Generated Code

> 📂 Source: [📄 `ProductReviews.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductReviews.go) · [📄 `ProductReviews.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductReviews.go) · [📄 `ProductReviews.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/ProductReviews.go)

### Structs

:::tabs

== Form

#### ProductReviewsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductReviews.go#:~:text=type%20ProductReviewsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `Rating` | `int` | `rating` | NO |
| `Title` | `string` | `title` | NO |
| `Body` | `string` | `body` | NO |
| `IsVerified` | `bool` | `isVerified` | NO |
| `HelpfulCount` | `int` | `helpfulCount` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### ProductReviews [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductReviews.go#:~:text=type%20ProductReviews%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `Rating` | `int` | `rating` | NO |
| `Title` | `string` | `title` | NO |
| `Body` | `string` | `body` | NO |
| `IsVerified` | `bool` | `isVerified` | NO |
| `HelpfulCount` | `int` | `helpfulCount` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### ProductReviewsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductReviews.go#:~:text=type%20ProductReviewsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Rating` | `*int` | `rating` | YES |
| `Title` | `*string` | `title` | YES |
| `Body` | `*string` | `body` | YES |
| `IsVerified` | `*bool` | `isVerified` | YES |
| `HelpfulCount` | `*int` | `helpfulCount` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### ProductReviewsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductReviews.go#:~:text=type%20ProductReviewsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Rating` | `*int` | `rating` | YES |
| `Title` | `*string` | `title` | YES |
| `Body` | `*string` | `body` | YES |
| `IsVerified` | `*bool` | `isVerified` | YES |
| `HelpfulCount` | `*int` | `helpfulCount` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### ProductReviewsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductReviews.go#:~:text=type%20ProductReviewsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `Rating` | `int` | `rating` | NO |
| `Title` | `string` | `title` | NO |
| `Body` | `string` | `body` | NO |
| `IsVerified` | `bool` | `isVerified` | NO |
| `HelpfulCount` | `int` | `helpfulCount` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### ProductReviewsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductReviews.go#:~:text=type%20ProductReviewsBatchUpdate%20struct)

```go
type ProductReviewsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductReviews.go#:~:text=)%20CreateProductReviews() | `(ProductReviewsService) CreateProductReviews(data ProductReviewsForm) (ProductReviewsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductReviews.go#:~:text=)%20CreateProductReviewsMultiple() | `(ProductReviewsService) CreateProductReviewsMultiple(data []ProductReviewsForm) ([]ProductReviewsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductReviews.go#:~:text=)%20UpdateProductReviews() | `(ProductReviewsService) UpdateProductReviews(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductReviews.go#:~:text=)%20UpdateProductReviewsMultiple() | `(ProductReviewsService) UpdateProductReviewsMultiple(data []ProductReviewsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductReviews.go#:~:text=)%20DeleteProductReviews() | `(ProductReviewsService) DeleteProductReviews(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/product-reviews/` | Search with query params |
| `GET` | `/product-reviews/pagination` | Paginated listing |
| `POST` | `/product-reviews/` | Create single record |
| `POST` | `/product-reviews/bulk/` | Create multiple records |
| `PUT` | `/product-reviews/bulk/` | Batch update |
| `GET` | `/product-reviews/with-id/:id` | Get by ID |
| `PUT` | `/product-reviews/with-id/:id` | Update by ID |
| `DELETE` | `/product-reviews/with-id/:id` | Delete by ID |

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
export interface ProductReviews {
  id: string;
  name: string;
  productId: string;
  userId: string;
  rating: number;
  title: string;
  body: string;
  isVerified: boolean;
  helpfulCount: number;
  createdAt: string;
  updatedAt: string;
}

export interface ProductReviewsForm {
  name: string;
  productId: string;
  userId: string;
  rating: number;
  title: string;
  body: string;
  isVerified: boolean;
  helpfulCount: number;
  createdAt: string;
  updatedAt: string;
}

export interface ProductReviewsEdit {
  id: string;
  name: string;
  productId: string;
  userId: string;
  rating: number;
  title: string;
  body: string;
  isVerified: boolean;
  helpfulCount: number;
  createdAt: string;
  updatedAt: string;
}

export interface ProductReviewsPage {
  data: ProductReviews[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type ProductReviewsPathQuery = {
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

const ProductReviewsKeys = {
  all: ["product_reviews"] as const,
  lists: () => [...ProductReviewsKeys.all, "list"] as const,
  detail: (id: any) => [...ProductReviewsKeys.all, "detail", id] as const,
} as const;

export function useProductReviewsList(query?: ProductReviewsPathQuery) {
  return useQuery({
    queryKey: [...ProductReviewsKeys.lists(), query],
    queryFn: () => fetch(`/product-reviews/pagination`, { method: "GET" }).then(r => r.json()) as Promise<ProductReviewsPage>,
  });
}

export function useProductReviewsDetail(id: any) {
  return useQuery({
    queryKey: ProductReviewsKeys.detail(id),
    queryFn: () => fetch(`/product-reviews/with-id/:id`).then(r => r.json()) as Promise<ProductReviews>,
  });
}

export function useCreateProductReviews() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: ProductReviewsForm) =>
      fetch("/product-reviews/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductReviewsKeys.lists() }),
  });
}

export function useUpdateProductReviews() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: ProductReviewsEdit }) =>
      fetch(`/product-reviews/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductReviewsKeys.all }),
  });
}

export function useDeleteProductReviews() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/product-reviews/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductReviewsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const ProductReviewsFormSchema = z.object({
  name: z.string(),
  productId: z.string().uuid(),
  userId: z.string().uuid(),
  rating: z.number().int(),
  title: z.string(),
  body: z.string(),
  isVerified: z.boolean(),
  helpfulCount: z.number().int(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type ProductReviewsFormInput = z.infer<typeof ProductReviewsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './product_reviews.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search ProductReviews

```
GET /api/v1/product-reviews/
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
| `userId` | `string (uuid)` | No | Filter by user_id |
| `rating` | `integer` | No | Filter by rating |
| `title` | `string` | No | Filter by title |
| `body` | `string` | No | Filter by body |
| `isVerified` | `boolean` | No | Filter by is_verified |
| `helpfulCount` | `integer` | No | Filter by helpful_count |

**Response:** `ProductReviews[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-reviews/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search ProductReviews (POST)

```
POST /api/v1/product-reviews/search
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

**Response:** `ProductReviews[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-reviews/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate ProductReviews

```
GET /api/v1/product-reviews/pagination
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
| `userId` | `string (uuid)` | No | Filter by user_id |
| `rating` | `integer` | No | Filter by rating |
| `title` | `string` | No | Filter by title |
| `body` | `string` | No | Filter by body |
| `isVerified` | `boolean` | No | Filter by is_verified |
| `helpfulCount` | `integer` | No | Filter by helpful_count |

**Response:** `PaginationResponse<ProductReviews>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-reviews/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate ProductReviews (POST)

```
POST /api/v1/product-reviews/pagination
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

**Response:** `PaginationResponse<ProductReviews>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-reviews/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create ProductReviews

```
POST /api/v1/product-reviews/
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
  userId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  rating: number  // e.g. 1
  title?: string  // e.g. example_title
  body?: string  // e.g. example_body
  isVerified?: boolean  // e.g. true
  helpfulCount?: number  // e.g. 1
}
```

**Response:** `ProductReviews`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-reviews/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create ProductReviews

```
POST /api/v1/product-reviews/bulk/
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
  userId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  rating: number  // e.g. 1
  title?: string  // e.g. example_title
  body?: string  // e.g. example_body
  isVerified?: boolean  // e.g. true
  helpfulCount?: number  // e.g. 1
}
```

**Response:** `ProductReviews[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-reviews/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find ProductReviews by ID

```
GET /api/v1/product-reviews/with-id/:id
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

**Response:** `ProductReviews`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-reviews/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update ProductReviews

```
PUT /api/v1/product-reviews/with-id/:id
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
  userId?: string
  rating?: number
  title?: string
  body?: string
  isVerified?: boolean
  helpfulCount?: number
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
  "http://localhost:3000/api/v1/product-reviews/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update ProductReviews

```
PUT /api/v1/product-reviews/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: ProductReviewsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-reviews/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete ProductReviews

```
DELETE /api/v1/product-reviews/with-id/:id
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
  "http://localhost:3000/api/v1/product-reviews/with-id/:id"
```

</details>

---

:::


::::
