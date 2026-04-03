---
title: OrderItems
---

# OrderItems

**Table:** `orders.order_items`

**Base path:** `/order-items`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `orders` | `order_id` | `order_items_order_id_fkey` | [Orders](./orders) |
| `products` | `product_id` | `order_items_product_id_fkey` | [Products](./products) |
| `product_variants` | `variant_id` | `order_items_variant_id_fkey` | [ProductVariants](./product_variants) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `shipment_items` | `order_item_id` | `shipment_items_order_item_id_fkey` | [ShipmentItems](./shipment_items) |


## Entity Relationship Diagram

```mermaid
erDiagram
    OrderItems }o--|| Orders : "FK"
    OrderItems }o--|| Products : "FK"
    OrderItems }o--|| ProductVariants : "FK"
```

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `order_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `orders` |
| 4 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `products` |
| 5 | `variant_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `product_variants` |
| 6 | `quantity` | `integer` | `int` | `number` | NO | `1` | - | - |
| 7 | `unit_price` | `numeric` | `float64` | `number` | NO | `0.00` | - | - |
| 8 | `total_price` | `numeric` | `float64` | `number` | NO | `0.00` | - | - |
| 9 | `discount` | `numeric` | `float64` | `number` | NO | `0.00` | - | - |
| 10 | `metadata` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 11 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 12 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `order_id` | `orders` | `order_items_order_id_fkey` |
| `product_id` | `products` | `order_items_product_id_fkey` |
| `variant_id` | `product_variants` | `order_items_variant_id_fkey` |


## Go Generated Code

