
# Documentation Engine

## Multi-Format Documentation Generation

The doc-generator produces six distinct output formats from a
single PostgreSQL schema, serving different audiences and use cases.

| Format | Audience | Purpose |
|--------|----------|---------|
| VitePress Markdown | Developers | Browsable documentation site |
| OpenAPI JSON | API Consumers | Per-table API specification |
| TypeScript Types | Frontend Devs | Typed interfaces + React hooks |
| MCP Knowledge Graph | AI Agents | Structured knowledge for RAG |
| Training Data JSONL | ML Engineers | LLM fine-tuning datasets |
| Mermaid ER Diagrams | Architects | Visual schema relationships |

---

# VitePress Site Generation

Each table produces a dedicated documentation page with:

1. **Schema Overview** -- Table description, row count, constraints
2. **Column Reference** -- Type mapping across SQL, Go, and TypeScript
3. **API Endpoints** -- Generated REST operations with examples
4. **Relations** -- FK references with navigation links
5. **Mermaid ER Diagram** -- Visual relationship graph

**Sidebar Configuration** is generated automatically:

```json
[
  {
    "text": "IAM Schema",
    "items": [
      { "text": "Users", "link": "/iam/users" },
      { "text": "Roles", "link": "/iam/roles" },
      { "text": "Permissions", "link": "/iam/permissions" }
    ]
  }
]
```

---

# Column Reference Table

For each table, the doc generator produces a comprehensive column table:

| Column | SQL Type | Go Type | TS Type | Nullable | Default | Constraints |
|--------|----------|---------|---------|----------|---------|-------------|
| id | uuid | types.URID | string | No | gen_random_uuid() | PK |
| email | varchar(255) | string | string | No | - | UNIQUE |
| role_id | uuid | *types.URID | string? | Yes | - | FK(roles) |
| status | user_status | UserStatus | enum | No | 'active' | ENUM |
| metadata | jsonb | types.JSONB | Record | Yes | '{}' | - |
| location | geometry | types.Geometry | GeoJSON | Yes | - | - |

Three type systems (SQL, Go, TypeScript) are presented side-by-side,
eliminating the need for manual cross-referencing.

---

# API Documentation Tabs

Each table page includes tabbed API documentation:

### Tab: Search

```
GET /api/v1/{schema}/{table}/search
  ?page=1&size=20&sort=created_at:desc
  &filters[0][field]=status&filters[0][op]=eq&filters[0][value]=active

Response: Page[Model]
```

### Tab: Create

```
POST /api/v1/{schema}/{table}
Content-Type: application/json

Body: FormStruct (required fields highlighted)
Response: 201 Model
```

### Tab: Find / Update / Delete

```
GET    /api/v1/{schema}/{table}/:id  -> Model
PUT    /api/v1/{schema}/{table}/:id  -> Model (EditStruct body)
DELETE /api/v1/{schema}/{table}/:id  -> 204 No Content
```

---

# TypeScript Documentation

The doc generator produces frontend-ready TypeScript artifacts:

**Interfaces:**
```typescript
interface User {
  id: string;
  email: string;
  full_name: string;
  role_id?: string;
  status?: UserStatus;
  created_at: string;
  updated_at: string;
}

type UserForm = Omit<User, 'id' | 'created_at' | 'updated_at'>;
type UserEdit = Partial<UserForm>;
```

**React Query Hooks:**
```typescript
const useUsers = (params?: PaginationParams) =>
  useQuery(['users', params], () => fetchUsers(params));

const useCreateUser = () =>
  useMutation((form: UserForm) => createUser(form));
```

**Zod Validation Schemas:**
```typescript
const UserFormSchema = z.object({
  email: z.string().email(),
  full_name: z.string().min(1).max(100),
  role_id: z.string().uuid().optional(),
});
```

---

# MCP Knowledge Graph Export

Model Context Protocol export for AI agent consumption:

### Three JSONL Files

**graph_nodes.jsonl** -- Every entity in the schema:
```json
{"id": "iam.users", "type": "table", "schema": "iam",
 "name": "users", "columns": 7, "pk": "id",
 "description": "User accounts and authentication"}
```

**graph_edges.jsonl** -- Relationships between entities:
```json
{"source": "iam.users", "target": "iam.roles",
 "type": "belongs_to", "via": "role_id",
 "cardinality": "many_to_one"}
```

**doc_chunks.jsonl** -- Semantic text chunks for RAG:
```json
{"content": "The users table stores authentication...",
 "metadata": {"schema": "iam", "table": "users",
              "section": "overview"}}
```

**manifest.json** -- Index with node/edge/chunk counts.

---

# Training Data Generation

JSONL dataset for LLM fine-tuning with system/user/assistant triplets:

```json
{
  "messages": [
    {"role": "system", "content": "You are an API documentation expert..."},
    {"role": "user", "content": "What columns does the users table have?"},
    {"role": "assistant", "content": "The users table has 7 columns: id (UUID PK), email (VARCHAR UNIQUE), full_name (VARCHAR), role_id (UUID FK->roles), status (ENUM: active|inactive|suspended), created_at (TIMESTAMPTZ), updated_at (TIMESTAMPTZ)."}
  ]
}
```

### Generated Question Categories

| Category | Example Questions |
|----------|-----------------|
| Schema QA | "What tables are in the IAM schema?" |
| Column Info | "What type is the users.status column?" |
| Endpoint Usage | "How do I create a new user?" |
| Filter Examples | "How do I search for active users?" |
| React Query | "How do I fetch paginated products?" |
| Zod Schemas | "What validation does UserForm need?" |
| RPC Usage | "How do I call the calculate_totals function?" |

---

# Mermaid ER Diagrams

Automatically generated entity-relationship diagrams:

```
erDiagram
    USERS {
        uuid id PK
        varchar email UK
        varchar full_name
        uuid role_id FK
        user_status status
        timestamptz created_at
        timestamptz updated_at
    }
    ROLES {
        uuid id PK
        varchar name UK
        varchar slug UK
    }
    PERMISSIONS {
        uuid id PK
        varchar name UK
        varchar resource
        varchar action
    }
    ROLE_PERMISSIONS {
        uuid role_id FK
        uuid permission_id FK
    }

    USERS ||--o{ ROLES : "role_id"
    ROLE_PERMISSIONS }o--|| ROLES : "role_id"
    ROLE_PERMISSIONS }o--|| PERMISSIONS : "permission_id"
```

Diagrams are embedded in VitePress pages and rendered as SVG.

---

# Source Linker

The `SourceLinker` generates GitHub permalink badges for each
generated file, connecting documentation to source code:

```go
type SourceLinker struct {
    GitHubRepo string  // "github.com/org/repo"
    Branch     string  // "main"
    BasePath   string  // "api_v2"
}

func (sl *SourceLinker) StructLink(schema, table string) string {
    return fmt.Sprintf(
        "[![Source](https://img.shields.io/badge/source-GitHub-blue)]"+
        "(https://%s/blob/%s/%s/%s/structures/%s.go)",
        sl.GitHubRepo, sl.Branch, sl.BasePath, schema, table,
    )
}
```

Every documentation page includes direct links to:
- Struct definition file
- Service implementation file
- Controller handler file
- Endpoint registration file
