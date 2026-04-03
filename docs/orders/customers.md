---
title: Customers
---

# Customers

**Table:** `orders.customers`

**Base path:** `/customers`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `users` | `user_id` | `customers_user_id_fkey` | [Users](./users) |
| `organizations` | `organization_id` | `customers_organization_id_fkey` | [Organizations](./organizations) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `carts` | `customer_id` | `carts_customer_id_fkey` | [Carts](./carts) |
| `orders` | `customer_id` | `orders_customer_id_fkey` | [Orders](./orders) |


## Entity Relationship Diagram

```mermaid
erDiagram
    Customers }o--|| Users : "FK"
    Customers }o--|| Organizations : "FK"
    Customers ||--o{ Carts : "ref"
    Customers ||--o{ Orders : "ref"
```

::::tabs

=== FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `user_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `users` |
| 4 | `organization_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `organizations` |
| 5 | `billing_address` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 6 | `shipping_address` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 7 | `tax_id` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 8 | `loyalty_points` | `integer` | `int` | `number` | NO | `0` | - | - |
| 9 | `tier` | `text` | `string` | `string` | NO | `'standard'::text` | - | - |
| 10 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 11 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `user_id` | `users` | `customers_user_id_fkey` |
| `organization_id` | `organizations` | `customers_organization_id_fkey` |


## Go Generated Code

