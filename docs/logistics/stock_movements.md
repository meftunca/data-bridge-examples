---
title: StockMovements
---

# StockMovements

**Table:** `logistics.stock_movements`

**Base path:** `/stock-movements`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `inventory` | `inventory_id` | `stock_movements_inventory_id_fkey` | [Inventory](./inventory) |
| `users` | `performed_by` | `stock_movements_performed_by_fkey` | [Users](./users) |


## Entity Relationship Diagram

erDiagram
    StockMovements }o--|| Inventory : "FK"
    StockMovements }o--|| Users : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `inventory_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `inventory` |
| 4 | `movement_type` | `text` | `string` | `string` | NO | `'adjustment'::text` | - | - |
| 5 | `quantity` | `integer` | `int` | `number` | NO | `0` | - | - |
| 6 | `reference_type` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `reference_id` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 8 | `performed_by` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `users` |
| 9 | `notes` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 10 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `inventory_id` | `inventory` | `stock_movements_inventory_id_fkey` |
| `performed_by` | `users` | `stock_movements_performed_by_fkey` |


## Go Generated Code

> 📂 Source: [📄 `StockMovements.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StockMovements.go) · [📄 `StockMovements.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StockMovements.go) · [📄 `StockMovements.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/controllers/StockMovements.go)

### Structs

:::tabs

== Form

#### StockMovementsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StockMovements.go#:~:text=type%20StockMovementsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `InventoryId` | `uuid.UUID` | `inventoryId` | NO |
| `MovementType` | `string` | `movementType` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `ReferenceType` | `string` | `referenceType` | NO |
| `ReferenceId` | `string` | `referenceId` | NO |
| `PerformedBy` | `*uuid.UUID` | `performedBy` | YES |
| `Notes` | `string` | `notes` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### StockMovements [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StockMovements.go#:~:text=type%20StockMovements%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `InventoryId` | `uuid.UUID` | `inventoryId` | NO |
| `MovementType` | `string` | `movementType` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `ReferenceType` | `string` | `referenceType` | NO |
| `ReferenceId` | `string` | `referenceId` | NO |
| `PerformedBy` | `*uuid.UUID` | `performedBy` | YES |
| `Notes` | `string` | `notes` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### StockMovementsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StockMovements.go#:~:text=type%20StockMovementsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `InventoryId` | `*uuid.UUID` | `inventoryId` | YES |
| `MovementType` | `*string` | `movementType` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `ReferenceType` | `*string` | `referenceType` | YES |
| `ReferenceId` | `*string` | `referenceId` | YES |
| `PerformedBy` | `*uuid.UUID` | `performedBy` | YES |
| `Notes` | `*string` | `notes` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### StockMovementsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StockMovements.go#:~:text=type%20StockMovementsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `InventoryId` | `*uuid.UUID` | `inventoryId` | YES |
| `MovementType` | `*string` | `movementType` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `ReferenceType` | `*string` | `referenceType` | YES |
| `ReferenceId` | `*string` | `referenceId` | YES |
| `PerformedBy` | `*uuid.UUID` | `performedBy` | YES |
| `Notes` | `*string` | `notes` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### StockMovementsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StockMovements.go#:~:text=type%20StockMovementsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `InventoryId` | `uuid.UUID` | `inventoryId` | NO |
| `MovementType` | `string` | `movementType` | NO |
| `Quantity` | `int` | `quantity` | NO |
| `ReferenceType` | `string` | `referenceType` | NO |
| `ReferenceId` | `string` | `referenceId` | NO |
| `PerformedBy` | `*uuid.UUID` | `performedBy` | YES |
| `Notes` | `string` | `notes` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### StockMovementsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/StockMovements.go#:~:text=type%20StockMovementsBatchUpdate%20struct)