> 📂 Source: [📄 `OrderItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderItems.go) · [📄 `OrderItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderItems.go) · [📄 `OrderItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/controllers/OrderItems.go)

### Structs

::::tabs

:::tab Form

#### OrderItemsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderItems.go#:~:text=type%20OrderItemsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `OrderId` | `uuid.UUID` | `orderId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `UnitPrice` | `float64` | `unitPrice` | NO |
| `TotalPrice` | `float64` | `totalPrice` | NO |
| `Discount` | `float64` | `discount` | NO |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Model

#### OrderItems [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderItems.go#:~:text=type%20OrderItems%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `OrderId` | `uuid.UUID` | `orderId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `UnitPrice` | `float64` | `unitPrice` | NO |
| `TotalPrice` | `float64` | `totalPrice` | NO |
| `Discount` | `float64` | `discount` | NO |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Edit

#### OrderItemsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderItems.go#:~:text=type%20OrderItemsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `OrderId` | `*uuid.UUID` | `orderId` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `UnitPrice` | `*float64` | `unitPrice` | YES |
| `TotalPrice` | `*float64` | `totalPrice` | YES |
| `Discount` | `*float64` | `discount` | YES |
| `Metadata` | `*json.RawMessage` | `metadata` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Filter

#### OrderItemsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderItems.go#:~:text=type%20OrderItemsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `OrderId` | `*uuid.UUID` | `orderId` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `UnitPrice` | `*float64` | `unitPrice` | YES |
| `TotalPrice` | `*float64` | `totalPrice` | YES |
| `Discount` | `*float64` | `discount` | YES |
| `Metadata` | `*json.RawMessage` | `metadata` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Page

#### OrderItemsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderItems.go#:~:text=type%20OrderItemsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `OrderId` | `uuid.UUID` | `orderId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `UnitPrice` | `float64` | `unitPrice` | NO |
| `TotalPrice` | `float64` | `totalPrice` | NO |
| `Discount` | `float64` | `discount` | NO |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab BatchUpdate

#### OrderItemsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderItems.go#:~:text=type%20OrderItemsBatchUpdate%20struct)

```go
type OrderItemsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderItems.go#:~:text=)%20CreateOrderItems() | `(OrderItemsService) CreateOrderItems(data OrderItemsForm) (OrderItemsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderItems.go#:~:text=)%20CreateOrderItemsMultiple() | `(OrderItemsService) CreateOrderItemsMultiple(data []OrderItemsForm) ([]OrderItemsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderItems.go#:~:text=)%20UpdateOrderItems() | `(OrderItemsService) UpdateOrderItems(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderItems.go#:~:text=)%20UpdateOrderItemsMultiple() | `(OrderItemsService) UpdateOrderItemsMultiple(data []OrderItemsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderItems.go#:~:text=)%20DeleteOrderItems() | `(OrderItemsService) DeleteOrderItems(id uuid.UUID) error` |

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/order-items/` | Search with query params |
| `GET` | `/order-items/pagination` | Paginated listing |
| `POST` | `/order-items/` | Create single record |
| `POST` | `/order-items/bulk/` | Create multiple records |
| `PUT` | `/order-items/bulk/` | Batch update |
| `GET` | `/order-items/with-id/:id` | Get by ID |
| `PUT` | `/order-items/with-id/:id` | Update by ID |
| `DELETE` | `/order-items/with-id/:id` | Delete by ID |

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
| `customer_total_spent` | `p_customer_id uuid` | `numeric` | `/rpc/customer_total_spent` |
| `orders_by_status` | `p_status text` | `integer` | `/rpc/orders_by_status` |
| `total_revenue` | - | `numeric` | `/rpc/total_revenue` |


:::tab Frontend

## TypeScript Types & Hooks

::::tabs

:::tab Interfaces

```typescript
export interface OrderItems {
  id: string;
  name: string;
  orderId: string;
  productId: string;
  variantId?: string;
  quantity: number;
  unitPrice: number;
  totalPrice: number;
  discount: number;
  metadata: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface OrderItemsForm {
  name: string;
  orderId: string;
  productId: string;
  variantId?: string;
  quantity: number;
  unitPrice: number;
  totalPrice: number;
  discount: number;
  metadata: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface OrderItemsEdit {
  id: string;
  name: string;
  orderId: string;
  productId: string;
  variantId?: string;
  quantity: number;
  unitPrice: number;
  totalPrice: number;
  discount: number;
  metadata: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface OrderItemsPage {
  data: OrderItems[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type OrderItemsPathQuery = {
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

const OrderItemsKeys = {
  all: ["order_items"] as const,
  lists: () => [...OrderItemsKeys.all, "list"] as const,
  detail: (id: any) => [...OrderItemsKeys.all, "detail", id] as const,
} as const;

export function useOrderItemsList(query?: OrderItemsPathQuery) {
  return useQuery({
    queryKey: [...OrderItemsKeys.lists(), query],
    queryFn: () => fetch(`/order-items/pagination`, { method: "GET" }).then(r => r.json()) as Promise<OrderItemsPage>,
  });
}

export function useOrderItemsDetail(id: any) {
  return useQuery({
    queryKey: OrderItemsKeys.detail(id),
    queryFn: () => fetch(`/order-items/with-id/:id`).then(r => r.json()) as Promise<OrderItems>,
  });
}

export function useCreateOrderItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: OrderItemsForm) =>
      fetch("/order-items/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: OrderItemsKeys.lists() }),
  });
}

export function useUpdateOrderItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: OrderItemsEdit }) =>
      fetch(`/order-items/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: OrderItemsKeys.all }),
  });
}

export function useDeleteOrderItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/order-items/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: OrderItemsKeys.all }),
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

export const OrderItemsFormSchema = z.object({
  name: z.string(),
  orderId: z.string().uuid(),
  productId: z.string().uuid(),
  variantId: z.string().uuid().optional(),
  quantity: z.number().int(),
  unitPrice: z.number(),
  totalPrice: z.number(),
  discount: z.number(),
  metadata: z.record(z.unknown()),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type OrderItemsFormInput = z.infer<typeof OrderItemsFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './order_items.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search OrderItems

```
GET /api/v1/order-items/
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
| `preloads` | `string` | No | Available: OrderIdDetail, OrderIdDetail.OrderItemsList, OrderIdDetail.OrderItemsList.OrderIdDetail, OrderIdDetail.PaymentsList, OrderIdDetail.PaymentsList.RefundsList, OrderIdDetail.PaymentsList.OrderIdDetail, OrderIdDetail.RefundsList, OrderIdDetail.RefundsList.OrderIdDetail, OrderIdDetail.RefundsList.PaymentIdDetail, OrderIdDetail.OrderStatusHistoryList, OrderIdDetail.OrderStatusHistoryList.OrderIdDetail, OrderIdDetail.CustomerIdDetail, OrderIdDetail.CustomerIdDetail.OrdersList, OrderIdDetail.CustomerIdDetail.CartsList, OrderIdDetail.CouponIdDetail, OrderIdDetail.CouponIdDetail.OrdersList |
| `joins` | `string` | No | Available: Orders, Orders.Customers, Orders.Customers.Users, Orders.Customers.Organizations, Orders.Coupons, Orders.Users, Products, ProductVariants |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `orderId` | `string (uuid)` | No | Filter by order_id |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `variantId` | `string (uuid)` | No | Filter by variant_id |
| `quantity` | `integer` | No | Filter by quantity |
| `unitPrice` | `number` | No | Filter by unit_price |
| `totalPrice` | `number` | No | Filter by total_price |
| `discount` | `number` | No | Filter by discount |
| `metadata` | `string` | No | Filter by metadata |

**Response:** `OrderItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/order-items/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search OrderItems (POST)

```
POST /api/v1/order-items/search
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

**Response:** `OrderItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-items/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate OrderItems

```
GET /api/v1/order-items/pagination
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
| `preloads` | `string` | No | Available: OrderIdDetail, OrderIdDetail.OrderItemsList, OrderIdDetail.OrderItemsList.OrderIdDetail, OrderIdDetail.PaymentsList, OrderIdDetail.PaymentsList.RefundsList, OrderIdDetail.PaymentsList.OrderIdDetail, OrderIdDetail.RefundsList, OrderIdDetail.RefundsList.OrderIdDetail, OrderIdDetail.RefundsList.PaymentIdDetail, OrderIdDetail.OrderStatusHistoryList, OrderIdDetail.OrderStatusHistoryList.OrderIdDetail, OrderIdDetail.CustomerIdDetail, OrderIdDetail.CustomerIdDetail.OrdersList, OrderIdDetail.CustomerIdDetail.CartsList, OrderIdDetail.CouponIdDetail, OrderIdDetail.CouponIdDetail.OrdersList |
| `joins` | `string` | No | Available: Orders, Orders.Customers, Orders.Customers.Users, Orders.Customers.Organizations, Orders.Coupons, Orders.Users, Products, ProductVariants |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `orderId` | `string (uuid)` | No | Filter by order_id |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `variantId` | `string (uuid)` | No | Filter by variant_id |
| `quantity` | `integer` | No | Filter by quantity |
| `unitPrice` | `number` | No | Filter by unit_price |
| `totalPrice` | `number` | No | Filter by total_price |
| `discount` | `number` | No | Filter by discount |
| `metadata` | `string` | No | Filter by metadata |

**Response:** `PaginationResponse<OrderItems>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/order-items/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate OrderItems (POST)

```
POST /api/v1/order-items/pagination
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

**Response:** `PaginationResponse<OrderItems>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-items/pagination"
```

</details>

---

:::tab Create

#### <Badge type="tip" text="POST" /> Create OrderItems

```
POST /api/v1/order-items/
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
  orderId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  variantId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
  unitPrice?: number  // e.g. 99.99
  totalPrice?: number  // e.g. 99.99
  discount?: number  // e.g. 99.99
  metadata?: Record<string, unknown>  // e.g. map[]
}
```

**Response:** `OrderItems`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-items/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create OrderItems

```
POST /api/v1/order-items/bulk/
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
  orderId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  variantId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
  unitPrice?: number  // e.g. 99.99
  totalPrice?: number  // e.g. 99.99
  discount?: number  // e.g. 99.99
  metadata?: Record<string, unknown>  // e.g. map[]
}
```

**Response:** `OrderItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-items/bulk/"
```

</details>

---

:::tab Find & Update

#### <Badge type="info" text="GET" /> Find OrderItems by ID

```
GET /api/v1/order-items/with-id/:id
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

**Response:** `OrderItems`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/order-items/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update OrderItems

```
PUT /api/v1/order-items/with-id/:id
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
  orderId?: string
  productId?: string
  variantId?: string
  quantity?: number
  unitPrice?: number
  totalPrice?: number
  discount?: number
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
  "http://localhost:3000/api/v1/order-items/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update OrderItems

```
PUT /api/v1/order-items/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: OrderItemsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-items/bulk/"
```

</details>

---

:::tab Delete

#### <Badge type="danger" text="DELETE" /> Delete OrderItems

```
DELETE /api/v1/order-items/with-id/:id
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
  "http://localhost:3000/api/v1/order-items/with-id/:id"
```

</details>

---

::::


::::
