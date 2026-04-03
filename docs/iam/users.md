---
title: Users
---

# Users

**Table:** `iam.users`

**Base path:** `/users`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `organizations` | `organization_id` | `users_organization_id_fkey` | [Organizations](./organizations) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `api_keys` | `user_id` | `api_keys_user_id_fkey` | [ApiKeys](./api_keys) |
| `invitations` | `invited_by` | `invitations_invited_by_fkey` | [Invitations](./invitations) |
| `sessions` | `user_id` | `sessions_user_id_fkey` | [Sessions](./sessions) |
| `team_members` | `user_id` | `team_members_user_id_fkey` | [TeamMembers](./team_members) |
| `teams` | `lead_id` | `teams_lead_id_fkey` | [Teams](./teams) |
| `user_roles` | `user_id` | `user_roles_user_id_fkey` | [UserRoles](./user_roles) |
| `user_roles` | `granted_by` | `user_roles_granted_by_fkey` | [UserRoles](./user_roles) |
| `collections` | `created_by` | `collections_created_by_fkey` | [Collections](./collections) |
| `price_history` | `changed_by` | `price_history_changed_by_fkey` | [PriceHistory](./price_history) |
| `product_reviews` | `user_id` | `product_reviews_user_id_fkey` | [ProductReviews](./product_reviews) |
| `products` | `created_by` | `products_created_by_fkey` | [Products](./products) |
| `products` | `updated_by` | `products_updated_by_fkey` | [Products](./products) |
| `customers` | `user_id` | `customers_user_id_fkey` | [Customers](./customers) |
| `order_status_history` | `changed_by` | `order_status_history_changed_by_fkey` | [OrderStatusHistory](./order_status_history) |
| `orders` | `created_by` | `orders_created_by_fkey` | [Orders](./orders) |
| `refunds` | `processed_by` | `refunds_processed_by_fkey` | [Refunds](./refunds) |
| `purchase_orders` | `created_by` | `purchase_orders_created_by_fkey` | [PurchaseOrders](./purchase_orders) |
| `purchase_orders` | `approved_by` | `purchase_orders_approved_by_fkey` | [PurchaseOrders](./purchase_orders) |
| `stock_movements` | `performed_by` | `stock_movements_performed_by_fkey` | [StockMovements](./stock_movements) |
| `warehouses` | `manager_id` | `warehouses_manager_id_fkey` | [Warehouses](./warehouses) |
| `alert_history` | `resolved_by` | `alert_history_resolved_by_fkey` | [AlertHistory](./alert_history) |
| `alert_rules` | `owner_id` | `alert_rules_owner_id_fkey` | [AlertRules](./alert_rules) |
| `audit_logs` | `user_id` | `audit_logs_user_id_fkey` | [AuditLogs](./audit_logs) |
| `dashboards` | `owner_id` | `dashboards_owner_id_fkey` | [Dashboards](./dashboards) |
| `events` | `actor_id` | `events_actor_id_fkey` | [Events](./events) |
| `notifications` | `user_id` | `notifications_user_id_fkey` | [Notifications](./notifications) |
| `report_executions` | `executed_by` | `report_executions_executed_by_fkey` | [ReportExecutions](./report_executions) |
| `reports` | `owner_id` | `reports_owner_id_fkey` | [Reports](./reports) |


## Entity Relationship Diagram

