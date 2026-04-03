---
title: AlertRules
---

# AlertRules

**Table:** `analytics.alert_rules`

**Base path:** `/alert-rules`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `users` | `owner_id` | `alert_rules_owner_id_fkey` | [Users](./users) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `alert_history` | `alert_rule_id` | `alert_history_alert_rule_id_fkey` | [AlertHistory](./alert_history) |


## Entity Relationship Diagram

erDiagram
    AlertRules }o--|| Users : "FK"
    AlertRules ||--o{ AlertHistory : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `description` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 4 | `owner_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `users` |
| 5 | `condition` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 6 | `action` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 7 | `severity` | `USER-DEFINED` | `AnalyticsEventSeverity` | `"debug" \| "info" \| "warning" \| "error" \| "critical"` | NO | `'warning'::analytics.event_severity` | - | - |
| 8 | `is_active` | `boolean` | `bool` | `boolean` | NO | `true` | - | - |
| 9 | `last_triggered` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 10 | `cooldown_minutes` | `integer` | `int` | `number` | NO | `60` | - | - |
| 11 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 12 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `owner_id` | `users` | `alert_rules_owner_id_fkey` |

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

> 📂 Source: [📄 `AlertRules.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertRules.go) · [📄 `AlertRules.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertRules.go) · [📄 `AlertRules.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/AlertRules.go)

### Structs

:::tabs

== Form

#### AlertRulesForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertRules.go#:~:text=type%20AlertRulesForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Description` | `string` | `description` | NO |
| `OwnerId` | `uuid.UUID` | `ownerId` | NO |
| `Condition` | `json.RawMessage` | `condition` | NO |
| `Action` | `json.RawMessage` | `action` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `LastTriggered` | `*time.Time` | `lastTriggered` | YES |
| `CooldownMinutes` | `int` | `cooldownMinutes` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### AlertRules [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertRules.go#:~:text=type%20AlertRules%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Description` | `string` | `description` | NO |
| `OwnerId` | `uuid.UUID` | `ownerId` | NO |
| `Condition` | `json.RawMessage` | `condition` | NO |
| `Action` | `json.RawMessage` | `action` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `LastTriggered` | `*time.Time` | `lastTriggered` | YES |
| `CooldownMinutes` | `int` | `cooldownMinutes` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### AlertRulesEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertRules.go#:~:text=type%20AlertRulesEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Description` | `*string` | `description` | YES |
| `OwnerId` | `*uuid.UUID` | `ownerId` | YES |
| `Condition` | `*json.RawMessage` | `condition` | YES |
| `Action` | `*json.RawMessage` | `action` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `LastTriggered` | `*time.Time` | `lastTriggered` | YES |
| `CooldownMinutes` | `*int` | `cooldownMinutes` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### AlertRulesFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertRules.go#:~:text=type%20AlertRulesFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Description` | `*string` | `description` | YES |
| `OwnerId` | `*uuid.UUID` | `ownerId` | YES |
| `Condition` | `*json.RawMessage` | `condition` | YES |
| `Action` | `*json.RawMessage` | `action` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `LastTriggered` | `*time.Time` | `lastTriggered` | YES |
| `CooldownMinutes` | `*int` | `cooldownMinutes` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### AlertRulesPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertRules.go#:~:text=type%20AlertRulesPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Description` | `string` | `description` | NO |
| `OwnerId` | `uuid.UUID` | `ownerId` | NO |
| `Condition` | `json.RawMessage` | `condition` | NO |
| `Action` | `json.RawMessage` | `action` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `LastTriggered` | `*time.Time` | `lastTriggered` | YES |
| `CooldownMinutes` | `int` | `cooldownMinutes` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### AlertRulesBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertRules.go#:~:text=type%20AlertRulesBatchUpdate%20struct)

