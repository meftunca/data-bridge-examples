---
title: ProductVariants
---

# ProductVariants

**Table:** `catalog.product_variants`

**Base path:** `/product-variants`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `products` | `product_id` | `product_variants_product_id_fkey` | [Products](./products) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `inventory` | `variant_id` | `inventory_variant_id_fkey` | [Inventory](./inventory) |
| `cart_items` | `variant_id` | `cart_items_variant_id_fkey` | [CartItems](./cart_items) |
| `order_items` | `variant_id` | `order_items_variant_id_fkey` | [OrderItems](./order_items) |


## Entity Relationship Diagram

```mermaid
erDiagram
    ProductVariants }o--|| Products : "FK"
```

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `products` |
| 4 | `sku` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 5 | `price_override` | `numeric` | `float64` | `number` | YES | - | - | - |
| 6 | `attributes` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 7 | `is_active` | `boolean` | `bool` | `boolean` | NO | `true` | - | - |
| 8 | `stock_quantity` | `integer` | `int` | `number` | NO | `0` | - | - |
| 9 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 10 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `product_id` | `products` | `product_variants_product_id_fkey` |

## Unique Keys

- `sku` (`text`)


## Go Generated Code

> 📂 Source: [📄 `ProductVariants.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductVariants.go) · [📄 `ProductVariants.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductVariants.go) · [📄 `ProductVariants.go`](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/controllers/ProductVariants.go)

### Structs

::::tabs

:::tab Form

#### ProductVariantsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductVariants.go#:~:text=type%20ProductVariantsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `Sku` | `string` | `sku` | NO |
| `PriceOverride` | `*float64` | `priceOverride` | YES |
| `Attributes` | `json.RawMessage` | `attributes` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `StockQuantity` | `int` | `stockQuantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Model

#### ProductVariants [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductVariants.go#:~:text=type%20ProductVariants%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `Sku` | `string` | `sku` | NO |
| `PriceOverride` | `*float64` | `priceOverride` | YES |
| `Attributes` | `json.RawMessage` | `attributes` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `StockQuantity` | `int` | `stockQuantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Edit

#### ProductVariantsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductVariants.go#:~:text=type%20ProductVariantsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `Sku` | `*string` | `sku` | YES |
| `PriceOverride` | `*float64` | `priceOverride` | YES |
| `Attributes` | `*json.RawMessage` | `attributes` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `StockQuantity` | `*int` | `stockQuantity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Filter

#### ProductVariantsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductVariants.go#:~:text=type%20ProductVariantsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `Sku` | `*string` | `sku` | YES |
| `PriceOverride` | `*float64` | `priceOverride` | YES |
| `Attributes` | `*json.RawMessage` | `attributes` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `StockQuantity` | `*int` | `stockQuantity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Page

#### ProductVariantsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductVariants.go#:~:text=type%20ProductVariantsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `Sku` | `string` | `sku` | NO |
| `PriceOverride` | `*float64` | `priceOverride` | YES |
| `Attributes` | `json.RawMessage` | `attributes` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `StockQuantity` | `int` | `stockQuantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab BatchUpdate

#### ProductVariantsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/structures/ProductVariants.go#:~:text=type%20ProductVariantsBatchUpdate%20struct)

```go
type ProductVariantsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductVariants.go#:~:text=)%20CreateProductVariants() | `(ProductVariantsService) CreateProductVariants(data ProductVariantsForm) (ProductVariantsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductVariants.go#:~:text=)%20CreateProductVariantsMultiple() | `(ProductVariantsService) CreateProductVariantsMultiple(data []ProductVariantsForm) ([]ProductVariantsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductVariants.go#:~:text=)%20UpdateProductVariants() | `(ProductVariantsService) UpdateProductVariants(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductVariants.go#:~:text=)%20UpdateProductVariantsMultiple() | `(ProductVariantsService) UpdateProductVariantsMultiple(data []ProductVariantsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//catalog/services/ProductVariants.go#:~:text=)%20DeleteProductVariants() | `(ProductVariantsService) DeleteProductVariants(id uuid.UUID) error` |

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/product-variants/` | Search with query params |
| `GET` | `/product-variants/pagination` | Paginated listing |
| `POST` | `/product-variants/` | Create single record |
| `POST` | `/product-variants/bulk/` | Create multiple records |
| `PUT` | `/product-variants/bulk/` | Batch update |
| `GET` | `/product-variants/with-id/:id` | Get by ID |
| `PUT` | `/product-variants/with-id/:id` | Update by ID |
| `DELETE` | `/product-variants/with-id/:id` | Delete by ID |

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
export interface ProductVariants {
  id: string;
  name: string;
  productId: string;
  sku: string;
  priceOverride?: number;
  attributes: Record<string, unknown>;
  isActive: boolean;
  stockQuantity: number;
  createdAt: string;
  updatedAt: string;
}

export interface ProductVariantsForm {
  name: string;
  productId: string;
  sku: string;
  priceOverride?: number;
  attributes: Record<string, unknown>;
  isActive: boolean;
  stockQuantity: number;
  createdAt: string;
  updatedAt: string;
}

export interface ProductVariantsEdit {
  id: string;
  name: string;
  productId: string;
  sku: string;
  priceOverride?: number;
  attributes: Record<string, unknown>;
  isActive: boolean;
  stockQuantity: number;
  createdAt: string;
  updatedAt: string;
}

export interface ProductVariantsPage {
  data: ProductVariants[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type ProductVariantsPathQuery = {
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

const ProductVariantsKeys = {
  all: ["product_variants"] as const,
  lists: () => [...ProductVariantsKeys.all, "list"] as const,
  detail: (id: any) => [...ProductVariantsKeys.all, "detail", id] as const,
} as const;

export function useProductVariantsList(query?: ProductVariantsPathQuery) {
  return useQuery({
    queryKey: [...ProductVariantsKeys.lists(), query],
    queryFn: () => fetch(`/product-variants/pagination`, { method: "GET" }).then(r => r.json()) as Promise<ProductVariantsPage>,
  });
}

export function useProductVariantsDetail(id: any) {
  return useQuery({
    queryKey: ProductVariantsKeys.detail(id),
    queryFn: () => fetch(`/product-variants/with-id/:id`).then(r => r.json()) as Promise<ProductVariants>,
  });
}

export function useCreateProductVariants() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: ProductVariantsForm) =>
      fetch("/product-variants/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductVariantsKeys.lists() }),
  });
}

export function useUpdateProductVariants() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: ProductVariantsEdit }) =>
      fetch(`/product-variants/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductVariantsKeys.all }),
  });
}

