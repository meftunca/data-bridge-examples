---
title: AlertHistory
---

# AlertHistory

**Table:** `analytics.alert_history`

**Base path:** `/alert-history`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `alert_rules` | `alert_rule_id` | `alert_history_alert_rule_id_fkey` | [AlertRules](./alert_rules) |
| `users` | `resolved_by` | `alert_history_resolved_by_fkey` | [Users](./users) |


## Entity Relationship Diagram

```mermaid
erDiagram
    AlertHistory }o--|| AlertRules : "FK"
    AlertHistory }o--|| Users : "FK"
```

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `alert_rule_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `alert_rules` |
| 4 | `triggered_value` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 5 | `resolved` | `boolean` | `bool` | `boolean` | NO | `false` | - | - |
| 6 | `resolved_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 7 | `resolved_by` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `users` |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `alert_rule_id` | `alert_rules` | `alert_history_alert_rule_id_fkey` |
| `resolved_by` | `users` | `alert_history_resolved_by_fkey` |


## Go Generated Code

> 📂 Source: [📄 `AlertHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertHistory.go) · [📄 `AlertHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertHistory.go) · [📄 `AlertHistory.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/AlertHistory.go)

### Structs

::::tabs

:::tab Form

#### AlertHistoryForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertHistory.go#:~:text=type%20AlertHistoryForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `AlertRuleId` | `uuid.UUID` | `alertRuleId` | NO |
| `TriggeredValue` | `json.RawMessage` | `triggeredValue` | NO |
| `Resolved` | `bool` | `resolved` | NO |
| `ResolvedAt` | `*time.Time` | `resolvedAt` | YES |
| `ResolvedBy` | `*uuid.UUID` | `resolvedBy` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab Model

#### AlertHistory [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertHistory.go#:~:text=type%20AlertHistory%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `AlertRuleId` | `uuid.UUID` | `alertRuleId` | NO |
| `TriggeredValue` | `json.RawMessage` | `triggeredValue` | NO |
| `Resolved` | `bool` | `resolved` | NO |
| `ResolvedAt` | `*time.Time` | `resolvedAt` | YES |
| `ResolvedBy` | `*uuid.UUID` | `resolvedBy` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab Edit

#### AlertHistoryEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertHistory.go#:~:text=type%20AlertHistoryEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `AlertRuleId` | `*uuid.UUID` | `alertRuleId` | YES |
| `TriggeredValue` | `*json.RawMessage` | `triggeredValue` | YES |
| `Resolved` | `*bool` | `resolved` | YES |
| `ResolvedAt` | `*time.Time` | `resolvedAt` | YES |
| `ResolvedBy` | `*uuid.UUID` | `resolvedBy` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Filter

#### AlertHistoryFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertHistory.go#:~:text=type%20AlertHistoryFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `AlertRuleId` | `*uuid.UUID` | `alertRuleId` | YES |
| `TriggeredValue` | `*json.RawMessage` | `triggeredValue` | YES |
| `Resolved` | `*bool` | `resolved` | YES |
| `ResolvedAt` | `*time.Time` | `resolvedAt` | YES |
| `ResolvedBy` | `*uuid.UUID` | `resolvedBy` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Page

#### AlertHistoryPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertHistory.go#:~:text=type%20AlertHistoryPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `AlertRuleId` | `uuid.UUID` | `alertRuleId` | NO |
| `TriggeredValue` | `json.RawMessage` | `triggeredValue` | NO |
| `Resolved` | `bool` | `resolved` | NO |
| `ResolvedAt` | `*time.Time` | `resolvedAt` | YES |
| `ResolvedBy` | `*uuid.UUID` | `resolvedBy` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

:::tab BatchUpdate

#### AlertHistoryBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/AlertHistory.go#:~:text=type%20AlertHistoryBatchUpdate%20struct)

```go
type AlertHistoryBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertHistory.go#:~:text=)%20CreateAlertHistory() | `(AlertHistoryService) CreateAlertHistory(data AlertHistoryForm) (AlertHistoryForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertHistory.go#:~:text=)%20CreateAlertHistoryMultiple() | `(AlertHistoryService) CreateAlertHistoryMultiple(data []AlertHistoryForm) ([]AlertHistoryForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertHistory.go#:~:text=)%20UpdateAlertHistory() | `(AlertHistoryService) UpdateAlertHistory(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertHistory.go#:~:text=)%20UpdateAlertHistoryMultiple() | `(AlertHistoryService) UpdateAlertHistoryMultiple(data []AlertHistoryBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/AlertHistory.go#:~:text=)%20DeleteAlertHistory() | `(AlertHistoryService) DeleteAlertHistory(id uuid.UUID) error` |

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/alert-history/` | Search with query params |
| `GET` | `/alert-history/pagination` | Paginated listing |
| `POST` | `/alert-history/` | Create single record |
| `POST` | `/alert-history/bulk/` | Create multiple records |
| `PUT` | `/alert-history/bulk/` | Batch update |
| `GET` | `/alert-history/with-id/:id` | Get by ID |
| `PUT` | `/alert-history/with-id/:id` | Update by ID |
| `DELETE` | `/alert-history/with-id/:id` | Delete by ID |

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
| `dashboard_count` | - | `integer` | `/rpc/dashboard_count` |
| `event_count_by_severity` | `p_severity text` | `integer` | `/rpc/event_count_by_severity` |
| `unread_notification_count` | `p_user_id uuid` | `integer` | `/rpc/unread_notification_count` |


:::tab Frontend

## TypeScript Types & Hooks

::::tabs

:::tab Interfaces

```typescript
export interface AlertHistory {
  id: string;
  name: string;
  alertRuleId: string;
  triggeredValue: Record<string, unknown>;
  resolved: boolean;
  resolvedAt?: string;
  resolvedBy?: string;
  createdAt: string;
}

export interface AlertHistoryForm {
  name: string;
  alertRuleId: string;
  triggeredValue: Record<string, unknown>;
  resolved: boolean;
  resolvedAt?: string;
  resolvedBy?: string;
  createdAt: string;
}

export interface AlertHistoryEdit {
  id: string;
  name: string;
  alertRuleId: string;
  triggeredValue: Record<string, unknown>;
  resolved: boolean;
  resolvedAt?: string;
  resolvedBy?: string;
  createdAt: string;
}

export interface AlertHistoryPage {
  data: AlertHistory[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type AlertHistoryPathQuery = {
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

const AlertHistoryKeys = {
  all: ["alert_history"] as const,
  lists: () => [...AlertHistoryKeys.all, "list"] as const,
  detail: (id: any) => [...AlertHistoryKeys.all, "detail", id] as const,
} as const;

export function useAlertHistoryList(query?: AlertHistoryPathQuery) {
  return useQuery({
    queryKey: [...AlertHistoryKeys.lists(), query],
    queryFn: () => fetch(`/alert-history/pagination`, { method: "GET" }).then(r => r.json()) as Promise<AlertHistoryPage>,
  });
}

export function useAlertHistoryDetail(id: any) {
  return useQuery({
    queryKey: AlertHistoryKeys.detail(id),
    queryFn: () => fetch(`/alert-history/with-id/:id`).then(r => r.json()) as Promise<AlertHistory>,
  });
}

export function useCreateAlertHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: AlertHistoryForm) =>
      fetch("/alert-history/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AlertHistoryKeys.lists() }),
  });
}

export function useUpdateAlertHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: AlertHistoryEdit }) =>
      fetch(`/alert-history/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AlertHistoryKeys.all }),
  });
}

export function useDeleteAlertHistory() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/alert-history/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: AlertHistoryKeys.all }),
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

export const AlertHistoryFormSchema = z.object({
  name: z.string(),
  alertRuleId: z.string().uuid(),
  triggeredValue: z.record(z.unknown()),
  resolved: z.boolean(),
  resolvedAt: z.string().datetime().optional(),
  resolvedBy: z.string().uuid().optional(),
  createdAt: z.string().datetime(),
});

export type AlertHistoryFormInput = z.infer<typeof AlertHistoryFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './alert_history.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search AlertHistory

```
GET /api/v1/alert-history/
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
| `preloads` | `string` | No | Available: AlertRuleIdDetail, AlertRuleIdDetail.AlertHistoryList, AlertRuleIdDetail.AlertHistoryList.AlertRuleIdDetail |
| `joins` | `string` | No | Available: AlertRules, AlertRules.Users, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `alertRuleId` | `string (uuid)` | No | Filter by alert_rule_id |
| `triggeredValue` | `string` | No | Filter by triggered_value |
| `resolved` | `boolean` | No | Filter by resolved |
| `resolvedAt` | `string (date-time)` | No | Filter by resolved_at |
| `resolvedBy` | `string (uuid)` | No | Filter by resolved_by |

**Response:** `AlertHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/alert-history/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search AlertHistory (POST)

```
POST /api/v1/alert-history/search
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

**Response:** `AlertHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-history/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate AlertHistory

```
GET /api/v1/alert-history/pagination
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
| `preloads` | `string` | No | Available: AlertRuleIdDetail, AlertRuleIdDetail.AlertHistoryList, AlertRuleIdDetail.AlertHistoryList.AlertRuleIdDetail |
| `joins` | `string` | No | Available: AlertRules, AlertRules.Users, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `alertRuleId` | `string (uuid)` | No | Filter by alert_rule_id |
| `triggeredValue` | `string` | No | Filter by triggered_value |
| `resolved` | `boolean` | No | Filter by resolved |
| `resolvedAt` | `string (date-time)` | No | Filter by resolved_at |
| `resolvedBy` | `string (uuid)` | No | Filter by resolved_by |

**Response:** `PaginationResponse<AlertHistory>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/alert-history/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate AlertHistory (POST)

```
POST /api/v1/alert-history/pagination
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

**Response:** `PaginationResponse<AlertHistory>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-history/pagination"
```

</details>

---

:::tab Create

#### <Badge type="tip" text="POST" /> Create AlertHistory

```
POST /api/v1/alert-history/
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
  alertRuleId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  triggeredValue?: Record<string, unknown>  // e.g. map[]
  resolved?: boolean  // e.g. true
  resolvedAt?: string  // e.g. 2026-01-15T10:30:00Z
  resolvedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
}
```

**Response:** `AlertHistory`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-history/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create AlertHistory

```
POST /api/v1/alert-history/bulk/
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
  alertRuleId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  triggeredValue?: Record<string, unknown>  // e.g. map[]
  resolved?: boolean  // e.g. true
  resolvedAt?: string  // e.g. 2026-01-15T10:30:00Z
  resolvedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
}
```

**Response:** `AlertHistory[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-history/bulk/"
```

</details>

---

:::tab Find & Update

#### <Badge type="info" text="GET" /> Find AlertHistory by ID

```
GET /api/v1/alert-history/with-id/:id
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

**Response:** `AlertHistory`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/alert-history/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update AlertHistory

```
PUT /api/v1/alert-history/with-id/:id
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
  alertRuleId?: string
  triggeredValue?: Record<string, unknown>
  resolved?: boolean
  resolvedAt?: string
  resolvedBy?: string
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
  "http://localhost:3000/api/v1/alert-history/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update AlertHistory

```
PUT /api/v1/alert-history/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: AlertHistoryEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/alert-history/bulk/"
```

</details>

---

:::tab Delete

#### <Badge type="danger" text="DELETE" /> Delete AlertHistory

```
DELETE /api/v1/alert-history/with-id/:id
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
  "http://localhost:3000/api/v1/alert-history/with-id/:id"
```

</details>

---

::::


::::