```go
type StockMovementsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StockMovements.go#:~:text=)%20CreateStockMovements() | `(StockMovementsService) CreateStockMovements(data StockMovementsForm) (StockMovementsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StockMovements.go#:~:text=)%20CreateStockMovementsMultiple() | `(StockMovementsService) CreateStockMovementsMultiple(data []StockMovementsForm) ([]StockMovementsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StockMovements.go#:~:text=)%20UpdateStockMovements() | `(StockMovementsService) UpdateStockMovements(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StockMovements.go#:~:text=)%20UpdateStockMovementsMultiple() | `(StockMovementsService) UpdateStockMovementsMultiple(data []StockMovementsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/StockMovements.go#:~:text=)%20DeleteStockMovements() | `(StockMovementsService) DeleteStockMovements(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/stock-movements/` | Search with query params |
| `GET` | `/stock-movements/pagination` | Paginated listing |
| `POST` | `/stock-movements/` | Create single record |
| `POST` | `/stock-movements/bulk/` | Create multiple records |
| `PUT` | `/stock-movements/bulk/` | Batch update |
| `GET` | `/stock-movements/with-id/:id` | Get by ID |
| `PUT` | `/stock-movements/with-id/:id` | Update by ID |
| `DELETE` | `/stock-movements/with-id/:id` | Delete by ID |

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
export interface StockMovements {
  id: string;
  name: string;
  inventoryId: string;
  movementType: string;
  quantity: number;
  referenceType: string;
  referenceId: string;
  performedBy?: string;
  notes: string;
  createdAt: string;
}

export interface StockMovementsForm {
  name: string;
  inventoryId: string;
  movementType: string;
  quantity: number;
  referenceType: string;
  referenceId: string;
  performedBy?: string;
  notes: string;
  createdAt: string;
}

export interface StockMovementsEdit {
  id: string;
  name: string;
  inventoryId: string;
  movementType: string;
  quantity: number;
  referenceType: string;
  referenceId: string;
  performedBy?: string;
  notes: string;
  createdAt: string;
}

export interface StockMovementsPage {
  data: StockMovements[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type StockMovementsPathQuery = {
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

const StockMovementsKeys = {
  all: ["stock_movements"] as const,
  lists: () => [...StockMovementsKeys.all, "list"] as const,
  detail: (id: any) => [...StockMovementsKeys.all, "detail", id] as const,
} as const;

export function useStockMovementsList(query?: StockMovementsPathQuery) {
  return useQuery({
    queryKey: [...StockMovementsKeys.lists(), query],
    queryFn: () => fetch(`/stock-movements/pagination`, { method: "GET" }).then(r => r.json()) as Promise<StockMovementsPage>,
  });
}

export function useStockMovementsDetail(id: any) {
  return useQuery({
    queryKey: StockMovementsKeys.detail(id),
    queryFn: () => fetch(`/stock-movements/with-id/:id`).then(r => r.json()) as Promise<StockMovements>,
  });
}

export function useCreateStockMovements() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: StockMovementsForm) =>
      fetch("/stock-movements/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: StockMovementsKeys.lists() }),
  });
}

export function useUpdateStockMovements() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: StockMovementsEdit }) =>
      fetch(`/stock-movements/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: StockMovementsKeys.all }),
  });
}

export function useDeleteStockMovements() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/stock-movements/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: StockMovementsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const StockMovementsFormSchema = z.object({
  name: z.string(),
  inventoryId: z.string().uuid(),
  movementType: z.string(),
  quantity: z.number().int(),
  referenceType: z.string(),
  referenceId: z.string(),
  performedBy: z.string().uuid().optional(),
  notes: z.string(),
  createdAt: z.string().datetime(),
});

export type StockMovementsFormInput = z.infer<typeof StockMovementsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './stock_movements.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search StockMovements

```
GET /api/v1/stock-movements/
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
| `preloads` | `string` | No | Available: InventoryIdDetail, InventoryIdDetail.StockMovementsList, InventoryIdDetail.StockMovementsList.InventoryIdDetail, InventoryIdDetail.WarehouseIdDetail, InventoryIdDetail.WarehouseIdDetail.StorageZonesList, InventoryIdDetail.WarehouseIdDetail.InventoryList, InventoryIdDetail.WarehouseIdDetail.PurchaseOrdersList, InventoryIdDetail.WarehouseIdDetail.ShipmentsList, InventoryIdDetail.BinIdDetail, InventoryIdDetail.BinIdDetail.InventoryList, InventoryIdDetail.BinIdDetail.ZoneIdDetail |
| `joins` | `string` | No | Available: Inventory, Inventory.Products, Inventory.ProductVariants, Inventory.Warehouses, Inventory.Warehouses.Organizations, Inventory.Warehouses.Users, Inventory.StorageBins, Inventory.StorageBins.StorageZones, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `inventoryId` | `string (uuid)` | No | Filter by inventory_id |
| `movementType` | `string` | No | Filter by movement_type |
| `quantity` | `integer` | No | Filter by quantity |
| `referenceType` | `string` | No | Filter by reference_type |
| `referenceId` | `string` | No | Filter by reference_id |
| `performedBy` | `string (uuid)` | No | Filter by performed_by |
| `notes` | `string` | No | Filter by notes |

**Response:** `StockMovements[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/stock-movements/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search StockMovements (POST)

```
POST /api/v1/stock-movements/search
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

**Response:** `StockMovements[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/stock-movements/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate StockMovements

```
GET /api/v1/stock-movements/pagination
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
| `preloads` | `string` | No | Available: InventoryIdDetail, InventoryIdDetail.StockMovementsList, InventoryIdDetail.StockMovementsList.InventoryIdDetail, InventoryIdDetail.WarehouseIdDetail, InventoryIdDetail.WarehouseIdDetail.StorageZonesList, InventoryIdDetail.WarehouseIdDetail.InventoryList, InventoryIdDetail.WarehouseIdDetail.PurchaseOrdersList, InventoryIdDetail.WarehouseIdDetail.ShipmentsList, InventoryIdDetail.BinIdDetail, InventoryIdDetail.BinIdDetail.InventoryList, InventoryIdDetail.BinIdDetail.ZoneIdDetail |
| `joins` | `string` | No | Available: Inventory, Inventory.Products, Inventory.ProductVariants, Inventory.Warehouses, Inventory.Warehouses.Organizations, Inventory.Warehouses.Users, Inventory.StorageBins, Inventory.StorageBins.StorageZones, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `inventoryId` | `string (uuid)` | No | Filter by inventory_id |
| `movementType` | `string` | No | Filter by movement_type |
| `quantity` | `integer` | No | Filter by quantity |
| `referenceType` | `string` | No | Filter by reference_type |
| `referenceId` | `string` | No | Filter by reference_id |
| `performedBy` | `string (uuid)` | No | Filter by performed_by |
| `notes` | `string` | No | Filter by notes |

**Response:** `PaginationResponse<StockMovements>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/stock-movements/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate StockMovements (POST)

```
POST /api/v1/stock-movements/pagination
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

**Response:** `PaginationResponse<StockMovements>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/stock-movements/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create StockMovements

```
POST /api/v1/stock-movements/
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
  inventoryId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  movementType?: string  // e.g. example_movement_type
  quantity?: number  // e.g. 1
  referenceType?: string  // e.g. example_reference_type
  referenceId?: string  // e.g. example_reference_id
  performedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  notes?: string  // e.g. example_notes
}
```

**Response:** `StockMovements`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/stock-movements/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create StockMovements

```
POST /api/v1/stock-movements/bulk/
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
  inventoryId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  movementType?: string  // e.g. example_movement_type
  quantity?: number  // e.g. 1
  referenceType?: string  // e.g. example_reference_type
  referenceId?: string  // e.g. example_reference_id
  performedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  notes?: string  // e.g. example_notes
}
```

**Response:** `StockMovements[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/stock-movements/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find StockMovements by ID

```
GET /api/v1/stock-movements/with-id/:id
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

**Response:** `StockMovements`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/stock-movements/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update StockMovements

```
PUT /api/v1/stock-movements/with-id/:id
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
  inventoryId?: string
  movementType?: string
  quantity?: number
  referenceType?: string
  referenceId?: string
  performedBy?: string
  notes?: string
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
  "http://localhost:3000/api/v1/stock-movements/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update StockMovements

```
PUT /api/v1/stock-movements/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: StockMovementsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/stock-movements/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete StockMovements

```
DELETE /api/v1/stock-movements/with-id/:id
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
  "http://localhost:3000/api/v1/stock-movements/with-id/:id"
```

</details>

---

:::


::::
