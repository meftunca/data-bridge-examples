---
title: OrderStatusHistory
---

# OrderStatusHistory

**Table:** `orders.order_status_history`

**Base path:** `/order-status-history`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `orders` | `order_id` | `order_status_history_order_id_fkey` | [Orders](./orders) |
| `users` | `changed_by` | `order_status_history_changed_by_fkey` | [Users](./users) |


## Entity Relationship Diagram

```mermaid
erDiagram
    OrderStatusHistory }o--|| Orders : "FK"
    OrderStatusHistory }o--|| Users : "FK"
```

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `order_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `orders` |
| 4 | `from_status` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `to_status` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 6 | `changed_by` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `users` |
| 7 | `note` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `order_id` | `orders` | `order_status_history_order_id_fkey` |
| `changed_by` | `users` | `order_status_history_changed_by_fkey` |


## Go Generated Code

> 📂 Source: [📄 `OrderStatusHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderStatusHistory.go) · [📄 `OrderStatusHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderStatusHistory.go) · [📄 `OrderStatusHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/controllers/OrderStatusHistory.go)

### Structs

::::tabs

:::tab Form

#### OrderStatusHistoryForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderStatusHistory.go#:~:text=type%20OrderStatusHistoryForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `OrderId` | `uuid.UUID` | `orderId` | NO |
| `FromStatus` | `string` | `fromStatus` | NO |
| `ToStatus` | `string` | `toStatus` | NO |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Note` | `string` | `note` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab Model

#### OrderStatusHistory [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderStatusHistory.go#:~:text=type%20OrderStatusHistory%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `OrderId` | `uuid.UUID` | `orderId` | NO |
| `FromStatus` | `string` | `fromStatus` | NO |
| `ToStatus` | `string` | `toStatus` | NO |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Note` | `string` | `note` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab Edit

#### OrderStatusHistoryEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderStatusHistory.go#:~:text=type%20OrderStatusHistoryEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `OrderId` | `*uuid.UUID` | `orderId` | YES |
| `FromStatus` | `*string` | `fromStatus` | YES |
| `ToStatus` | `*string` | `toStatus` | YES |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Note` | `*string` | `note` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Filter

#### OrderStatusHistoryFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderStatusHistory.go#:~:text=type%20OrderStatusHistoryFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `OrderId` | `*uuid.UUID` | `orderId` | YES |
| `FromStatus` | `*string` | `fromStatus` | YES |
| `ToStatus` | `*string` | `toStatus` | YES |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Note` | `*string` | `note` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Page

#### OrderStatusHistoryPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderStatusHistory.go#:~:text=type%20OrderStatusHistoryPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `OrderId` | `uuid.UUID` | `orderId` | NO |
| `FromStatus` | `string` | `fromStatus` | NO |
| `ToStatus` | `string` | `toStatus` | NO |
| `ChangedBy` | `*uuid.UUID` | `changedBy` | YES |
| `Note` | `string` | `note` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab BatchUpdate

#### OrderStatusHistoryBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/OrderStatusHistory.go#:~:text=type%20OrderStatusHistoryBatchUpdate%20struct)

```go
type OrderStatusHistoryBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderStatusHistory.go#:~:text=)%20CreateOrderStatusHistory() | `(OrderStatusHistoryService) CreateOrderStatusHistory(data OrderStatusHistoryForm) (OrderStatusHistoryForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderStatusHistory.go#:~:text=)%20CreateOrderStatusHistoryMultiple() | `(OrderStatusHistoryService) CreateOrderStatusHistoryMultiple(data []OrderStatusHistoryForm) ([]OrderStatusHistoryForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderStatusHistory.go#:~:text=)%20UpdateOrderStatusHistory() | `(OrderStatusHistoryService) UpdateOrderStatusHistory(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderStatusHistory.go#:~:text=)%20UpdateOrderStatusHistoryMultiple() | `(OrderStatusHistoryService) UpdateOrderStatusHistoryMultiple(data []OrderStatusHistoryBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/OrderStatusHistory.go#:~:text=)%20DeleteOrderStatusHistory() | `(OrderStatusHistoryService) DeleteOrderStatusHistory(id uuid.UUID) error` |

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/order-status-history/` | Search with query params |
| `GET` | `/order-status-history/pagination` | Paginated listing |
| `POST` | `/order-status-history/` | Create single record |
| `POST` | `/order-status-history/bulk/` | Create multiple records |
| `PUT` | `/order-status-history/bulk/` | Batch update |
| `GET` | `/order-status-history/with-id/:id` | Get by ID |
| `PUT` | `/order-status-history/with-id/:id` | Update by ID |
| `DELETE` | `/order-status-history/with-id/:id` | Delete by ID |

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
export interface OrderStatusHistory {
  id: string;
  name: string;
  orderId: string;
  fromStatus: string;
  toStatus: string;
  changedBy?: string;
  note: string;
  createdAt: string;
}

export interface OrderStatusHistoryForm {
  name: string;
  orderId: string;
  fromStatus: string;
  toStatus: string;
  changedBy?: string;
  note: string;
  createdAt: string;
}

export interface OrderStatusHistoryEdit {
  id: string;
  name: string;
  orderId: string;
  fromStatus: string;
  toStatus: string;
  changedBy?: string;
  note: string;
  createdAt: string;
}

export interface OrderStatusHistoryPage {
  data: OrderStatusHistory[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type OrderStatusHistoryPathQuery = {
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

const OrderStatusHistoryKeys = {
  all: ["order_status_history"] as const,
  lists: () => [...OrderStatusHistoryKeys.all, "list"] as const,
  detail: (id: any) => [...OrderStatusHistoryKeys.all, "detail", id] as const,
} as const;

export function useOrderStatusHistoryList(query?: OrderStatusHistoryPathQuery) {
  return useQuery({
    queryKey: [...OrderStatusHistoryKeys.lists(), query],
    queryFn: () => fetch(`/order-status-history/pagination`, { method: "GET" }).then(r => r.json()) as Promise<OrderStatusHistoryPage>,
  });
}

export function useOrderStatusHistoryDetail(id: any) {
  return useQuery({
    queryKey: OrderStatusHistoryKeys.detail(id),
    queryFn: () => fetch(`/order-status-history/with-id/:id`).then(r => r.json()) as Promise<OrderStatusHistory>,
  });
}

export function useCreateOrderStatusHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: OrderStatusHistoryForm) =>
      fetch("/order-status-history/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: OrderStatusHistoryKeys.lists() }),
  });
}

export function useUpdateOrderStatusHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: OrderStatusHistoryEdit }) =>
      fetch(`/order-status-history/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: OrderStatusHistoryKeys.all }),
  });
}

