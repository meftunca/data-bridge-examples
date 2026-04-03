---
title: Metrics
---

# Metrics

**Table:** `analytics.metrics`

**Base path:** `/metrics`

::::tabs

=== FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | NO | `gen_random_uuid()` | `PK` | Primary key |
| 2 | `name` | `text` | `string` | `string` | NO | - | - | - |
| 3 | `metric_key` | `text` | `string` | `string` | NO | - | - | - |
| 4 | `value` | `numeric` | `float64` | `number` | NO | `0` | - | - |
| 5 | `dimensions` | `jsonb` | `json.RawMessage` | `Record<string, unknown>` | NO | `'{}'::jsonb` | - | - |
| 6 | `recorded_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | - |
| 7 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | NO | `now()` | - | Auto-filled from session |

## Primary Keys

- `id` (`uuid`)


## Go Generated Code

> 📂 Source: [📄 `Metrics.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Metrics.go) · [📄 `Metrics.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Metrics.go) · [📄 `Metrics.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/Metrics.go)

### Structs

:::tabs

== Form

#### MetricsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Metrics.go#:~:text=type%20MetricsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Name` | `string` | `name` | NO |
| `MetricKey` | `string` | `metricKey` | NO |
| `Value` | `float64` | `value` | NO |
| `Dimensions` | `json.RawMessage` | `dimensions` | NO |
| `RecordedAt` | `time.Time` | `recordedAt` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Model

#### Metrics [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Metrics.go#:~:text=type%20Metrics%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `MetricKey` | `string` | `metricKey` | NO |
| `Value` | `float64` | `value` | NO |
| `Dimensions` | `json.RawMessage` | `dimensions` | NO |
| `RecordedAt` | `time.Time` | `recordedAt` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== Edit

#### MetricsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Metrics.go#:~:text=type%20MetricsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `MetricKey` | `*string` | `metricKey` | YES |
| `Value` | `*float64` | `value` | YES |
| `Dimensions` | `*json.RawMessage` | `dimensions` | YES |
| `RecordedAt` | `*time.Time` | `recordedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Filter

#### MetricsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Metrics.go#:~:text=type%20MetricsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `MetricKey` | `*string` | `metricKey` | YES |
| `Value` | `*float64` | `value` | YES |
| `Dimensions` | `*json.RawMessage` | `dimensions` | YES |
| `RecordedAt` | `*time.Time` | `recordedAt` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

== Page

#### MetricsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Metrics.go#:~:text=type%20MetricsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `uuid.UUID` | `id` | NO |
| `Name` | `string` | `name` | NO |
| `MetricKey` | `string` | `metricKey` | NO |
| `Value` | `float64` | `value` | NO |
| `Dimensions` | `json.RawMessage` | `dimensions` | NO |
| `RecordedAt` | `time.Time` | `recordedAt` | NO |
| `CreatedAt` | `time.Time` | `createdAt` | NO |

== BatchUpdate

#### MetricsBatchUpdate [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/Metrics.go#:~:text=type%20MetricsBatchUpdate%20struct)

```go
type MetricsBatchUpdate struct {
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
| [Create](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Metrics.go#:~:text=%29%20CreateMetrics%28%29) | `(MetricsService) CreateMetrics(data MetricsForm) (MetricsForm, error)` |
| [Create Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Metrics.go#:~:text=%29%20CreateMetricsMultiple%28%29) | `(MetricsService) CreateMetricsMultiple(data []MetricsForm) ([]MetricsForm, error)` |
| [Update](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Metrics.go#:~:text=%29%20UpdateMetrics%28%29) | `(MetricsService) UpdateMetrics(id uuid.UUID, data interface{}) error` |
| [Update Multiple](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Metrics.go#:~:text=%29%20UpdateMetricsMultiple%28%29) | `(MetricsService) UpdateMetricsMultiple(data []MetricsBatchUpdate) error` |
| [Delete](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/Metrics.go#:~:text=%29%20DeleteMetrics%28%29) | `(MetricsService) DeleteMetrics(id uuid.UUID) error` |

== Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/metrics/` | Search with query params |
| `GET` | `/metrics/pagination` | Paginated listing |
| `POST` | `/metrics/` | Create single record |
| `POST` | `/metrics/bulk/` | Create multiple records |
| `PUT` | `/metrics/bulk/` | Batch update |
| `GET` | `/metrics/with-id/:id` | Get by ID |
| `PUT` | `/metrics/with-id/:id` | Update by ID |
| `DELETE` | `/metrics/with-id/:id` | Delete by ID |

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
export interface Metrics {
  id: string;
  name: string;
  metricKey: string;
  value: number;
  dimensions: Record<string, unknown>;
  recordedAt: string;
  createdAt: string;
}

export interface MetricsForm {
  name: string;
  metricKey: string;
  value: number;
  dimensions: Record<string, unknown>;
  recordedAt: string;
  createdAt: string;
}

export interface MetricsEdit {
  id: string;
  name: string;
  metricKey: string;
  value: number;
  dimensions: Record<string, unknown>;
  recordedAt: string;
  createdAt: string;
}

export interface MetricsPage {
  data: Metrics[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type MetricsPathQuery = {
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

const MetricsKeys = {
  all: ["metrics"] as const,
  lists: () => [...MetricsKeys.all, "list"] as const,
  detail: (id: any) => [...MetricsKeys.all, "detail", id] as const,
} as const;

export function useMetricsList(query?: MetricsPathQuery) {
  return useQuery({
    queryKey: [...MetricsKeys.lists(), query],
    queryFn: () => fetch(`/metrics/pagination`, { method: "GET" }).then(r => r.json()) as Promise<MetricsPage>,
  });
}

export function useMetricsDetail(id: any) {
  return useQuery({
    queryKey: MetricsKeys.detail(id),
    queryFn: () => fetch(`/metrics/with-id/:id`).then(r => r.json()) as Promise<Metrics>,
  });
}

export function useCreateMetrics() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: MetricsForm) =>
      fetch("/metrics/", { method: "POST", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: MetricsKeys.lists() }),
  });
}

export function useUpdateMetrics() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: any: any; data: MetricsEdit }) =>
      fetch(`/metrics/with-id/:id`, { method: "PUT", body: JSON.stringify(data) }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: MetricsKeys.all }),
  });
}

export function useDeleteMetrics() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: any) =>
      fetch(`/metrics/with-id/:id`, { method: "DELETE" }).then(r => r.json()),
    onSuccess: () => qc.invalidateQueries({ queryKey: MetricsKeys.all }),
  });
}

```

== Zod Validation

```typescript
import { z } from "zod";

export const MetricsFormSchema = z.object({
  name: z.string(),
  metricKey: z.string(),
  value: z.number(),
  dimensions: z.record(z.unknown()),
  recordedAt: z.string().datetime(),
  createdAt: z.string().datetime(),
});

export type MetricsFormInput = z.infer<typeof MetricsFormSchema>;

```

:::


=== API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './metrics.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

:::tabs

== Search

#### <Badge type="info" text="GET" /> Search Metrics

```
GET /api/v1/metrics/
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
| `metricKey` | `string` | No | Filter by metric_key |
| `value` | `number` | No | Filter by value |
| `dimensions` | `string` | No | Filter by dimensions |
| `recordedAt` | `string (date-time)` | No | Filter by recorded_at |

**Response:** `Metrics[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/metrics/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search Metrics (POST)

```
POST /api/v1/metrics/search
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

**Response:** `Metrics[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/metrics/search"
```

</details>

---

== Pagination

#### <Badge type="info" text="GET" /> Paginate Metrics

```
GET /api/v1/metrics/pagination
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
| `metricKey` | `string` | No | Filter by metric_key |
| `value` | `number` | No | Filter by value |
| `dimensions` | `string` | No | Filter by dimensions |
| `recordedAt` | `string (date-time)` | No | Filter by recorded_at |

**Response:** `PaginationResponse<Metrics>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/metrics/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate Metrics (POST)

```
POST /api/v1/metrics/pagination
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

**Response:** `PaginationResponse<Metrics>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/metrics/pagination"
```

</details>

---

== Create

#### <Badge type="tip" text="POST" /> Create Metrics

```
POST /api/v1/metrics/
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
  metricKey: string  // e.g. example_metric_key
  value?: number  // e.g. 99.99
  dimensions?: Record<string, unknown>  // e.g. map[]
  recordedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Metrics`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/metrics/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Bulk Create Metrics

```
POST /api/v1/metrics/bulk/
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
  metricKey: string  // e.g. example_metric_key
  value?: number  // e.g. 99.99
  dimensions?: Record<string, unknown>  // e.g. map[]
  recordedAt?: string  // e.g. 2026-01-15T10:30:00Z
}
```

**Response:** `Metrics[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/metrics/bulk/"
```

</details>

---

== Find & Update

#### <Badge type="info" text="GET" /> Find Metrics by ID

```
GET /api/v1/metrics/with-id/:id
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

**Response:** `Metrics`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/metrics/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Update Metrics

```
PUT /api/v1/metrics/with-id/:id
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
  metricKey?: string
  value?: number
  dimensions?: Record<string, unknown>
  recordedAt?: string
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
  "http://localhost:3000/api/v1/metrics/with-id/:id"
```

</details>

---

#### <Badge type="warning" text="PUT" /> Bulk Update Metrics

```
PUT /api/v1/metrics/bulk/
```

> Batch update multiple records.

**Headers:**

| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes | Bearer token |
| `x-company` | Yes | Company ID |

**Request Body:** Array of { pathParams, data: MetricsEdit }

**Response:** `Success`

<details>
<summary>curl example</summary>

```bash
curl -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/metrics/bulk/"
```

</details>

---

== Delete

#### <Badge type="danger" text="DELETE" /> Delete Metrics

```
DELETE /api/v1/metrics/with-id/:id
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
  "http://localhost:3000/api/v1/metrics/with-id/:id"
```

</details>

---

:::


::::
