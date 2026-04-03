---
title: Events
---

# Events

**Table:** `analytics.events`

**Base path:** `/events`

## Related Tables

### Parent Tables

_Tables this table references via foreign keys._

| Parent Table | FK Column | References | Link |
|-------------|-----------|------------|------|
| `users` | `actor_id` | `events_actor_id_fkey` | [Users](./users) |


## Entity Relationship Diagram

erDiagram
    Events }o--|| Users : "FK"

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 3 | `event_type` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 4 | `source_schema` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 5 | `source_table` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 6 | `source_id` | `text` | `string` | `string` | NO | `''::text` | - | - |
| 7 | `actor_id` | `uuid` | `uuid.UUID` | `string` | YES | - | `FK` | → References `users` |
| 8 | `payload` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 9 | `severity` | `USER-DEFINED` | `AnalyticsEventSeverity` | `"debug" \| "info" \| "warning" \| "error" \| "critical"` | NO | `'info'::analytics.event_severity` | - | - |
| 10 | `processed` | `boolean` | `bool` | `boolean` | NO | `false` | - | - |
| 11 | `processed_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | - |
| 12 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)

## Foreign Keys & Relationships

| Column | References | Constraint |
|--------|-----------|------------|
| `actor_id` | `users` | `events_actor_id_fkey` |

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

> 📂 Source: [📄 `Events.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Events.go) · [📄 `Events.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Events.go) · [📄 `Events.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/Events.go)

### Structs

:::tabs

== Form

#### EventsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Events.go#:~:text=type%20EventsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `EventType` | `string` | `eventType` | NO |
| `SourceSchema` | `string` | `sourceSchema` | NO |
| `SourceTable` | `string` | `sourceTable` | NO |
| `SourceId` | `string` | `sourceId` | NO |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Payload` | `json.RawMessage` | `payload` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `Processed` | `bool` | `processed` | NO |
| `ProcessedAt` | `*time.Time` | `processedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### Events [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Events.go#:~:text=type%20Events%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `EventType` | `string` | `eventType` | NO |
| `SourceSchema` | `string` | `sourceSchema` | NO |
| `SourceTable` | `string` | `sourceTable` | NO |
| `SourceId` | `string` | `sourceId` | NO |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Payload` | `json.RawMessage` | `payload` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `Processed` | `bool` | `processed` | NO |
| `ProcessedAt` | `*time.Time` | `processedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### EventsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Events.go#:~:text=type%20EventsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `EventType` | `*string` | `eventType` | YES |
| `SourceSchema` | `*string` | `sourceSchema` | YES |
| `SourceTable` | `*string` | `sourceTable` | YES |
| `SourceId` | `*string` | `sourceId` | YES |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Payload` | `*json.RawMessage` | `payload` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `Processed` | `*bool` | `processed` | YES |
| `ProcessedAt` | `*time.Time` | `processedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### EventsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Events.go#:~:text=type%20EventsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `EventType` | `*string` | `eventType` | YES |
| `SourceSchema` | `*string` | `sourceSchema` | YES |
| `SourceTable` | `*string` | `sourceTable` | YES |
| `SourceId` | `*string` | `sourceId` | YES |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Payload` | `*json.RawMessage` | `payload` | YES |
| `Severity` | `*AnalyticsEventSeverity` | `severity` | YES |
| `Processed` | `*bool` | `processed` | YES |
| `ProcessedAt` | `*time.Time` | `processedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### EventsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Events.go#:~:text=type%20EventsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `EventType` | `string` | `eventType` | NO |
| `SourceSchema` | `string` | `sourceSchema` | NO |
| `SourceTable` | `string` | `sourceTable` | NO |
| `SourceId` | `string` | `sourceId` | NO |
| `ActorId` | `*uuid.UUID` | `actorId` | YES |
| `Payload` | `json.RawMessage` | `payload` | NO |
| `Severity` | `AnalyticsEventSeverity` | `severity` | NO |
| `Processed` | `bool` | `processed` | NO |
| `ProcessedAt` | `*time.Time` | `processedAt` | YES |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### EventsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Events.go#:~:text=type%20EventsBatchUpdate%20struct)