export function useDeleteProductVariants() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/product-variants/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ProductVariantsKeys.all }),
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

export const ProductVariantsFormSchema = z.object({
  name: z.string(),
  productId: z.string().uuid(),
  sku: z.string(),
  priceOverride: z.number().optional(),
  attributes: z.record(z.unknown()),
  isActive: z.boolean(),
  stockQuantity: z.number().int(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type ProductVariantsFormInput = z.infer<typeof ProductVariantsFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './product_variants.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search ProductVariants

```
GET /api/v1/product-variants/
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
| `sku` | `string` | No | Filter by sku |
| `priceOverride` | `number` | No | Filter by price_override |
| `attributes` | `string` | No | Filter by attributes |
| `isActive` | `boolean` | No | Filter by is_active |
| `stockQuantity` | `integer` | No | Filter by stock_quantity |

**Response:** `ProductVariants[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-variants/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search ProductVariants (POST)

```
POST /api/v1/product-variants/search
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

**Response:** `ProductVariants[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-variants/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate ProductVariants

```
GET /api/v1/product-variants/pagination
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
| `sku` | `string` | No | Filter by sku |
| `priceOverride` | `number` | No | Filter by price_override |
| `attributes` | `string` | No | Filter by attributes |
| `isActive` | `boolean` | No | Filter by is_active |
| `stockQuantity` | `integer` | No | Filter by stock_quantity |

**Response:** `PaginationResponse<ProductVariants>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-variants/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate ProductVariants (POST)

```
POST /api/v1/product-variants/pagination
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

**Response:** `PaginationResponse<ProductVariants>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-variants/pagination"
```

</details>

---

:::tab Create

#### <Badge type="tip" text="POST" /> Create ProductVariants

```
POST /api/v1/product-variants/
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
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  sku: string  // e.g. example_sku
  priceOverride?: number  // e.g. 99.99
  attributes?: Record<string, unknown>  // e.g. map[]
  isActive?: boolean  // e.g. true
  stockQuantity?: number  // e.g. 1
}
```

**Response:** `ProductVariants`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-variants/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create ProductVariants

```
POST /api/v1/product-variants/bulk/
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
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  sku: string  // e.g. example_sku
  priceOverride?: number  // e.g. 99.99
  attributes?: Record<string, unknown>  // e.g. map[]
  isActive?: boolean  // e.g. true
  stockQuantity?: number  // e.g. 1
}
```

**Response:** `ProductVariants[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-variants/bulk/"
```

</details>

---

:::tab Find & Update

#### <Badge type="info" text="GET" /> Find ProductVariants by ID

```
GET /api/v1/product-variants/with-id/:id
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

**Response:** `ProductVariants`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/product-variants/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update ProductVariants

```
PUT /api/v1/product-variants/with-id/:id
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
  sku?: string
  priceOverride?: number
  attributes?: Record<string, unknown>
  isActive?: boolean
  stockQuantity?: number
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
  "http://localhost:3000/api/v1/product-variants/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update ProductVariants

```
PUT /api/v1/product-variants/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: ProductVariantsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/product-variants/bulk/"
```

</details>

---

:::tab Delete

#### <Badge type="danger" text="DELETE" /> Delete ProductVariants

```
DELETE /api/v1/product-variants/with-id/:id
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
  "http://localhost:3000/api/v1/product-variants/with-id/:id"
```

</details>

---

::::


::::
