---
title: Inventory
---

# Inventory

**Table:** `logistics.inventory`

**Base path:** `/inventory`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `products` | `product_id` | `inventory_product_id_fkey` | [Products](./products) |
| `product_variants` | `variant_id` | `inventory_variant_id_fkey` | [ProductVariants](./product_variants) |
| `warehouses` | `warehouse_id` | `inventory_warehouse_id_fkey` | [Warehouses](./warehouses) |
| `storage_bins` | `bin_id` | `inventory_bin_id_fkey` | [StorageBins](./storage_bins) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `stock_movements` | `inventory_id` | `stock_movements_inventory_id_fkey` | [StockMovements](./stock_movements) |


## Entity Relationship Diagram

erDiagram
    Inventory }o--|| Products : "FK"
    Inventory }o--|| ProductVariants : "FK"
    Inventory }o--|| Warehouses : "FK"
    Inventory }o--|| StorageBins : "FK"
    Inventory ||--o{ StockMovements : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `product_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` `UQ` | → References `products` |
| 4 | `variant_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `UQ` `FK` | → References `product_variants` |
| 5 | `warehouse_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `UQ` `FK` | → References `warehouses` |
| 6 | `bin_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `storage_bins` |
| 7 | `quantity` | `integer` | `int` | `number` | NO | `0` | - | - |
| 8 | `reserved` | `integer` | `int` | `number` | NO | `0` | - | - |
| 9 | `reorder_level` | `integer` | `int` | `number` | NO | `10` | - | - |
| 10 | `reorder_quantity` | `integer` | `int` | `number` | NO | `50` | - | - |
| 11 | `last_counted_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 12 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 13 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `product_id` | `products` | `inventory_product_id_fkey` |
| `variant_id` | `product_variants` | `inventory_variant_id_fkey` |
| `warehouse_id` | `warehouses` | `inventory_warehouse_id_fkey` |
| `bin_id` | `storage_bins` | `inventory_bin_id_fkey` |

## Unique Keys

- `product_id` (`uuid`)
- `variant_id` (`uuid`)
- `warehouse_id` (`uuid`)


## Go Generated Code

> 📂 Source: [📄 `Inventory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Inventory.go) · [📄 `Inventory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Inventory.go) · [📄 `Inventory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/controllers/Inventory.go)

### Structs

:::tabs

== Form

#### InventoryForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Inventory.go#:~:text=type%20InventoryForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `WarehouseId` | `uuid.UUID` | `warehouseId` | NO |
| `BinId` | `*uuid.UUID` | `binId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `Reserved` | `int` | `reserved` | NO |
| `ReorderLevel` | `int` | `reorderLevel` | NO |
| `ReorderQuantity` | `int` | `reorderQuantity` | NO |
| `LastCountedAt` | `*time.Time` | `lastCountedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### Inventory [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Inventory.go#:~:text=type%20Inventory%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `WarehouseId` | `uuid.UUID` | `warehouseId` | NO |
| `BinId` | `*uuid.UUID` | `binId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `Reserved` | `int` | `reserved` | NO |
| `ReorderLevel` | `int` | `reorderLevel` | NO |
| `ReorderQuantity` | `int` | `reorderQuantity` | NO |
| `LastCountedAt` | `*time.Time` | `lastCountedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### InventoryEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Inventory.go#:~:text=type%20InventoryEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `WarehouseId` | `*uuid.UUID` | `warehouseId` | YES |
| `BinId` | `*uuid.UUID` | `binId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `Reserved` | `*int` | `reserved` | YES |
| `ReorderLevel` | `*int` | `reorderLevel` | YES |
| `ReorderQuantity` | `*int` | `reorderQuantity` | YES |
| `LastCountedAt` | `*time.Time` | `lastCountedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### InventoryFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Inventory.go#:~:text=type%20InventoryFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ProductId` | `*uuid.UUID` | `productId` | YES |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `WarehouseId` | `*uuid.UUID` | `warehouseId` | YES |
| `BinId` | `*uuid.UUID` | `binId` | YES |
| `Quantity` | `*int` | `quantity` | YES |
| `Reserved` | `*int` | `reserved` | YES |
| `ReorderLevel` | `*int` | `reorderLevel` | YES |
| `ReorderQuantity` | `*int` | `reorderQuantity` | YES |
| `LastCountedAt` | `*time.Time` | `lastCountedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### InventoryPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Inventory.go#:~:text=type%20InventoryPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ProductId` | `uuid.UUID` | `productId` | NO |
| `VariantId` | `*uuid.UUID` | `variantId` | YES |
| `WarehouseId` | `uuid.UUID` | `warehouseId` | NO |
| `BinId` | `*uuid.UUID` | `binId` | YES |
| `Quantity` | `int` | `quantity` | NO |
| `Reserved` | `int` | `reserved` | NO |
| `ReorderLevel` | `int` | `reorderLevel` | NO |
| `ReorderQuantity` | `int` | `reorderQuantity` | NO |
| `LastCountedAt` | `*time.Time` | `lastCountedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### InventoryBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Inventory.go#:~:text=type%20InventoryBatchUpdate%20struct)

```go
type InventoryBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Inventory.go#:~:text=)%20CreateInventory() | `(InventoryService) CreateInventory(data InventoryForm) (InventoryForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Inventory.go#:~:text=)%20CreateInventoryMultiple() | `(InventoryService) CreateInventoryMultiple(data []InventoryForm) ([]InventoryForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Inventory.go#:~:text=)%20UpdateInventory() | `(InventoryService) UpdateInventory(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Inventory.go#:~:text=)%20UpdateInventoryMultiple() | `(InventoryService) UpdateInventoryMultiple(data []InventoryBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Inventory.go#:~:text=)%20DeleteInventory() | `(InventoryService) DeleteInventory(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/inventory/` | Search with query params |
| `GET` | `/inventory/pagination` | Paginated listing |
| `POST` | `/inventory/` | Create single record |
| `POST` | `/inventory/bulk/` | Create multiple records |
| `PUT` | `/inventory/bulk/` | Batch update |
| `GET` | `/inventory/with-id/:id` | Get by ID |
| `PUT` | `/inventory/with-id/:id` | Update by ID |
| `DELETE` | `/inventory/with-id/:id` | Delete by ID |

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
export interface Inventory {
  id: string;
  name: string;
  productId: string;
  variantId?: string;
  warehouseId: string;
  binId?: string;
  quantity: number;
  reserved: number;
  reorderLevel: number;
  reorderQuantity: number;
  lastCountedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface InventoryForm {
  name: string;
  productId: string;
  variantId?: string;
  warehouseId: string;
  binId?: string;
  quantity: number;
  reserved: number;
  reorderLevel: number;
  reorderQuantity: number;
  lastCountedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface InventoryEdit {
  id: string;
  name: string;
  productId: string;
  variantId?: string;
  warehouseId: string;
  binId?: string;
  quantity: number;
  reserved: number;
  reorderLevel: number;
  reorderQuantity: number;
  lastCountedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface InventoryPage {
  data: Inventory[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type InventoryPathQuery = {
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

const InventoryKeys = {
  all: ["inventory"] as const,
  lists: () => [...InventoryKeys.all, "list"] as const,
  detail: (id: any) => [...InventoryKeys.all, "detail", id] as const,
} as const;

export function useInventoryList(query?: InventoryPathQuery) {
  return useQuery({
    queryKey: [...InventoryKeys.lists(), query],
    queryFn: () => fetch(`/inventory/pagination`, { method: "GET" }).then(r => r.json()) as Promise<InventoryPage>,
  });
}

export function useInventoryDetail(id: any) {
  return useQuery({
    queryKey: InventoryKeys.detail(id),
    queryFn: () => fetch(`/inventory/with-id/:id`).then(r => r.json()) as Promise<Inventory>,
  });
}

export function useCreateInventory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: InventoryForm) =>
      fetch("/inventory/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: InventoryKeys.lists() }),
  });
}

export function useUpdateInventory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: InventoryEdit }) =>
      fetch(`/inventory/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: InventoryKeys.all }),
  });
}

export function useDeleteInventory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/inventory/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: InventoryKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const InventoryFormSchema = z.object({
  name: z.string(),
  productId: z.string().uuid(),
  variantId: z.string().uuid().optional(),
  warehouseId: z.string().uuid(),
  binId: z.string().uuid().optional(),
  quantity: z.number().int(),
  reserved: z.number().int(),
  reorderLevel: z.number().int(),
  reorderQuantity: z.number().int(),
  lastCountedAt: z.string().datetime().optional(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type InventoryFormInput = z.infer<typeof InventoryFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './inventory.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Inventory

```
GET /api/v1/inventory/
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
| `preloads` | `string` | No | Available: StockMovementsList, StockMovementsList.InventoryIdDetail, StockMovementsList.InventoryIdDetail.StockMovementsList, StockMovementsList.InventoryIdDetail.WarehouseIdDetail, StockMovementsList.InventoryIdDetail.BinIdDetail, WarehouseIdDetail, WarehouseIdDetail.StorageZonesList, WarehouseIdDetail.StorageZonesList.StorageBinsList, WarehouseIdDetail.StorageZonesList.WarehouseIdDetail, WarehouseIdDetail.InventoryList, WarehouseIdDetail.InventoryList.StockMovementsList, WarehouseIdDetail.InventoryList.WarehouseIdDetail, WarehouseIdDetail.InventoryList.BinIdDetail, WarehouseIdDetail.PurchaseOrdersList, WarehouseIdDetail.PurchaseOrdersList.PurchaseOrderItemsList, WarehouseIdDetail.PurchaseOrdersList.SupplierIdDetail, WarehouseIdDetail.PurchaseOrdersList.WarehouseIdDetail, WarehouseIdDetail.ShipmentsList, WarehouseIdDetail.ShipmentsList.ShipmentItemsList, WarehouseIdDetail.ShipmentsList.ShipmentTrackingList, WarehouseIdDetail.ShipmentsList.WarehouseIdDetail, BinIdDetail, BinIdDetail.InventoryList, BinIdDetail.InventoryList.StockMovementsList, BinIdDetail.InventoryList.WarehouseIdDetail, BinIdDetail.InventoryList.BinIdDetail, BinIdDetail.ZoneIdDetail, BinIdDetail.ZoneIdDetail.StorageBinsList, BinIdDetail.ZoneIdDetail.WarehouseIdDetail |
| `joins` | `string` | No | Available: Products, ProductVariants, Warehouses, Warehouses.Organizations, Warehouses.Users, StorageBins, StorageBins.StorageZones, StorageBins.StorageZones.Warehouses |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `variantId` | `string (uuid)` | No | Filter by variant_id |
| `warehouseId` | `string (uuid)` | No | Filter by warehouse_id |
| `binId` | `string (uuid)` | No | Filter by bin_id |
| `quantity` | `integer` | No | Filter by quantity |
| `reserved` | `integer` | No | Filter by reserved |
| `reorderLevel` | `integer` | No | Filter by reorder_level |
| `reorderQuantity` | `integer` | No | Filter by reorder_quantity |
| `lastCountedAt` | `string (date-time)` | No | Filter by last_counted_at |

**Response:** `Inventory[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/inventory/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Inventory (POST)

```
POST /api/v1/inventory/search
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

**Response:** `Inventory[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/inventory/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Inventory

```
GET /api/v1/inventory/pagination
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
| `preloads` | `string` | No | Available: StockMovementsList, StockMovementsList.InventoryIdDetail, StockMovementsList.InventoryIdDetail.StockMovementsList, StockMovementsList.InventoryIdDetail.WarehouseIdDetail, StockMovementsList.InventoryIdDetail.BinIdDetail, WarehouseIdDetail, WarehouseIdDetail.StorageZonesList, WarehouseIdDetail.StorageZonesList.StorageBinsList, WarehouseIdDetail.StorageZonesList.WarehouseIdDetail, WarehouseIdDetail.InventoryList, WarehouseIdDetail.InventoryList.StockMovementsList, WarehouseIdDetail.InventoryList.WarehouseIdDetail, WarehouseIdDetail.InventoryList.BinIdDetail, WarehouseIdDetail.PurchaseOrdersList, WarehouseIdDetail.PurchaseOrdersList.PurchaseOrderItemsList, WarehouseIdDetail.PurchaseOrdersList.SupplierIdDetail, WarehouseIdDetail.PurchaseOrdersList.WarehouseIdDetail, WarehouseIdDetail.ShipmentsList, WarehouseIdDetail.ShipmentsList.ShipmentItemsList, WarehouseIdDetail.ShipmentsList.ShipmentTrackingList, WarehouseIdDetail.ShipmentsList.WarehouseIdDetail, BinIdDetail, BinIdDetail.InventoryList, BinIdDetail.InventoryList.StockMovementsList, BinIdDetail.InventoryList.WarehouseIdDetail, BinIdDetail.InventoryList.BinIdDetail, BinIdDetail.ZoneIdDetail, BinIdDetail.ZoneIdDetail.StorageBinsList, BinIdDetail.ZoneIdDetail.WarehouseIdDetail |
| `joins` | `string` | No | Available: Products, ProductVariants, Warehouses, Warehouses.Organizations, Warehouses.Users, StorageBins, StorageBins.StorageZones, StorageBins.StorageZones.Warehouses |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `productId` | `string (uuid)` | No | Filter by product_id |
| `variantId` | `string (uuid)` | No | Filter by variant_id |
| `warehouseId` | `string (uuid)` | No | Filter by warehouse_id |
| `binId` | `string (uuid)` | No | Filter by bin_id |
| `quantity` | `integer` | No | Filter by quantity |
| `reserved` | `integer` | No | Filter by reserved |
| `reorderLevel` | `integer` | No | Filter by reorder_level |
| `reorderQuantity` | `integer` | No | Filter by reorder_quantity |
| `lastCountedAt` | `string (date-time)` | No | Filter by last_counted_at |

**Response:** `PaginationResponse<Inventory>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/inventory/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Inventory (POST)

```
POST /api/v1/inventory/pagination
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

**Response:** `PaginationResponse<Inventory>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/inventory/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Inventory

```
POST /api/v1/inventory/
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
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  variantId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  warehouseId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  binId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
  reserved?: number  // e.g. 1
  reorderLevel?: number  // e.g. 1
  reorderQuantity?: number  // e.g. 1
  lastCountedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Inventory`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/inventory/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Inventory

```
POST /api/v1/inventory/bulk/
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
  productId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  variantId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  warehouseId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  binId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  quantity?: number  // e.g. 1
  reserved?: number  // e.g. 1
  reorderLevel?: number  // e.g. 1
  reorderQuantity?: number  // e.g. 1
  lastCountedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Inventory[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/inventory/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Inventory by ID

```
GET /api/v1/inventory/with-id/:id
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

**Response:** `Inventory`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/inventory/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Inventory

```
PUT /api/v1/inventory/with-id/:id
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
  variantId?: string
  warehouseId?: string
  binId?: string
  quantity?: number
  reserved?: number
  reorderLevel?: number
  reorderQuantity?: number
  lastCountedAt?: string
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
  "http://localhost:3000/api/v1/inventory/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Inventory

```
PUT /api/v1/inventory/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: InventoryEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/inventory/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Inventory

```
DELETE /api/v1/inventory/with-id/:id
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
  "http://localhost:3000/api/v1/inventory/with-id/:id"
```

</details>

---

:::


::::
