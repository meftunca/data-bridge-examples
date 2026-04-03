---
title: StorageZones
---

# StorageZones

**Table:** `logistics.storage_zones`

**Base path:** `/storage-zones`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `warehouses` | `warehouse_id` | `storage_zones_warehouse_id_fkey` | [Warehouses](./warehouses) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `storage_bins` | `zone_id` | `storage_bins_zone_id_fkey` | [StorageBins](./storage_bins) |


## Entity Relationship Diagram

erDiagram
    StorageZones }o--|| Warehouses : "FK"
    StorageZones ||--o{ StorageBins : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `warehouse_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `UQ` `FK` | → References `warehouses` |
| 4 | `zone_code` | `text` | `string` | `string` | NO | `'A'::text` | `UQ` | - |
| 5 | `zone_type` | `text` | `string` | `string` | NO | `'general'::text` | - | - |
| 6 | `temperature_min` | `numeric` | `float64` | `number` | YES | - | - | - |
| 7 | `temperature_max` | `numeric` | `float64` | `number` | YES | - | - | - |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 9 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `warehouse_id` | `warehouses` | `storage_zones_warehouse_id_fkey` |

## Unique Keys

- `warehouse_id` (`uuid`)
- `zone_code` (`text`)


## Go Generated Code

> 📂 Source: [📄 `StorageZones.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StorageZones.go) · [📄 `StorageZones.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StorageZones.go) · [📄 `StorageZones.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/controllers/StorageZones.go)

### Structs

:::tabs

== Form

#### StorageZonesForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StorageZones.go#:~:text=type%20StorageZonesForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `WarehouseId` | `uuid.UUID` | `warehouseId` | NO |
| `ZoneCode` | `string` | `zoneCode` | NO |
| `ZoneType` | `string` | `zoneType` | NO |
| `TemperatureMin` | `*float64` | `temperatureMin` | YES |
| `TemperatureMax` | `*float64` | `temperatureMax` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### StorageZones [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StorageZones.go#:~:text=type%20StorageZones%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `WarehouseId` | `uuid.UUID` | `warehouseId` | NO |
| `ZoneCode` | `string` | `zoneCode` | NO |
| `ZoneType` | `string` | `zoneType` | NO |
| `TemperatureMin` | `*float64` | `temperatureMin` | YES |
| `TemperatureMax` | `*float64` | `temperatureMax` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### StorageZonesEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StorageZones.go#:~:text=type%20StorageZonesEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `WarehouseId` | `*uuid.UUID` | `warehouseId` | YES |
| `ZoneCode` | `*string` | `zoneCode` | YES |
| `ZoneType` | `*string` | `zoneType` | YES |
| `TemperatureMin` | `*float64` | `temperatureMin` | YES |
| `TemperatureMax` | `*float64` | `temperatureMax` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### StorageZonesFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StorageZones.go#:~:text=type%20StorageZonesFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `WarehouseId` | `*uuid.UUID` | `warehouseId` | YES |
| `ZoneCode` | `*string` | `zoneCode` | YES |
| `ZoneType` | `*string` | `zoneType` | YES |
| `TemperatureMin` | `*float64` | `temperatureMin` | YES |
| `TemperatureMax` | `*float64` | `temperatureMax` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### StorageZonesPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StorageZones.go#:~:text=type%20StorageZonesPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `WarehouseId` | `uuid.UUID` | `warehouseId` | NO |
| `ZoneCode` | `string` | `zoneCode` | NO |
| `ZoneType` | `string` | `zoneType` | NO |
| `TemperatureMin` | `*float64` | `temperatureMin` | YES |
| `TemperatureMax` | `*float64` | `temperatureMax` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### StorageZonesBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StorageZones.go#:~:text=type%20StorageZonesBatchUpdate%20struct)

