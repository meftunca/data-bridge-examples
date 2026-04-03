---
title: Reports
---

# Reports

**Table:** `analytics.reports`

**Base path:** `/reports`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `users` | `owner_id` | `reports_owner_id_fkey` | [Users](./users) |
| `organizations` | `organization_id` | `reports_organization_id_fkey` | [Organizations](./organizations) |

### Child Tables

_Tables that reference this table via foreign keys._

| Child Table | FK Column | References | Link |
|------------|-----------|------------|------|
| `report_executions` | `report_id` | `report_executions_report_id_fkey` | [ReportExecutions](./report_executions) |


## Entity Relationship Diagram

erDiagram
    Reports }o--|| Users : "FK"
    Reports }o--|| Organizations : "FK"
    Reports ||--o{ ReportExecutions : "ref"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `description` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 4 | `report_type` | `USER-DEFINED` | `AnalyticsReportType` | `"daily" \| "weekly" \| "monthly" \| "quarterly" \| "annual" \| "custom"` | NO | `'monthly'::analytics.report_type` | - | - |
| 5 | `owner_id` | `uuid` | `uuid.UUID` | `string` | NO | - | `FK` | → References `users` |
| 6 | `organization_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `organizations` |
| 7 | `query_config` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 8 | `schedule` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 9 | `is_active` | `boolean` | `bool` | `boolean` | NO | `true` | - | - |
| 10 | `last_run_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 11 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |
| 12 | `updated_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `owner_id` | `users` | `reports_owner_id_fkey` |
| `organization_id` | `organizations` | `reports_organization_id_fkey` |

## Enum Types

### ReportType

| Value | Go Constant |
|-------|-------------|
| `daily` | `AnalyticsReportTypeDaily` |
| `weekly` | `AnalyticsReportTypeWeekly` |
| `monthly` | `AnalyticsReportTypeMonthly` |
| `quarterly` | `AnalyticsReportTypeQuarterly` |
| `annual` | `AnalyticsReportTypeAnnual` |
| `custom` | `AnalyticsReportTypeCustom` |


## Go Generated Code

> 📂 Source: [📄 `Reports.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Reports.go) · [📄 `Reports.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Reports.go) · [📄 `Reports.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/Reports.go)

### Structs

:::tabs

== Form

#### ReportsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Reports.go#:~:text=type%20ReportsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `Description` | `string` | `description` | NO |
| `ReportType` | `AnalyticsReportType` | `reportType` | NO |
| `OwnerId` | `uuid.UUID` | `ownerId` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `QueryConfig` | `json.RawMessage` | `queryConfig` | NO |
| `Schedule` | `json.RawMessage` | `schedule` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `LastRunAt` | `*time.Time` | `lastRunAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Model

#### Reports [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Reports.go#:~:text=type%20Reports%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Description` | `string` | `description` | NO |
| `ReportType` | `AnalyticsReportType` | `reportType` | NO |
| `OwnerId` | `uuid.UUID` | `ownerId` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `QueryConfig` | `json.RawMessage` | `queryConfig` | NO |
| `Schedule` | `json.RawMessage` | `schedule` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `LastRunAt` | `*time.Time` | `lastRunAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== Edit

#### ReportsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Reports.go#:~:text=type%20ReportsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Description` | `*string` | `description` | YES |
| `ReportType` | `*AnalyticsReportType` | `reportType` | YES |
| `OwnerId` | `*uuid.UUID` | `ownerId` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `QueryConfig` | `*json.RawMessage` | `queryConfig` | YES |
| `Schedule` | `*json.RawMessage` | `schedule` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `LastRunAt` | `*time.Time` | `lastRunAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Filter

#### ReportsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Reports.go#:~:text=type%20ReportsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `Description` | `*string` | `description` | YES |
| `ReportType` | `*AnalyticsReportType` | `reportType` | YES |
| `OwnerId` | `*uuid.UUID` | `ownerId` | YES |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `QueryConfig` | `*json.RawMessage` | `queryConfig` | YES |
| `Schedule` | `*json.RawMessage` | `schedule` | YES |
| `IsActive` | `*bool` | `isActive` | YES |
| `LastRunAt` | `*time.Time` | `lastRunAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |
| `UpdatedAt` | `*time.Time` | `updatedAt` | YES |

