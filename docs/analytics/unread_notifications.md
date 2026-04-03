---
title: UnreadNotifications
---

# UnreadNotifications

**Table:** `analytics.unread_notifications`

::::tabs

:::tab FullStack

## Columns

| # | Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints | Description |
|---|--------|----------|---------|---------|----------|---------|-------------|-------------|
| 1 | `id` | `uuid` | `uuid.UUID` | `string` | YES | - | - | - |
| 2 | `name` | `text` | `string` | `string` | YES | - | - | - |
| 3 | `user_id` | `uuid` | `uuid.UUID` | `string` | YES | - | - | - |
| 4 | `title` | `text` | `string` | `string` | YES | - | - | - |
| 5 | `message` | `text` | `string` | `string` | YES | - | - | - |
| 6 | `channel` | `text` | `string` | `string` | YES | - | - | - |
| 7 | `action_url` | `text` | `string` | `string` | YES | - | - | - |
| 8 | `created_at` | `timestamp with time zone` | `time.Time` | `string` | YES | - | - | Auto-filled from session |


## Go Generated Code

> 📂 Source: [📄 `UnreadNotifications.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/UnreadNotifications.go) · [📄 `UnreadNotifications.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/services/UnreadNotifications.go) · [📄 `UnreadNotifications.go`](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/controllers/UnreadNotifications.go)

### Structs

::::tabs

:::tab Form

#### UnreadNotificationsForm [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/UnreadNotifications.go#:~:text=type%20UnreadNotificationsForm%20struct)

_Create payload — excludes auto-generated PK fields_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Title` | `*string` | `title` | YES |
| `Message` | `*string` | `message` | YES |
| `Channel` | `*string` | `channel` | YES |
| `ActionUrl` | `*string` | `actionUrl` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Model

#### UnreadNotifications [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/UnreadNotifications.go#:~:text=type%20UnreadNotifications%20struct)

_Full model — all columns + GORM/JSON tags + preload relations_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Title` | `*string` | `title` | YES |
| `Message` | `*string` | `message` | YES |
| `Channel` | `*string` | `channel` | YES |
| `ActionUrl` | `*string` | `actionUrl` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Edit

#### UnreadNotificationsEdit [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/UnreadNotifications.go#:~:text=type%20UnreadNotificationsEdit%20struct)

_Update payload — all fields are pointers (partial update)_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Title` | `*string` | `title` | YES |
| `Message` | `*string` | `message` | YES |
| `Channel` | `*string` | `channel` | YES |
| `ActionUrl` | `*string` | `actionUrl` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Filter

#### UnreadNotificationsFilter [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/UnreadNotifications.go#:~:text=type%20UnreadNotificationsFilter%20struct)

_Query filter — all fields are pointers_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Title` | `*string` | `title` | YES |
| `Message` | `*string` | `message` | YES |
| `Channel` | `*string` | `channel` | YES |
| `ActionUrl` | `*string` | `actionUrl` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

:::tab Page

#### UnreadNotificationsPage [![source](https://img.shields.io/badge/source-gray?style=flat-square&logo=github)](https://github.com/meftunca/data-bridge-examples/blob/main//analytics/structures/UnreadNotifications.go#:~:text=type%20UnreadNotificationsPage%20struct)

_Paginated response wrapper_

| Field | Go Type | JSON Key | Nullable |
|-------|---------|----------|----------|
| `Id` | `*uuid.UUID` | `id` | YES |
| `Name` | `*string` | `name` | YES |
| `UserId` | `*uuid.UUID` | `userId` | YES |
| `Title` | `*string` | `title` | YES |
| `Message` | `*string` | `message` | YES |
| `Channel` | `*string` | `channel` | YES |
| `ActionUrl` | `*string` | `actionUrl` | YES |
| `CreatedAt` | `*time.Time` | `createdAt` | YES |

::::

### Service & Endpoints

::::tabs

:::tab Service Methods

_View — read-only, no write service methods._