```go
type StorageZonesBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StorageZones.go#:~:text=)%20CreateStorageZones() | `(StorageZonesService) CreateStorageZones(data StorageZonesForm) (StorageZonesForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StorageZones.go#:~:text=)%20CreateStorageZonesMultiple() | `(StorageZonesService) CreateStorageZonesMultiple(data []StorageZonesForm) ([]StorageZonesForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StorageZones.go#:~:text=)%20UpdateStorageZones() | `(StorageZonesService) UpdateStorageZones(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StorageZones.go#:~:text=)%20UpdateStorageZonesMultiple() | `(StorageZonesService) UpdateStorageZonesMultiple(data []StorageZonesBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StorageZones.go#:~:text=)%20DeleteStorageZones() | `(StorageZonesService) DeleteStorageZones(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/storage-zones/` | Search with query params |
| `GET` | `/storage-zones/pagination` | Paginated listing |
| `POST` | `/storage-zones/` | Create single record |
| `POST` | `/storage-zones/bulk/` | Create multiple records |
| `PUT` | `/storage-zones/bulk/` | Batch update |
| `GET` | `/storage-zones/with-id/:id` | Get by ID |
| `PUT` | `/storage-zones/with-id/:id` | Update by ID |
| `DELETE` | `/storage-zones/with-id/:id` | Delete by ID |

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
export interface StorageZones {
  id: string;
  name: string;
  warehouseId: string;
  zoneCode: string;
  zoneType: string;
  temperatureMin?: number;
  temperatureMax?: number;
  createdAt: string;
  updatedAt: string;
}

export interface StorageZonesForm {
  name: string;
  warehouseId: string;
  zoneCode: string;
  zoneType: string;
  temperatureMin?: number;
  temperatureMax?: number;
  createdAt: string;
  updatedAt: string;
}

export interface StorageZonesEdit {
  id: string;
  name: string;
  warehouseId: string;
  zoneCode: string;
  zoneType: string;
  temperatureMin?: number;
  temperatureMax?: number;
  createdAt: string;
  updatedAt: string;
}

export interface StorageZonesPage {
  data: StorageZones[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type StorageZonesPathQuery = {
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

const StorageZonesKeys = {
  all: ["storage_zones"] as const,
  lists: () => [...StorageZonesKeys.all, "list"] as const,
  detail: (id: any) => [...StorageZonesKeys.all, "detail", id] as const,
} as const;

export function useStorageZonesList(query?: StorageZonesPathQuery) {
  return useQuery({
    queryKey: [...StorageZonesKeys.lists(), query],
    queryFn: () => fetch(`/storage-zones/pagination`, { method: "GET" }).then(r => r.json()) as Promise<StorageZonesPage>,
  });
}

export function useStorageZonesDetail(id: any) {
  return useQuery({
    queryKey: StorageZonesKeys.detail(id),
    queryFn: () => fetch(`/storage-zones/with-id/:id`).then(r => r.json()) as Promise<StorageZones>,
  });
}

export function useCreateStorageZones() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: StorageZonesForm) =>
      fetch("/storage-zones/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: StorageZonesKeys.lists() }),
  });
}

export function useUpdateStorageZones() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: StorageZonesEdit }) =>
      fetch(`/storage-zones/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: StorageZonesKeys.all }),
  });
}

export function useDeleteStorageZones() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/storage-zones/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: StorageZonesKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const StorageZonesFormSchema = z.object({
  name: z.string(),
  warehouseId: z.string().uuid(),
  zoneCode: z.string(),
  zoneType: z.string(),
  temperatureMin: z.number().optional(),
  temperatureMax: z.number().optional(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type StorageZonesFormInput = z.infer<typeof StorageZonesFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './storage_zones.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search StorageZones

```
GET /api/v1/storage-zones/
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
| `preloads` | `string` | No | Available: StorageBinsList, StorageBinsList.InventoryList, StorageBinsList.InventoryList.StockMovementsList, StorageBinsList.InventoryList.WarehouseIdDetail, StorageBinsList.InventoryList.BinIdDetail, StorageBinsList.ZoneIdDetail, StorageBinsList.ZoneIdDetail.StorageBinsList, StorageBinsList.ZoneIdDetail.WarehouseIdDetail, WarehouseIdDetail, WarehouseIdDetail.StorageZonesList, WarehouseIdDetail.StorageZonesList.StorageBinsList, WarehouseIdDetail.StorageZonesList.WarehouseIdDetail, WarehouseIdDetail.InventoryList, WarehouseIdDetail.InventoryList.StockMovementsList, WarehouseIdDetail.InventoryList.WarehouseIdDetail, WarehouseIdDetail.InventoryList.BinIdDetail, WarehouseIdDetail.PurchaseOrdersList, WarehouseIdDetail.PurchaseOrdersList.PurchaseOrderItemsList, WarehouseIdDetail.PurchaseOrdersList.SupplierIdDetail, WarehouseIdDetail.PurchaseOrdersList.WarehouseIdDetail, WarehouseIdDetail.ShipmentsList, WarehouseIdDetail.ShipmentsList.ShipmentItemsList, WarehouseIdDetail.ShipmentsList.ShipmentTrackingList, WarehouseIdDetail.ShipmentsList.WarehouseIdDetail |
| `joins` | `string` | No | Available: Warehouses, Warehouses.Organizations, Warehouses.Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `warehouseId` | `string (uuid)` | No | Filter by warehouse_id |
| `zoneCode` | `string` | No | Filter by zone_code |
| `zoneType` | `string` | No | Filter by zone_type |
| `temperatureMin` | `number` | No | Filter by temperature_min |
| `temperatureMax` | `number` | No | Filter by temperature_max |

**Response:** `StorageZones[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/storage-zones/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search StorageZones (POST)

```
POST /api/v1/storage-zones/search
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

**Response:** `StorageZones[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/storage-zones/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate StorageZones

```
GET /api/v1/storage-zones/pagination
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
| `preloads` | `string` | No | Available: StorageBinsList, StorageBinsList.InventoryList, StorageBinsList.InventoryList.StockMovementsList, StorageBinsList.InventoryList.WarehouseIdDetail, StorageBinsList.InventoryList.BinIdDetail, StorageBinsList.ZoneIdDetail, StorageBinsList.ZoneIdDetail.StorageBinsList, StorageBinsList.ZoneIdDetail.WarehouseIdDetail, WarehouseIdDetail, WarehouseIdDetail.StorageZonesList, WarehouseIdDetail.StorageZonesList.StorageBinsList, WarehouseIdDetail.StorageZonesList.WarehouseIdDetail, WarehouseIdDetail.InventoryList, WarehouseIdDetail.InventoryList.StockMovementsList, WarehouseIdDetail.InventoryList.WarehouseIdDetail, WarehouseIdDetail.InventoryList.BinIdDetail, WarehouseIdDetail.PurchaseOrdersList, WarehouseIdDetail.PurchaseOrdersList.PurchaseOrderItemsList, WarehouseIdDetail.PurchaseOrdersList.SupplierIdDetail, WarehouseIdDetail.PurchaseOrdersList.WarehouseIdDetail, WarehouseIdDetail.ShipmentsList, WarehouseIdDetail.ShipmentsList.ShipmentItemsList, WarehouseIdDetail.ShipmentsList.ShipmentTrackingList, WarehouseIdDetail.ShipmentsList.WarehouseIdDetail |
| `joins` | `string` | No | Available: Warehouses, Warehouses.Organizations, Warehouses.Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `warehouseId` | `string (uuid)` | No | Filter by warehouse_id |
| `zoneCode` | `string` | No | Filter by zone_code |
| `zoneType` | `string` | No | Filter by zone_type |
| `temperatureMin` | `number` | No | Filter by temperature_min |
| `temperatureMax` | `number` | No | Filter by temperature_max |

**Response:** `PaginationResponse<StorageZones>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/storage-zones/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate StorageZones (POST)

```
POST /api/v1/storage-zones/pagination
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

**Response:** `PaginationResponse<StorageZones>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/storage-zones/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create StorageZones

```
POST /api/v1/storage-zones/
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
  warehouseId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  zoneCode?: string  // e.g. example_zone_code
  zoneType?: string  // e.g. example_zone_type
  temperatureMin?: number  // e.g. 99.99
  temperatureMax?: number  // e.g. 99.99
}
```

**Response:** `StorageZones`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/storage-zones/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create StorageZones

```
POST /api/v1/storage-zones/bulk/
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
  warehouseId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  zoneCode?: string  // e.g. example_zone_code
  zoneType?: string  // e.g. example_zone_type
  temperatureMin?: number  // e.g. 99.99
  temperatureMax?: number  // e.g. 99.99
}
```

**Response:** `StorageZones[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/storage-zones/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find StorageZones by ID

```
GET /api/v1/storage-zones/with-id/:id
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

**Response:** `StorageZones`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/storage-zones/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update StorageZones

```
PUT /api/v1/storage-zones/with-id/:id
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
  warehouseId?: string
  zoneCode?: string
  zoneType?: string
  temperatureMin?: number
  temperatureMax?: number
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
  "http://localhost:3000/api/v1/storage-zones/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update StorageZones

```
PUT /api/v1/storage-zones/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: StorageZonesEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/storage-zones/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete StorageZones

```
DELETE /api/v1/storage-zones/with-id/:id
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
  "http://localhost:3000/api/v1/storage-zones/with-id/:id"
```

</details>

---

:::


::::