> 📂 Source: [📄 `Customers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Customers.go) · [📄 `Customers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Customers.go) · [📄 `Customers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//orders/controllers/Customers.go)

### Structs

:::tabs

== Form

#### CustomersForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Customers.go#:~:text=type%20CustomersForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `BillingAddress` | `json.RawMessage` | `billingAddress` | NO |
| `ShippingAddress` | `json.RawMessage` | `shippingAddress` | NO |
| `TaxId` | `string` | `taxId` | NO |
| `LoyaltyPoints` | `int` | `loyaltyPoints` | NO |
| `Tier` | `string` | `tier` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### Customers [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Customers.go#:~:text=type%20Customers%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `BillingAddress` | `json.RawMessage` | `billingAddress` | NO |
| `ShippingAddress` | `json.RawMessage` | `shippingAddress` | NO |
| `TaxId` | `string` | `taxId` | NO |
| `LoyaltyPoints` | `int` | `loyaltyPoints` | NO |
| `Tier` | `string` | `tier` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### CustomersEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Customers.go#:~:text=type%20CustomersEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `BillingAddress` | `*json.RawMessage` | `billingAddress` | YES |
| `ShippingAddress` | `*json.RawMessage` | `shippingAddress` | YES |
| `TaxId` | `*string` | `taxId` | YES |
| `LoyaltyPoints` | `*int` | `loyaltyPoints` | YES |
| `Tier` | `*string` | `tier` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### CustomersFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Customers.go#:~:text=type%20CustomersFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `BillingAddress` | `*json.RawMessage` | `billingAddress` | YES |
| `ShippingAddress` | `*json.RawMessage` | `shippingAddress` | YES |
| `TaxId` | `*string` | `taxId` | YES |
| `LoyaltyPoints` | `*int` | `loyaltyPoints` | YES |
| `Tier` | `*string` | `tier` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### CustomersPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Customers.go#:~:text=type%20CustomersPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `BillingAddress` | `json.RawMessage` | `billingAddress` | NO |
| `ShippingAddress` | `json.RawMessage` | `shippingAddress` | NO |
| `TaxId` | `string` | `taxId` | NO |
| `LoyaltyPoints` | `int` | `loyaltyPoints` | NO |
| `Tier` | `string` | `tier` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### CustomersBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//orders/structures/Customers.go#:~:text=type%20CustomersBatchUpdate%20struct)

```go
type CustomersBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Customers.go#:~:text=%29%20CreateCustomers%28%29) | `(CustomersService) CreateCustomers(data CustomersForm) (CustomersForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Customers.go#:~:text=%29%20CreateCustomersMultiple%28%29) | `(CustomersService) CreateCustomersMultiple(data []CustomersForm) ([]CustomersForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Customers.go#:~:text=%29%20UpdateCustomers%28%29) | `(CustomersService) UpdateCustomers(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Customers.go#:~:text=%29%20UpdateCustomersMultiple%28%29) | `(CustomersService) UpdateCustomersMultiple(data []CustomersBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//orders/services/Customers.go#:~:text=%29%20DeleteCustomers%28%29) | `(CustomersService) DeleteCustomers(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/customers/` | Search with query params |
| `GET` | `/customers/pagination` | Paginated listing |
| `POST` | `/customers/` | Create single record |
| `POST` | `/customers/bulk/` | Create multiple records |
| `PUT` | `/customers/bulk/` | Batch update |
| `GET` | `/customers/with-id/:id` | Get by ID |
| `PUT` | `/customers/with-id/:id` | Update by ID |
| `DELETE` | `/customers/with-id/:id` | Delete by ID |

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
| `customer_total_spent` | `p_customer_id uuid` | `numeric` | `/rpc/customer_total_spent` |
| `orders_by_status` | `p_status text` | `integer` | `/rpc/orders_by_status` |
| `total_revenue` | - | `numeric` | `/rpc/total_revenue` |


=== Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export interface Customers {
  id: string;
  name: string;
  userId: string;
  organizationId?: string;
  billingAddress: Record<string, unknown>;
  shippingAddress: Record<string, unknown>;
  taxId: string;
  loyaltyPoints: number;
  tier: string;
  createdAt: string;
  updatedAt: string;
}

export interface CustomersForm {
  name: string;
  userId: string;
  organizationId?: string;
  billingAddress: Record<string, unknown>;
  shippingAddress: Record<string, unknown>;
  taxId: string;
  loyaltyPoints: number;
  tier: string;
  createdAt: string;
  updatedAt: string;
}

export interface CustomersEdit {
  id: string;
  name: string;
  userId: string;
  organizationId?: string;
  billingAddress: Record<string, unknown>;
  shippingAddress: Record<string, unknown>;
  taxId: string;
  loyaltyPoints: number;
  tier: string;
  createdAt: string;
  updatedAt: string;
}

export interface CustomersPage {
  data: Customers[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type CustomersPathQuery = {
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

const CustomersKeys = {
  all: ["customers"] as const,
  lists: () => [...CustomersKeys.all, "list"] as const,
  detail: (id: any) => [...CustomersKeys.all, "detail", id] as const,
} as const;

export function useCustomersList(query?: CustomersPathQuery) {
  return useQuery({
    queryKey: [...CustomersKeys.lists(), query],
    queryFn: () => fetch(`/customers/pagination`, { method: "GET" }).then(r => r.json()) as Promise<CustomersPage>,
  });
}

export function useCustomersDetail(id: any) {
  return useQuery({
    queryKey: CustomersKeys.detail(id),
    queryFn: () => fetch(`/customers/with-id/:id`).then(r => r.json()) as Promise<Customers>,
  });
}

export function useCreateCustomers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: CustomersForm) =>
      fetch("/customers/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CustomersKeys.lists() }),
  });
}

export function useUpdateCustomers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: CustomersEdit }) =>
      fetch(`/customers/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CustomersKeys.all }),
  });
}

export function useDeleteCustomers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/customers/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: CustomersKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const CustomersFormSchema = z.object({
  name: z.string(),
  userId: z.string().uuid(),
  organizationId: z.string().uuid().optional(),
  billingAddress: z.record(z.unknown()),
  shippingAddress: z.record(z.unknown()),
  taxId: z.string(),
  loyaltyPoints: z.number().int(),
  tier: z.string(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type CustomersFormInput = z.infer<typeof CustomersFormSchema>;

```

:::


=== API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './customers.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Customers

```
GET /api/v1/customers/
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
| `preloads` | `string` | No | Available: OrdersList, OrdersList.OrderItemsList, OrdersList.OrderItemsList.OrderIdDetail, OrdersList.PaymentsList, OrdersList.PaymentsList.RefundsList, OrdersList.PaymentsList.OrderIdDetail, OrdersList.RefundsList, OrdersList.RefundsList.OrderIdDetail, OrdersList.RefundsList.PaymentIdDetail, OrdersList.OrderStatusHistoryList, OrdersList.OrderStatusHistoryList.OrderIdDetail, OrdersList.CustomerIdDetail, OrdersList.CustomerIdDetail.OrdersList, OrdersList.CustomerIdDetail.CartsList, OrdersList.CouponIdDetail, OrdersList.CouponIdDetail.OrdersList, CartsList, CartsList.CartItemsList, CartsList.CartItemsList.CartIdDetail, CartsList.CustomerIdDetail, CartsList.CustomerIdDetail.OrdersList, CartsList.CustomerIdDetail.CartsList |
| `joins` | `string` | No | Available: Users, Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `userId` | `string (uuid)` | No | Filter by user_id |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `billingAddress` | `string` | No | Filter by billing_address |
| `shippingAddress` | `string` | No | Filter by shipping_address |
| `taxId` | `string` | No | Filter by tax_id |
| `loyaltyPoints` | `integer` | No | Filter by loyalty_points |
| `tier` | `string` | No | Filter by tier |

**Response:** `Customers[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/customers/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Customers (POST)

```
POST /api/v1/customers/search
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

**Response:** `Customers[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/customers/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Customers

```
GET /api/v1/customers/pagination
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
| `preloads` | `string` | No | Available: OrdersList, OrdersList.OrderItemsList, OrdersList.OrderItemsList.OrderIdDetail, OrdersList.PaymentsList, OrdersList.PaymentsList.RefundsList, OrdersList.PaymentsList.OrderIdDetail, OrdersList.RefundsList, OrdersList.RefundsList.OrderIdDetail, OrdersList.RefundsList.PaymentIdDetail, OrdersList.OrderStatusHistoryList, OrdersList.OrderStatusHistoryList.OrderIdDetail, OrdersList.CustomerIdDetail, OrdersList.CustomerIdDetail.OrdersList, OrdersList.CustomerIdDetail.CartsList, OrdersList.CouponIdDetail, OrdersList.CouponIdDetail.OrdersList, CartsList, CartsList.CartItemsList, CartsList.CartItemsList.CartIdDetail, CartsList.CustomerIdDetail, CartsList.CustomerIdDetail.OrdersList, CartsList.CustomerIdDetail.CartsList |
| `joins` | `string` | No | Available: Users, Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `userId` | `string (uuid)` | No | Filter by user_id |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `billingAddress` | `string` | No | Filter by billing_address |
| `shippingAddress` | `string` | No | Filter by shipping_address |
| `taxId` | `string` | No | Filter by tax_id |
| `loyaltyPoints` | `integer` | No | Filter by loyalty_points |
| `tier` | `string` | No | Filter by tier |

**Response:** `PaginationResponse<Customers>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/customers/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Customers (POST)

```
POST /api/v1/customers/pagination
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

**Response:** `PaginationResponse<Customers>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/customers/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Customers

```
POST /api/v1/customers/
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
  userId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  billingAddress?: Record<string, unknown>  // e.g. map[]
  shippingAddress?: Record<string, unknown>  // e.g. map[]
  taxId?: string  // e.g. example_tax_id
  loyaltyPoints?: number  // e.g. 1
  tier?: string  // e.g. example_tier
}
```

**Response:** `Customers`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/customers/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Customers

```
POST /api/v1/customers/bulk/
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
  userId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  billingAddress?: Record<string, unknown>  // e.g. map[]
  shippingAddress?: Record<string, unknown>  // e.g. map[]
  taxId?: string  // e.g. example_tax_id
  loyaltyPoints?: number  // e.g. 1
  tier?: string  // e.g. example_tier
}
```

**Response:** `Customers[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/customers/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Customers by ID

```
GET /api/v1/customers/with-id/:id
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

**Response:** `Customers`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/customers/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Customers

```
PUT /api/v1/customers/with-id/:id
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
  userId?: string
  organizationId?: string
  billingAddress?: Record<string, unknown>
  shippingAddress?: Record<string, unknown>
  taxId?: string
  loyaltyPoints?: number
  tier?: string
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
  "http://localhost:3000/api/v1/customers/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Customers

```
PUT /api/v1/customers/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: CustomersEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/customers/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Customers

```
DELETE /api/v1/customers/with-id/:id
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
  "http://localhost:3000/api/v1/customers/with-id/:id"
```

</details>

---

:::


::::