:::tab Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/unread-notifications/` | Search with query params |
| `GET` | `/unread-notifications/pagination` | Paginated listing |

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
export interface UnreadNotifications {
  id?: string;
  name?: string;
  userId?: string;
  title?: string;
  message?: string;
  channel?: string;
  actionUrl?: string;
  createdAt?: string;
}

export interface UnreadNotificationsForm {
  id?: string;
  name?: string;
  userId?: string;
  title?: string;
  message?: string;
  channel?: string;
  actionUrl?: string;
  createdAt?: string;
}

export interface UnreadNotificationsEdit {
  id?: string;
  name?: string;
  userId?: string;
  title?: string;
  message?: string;
  channel?: string;
  actionUrl?: string;
  createdAt?: string;
}

export interface UnreadNotificationsPage {
  data: UnreadNotifications[];
  total: number;
  page: number;
  size: number;
  totalPages: number;
}

export type UnreadNotificationsPathQuery = {
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

const UnreadNotificationsKeys = {
  all: ["unread_notifications"] as const,
  lists: () => [...UnreadNotificationsKeys.all, "list"] as const,
  detail: (id: any) => [...UnreadNotificationsKeys.all, "detail", id] as const,
} as const;

export function useUnreadNotificationsList(query?: UnreadNotificationsPathQuery) {
  return useQuery({
    queryKey: [...UnreadNotificationsKeys.lists(), query],
    queryFn: () => fetch(`/unread-notifications/pagination`, { method: "GET" }).then(r => r.json()) as Promise<UnreadNotificationsPage>,
  });
}

export function useUnreadNotificationsDetail(id: any) {
  return useQuery({
    queryKey: UnreadNotificationsKeys.detail(id),
    queryFn: () => fetch(`/unread-notifications/with-id/:id`).then(r => r.json()) as Promise<UnreadNotifications>,
  });
}

```

:::tab Zod Validation

```typescript
import { z } from "zod";

export const UnreadNotificationsFormSchema = z.object({
  id: z.string().uuid().optional(),
  name: z.string().optional(),
  userId: z.string().uuid().optional(),
  title: z.string().optional(),
  message: z.string().optional(),
  channel: z.string().optional(),
  actionUrl: z.string().optional(),
  createdAt: z.string().datetime().optional(),
});

export type UnreadNotificationsFormInput = z.infer<typeof UnreadNotificationsFormSchema>;

```

::::


:::tab API

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from './unread_notifications.openapi.json'
useOpenapi({ spec })
</script>


## API Reference

::::tabs

:::tab Search

#### <Badge type="info" text="GET" /> Search UnreadNotifications

```
GET /api/v1/unread-notifications/
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
| `userId` | `string (uuid)` | No | Filter by user_id |
| `title` | `string` | No | Filter by title |
| `message` | `string` | No | Filter by message |
| `channel` | `string` | No | Filter by channel |
| `actionUrl` | `string` | No | Filter by action_url |

**Response:** `UnreadNotifications[]`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/unread-notifications/"
```

</details>

---

#### <Badge type="tip" text="POST" /> Search UnreadNotifications (POST)

```
POST /api/v1/unread-notifications/search
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

**Response:** `UnreadNotifications[]`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/unread-notifications/search"
```

</details>

---

:::tab Pagination

#### <Badge type="info" text="GET" /> Paginate UnreadNotifications

```
GET /api/v1/unread-notifications/pagination
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
| `userId` | `string (uuid)` | No | Filter by user_id |
| `title` | `string` | No | Filter by title |
| `message` | `string` | No | Filter by message |
| `channel` | `string` | No | Filter by channel |
| `actionUrl` | `string` | No | Filter by action_url |

**Response:** `PaginationResponse<UnreadNotifications>`

<details>
<summary>curl example</summary>

```bash
curl -X GET \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  "http://localhost:3000/api/v1/unread-notifications/pagination"
```

</details>

---

#### <Badge type="tip" text="POST" /> Paginate UnreadNotifications (POST)

```
POST /api/v1/unread-notifications/pagination
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

**Response:** `PaginationResponse<UnreadNotifications>`

<details>
<summary>curl example</summary>

```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "x-company: $COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{}' \
  "http://localhost:3000/api/v1/unread-notifications/pagination"
```

</details>

---

::::


::::
