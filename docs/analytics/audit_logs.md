---
title: AuditLogs
---

# AuditLogs

**Table:** `analytics.audit_logs`

**Base path:** `/audit-logs`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `users` | `user_id` | `audit_logs_user_id_fkey` | [Users](./users) |


## Entity Relationship Diagram

erDiagram
    AuditLogs }o--|| Users : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `user_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `users` |
| 4 | `action` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `resource_type` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 6 | `resource_id` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `severity` | `USER-DEFINED` | `AnalyticsEventSeverity` | `"debug" \| "info" \| "warning" \| "error" \| "critical"` | NO | `'info'::analytics.event_severity` | - | - |
| 8 | `ip_address` | `inet` | `interface{}` | `unknown` | YES | - | - | - |
| 9 | `user_agent` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 10 | `old_values` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | YES | - | - | - |
| 11 | `new_values` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | YES | - | - | - |
| 12 | `metadata` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 13 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `user_id` | `users` | `audit_logs_user_id_fkey` |

## Enum Types

### EventSeverity

| Value | Go Constant |
|-------|-------------|
| `debug` | `AnalyticsEventSeverityDebug` |
| `info` | `AnalyticsEventSeverityInfo` |
| `warning` | `AnalyticsEventSeverityWarning` |
| `error` | `AnalyticsEventSeverityError` |
| `critical` | `AnalyticsEventSeverityCritical` |


## Go Generated Code

> 📂 Source: [📄 `AuditLogs.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AuditLogs.go) · [📄 `AuditLogs.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AuditLogs.go) · [📄 `AuditLogs.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/AuditLogs.go)

### Structs

:::tabs

== Form

#### AuditLogsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AuditLogs.go#:~:text=type%20AuditLogsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Action` | `string` | `action` | NO |
| `ResourceType` | `string` | `resourceType` | NO |
| `ResourceId` | `string` | `resourceId` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `IpAddress` | `*interface{}` | `ipAddress` | YES |
| `UserAgent` | `string` | `userAgent` | NO |
| `OldValues` | `*json.RawMessage` | `oldValues` | YES |
| `NewValues` | `*json.RawMessage` | `newValues` | YES |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### AuditLogs [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AuditLogs.go#:~:text=type%20AuditLogs%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Action` | `string` | `action` | NO |
| `ResourceType` | `string` | `resourceType` | NO |
| `ResourceId` | `string` | `resourceId` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `IpAddress` | `*interface{}` | `ipAddress` | YES |
| `UserAgent` | `string` | `userAgent` | NO |
| `OldValues` | `*json.RawMessage` | `oldValues` | YES |
| `NewValues` | `*json.RawMessage` | `newValues` | YES |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### AuditLogsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AuditLogs.go#:~:text=type%20AuditLogsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Action` | `*string` | `action` | YES |
| `ResourceType` | `*string` | `resourceType` | YES |
| `ResourceId` | `*string` | `resourceId` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `IpAddress` | `*interface{}` | `ipAddress` | YES |
| `UserAgent` | `*string` | `userAgent` | YES |
| `OldValues` | `*json.RawMessage` | `oldValues` | YES |
| `NewValues` | `*json.RawMessage` | `newValues` | YES |
| `Metadata` | `*json.RawMessage` | `metadata` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### AuditLogsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AuditLogs.go#:~:text=type%20AuditLogsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Action` | `*string` | `action` | YES |
| `ResourceType` | `*string` | `resourceType` | YES |
| `ResourceId` | `*string` | `resourceId` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `IpAddress` | `*interface{}` | `ipAddress` | YES |
| `UserAgent` | `*string` | `userAgent` | YES |
| `OldValues` | `*json.RawMessage` | `oldValues` | YES |
| `NewValues` | `*json.RawMessage` | `newValues` | YES |
| `Metadata` | `*json.RawMessage` | `metadata` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### AuditLogsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AuditLogs.go#:~:text=type%20AuditLogsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Action` | `string` | `action` | NO |
| `ResourceType` | `string` | `resourceType` | NO |
| `ResourceId` | `string` | `resourceId` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `IpAddress` | `*interface{}` | `ipAddress` | YES |
| `UserAgent` | `string` | `userAgent` | NO |
| `OldValues` | `*json.RawMessage` | `oldValues` | YES |
| `NewValues` | `*json.RawMessage` | `newValues` | YES |
| `Metadata` | `json.RawMessage` | `metadata` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### AuditLogsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AuditLogs.go#:~:text=type%20AuditLogsBatchUpdate%20struct)