== Page

#### ReportsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Reports.go#:~:text=type%20ReportsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `Description` | `string` | `description` | NO |
| `ReportType` | `AnalyticsReportType` | `reportType` | NO |
| `OwnerId` | `uuid.UUID` | `ownerId` | NO |
| `OrganizationId` | `*uuid.UUID` | `organizationId` | YES |
| `QueryConfig` | `json.RawMessage` | `queryConfig` | NO |
| `Schedule` | `json.RawMessage` | `schedule` | NO |
| `IsActive` | `bool` | `isActive` | NO |
| `LastRunAt` | `*time.Time` | `lastRunAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |
| `UpdatedAt` | `time.Time` | `updatedAt` | NO |

== BatchUpdate

#### ReportsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Reports.go#:~:text=type%20ReportsBatchUpdate%20struct)

```go
type ReportsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Reports.go#:~:text=)%20CreateReports() | `(ReportsService) CreateReports(data ReportsForm) (ReportsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Reports.go#:~:text=)%20CreateReportsMultiple() | `(ReportsService) CreateReportsMultiple(data []ReportsForm) ([]ReportsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Reports.go#:~:text=)%20UpdateReports() | `(ReportsService) UpdateReports(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Reports.go#:~:text=)%20UpdateReportsMultiple() | `(ReportsService) UpdateReportsMultiple(data []ReportsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Reports.go#:~:text=)%20DeleteReports() | `(ReportsService) DeleteReports(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/reports/` | Search with query params |
| `GET` | `/reports/pagination` | Paginated listing |
| `POST` | `/reports/` | Create single record |
| `POST` | `/reports/bulk/` | Create multiple records |
| `PUT` | `/reports/bulk/` | Batch update |
| `GET` | `/reports/with-id/:id` | Get by ID |
| `PUT` | `/reports/with-id/:id` | Update by ID |
| `DELETE` | `/reports/with-id/:id` | Delete by ID |

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
export type AnalyticsReportType =
  | "daily"
  | "weekly"
  | "monthly"
  | "quarterly"
  | "annual"
  | "custom"

export const AnalyticsReportTypeValues = ["daily", "weekly", "monthly", "quarterly", "annual", "custom"] as const;

