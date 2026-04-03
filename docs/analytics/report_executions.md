---
title: ReportExecutions
---

# ReportExecutions

**Table:** `analytics.report_executions`

**Base path:** `/report-executions`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `reports` | `report_id` | `report_executions_report_id_fkey` | [Reports](./reports) |
| `users` | `executed_by` | `report_executions_executed_by_fkey` | [Users](./users) |


## Entity Relationship Diagram

erDiagram
    ReportExecutions }o--|| Reports : "FK"
    ReportExecutions }o--|| Users : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `report_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `reports` |
| 4 | `status` | `text` | `string` | `string` | NO | `'pending'::text` | - | - |
| 5 | `result_data` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | YES | - | - | - |
| 6 | `row_count` | `integer` | `int` | `number` | NO | `0` | - | - |
| 7 | `duration_ms` | `integer` | `int` | `number` | NO | `0` | - | - |
| 8 | `error_message` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 9 | `executed_by` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `users` |
| 10 | `started_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 11 | `completed_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 12 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `report_id` | `reports` | `report_executions_report_id_fkey` |
| `executed_by` | `users` | `report_executions_executed_by_fkey` |


## Go Generated Code

> 📂 Source: [📄 `ReportExecutions.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/ReportExecutions.go) · [📄 `ReportExecutions.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/ReportExecutions.go) · [📄 `ReportExecutions.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/ReportExecutions.go)

### Structs

:::tabs

== Form

#### ReportExecutionsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/ReportExecutions.go#:~:text=type%20ReportExecutionsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `ReportId` | `uuid.UUID` | `reportId` | NO |
| `Status` | `string` | `status` | NO |
| `ResultData` | `*json.RawMessage` | `resultData` | YES |
| `RowCount` | `int` | `rowCount` | NO |
| `DurationMs` | `int` | `durationMs` | NO |
| `ErrorMessage` | `string` | `errorMessage` | NO |
| `ExecutedBy` | `*uuid.UUID` | `executedBy` | YES |
| `StartedAt` | `*time.Time` | `startedAt` | YES |
| `CompletedAt` | `*time.Time` | `completedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### ReportExecutions [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/ReportExecutions.go#:~:text=type%20ReportExecutions%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ReportId` | `uuid.UUID` | `reportId` | NO |
| `Status` | `string` | `status` | NO |
| `ResultData` | `*json.RawMessage` | `resultData` | YES |
| `RowCount` | `int` | `rowCount` | NO |
| `DurationMs` | `int` | `durationMs` | NO |
| `ErrorMessage` | `string` | `errorMessage` | NO |
| `ExecutedBy` | `*uuid.UUID` | `executedBy` | YES |
| `StartedAt` | `*time.Time` | `startedAt` | YES |
| `CompletedAt` | `*time.Time` | `completedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### ReportExecutionsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/ReportExecutions.go#:~:text=type%20ReportExecutionsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ReportId` | `*uuid.UUID` | `reportId` | YES |
| `Status` | `*string` | `status` | YES |
| `ResultData` | `*json.RawMessage` | `resultData` | YES |
| `RowCount` | `*int` | `rowCount` | YES |
| `DurationMs` | `*int` | `durationMs` | YES |
| `ErrorMessage` | `*string` | `errorMessage` | YES |
| `ExecutedBy` | `*uuid.UUID` | `executedBy` | YES |
| `StartedAt` | `*time.Time` | `startedAt` | YES |
| `CompletedAt` | `*time.Time` | `completedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### ReportExecutionsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/ReportExecutions.go#:~:text=type%20ReportExecutionsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `ReportId` | `*uuid.UUID` | `reportId` | YES |
| `Status` | `*string` | `status` | YES |
| `ResultData` | `*json.RawMessage` | `resultData` | YES |
| `RowCount` | `*int` | `rowCount` | YES |
| `DurationMs` | `*int` | `durationMs` | YES |
| `ErrorMessage` | `*string` | `errorMessage` | YES |
| `ExecutedBy` | `*uuid.UUID` | `executedBy` | YES |
| `StartedAt` | `*time.Time` | `startedAt` | YES |
| `CompletedAt` | `*time.Time` | `completedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### ReportExecutionsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/ReportExecutions.go#:~:text=type%20ReportExecutionsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `ReportId` | `uuid.UUID` | `reportId` | NO |
| `Status` | `string` | `status` | NO |
| `ResultData` | `*json.RawMessage` | `resultData` | YES |
| `RowCount` | `int` | `rowCount` | NO |
| `DurationMs` | `int` | `durationMs` | NO |
| `ErrorMessage` | `string` | `errorMessage` | NO |
| `ExecutedBy` | `*uuid.UUID` | `executedBy` | YES |
| `StartedAt` | `*time.Time` | `startedAt` | YES |
| `CompletedAt` | `*time.Time` | `completedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### ReportExecutionsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/ReportExecutions.go#:~:text=type%20ReportExecutionsBatchUpdate%20struct)

```go
type ReportExecutionsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/ReportExecutions.go#:~:text=)%20CreateReportExecutions() | `(ReportExecutionsService) CreateReportExecutions(data ReportExecutionsForm) (ReportExecutionsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/ReportExecutions.go#:~:text=)%20CreateReportExecutionsMultiple() | `(ReportExecutionsService) CreateReportExecutionsMultiple(data []ReportExecutionsForm) ([]ReportExecutionsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/ReportExecutions.go#:~:text=)%20UpdateReportExecutions() | `(ReportExecutionsService) UpdateReportExecutions(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/ReportExecutions.go#:~:text=)%20UpdateReportExecutionsMultiple() | `(ReportExecutionsService) UpdateReportExecutionsMultiple(data []ReportExecutionsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/ReportExecutions.go#:~:text=)%20DeleteReportExecutions() | `(ReportExecutionsService) DeleteReportExecutions(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/report-executions/` | Search with query params |
| `GET` | `/report-executions/pagination` | Paginated listing |
| `POST` | `/report-executions/` | Create single record |
| `POST` | `/report-executions/bulk/` | Create multiple records |
| `PUT` | `/report-executions/bulk/` | Batch update |
| `GET` | `/report-executions/with-id/:id` | Get by ID |
| `PUT` | `/report-executions/with-id/:id` | Update by ID |
| `DELETE` | `/report-executions/with-id/:id` | Delete by ID |

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
export interface ReportExecutions {
  id: string;
  name: string;
  reportId: string;
  status: string;
  resultData?: Record<string, unknown>;
  rowCount: number;
  durationMs: number;
  errorMessage: string;
  executedBy?: string;
  startedAt?: string;
  completedAt?: string;
  createdAt: string;
}

