---
title: TeamMembers
---

# TeamMembers

**Table:** `iam.team_members`

**Base path:** `/team-members`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `teams` | `team_id` | `team_members_team_id_fkey` | [Teams](./teams) |
| `users` | `user_id` | `team_members_user_id_fkey` | [Users](./users) |


## Entity Relationship Diagram

erDiagram
    TeamMembers }o--|| Teams : "FK"
    TeamMembers }o--|| Users : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `team_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `UQ` `FK` | → References `teams` |
| 4 | `user_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `UQ` `FK` | → References `users` |
| 5 | `role` | `text` | `string` | `string` | NO | `'member'::text` | - | - |
| 6 | `joined_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | - |
| 7 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `team_id` | `teams` | `team_members_team_id_fkey` |
| `user_id` | `users` | `team_members_user_id_fkey` |

## Unique Keys

- `team_id` (`uuid`)
- `user_id` (`uuid`)


## Go Generated Code

> 📂 Source: [📄 `TeamMembers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/TeamMembers.go) · [📄 `TeamMembers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/TeamMembers.go) · [📄 `TeamMembers.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/controllers/TeamMembers.go)

### Structs

:::tabs

== Form

#### TeamMembersForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/TeamMembers.go#:~:text=type%20TeamMembersForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `TeamId` | `uuid.UUID` | `teamId` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `Role` | `string` | `role` | NO |
| `JoinedAt` | `time.Time` | `joinedAt` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### TeamMembers [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/TeamMembers.go#:~:text=type%20TeamMembers%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `TeamId` | `uuid.UUID` | `teamId` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `Role` | `string` | `role` | NO |
| `JoinedAt` | `time.Time` | `joinedAt` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### TeamMembersEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/TeamMembers.go#:~:text=type%20TeamMembersEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `TeamId` | `*uuid.UUID` | `teamId` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Role` | `*string` | `role` | YES |
| `JoinedAt` | `*time.Time` | `joinedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### TeamMembersFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/TeamMembers.go#:~:text=type%20TeamMembersFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `TeamId` | `*uuid.UUID` | `teamId` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Role` | `*string` | `role` | YES |
| `JoinedAt` | `*time.Time` | `joinedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### TeamMembersPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/TeamMembers.go#:~:text=type%20TeamMembersPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `TeamId` | `uuid.UUID` | `teamId` | NO |
| `UserId` | `uuid.UUID` | `userId` | NO |
| `Role` | `string` | `role` | NO |
| `JoinedAt` | `time.Time` | `joinedAt` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### TeamMembersBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/TeamMembers.go#:~:text=type%20TeamMembersBatchUpdate%20struct)

```go
type TeamMembersBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/TeamMembers.go#:~:text=)%20CreateTeamMembers() | `(TeamMembersService) CreateTeamMembers(data TeamMembersForm) (TeamMembersForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/TeamMembers.go#:~:text=)%20CreateTeamMembersMultiple() | `(TeamMembersService) CreateTeamMembersMultiple(data []TeamMembersForm) ([]TeamMembersForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/TeamMembers.go#:~:text=)%20UpdateTeamMembers() | `(TeamMembersService) UpdateTeamMembers(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/TeamMembers.go#:~:text=)%20UpdateTeamMembersMultiple() | `(TeamMembersService) UpdateTeamMembersMultiple(data []TeamMembersBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/TeamMembers.go#:~:text=)%20DeleteTeamMembers() | `(TeamMembersService) DeleteTeamMembers(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/team-members/` | Search with query params |
| `GET` | `/team-members/pagination` | Paginated listing |
| `POST` | `/team-members/` | Create single record |
| `POST` | `/team-members/bulk/` | Create multiple records |
| `PUT` | `/team-members/bulk/` | Batch update |
| `GET` | `/team-members/with-id/:id` | Get by ID |
| `PUT` | `/team-members/with-id/:id` | Update by ID |
| `DELETE` | `/team-members/with-id/:id` | Delete by ID |

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
export interface TeamMembers {
  id: string;
  name: string;
  teamId: string;
  userId: string;
  role: string;
  joinedAt: string;
  createdAt: string;
}

export interface TeamMembersForm {
  name: string;
  teamId: string;
  userId: string;
  role: string;
  joinedAt: string;
  createdAt: string;
}

export interface TeamMembersEdit {
  id: string;
  name: string;
  teamId: string;
  userId: string;
  role: string;
  joinedAt: string;
  createdAt: string;
}

