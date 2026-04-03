---
title: ShipmentTracking
---

# ShipmentTracking

**Table:** `logistics.shipment_tracking`

**Base path:** `/shipment-tracking`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `shipments` | `shipment_id` | `shipment_tracking_shipment_id_fkey` | [Shipments](./shipments) |


## Entity Relationship Diagram

erDiagram
    ShipmentTracking }o--|| Shipments : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `shipment_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `shipments` |
| 4 | `status` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `location` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 6 | `description` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `event_time` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | - |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `shipment_id` | `shipments` | `shipment_tracking_shipment_id_fkey` |


## Go Generated Code

> 📂 Source: [📄 `ShipmentTracking.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentTracking.go) · [📄 `ShipmentTracking.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentTracking.go) · [📄 `ShipmentTracking.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/controllers/ShipmentTracking.go)

### Structs

:::tabs

== Form

#### ShipmentTrackingForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentTracking.go#:~:text=type%20ShipmentTrackingForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `ShipmentId` | `uuid.UUID` | `shipmentId` | NO |
| `Status` | `string` | `status` | NO |
| `Location` | `string` | `location` | NO |
| `Description` | `string` | `description` | NO |
| `EventTime` | `time.Time` | `eventTime` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### ShipmentTracking [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentTracking.go#:~:text=type%20ShipmentTracking%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ShipmentId` | `uuid.UUID` | `shipmentId` | NO |
| `Status` | `string` | `status` | NO |
| `Location` | `string` | `location` | NO |
| `Description` | `string` | `description` | NO |
| `EventTime` | `time.Time` | `eventTime` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### ShipmentTrackingEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentTracking.go#:~:text=type%20ShipmentTrackingEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ShipmentId` | `*uuid.UUID` | `shipmentId` | YES |
| `Status` | `*string` | `status` | YES |
| `Location` | `*string` | `location` | YES |
| `Description` | `*string` | `description` | YES |
| `EventTime` | `*time.Time` | `eventTime` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### ShipmentTrackingFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentTracking.go#:~:text=type%20ShipmentTrackingFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ShipmentId` | `*uuid.UUID` | `shipmentId` | YES |
| `Status` | `*string` | `status` | YES |
| `Location` | `*string` | `location` | YES |
| `Description` | `*string` | `description` | YES |
| `EventTime` | `*time.Time` | `eventTime` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### ShipmentTrackingPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentTracking.go#:~:text=type%20ShipmentTrackingPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ShipmentId` | `uuid.UUID` | `shipmentId` | NO |
| `Status` | `string` | `status` | NO |
| `Location` | `string` | `location` | NO |
| `Description` | `string` | `description` | NO |
| `EventTime` | `time.Time` | `eventTime` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### ShipmentTrackingBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/ShipmentTracking.go#:~:text=type%20ShipmentTrackingBatchUpdate%20struct)

