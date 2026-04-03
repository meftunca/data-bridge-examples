---
title: CartItems
---

# CartItems

**Table:** `orders.cart_items`

**Base path:** `/cart-items`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `carts` | `cart_id` | `cart_items_cart_id_fkey` | [Carts](./carts) |
| `products` | `product_id` | `cart_items_product_id_fkey` | [Products](./products) |
| `product_variants` | `variant_id` | `cart_items_variant_id_fkey` | [ProductVariants](./product_variants) |


## Entity Relationship Diagram

```mermaid
erDiagram
    CartItems }o--|| Carts : "FK"
    CartItems }o--|| Products : "FK"
    CartItems }o--|| ProductVariants : "FK"
```

::::tabs

=== FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `cart_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `carts` |
| 4 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `products` |
| 5 | `variant_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `product_variants` |
| 6 | `quantity` | `integer` | `int` | `number` | NO | `1` | - | - |
| 7 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 8 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `cart_id` | `carts` | `cart_items_cart_id_fkey` |
| `product_id` | `products` | `cart_items_product_id_fkey` |
| `variant_id` | `product_variants` | `cart_items_variant_id_fkey` |


## Go Generated Code

> 📂 Source: [📄 `CartItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/CartItems.go) · [📄 `CartItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/CartItems.go) · [📄 `CartItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/controllers/CartItems.go)

### Structs

:::tabs

== Form

#### CartItemsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/CartItems.go#:~:text=type%20CartItemsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `CartId` | `uuid.UUID` | `cartId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### CartItems [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/CartItems.go#:~:text=type%20CartItems%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `CartId` | `uuid.UUID` | `cartId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### CartItemsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/CartItems.go#:~:text=type%20CartItemsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `CartId` | `*uuid.UUID` | `cartId` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### CartItemsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/CartItems.go#:~:text=type%20CartItemsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `CartId` | `*uuid.UUID` | `cartId` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### CartItemsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/CartItems.go#:~:text=type%20CartItemsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `CartId` | `uuid.UUID` | `cartId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### CartItemsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/CartItems.go#:~:text=type%20CartItemsBatchUpdate%20struct)

```go
type CartItemsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/CartItems.go#:~:text=%29%20CreateCartItems%28%29) | `(CartItemsService) CreateCartItems(data CartItemsForm) (CartItemsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/CartItems.go#:~:text=%29%20CreateCartItemsMultiple%28%29) | `(CartItemsService) CreateCartItemsMultiple(data []CartItemsForm) ([]CartItemsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/CartItems.go#:~:text=%29%20UpdateCartItems%28%29) | `(CartItemsService) UpdateCartItems(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/CartItems.go#:~:text=%29%20UpdateCartItemsMultiple%28%29) | `(CartItemsService) UpdateCartItemsMultiple(data []CartItemsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/CartItems.go#:~:text=%29%20DeleteCartItems%28%29) | `(CartItemsService) DeleteCartItems(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/cart-items/` | Search with query params |
| `GET` | `/cart-items/pagination` | Paginated listing |
| `POST` | `/cart-items/` | Create single record |
| `POST` | `/cart-items/bulk/` | Create multiple records |
| `PUT` | `/cart-items/bulk/` | Batch update |
| `GET` | `/cart-items/with-id/:id` | Get by ID |
| `PUT` | `/cart-items/with-id/:id` | Update by ID |
| `DELETE` | `/cart-items/with-id/:id` | Delete by ID |

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
| `customer_total_spent` | `p_customer_id uuid` | `numeric` | `/rpc/customer_total_spent` |
| `orders_by_status` | `p_status text` | `integer` | `/rpc/orders_by_status` |
| `total_revenue` | - | `numeric` | `/rpc/total_revenue` |


=== Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export interface CartItems {
  id: string;
  name: string;
  cartId: string;
  productId: string;
  variantId?: string;
  quantity: number;
  createdAt: string;
  updatedAt: string;
}

export interface CartItemsForm {
  name: string;
  cartId: string;
  productId: string;
  variantId?: string;
  quantity: number;
  createdAt: string;
  updatedAt: string;
}

export interface CartItemsEdit {
  id: string;
  name: string;
  cartId: string;
  productId: string;
  variantId?: string;
  quantity: number;
  createdAt: string;
  updatedAt: string;
}

export interface CartItemsPage {
  data: CartItems[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type CartItemsPathQuery = {
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

const CartItemsKeys = {
  all: ["cart_items"] as const,
  lists: () => [...CartItemsKeys.all, "list"] as const,
  detail: (id: any) => [...CartItemsKeys.all, "detail", id] as const,
} as const;

export function useCartItemsList(query?: CartItemsPathQuery) {
  return useQuery({
    queryKey: [...CartItemsKeys.lists(), query],
    queryFn: () => fetch(`/cart-items/pagination`, { method: "GET" }).then(r => r.json()) as Promise<CartItemsPage>,
  });
}

export function useCartItemsDetail(id: any) {
  return useQuery({
    queryKey: CartItemsKeys.detail(id),
    queryFn: () => fetch(`/cart-items/with-id/:id`).then(r => r.json()) as Promise<CartItems>,
  });
}

export function useCreateCartItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: CartItemsForm) =>
      fetch("/cart-items/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CartItemsKeys.lists() }),
  });
}

export function useUpdateCartItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: CartItemsEdit }) =>
      fetch(`/cart-items/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CartItemsKeys.all }),
  });
}

export function useDeleteCartItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/cart-items/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CartItemsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const CartItemsFormSchema = z.object({
  name: z.string(),
  cartId: z.string().uuid(),
  productId: z.string().uuid(),
  variantId: z.string().uuid().optional(),
  quantity: z.number().int(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type CartItemsFormInput = z.infer<typeof CartItemsFormSchema>;

```

:::


=== API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './cart_items.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search CartItems

```
GET /api/v1/cart-items/
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
| `preloads` | `string` | No | Available: CartIdDetail, CartIdDetail.CartItemsList, CartIdDetail.CartItemsList.CartIdDetail, CartIdDetail.CustomerIdDetail, CartIdDetail.CustomerIdDetail.OrdersList, CartIdDetail.CustomerIdDetail.CartsList |
| `joins` | `string` | No | Available: Carts, Carts.Customers, Carts.Customers.Users, Carts.Customers.Organizations, Products, ProductVariants |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `cartId` | `string (uuid)` | No | Filter by cart_id |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `variantId` | `string (uuid)` | No | Filter by variant_id |
| `quantity` | `integer` | No | Filter by quantity |

**Response:** `CartItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/cart-items/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search CartItems (POST)

```
POST /api/v1/cart-items/search
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

**Response:** `CartItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/cart-items/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate CartItems

```
GET /api/v1/cart-items/pagination
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
| `preloads` | `string` | No | Available: CartIdDetail, CartIdDetail.CartItemsList, CartIdDetail.CartItemsList.CartIdDetail, CartIdDetail.CustomerIdDetail, CartIdDetail.CustomerIdDetail.OrdersList, CartIdDetail.CustomerIdDetail.CartsList |
| `joins` | `string` | No | Available: Carts, Carts.Customers, Carts.Customers.Users, Carts.Customers.Organizations, Products, ProductVariants |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `cartId` | `string (uuid)` | No | Filter by cart_id |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `variantId` | `string (uuid)` | No | Filter by variant_id |
| `quantity` | `integer` | No | Filter by quantity |

**Response:** `PaginationResponse<CartItems>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/cart-items/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate CartItems (POST)

```
POST /api/v1/cart-items/pagination
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

**Response:** `PaginationResponse<CartItems>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/cart-items/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create CartItems

```
POST /api/v1/cart-items/
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
  cartId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  variantId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
}
```

**Response:** `CartItems`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/cart-items/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create CartItems

```
POST /api/v1/cart-items/bulk/
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
  cartId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  variantId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
}
```

**Response:** `CartItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/cart-items/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find CartItems by ID

```
GET /api/v1/cart-items/with-id/:id
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

**Response:** `CartItems`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/cart-items/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update CartItems

```
PUT /api/v1/cart-items/with-id/:id
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
  cartId?: string
  productId?: string
  variantId?: string
  quantity?: number
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
  "http://localhost:3000/api/v1/cart-items/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update CartItems

```
PUT /api/v1/cart-items/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: CartItemsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/cart-items/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete CartItems

```
DELETE /api/v1/cart-items/with-id/:id
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
  "http://localhost:3000/api/v1/cart-items/with-id/:id"
```

</details>

---

:::


::::
