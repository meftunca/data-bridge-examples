---
layout: home

hero:
  name: Innovation Hub API
  text: DataBridge V2 Demo
  tagline: 5 schemas · 54 tables · 14 RPC functions · Fully auto-generated API layer
  actions:
    - theme: brand
      text: API Reference
      link: /API_REFERENCE
    - theme: alt
      text: Swagger UI
      link: /swagger
    - theme: alt
      text: GitHub
      link: https://github.com/meftunca/data-bridge-examples

features:
  - icon: 🔐
    title: IAM — Identity & Access Management
    details: Organizations, users, roles, permissions, teams, API keys, sessions, invitations
    link: /iam/
  - icon: 📦
    title: Catalog — Product Catalog
    details: Brands, categories, products, variants, media, reviews, collections, tags, price history
    link: /catalog/
  - icon: 🛒
    title: Orders — Order Management
    details: Customers, coupons, orders, payments, refunds, carts, order status history
    link: /orders/
  - icon: 🚚
    title: Logistics — Logistics & Warehouse
    details: Warehouses, storage zones, inventory, stock movements, suppliers, purchase orders, shipments
    link: /logistics/
  - icon: 📊
    title: Analytics — Analytics & Reporting
    details: Audit logs, events, metrics, dashboards, alerts, notifications, reports
    link: /analytics/
  - icon: ⚡
    title: Technical Features
    details: "EventManager · SSE · Outbox Pattern · 25+ filter operators · 22 aggregations · Bulk CRUD · Swagger · TypeScript client"
---

## Quick Start

```bash
# Start the database
docker compose up -d

# Run SQL migrations
for f in sql/*.sql; do
  PGPASSWORD=databridge_demo_2026 psql -h localhost -p 55433 \
    -U databridge -d innovation_hub -f "$f"
done

# Start the API server
go run main.go
```

## Generated Structure

```
api_v2/
├── api.go                    # Setup() — wires all schemas
├── iam/                      # 11 tables
│   ├── structures/           # Go structs + GORM tags
│   ├── services/             # CRUD + RPC business logic
│   ├── controllers/          # HTTP handlers + Swagger annotations
│   ├── tests/                # API tests
│   └── pkg/iam_events/       # Event handler scaffolding
├── catalog/                  # 11 tables
├── orders/                   # 9 tables
├── logistics/                # 11 tables
├── analytics/                # 12 tables
└── shared/types/             # Cross-schema shared types
```