erDiagram
    Users }o--|| Organizations : "FK"
    Users ||--o{ ApiKeys : "ref"
    Users ||--o{ Invitations : "ref"
    Users ||--o{ Sessions : "ref"
    Users ||--o{ TeamMembers : "ref"
    Users ||--o{ Teams : "ref"
    Users ||--o{ UserRoles : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `email` | `text` | `string` | `string` | NO | - | `UQ` | - |
| 3 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 4 | `display_name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `avatar_url` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 6 | `phone` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `status` | `USER-DEFINED` | `IamUserStatus` | `"active" \| "inactive" \| "suspended" \| "pending_verification"` | NO | `'pending_verification'::iam.user_status` | - | - |
| 8 | `auth_provider` | `USER-DEFINED` | `IamAuthProvider` | `"local" \| "google" \| "github" \| "azure_ad" \| "saml"` | NO | `'local'::iam.auth_provider` | - | - |
| 9 | `organization_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `organizations` |
| 10 | `metadata` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 11 | `last_login_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 12 | `email_verified` | `boolean` | `bool` | `boolean` | NO | `false` | - | - |
| 13 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 14 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 15 | `deleted_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `organization_id` | `organizations` | `users_organization_id_fkey` |

## Unique Keys

- `email` (`text`)

## Enum Types

### UserStatus

| Value | Go Constant |
|-------|-------------|
| `active` | `IamUserStatusActive` |
| `inactive` | `IamUserStatusInactive` |
| `suspended` | `IamUserStatusSuspended` |
| `pending_verification` | `IamUserStatusPendingVerification` |

### AuthProvider

| Value | Go Constant |
|-------|-------------|
| `local` | `IamAuthProviderLocal` |
| `google` | `IamAuthProviderGoogle` |
| `github` | `IamAuthProviderGithub` |
| `azure_ad` | `IamAuthProviderAzureAd` |
| `saml` | `IamAuthProviderSaml` |


## Go Generated Code

> 📂 Source: [📄 `Users.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Users.go) · [📄 `Users.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Users.go) · [📄 `Users.go`](https://github.com/meftunca/data-bridge-examples/blob/main//iam/controllers/Users.go)

### Structs

:::tabs

== Form

#### UsersForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Users.go#:~:text=type%20UsersForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Email` | `string` | `email` | NO |
| `Name` | `string` | `name` | NO |
| `DisplayName` | `string` | `displayName` | NO |
| `AvatarUrl` | `string` | `avatarUrl` | NO |
| `Phone` | `string` | `phone` | NO |
| `Status` | `IamUserStatus` | `status` | NO |
| `AuthProvider` | `IamAuthProvider` | `authProvider` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `LastLoginAt` | `*time.Time` | `lastLoginAt` | YES |
| `EmailVerified` | `bool` | `emailVerified` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |
| `DeletedAt` | `*time.Time` | `deletedAt` | YES |

== Model

#### Users [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Users.go#:~:text=type%20Users%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Email` | `string` | `email` | NO |
| `Name` | `string` | `name` | NO |
| `DisplayName` | `string` | `displayName` | NO |
| `AvatarUrl` | `string` | `avatarUrl` | NO |
| `Phone` | `string` | `phone` | NO |
| `Status` | `IamUserStatus` | `status` | NO |
| `AuthProvider` | `IamAuthProvider` | `authProvider` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `LastLoginAt` | `*time.Time` | `lastLoginAt` | YES |
| `EmailVerified` | `bool` | `emailVerified` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |
| `DeletedAt` | `*time.Time` | `deletedAt` | YES |

== Edit

#### UsersEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Users.go#:~:text=type%20UsersEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Email` | `*string` | `email` | YES |
| `Name` | `*string` | `name` | YES |
| `DisplayName` | `*string` | `displayName` | YES |
| `AvatarUrl` | `*string` | `avatarUrl` | YES |
| `Phone` | `*string` | `phone` | YES |
| `Status` | `*IamUserStatus` | `status` | YES |
| `AuthProvider` | `*IamAuthProvider` | `authProvider` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `Metadata` | `*json.RawMessage` | `metadata` | YES |
| `LastLoginAt` | `*time.Time` | `lastLoginAt` | YES |
| `EmailVerified` | `*bool` | `emailVerified` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |
| `DeletedAt` | `*time.Time` | `deletedAt` | YES |

== Filter

#### UsersFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Users.go#:~:text=type%20UsersFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Email` | `*string` | `email` | YES |
| `Name` | `*string` | `name` | YES |
| `DisplayName` | `*string` | `displayName` | YES |
| `AvatarUrl` | `*string` | `avatarUrl` | YES |
| `Phone` | `*string` | `phone` | YES |
| `Status` | `*IamUserStatus` | `status` | YES |
| `AuthProvider` | `*IamAuthProvider` | `authProvider` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `Metadata` | `*json.RawMessage` | `metadata` | YES |
| `LastLoginAt` | `*time.Time` | `lastLoginAt` | YES |
| `EmailVerified` | `*bool` | `emailVerified` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |
| `DeletedAt` | `*time.Time` | `deletedAt` | YES |

== Page

#### UsersPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Users.go#:~:text=type%20UsersPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Email` | `string` | `email` | NO |
| `Name` | `string` | `name` | NO |
| `DisplayName` | `string` | `displayName` | NO |
| `AvatarUrl` | `string` | `avatarUrl` | NO |
| `Phone` | `string` | `phone` | NO |
| `Status` | `IamUserStatus` | `status` | NO |
| `AuthProvider` | `IamAuthProvider` | `authProvider` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `LastLoginAt` | `*time.Time` | `lastLoginAt` | YES |
| `EmailVerified` | `bool` | `emailVerified` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |
| `DeletedAt` | `*time.Time` | `deletedAt` | YES |

== BatchUpdate

#### UsersBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//iam/structures/Users.go#:~:text=type%20UsersBatchUpdate%20struct)

```go
type UsersBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Users.go#:~:text=)%20CreateUsers() | `(UsersService) CreateUsers(data UsersForm) (UsersForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Users.go#:~:text=)%20CreateUsersMultiple() | `(UsersService) CreateUsersMultiple(data []UsersForm) ([]UsersForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Users.go#:~:text=)%20UpdateUsers() | `(UsersService) UpdateUsers(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Users.go#:~:text=)%20UpdateUsersMultiple() | `(UsersService) UpdateUsersMultiple(data []UsersBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//iam/services/Users.go#:~:text=)%20DeleteUsers() | `(UsersService) DeleteUsers(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/users/` | Search with query params |
| `GET` | `/users/pagination` | Paginated listing |
| `POST` | `/users/` | Create single record |
| `POST` | `/users/bulk/` | Create multiple records |
| `PUT` | `/users/bulk/` | Batch update |
| `GET` | `/users/with-id/:id` | Get by ID |
| `PUT` | `/users/with-id/:id` | Update by ID |
| `DELETE` | `/users/with-id/:id` | Delete by ID |

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
export type IamAuthProvider =
  | "local"
  | "google"
  | "github"
  | "azure_ad"
  | "saml"

export const IamAuthProviderValues = ["local", "google", "github", "azure_ad", "saml"] as const;

export type IamUserStatus =
  | "active"
  | "inactive"
  | "suspended"
  | "pending_verification"

export const IamUserStatusValues = ["active", "inactive", "suspended", "pending_verification"] as const;

export interface Users {
  id: string;
  email: string;
  name: string;
  displayName: string;
  avatarUrl: string;
  phone: string;
  status: IamUserStatus;
  authProvider: IamAuthProvider;
  organizationId?: string;
  metadata: Record<string, unknown>;
  lastLoginAt?: string;
  emailVerified: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface UsersForm {
  email: string;
  name: string;
  displayName: string;
  avatarUrl: string;
  phone: string;
  status: IamUserStatus;
  authProvider: IamAuthProvider;
  organizationId?: string;
  metadata: Record<string, unknown>;
  lastLoginAt?: string;
  emailVerified: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface UsersEdit {
  id: string;
  email: string;
  name: string;
  displayName: string;
  avatarUrl: string;
  phone: string;
  status: IamUserStatus;
  authProvider: IamAuthProvider;
  organizationId?: string;
  metadata: Record<string, unknown>;
  lastLoginAt?: string;
  emailVerified: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface UsersPage {
  data: Users[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type UsersPathQuery = {
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

const UsersKeys = {
  all: ["users"] as const,
  lists: () => [...UsersKeys.all, "list"] as const,
  detail: (id: any) => [...UsersKeys.all, "detail", id] as const,
} as const;

export function useUsersList(query?: UsersPathQuery) {
  return useQuery({
    queryKey: [...UsersKeys.lists(), query],
    queryFn: () => fetch(`/users/pagination`, { method: "GET" }).then(r => r.json()) as Promise<UsersPage>,
  });
}

export function useUsersDetail(id: any) {
  return useQuery({
    queryKey: UsersKeys.detail(id),
    queryFn: () => fetch(`/users/with-id/:id`).then(r => r.json()) as Promise<Users>,
  });
}

export function useCreateUsers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: UsersForm) =>
      fetch("/users/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: UsersKeys.lists() }),
  });
}

export function useUpdateUsers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: UsersEdit }) =>
      fetch(`/users/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: UsersKeys.all }),
  });
}

export function useDeleteUsers() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/users/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: UsersKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

const IamAuthProviderSchema = z.enum(["local", "google", "github", "azure_ad", "saml"]);

const IamUserStatusSchema = z.enum(["active", "inactive", "suspended", "pending_verification"]);

export const UsersFormSchema = z.object({
  email: z.string(),
  name: z.string(),
  displayName: z.string(),
  avatarUrl: z.string(),
  phone: z.string(),
  status: IamUserStatusSchema,
  authProvider: IamAuthProviderSchema,
  organizationId: z.string().uuid().optional(),
  metadata: z.record(z.unknown()),
  lastLoginAt: z.string().datetime().optional(),
  emailVerified: z.boolean(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
  deletedAt: z.string().datetime().optional(),
});

export type UsersFormInput = z.infer<typeof UsersFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './users.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Users

```
GET /api/v1/users/
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
| `preloads` | `string` | No | Available: UserRolesList, UserRolesList.UserIdDetail, UserRolesList.UserIdDetail.UserRolesList, UserRolesList.UserIdDetail.TeamsList, UserRolesList.UserIdDetail.TeamMembersList, UserRolesList.UserIdDetail.ApiKeysList, UserRolesList.UserIdDetail.SessionsList, UserRolesList.UserIdDetail.InvitationsList, UserRolesList.UserIdDetail.OrganizationIdDetail, UserRolesList.RoleIdDetail, UserRolesList.RoleIdDetail.RolePermissionsList, UserRolesList.RoleIdDetail.UserRolesList, UserRolesList.RoleIdDetail.InvitationsList, UserRolesList.RoleIdDetail.OrganizationIdDetail, UserRolesList.GrantedByDetail, UserRolesList.GrantedByDetail.UserRolesList, UserRolesList.GrantedByDetail.TeamsList, UserRolesList.GrantedByDetail.TeamMembersList, UserRolesList.GrantedByDetail.ApiKeysList, UserRolesList.GrantedByDetail.SessionsList, UserRolesList.GrantedByDetail.InvitationsList, UserRolesList.GrantedByDetail.OrganizationIdDetail, TeamsList, TeamsList.TeamMembersList, TeamsList.TeamMembersList.TeamIdDetail, TeamsList.TeamMembersList.UserIdDetail, TeamsList.OrganizationIdDetail, TeamsList.OrganizationIdDetail.OrganizationsList, TeamsList.OrganizationIdDetail.UsersList, TeamsList.OrganizationIdDetail.RolesList, TeamsList.OrganizationIdDetail.TeamsList, TeamsList.OrganizationIdDetail.ApiKeysList, TeamsList.OrganizationIdDetail.InvitationsList, TeamsList.OrganizationIdDetail.ParentIdDetail, TeamsList.LeadIdDetail, TeamsList.LeadIdDetail.UserRolesList, TeamsList.LeadIdDetail.TeamsList, TeamsList.LeadIdDetail.TeamMembersList, TeamsList.LeadIdDetail.ApiKeysList, TeamsList.LeadIdDetail.SessionsList, TeamsList.LeadIdDetail.InvitationsList, TeamsList.LeadIdDetail.OrganizationIdDetail, TeamMembersList, TeamMembersList.TeamIdDetail, TeamMembersList.TeamIdDetail.TeamMembersList, TeamMembersList.TeamIdDetail.OrganizationIdDetail, TeamMembersList.TeamIdDetail.LeadIdDetail, TeamMembersList.UserIdDetail, TeamMembersList.UserIdDetail.UserRolesList, TeamMembersList.UserIdDetail.TeamsList, TeamMembersList.UserIdDetail.TeamMembersList, TeamMembersList.UserIdDetail.ApiKeysList, TeamMembersList.UserIdDetail.SessionsList, TeamMembersList.UserIdDetail.InvitationsList, TeamMembersList.UserIdDetail.OrganizationIdDetail, ApiKeysList, ApiKeysList.UserIdDetail, ApiKeysList.UserIdDetail.UserRolesList, ApiKeysList.UserIdDetail.TeamsList, ApiKeysList.UserIdDetail.TeamMembersList, ApiKeysList.UserIdDetail.ApiKeysList, ApiKeysList.UserIdDetail.SessionsList, ApiKeysList.UserIdDetail.InvitationsList, ApiKeysList.UserIdDetail.OrganizationIdDetail, ApiKeysList.OrganizationIdDetail, ApiKeysList.OrganizationIdDetail.OrganizationsList, ApiKeysList.OrganizationIdDetail.UsersList, ApiKeysList.OrganizationIdDetail.RolesList, ApiKeysList.OrganizationIdDetail.TeamsList, ApiKeysList.OrganizationIdDetail.ApiKeysList, ApiKeysList.OrganizationIdDetail.InvitationsList, ApiKeysList.OrganizationIdDetail.ParentIdDetail, SessionsList, SessionsList.UserIdDetail, SessionsList.UserIdDetail.UserRolesList, SessionsList.UserIdDetail.TeamsList, SessionsList.UserIdDetail.TeamMembersList, SessionsList.UserIdDetail.ApiKeysList, SessionsList.UserIdDetail.SessionsList, SessionsList.UserIdDetail.InvitationsList, SessionsList.UserIdDetail.OrganizationIdDetail, InvitationsList, InvitationsList.OrganizationIdDetail, InvitationsList.OrganizationIdDetail.OrganizationsList, InvitationsList.OrganizationIdDetail.UsersList, InvitationsList.OrganizationIdDetail.RolesList, InvitationsList.OrganizationIdDetail.TeamsList, InvitationsList.OrganizationIdDetail.ApiKeysList, InvitationsList.OrganizationIdDetail.InvitationsList, InvitationsList.OrganizationIdDetail.ParentIdDetail, InvitationsList.InvitedByDetail, InvitationsList.InvitedByDetail.UserRolesList, InvitationsList.InvitedByDetail.TeamsList, InvitationsList.InvitedByDetail.TeamMembersList, InvitationsList.InvitedByDetail.ApiKeysList, InvitationsList.InvitedByDetail.SessionsList, InvitationsList.InvitedByDetail.InvitationsList, InvitationsList.InvitedByDetail.OrganizationIdDetail, InvitationsList.RoleIdDetail, InvitationsList.RoleIdDetail.RolePermissionsList, InvitationsList.RoleIdDetail.UserRolesList, InvitationsList.RoleIdDetail.InvitationsList, InvitationsList.RoleIdDetail.OrganizationIdDetail, OrganizationIdDetail, OrganizationIdDetail.OrganizationsList, OrganizationIdDetail.UsersList, OrganizationIdDetail.UsersList.UserRolesList, OrganizationIdDetail.UsersList.TeamsList, OrganizationIdDetail.UsersList.TeamMembersList, OrganizationIdDetail.UsersList.ApiKeysList, OrganizationIdDetail.UsersList.SessionsList, OrganizationIdDetail.UsersList.InvitationsList, OrganizationIdDetail.UsersList.OrganizationIdDetail, OrganizationIdDetail.RolesList, OrganizationIdDetail.RolesList.RolePermissionsList, OrganizationIdDetail.RolesList.UserRolesList, OrganizationIdDetail.RolesList.InvitationsList, OrganizationIdDetail.RolesList.OrganizationIdDetail, OrganizationIdDetail.TeamsList, OrganizationIdDetail.TeamsList.TeamMembersList, OrganizationIdDetail.TeamsList.OrganizationIdDetail, OrganizationIdDetail.TeamsList.LeadIdDetail, OrganizationIdDetail.ApiKeysList, OrganizationIdDetail.ApiKeysList.UserIdDetail, OrganizationIdDetail.ApiKeysList.OrganizationIdDetail, OrganizationIdDetail.InvitationsList, OrganizationIdDetail.InvitationsList.OrganizationIdDetail, OrganizationIdDetail.InvitationsList.InvitedByDetail, OrganizationIdDetail.InvitationsList.RoleIdDetail, OrganizationIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Organizations, Organizations.Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `email` | `string` | No | Filter by email |
| `name` | `string` | No | Filter by name |
| `displayName` | `string` | No | Filter by display_name |
| `avatarUrl` | `string` | No | Filter by avatar_url |
| `phone` | `string` | No | Filter by phone |
| `status` | `string` | No | Filter by status |
| `authProvider` | `string` | No | Filter by auth_provider |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `metadata` | `string` | No | Filter by metadata |
| `lastLoginAt` | `string (date-time)` | No | Filter by last_login_at |
| `emailVerified` | `boolean` | No | Filter by email_verified |

**Response:** `Users[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/users/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Users (POST)

```
POST /api/v1/users/search
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

**Response:** `Users[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/users/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Users

```
GET /api/v1/users/pagination
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
| `preloads` | `string` | No | Available: UserRolesList, UserRolesList.UserIdDetail, UserRolesList.UserIdDetail.UserRolesList, UserRolesList.UserIdDetail.TeamsList, UserRolesList.UserIdDetail.TeamMembersList, UserRolesList.UserIdDetail.ApiKeysList, UserRolesList.UserIdDetail.SessionsList, UserRolesList.UserIdDetail.InvitationsList, UserRolesList.UserIdDetail.OrganizationIdDetail, UserRolesList.RoleIdDetail, UserRolesList.RoleIdDetail.RolePermissionsList, UserRolesList.RoleIdDetail.UserRolesList, UserRolesList.RoleIdDetail.InvitationsList, UserRolesList.RoleIdDetail.OrganizationIdDetail, UserRolesList.GrantedByDetail, UserRolesList.GrantedByDetail.UserRolesList, UserRolesList.GrantedByDetail.TeamsList, UserRolesList.GrantedByDetail.TeamMembersList, UserRolesList.GrantedByDetail.ApiKeysList, UserRolesList.GrantedByDetail.SessionsList, UserRolesList.GrantedByDetail.InvitationsList, UserRolesList.GrantedByDetail.OrganizationIdDetail, TeamsList, TeamsList.TeamMembersList, TeamsList.TeamMembersList.TeamIdDetail, TeamsList.TeamMembersList.UserIdDetail, TeamsList.OrganizationIdDetail, TeamsList.OrganizationIdDetail.OrganizationsList, TeamsList.OrganizationIdDetail.UsersList, TeamsList.OrganizationIdDetail.RolesList, TeamsList.OrganizationIdDetail.TeamsList, TeamsList.OrganizationIdDetail.ApiKeysList, TeamsList.OrganizationIdDetail.InvitationsList, TeamsList.OrganizationIdDetail.ParentIdDetail, TeamsList.LeadIdDetail, TeamsList.LeadIdDetail.UserRolesList, TeamsList.LeadIdDetail.TeamsList, TeamsList.LeadIdDetail.TeamMembersList, TeamsList.LeadIdDetail.ApiKeysList, TeamsList.LeadIdDetail.SessionsList, TeamsList.LeadIdDetail.InvitationsList, TeamsList.LeadIdDetail.OrganizationIdDetail, TeamMembersList, TeamMembersList.TeamIdDetail, TeamMembersList.TeamIdDetail.TeamMembersList, TeamMembersList.TeamIdDetail.OrganizationIdDetail, TeamMembersList.TeamIdDetail.LeadIdDetail, TeamMembersList.UserIdDetail, TeamMembersList.UserIdDetail.UserRolesList, TeamMembersList.UserIdDetail.TeamsList, TeamMembersList.UserIdDetail.TeamMembersList, TeamMembersList.UserIdDetail.ApiKeysList, TeamMembersList.UserIdDetail.SessionsList, TeamMembersList.UserIdDetail.InvitationsList, TeamMembersList.UserIdDetail.OrganizationIdDetail, ApiKeysList, ApiKeysList.UserIdDetail, ApiKeysList.UserIdDetail.UserRolesList, ApiKeysList.UserIdDetail.TeamsList, ApiKeysList.UserIdDetail.TeamMembersList, ApiKeysList.UserIdDetail.ApiKeysList, ApiKeysList.UserIdDetail.SessionsList, ApiKeysList.UserIdDetail.InvitationsList, ApiKeysList.UserIdDetail.OrganizationIdDetail, ApiKeysList.OrganizationIdDetail, ApiKeysList.OrganizationIdDetail.OrganizationsList, ApiKeysList.OrganizationIdDetail.UsersList, ApiKeysList.OrganizationIdDetail.RolesList, ApiKeysList.OrganizationIdDetail.TeamsList, ApiKeysList.OrganizationIdDetail.ApiKeysList, ApiKeysList.OrganizationIdDetail.InvitationsList, ApiKeysList.OrganizationIdDetail.ParentIdDetail, SessionsList, SessionsList.UserIdDetail, SessionsList.UserIdDetail.UserRolesList, SessionsList.UserIdDetail.TeamsList, SessionsList.UserIdDetail.TeamMembersList, SessionsList.UserIdDetail.ApiKeysList, SessionsList.UserIdDetail.SessionsList, SessionsList.UserIdDetail.InvitationsList, SessionsList.UserIdDetail.OrganizationIdDetail, InvitationsList, InvitationsList.OrganizationIdDetail, InvitationsList.OrganizationIdDetail.OrganizationsList, InvitationsList.OrganizationIdDetail.UsersList, InvitationsList.OrganizationIdDetail.RolesList, InvitationsList.OrganizationIdDetail.TeamsList, InvitationsList.OrganizationIdDetail.ApiKeysList, InvitationsList.OrganizationIdDetail.InvitationsList, InvitationsList.OrganizationIdDetail.ParentIdDetail, InvitationsList.InvitedByDetail, InvitationsList.InvitedByDetail.UserRolesList, InvitationsList.InvitedByDetail.TeamsList, InvitationsList.InvitedByDetail.TeamMembersList, InvitationsList.InvitedByDetail.ApiKeysList, InvitationsList.InvitedByDetail.SessionsList, InvitationsList.InvitedByDetail.InvitationsList, InvitationsList.InvitedByDetail.OrganizationIdDetail, InvitationsList.RoleIdDetail, InvitationsList.RoleIdDetail.RolePermissionsList, InvitationsList.RoleIdDetail.UserRolesList, InvitationsList.RoleIdDetail.InvitationsList, InvitationsList.RoleIdDetail.OrganizationIdDetail, OrganizationIdDetail, OrganizationIdDetail.OrganizationsList, OrganizationIdDetail.UsersList, OrganizationIdDetail.UsersList.UserRolesList, OrganizationIdDetail.UsersList.TeamsList, OrganizationIdDetail.UsersList.TeamMembersList, OrganizationIdDetail.UsersList.ApiKeysList, OrganizationIdDetail.UsersList.SessionsList, OrganizationIdDetail.UsersList.InvitationsList, OrganizationIdDetail.UsersList.OrganizationIdDetail, OrganizationIdDetail.RolesList, OrganizationIdDetail.RolesList.RolePermissionsList, OrganizationIdDetail.RolesList.UserRolesList, OrganizationIdDetail.RolesList.InvitationsList, OrganizationIdDetail.RolesList.OrganizationIdDetail, OrganizationIdDetail.TeamsList, OrganizationIdDetail.TeamsList.TeamMembersList, OrganizationIdDetail.TeamsList.OrganizationIdDetail, OrganizationIdDetail.TeamsList.LeadIdDetail, OrganizationIdDetail.ApiKeysList, OrganizationIdDetail.ApiKeysList.UserIdDetail, OrganizationIdDetail.ApiKeysList.OrganizationIdDetail, OrganizationIdDetail.InvitationsList, OrganizationIdDetail.InvitationsList.OrganizationIdDetail, OrganizationIdDetail.InvitationsList.InvitedByDetail, OrganizationIdDetail.InvitationsList.RoleIdDetail, OrganizationIdDetail.ParentIdDetail |
| `joins` | `string` | No | Available: Organizations, Organizations.Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `email` | `string` | No | Filter by email |
| `name` | `string` | No | Filter by name |
| `displayName` | `string` | No | Filter by display_name |
| `avatarUrl` | `string` | No | Filter by avatar_url |
| `phone` | `string` | No | Filter by phone |
| `status` | `string` | No | Filter by status |
| `authProvider` | `string` | No | Filter by auth_provider |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `metadata` | `string` | No | Filter by metadata |
| `lastLoginAt` | `string (date-time)` | No | Filter by last_login_at |
| `emailVerified` | `boolean` | No | Filter by email_verified |

**Response:** `PaginationResponse<Users>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/users/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Users (POST)

```
POST /api/v1/users/pagination
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

**Response:** `PaginationResponse<Users>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/users/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Users

```
POST /api/v1/users/
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
  email: string  // e.g. example_email
  name: string  // e.g. example_name
  displayName?: string  // e.g. example_display_name
  avatarUrl?: string  // e.g. example_avatar_url
  phone?: string  // e.g. example_phone
  status?: "active" | "inactive" | "suspended" | "pending_verification"  // e.g. active
  authProvider?: "local" | "google" | "github" | "azure_ad" | "saml"  // e.g. local
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  metadata?: Record<string, unknown>  // e.g. map[]
  lastLoginAt?: string  // e.g. 2026-01-15T10:30:00Z
  emailVerified?: boolean  // e.g. true
}
```

**Response:** `Users`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/users/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Users

```
POST /api/v1/users/bulk/
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
  email: string  // e.g. example_email
  name: string  // e.g. example_name
  displayName?: string  // e.g. example_display_name
  avatarUrl?: string  // e.g. example_avatar_url
  phone?: string  // e.g. example_phone
  status?: "active" | "inactive" | "suspended" | "pending_verification"  // e.g. active
  authProvider?: "local" | "google" | "github" | "azure_ad" | "saml"  // e.g. local
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  metadata?: Record<string, unknown>  // e.g. map[]
  lastLoginAt?: string  // e.g. 2026-01-15T10:30:00Z
  emailVerified?: boolean  // e.g. true
}
```

**Response:** `Users[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/users/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Users by ID

```
GET /api/v1/users/with-id/:id
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

**Response:** `Users`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/users/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Users

```
PUT /api/v1/users/with-id/:id
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
  email?: string
  name?: string
  displayName?: string
  avatarUrl?: string
  phone?: string
  status?: "active" | "inactive" | "suspended" | "pending_verification"
  authProvider?: "local" | "google" | "github" | "azure_ad" | "saml"
  organizationId?: string
  metadata?: Record<string, unknown>
  lastLoginAt?: string
  emailVerified?: boolean
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
  "http://localhost:3000/api/v1/users/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Users

```
PUT /api/v1/users/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: UsersEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/users/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Users

```
DELETE /api/v1/users/with-id/:id
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
  "http://localhost:3000/api/v1/users/with-id/:id"
```

</details>

---

:::


::::