```go
type AuditLogsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AuditLogs.go#:~:text=)%20CreateAuditLogs() | `(AuditLogsService) CreateAuditLogs(data AuditLogsForm) (AuditLogsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AuditLogs.go#:~:text=)%20CreateAuditLogsMultiple() | `(AuditLogsService) CreateAuditLogsMultiple(data []AuditLogsForm) ([]AuditLogsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AuditLogs.go#:~:text=)%20UpdateAuditLogs() | `(AuditLogsService) UpdateAuditLogs(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AuditLogs.go#:~:text=)%20UpdateAuditLogsMultiple() | `(AuditLogsService) UpdateAuditLogsMultiple(data []AuditLogsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AuditLogs.go#:~:text=)%20DeleteAuditLogs() | `(AuditLogsService) DeleteAuditLogs(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/audit-logs/` | Search with query params |
| `GET` | `/audit-logs/pagination` | Paginated listing |
| `POST` | `/audit-logs/` | Create single record |
| `POST` | `/audit-logs/bulk/` | Create multiple records |
| `PUT` | `/audit-logs/bulk/` | Batch update |
| `GET` | `/audit-logs/with-id/:id` | Get by ID |
| `PUT` | `/audit-logs/with-id/:id` | Update by ID |
| `DELETE` | `/audit-logs/with-id/:id` | Delete by ID |

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
| `dashboard_count` | - | `integer` | `/rpc/dashboard_count` |
| `event_count_by_severity` | `p_severity text` | `integer` | `/rpc/event_count_by_severity` |
| `unread_notification_count` | `p_user_id uuid` | `integer` | `/rpc/unread_notification_count` |


:::tab Frontend

## TypeScript Types & Hooks

:::tabs

== Interfaces

```typescript
export type AnalyticsEventSeverity =
  | "debug"
  | "info"
  | "warning"
  | "error"
  | "critical"

export const AnalyticsEventSeverityValues = ["debug", "info", "warning", "error", "critical"] as const;

export interface AuditLogs {
  id: string;
  name: string;
  userId?: string;
  action: string;
  resourceType: string;
  resourceId: string;
  severity: AnalyticsEventSeverity;
  ipAddress?: unknown;
  userAgent: string;
  oldValues?: Record<string, unknown>;
  newValues?: Record<string, unknown>;
  metadata: Record<string, unknown>;
  createdAt: string;
}

export interface AuditLogsForm {
  name: string;
  userId?: string;
  action: string;
  resourceType: string;
  resourceId: string;
  severity: AnalyticsEventSeverity;
  ipAddress?: unknown;
  userAgent: string;
  oldValues?: Record<string, unknown>;
  newValues?: Record<string, unknown>;
  metadata: Record<string, unknown>;
  createdAt: string;
}

export interface AuditLogsEdit {
  id: string;
  name: string;
  userId?: string;
  action: string;
  resourceType: string;
  resourceId: string;
  severity: AnalyticsEventSeverity;
  ipAddress?: unknown;
  userAgent: string;
  oldValues?: Record<string, unknown>;
  newValues?: Record<string, unknown>;
  metadata: Record<string, unknown>;
  createdAt: string;
}

export interface AuditLogsPage {
  data: AuditLogs[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type AuditLogsPathQuery = {
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

const AuditLogsKeys = {
  all: ["audit_logs"] as const,
  lists: () => [...AuditLogsKeys.all, "list"] as const,
  detail: (id: any) => [...AuditLogsKeys.all, "detail", id] as const,
} as const;

export function useAuditLogsList(query?: AuditLogsPathQuery) {
  return useQuery({
    queryKey: [...AuditLogsKeys.lists(), query],
    queryFn: () => fetch(`/audit-logs/pagination`, { method: "GET" }).then(r => r.json()) as Promise<AuditLogsPage>,
  });
}

export function useAuditLogsDetail(id: any) {
  return useQuery({
    queryKey: AuditLogsKeys.detail(id),
    queryFn: () => fetch(`/audit-logs/with-id/:id`).then(r => r.json()) as Promise<AuditLogs>,
  });
}

export function useCreateAuditLogs() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: AuditLogsForm) =>
      fetch("/audit-logs/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AuditLogsKeys.lists() }),
  });
}

export function useUpdateAuditLogs() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: AuditLogsEdit }) =>
      fetch(`/audit-logs/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AuditLogsKeys.all }),
  });
}

export function useDeleteAuditLogs() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/audit-logs/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AuditLogsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

const AnalyticsEventSeveritySchema = z.enum(["debug", "info", "warning", "error", "critical"]);