export interface Reports {
  id: string;
  name: string;
  description: string;
  reportType: AnalyticsReportType;
  ownerId: string;
  organizationId?: string;
  queryConfig: Record<string, unknown>;
  schedule: Record<string, unknown>;
  isActive: boolean;
  lastRunAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface ReportsForm {
  name: string;
  description: string;
  reportType: AnalyticsReportType;
  ownerId: string;
  organizationId?: string;
  queryConfig: Record<string, unknown>;
  schedule: Record<string, unknown>;
  isActive: boolean;
  lastRunAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface ReportsEdit {
  id: string;
  name: string;
  description: string;
  reportType: AnalyticsReportType;
  ownerId: string;
  organizationId?: string;
  queryConfig: Record<string, unknown>;
  schedule: Record<string, unknown>;
  isActive: boolean;
  lastRunAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface ReportsPage {
  data: Reports[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type ReportsPathQuery = {
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

const ReportsKeys = {
  all: ["reports"] as const,
  lists: () => [...ReportsKeys.all, "list"] as const,
  detail: (id: any) => [...ReportsKeys.all, "detail", id] as const,
} as const;

export function useReportsList(query?: ReportsPathQuery) {
  return useQuery({
    queryKey: [...ReportsKeys.lists(), query],
    queryFn: () => fetch(`/reports/pagination`, { method: "GET" }).then(r => r.json()) as Promise<ReportsPage>,
  });
}

export function useReportsDetail(id: any) {
  return useQuery({
    queryKey: ReportsKeys.detail(id),
    queryFn: () => fetch(`/reports/with-id/:id`).then(r => r.json()) as Promise<Reports>,
  });
}

export function useCreateReports() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: ReportsForm) =>
      fetch("/reports/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ReportsKeys.lists() }),
  });
}

export function useUpdateReports() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: ReportsEdit }) =>
      fetch(`/reports/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ReportsKeys.all }),
  });
}

export function useDeleteReports() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/reports/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: ReportsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

const AnalyticsReportTypeSchema = z.enum(["daily", "weekly", "monthly", "quarterly", "annual", "custom"]);

export const ReportsFormSchema = z.object({
  name: z.string(),
  description: z.string(),
  reportType: AnalyticsReportTypeSchema,
  ownerId: z.string().uuid(),
  organizationId: z.string().uuid().optional(),
  queryConfig: z.record(z.unknown()),
  schedule: z.record(z.unknown()),
  isActive: z.boolean(),
  lastRunAt: z.string().datetime().optional(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type ReportsFormInput = z.infer<typeof ReportsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './reports.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Reports

```
GET /api/v1/reports/
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
| `preloads` | `string` | No | Available: ReportExecutionsList, ReportExecutionsList.ReportIdDetail, ReportExecutionsList.ReportIdDetail.ReportExecutionsList |
| `joins` | `string` | No | Available: Users, Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `description` | `string` | No | Filter by description |
| `reportType` | `string` | No | Filter by report_type |
| `ownerId` | `string (uuid)` | No | Filter by owner_id |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `queryConfig` | `string` | No | Filter by query_config |
| `schedule` | `string` | No | Filter by schedule |
| `isActive` | `boolean` | No | Filter by is_active |
| `lastRunAt` | `string (date-time)` | No | Filter by last_run_at |

**Response:** `Reports[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/reports/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Reports (POST)

```
POST /api/v1/reports/search
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

**Response:** `Reports[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/reports/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Reports

```
GET /api/v1/reports/pagination
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
| `preloads` | `string` | No | Available: ReportExecutionsList, ReportExecutionsList.ReportIdDetail, ReportExecutionsList.ReportIdDetail.ReportExecutionsList |
| `joins` | `string` | No | Available: Users, Organizations |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `description` | `string` | No | Filter by description |
| `reportType` | `string` | No | Filter by report_type |
| `ownerId` | `string (uuid)` | No | Filter by owner_id |
| `organizationId` | `string (uuid)` | No | Filter by organization_id |
| `queryConfig` | `string` | No | Filter by query_config |
| `schedule` | `string` | No | Filter by schedule |
| `isActive` | `boolean` | No | Filter by is_active |
| `lastRunAt` | `string (date-time)` | No | Filter by last_run_at |

**Response:** `PaginationResponse<Reports>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/reports/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Reports (POST)

```
POST /api/v1/reports/pagination
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

**Response:** `PaginationResponse<Reports>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/reports/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Reports

```
POST /api/v1/reports/
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
  reportType?: "daily" | "weekly" | "monthly" | "quarterly" | "annual" | "custom"  // e.g. daily
  ownerId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  queryConfig?: Record<string, unknown>  // e.g. map[]
  schedule?: Record<string, unknown>  // e.g. map[]
  isActive?: boolean  // e.g. true
  lastRunAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Reports`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/reports/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Reports

```
POST /api/v1/reports/bulk/
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
  reportType?: "daily" | "weekly" | "monthly" | "quarterly" | "annual" | "custom"  // e.g. daily
  ownerId: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  organizationId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  queryConfig?: Record<string, unknown>  // e.g. map[]
  schedule?: Record<string, unknown>  // e.g. map[]
  isActive?: boolean  // e.g. true
  lastRunAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Reports[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/reports/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Reports by ID

```
GET /api/v1/reports/with-id/:id
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

**Response:** `Reports`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/reports/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Reports

```
PUT /api/v1/reports/with-id/:id
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
  reportType?: "daily" | "weekly" | "monthly" | "quarterly" | "annual" | "custom"
  ownerId?: string
  organizationId?: string
  queryConfig?: Record<string, unknown>
  schedule?: Record<string, unknown>
  isActive?: boolean
  lastRunAt?: string
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
  "http://localhost:3000/api/v1/reports/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Reports

```
PUT /api/v1/reports/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: ReportsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/reports/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Reports

```
DELETE /api/v1/reports/with-id/:id
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
  "http://localhost:3000/api/v1/reports/with-id/:id"
```

</details>

---

:::


::::
