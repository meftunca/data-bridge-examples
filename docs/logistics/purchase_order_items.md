---
title: PurchaseOrderItems
---

# PurchaseOrderItems

**Table:** `logistics.purchase_order_items`

**Base path:** `/purchase-order-items`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `purchase_orders` | `purchase_order_id` | `purchase_order_items_purchase_order_id_fkey` | [PurchaseOrders](./purchase_orders) |
| `products` | `product_id` | `purchase_order_items_product_id_fkey` | [Products](./products) |


## Entity Relationship Diagram

```mermaid
erDiagram
    PurchaseOrderItems }o--|| PurchaseOrders : "FK"
    PurchaseOrderItems }o--|| Products : "FK"
```

::::tabs

=== FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `purchase_order_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `purchase_orders` |
| 4 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `products` |
| 5 | `quantity` | `integer` | `int` | `number` | NO | `1` | - | - |
| 6 | `unit_cost` | `numeric` | `float64` | `number` | NO | `0.00` | - | - |
| 7 | `received_qty` | `integer` | `int` | `number` | NO | `0` | - | - |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 9 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `purchase_order_id` | `purchase_orders` | `purchase_order_items_purchase_order_id_fkey` |
| `product_id` | `products` | `purchase_order_items_product_id_fkey` |


## Go Generated Code

> 📂 Source: [📄 `PurchaseOrderItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/PurchaseOrderItems.go) · [📄 `PurchaseOrderItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/PurchaseOrderItems.go) · [📄 `PurchaseOrderItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/controllers/PurchaseOrderItems.go)

### Structs

:::tabs

== Form

#### PurchaseOrderItemsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/PurchaseOrderItems.go#:~:text=type%20PurchaseOrderItemsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `PurchaseOrderId` | `uuid.UUID` | `purchaseOrderId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `UnitCost` | `float64` | `unitCost` | NO |
| `ReceivedQty` | `int` | `receivedQty` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### PurchaseOrderItems [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/PurchaseOrderItems.go#:~:text=type%20PurchaseOrderItems%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `PurchaseOrderId` | `uuid.UUID` | `purchaseOrderId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `UnitCost` | `float64` | `unitCost` | NO |
| `ReceivedQty` | `int` | `receivedQty` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### PurchaseOrderItemsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/PurchaseOrderItems.go#:~:text=type%20PurchaseOrderItemsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `PurchaseOrderId` | `*uuid.UUID` | `purchaseOrderId` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `UnitCost` | `*float64` | `unitCost` | YES |
| `ReceivedQty` | `*int` | `receivedQty` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### PurchaseOrderItemsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/PurchaseOrderItems.go#:~:text=type%20PurchaseOrderItemsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `PurchaseOrderId` | `*uuid.UUID` | `purchaseOrderId` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `UnitCost` | `*float64` | `unitCost` | YES |
| `ReceivedQty` | `*int` | `receivedQty` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### PurchaseOrderItemsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/PurchaseOrderItems.go#:~:text=type%20PurchaseOrderItemsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `PurchaseOrderId` | `uuid.UUID` | `purchaseOrderId` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `UnitCost` | `float64` | `unitCost` | NO |
| `ReceivedQty` | `int` | `receivedQty` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### PurchaseOrderItemsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/PurchaseOrderItems.go#:~:text=type%20PurchaseOrderItemsBatchUpdate%20struct)

