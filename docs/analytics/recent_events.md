---
title: RecentEvents
---

# RecentEvents

**Table:** `analytics.recent_events`

::::tabs

=== FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | YES | - | - | - |
| 2 | `name` | `text` | `string` | `string` | YES | - | - | - |
| 3 | `event_type` | `text` | `string` | `string` | YES | - | - | - |
| 4 | `source_schema` | `text` | `string` | `string` | YES | - | - | - |
| 5 | `source_table` | `text` | `string` | `string` | YES | - | - | - |
| 6 | `actor_id` | `uuid` | `uuid.UUID` | `string` | YES | - | - | - |
| 7 | `severity` | `USER-DEFINED` | `AnalyticsEventSeverity` | `"debug" \| "info" \| "warning" \| "error" \| "critical"` | YES | - | - | - |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | Auto-filled from session |

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

> 📂 Source: [📄 `RecentEvents.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/RecentEvents.go) · [📄 `RecentEvents.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/RecentEvents.go) · [📄 `RecentEvents.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/RecentEvents.go)

### Structs

:::tabs

== Form

#### RecentEventsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/RecentEvents.go#:~:text=type%20RecentEventsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `EventType` | `*string` | `eventType` | YES |
| `SourceSchema` | `*string` | `sourceSchema` | YES |
| `SourceTable` | `*string` | `sourceTable` | YES |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Model

#### RecentEvents [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/RecentEvents.go#:~:text=type%20RecentEvents%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `EventType` | `*string` | `eventType` | YES |
| `SourceSchema` | `*string` | `sourceSchema` | YES |
| `SourceTable` | `*string` | `sourceTable` | YES |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Edit

#### RecentEventsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/RecentEvents.go#:~:text=type%20RecentEventsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `EventType` | `*string` | `eventType` | YES |
| `SourceSchema` | `*string` | `sourceSchema` | YES |
| `SourceTable` | `*string` | `sourceTable` | YES |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### RecentEventsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/RecentEvents.go#:~:text=type%20RecentEventsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `EventType` | `*string` | `eventType` | YES |
| `SourceSchema` | `*string` | `sourceSchema` | YES |
| `SourceTable` | `*string` | `sourceTable` | YES |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### RecentEventsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/RecentEvents.go#:~:text=type%20RecentEventsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `EventType` | `*string` | `eventType` | YES |
| `SourceSchema` | `*string` | `sourceSchema` | YES |
| `SourceTable` | `*string` | `sourceTable` | YES |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::

### Service & Endpoints

:::tabs

== Service Methods

_View — read-only, no write service methods._

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/recent-events/` | Search with query params |
| `GET` | `/recent-events/pagination` | Paginated listing |

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


=== Frontend

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

export interface RecentEvents {
  id?: string;
  name?: string;
  eventType?: string;
  sourceSchema?: string;
  sourceTable?: string;
  actorId?: string;
  severity?: AnalyticsEventSeverity;
  createdAt?: string;
}

export interface RecentEventsForm {
  id?: string;
  name?: string;
  eventType?: string;
  sourceSchema?: string;
  sourceTable?: string;
  actorId?: string;
  severity?: AnalyticsEventSeverity;
  createdAt?: string;
}

export interface RecentEventsEdit {
  id?: string;
  name?: string;
  eventType?: string;
  sourceSchema?: string;
  sourceTable?: string;
  actorId?: string;
  severity?: AnalyticsEventSeverity;
  createdAt?: string;
}

export interface RecentEventsPage {
  data: RecentEvents[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type RecentEventsPathQuery = {
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

const RecentEventsKeys = {
  all: ["recent_events"] as const,
  lists: () => [...RecentEventsKeys.all, "list"] as const,
  detail: (id: any) => [...RecentEventsKeys.all, "detail", id] as const,
} as const;

export function useRecentEventsList(query?: RecentEventsPathQuery) {
  return useQuery({
    queryKey: [...RecentEventsKeys.lists(), query],
    queryFn: () => fetch(`/recent-events/pagination`, { method: "GET" }).then(r => r.json()) as Promise<RecentEventsPage>,
  });
}

export function useRecentEventsDetail(id: any) {
  return useQuery({
    queryKey: RecentEventsKeys.detail(id),
    queryFn: () => fetch(`/recent-events/with-id/:id`).then(r => r.json()) as Promise<RecentEvents>,
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

const AnalyticsEventSeveritySchema = z.enum(["debug", "info", "warning", "error", "critical"]);

export const RecentEventsFormSchema = z.object({
  id: z.string().uuid().optional(),
  name: z.string().optional(),
  eventType: z.string().optional(),
  sourceSchema: z.string().optional(),
  sourceTable: z.string().optional(),
  actorId: z.string().uuid().optional(),
  severity: AnalyticsEventSeveritySchema.optional(),
  createdAt: z.string().datetime().optional(),
});

export type RecentEventsFormInput = z.infer<typeof RecentEventsFormSchema>;

```

:::


=== API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './recent_events.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search RecentEvents

```
GET /api/v1/recent-events/
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
| `joins` | `string` | No | Comma-separated relations to join |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `eventType` | `string` | No | Filter by event_type |
| `sourceSchema` | `string` | No | Filter by source_schema |
| `sourceTable` | `string` | No | Filter by source_table |
| `actorId` | `string (uuid)` | No | Filter by actor_id |
| `severity` | `string` | No | Filter by severity |

**Response:** `RecentEvents[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/recent-events/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search RecentEvents (POST)

```
POST /api/v1/recent-events/search
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

**Response:** `RecentEvents[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/recent-events/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate RecentEvents

```
GET /api/v1/recent-events/pagination
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
| `joins` | `string` | No | Comma-separated relations to join |
| `id` | `string (uuid)` | No | Filter by id |
| `name` | `string` | No | Filter by name |
| `eventType` | `string` | No | Filter by event_type |
| `sourceSchema` | `string` | No | Filter by source_schema |
| `sourceTable` | `string` | No | Filter by source_table |
| `actorId` | `string (uuid)` | No | Filter by actor_id |
| `severity` | `string` | No | Filter by severity |

**Response:** `PaginationResponse<RecentEvents>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/recent-events/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate RecentEvents (POST)

```
POST /api/v1/recent-events/pagination
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

**Response:** `PaginationResponse<RecentEvents>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/recent-events/pagination"
```

</details>

---

:::


::::