export interface TeamMembersPage {
  data: TeamMembers[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type TeamMembersPathQuery = {
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

const TeamMembersKeys = {
  all: ["team_members"] as const,
  lists: () => [...TeamMembersKeys.all, "list"] as const,
  detail: (id: any) => [...TeamMembersKeys.all, "detail", id] as const,
} as const;

export function useTeamMembersList(query?: TeamMembersPathQuery) {
  return useQuery({
    queryKey: [...TeamMembersKeys.lists(), query],
    queryFn: () => fetch(`/team-members/pagination`, { method: "GET" }).then(r => r.json()) as Promise<TeamMembersPage>,
  });
}

export function useTeamMembersDetail(id: any) {
  return useQuery({
    queryKey: TeamMembersKeys.detail(id),
    queryFn: () => fetch(`/team-members/with-id/:id`).then(r => r.json()) as Promise<TeamMembers>,
  });
}

export function useCreateTeamMembers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: TeamMembersForm) =>
      fetch("/team-members/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: TeamMembersKeys.lists() }),
  });
}

export function useUpdateTeamMembers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: TeamMembersEdit }) =>
      fetch(`/team-members/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: TeamMembersKeys.all }),
  });
}

export function useDeleteTeamMembers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/team-members/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: TeamMembersKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const TeamMembersFormSchema = z.object({
  name: z.string(),
  teamId: z.string().uuid(),
  userId: z.string().uuid(),
  role: z.string(),
  joinedAt: z.string().datetime(),
  createdAt: z.string().datetime(),
});

export type TeamMembersFormInput = z.infer<typeof TeamMembersFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './team_members.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search TeamMembers

```
GET /api/v1/team-members/
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
| `preloads` | `string` | No | Available: TeamIdDetail, TeamIdDetail.TeamMembersList, TeamIdDetail.TeamMembersList.TeamIdDetail, TeamIdDetail.TeamMembersList.UserIdDetail, TeamIdDetail.OrganizationIdDetail, TeamIdDetail.OrganizationIdDetail.OrganizationsList, TeamIdDetail.OrganizationIdDetail.UsersList, TeamIdDetail.OrganizationIdDetail.RolesList, TeamIdDetail.OrganizationIdDetail.TeamsList, TeamIdDetail.OrganizationIdDetail.ApiKeysList, TeamIdDetail.OrganizationIdDetail.InvitationsList, TeamIdDetail.OrganizationIdDetail.ParentIdDetail, TeamIdDetail.LeadIdDetail, TeamIdDetail.LeadIdDetail.UserRolesList, TeamIdDetail.LeadIdDetail.TeamsList, TeamIdDetail.LeadIdDetail.TeamMembersList, TeamIdDetail.LeadIdDetail.ApiKeysList, TeamIdDetail.LeadIdDetail.SessionsList, TeamIdDetail.LeadIdDetail.InvitationsList, TeamIdDetail.LeadIdDetail.OrganizationIdDetail, UserIdDetail, UserIdDetail.UserRolesList, UserIdDetail.UserRolesList.UserIdDetail, UserIdDetail.UserRolesList.RoleIdDetail, UserIdDetail.UserRolesList.GrantedByDetail, UserIdDetail.TeamsList, UserIdDetail.TeamsList.TeamMembersList, UserIdDetail.TeamsList.OrganizationIdDetail, UserIdDetail.TeamsList.LeadIdDetail, UserIdDetail.TeamMembersList, UserIdDetail.TeamMembersList.TeamIdDetail, UserIdDetail.TeamMembersList.UserIdDetail, UserIdDetail.ApiKeysList, UserIdDetail.ApiKeysList.UserIdDetail, UserIdDetail.ApiKeysList.OrganizationIdDetail, UserIdDetail.SessionsList, UserIdDetail.SessionsList.UserIdDetail, UserIdDetail.InvitationsList, UserIdDetail.InvitationsList.OrganizationIdDetail, UserIdDetail.InvitationsList.InvitedByDetail, UserIdDetail.InvitationsList.RoleIdDetail, UserIdDetail.OrganizationIdDetail, UserIdDetail.OrganizationIdDetail.OrganizationsList, UserIdDetail.OrganizationIdDetail.UsersList, UserIdDetail.OrganizationIdDetail.RolesList, UserIdDetail.OrganizationIdDetail.TeamsList, UserIdDetail.OrganizationIdDetail.ApiKeysList, UserIdDetail.OrganizationIdDetail.InvitationsList, UserIdDetail.OrganizationIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Teams, Teams.Organizations, Teams.Organizations.Organizations, Teams.Users, Teams.Users.Organizations, Users, Users.Organizations, Users.Organizations.Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `teamId` | `string (uuid)` | No | Filter by team_id |
| `userId` | `string (uuid)` | No | Filter by user_id |
| `role` | `string` | No | Filter by role |
| `joinedAt` | `string (date-time)` | No | Filter by joined_at |

**Response:** `TeamMembers[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/team-members/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search TeamMembers (POST)

```
POST /api/v1/team-members/search
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

**Response:** `TeamMembers[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/team-members/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate TeamMembers

```
GET /api/v1/team-members/pagination
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
| `preloads` | `string` | No | Available: TeamIdDetail, TeamIdDetail.TeamMembersList, TeamIdDetail.TeamMembersList.TeamIdDetail, TeamIdDetail.TeamMembersList.UserIdDetail, TeamIdDetail.OrganizationIdDetail, TeamIdDetail.OrganizationIdDetail.OrganizationsList, TeamIdDetail.OrganizationIdDetail.UsersList, TeamIdDetail.OrganizationIdDetail.RolesList, TeamIdDetail.OrganizationIdDetail.TeamsList, TeamIdDetail.OrganizationIdDetail.ApiKeysList, TeamIdDetail.OrganizationIdDetail.InvitationsList, TeamIdDetail.OrganizationIdDetail.ParentIdDetail, TeamIdDetail.LeadIdDetail, TeamIdDetail.LeadIdDetail.UserRolesList, TeamIdDetail.LeadIdDetail.TeamsList, TeamIdDetail.LeadIdDetail.TeamMembersList, TeamIdDetail.LeadIdDetail.ApiKeysList, TeamIdDetail.LeadIdDetail.SessionsList, TeamIdDetail.LeadIdDetail.InvitationsList, TeamIdDetail.LeadIdDetail.OrganizationIdDetail, UserIdDetail, UserIdDetail.UserRolesList, UserIdDetail.UserRolesList.UserIdDetail, UserIdDetail.UserRolesList.RoleIdDetail, UserIdDetail.UserRolesList.GrantedByDetail, UserIdDetail.TeamsList, UserIdDetail.TeamsList.TeamMembersList, UserIdDetail.TeamsList.OrganizationIdDetail, UserIdDetail.TeamsList.LeadIdDetail, UserIdDetail.TeamMembersList, UserIdDetail.TeamMembersList.TeamIdDetail, UserIdDetail.TeamMembersList.UserIdDetail, UserIdDetail.ApiKeysList, UserIdDetail.ApiKeysList.UserIdDetail, UserIdDetail.ApiKeysList.OrganizationIdDetail, UserIdDetail.SessionsList, UserIdDetail.SessionsList.UserIdDetail, UserIdDetail.InvitationsList, UserIdDetail.InvitationsList.OrganizationIdDetail, UserIdDetail.InvitationsList.InvitedByDetail, UserIdDetail.InvitationsList.RoleIdDetail, UserIdDetail.OrganizationIdDetail, UserIdDetail.OrganizationIdDetail.OrganizationsList, UserIdDetail.OrganizationIdDetail.UsersList, UserIdDetail.OrganizationIdDetail.RolesList, UserIdDetail.OrganizationIdDetail.TeamsList, UserIdDetail.OrganizationIdDetail.ApiKeysList, UserIdDetail.OrganizationIdDetail.InvitationsList, UserIdDetail.OrganizationIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Teams, Teams.Organizations, Teams.Organizations.Organizations, Teams.Users, Teams.Users.Organizations, Users, Users.Organizations, Users.Organizations.Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `teamId` | `string (uuid)` | No | Filter by team_id |
| `userId` | `string (uuid)` | No | Filter by user_id |
| `role` | `string` | No | Filter by role |
| `joinedAt` | `string (date-time)` | No | Filter by joined_at |

**Response:** `PaginationResponse<TeamMembers>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/team-members/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate TeamMembers (POST)

```
POST /api/v1/team-members/pagination
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

**Response:** `PaginationResponse<TeamMembers>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/team-members/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create TeamMembers

```
POST /api/v1/team-members/
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
  teamId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  userId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  role?: string  // e.g. example_role
  joinedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `TeamMembers`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/team-members/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create TeamMembers

```
POST /api/v1/team-members/bulk/
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
  teamId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  userId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  role?: string  // e.g. example_role
  joinedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `TeamMembers[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/team-members/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find TeamMembers by ID

```
GET /api/v1/team-members/with-id/:id
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

**Response:** `TeamMembers`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/team-members/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update TeamMembers

```
PUT /api/v1/team-members/with-id/:id
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
  teamId?: string
  userId?: string
  role?: string
  joinedAt?: string
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
  "http://localhost:3000/api/v1/team-members/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update TeamMembers

```
PUT /api/v1/team-members/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: TeamMembersEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/team-members/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete TeamMembers

```
DELETE /api/v1/team-members/with-id/:id
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
  "http://localhost:3000/api/v1/team-members/with-id/:id"
```

</details>

---

:::


::::
