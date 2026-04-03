
# Generated Code Walkthrough

## Six Struct Variants Per Table

For each database table, the generator produces six Go struct variants,
each serving a distinct purpose in the API lifecycle.

| Variant | Purpose | Fields |
|---------|---------|--------|
| `Form` | Create input (POST body) | Non-auto columns, no PK, no timestamps |
| `Model` | Database record (GORM model) | All columns with GORM tags |
| `Edit` | Update input (PUT body) | All mutable fields as pointers |
| `Filter` | Search/filter parameters | All fields as pointers for optional filtering |
| `Identity` | Minimal FK reference | PK + display field only |
| `Page` | Paginated response wrapper | Model + Count + computed fields |

---

# Struct Generation Example

**PostgreSQL Table:**
```sql
CREATE TABLE iam.users (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email       VARCHAR(255) NOT NULL UNIQUE,
    full_name   VARCHAR(100) NOT NULL,
    role_id     UUID REFERENCES iam.roles(id),
    status      iam.user_status DEFAULT 'active',
    created_at  TIMESTAMPTZ DEFAULT now(),
    updated_at  TIMESTAMPTZ DEFAULT now()
);
```

**Generated Form Struct:**
```go
type UserForm struct {
    Email    string      `json:"email" example:"sarah@acme.com"`
    FullName string      `json:"full_name" example:"Sarah Connor"`
    RoleID   *types.URID `json:"role_id,omitempty"`
    Status   *UserStatus `json:"status,omitempty"`
}
```

---

# Generated Model and Edit Structs

**Model (full GORM entity):**
```go
type User struct {
    ID        types.URID    `gorm:"primaryKey;type:uuid" json:"id"`
    Email     string        `gorm:"not null;uniqueIndex" json:"email"`
    FullName  string        `gorm:"not null" json:"full_name"`
    RoleID    *types.URID   `gorm:"type:uuid" json:"role_id,omitempty"`
    Status    *UserStatus   `gorm:"type:user_status" json:"status,omitempty"`
    CreatedAt types.NullTime `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt types.NullTime `gorm:"autoUpdateTime" json:"updated_at"`

    // Relations (Preload)
    Role *RoleIdentity `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}
```

**Edit (pointer fields for partial updates):**
```go
type UserEdit struct {
    Email    *string     `json:"email,omitempty"`
    FullName *string     `json:"full_name,omitempty"`
    RoleID   *types.URID `json:"role_id,omitempty"`
    Status   *UserStatus `json:"status,omitempty"`
}
```

---

# Service Layer Generation

Each table gets an interface and an implementation:

```go
type IUserService interface {
    Create(form *UserForm) (*User, error)
    BulkCreate(forms []*UserForm) ([]*User, error)
    Update(id types.URID, edit *UserEdit) (*User, error)
    BulkUpdate(edits map[types.URID]*UserEdit) ([]*User, error)
    Delete(id types.URID) error
    Find(id types.URID) (*User, error)
}

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) IUserService {
    return &UserService{db: db}
}
```

The interface enables mock-based testing without database dependencies.

---

# Service Implementation Details

**Create with Outbox Pattern (when enabled):**
```go
func (s *UserService) Create(form *UserForm) (*User, error) {
    record := &User{
        Email:    form.Email,
        FullName: form.FullName,
        RoleID:   form.RoleID,
        Status:   form.Status,
    }

    err := s.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(record).Error; err != nil {
            return err
        }
        // Outbox: persist event for async processing
        outboxEntry := OutboxEvent{
            AggregateType: "user",
            AggregateID:   record.ID.String(),
            EventType:     "user.created",
            Payload:       sonic.Marshal(record),
        }
        return tx.Create(&outboxEntry).Error
    })

    return record, err
}
```

---

# Service: Associations and Nested Creation

When `AssociationsCreating` is enabled, the service handles
parent-child relationships within a single transaction:

```go
func (s *OrderService) Create(form *OrderForm) (*Order, error) {
    return s.createWithAssociations(form)
}

func (s *OrderService) createWithAssociations(form *OrderForm) (*Order, error) {
    record := mapFormToRecord(form)

    err := s.db.Transaction(func(tx *gorm.DB) error {
        session := tx.Session(&gorm.Session{
            FullSaveAssociations: true,
        })
        return session.Create(record).Error
    })

    return record, err
}
```

Child records (e.g., `order_items`) inherit `company_id` and `created_by`
from the parent context automatically via `getChildListInfos()`.

---

# Controller Generation

Each controller wraps a service with HTTP handling:

```go
type UserController struct {
    Service IUserService
    EM      *events.EventManager  // optional
}

// @Summary Create a new user
// @Tags    IAM - Users
// @Accept  json
// @Produce json
// @Param   body body UserForm true "User data"
// @Success 201 {object} User
// @Failure 400 {object} swag.ErrorResponse
// @Router  /api/v1/iam/users [post]
func (ctrl *UserController) Create(c *fiber.Ctx) error {
    var form UserForm
    if err := c.BodyParser(&form); err != nil {
        return c.Status(400).JSON(swag.ErrorResponse{
            Message: "Invalid request body",
        })
    }

    result, err := ctrl.Service.Create(&form)
    if err != nil {
        ctrl.EM.EmitWithData(c, events.CreationError, err, result)
        return c.Status(500).JSON(swag.ErrorResponse{...})
    }

    ctrl.EM.EmitWithData(c, events.CreationSuccess, nil, result)
    return c.Status(201).JSON(result)
}
```

---

# Controller: Pagination Endpoint

The pagination controller leverages the generic pagination engine:

```go
// @Summary  Search users with pagination
// @Tags     IAM - Users
// @Produce  json
// @Param    page     query int    false "Page number"
// @Param    size     query int    false "Page size"
// @Param    sort     query string false "Sort field"
// @Param    fields   query string false "Select fields"
// @Param    filters  query string false "Filter conditions"
// @Param    preloads query string false "Preload relations"
// @Success  200 {object} pagination.Page[User]
// @Router   /api/v1/iam/users/search [get]
func (ctrl *UserController) Paginate(c *fiber.Ctx) error {
    result, err := pagination.NewPagination[User]().
        With(ctrl.DB).
        Request(c).
        Response()

    if err != nil {
        return c.Status(500).JSON(swag.ErrorResponse{...})
    }

    return c.JSON(result)
}
```

A single generic call handles filtering, sorting, preloading,
field selection, and response formatting.

---

# Endpoint Registration

Route registration uses recursive tree grouping:

```go
func Run(app fiber.Router, db *gorm.DB,
         em *events.EventManager) {

    api := app.Group("/iam")

    // Users group
    users := api.Group("/users")
    userCtrl := &UserController{
        Service: NewUserService(db), EM: em,
    }
    users.Post("/", userCtrl.Create)
    users.Post("/bulk", userCtrl.BulkCreate)
    users.Get("/search", userCtrl.Paginate)
    users.Get("/:id", userCtrl.Find)
    users.Put("/:id", userCtrl.Update)
    users.Delete("/:id", userCtrl.Delete)

    // Nested: user_settings under /users
    settings := users.Group("/settings")
    settingsCtrl := &UserSettingController{...}
    settings.Post("/", settingsCtrl.Create)
    // ... more routes

    // SSE streams (when enabled)
    users.Get("/stream", sse.Hub.StreamResource("iam.users"))
}
```
