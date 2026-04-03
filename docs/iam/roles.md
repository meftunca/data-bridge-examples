---
title: Roles
---

# Roles

**Table:** `iam.roles`

**Base path:** `/roles`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `organizations` | `organization_id` | `roles_organization_id_fkey` | [Organizations](./organizations) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `invitations` | `role_id` | `invitations_role_id_fkey` | [Invitations](./invitations) |
| `role_permissions` | `role_id` | `role_permissions_role_id_fkey` | [RolePermissions](./role_permissions) |
| `user_roles` | `role_id` | `user_roles_role_id_fkey` | [UserRoles](./user_roles) |


## Entity Relationship Diagram

```mermaid
erDiagram
    Roles }o--|| Organizations : "FK"
    Roles ||--o{ Invitations : "ref"
    Roles ||--o{ RolePermissions : "ref"
    Roles ||--o{ UserRoles : "ref"
```

::::tabs

=== FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `slug` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 4 | `description` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `organization_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `organizations` |
| 6 | `is_system` | `boolean` | `bool` | `boolean` | NO | `false` | - | - |
| 7 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 8 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `organization_id` | `organizations` | `roles_organization_id_fkey` |

## Unique Keys

- `slug` (`text`)


## Go Generated Code

> 📂 Source: [📄 `Roles.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Roles.go) · [📄 `Roles.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Roles.go) · [📄 `Roles.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/controllers/Roles.go)

### Structs

:::tabs

== Form

#### RolesForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Roles.go#:~:text=type%20RolesForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `IsSystem` | `bool` | `isSystem` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### Roles [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Roles.go#:~:text=type%20Roles%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `IsSystem` | `bool` | `isSystem` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### RolesEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Roles.go#:~:text=type%20RolesEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `Description` | `*string` | `description` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `IsSystem` | `*bool` | `isSystem` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### RolesFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Roles.go#:~:text=type%20RolesFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Slug` | `*string` | `slug` | YES |
| `Description` | `*string` | `description` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `IsSystem` | `*bool` | `isSystem` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### RolesPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Roles.go#:~:text=type%20RolesPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Slug` | `string` | `slug` | NO |
| `Description` | `string` | `description` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `IsSystem` | `bool` | `isSystem` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### RolesBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Roles.go#:~:text=type%20RolesBatchUpdate%20struct)

