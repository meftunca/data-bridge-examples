
# Testing and Quality Assurance

## Multi-Layer Test Strategy

DataBridge V2 generates and maintains tests at four distinct levels,
ensuring quality from template correctness to end-to-end integration.

```
+-------------------------------------------+
|  E2E Tests (app/e2e_test.go)              |  Docker PostgreSQL
|  Full pipeline: SQL -> Generate -> Compile |  + generated test suite
+-------------------------------------------+
|  Integration Tests (*_test.go per table)  |  HTTP tests with
|  Real Fiber app + real DB queries          |  httptest.Server
+-------------------------------------------+
|  Unit Tests (controller_test.go per table)|  Mock services
|  Isolated handler logic validation         |  testify/mock
+-------------------------------------------+
|  Generator Tests (parser/*_test.go)       |  Template validation
|  Template output correctness               |  Output assertions
+-------------------------------------------+
```

---

# Generated Unit Tests

For each table controller, the generator produces unit tests
with mock service implementations:

```go
type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) Create(form *UserForm) (*User, error) {
    args := m.Called(form)
    return args.Get(0).(*User), args.Error(1)
}

func TestUserController_Create(t *testing.T) {
    mockService := new(MockUserService)
    ctrl := &UserController{Service: mockService}

    expectedUser := &User{ID: types.NewURID(), Email: "test@example.com"}
    mockService.On("Create", mock.AnythingOfType("*UserForm")).
        Return(expectedUser, nil)

    app := fiber.New()
    app.Post("/users", ctrl.Create)

    req := httptest.NewRequest("POST", "/users",
        strings.NewReader(`{"email":"test@example.com"}`))
    req.Header.Set("Content-Type", "application/json")

    resp, _ := app.Test(req)
    assert.Equal(t, 201, resp.StatusCode)
}
```

---

# Generated Integration Tests

Full HTTP round-trip tests against a real database:

```go
func TestUserAPI_Integration(t *testing.T) {
    // Setup: real DB + real Fiber app
    db := setupTestDB(t)
    app := fiber.New()
    service := NewUserService(db)
    ctrl := &UserController{Service: service}
    app.Post("/users", ctrl.Create)
    app.Get("/users/:id", ctrl.Find)

    // Test Create
    createResp, _ := app.Test(httptest.NewRequest("POST", "/users",
        strings.NewReader(`{
            "email": "integration@test.com",
            "full_name": "Integration User"
        }`)))
    assert.Equal(t, 201, createResp.StatusCode)

    var created User
    json.NewDecoder(createResp.Body).Decode(&created)
    assert.NotEmpty(t, created.ID)

    // Test Find
    findResp, _ := app.Test(httptest.NewRequest("GET",
        "/users/"+created.ID.String(), nil))
    assert.Equal(t, 200, findResp.StatusCode)
}
```

---

# Generated Test Helpers

Shared infrastructure for all test files within a schema:

```go
// test_helpers.go -- generated per schema

func setupTestDB(t *testing.T) *gorm.DB {
    dsn := os.Getenv("TEST_DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    require.NoError(t, err)

    // Auto-migrate all schema tables
    db.AutoMigrate(&User{}, &Role{}, &Permission{}, ...)

    t.Cleanup(func() {
        sqlDB, _ := db.DB()
        sqlDB.Close()
    })
    return db
}

func createTestUser(t *testing.T, db *gorm.DB) *User {
    user := &User{
        Email:    gofakeit.Email(),
        FullName: gofakeit.Name(),
    }
    require.NoError(t, db.Create(user).Error)
    return user
}
```

---

# E2E Test Pipeline

The `app/e2e_test.go` validates the entire generation pipeline:

```
1. Spin up PostgreSQL      Docker container with test schemas
       |
2. Write test config       Temporary config.yaml with all features
       |
3. Run generator           app.Run() with test config
       |
4. Assert file existence   Check all expected output files
       |
5. Initialize Go module    go mod init + go mod tidy
       |
6. Compile generated code  go build ./...
       |
7. Run generated tests     go test ./... on generated test suite
       |
8. Cleanup                 Remove temp files, stop Docker
```

This ensures that template changes never produce code that fails
to compile or whose generated tests fail.

---

# Generator Unit Tests

The generator itself has comprehensive tests for each component:

### Template Helper Tests
```go
func TestToCamel(t *testing.T) {
    tests := []struct{ input, expected string }{
        {"user_settings", "UserSettings"},
        {"api_v2_config", "ApiV2Config"},
        {"id", "ID"},
        {"url", "URL"},
    }
    for _, tt := range tests {
        assert.Equal(t, tt.expected, ToCamel(tt.input))
    }
}
```

### Creator Tests
```go
func TestControllerCreator_Create(t *testing.T) {
    table := models.Table{Name: "users", Columns: testColumns}
    creator := &ControllerCreator{Context: testContext}

    output, err := creator.CreateTableController(table, "iam")
    require.NoError(t, err)

    assert.Contains(t, output, "func (ctrl *UserController) Create")
    assert.Contains(t, output, "@Router /api/v1/iam/users [post]")
    assert.Contains(t, output, "c.BodyParser")
}
```

---

# Pagination Engine Tests

The pagination package has the most extensive test suite: 20+ test files
covering every operator, aggregation, and edge case.

| Test File | Coverage |
|-----------|----------|
| `core_test.go` | Basic pagination flow |
| `executor_test.go` | Query execution paths |
| `filter_test.go` | All filter operators |
| `recursive_filter_test.go` | Nested AND/OR logic |
| `aggregation_test.go` | All 22+ aggregation functions |
| `vector_search_test.go` | pgvector operators |
| `geo_search_test.go` | PostGIS operators |
| `search_language_test.go` | Full-text search languages |
| `lru_cache_test.go` | Cache hit/miss/eviction/TTL |
| `relations_integration_test.go` | Preload with nested relations |
| `query_builder_test.go` | GORM reflection and schema mapping |
| `benchmark_test.go` | Performance benchmarks |
| `simple_benchmark_test.go` | Throughput measurements |

---

# Quality Gates

```
Developer Workflow:
                                          Merge
  Code Change --> Unit Tests --> E2E --> Compile --> Deploy
       |              |          |         |
       |         Template    Full        go build
       |         helpers    pipeline     ./...
       |         Creators   Docker PG
       |         Pagination
       |
       v
  [Generator Change]
       |
       v
  Re-generate all projects
       |
       v
  Compile + test generated output
```

**No template change ships without:**
1. Generator unit tests passing
2. E2E pipeline completing (Docker PostgreSQL)
3. Generated code compiling cleanly
4. Generated tests passing

This feedback loop ensures that the generator and its output
remain in sync at all times.