export interface ReportExecutionsForm {
  name: string;
  reportId: string;
  status: string;
  resultData?: Record<string, unknown>;
  rowCount: number;
  durationMs: number;
  errorMessage: string;
  executedBy?: string;
  startedAt?: string;
  completedAt?: string;
  createdAt: string;
}

export interface ReportExecutionsEdit {
  id: string;
  name: string;
  reportId: string;
  status: string;
  resultData?: Record<string, unknown>;
  rowCount: number;
  durationMs: number;
  errorMessage: string;
  executedBy?: string;
  startedAt?: string;
  completedAt?: string;
  createdAt: string;
}

export interface ReportExecutionsPage {
  data: ReportExecutions[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type ReportExecutionsPathQuery = {
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

const ReportExecutionsKeys = {
  all: ["report_executions"] as const,
  lists: () => [...ReportExecutionsKeys.all, "list"] as const,
  detail: (id: any) => [...ReportExecutionsKeys.all, "detail", id] as const,
} as const;

export function useReportExecutionsList(query?: ReportExecutionsPathQuery) {
  return useQuery({
    queryKey: [...ReportExecutionsKeys.lists(), query],
    queryFn: () => fetch(`/report-executions/pagination`, { method: "GET" }).then(r => r.json()) as Promise<ReportExecutionsPage>,
  });
}

export function useReportExecutionsDetail(id: any) {
  return useQuery({
    queryKey: ReportExecutionsKeys.detail(id),
    queryFn: () => fetch(`/report-executions/with-id/:id`).then(r => r.json()) as Promise<ReportExecutions>,
  });
}

export function useCreateReportExecutions() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: ReportExecutionsForm) =>
      fetch("/report-executions/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ReportExecutionsKeys.lists() }),
  });
}