```go
type AlertRulesBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertRules.go#:~:text=)%20CreateAlertRules() | `(AlertRulesService) CreateAlertRules(data AlertRulesForm) (AlertRulesForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertRules.go#:~:text=)%20CreateAlertRulesMultiple() | `(AlertRulesService) CreateAlertRulesMultiple(data []AlertRulesForm) ([]AlertRulesForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertRules.go#:~:text=)%20UpdateAlertRules() | `(AlertRulesService) UpdateAlertRules(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertRules.go#:~:text=)%20UpdateAlertRulesMultiple() | `(AlertRulesService) UpdateAlertRulesMultiple(data []AlertRulesBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertRules.go#:~:text=)%20DeleteAlertRules() | `(AlertRulesService) DeleteAlertRules(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/alert-rules/` | Search with query params |
| `GET` | `/alert-rules/pagination` | Paginated listing |
| `POST` | `/alert-rules/` | Create single record |
| `POST` | `/alert-rules/bulk/` | Create multiple records |
| `PUT` | `/alert-rules/bulk/` | Batch update |
| `GET` | `/alert-rules/with-id/:id` | Get by ID |
| `PUT` | `/alert-rules/with-id/:id` | Update by ID |
| `DELETE` | `/alert-rules/with-id/:id` | Delete by ID |

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

export interface AlertRules {
  id: string;
  name: string;
  description: string;
  ownerId: string;
  condition: Record<string, unknown>;
  action: Record<string, unknown>;
  severity: AnalyticsEventSeverity;
  isActive: boolean;
  lastTriggered?: string;
  cooldownMinutes: number;
  createdAt: string;
  updatedAt: string;
}

export interface AlertRulesForm {
  name: string;
  description: string;
  ownerId: string;
  condition: Record<string, unknown>;
  action: Record<string, unknown>;
  severity: AnalyticsEventSeverity;
  isActive: boolean;
  lastTriggered?: string;
  cooldownMinutes: number;
  createdAt: string;
  updatedAt: string;
}

export interface AlertRulesEdit {
  id: string;
  name: string;
  description: string;
  ownerId: string;
  condition: Record<string, unknown>;
  action: Record<string, unknown>;
  severity: AnalyticsEventSeverity;
  isActive: boolean;
  lastTriggered?: string;
  cooldownMinutes: number;
  createdAt: string;
  updatedAt: string;
}

export interface AlertRulesPage {
  data: AlertRules[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type AlertRulesPathQuery = {
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

const AlertRulesKeys = {
  all: ["alert_rules"] as const,
  lists: () => [...AlertRulesKeys.all, "list"] as const,
  detail: (id: any) => [...AlertRulesKeys.all, "detail", id] as const,
} as const;

export function useAlertRulesList(query?: AlertRulesPathQuery) {
  return useQuery({
    queryKey: [...AlertRulesKeys.lists(), query],
    queryFn: () => fetch(`/alert-rules/pagination`, { method: "GET" }).then(r => r.json()) as Promise<AlertRulesPage>,
  });
}

export function useAlertRulesDetail(id: any) {
  return useQuery({
    queryKey: AlertRulesKeys.detail(id),
    queryFn: () => fetch(`/alert-rules/with-id/:id`).then(r => r.json()) as Promise<AlertRules>,
  });
}

export function useCreateAlertRules() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: AlertRulesForm) =>
      fetch("/alert-rules/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AlertRulesKeys.lists() }),
  });
}

export function useUpdateAlertRules() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: AlertRulesEdit }) =>
      fetch(`/alert-rules/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AlertRulesKeys.all }),
  });
}

export function useDeleteAlertRules() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/alert-rules/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AlertRulesKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

const AnalyticsEventSeveritySchema = z.enum(["debug", "info", "warning", "error", "critical"]);

export const AlertRulesFormSchema = z.object({
  name: z.string(),
  description: z.string(),
  ownerId: z.string().uuid(),
  condition: z.record(z.unknown()),
  action: z.record(z.unknown()),
  severity: AnalyticsEventSeveritySchema,
  isActive: z.boolean(),
  lastTriggered: z.string().datetime().optional(),
  cooldownMinutes: z.number().int(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type AlertRulesFormInput = z.infer<typeof AlertRulesFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './alert_rules.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search AlertRules

```
GET /api/v1/alert-rules/
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
| `preloads` | `string` | No | Available: AlertHistoryList, AlertHistoryList.AlertRuleIdDetail, AlertHistoryList.AlertRuleIdDetail.AlertHistoryList |
| `joins` | `string` | No | Available: Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `description` | `string` | No | Filter by description |
| `ownerId` | `string (uuid)` | No | Filter by owner_id |
| `condition` | `string` | No | Filter by condition |
| `action` | `string` | No | Filter by action |
| `severity` | `string` | No | Filter by severity |
| `isActive` | `boolean` | No | Filter by is_active |
| `lastTriggered` | `string (date-time)` | No | Filter by last_triggered |
| `cooldownMinutes` | `integer` | No | Filter by cooldown_minutes |

**Response:** `AlertRules[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/alert-rules/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search AlertRules (POST)

```
POST /api/v1/alert-rules/search
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

**Response:** `AlertRules[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-rules/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate AlertRules

```
GET /api/v1/alert-rules/pagination
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
| `preloads` | `string` | No | Available: AlertHistoryList, AlertHistoryList.AlertRuleIdDetail, AlertHistoryList.AlertRuleIdDetail.AlertHistoryList |
| `joins` | `string` | No | Available: Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `description` | `string` | No | Filter by description |
| `ownerId` | `string (uuid)` | No | Filter by owner_id |
| `condition` | `string` | No | Filter by condition |
| `action` | `string` | No | Filter by action |
| `severity` | `string` | No | Filter by severity |
| `isActive` | `boolean` | No | Filter by is_active |
| `lastTriggered` | `string (date-time)` | No | Filter by last_triggered |
| `cooldownMinutes` | `integer` | No | Filter by cooldown_minutes |

**Response:** `PaginationResponse<AlertRules>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/alert-rules/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate AlertRules (POST)

```
POST /api/v1/alert-rules/pagination
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

**Response:** `PaginationResponse<AlertRules>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-rules/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create AlertRules

```
POST /api/v1/alert-rules/
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
  description?: string  // e.g. example_description
  ownerId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  condition?: Record<string, unknown>  // e.g. map[]
  action?: Record<string, unknown>  // e.g. map[]
  severity?: "debug" | "info" | "warning" | "error" | "critical"  // e.g. debug
  isActive?: boolean  // e.g. true
  lastTriggered?: string  // e.g. 2026-01-15T10:30:00Z
  cooldownMinutes?: number  // e.g. 1
}
```

**Response:** `AlertRules`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-rules/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create AlertRules

```
POST /api/v1/alert-rules/bulk/
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
  description?: string  // e.g. example_description
  ownerId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  condition?: Record<string, unknown>  // e.g. map[]
  action?: Record<string, unknown>  // e.g. map[]
  severity?: "debug" | "info" | "warning" | "error" | "critical"  // e.g. debug
  isActive?: boolean  // e.g. true
  lastTriggered?: string  // e.g. 2026-01-15T10:30:00Z
  cooldownMinutes?: number  // e.g. 1
}
```

**Response:** `AlertRules[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-rules/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find AlertRules by ID

```
GET /api/v1/alert-rules/with-id/:id
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

**Response:** `AlertRules`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/alert-rules/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update AlertRules

```
PUT /api/v1/alert-rules/with-id/:id
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
  description?: string
  ownerId?: string
  condition?: Record<string, unknown>
  action?: Record<string, unknown>
  severity?: "debug" | "info" | "warning" | "error" | "critical"
  isActive?: boolean
  lastTriggered?: string
  cooldownMinutes?: number
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
  "http://localhost:3000/api/v1/alert-rules/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update AlertRules

```
PUT /api/v1/alert-rules/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: AlertRulesEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-rules/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete AlertRules

```
DELETE /api/v1/alert-rules/with-id/:id
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
  "http://localhost:3000/api/v1/alert-rules/with-id/:id"
```

</details>

---

:::


::::
