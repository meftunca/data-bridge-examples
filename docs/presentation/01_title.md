
<!-- _class: lead -->
<!-- _paginate: false -->
<!-- _backgroundColor: #0f3460 -->
<!-- _color: #ffffff -->

# DataBridge V2

## From PostgreSQL Schema to Production-Ready API

**Automated Backend Code Generation Platform**

Maple Technologies -- Engineering Division

---

# Agenda

1. Architecture Overview and Design Philosophy
2. Configuration System
3. Parser Engine and Schema Introspection
4. Template System and Code Generation
5. Generated Code Walkthrough
6. Pagination Engine
7. Real-Time Events and SSE
8. Documentation Engine
9. Testing and Quality Assurance
10. Performance Engineering
11. Live Demo Showcase
12. Summary and Roadmap

---

# Executive Summary

**DataBridge V2** is a schema-driven code generation platform that transforms
PostgreSQL database schemas into fully operational REST APIs.

### What It Generates

- **Go Structs** with GORM tags, validation, and JSON serialization
- **Service Layer** with CRUD, bulk operations, and transaction support
- **HTTP Controllers** with Swagger annotations and error handling
- **Route Registration** with recursive URL tree grouping
- **Unit and Integration Tests** with mock infrastructure
- **Event Handlers** with sync/async dispatch
- **RPC Wrappers** for PostgreSQL stored procedures
- **API Documentation** (VitePress, OpenAPI, MCP, Training Data)

---

# Key Value Proposition

### Before DataBridge V2

- Manual struct definition from SQL schemas
- Hand-written CRUD services with inconsistent patterns
- Copy-paste controller boilerplate across endpoints
- Swagger annotations maintained separately from code
- No standardized test generation

### After DataBridge V2

- **One command** generates the entire API layer
- Consistent patterns across all tables and schemas
- Swagger docs embedded in generated code
- Tests generated alongside business logic
- Real-time event hooks and SSE built in

---

# Technical Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.25+ |
| HTTP Framework | Fiber (Express-style, built on fasthttp) |
| ORM | GORM with PostgreSQL driver |
| Database | PostgreSQL 14+ (with pgvector, PostGIS support) |
| Documentation | Swagger/OpenAPI 3.x, VitePress |
| Template Engine | Go `text/template` with 40+ custom functions |
| Serialization | bytedance/sonic (high-performance JSON) |
| Testing | Go testing + httptest + testify |
| CI/CD | GitHub Actions |

---

# Project Scale

| Metric | Value |
|--------|-------|
| Core Source Files | 80+ Go files |
| Template Files | 13 code generation templates |
| Pagination Engine | 30+ files, 3000+ LOC |
| Template Functions | 40+ custom FuncMap entries |
| Supported SQL Types | 25+ PostgreSQL types mapped to Go |
| Doc Generator Outputs | 6 formats (MD, OpenAPI, MCP, TS, Training, Mermaid) |
| Test Coverage | Unit, Integration, E2E, Benchmark |
| Generated Output per Table | 6-8 files (struct, service, controller, tests, events) |
