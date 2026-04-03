---
title: Suppliers
---

# Suppliers

**Table:** `logistics.suppliers`

**Base path:** `/suppliers`

## Related Tables

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `purchase_orders` | `supplier_id` | `purchase_orders_supplier_id_fkey` | [PurchaseOrders](./purchase_orders) |


## Entity Relationship Diagram

erDiagram
    Suppliers ||--o{ PurchaseOrders : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `code` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 4 | `contact_name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `contact_email` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 6 | `contact_phone` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `address` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 8 | `is_active` | `boolean` | `bool` | `boolean` | NO | `true` | - | - |
| 9 | `rating` | `numeric` | `float64` | `number` | NO | `0.00` | - | - |
| 10 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 11 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Unique Keys

- `code` (`text`)


## Go Generated Code

> 📂 Source: [📄 `Suppliers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Suppliers.go) · [📄 `Suppliers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Suppliers.go) · [📄 `Suppliers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/controllers/Suppliers.go)

### Structs

:::tabs

== Form

#### SuppliersForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Suppliers.go#:~:text=type%20SuppliersForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Code` | `string` | `code` | NO |
| `ContactName` | `string` | `contactName` | NO |
| `ContactEmail` | `string` | `contactEmail` | NO |
| `ContactPhone` | `string` | `contactPhone` | NO |
| `Address` | `json.RawMessage` | `address` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `Rating` | `float64` | `rating` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### Suppliers [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Suppliers.go#:~:text=type%20Suppliers%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Code` | `string` | `code` | NO |
| `ContactName` | `string` | `contactName` | NO |
| `ContactEmail` | `string` | `contactEmail` | NO |
| `ContactPhone` | `string` | `contactPhone` | NO |
| `Address` | `json.RawMessage` | `address` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `Rating` | `float64` | `rating` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### SuppliersEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Suppliers.go#:~:text=type%20SuppliersEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Code` | `*string` | `code` | YES |
| `ContactName` | `*string` | `contactName` | YES |
| `ContactEmail` | `*string` | `contactEmail` | YES |
| `ContactPhone` | `*string` | `contactPhone` | YES |
| `Address` | `*json.RawMessage` | `address` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `Rating` | `*float64` | `rating` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### SuppliersFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Suppliers.go#:~:text=type%20SuppliersFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Code` | `*string` | `code` | YES |
| `ContactName` | `*string` | `contactName` | YES |
| `ContactEmail` | `*string` | `contactEmail` | YES |
| `ContactPhone` | `*string` | `contactPhone` | YES |
| `Address` | `*json.RawMessage` | `address` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `Rating` | `*float64` | `rating` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### SuppliersPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Suppliers.go#:~:text=type%20SuppliersPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Code` | `string` | `code` | NO |
| `ContactName` | `string` | `contactName` | NO |
| `ContactEmail` | `string` | `contactEmail` | NO |
| `ContactPhone` | `string` | `contactPhone` | NO |
| `Address` | `json.RawMessage` | `address` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `Rating` | `float64` | `rating` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### SuppliersBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/structures/Suppliers.go#:~:text=type%20SuppliersBatchUpdate%20struct)

