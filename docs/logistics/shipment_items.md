---
title: ShipmentItems
---

# ShipmentItems

**Table:** `logistics.shipment_items`

**Base path:** `/shipment-items`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `shipments` | `shipment_id` | `shipment_items_shipment_id_fkey` | [Shipments](./shipments) |
| `order_items` | `order_item_id` | `shipment_items_order_item_id_fkey` | [OrderItems](./order_items) |


## Entity Relationship Diagram

erDiagram
    ShipmentItems }o--|| Shipments : "FK"
    ShipmentItems }o--|| OrderItems : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `shipment_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `shipments` |
| 4 | `order_item_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `order_items` |
| 5 | `quantity` | `integer` | `int` | `number` | NO | `1` | - | - |
| 6 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `shipment_id` | `shipments` | `shipment_items_shipment_id_fkey` |
| `order_item_id` | `order_items` | `shipment_items_order_item_id_fkey` |


## Go Generated Code

> 📂 Source: [📄 `ShipmentItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentItems.go) · [📄 `ShipmentItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentItems.go) · [📄 `ShipmentItems.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/controllers/ShipmentItems.go)

### Structs

:::tabs

== Form

#### ShipmentItemsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentItems.go#:~:text=type%20ShipmentItemsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `ShipmentId` | `uuid.UUID` | `shipmentId` | NO |
| `OrderItemId` | `uuid.UUID` | `orderItemId` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### ShipmentItems [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentItems.go#:~:text=type%20ShipmentItems%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ShipmentId` | `uuid.UUID` | `shipmentId` | NO |
| `OrderItemId` | `uuid.UUID` | `orderItemId` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### ShipmentItemsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentItems.go#:~:text=type%20ShipmentItemsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ShipmentId` | `*uuid.UUID` | `shipmentId` | YES |
| `OrderItemId` | `*uuid.UUID` | `orderItemId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### ShipmentItemsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentItems.go#:~:text=type%20ShipmentItemsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ShipmentId` | `*uuid.UUID` | `shipmentId` | YES |
| `OrderItemId` | `*uuid.UUID` | `orderItemId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### ShipmentItemsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentItems.go#:~:text=type%20ShipmentItemsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ShipmentId` | `uuid.UUID` | `shipmentId` | NO |
| `OrderItemId` | `uuid.UUID` | `orderItemId` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### ShipmentItemsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentItems.go#:~:text=type%20ShipmentItemsBatchUpdate%20struct)

```go
type ShipmentItemsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentItems.go#:~:text=)%20CreateShipmentItems() | `(ShipmentItemsService) CreateShipmentItems(data ShipmentItemsForm) (ShipmentItemsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentItems.go#:~:text=)%20CreateShipmentItemsMultiple() | `(ShipmentItemsService) CreateShipmentItemsMultiple(data []ShipmentItemsForm) ([]ShipmentItemsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentItems.go#:~:text=)%20UpdateShipmentItems() | `(ShipmentItemsService) UpdateShipmentItems(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentItems.go#:~:text=)%20UpdateShipmentItemsMultiple() | `(ShipmentItemsService) UpdateShipmentItemsMultiple(data []ShipmentItemsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentItems.go#:~:text=)%20DeleteShipmentItems() | `(ShipmentItemsService) DeleteShipmentItems(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/shipment-items/` | Search with query params |
| `GET` | `/shipment-items/pagination` | Paginated listing |
| `POST` | `/shipment-items/` | Create single record |
| `POST` | `/shipment-items/bulk/` | Create multiple records |
| `PUT` | `/shipment-items/bulk/` | Batch update |
| `GET` | `/shipment-items/with-id/:id` | Get by ID |
| `PUT` | `/shipment-items/with-id/:id` | Update by ID |
| `DELETE` | `/shipment-items/with-id/:id` | Delete by ID |

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


:::tab Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export interface ShipmentItems {
  id: string;
  name: string;
  shipmentId: string;
  orderItemId: string;
  quantity: number;
  createdAt: string;
}

export interface ShipmentItemsForm {
  name: string;
  shipmentId: string;
  orderItemId: string;
  quantity: number;
  createdAt: string;
}

export interface ShipmentItemsEdit {
  id: string;
  name: string;
  shipmentId: string;
  orderItemId: string;
  quantity: number;
  createdAt: string;
}

export interface ShipmentItemsPage {
  data: ShipmentItems[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type ShipmentItemsPathQuery = {
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

const ShipmentItemsKeys = {
  all: ["shipment_items"] as const,
  lists: () => [...ShipmentItemsKeys.all, "list"] as const,
  detail: (id: any) => [...ShipmentItemsKeys.all, "detail", id] as const,
} as const;

export function useShipmentItemsList(query?: ShipmentItemsPathQuery) {
  return useQuery({
    queryKey: [...ShipmentItemsKeys.lists(), query],
    queryFn: () => fetch(`/shipment-items/pagination`, { method: "GET" }).then(r => r.json()) as Promise<ShipmentItemsPage>,
  });
}

export function useShipmentItemsDetail(id: any) {
  return useQuery({
    queryKey: ShipmentItemsKeys.detail(id),
    queryFn: () => fetch(`/shipment-items/with-id/:id`).then(r => r.json()) as Promise<ShipmentItems>,
  });
}

export function useCreateShipmentItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: ShipmentItemsForm) =>
      fetch("/shipment-items/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ShipmentItemsKeys.lists() }),
  });
}

export function useUpdateShipmentItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: ShipmentItemsEdit }) =>
      fetch(`/shipment-items/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ShipmentItemsKeys.all }),
  });
}

export function useDeleteShipmentItems() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/shipment-items/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ShipmentItemsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const ShipmentItemsFormSchema = z.object({
  name: z.string(),
  shipmentId: z.string().uuid(),
  orderItemId: z.string().uuid(),
  quantity: z.number().int(),
  createdAt: z.string().datetime(),
});

export type ShipmentItemsFormInput = z.infer<typeof ShipmentItemsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './shipment_items.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search ShipmentItems

```
GET /api/v1/shipment-items/
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
| `preloads` | `string` | No | Available: ShipmentIdDetail, ShipmentIdDetail.ShipmentItemsList, ShipmentIdDetail.ShipmentItemsList.ShipmentIdDetail, ShipmentIdDetail.ShipmentTrackingList, ShipmentIdDetail.ShipmentTrackingList.ShipmentIdDetail, ShipmentIdDetail.WarehouseIdDetail, ShipmentIdDetail.WarehouseIdDetail.StorageZonesList, ShipmentIdDetail.WarehouseIdDetail.InventoryList, ShipmentIdDetail.WarehouseIdDetail.PurchaseOrdersList, ShipmentIdDetail.WarehouseIdDetail.ShipmentsList |
| `joins` | `string` | No | Available: Shipments, Shipments.Orders, Shipments.Warehouses, Shipments.Warehouses.Organizations, Shipments.Warehouses.Users, OrderItems |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `shipmentId` | `string (uuid)` | No | Filter by shipment_id |
| `orderItemId` | `string (uuid)` | No | Filter by order_item_id |
| `quantity` | `integer` | No | Filter by quantity |

**Response:** `ShipmentItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/shipment-items/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search ShipmentItems (POST)

```
POST /api/v1/shipment-items/search
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

**Response:** `ShipmentItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-items/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate ShipmentItems

```
GET /api/v1/shipment-items/pagination
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
| `preloads` | `string` | No | Available: ShipmentIdDetail, ShipmentIdDetail.ShipmentItemsList, ShipmentIdDetail.ShipmentItemsList.ShipmentIdDetail, ShipmentIdDetail.ShipmentTrackingList, ShipmentIdDetail.ShipmentTrackingList.ShipmentIdDetail, ShipmentIdDetail.WarehouseIdDetail, ShipmentIdDetail.WarehouseIdDetail.StorageZonesList, ShipmentIdDetail.WarehouseIdDetail.InventoryList, ShipmentIdDetail.WarehouseIdDetail.PurchaseOrdersList, ShipmentIdDetail.WarehouseIdDetail.ShipmentsList |
| `joins` | `string` | No | Available: Shipments, Shipments.Orders, Shipments.Warehouses, Shipments.Warehouses.Organizations, Shipments.Warehouses.Users, OrderItems |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `shipmentId` | `string (uuid)` | No | Filter by shipment_id |
| `orderItemId` | `string (uuid)` | No | Filter by order_item_id |
| `quantity` | `integer` | No | Filter by quantity |

**Response:** `PaginationResponse<ShipmentItems>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/shipment-items/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate ShipmentItems (POST)

```
POST /api/v1/shipment-items/pagination
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

**Response:** `PaginationResponse<ShipmentItems>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-items/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create ShipmentItems

```
POST /api/v1/shipment-items/
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
  shipmentId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  orderItemId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
}
```

**Response:** `ShipmentItems`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-items/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create ShipmentItems

```
POST /api/v1/shipment-items/bulk/
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
  shipmentId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  orderItemId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
}
```

**Response:** `ShipmentItems[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-items/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find ShipmentItems by ID

```
GET /api/v1/shipment-items/with-id/:id
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

**Response:** `ShipmentItems`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/shipment-items/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update ShipmentItems

```
PUT /api/v1/shipment-items/with-id/:id
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
  shipmentId?: string
  orderItemId?: string
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
  "http://localhost:3000/api/v1/shipment-items/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update ShipmentItems

```
PUT /api/v1/shipment-items/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: ShipmentItemsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-items/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete ShipmentItems

```
DELETE /api/v1/shipment-items/with-id/:id
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
  "http://localhost:3000/api/v1/shipment-items/with-id/:id"
```

</details>

---

:::


::::