```go
type PurchaseOrderItemsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/PurchaseOrderItems.go#:~:text=%29%20CreatePurchaseOrderItems%28%29) | `(PurchaseOrderItemsService) CreatePurchaseOrderItems(data PurchaseOrderItemsForm) (PurchaseOrderItemsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/PurchaseOrderItems.go#:~:text=%29%20CreatePurchaseOrderItemsMultiple%28%29) | `(PurchaseOrderItemsService) CreatePurchaseOrderItemsMultiple(data []PurchaseOrderItemsForm) ([]PurchaseOrderItemsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/PurchaseOrderItems.go#:~:text=%29%20UpdatePurchaseOrderItems%28%29) | `(PurchaseOrderItemsService) UpdatePurchaseOrderItems(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/PurchaseOrderItems.go#:~:text=%29%20UpdatePurchaseOrderItemsMultiple%28%29) | `(PurchaseOrderItemsService) UpdatePurchaseOrderItemsMultiple(data []PurchaseOrderItemsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/PurchaseOrderItems.go#:~:text=%29%20DeletePurchaseOrderItems%28%29) | `(PurchaseOrderItemsService) DeletePurchaseOrderItems(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/purchase-order-items/` | Search with query params |
| `GET` | `/purchase-order-items/pagination` | Paginated listing |
| `POST` | `/purchase-order-items/` | Create single record |
| `POST` | `/purchase-order-items/bulk/` | Create multiple records |
| `PUT` | `/purchase-order-items/bulk/` | Batch update |
| `GET` | `/purchase-order-items/with-id/:id` | Get by ID |
| `PUT` | `/purchase-order-items/with-id/:id` | Update by ID |
| `DELETE` | `/purchase-order-items/with-id/:id` | Delete by ID |

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
| `low_stock_count` | `p_warehouse_id uuid` | `integer` | `/rpc/low_stock_count` |
| `warehouse_utilization` | `p_warehouse_id uuid` | `numeric` | `/rpc/warehouse_utilization` |


=== Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export interface PurchaseOrderItems {
  id: string;
  name: string;
  purchaseOrderId: string;
  productId: string;
  quantity: number;
  unitCost: number;
  receivedQty: number;
  createdAt: string;
  updatedAt: string;
}

export interface PurchaseOrderItemsForm {
  name: string;
  purchaseOrderId: string;
  productId: string;
  quantity: number;
  unitCost: number;
  receivedQty: number;
  createdAt: string;
  updatedAt: string;
}

export interface PurchaseOrderItemsEdit {
  id: string;
  name: string;
  purchaseOrderId: string;
  productId: string;
  quantity: number;
  unitCost: number;
  receivedQty: number;
  createdAt: string;
  updatedAt: string;
}

export interface PurchaseOrderItemsPage {
  data: PurchaseOrderItems[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type PurchaseOrderItemsPathQuery = {
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

const PurchaseOrderItemsKeys = {
  all: ["purchase_order_items"] as const,
  lists: () => [...PurchaseOrderItemsKeys.all, "list"] as const,
  detail: (id: any) => [...PurchaseOrderItemsKeys.all, "detail", id] as const,
} as const;

export function usePurchaseOrderItemsList(query?: PurchaseOrderItemsPathQuery) {
  return useQuery({
    queryKey: [...PurchaseOrderItemsKeys.lists(), query],
    queryFn: () => fetch(`/purchase-order-items/pagination`, { method: "GET" }).then(r => r.json()) as Promise<PurchaseOrderItemsPage>,
  });
}

export function usePurchaseOrderItemsDetail(id: any) {
  return useQuery({
    queryKey: PurchaseOrderItemsKeys.detail(id),
    queryFn: () => fetch(`/purchase-order-items/with-id/:id`).then(r => r.json()) as Promise<PurchaseOrderItems>,
  });
}

export function useCreatePurchaseOrderItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: PurchaseOrderItemsForm) =>
      fetch("/purchase-order-items/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: PurchaseOrderItemsKeys.lists() }),
  });
}

export function useUpdatePurchaseOrderItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: PurchaseOrderItemsEdit }) =>
      fetch(`/purchase-order-items/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: PurchaseOrderItemsKeys.all }),
  });
}

export function useDeletePurchaseOrderItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/purchase-order-items/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: PurchaseOrderItemsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const PurchaseOrderItemsFormSchema = z.object({
  name: z.string(),
  purchaseOrderId: z.string().uuid(),
  productId: z.string().uuid(),
  quantity: z.number().int(),
  unitCost: z.number(),
  receivedQty: z.number().int(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type PurchaseOrderItemsFormInput = z.infer<typeof PurchaseOrderItemsFormSchema>;

```

:::


=== API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './purchase_order_items.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search PurchaseOrderItems

```
GET /api/v1/purchase-order-items/
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
| `preloads` | `string` | No | Available: PurchaseOrderIdDetail, PurchaseOrderIdDetail.PurchaseOrderItemsList, PurchaseOrderIdDetail.PurchaseOrderItemsList.PurchaseOrderIdDetail, PurchaseOrderIdDetail.SupplierIdDetail, PurchaseOrderIdDetail.SupplierIdDetail.PurchaseOrdersList, PurchaseOrderIdDetail.WarehouseIdDetail, PurchaseOrderIdDetail.WarehouseIdDetail.StorageZonesList, PurchaseOrderIdDetail.WarehouseIdDetail.InventoryList, PurchaseOrderIdDetail.WarehouseIdDetail.PurchaseOrdersList, PurchaseOrderIdDetail.WarehouseIdDetail.ShipmentsList |
| `joins` | `string` | No | Available: PurchaseOrders, PurchaseOrders.Suppliers, PurchaseOrders.Warehouses, PurchaseOrders.Warehouses.Organizations, PurchaseOrders.Warehouses.Users, PurchaseOrders.Users, Products |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `purchaseOrderId` | `string (uuid)` | No | Filter by purchase_order_id |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `quantity` | `integer` | No | Filter by quantity |
| `unitCost` | `number` | No | Filter by unit_cost |
| `receivedQty` | `integer` | No | Filter by received_qty |

**Response:** `PurchaseOrderItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/purchase-order-items/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search PurchaseOrderItems (POST)

```
POST /api/v1/purchase-order-items/search
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

**Response:** `PurchaseOrderItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/purchase-order-items/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate PurchaseOrderItems

```
GET /api/v1/purchase-order-items/pagination
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
| `preloads` | `string` | No | Available: PurchaseOrderIdDetail, PurchaseOrderIdDetail.PurchaseOrderItemsList, PurchaseOrderIdDetail.PurchaseOrderItemsList.PurchaseOrderIdDetail, PurchaseOrderIdDetail.SupplierIdDetail, PurchaseOrderIdDetail.SupplierIdDetail.PurchaseOrdersList, PurchaseOrderIdDetail.WarehouseIdDetail, PurchaseOrderIdDetail.WarehouseIdDetail.StorageZonesList, PurchaseOrderIdDetail.WarehouseIdDetail.InventoryList, PurchaseOrderIdDetail.WarehouseIdDetail.PurchaseOrdersList, PurchaseOrderIdDetail.WarehouseIdDetail.ShipmentsList |
| `joins` | `string` | No | Available: PurchaseOrders, PurchaseOrders.Suppliers, PurchaseOrders.Warehouses, PurchaseOrders.Warehouses.Organizations, PurchaseOrders.Warehouses.Users, PurchaseOrders.Users, Products |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `purchaseOrderId` | `string (uuid)` | No | Filter by purchase_order_id |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `quantity` | `integer` | No | Filter by quantity |
| `unitCost` | `number` | No | Filter by unit_cost |
| `receivedQty` | `integer` | No | Filter by received_qty |

**Response:** `PaginationResponse<PurchaseOrderItems>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/purchase-order-items/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate PurchaseOrderItems (POST)

```
POST /api/v1/purchase-order-items/pagination
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

**Response:** `PaginationResponse<PurchaseOrderItems>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/purchase-order-items/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create PurchaseOrderItems

```
POST /api/v1/purchase-order-items/
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
  purchaseOrderId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
  unitCost?: number  // e.g. 99.99
  receivedQty?: number  // e.g. 1
}
```

**Response:** `PurchaseOrderItems`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/purchase-order-items/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create PurchaseOrderItems

```
POST /api/v1/purchase-order-items/bulk/
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
  purchaseOrderId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
  unitCost?: number  // e.g. 99.99
  receivedQty?: number  // e.g. 1
}
```

**Response:** `PurchaseOrderItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/purchase-order-items/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find PurchaseOrderItems by ID

```
GET /api/v1/purchase-order-items/with-id/:id
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

**Response:** `PurchaseOrderItems`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/purchase-order-items/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update PurchaseOrderItems

```
PUT /api/v1/purchase-order-items/with-id/:id
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
  purchaseOrderId?: string
  productId?: string
  quantity?: number
  unitCost?: number
  receivedQty?: number
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
  "http://localhost:3000/api/v1/purchase-order-items/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update PurchaseOrderItems

```
PUT /api/v1/purchase-order-items/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: PurchaseOrderItemsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/purchase-order-items/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete PurchaseOrderItems

```
DELETE /api/v1/purchase-order-items/with-id/:id
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
  "http://localhost:3000/api/v1/purchase-order-items/with-id/:id"
```

</details>

---

:::


::::