```go
type SuppliersBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Suppliers.go#:~:text=)%20CreateSuppliers() | `(SuppliersService) CreateSuppliers(data SuppliersForm) (SuppliersForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Suppliers.go#:~:text=)%20CreateSuppliersMultiple() | `(SuppliersService) CreateSuppliersMultiple(data []SuppliersForm) ([]SuppliersForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Suppliers.go#:~:text=)%20UpdateSuppliers() | `(SuppliersService) UpdateSuppliers(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Suppliers.go#:~:text=)%20UpdateSuppliersMultiple() | `(SuppliersService) UpdateSuppliersMultiple(data []SuppliersBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//logistics/services/Suppliers.go#:~:text=)%20DeleteSuppliers() | `(SuppliersService) DeleteSuppliers(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/suppliers/` | Search with query params |
| `GET` | `/suppliers/pagination` | Paginated listing |
| `POST` | `/suppliers/` | Create single record |
| `POST` | `/suppliers/bulk/` | Create multiple records |
| `PUT` | `/suppliers/bulk/` | Batch update |
| `GET` | `/suppliers/with-id/:id` | Get by ID |
| `PUT` | `/suppliers/with-id/:id` | Update by ID |
| `DELETE` | `/suppliers/with-id/:id` | Delete by ID |

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
export interface Suppliers {
  id: string;
  name: string;
  code: string;
  contactName: string;
  contactEmail: string;
  contactPhone: string;
  address: Record<string, unknown>;
  isActive: boolean;
  rating: number;
  createdAt: string;
  updatedAt: string;
}

export interface SuppliersForm {
  name: string;
  code: string;
  contactName: string;
  contactEmail: string;
  contactPhone: string;
  address: Record<string, unknown>;
  isActive: boolean;
  rating: number;
  createdAt: string;
  updatedAt: string;
}

export interface SuppliersEdit {
  id: string;
  name: string;
  code: string;
  contactName: string;
  contactEmail: string;
  contactPhone: string;
  address: Record<string, unknown>;
  isActive: boolean;
  rating: number;
  createdAt: string;
  updatedAt: string;
}

export interface SuppliersPage {
  data: Suppliers[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type SuppliersPathQuery = {
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

const SuppliersKeys = {
  all: ["suppliers"] as const,
  lists: () => [...SuppliersKeys.all, "list"] as const,
  detail: (id: any) => [...SuppliersKeys.all, "detail", id] as const,
} as const;

export function useSuppliersList(query?: SuppliersPathQuery) {
  return useQuery({
    queryKey: [...SuppliersKeys.lists(), query],
    queryFn: () => fetch(`/suppliers/pagination`, { method: "GET" }).then(r => r.json()) as Promise<SuppliersPage>,
  });
}

export function useSuppliersDetail(id: any) {
  return useQuery({
    queryKey: SuppliersKeys.detail(id),
    queryFn: () => fetch(`/suppliers/with-id/:id`).then(r => r.json()) as Promise<Suppliers>,
  });
}

export function useCreateSuppliers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: SuppliersForm) =>
      fetch("/suppliers/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: SuppliersKeys.lists() }),
  });
}

export function useUpdateSuppliers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: SuppliersEdit }) =>
      fetch(`/suppliers/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: SuppliersKeys.all }),
  });
}

export function useDeleteSuppliers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/suppliers/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: SuppliersKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const SuppliersFormSchema = z.object({
  name: z.string(),
  code: z.string(),
  contactName: z.string(),
  contactEmail: z.string(),
  contactPhone: z.string(),
  address: z.record(z.unknown()),
  isActive: z.boolean(),
  rating: z.number(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type SuppliersFormInput = z.infer<typeof SuppliersFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './suppliers.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Suppliers

```
GET /api/v1/suppliers/
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
| `preloads` | `string` | No | Available: PurchaseOrdersList, PurchaseOrdersList.PurchaseOrderItemsList, PurchaseOrdersList.PurchaseOrderItemsList.PurchaseOrderIdDetail, PurchaseOrdersList.SupplierIdDetail, PurchaseOrdersList.SupplierIdDetail.PurchaseOrdersList, PurchaseOrdersList.WarehouseIdDetail, PurchaseOrdersList.WarehouseIdDetail.StorageZonesList, PurchaseOrdersList.WarehouseIdDetail.InventoryList, PurchaseOrdersList.WarehouseIdDetail.PurchaseOrdersList, PurchaseOrdersList.WarehouseIdDetail.ShipmentsList |
| `joins` | `string` | No | Comma-separated relations to join |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `code` | `string` | No | Filter by code |
| `contactName` | `string` | No | Filter by contact_name |
| `contactEmail` | `string` | No | Filter by contact_email |
| `contactPhone` | `string` | No | Filter by contact_phone |
| `address` | `string` | No | Filter by address |
| `isActive` | `boolean` | No | Filter by is_active |
| `rating` | `number` | No | Filter by rating |

**Response:** `Suppliers[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/suppliers/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Suppliers (POST)

```
POST /api/v1/suppliers/search
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

**Response:** `Suppliers[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/suppliers/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Suppliers

```
GET /api/v1/suppliers/pagination
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
| `preloads` | `string` | No | Available: PurchaseOrdersList, PurchaseOrdersList.PurchaseOrderItemsList, PurchaseOrdersList.PurchaseOrderItemsList.PurchaseOrderIdDetail, PurchaseOrdersList.SupplierIdDetail, PurchaseOrdersList.SupplierIdDetail.PurchaseOrdersList, PurchaseOrdersList.WarehouseIdDetail, PurchaseOrdersList.WarehouseIdDetail.StorageZonesList, PurchaseOrdersList.WarehouseIdDetail.InventoryList, PurchaseOrdersList.WarehouseIdDetail.PurchaseOrdersList, PurchaseOrdersList.WarehouseIdDetail.ShipmentsList |
| `joins` | `string` | No | Comma-separated relations to join |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `code` | `string` | No | Filter by code |
| `contactName` | `string` | No | Filter by contact_name |
| `contactEmail` | `string` | No | Filter by contact_email |
| `contactPhone` | `string` | No | Filter by contact_phone |
| `address` | `string` | No | Filter by address |
| `isActive` | `boolean` | No | Filter by is_active |
| `rating` | `number` | No | Filter by rating |

**Response:** `PaginationResponse<Suppliers>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/suppliers/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Suppliers (POST)

```
POST /api/v1/suppliers/pagination
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

**Response:** `PaginationResponse<Suppliers>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/suppliers/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Suppliers

```
POST /api/v1/suppliers/
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
  contactName?: string  // e.g. example_contact_name
  contactEmail?: string  // e.g. example_contact_email
  contactPhone?: string  // e.g. example_contact_phone
  address?: Record<string, unknown>  // e.g. map[]
  isActive?: boolean  // e.g. true
  rating?: number  // e.g. 99.99
}
```

**Response:** `Suppliers`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/suppliers/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Suppliers

```
POST /api/v1/suppliers/bulk/
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
  contactName?: string  // e.g. example_contact_name
  contactEmail?: string  // e.g. example_contact_email
  contactPhone?: string  // e.g. example_contact_phone
  address?: Record<string, unknown>  // e.g. map[]
  isActive?: boolean  // e.g. true
  rating?: number  // e.g. 99.99
}
```

**Response:** `Suppliers[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/suppliers/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Suppliers by ID

```
GET /api/v1/suppliers/with-id/:id
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

**Response:** `Suppliers`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/suppliers/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Suppliers

```
PUT /api/v1/suppliers/with-id/:id
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
  contactName?: string
  contactEmail?: string
  contactPhone?: string
  address?: Record<string, unknown>
  isActive?: boolean
  rating?: number
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
  "http://localhost:3000/api/v1/suppliers/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Suppliers

```
PUT /api/v1/suppliers/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: SuppliersEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/suppliers/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Suppliers

```
DELETE /api/v1/suppliers/with-id/:id
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
  "http://localhost:3000/api/v1/suppliers/with-id/:id"
```

</details>

---

:::


::::