```go
type EventsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Events.go#:~:text=)%20CreateEvents() | `(EventsService) CreateEvents(data EventsForm) (EventsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Events.go#:~:text=)%20CreateEventsMultiple() | `(EventsService) CreateEventsMultiple(data []EventsForm) ([]EventsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Events.go#:~:text=)%20UpdateEvents() | `(EventsService) UpdateEvents(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Events.go#:~:text=)%20UpdateEventsMultiple() | `(EventsService) UpdateEventsMultiple(data []EventsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Events.go#:~:text=)%20DeleteEvents() | `(EventsService) DeleteEvents(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/events/` | Search with query params |
| `GET` | `/events/pagination` | Paginated listing |
| `POST` | `/events/` | Create single record |
| `POST` | `/events/bulk/` | Create multiple records |
| `PUT` | `/events/bulk/` | Batch update |
| `GET` | `/events/with-id/:id` | Get by ID |
| `PUT` | `/events/with-id/:id` | Update by ID |
| `DELETE` | `/events/with-id/:id` | Delete by ID |

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

export interface Events {
  id: string;
  name: string;
  eventType: string;
  sourceSchema: string;
  sourceTable: string;
  sourceId: string;
  actorId?: string;
  payload: Record<string, unknown>;
  severity: AnalyticsEventSeverity;
  processed: boolean;
  processedAt?: string;
  createdAt: string;
}

export interface EventsForm {
  name: string;
  eventType: string;
  sourceSchema: string;
  sourceTable: string;
  sourceId: string;
  actorId?: string;
  payload: Record<string, unknown>;
  severity: AnalyticsEventSeverity;
  processed: boolean;
  processedAt?: string;
  createdAt: string;
}

export interface EventsEdit {
  id: string;
  name: string;
  eventType: string;
  sourceSchema: string;
  sourceTable: string;
  sourceId: string;
  actorId?: string;
  payload: Record<string, unknown>;
  severity: AnalyticsEventSeverity;
  processed: boolean;
  processedAt?: string;
  createdAt: string;
}

export interface EventsPage {
  data: Events[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type EventsPathQuery = {
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

const EventsKeys = {
  all: ["events"] as const,
  lists: () => [...EventsKeys.all, "list"] as const,
  detail: (id: any) => [...EventsKeys.all, "detail", id] as const,
} as const;

export function useEventsList(query?: EventsPathQuery) {
  return useQuery({
    queryKey: [...EventsKeys.lists(), query],
    queryFn: () => fetch(`/events/pagination`, { method: "GET" }).then(r => r.json()) as Promise<EventsPage>,
  });
}

export function useEventsDetail(id: any) {
  return useQuery({
    queryKey: EventsKeys.detail(id),
    queryFn: () => fetch(`/events/with-id/:id`).then(r => r.json()) as Promise<Events>,
  });
}

export function useCreateEvents() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: EventsForm) =>
      fetch("/events/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: EventsKeys.lists() }),
  });
}

export function useUpdateEvents() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: EventsEdit }) =>
      fetch(`/events/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: EventsKeys.all }),
  });
}

export function useDeleteEvents() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/events/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: EventsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

const AnalyticsEventSeveritySchema = z.enum(["debug", "info", "warning", "error", "critical"]);

export const EventsFormSchema = z.object({
  name: z.string(),
  eventType: z.string(),
  sourceSchema: z.string(),
  sourceTable: z.string(),
  sourceId: z.string(),
  actorId: z.string().uuid().optional(),
  payload: z.record(z.unknown()),
  severity: AnalyticsEventSeveritySchema,
  processed: z.boolean(),
  processedAt: z.string().datetime().optional(),
  createdAt: z.string().datetime(),
});

export type EventsFormInput = z.infer<typeof EventsFormSchema>;

```

:::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './events.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Events

```
GET /api/v1/events/
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
| `eventType` | `string` | No | Filter by event_type |
| `sourceSchema` | `string` | No | Filter by source_schema |
| `sourceTable` | `string` | No | Filter by source_table |
| `sourceId` | `string` | No | Filter by source_id |
| `actorId` | `string (uuid)` | No | Filter by actor_id |
| `payload` | `string` | No | Filter by payload |
| `severity` | `string` | No | Filter by severity |
| `processed` | `boolean` | No | Filter by processed |
| `processedAt` | `string (date-time)` | No | Filter by processed_at |

**Response:** `Events[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/events/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Events (POST)

```
POST /api/v1/events/search
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

**Response:** `Events[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/events/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Events

```
GET /api/v1/events/pagination
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
| `eventType` | `string` | No | Filter by event_type |
| `sourceSchema` | `string` | No | Filter by source_schema |
| `sourceTable` | `string` | No | Filter by source_table |
| `sourceId` | `string` | No | Filter by source_id |
| `actorId` | `string (uuid)` | No | Filter by actor_id |
| `payload` | `string` | No | Filter by payload |
| `severity` | `string` | No | Filter by severity |
| `processed` | `boolean` | No | Filter by processed |
| `processedAt` | `string (date-time)` | No | Filter by processed_at |

**Response:** `PaginationResponse<Events>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/events/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Events (POST)

```
POST /api/v1/events/pagination
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

**Response:** `PaginationResponse<Events>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/events/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Events

```
POST /api/v1/events/
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
  eventType?: string  // e.g. example_event_type
  sourceSchema?: string  // e.g. example_source_schema
  sourceTable?: string  // e.g. example_source_table
  sourceId?: string  // e.g. example_source_id
  actorId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  payload?: Record<string, unknown>  // e.g. map[]
  severity?: "debug" | "info" | "warning" | "error" | "critical"  // e.g. debug
  processed?: boolean  // e.g. true
  processedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Events`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/events/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Events

```
POST /api/v1/events/bulk/
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
  eventType?: string  // e.g. example_event_type
  sourceSchema?: string  // e.g. example_source_schema
  sourceTable?: string  // e.g. example_source_table
  sourceId?: string  // e.g. example_source_id
  actorId?: string  // e.g. 550e8400-e29b-41d4-a716-446655440000
  payload?: Record<string, unknown>  // e.g. map[]
  severity?: "debug" | "info" | "warning" | "error" | "critical"  // e.g. debug
  processed?: boolean  // e.g. true
  processedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Events[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/events/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Events by ID

```
GET /api/v1/events/with-id/:id
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

**Response:** `Events`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/events/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Events

```
PUT /api/v1/events/with-id/:id
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
  eventType?: string
  sourceSchema?: string
  sourceTable?: string
  sourceId?: string
  actorId?: string
  payload?: Record<string, unknown>
  severity?: "debug" | "info" | "warning" | "error" | "critical"
  processed?: boolean
  processedAt?: string
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
  "http://localhost:3000/api/v1/events/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Events

```
PUT /api/v1/events/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: EventsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/events/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Events

```
DELETE /api/v1/events/with-id/:id
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
  "http://localhost:3000/api/v1/events/with-id/:id"
```

</details>

---

:::


::::
