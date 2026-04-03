---
title: Warehouses
---

# Warehouses

**Table:** `logistics.warehouses`

**Base path:** `/warehouses`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `organizations` | `organization_id` | `warehouses_organization_id_fkey` | [Organizations](./organizations) |
| `users` | `manager_id` | `warehouses_manager_id_fkey` | [Users](./users) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `inventory` | `warehouse_id` | `inventory_warehouse_id_fkey` | [Inventory](./inventory) |
| `purchase_orders` | `warehouse_id` | `purchase_orders_warehouse_id_fkey` | [PurchaseOrders](./purchase_orders) |
| `shipments` | `warehouse_id` | `shipments_warehouse_id_fkey` | [Shipments](./shipments) |
| `storage_zones` | `warehouse_id` | `storage_zones_warehouse_id_fkey` | [StorageZones](./storage_zones) |


## Entity Relationship Diagram

erDiagram
    Warehouses }o--|| Organizations : "FK"
    Warehouses }o--|| Users : "FK"
    Warehouses ||--o{ Inventory : "ref"
    Warehouses ||--o{ PurchaseOrders : "ref"
    Warehouses ||--o{ Shipments : "ref"
    Warehouses ||--o{ StorageZones : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `code` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 4 | `address` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 5 | `organization_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `organizations` |
| 6 | `manager_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `users` |
| 7 | `is_active` | `boolean` | `bool` | `boolean` | NO | `true` | - | - |
| 8 | `capacity` | `integer` | `int` | `number` | NO | `0` | - | - |
| 9 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 10 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `organization_id` | `organizations` | `warehouses_organization_id_fkey` |
| `manager_id` | `users` | `warehouses_manager_id_fkey` |

## Unique Keys

- `code` (`text`)


## Go Generated Code

> 📂 Source: [📄 `Warehouses.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Warehouses.go) · [📄 `Warehouses.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Warehouses.go) · [📄 `Warehouses.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/controllers/Warehouses.go)

### Structs

:::tabs

== Form

#### WarehousesForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Warehouses.go#:~:text=type%20WarehousesForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Code` | `string` | `code` | NO |
| `Address` | `json.RawMessage` | `address` | NO |
| `OrganizationId` | `uuid.UUID` | `organizationId` | NO |
| `ManagerId` | `*uuid.UUID` | `managerId` | YES |
| `IsActive` | `bool` | `isActive` | NO |
| `Capacity` | `int` | `capacity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### Warehouses [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Warehouses.go#:~:text=type%20Warehouses%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Code` | `string` | `code` | NO |
| `Address` | `json.RawMessage` | `address` | NO |
| `OrganizationId` | `uuid.UUID` | `organizationId` | NO |
| `ManagerId` | `*uuid.UUID` | `managerId` | YES |
| `IsActive` | `bool` | `isActive` | NO |
| `Capacity` | `int` | `capacity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### WarehousesEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Warehouses.go#:~:text=type%20WarehousesEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Code` | `*string` | `code` | YES |
| `Address` | `*json.RawMessage` | `address` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `ManagerId` | `*uuid.UUID` | `managerId` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `Capacity` | `*int` | `capacity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### WarehousesFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Warehouses.go#:~:text=type%20WarehousesFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Code` | `*string` | `code` | YES |
| `Address` | `*json.RawMessage` | `address` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `ManagerId` | `*uuid.UUID` | `managerId` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `Capacity` | `*int` | `capacity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### WarehousesPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Warehouses.go#:~:text=type%20WarehousesPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Code` | `string` | `code` | NO |
| `Address` | `json.RawMessage` | `address` | NO |
| `OrganizationId` | `uuid.UUID` | `organizationId` | NO |
| `ManagerId` | `*uuid.UUID` | `managerId` | YES |
| `IsActive` | `bool` | `isActive` | NO |
| `Capacity` | `int` | `capacity` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### WarehousesBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Warehouses.go#:~:text=type%20WarehousesBatchUpdate%20struct)

```go
type WarehousesBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Warehouses.go#:~:text=)%20CreateWarehouses() | `(WarehousesService) CreateWarehouses(data WarehousesForm) (WarehousesForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Warehouses.go#:~:text=)%20CreateWarehousesMultiple() | `(WarehousesService) CreateWarehousesMultiple(data []WarehousesForm) ([]WarehousesForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Warehouses.go#:~:text=)%20UpdateWarehouses() | `(WarehousesService) UpdateWarehouses(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Warehouses.go#:~:text=)%20UpdateWarehousesMultiple() | `(WarehousesService) UpdateWarehousesMultiple(data []WarehousesBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Warehouses.go#:~:text=)%20DeleteWarehouses() | `(WarehousesService) DeleteWarehouses(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/warehouses/` | Search with query params |
| `GET` | `/warehouses/pagination` | Paginated listing |
| `POST` | `/warehouses/` | Create single record |
| `POST` | `/warehouses/bulk/` | Create multiple records |
| `PUT` | `/warehouses/bulk/` | Batch update |
| `GET` | `/warehouses/with-id/:id` | Get by ID |
| `PUT` | `/warehouses/with-id/:id` | Update by ID |
| `DELETE` | `/warehouses/with-id/:id` | Delete by ID |

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
export interface Warehouses {
  id: string;
  name: string;
  code: string;
  address: Record<string, unknown>;
  organizationId: string;
  managerId?: string;
  isActive: boolean;
  capacity: number;
  createdAt: string;
  updatedAt: string;
}

export interface WarehousesForm {
  name: string;
  code: string;
  address: Record<string, unknown>;
  organizationId: string;
  managerId?: string;
  isActive: boolean;
  capacity: number;
  createdAt: string;
  updatedAt: string;
}

export interface WarehousesEdit {
  id: string;
  name: string;
  code: string;
  address: Record<string, unknown>;
  organizationId: string;
  managerId?: string;
  isActive: boolean;
  capacity: number;
  createdAt: string;
  updatedAt: string;
}

export interface WarehousesPage {
  data: Warehouses[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type WarehousesPathQuery = {
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

const WarehousesKeys = {
  all: ["warehouses"] as const,
  lists: () => [...WarehousesKeys.all, "list"] as const,
  detail: (id: any) => [...WarehousesKeys.all, "detail", id] as const,
} as const;

export function useWarehousesList(query?: WarehousesPathQuery) {
  return useQuery({
    queryKey: [...WarehousesKeys.lists(), query],
    queryFn: () => fetch(`/warehouses/pagination`, { method: "GET" }).then(r => r.json()) as Promise<WarehousesPage>,
  });
}

export function useWarehousesDetail(id: any) {
  return useQuery({
    queryKey: WarehousesKeys.detail(id),
    queryFn: () => fetch(`/warehouses/with-id/:id`).then(r => r.json()) as Promise<Warehouses>,
  });
}

export function useCreateWarehouses() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: WarehousesForm) =>
      fetch("/warehouses/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: WarehousesKeys.lists() }),
  });
}

export function useUpdateWarehouses() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: WarehousesEdit }) =>
      fetch(`/warehouses/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: WarehousesKeys.all }),
  });
}

export function useDeleteWarehouses() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/warehouses/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: WarehousesKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const WarehousesFormSchema = z.object({
  name: z.string(),
  code: z.string(),
  address: z.record(z.unknown()),
  organizationId: z.string().uuid(),
  managerId: z.string().uuid().optional(),
  isActive: z.boolean(),
  capacity: z.number().int(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type WarehousesFormInput = z.infer<typeof WarehousesFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './warehouses.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Warehouses

```
GET /api/v1/warehouses/
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
| `preloads` | `string` | No | Available: StorageZonesList, StorageZonesList.StorageBinsList, StorageZonesList.StorageBinsList.InventoryList, StorageZonesList.StorageBinsList.ZoneIdDetail, StorageZonesList.WarehouseIdDetail, StorageZonesList.WarehouseIdDetail.StorageZonesList, StorageZonesList.WarehouseIdDetail.InventoryList, StorageZonesList.WarehouseIdDetail.PurchaseOrdersList, StorageZonesList.WarehouseIdDetail.ShipmentsList, InventoryList, InventoryList.StockMovementsList, InventoryList.StockMovementsList.InventoryIdDetail, InventoryList.WarehouseIdDetail, InventoryList.WarehouseIdDetail.StorageZonesList, InventoryList.WarehouseIdDetail.InventoryList, InventoryList.WarehouseIdDetail.PurchaseOrdersList, InventoryList.WarehouseIdDetail.ShipmentsList, InventoryList.BinIdDetail, InventoryList.BinIdDetail.InventoryList, InventoryList.BinIdDetail.ZoneIdDetail, PurchaseOrdersList, PurchaseOrdersList.PurchaseOrderItemsList, PurchaseOrdersList.PurchaseOrderItemsList.PurchaseOrderIdDetail, PurchaseOrdersList.SupplierIdDetail, PurchaseOrdersList.SupplierIdDetail.PurchaseOrdersList, PurchaseOrdersList.WarehouseIdDetail, PurchaseOrdersList.WarehouseIdDetail.StorageZonesList, PurchaseOrdersList.WarehouseIdDetail.InventoryList, PurchaseOrdersList.WarehouseIdDetail.PurchaseOrdersList, PurchaseOrdersList.WarehouseIdDetail.ShipmentsList, ShipmentsList, ShipmentsList.ShipmentItemsList, ShipmentsList.ShipmentItemsList.ShipmentIdDetail, ShipmentsList.ShipmentTrackingList, ShipmentsList.ShipmentTrackingList.ShipmentIdDetail, ShipmentsList.WarehouseIdDetail, ShipmentsList.WarehouseIdDetail.StorageZonesList, ShipmentsList.WarehouseIdDetail.InventoryList, ShipmentsList.WarehouseIdDetail.PurchaseOrdersList, ShipmentsList.WarehouseIdDetail.ShipmentsList |
| `joins` | `string` | No | Available: Organizations, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `code` | `string` | No | Filter by code |
| `address` | `string` | No | Filter by address |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `managerId` | `string (uuid)` | No | Filter by manager_id |
| `isActive` | `boolean` | No | Filter by is_active |
| `capacity` | `integer` | No | Filter by capacity |

**Response:** `Warehouses[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/warehouses/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Warehouses (POST)

```
POST /api/v1/warehouses/search
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

**Response:** `Warehouses[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/warehouses/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Warehouses

```
GET /api/v1/warehouses/pagination
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
| `preloads` | `string` | No | Available: StorageZonesList, StorageZonesList.StorageBinsList, StorageZonesList.StorageBinsList.InventoryList, StorageZonesList.StorageBinsList.ZoneIdDetail, StorageZonesList.WarehouseIdDetail, StorageZonesList.WarehouseIdDetail.StorageZonesList, StorageZonesList.WarehouseIdDetail.InventoryList, StorageZonesList.WarehouseIdDetail.PurchaseOrdersList, StorageZonesList.WarehouseIdDetail.ShipmentsList, InventoryList, InventoryList.StockMovementsList, InventoryList.StockMovementsList.InventoryIdDetail, InventoryList.WarehouseIdDetail, InventoryList.WarehouseIdDetail.StorageZonesList, InventoryList.WarehouseIdDetail.InventoryList, InventoryList.WarehouseIdDetail.PurchaseOrdersList, InventoryList.WarehouseIdDetail.ShipmentsList, InventoryList.BinIdDetail, InventoryList.BinIdDetail.InventoryList, InventoryList.BinIdDetail.ZoneIdDetail, PurchaseOrdersList, PurchaseOrdersList.PurchaseOrderItemsList, PurchaseOrdersList.PurchaseOrderItemsList.PurchaseOrderIdDetail, PurchaseOrdersList.SupplierIdDetail, PurchaseOrdersList.SupplierIdDetail.PurchaseOrdersList, PurchaseOrdersList.WarehouseIdDetail, PurchaseOrdersList.WarehouseIdDetail.StorageZonesList, PurchaseOrdersList.WarehouseIdDetail.InventoryList, PurchaseOrdersList.WarehouseIdDetail.PurchaseOrdersList, PurchaseOrdersList.WarehouseIdDetail.ShipmentsList, ShipmentsList, ShipmentsList.ShipmentItemsList, ShipmentsList.ShipmentItemsList.ShipmentIdDetail, ShipmentsList.ShipmentTrackingList, ShipmentsList.ShipmentTrackingList.ShipmentIdDetail, ShipmentsList.WarehouseIdDetail, ShipmentsList.WarehouseIdDetail.StorageZonesList, ShipmentsList.WarehouseIdDetail.InventoryList, ShipmentsList.WarehouseIdDetail.PurchaseOrdersList, ShipmentsList.WarehouseIdDetail.ShipmentsList |
| `joins` | `string` | No | Available: Organizations, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `code` | `string` | No | Filter by code |
| `address` | `string` | No | Filter by address |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `managerId` | `string (uuid)` | No | Filter by manager_id |
| `isActive` | `boolean` | No | Filter by is_active |
| `capacity` | `integer` | No | Filter by capacity |

**Response:** `PaginationResponse<Warehouses>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/warehouses/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Warehouses (POST)

```
POST /api/v1/warehouses/pagination
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

**Response:** `PaginationResponse<Warehouses>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/warehouses/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Warehouses

```
POST /api/v1/warehouses/
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
  code: string  // e.g. example_code
  address?: Record<string, unknown>  // e.g. map[]
  organizationId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  managerId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  isActive?: boolean  // e.g. true
  capacity?: number  // e.g. 1
}
```

**Response:** `Warehouses`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/warehouses/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Warehouses

```
POST /api/v1/warehouses/bulk/
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
  code: string  // e.g. example_code
  address?: Record<string, unknown>  // e.g. map[]
  organizationId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  managerId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  isActive?: boolean  // e.g. true
  capacity?: number  // e.g. 1
}
```

**Response:** `Warehouses[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/warehouses/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Warehouses by ID

```
GET /api/v1/warehouses/with-id/:id
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

**Response:** `Warehouses`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/warehouses/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Warehouses

```
PUT /api/v1/warehouses/with-id/:id
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
  code?: string
  address?: Record<string, unknown>
  organizationId?: string
  managerId?: string
  isActive?: boolean
  capacity?: number
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
  "http://localhost:3000/api/v1/warehouses/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Warehouses

```
PUT /api/v1/warehouses/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: WarehousesEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/warehouses/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Warehouses

```
DELETE /api/v1/warehouses/with-id/:id
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
  "http://localhost:3000/api/v1/warehouses/with-id/:id"
```

</details>

---

:::


::::