export function useUpdateReportExecutions() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: ReportExecutionsEdit }) =>
      fetch(`/report-executions/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ReportExecutionsKeys.all }),
  });
}

export function useDeleteReportExecutions() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/report-executions/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ReportExecutionsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const ReportExecutionsFormSchema = z.object({
  name: z.string(),
  reportId: z.string().uuid(),
  status: z.string(),
  resultData: z.record(z.unknown()).optional(),
  rowCount: z.number().int(),
  durationMs: z.number().int(),
  errorMessage: z.string(),
  executedBy: z.string().uuid().optional(),
  startedAt: z.string().datetime().optional(),
  completedAt: z.string().datetime().optional(),
  createdAt: z.string().datetime(),
});

export type ReportExecutionsFormInput = z.infer<typeof ReportExecutionsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './report_executions.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search ReportExecutions

```
GET /api/v1/report-executions/
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
| `preloads` | `string` | No | Available: ReportIdDetail, ReportIdDetail.ReportExecutionsList, ReportIdDetail.ReportExecutionsList.ReportIdDetail |
| `joins` | `string` | No | Available: Reports, Reports.Users, Reports.Organizations, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `reportId` | `string (uuid)` | No | Filter by report_id |
| `status` | `string` | No | Filter by status |
| `resultData` | `string` | No | Filter by result_data |
| `rowCount` | `integer` | No | Filter by row_count |
| `durationMs` | `integer` | No | Filter by duration_ms |
| `errorMessage` | `string` | No | Filter by error_message |
| `executedBy` | `string (uuid)` | No | Filter by executed_by |
| `startedAt` | `string (date-time)` | No | Filter by started_at |
| `completedAt` | `string (date-time)` | No | Filter by completed_at |

**Response:** `ReportExecutions[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/report-executions/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search ReportExecutions (POST)

```
POST /api/v1/report-executions/search
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

**Response:** `ReportExecutions[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/report-executions/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate ReportExecutions

```
GET /api/v1/report-executions/pagination
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
| `preloads` | `string` | No | Available: ReportIdDetail, ReportIdDetail.ReportExecutionsList, ReportIdDetail.ReportExecutionsList.ReportIdDetail |
| `joins` | `string` | No | Available: Reports, Reports.Users, Reports.Organizations, Users |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `reportId` | `string (uuid)` | No | Filter by report_id |
| `status` | `string` | No | Filter by status |
| `resultData` | `string` | No | Filter by result_data |
| `rowCount` | `integer` | No | Filter by row_count |
| `durationMs` | `integer` | No | Filter by duration_ms |
| `errorMessage` | `string` | No | Filter by error_message |
| `executedBy` | `string (uuid)` | No | Filter by executed_by |
| `startedAt` | `string (date-time)` | No | Filter by started_at |
| `completedAt` | `string (date-time)` | No | Filter by completed_at |

**Response:** `PaginationResponse<ReportExecutions>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/report-executions/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate ReportExecutions (POST)

```
POST /api/v1/report-executions/pagination
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

**Response:** `PaginationResponse<ReportExecutions>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/report-executions/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create ReportExecutions

```
POST /api/v1/report-executions/
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
  reportId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  status?: string  // e.g. example_status
  resultData?: Record<string, unknown>  // e.g. map[]
  rowCount?: number  // e.g. 1
  durationMs?: number  // e.g. 1
  errorMessage?: string  // e.g. example_error_message
  executedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  startedAt?: string  // e.g. 2026-01-15T10:30:00Z
  completedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `ReportExecutions`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/report-executions/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create ReportExecutions

```
POST /api/v1/report-executions/bulk/
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
  reportId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  status?: string  // e.g. example_status
  resultData?: Record<string, unknown>  // e.g. map[]
  rowCount?: number  // e.g. 1
  durationMs?: number  // e.g. 1
  errorMessage?: string  // e.g. example_error_message
  executedBy?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  startedAt?: string  // e.g. 2026-01-15T10:30:00Z
  completedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `ReportExecutions[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/report-executions/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find ReportExecutions by ID

```
GET /api/v1/report-executions/with-id/:id
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

**Response:** `ReportExecutions`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/report-executions/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update ReportExecutions

```
PUT /api/v1/report-executions/with-id/:id
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
  reportId?: string
  status?: string
  resultData?: Record<string, unknown>
  rowCount?: number
  durationMs?: number
  errorMessage?: string
  executedBy?: string
  startedAt?: string
  completedAt?: string
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
  "http://localhost:3000/api/v1/report-executions/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update ReportExecutions

```
PUT /api/v1/report-executions/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: ReportExecutionsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/report-executions/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete ReportExecutions

```
DELETE /api/v1/report-executions/with-id/:id
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
  "http://localhost:3000/api/v1/report-executions/with-id/:id"
```

</details>

---

:::


::::
