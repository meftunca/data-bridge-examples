---
title: RolePermissions
---

# RolePermissions

**Table:** `iam.role_permissions`

**Base path:** `/role-permissions`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `roles` | `role_id` | `role_permissions_role_id_fkey` | [Roles](./roles) |
| `permissions` | `permission_id` | `role_permissions_permission_id_fkey` | [Permissions](./permissions) |


## Entity Relationship Diagram

erDiagram
    RolePermissions }o--|| Roles : "FK"
    RolePermissions }o--|| Permissions : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `role_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `UQ` `FK` | → References `roles` |
| 4 | `permission_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `UQ` `FK` | → References `permissions` |
| 5 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `role_id` | `roles` | `role_permissions_role_id_fkey` |
| `permission_id` | `permissions` | `role_permissions_permission_id_fkey` |

## Unique Keys

- `role_id` (`uuid`)
- `permission_id` (`uuid`)


## Go Generated Code

> 📂 Source: [📄 `RolePermissions.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/RolePermissions.go) · [📄 `RolePermissions.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/RolePermissions.go) · [📄 `RolePermissions.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/controllers/RolePermissions.go)

### Structs

:::tabs

== Form

#### RolePermissionsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/RolePermissions.go#:~:text=type%20RolePermissionsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `RoleId` | `uuid.UUID` | `roleId` | NO |
| `PermissionId` | `uuid.UUID` | `permissionId` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### RolePermissions [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/RolePermissions.go#:~:text=type%20RolePermissions%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `RoleId` | `uuid.UUID` | `roleId` | NO |
| `PermissionId` | `uuid.UUID` | `permissionId` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### RolePermissionsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/RolePermissions.go#:~:text=type%20RolePermissionsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `RoleId` | `*uuid.UUID` | `roleId` | YES |
| `PermissionId` | `*uuid.UUID` | `permissionId` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### RolePermissionsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/RolePermissions.go#:~:text=type%20RolePermissionsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `RoleId` | `*uuid.UUID` | `roleId` | YES |
| `PermissionId` | `*uuid.UUID` | `permissionId` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### RolePermissionsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/RolePermissions.go#:~:text=type%20RolePermissionsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `RoleId` | `uuid.UUID` | `roleId` | NO |
| `PermissionId` | `uuid.UUID` | `permissionId` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### RolePermissionsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/RolePermissions.go#:~:text=type%20RolePermissionsBatchUpdate%20struct)

```go
type RolePermissionsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/RolePermissions.go#:~:text=)%20CreateRolePermissions() | `(RolePermissionsService) CreateRolePermissions(data RolePermissionsForm) (RolePermissionsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/RolePermissions.go#:~:text=)%20CreateRolePermissionsMultiple() | `(RolePermissionsService) CreateRolePermissionsMultiple(data []RolePermissionsForm) ([]RolePermissionsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/RolePermissions.go#:~:text=)%20UpdateRolePermissions() | `(RolePermissionsService) UpdateRolePermissions(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/RolePermissions.go#:~:text=)%20UpdateRolePermissionsMultiple() | `(RolePermissionsService) UpdateRolePermissionsMultiple(data []RolePermissionsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/RolePermissions.go#:~:text=)%20DeleteRolePermissions() | `(RolePermissionsService) DeleteRolePermissions(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/role-permissions/` | Search with query params |
| `GET` | `/role-permissions/pagination` | Paginated listing |
| `POST` | `/role-permissions/` | Create single record |
| `POST` | `/role-permissions/bulk/` | Create multiple records |
| `PUT` | `/role-permissions/bulk/` | Batch update |
| `GET` | `/role-permissions/with-id/:id` | Get by ID |
| `PUT` | `/role-permissions/with-id/:id` | Update by ID |
| `DELETE` | `/role-permissions/with-id/:id` | Delete by ID |

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
| `count_active_users` | - | `integer` | `/rpc/count_active_users` |
| `user_permissions` | `p_user_id uuid`, `resource text`, `action text` | `record` | `/rpc/user_permissions` |
| `users_by_organization` | `p_org_id uuid` | `integer` | `/rpc/users_by_organization` |


:::tab Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export interface RolePermissions {
  id: string;
  name: string;
  roleId: string;
  permissionId: string;
  createdAt: string;
}

export interface RolePermissionsForm {
  name: string;
  roleId: string;
  permissionId: string;
  createdAt: string;
}

export interface RolePermissionsEdit {
  id: string;
  name: string;
  roleId: string;
  permissionId: string;
  createdAt: string;
}

export interface RolePermissionsPage {
  data: RolePermissions[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type RolePermissionsPathQuery = {
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

const RolePermissionsKeys = {
  all: ["role_permissions"] as const,
  lists: () => [...RolePermissionsKeys.all, "list"] as const,
  detail: (id: any) => [...RolePermissionsKeys.all, "detail", id] as const,
} as const;

export function useRolePermissionsList(query?: RolePermissionsPathQuery) {
  return useQuery({
    queryKey: [...RolePermissionsKeys.lists(), query],
    queryFn: () => fetch(`/role-permissions/pagination`, { method: "GET" }).then(r => r.json()) as Promise<RolePermissionsPage>,
  });
}

export function useRolePermissionsDetail(id: any) {
  return useQuery({
    queryKey: RolePermissionsKeys.detail(id),
    queryFn: () => fetch(`/role-permissions/with-id/:id`).then(r => r.json()) as Promise<RolePermissions>,
  });
}

export function useCreateRolePermissions() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: RolePermissionsForm) =>
      fetch("/role-permissions/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: RolePermissionsKeys.lists() }),
  });
}

export function useUpdateRolePermissions() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: RolePermissionsEdit }) =>
      fetch(`/role-permissions/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: RolePermissionsKeys.all }),
  });
}

export function useDeleteRolePermissions() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/role-permissions/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: RolePermissionsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const RolePermissionsFormSchema = z.object({
  name: z.string(),
  roleId: z.string().uuid(),
  permissionId: z.string().uuid(),
  createdAt: z.string().datetime(),
});

export type RolePermissionsFormInput = z.infer<typeof RolePermissionsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './role_permissions.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search RolePermissions

```
GET /api/v1/role-permissions/
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
| `preloads` | `string` | No | Available: RoleIdDetail, RoleIdDetail.RolePermissionsList, RoleIdDetail.RolePermissionsList.RoleIdDetail, RoleIdDetail.RolePermissionsList.PermissionIdDetail, RoleIdDetail.UserRolesList, RoleIdDetail.UserRolesList.UserIdDetail, RoleIdDetail.UserRolesList.RoleIdDetail, RoleIdDetail.UserRolesList.GrantedByDetail, RoleIdDetail.InvitationsList, RoleIdDetail.InvitationsList.OrganizationIdDetail, RoleIdDetail.InvitationsList.InvitedByDetail, RoleIdDetail.InvitationsList.RoleIdDetail, RoleIdDetail.OrganizationIdDetail, RoleIdDetail.OrganizationIdDetail.OrganizationsList, RoleIdDetail.OrganizationIdDetail.UsersList, RoleIdDetail.OrganizationIdDetail.RolesList, RoleIdDetail.OrganizationIdDetail.TeamsList, RoleIdDetail.OrganizationIdDetail.ApiKeysList, RoleIdDetail.OrganizationIdDetail.InvitationsList, RoleIdDetail.OrganizationIdDetail.ParentIdDetail, PermissionIdDetail, PermissionIdDetail.RolePermissionsList, PermissionIdDetail.RolePermissionsList.RoleIdDetail, PermissionIdDetail.RolePermissionsList.PermissionIdDetail |
| `joins` | `string` | No | Available: Roles, Roles.Organizations, Roles.Organizations.Organizations, Permissions |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `roleId` | `string (uuid)` | No | Filter by role_id |
| `permissionId` | `string (uuid)` | No | Filter by permission_id |

**Response:** `RolePermissions[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/role-permissions/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search RolePermissions (POST)

```
POST /api/v1/role-permissions/search
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

**Response:** `RolePermissions[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/role-permissions/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate RolePermissions

```
GET /api/v1/role-permissions/pagination
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
| `preloads` | `string` | No | Available: RoleIdDetail, RoleIdDetail.RolePermissionsList, RoleIdDetail.RolePermissionsList.RoleIdDetail, RoleIdDetail.RolePermissionsList.PermissionIdDetail, RoleIdDetail.UserRolesList, RoleIdDetail.UserRolesList.UserIdDetail, RoleIdDetail.UserRolesList.RoleIdDetail, RoleIdDetail.UserRolesList.GrantedByDetail, RoleIdDetail.InvitationsList, RoleIdDetail.InvitationsList.OrganizationIdDetail, RoleIdDetail.InvitationsList.InvitedByDetail, RoleIdDetail.InvitationsList.RoleIdDetail, RoleIdDetail.OrganizationIdDetail, RoleIdDetail.OrganizationIdDetail.OrganizationsList, RoleIdDetail.OrganizationIdDetail.UsersList, RoleIdDetail.OrganizationIdDetail.RolesList, RoleIdDetail.OrganizationIdDetail.TeamsList, RoleIdDetail.OrganizationIdDetail.ApiKeysList, RoleIdDetail.OrganizationIdDetail.InvitationsList, RoleIdDetail.OrganizationIdDetail.ParentIdDetail, PermissionIdDetail, PermissionIdDetail.RolePermissionsList, PermissionIdDetail.RolePermissionsList.RoleIdDetail, PermissionIdDetail.RolePermissionsList.PermissionIdDetail |
| `joins` | `string` | No | Available: Roles, Roles.Organizations, Roles.Organizations.Organizations, Permissions |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `roleId` | `string (uuid)` | No | Filter by role_id |
| `permissionId` | `string (uuid)` | No | Filter by permission_id |

**Response:** `PaginationResponse<RolePermissions>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/role-permissions/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate RolePermissions (POST)

```
POST /api/v1/role-permissions/pagination
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

**Response:** `PaginationResponse<RolePermissions>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/role-permissions/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create RolePermissions

```
POST /api/v1/role-permissions/
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
  roleId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  permissionId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
}
```

**Response:** `RolePermissions`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/role-permissions/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create RolePermissions

```
POST /api/v1/role-permissions/bulk/
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
  roleId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  permissionId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
}
```

**Response:** `RolePermissions[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/role-permissions/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find RolePermissions by ID

```
GET /api/v1/role-permissions/with-id/:id
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

**Response:** `RolePermissions`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/role-permissions/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update RolePermissions

```
PUT /api/v1/role-permissions/with-id/:id
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
  roleId?: string
  permissionId?: string
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
  "http://localhost:3000/api/v1/role-permissions/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update RolePermissions

```
PUT /api/v1/role-permissions/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: RolePermissionsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/role-permissions/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete RolePermissions

```
DELETE /api/v1/role-permissions/with-id/:id
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
  "http://localhost:3000/api/v1/role-permissions/with-id/:id"
```

</details>

---

:::


::::