```go
type ShipmentTrackingBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentTracking.go#:~:text=)%20CreateShipmentTracking() | `(ShipmentTrackingService) CreateShipmentTracking(data ShipmentTrackingForm) (ShipmentTrackingForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentTracking.go#:~:text=)%20CreateShipmentTrackingMultiple() | `(ShipmentTrackingService) CreateShipmentTrackingMultiple(data []ShipmentTrackingForm) ([]ShipmentTrackingForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentTracking.go#:~:text=)%20UpdateShipmentTracking() | `(ShipmentTrackingService) UpdateShipmentTracking(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentTracking.go#:~:text=)%20UpdateShipmentTrackingMultiple() | `(ShipmentTrackingService) UpdateShipmentTrackingMultiple(data []ShipmentTrackingBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/ShipmentTracking.go#:~:text=)%20DeleteShipmentTracking() | `(ShipmentTrackingService) DeleteShipmentTracking(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/shipment-tracking/` | Search with query params |
| `GET` | `/shipment-tracking/pagination` | Paginated listing |
| `POST` | `/shipment-tracking/` | Create single record |
| `POST` | `/shipment-tracking/bulk/` | Create multiple records |
| `PUT` | `/shipment-tracking/bulk/` | Batch update |
| `GET` | `/shipment-tracking/with-id/:id` | Get by ID |
| `PUT` | `/shipment-tracking/with-id/:id` | Update by ID |
| `DELETE` | `/shipment-tracking/with-id/:id` | Delete by ID |

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
export interface ShipmentTracking {
  id: string;
  name: string;
  shipmentId: string;
  status: string;
  location: string;
  description: string;
  eventTime: string;
  createdAt: string;
}

export interface ShipmentTrackingForm {
  name: string;
  shipmentId: string;
  status: string;
  location: string;
  description: string;
  eventTime: string;
  createdAt: string;
}

export interface ShipmentTrackingEdit {
  id: string;
  name: string;
  shipmentId: string;
  status: string;
  location: string;
  description: string;
  eventTime: string;
  createdAt: string;
}

export interface ShipmentTrackingPage {
  data: ShipmentTracking[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type ShipmentTrackingPathQuery = {
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

const ShipmentTrackingKeys = {
  all: ["shipment_tracking"] as const,
  lists: () => [...ShipmentTrackingKeys.all, "list"] as const,
  detail: (id: any) => [...ShipmentTrackingKeys.all, "detail", id] as const,
} as const;

export function useShipmentTrackingList(query?: ShipmentTrackingPathQuery) {
  return useQuery({
    queryKey: [...ShipmentTrackingKeys.lists(), query],
    queryFn: () => fetch(`/shipment-tracking/pagination`, { method: "GET" }).then(r => r.json()) as Promise<ShipmentTrackingPage>,
  });
}

export function useShipmentTrackingDetail(id: any) {
  return useQuery({
    queryKey: ShipmentTrackingKeys.detail(id),
    queryFn: () => fetch(`/shipment-tracking/with-id/:id`).then(r => r.json()) as Promise<ShipmentTracking>,
  });
}

export function useCreateShipmentTracking() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: ShipmentTrackingForm) =>
      fetch("/shipment-tracking/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ShipmentTrackingKeys.lists() }),
  });
}

export function useUpdateShipmentTracking() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: ShipmentTrackingEdit }) =>
      fetch(`/shipment-tracking/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ShipmentTrackingKeys.all }),
  });
}

export function useDeleteShipmentTracking() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/shipment-tracking/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ShipmentTrackingKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const ShipmentTrackingFormSchema = z.object({
  name: z.string(),
  shipmentId: z.string().uuid(),
  status: z.string(),
  location: z.string(),
  description: z.string(),
  eventTime: z.string().datetime(),
  createdAt: z.string().datetime(),
});

export type ShipmentTrackingFormInput = z.infer<typeof ShipmentTrackingFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './shipment_tracking.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search ShipmentTracking

```
GET /api/v1/shipment-tracking/
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
| `joins` | `string` | No | Available: Shipments, Shipments.Orders, Shipments.Warehouses, Shipments.Warehouses.Organizations, Shipments.Warehouses.Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `shipmentId` | `string (uuid)` | No | Filter by shipment_id |
| `status` | `string` | No | Filter by status |
| `location` | `string` | No | Filter by location |
| `description` | `string` | No | Filter by description |
| `eventTime` | `string (date-time)` | No | Filter by event_time |

**Response:** `ShipmentTracking[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/shipment-tracking/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search ShipmentTracking (POST)

```
POST /api/v1/shipment-tracking/search
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

**Response:** `ShipmentTracking[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-tracking/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate ShipmentTracking

```
GET /api/v1/shipment-tracking/pagination
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
| `joins` | `string` | No | Available: Shipments, Shipments.Orders, Shipments.Warehouses, Shipments.Warehouses.Organizations, Shipments.Warehouses.Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `shipmentId` | `string (uuid)` | No | Filter by shipment_id |
| `status` | `string` | No | Filter by status |
| `location` | `string` | No | Filter by location |
| `description` | `string` | No | Filter by description |
| `eventTime` | `string (date-time)` | No | Filter by event_time |

**Response:** `PaginationResponse<ShipmentTracking>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/shipment-tracking/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate ShipmentTracking (POST)

```
POST /api/v1/shipment-tracking/pagination
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

**Response:** `PaginationResponse<ShipmentTracking>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-tracking/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create ShipmentTracking

```
POST /api/v1/shipment-tracking/
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
  status?: string  // e.g. example_status
  location?: string  // e.g. example_location
  description?: string  // e.g. example_description
  eventTime?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `ShipmentTracking`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-tracking/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create ShipmentTracking

```
POST /api/v1/shipment-tracking/bulk/
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
  status?: string  // e.g. example_status
  location?: string  // e.g. example_location
  description?: string  // e.g. example_description
  eventTime?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `ShipmentTracking[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-tracking/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find ShipmentTracking by ID

```
GET /api/v1/shipment-tracking/with-id/:id
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

**Response:** `ShipmentTracking`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/shipment-tracking/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update ShipmentTracking

```
PUT /api/v1/shipment-tracking/with-id/:id
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
  status?: string
  location?: string
  description?: string
  eventTime?: string
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
  "http://localhost:3000/api/v1/shipment-tracking/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update ShipmentTracking

```
PUT /api/v1/shipment-tracking/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: ShipmentTrackingEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/shipment-tracking/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete ShipmentTracking

```
DELETE /api/v1/shipment-tracking/with-id/:id
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
  "http://localhost:3000/api/v1/shipment-tracking/with-id/:id"
```

</details>

---

:::


::::
