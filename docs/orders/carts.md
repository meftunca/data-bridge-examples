---
title: Carts
---

# Carts

**Table:** `orders.carts`

**Base path:** `/carts`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `customers` | `customer_id` | `carts_customer_id_fkey` | [Customers](./customers) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `cart_items` | `cart_id` | `cart_items_cart_id_fkey` | [CartItems](./cart_items) |


## Entity Relationship Diagram

```mermaid
erDiagram
    Carts }o--|| Customers : "FK"
    Carts ||--o{ CartItems : "ref"
```

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `customer_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `customers` |
| 4 | `is_active` | `boolean` | `bool` | `boolean` | NO | `true` | - | - |
| 5 | `expires_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 6 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 7 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `customer_id` | `customers` | `carts_customer_id_fkey` |


## Go Generated Code

> 📂 Source: [📄 `Carts.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Carts.go) · [📄 `Carts.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Carts.go) · [📄 `Carts.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/controllers/Carts.go)

### Structs

::::tabs

:::tab Form

#### CartsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Carts.go#:~:text=type%20CartsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `CustomerId` | `uuid.UUID` | `customerId` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `ExpiresAt` | `*time.Time` | `expiresAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Model

#### Carts [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Carts.go#:~:text=type%20Carts%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `CustomerId` | `uuid.UUID` | `customerId` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `ExpiresAt` | `*time.Time` | `expiresAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab Edit

#### CartsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Carts.go#:~:text=type%20CartsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `CustomerId` | `*uuid.UUID` | `customerId` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `ExpiresAt` | `*time.Time` | `expiresAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Filter

#### CartsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Carts.go#:~:text=type%20CartsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `CustomerId` | `*uuid.UUID` | `customerId` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `ExpiresAt` | `*time.Time` | `expiresAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

:::tab Page

#### CartsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Carts.go#:~:text=type%20CartsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `CustomerId` | `uuid.UUID` | `customerId` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `ExpiresAt` | `*time.Time` | `expiresAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

:::tab BatchUpdate

#### CartsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Carts.go#:~:text=type%20CartsBatchUpdate%20struct)

```go
type CartsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Carts.go#:~:text=)%20CreateCarts() | `(CartsService) CreateCarts(data CartsForm) (CartsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Carts.go#:~:text=)%20CreateCartsMultiple() | `(CartsService) CreateCartsMultiple(data []CartsForm) ([]CartsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Carts.go#:~:text=)%20UpdateCarts() | `(CartsService) UpdateCarts(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Carts.go#:~:text=)%20UpdateCartsMultiple() | `(CartsService) UpdateCartsMultiple(data []CartsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Carts.go#:~:text=)%20DeleteCarts() | `(CartsService) DeleteCarts(id uuid.UUID) error` |

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/carts/` | Search with query params |
| `GET` | `/carts/pagination` | Paginated listing |
| `POST` | `/carts/` | Create single record |
| `POST` | `/carts/bulk/` | Create multiple records |
| `PUT` | `/carts/bulk/` | Batch update |
| `GET` | `/carts/with-id/:id` | Get by ID |
| `PUT` | `/carts/with-id/:id` | Update by ID |
| `DELETE` | `/carts/with-id/:id` | Delete by ID |

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
export interface Carts {
  id: string;
  name: string;
  customerId: string;
  isActive: boolean;
  expiresAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CartsForm {
  name: string;
  customerId: string;
  isActive: boolean;
  expiresAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CartsEdit {
  id: string;
  name: string;
  customerId: string;
  isActive: boolean;
  expiresAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CartsPage {
  data: Carts[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type CartsPathQuery = {
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

const CartsKeys = {
  all: ["carts"] as const,
  lists: () => [...CartsKeys.all, "list"] as const,
  detail: (id: any) => [...CartsKeys.all, "detail", id] as const,
} as const;

export function useCartsList(query?: CartsPathQuery) {
  return useQuery({
    queryKey: [...CartsKeys.lists(), query],
    queryFn: () => fetch(`/carts/pagination`, { method: "GET" }).then(r => r.json()) as Promise<CartsPage>,
  });
}

export function useCartsDetail(id: any) {
  return useQuery({
    queryKey: CartsKeys.detail(id),
    queryFn: () => fetch(`/carts/with-id/:id`).then(r => r.json()) as Promise<Carts>,
  });
}

export function useCreateCarts() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: CartsForm) =>
      fetch("/carts/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CartsKeys.lists() }),
  });
}

export function useUpdateCarts() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: CartsEdit }) =>
      fetch(`/carts/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CartsKeys.all }),
  });
}