export const AuditLogsFormSchema = z.object({
  name: z.string(),
  userId: z.string().uuid().optional(),
  action: z.string(),
  resourceType: z.string(),
  resourceId: z.string(),
  severity: AnalyticsEventSeveritySchema,
  ipAddress: z.unknown().optional(),
  userAgent: z.string(),
  oldValues: z.record(z.unknown()).optional(),
  newValues: z.record(z.unknown()).optional(),
  metadata: z.record(z.unknown()),
  createdAt: z.string().datetime(),
});

export type AuditLogsFormInput = z.infer<typeof AuditLogsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './audit_logs.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search AuditLogs

```
GET /api/v1/audit-logs/
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
| `preloads` | `string` | No | Comma-separated relations to preload |
| `joins` | `string` | No | Available: Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `userId` | `string (uuid)` | No | Filter by user_id |
| `action` | `string` | No | Filter by action |
| `resourceType` | `string` | No | Filter by resource_type |
| `resourceId` | `string` | No | Filter by resource_id |
| `severity` | `string` | No | Filter by severity |
| `ipAddress` | `string` | No | Filter by ip_address |
| `userAgent` | `string` | No | Filter by user_agent |
| `oldValues` | `string` | No | Filter by old_values |
| `newValues` | `string` | No | Filter by new_values |
| `metadata` | `string` | No | Filter by metadata |

**Response:** `AuditLogs[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/audit-logs/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search AuditLogs (POST)

```
POST /api/v1/audit-logs/search
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

**Response:** `AuditLogs[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/audit-logs/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate AuditLogs

```
GET /api/v1/audit-logs/pagination
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
| `preloads` | `string` | No | Comma-separated relations to preload |
| `joins` | `string` | No | Available: Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `userId` | `string (uuid)` | No | Filter by user_id |
| `action` | `string` | No | Filter by action |
| `resourceType` | `string` | No | Filter by resource_type |
| `resourceId` | `string` | No | Filter by resource_id |
| `severity` | `string` | No | Filter by severity |
| `ipAddress` | `string` | No | Filter by ip_address |
| `userAgent` | `string` | No | Filter by user_agent |
| `oldValues` | `string` | No | Filter by old_values |
| `newValues` | `string` | No | Filter by new_values |
| `metadata` | `string` | No | Filter by metadata |

**Response:** `PaginationResponse<AuditLogs>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/audit-logs/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate AuditLogs (POST)

```
POST /api/v1/audit-logs/pagination
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

**Response:** `PaginationResponse<AuditLogs>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/audit-logs/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create AuditLogs

```
POST /api/v1/audit-logs/
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
  userId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  action?: string  // e.g. example_action
  resourceType?: string  // e.g. example_resource_type
  resourceId?: string  // e.g. example_resource_id
  severity?: "debug" | "info" | "warning" | "error" | "critical"  // e.g. debug
  ipAddress?: unknown  // e.g. value
  userAgent?: string  // e.g. example_user_agent
  oldValues?: Record<string, unknown>  // e.g. map[]
  newValues?: Record<string, unknown>  // e.g. map[]
  metadata?: Record<string, unknown>  // e.g. map[]
}
```

**Response:** `AuditLogs`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/audit-logs/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create AuditLogs

```
POST /api/v1/audit-logs/bulk/
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
  userId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  action?: string  // e.g. example_action
  resourceType?: string  // e.g. example_resource_type
  resourceId?: string  // e.g. example_resource_id
  severity?: "debug" | "info" | "warning" | "error" | "critical"  // e.g. debug
  ipAddress?: unknown  // e.g. value
  userAgent?: string  // e.g. example_user_agent
  oldValues?: Record<string, unknown>  // e.g. map[]
  newValues?: Record<string, unknown>  // e.g. map[]
  metadata?: Record<string, unknown>  // e.g. map[]
}
```

**Response:** `AuditLogs[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/audit-logs/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find AuditLogs by ID

```
GET /api/v1/audit-logs/with-id/:id
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

**Response:** `AuditLogs`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/audit-logs/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update AuditLogs

```
PUT /api/v1/audit-logs/with-id/:id
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
  action?: string
  resourceType?: string
  resourceId?: string
  severity?: "debug" | "info" | "warning" | "error" | "critical"
  ipAddress?: unknown
  userAgent?: string
  oldValues?: Record<string, unknown>
  newValues?: Record<string, unknown>
  metadata?: Record<string, unknown>
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
  "http://localhost:3000/api/v1/audit-logs/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update AuditLogs

```
PUT /api/v1/audit-logs/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: AuditLogsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/audit-logs/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete AuditLogs

```
DELETE /api/v1/audit-logs/with-id/:id
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
  "http://localhost:3000/api/v1/audit-logs/with-id/:id"
```

</details>

---

:::


::::