```go
type RolesBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Roles.go#:~:text=%29%20CreateRoles%28%29) | `(RolesService) CreateRoles(data RolesForm) (RolesForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Roles.go#:~:text=%29%20CreateRolesMultiple%28%29) | `(RolesService) CreateRolesMultiple(data []RolesForm) ([]RolesForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Roles.go#:~:text=%29%20UpdateRoles%28%29) | `(RolesService) UpdateRoles(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Roles.go#:~:text=%29%20UpdateRolesMultiple%28%29) | `(RolesService) UpdateRolesMultiple(data []RolesBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Roles.go#:~:text=%29%20DeleteRoles%28%29) | `(RolesService) DeleteRoles(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/roles/` | Search with query params |
| `GET` | `/roles/pagination` | Paginated listing |
| `POST` | `/roles/` | Create single record |
| `POST` | `/roles/bulk/` | Create multiple records |
| `PUT` | `/roles/bulk/` | Batch update |
| `GET` | `/roles/with-id/:id` | Get by ID |
| `PUT` | `/roles/with-id/:id` | Update by ID |
| `DELETE` | `/roles/with-id/:id` | Delete by ID |

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


=== Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export interface Roles {
  id: string;
  name: string;
  slug: string;
  description: string;
  organizationId?: string;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface RolesForm {
  name: string;
  slug: string;
  description: string;
  organizationId?: string;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface RolesEdit {
  id: string;
  name: string;
  slug: string;
  description: string;
  organizationId?: string;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface RolesPage {
  data: Roles[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type RolesPathQuery = {
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

const RolesKeys = {
  all: ["roles"] as const,
  lists: () => [...RolesKeys.all, "list"] as const,
  detail: (id: any) => [...RolesKeys.all, "detail", id] as const,
} as const;

export function useRolesList(query?: RolesPathQuery) {
  return useQuery({
    queryKey: [...RolesKeys.lists(), query],
    queryFn: () => fetch(`/roles/pagination`, { method: "GET" }).then(r => r.json()) as Promise<RolesPage>,
  });
}

export function useRolesDetail(id: any) {
  return useQuery({
    queryKey: RolesKeys.detail(id),
    queryFn: () => fetch(`/roles/with-id/:id`).then(r => r.json()) as Promise<Roles>,
  });
}

export function useCreateRoles() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: RolesForm) =>
      fetch("/roles/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: RolesKeys.lists() }),
  });
}

export function useUpdateRoles() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: RolesEdit }) =>
      fetch(`/roles/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: RolesKeys.all }),
  });
}

export function useDeleteRoles() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/roles/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: RolesKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const RolesFormSchema = z.object({
  name: z.string(),
  slug: z.string(),
  description: z.string(),
  organizationId: z.string().uuid().optional(),
  isSystem: z.boolean(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type RolesFormInput = z.infer<typeof RolesFormSchema>;

```

:::


=== API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './roles.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Roles

```
GET /api/v1/roles/
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
| `preloads` | `string` | No | Available: RolePermissionsList, RolePermissionsList.RoleIdDetail, RolePermissionsList.RoleIdDetail.RolePermissionsList, RolePermissionsList.RoleIdDetail.UserRolesList, RolePermissionsList.RoleIdDetail.InvitationsList, RolePermissionsList.RoleIdDetail.OrganizationIdDetail, RolePermissionsList.PermissionIdDetail, RolePermissionsList.PermissionIdDetail.RolePermissionsList, UserRolesList, UserRolesList.UserIdDetail, UserRolesList.UserIdDetail.UserRolesList, UserRolesList.UserIdDetail.TeamsList, UserRolesList.UserIdDetail.TeamMembersList, UserRolesList.UserIdDetail.ApiKeysList, UserRolesList.UserIdDetail.SessionsList, UserRolesList.UserIdDetail.InvitationsList, UserRolesList.UserIdDetail.OrganizationIdDetail, UserRolesList.RoleIdDetail, UserRolesList.RoleIdDetail.RolePermissionsList, UserRolesList.RoleIdDetail.UserRolesList, UserRolesList.RoleIdDetail.InvitationsList, UserRolesList.RoleIdDetail.OrganizationIdDetail, UserRolesList.GrantedByDetail, UserRolesList.GrantedByDetail.UserRolesList, UserRolesList.GrantedByDetail.TeamsList, UserRolesList.GrantedByDetail.TeamMembersList, UserRolesList.GrantedByDetail.ApiKeysList, UserRolesList.GrantedByDetail.SessionsList, UserRolesList.GrantedByDetail.InvitationsList, UserRolesList.GrantedByDetail.OrganizationIdDetail, InvitationsList, InvitationsList.OrganizationIdDetail, InvitationsList.OrganizationIdDetail.OrganizationsList, InvitationsList.OrganizationIdDetail.UsersList, InvitationsList.OrganizationIdDetail.RolesList, InvitationsList.OrganizationIdDetail.TeamsList, InvitationsList.OrganizationIdDetail.ApiKeysList, InvitationsList.OrganizationIdDetail.InvitationsList, InvitationsList.OrganizationIdDetail.ParentIdDetail, InvitationsList.InvitedByDetail, InvitationsList.InvitedByDetail.UserRolesList, InvitationsList.InvitedByDetail.TeamsList, InvitationsList.InvitedByDetail.TeamMembersList, InvitationsList.InvitedByDetail.ApiKeysList, InvitationsList.InvitedByDetail.SessionsList, InvitationsList.InvitedByDetail.InvitationsList, InvitationsList.InvitedByDetail.OrganizationIdDetail, InvitationsList.RoleIdDetail, InvitationsList.RoleIdDetail.RolePermissionsList, InvitationsList.RoleIdDetail.UserRolesList, InvitationsList.RoleIdDetail.InvitationsList, InvitationsList.RoleIdDetail.OrganizationIdDetail, OrganizationIdDetail, OrganizationIdDetail.OrganizationsList, OrganizationIdDetail.UsersList, OrganizationIdDetail.UsersList.UserRolesList, OrganizationIdDetail.UsersList.TeamsList, OrganizationIdDetail.UsersList.TeamMembersList, OrganizationIdDetail.UsersList.ApiKeysList, OrganizationIdDetail.UsersList.SessionsList, OrganizationIdDetail.UsersList.InvitationsList, OrganizationIdDetail.UsersList.OrganizationIdDetail, OrganizationIdDetail.RolesList, OrganizationIdDetail.RolesList.RolePermissionsList, OrganizationIdDetail.RolesList.UserRolesList, OrganizationIdDetail.RolesList.InvitationsList, OrganizationIdDetail.RolesList.OrganizationIdDetail, OrganizationIdDetail.TeamsList, OrganizationIdDetail.TeamsList.TeamMembersList, OrganizationIdDetail.TeamsList.OrganizationIdDetail, OrganizationIdDetail.TeamsList.LeadIdDetail, OrganizationIdDetail.ApiKeysList, OrganizationIdDetail.ApiKeysList.UserIdDetail, OrganizationIdDetail.ApiKeysList.OrganizationIdDetail, OrganizationIdDetail.InvitationsList, OrganizationIdDetail.InvitationsList.OrganizationIdDetail, OrganizationIdDetail.InvitationsList.InvitedByDetail, OrganizationIdDetail.InvitationsList.RoleIdDetail, OrganizationIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Organizations, Organizations.Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |
| `description` | `string` | No | Filter by description |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `isSystem` | `boolean` | No | Filter by is_system |

**Response:** `Roles[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/roles/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Roles (POST)

```
POST /api/v1/roles/search
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

**Response:** `Roles[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/roles/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Roles

```
GET /api/v1/roles/pagination
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
| `preloads` | `string` | No | Available: RolePermissionsList, RolePermissionsList.RoleIdDetail, RolePermissionsList.RoleIdDetail.RolePermissionsList, RolePermissionsList.RoleIdDetail.UserRolesList, RolePermissionsList.RoleIdDetail.InvitationsList, RolePermissionsList.RoleIdDetail.OrganizationIdDetail, RolePermissionsList.PermissionIdDetail, RolePermissionsList.PermissionIdDetail.RolePermissionsList, UserRolesList, UserRolesList.UserIdDetail, UserRolesList.UserIdDetail.UserRolesList, UserRolesList.UserIdDetail.TeamsList, UserRolesList.UserIdDetail.TeamMembersList, UserRolesList.UserIdDetail.ApiKeysList, UserRolesList.UserIdDetail.SessionsList, UserRolesList.UserIdDetail.InvitationsList, UserRolesList.UserIdDetail.OrganizationIdDetail, UserRolesList.RoleIdDetail, UserRolesList.RoleIdDetail.RolePermissionsList, UserRolesList.RoleIdDetail.UserRolesList, UserRolesList.RoleIdDetail.InvitationsList, UserRolesList.RoleIdDetail.OrganizationIdDetail, UserRolesList.GrantedByDetail, UserRolesList.GrantedByDetail.UserRolesList, UserRolesList.GrantedByDetail.TeamsList, UserRolesList.GrantedByDetail.TeamMembersList, UserRolesList.GrantedByDetail.ApiKeysList, UserRolesList.GrantedByDetail.SessionsList, UserRolesList.GrantedByDetail.InvitationsList, UserRolesList.GrantedByDetail.OrganizationIdDetail, InvitationsList, InvitationsList.OrganizationIdDetail, InvitationsList.OrganizationIdDetail.OrganizationsList, InvitationsList.OrganizationIdDetail.UsersList, InvitationsList.OrganizationIdDetail.RolesList, InvitationsList.OrganizationIdDetail.TeamsList, InvitationsList.OrganizationIdDetail.ApiKeysList, InvitationsList.OrganizationIdDetail.InvitationsList, InvitationsList.OrganizationIdDetail.ParentIdDetail, InvitationsList.InvitedByDetail, InvitationsList.InvitedByDetail.UserRolesList, InvitationsList.InvitedByDetail.TeamsList, InvitationsList.InvitedByDetail.TeamMembersList, InvitationsList.InvitedByDetail.ApiKeysList, InvitationsList.InvitedByDetail.SessionsList, InvitationsList.InvitedByDetail.InvitationsList, InvitationsList.InvitedByDetail.OrganizationIdDetail, InvitationsList.RoleIdDetail, InvitationsList.RoleIdDetail.RolePermissionsList, InvitationsList.RoleIdDetail.UserRolesList, InvitationsList.RoleIdDetail.InvitationsList, InvitationsList.RoleIdDetail.OrganizationIdDetail, OrganizationIdDetail, OrganizationIdDetail.OrganizationsList, OrganizationIdDetail.UsersList, OrganizationIdDetail.UsersList.UserRolesList, OrganizationIdDetail.UsersList.TeamsList, OrganizationIdDetail.UsersList.TeamMembersList, OrganizationIdDetail.UsersList.ApiKeysList, OrganizationIdDetail.UsersList.SessionsList, OrganizationIdDetail.UsersList.InvitationsList, OrganizationIdDetail.UsersList.OrganizationIdDetail, OrganizationIdDetail.RolesList, OrganizationIdDetail.RolesList.RolePermissionsList, OrganizationIdDetail.RolesList.UserRolesList, OrganizationIdDetail.RolesList.InvitationsList, OrganizationIdDetail.RolesList.OrganizationIdDetail, OrganizationIdDetail.TeamsList, OrganizationIdDetail.TeamsList.TeamMembersList, OrganizationIdDetail.TeamsList.OrganizationIdDetail, OrganizationIdDetail.TeamsList.LeadIdDetail, OrganizationIdDetail.ApiKeysList, OrganizationIdDetail.ApiKeysList.UserIdDetail, OrganizationIdDetail.ApiKeysList.OrganizationIdDetail, OrganizationIdDetail.InvitationsList, OrganizationIdDetail.InvitationsList.OrganizationIdDetail, OrganizationIdDetail.InvitationsList.InvitedByDetail, OrganizationIdDetail.InvitationsList.RoleIdDetail, OrganizationIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Organizations, Organizations.Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `slug` | `string` | No | Filter by slug |
| `description` | `string` | No | Filter by description |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `isSystem` | `boolean` | No | Filter by is_system |

**Response:** `PaginationResponse<Roles>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/roles/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Roles (POST)

```
POST /api/v1/roles/pagination
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

**Response:** `PaginationResponse<Roles>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/roles/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Roles

```
POST /api/v1/roles/
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
  slug: string  // e.g. example_slug
  description?: string  // e.g. example_description
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  isSystem?: boolean  // e.g. true
}
```

**Response:** `Roles`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/roles/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Roles

```
POST /api/v1/roles/bulk/
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
  slug: string  // e.g. example_slug
  description?: string  // e.g. example_description
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  isSystem?: boolean  // e.g. true
}
```

**Response:** `Roles[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/roles/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Roles by ID

```
GET /api/v1/roles/with-id/:id
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

**Response:** `Roles`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/roles/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Roles

```
PUT /api/v1/roles/with-id/:id
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
  slug?: string
  description?: string
  organizationId?: string
  isSystem?: boolean
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
  "http://localhost:3000/api/v1/roles/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Roles

```
PUT /api/v1/roles/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: RolesEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/roles/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Roles

```
DELETE /api/v1/roles/with-id/:id
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
  "http://localhost:3000/api/v1/roles/with-id/:id"
```

</details>

---

:::


::::