export function useDeleteCarts() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/carts/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CartsKeys.all }),
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

export const CartsFormSchema = z.object({
  name: z.string(),
  customerId: z.string().uuid(),
  isActive: z.boolean(),
  expiresAt: z.string().datetime().optional(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type CartsFormInput = z.infer<typeof CartsFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './carts.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search Carts

```
GET /api/v1/carts/
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
| `preloads` | `string` | No | Available: CartItemsList, CartItemsList.CartIdDetail, CartItemsList.CartIdDetail.CartItemsList, CartItemsList.CartIdDetail.CustomerIdDetail, CustomerIdDetail, CustomerIdDetail.OrdersList, CustomerIdDetail.OrdersList.OrderItemsList, CustomerIdDetail.OrdersList.PaymentsList, CustomerIdDetail.OrdersList.RefundsList, CustomerIdDetail.OrdersList.OrderStatusHistoryList, CustomerIdDetail.OrdersList.CustomerIdDetail, CustomerIdDetail.OrdersList.CouponIdDetail, CustomerIdDetail.CartsList, CustomerIdDetail.CartsList.CartItemsList, CustomerIdDetail.CartsList.CustomerIdDetail |
| `joins` | `string` | No | Available: Customers, Customers.Users, Customers.Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `customerId` | `string (uuid)` | No | Filter by customer_id |
| `isActive` | `boolean` | No | Filter by is_active |
| `expiresAt` | `string (date-time)` | No | Filter by expires_at |

**Response:** `Carts[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/carts/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Carts (POST)

```
POST /api/v1/carts/search
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

**Response:** `Carts[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/carts/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate Carts

```
GET /api/v1/carts/pagination
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
| `preloads` | `string` | No | Available: CartItemsList, CartItemsList.CartIdDetail, CartItemsList.CartIdDetail.CartItemsList, CartItemsList.CartIdDetail.CustomerIdDetail, CustomerIdDetail, CustomerIdDetail.OrdersList, CustomerIdDetail.OrdersList.OrderItemsList, CustomerIdDetail.OrdersList.PaymentsList, CustomerIdDetail.OrdersList.RefundsList, CustomerIdDetail.OrdersList.OrderStatusHistoryList, CustomerIdDetail.OrdersList.CustomerIdDetail, CustomerIdDetail.OrdersList.CouponIdDetail, CustomerIdDetail.CartsList, CustomerIdDetail.CartsList.CartItemsList, CustomerIdDetail.CartsList.CustomerIdDetail |
| `joins` | `string` | No | Available: Customers, Customers.Users, Customers.Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `customerId` | `string (uuid)` | No | Filter by customer_id |
| `isActive` | `boolean` | No | Filter by is_active |
| `expiresAt` | `string (date-time)` | No | Filter by expires_at |

**Response:** `PaginationResponse<Carts>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/carts/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Carts (POST)

```
POST /api/v1/carts/pagination
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

**Response:** `PaginationResponse<Carts>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/carts/pagination"
```

</details>

---

:::tab Create

#### <Badge type="tip" text="POST" /> Create Carts

```
POST /api/v1/carts/
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
  customerId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  isActive?: boolean  // e.g. true
  expiresAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Carts`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/carts/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Carts

```
POST /api/v1/carts/bulk/
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
  customerId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  isActive?: boolean  // e.g. true
  expiresAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Carts[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/carts/bulk/"
```

</details>

---

:::tab Find & Update

#### <Badge type="info" text="GET" /> Find Carts by ID

```
GET /api/v1/carts/with-id/:id
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

**Response:** `Carts`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/carts/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Carts

```
PUT /api/v1/carts/with-id/:id
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
  customerId?: string
  isActive?: boolean
  expiresAt?: string
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
  "http://localhost:3000/api/v1/carts/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Carts

```
PUT /api/v1/carts/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: CartsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/carts/bulk/"
```

</details>

---

:::tab Delete

#### <Badge type="danger" text="DELETE" /> Delete Carts

```
DELETE /api/v1/carts/with-id/:id
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
  "http://localhost:3000/api/v1/carts/with-id/:id"
```

</details>

---

::::


::::