export function useDeleteOrderStatusHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/order-status-history/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: OrderStatusHistoryKeys.all }),
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

export const OrderStatusHistoryFormSchema = z.object({
  name: z.string(),
  orderId: z.string().uuid(),
  fromStatus: z.string(),
  toStatus: z.string(),
  changedBy: z.string().uuid().optional(),
  note: z.string(),
  createdAt: z.string().datetime(),
});

export type OrderStatusHistoryFormInput = z.infer<typeof OrderStatusHistoryFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './order_status_history.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search OrderStatusHistory

```
GET /api/v1/order-status-history/
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
| `joins` | `string` | No | Available: Orders, Orders.Customers, Orders.Customers.Users, Orders.Customers.Organizations, Orders.Coupons, Orders.Users, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `orderId` | `string (uuid)` | No | Filter by order_id |
| `fromStatus` | `string` | No | Filter by from_status |
| `toStatus` | `string` | No | Filter by to_status |
| `changedBy` | `string (uuid)` | No | Filter by changed_by |
| `note` | `string` | No | Filter by note |

**Response:** `OrderStatusHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/order-status-history/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search OrderStatusHistory (POST)

```
POST /api/v1/order-status-history/search
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

**Response:** `OrderStatusHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-status-history/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate OrderStatusHistory

```
GET /api/v1/order-status-history/pagination
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
| `joins` | `string` | No | Available: Orders, Orders.Customers, Orders.Customers.Users, Orders.Customers.Organizations, Orders.Coupons, Orders.Users, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `orderId` | `string (uuid)` | No | Filter by order_id |
| `fromStatus` | `string` | No | Filter by from_status |
| `toStatus` | `string` | No | Filter by to_status |
| `changedBy` | `string (uuid)` | No | Filter by changed_by |
| `note` | `string` | No | Filter by note |

**Response:** `PaginationResponse<OrderStatusHistory>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/order-status-history/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate OrderStatusHistory (POST)

```
POST /api/v1/order-status-history/pagination
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

**Response:** `PaginationResponse<OrderStatusHistory>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-status-history/pagination"
```

</details>

---

:::tab Create

#### <Badge type="tip" text="POST" /> Create OrderStatusHistory

```
POST /api/v1/order-status-history/
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
  fromStatus?: string  // e.g. example_from_status
  toStatus?: string  // e.g. example_to_status
  changedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  note?: string  // e.g. example_note
}
```

**Response:** `OrderStatusHistory`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-status-history/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create OrderStatusHistory

```
POST /api/v1/order-status-history/bulk/
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
  fromStatus?: string  // e.g. example_from_status
  toStatus?: string  // e.g. example_to_status
  changedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  note?: string  // e.g. example_note
}
```

**Response:** `OrderStatusHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-status-history/bulk/"
```

</details>

---

:::tab Find & Update

#### <Badge type="info" text="GET" /> Find OrderStatusHistory by ID

```
GET /api/v1/order-status-history/with-id/:id
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

**Response:** `OrderStatusHistory`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/order-status-history/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update OrderStatusHistory

```
PUT /api/v1/order-status-history/with-id/:id
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
  fromStatus?: string
  toStatus?: string
  changedBy?: string
  note?: string
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
  "http://localhost:3000/api/v1/order-status-history/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update OrderStatusHistory

```
PUT /api/v1/order-status-history/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: OrderStatusHistoryEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/order-status-history/bulk/"
```

</details>

---

:::tab Delete

#### <Badge type="danger" text="DELETE" /> Delete OrderStatusHistory

```
DELETE /api/v1/order-status-history/with-id/:id
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
  "http://localhost:3000/api/v1/order-status-history/with-id/:id"
```

</details>

---

::::


::::
