# 🚀 DataBridge V2 — Innovation Hub Demo

> **PostgreSQL şemasından production-ready Go API'ye tek komutla.**  
> Bu demo, DataBridge V2'nin tüm özelliklerini eksiksiz olarak sergilemektedir.

---

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Mimari](#mimari)
- [Özellik Haritası](#özellik-haritası)
- [Hızlı Başlangıç](#hızlı-başlangıç)
- [Veritabanı Şeması](#veritabanı-şeması)
- [Konfigürasyon](#konfigürasyon)
- [Üretilen Çıktılar](#üretilen-çıktılar)
- [API Örnekleri](#api-örnekleri)
- [Özellik Detayları](#özellik-detayları)
- [Performans](#performans)
- [SSS](#sss)

---

## Genel Bakış

**DataBridge V2**, PostgreSQL veritabanı şemasını analiz ederek tam donanımlı bir Go REST API üretir:

```
PostgreSQL Schema ──→ DataBridge V2 ──→ Production-Ready Go API
     (input)           (generator)          (output)
```

### Bu Demo Ne İçerir?

| Bileşen | Açıklama |
|---------|----------|
| **5 PostgreSQL şeması** | `iam`, `catalog`, `orders`, `logistics`, `analytics` |
| **50+ tablo** | İlişkisel, hiyerarşik, M2M, composite PK |
| **13 ENUM tipi** | Go type-safe enum dönüşümü |
| **12 RPC fonksiyonu** | Parametreli ve parametresiz |
| **4 View** | Normal & Materialized |
| **Seed data** | Test-ready örnek veriler |
| **Tüm özellikler** | RBAC, SSE, Outbox, Events, Tests, Swagger |

---

## Mimari

```
┌─────────────────────────────────────────────────────┐
│                    Innovation Hub                    │
├──────────┬──────────┬──────────┬──────────┬─────────┤
│   IAM    │ Catalog  │  Orders  │Logistics │Analytics│
│          │          │          │          │         │
│ Users    │ Products │ Orders   │Warehouses│ Audit   │
│ Roles    │ Variants │ Payments │Inventory │ Events  │
│ Teams    │ Reviews  │ Refunds  │Shipments │Dashboards│
│ Perms    │ Tags     │ Carts    │ PO/PO    │ Reports │
│ Sessions │ Collect. │ Coupons  │ Tracking │ Alerts  │
│ API Keys │ Media    │ History  │ Bins     │ Metrics │
├──────────┴──────────┴──────────┴──────────┴─────────┤
│                                                      │
│  ┌─────────┐ ┌─────────┐ ┌──────┐ ┌──────────────┐ │
│  │  RBAC   │ │  SSE    │ │Outbox│ │Event Manager │ │
│  │Middleware│ │Real-time│ │Events│ │  (Pub/Sub)   │ │
│  └─────────┘ └─────────┘ └──────┘ └──────────────┘ │
│                                                      │
│  ┌──────────────────────────────────────────────────┐│
│  │           Pagination Engine                      ││
│  │  Full-text Search · Vector Search · Geo Search   ││
│  │  22 Aggregations · 25+ Filter Operators          ││
│  └──────────────────────────────────────────────────┘│
│                                                      │
│  ┌─────────┐ ┌──────────┐ ┌────────┐ ┌───────────┐ │
│  │ Swagger │ │  Tests   │ │ Shared │ │  go fmt   │ │
│  │ OpenAPI │ │Unit + API│ │ Types  │ │ goimports │ │
│  └─────────┘ └──────────┘ └────────┘ └───────────┘ │
└─────────────────────────────────────────────────────┘
```

### Şema İlişki Haritası

```
iam.organizations ◄──────── iam.users ◄──────── iam.user_roles
       │                       │    │                  │
       │                       │    │                  │
       ▼                       ▼    │                  ▼
iam.teams ──► iam.team_members │    │            iam.roles
       │                       │    │                  │
       │                       │    │                  ▼
       │                       │    │         iam.role_permissions
       │                       │    │                  │
       │                       │    │                  ▼
       │                       │    │          iam.permissions
       │                       │    │
       │              ┌────────┘    └─────────────┐
       │              ▼                            ▼
       │    catalog.products ──► catalog.product_reviews
       │         │    │
       │         │    ├──► catalog.product_variants
       │         │    ├──► catalog.product_media
       │         │    └──► catalog.price_history
       │         │
       │         ▼
       │    orders.order_items
       │         │
       │         ▼
       │    orders.orders ──── orders.payments ──► orders.refunds
       │         │
       │         ▼
       │    logistics.shipments ──► logistics.shipment_items
       │         │                        │
       │         ▼                        ▼
       │    logistics.warehouses ──► logistics.shipment_tracking
       │         │
       │         ├──► logistics.storage_zones ──► logistics.storage_bins
       │         └──► logistics.inventory ──► logistics.stock_movements
       │
       └──► analytics.audit_logs
            analytics.events
            analytics.dashboards ──► analytics.dashboard_widgets
            analytics.reports ──► analytics.report_executions
            analytics.alert_rules ──► analytics.alert_history
```

---

## Özellik Haritası

### ✅ Tamamı Bu Demoda Aktif

| # | Özellik | Config Flag | Açıklama |
|---|---------|-------------|----------|
| 1 | **CRUD API Generation** | — (her zaman) | Her tablo için Create, Read, Update, Delete, List, Bulk endpoints |
| 2 | **Struct Generation** | — (her zaman) | `Form`, `Model`, `Edit`, `Filter`, `Identity`, `Page` struct varyantları |
| 3 | **Swagger/OpenAPI** | — (her zaman) | Otomatik Swagger annotations her controller'da |
| 4 | **ENUM Generation** | — (her zaman) | PostgreSQL ENUM → Go type-safe dönüşümü |
| 5 | **RPC Functions** | — (DB introspection) | `rpc__` prefix'li DB fonksiyonları → REST endpoint |
| 6 | **Cross-Schema Shared Types** | — (otomatik FK analizi) | Şemalar arası FK referansları → unified struct |
| 7 | **RBAC** | `rbacEnabled: true` | Permission middleware her endpoint'e eklenir |
| 8 | **SSE (Server-Sent Events)** | `sseEnabled: true` | Her resource için `/events` real-time endpoint |
| 9 | **Outbox Pattern** | `outboxEnabled: true` | Service'ler `outbox_events` tablosuna yazar |
| 10 | **Event Manager** | `eventManagerIntegration: true` | Centralized event bus (request/event handlers) |
| 11 | **Event Handler Stubs** | `eventHandlers.generate: true` | Her şema için event handler dosyaları |
| 12 | **Test Generation** | `generateTests: true` | Unit test + API test dosyaları |
| 13 | **Advanced Pagination** | — (her zaman) | Full-text, vector, geo search + 22 aggregation |
| 14 | **Route Grouping** | `relaxGrouping` | Hiyerarşik veya düz rota yapısı |
| 15 | **Schema Prefix Control** | `skipSchemaPrefixInOps` | Swagger operationId prefix kontrolü |
| 16 | **Soft Delete** | — (`deleted_at` algılama) | `deleted_at` sütunu varsa otomatik soft delete |
| 17 | **UUID Primary Keys** | — (otomatik algılama) | UUID PK desteği |
| 18 | **Composite Primary Keys** | — (otomatik algılama) | Multi-column PK desteği |
| 19 | **Self-Referencing FK** | — (otomatik algılama) | Kendi kendine FK (ağaç yapıları) |
| 20 | **JSONB Columns** | — (otomatik algılama) | JSONB → Go interface{} |
| 21 | **Array Columns** | — (otomatik algılama) | `text[]` → `pq.StringArray` |
| 22 | **INET Type** | — (otomatik algılama) | `INET` → `net.IP` |
| 23 | **Views & Mat. Views** | — (otomatik algılama) | View'lar için read-only endpoint |
| 24 | **Parallel Processing** | — (otomatik) | Projeler paralel olarak işlenir |
| 25 | **Template Caching** | — (otomatik) | Template'ler bir kez yüklenir |

---

## Hızlı Başlangıç

### 1. Veritabanını Başlat

```bash
# Demo dizinine git
cd data-bridge-examples

# PostgreSQL'i başlat
docker compose up -d

# Sağlık kontrolü
docker compose ps
```

### 2. Bağlantı Bilgileri

```
Host:     localhost
Port:     55433
Database: innovation_hub
User:     databridge
Password: databridge_demo_2026
```

```bash
# psql ile bağlan
psql "host=localhost port=55433 dbname=innovation_hub user=databridge password=databridge_demo_2026"

# Şemaları kontrol et
\dn

# Tabloları listele
\dt iam.*
\dt catalog.*
\dt orders.*
\dt logistics.*
\dt analytics.*
```

### 3. DataBridge V2'yi Çalıştır

```bash
# Ana dizine dön
cd ..

# Generator'ı çalıştır
go run main.go --config data-bridge-examples/config.yaml

# Veya Makefile ile
make generate CONFIG=data-bridge-examples/config.yaml
```

### 4. Üretilen Kodun İncelenmesi

```bash
# Üretilen dosyaları gör
find api_v2/ -name "*.go" | head -30

# Endpoint dosyalarını kontrol et
cat api_v2/iam/iam.go
cat api_v2/catalog/catalog.go

# Struct dosyalarını incele
cat api_v2/iam/structures/users.go
```

---

## Veritabanı Şeması

### IAM (Identity & Access Management)

| Tablo | Sütun Sayısı | Özellikler |
|-------|-------------|------------|
| `organizations` | 11 | UUID PK, Self-ref FK, JSONB, Soft Delete |
| `users` | 15 | UUID PK, ENUM (status, auth_provider), JSONB, Soft Delete |
| `roles` | 7 | UUID PK |
| `permissions` | 6 | UUID PK |
| `role_permissions` | 4 | UUID PK, M2M, Composite Unique |
| `user_roles` | 8 | UUID PK, M2M, Composite Unique |
| `teams` | 8 | UUID PK, Array (`text[]`) |
| `team_members` | 7 | UUID PK, Composite Unique |
| `api_keys` | 11 | UUID PK, Array (`text[]`) |
| `sessions` | 7 | UUID PK, INET type |
| `invitations` | 9 | UUID PK |

**RPC Functions:**
- `rpc__count_active_users()` → Parametresiz
- `rpc__users_by_organization(p_org_id UUID)` → Parametreli
- `rpc__user_permissions(p_user_id UUID)` → TABLE dönüş tipi

### Catalog (Product Catalog)

| Tablo | Sütun Sayısı | Özellikler |
|-------|-------------|------------|
| `brands` | 8 | UUID PK, Cross-schema FK (→ iam.organizations) |
| `categories` | 9 | UUID PK, Self-ref FK (ağaç yapısı) |
| `products` | 20 | UUID PK, ENUM, JSONB, TSVECTOR, Array, Soft Delete |
| `product_variants` | 9 | UUID PK, Child of products |
| `product_media` | 10 | UUID PK, ENUM (media_type) |
| `product_reviews` | 10 | UUID PK, CHECK constraint, Cross-schema FK |
| `collections` | 9 | UUID PK |
| `collection_products` | 5 | UUID PK, M2M, Composite Unique |
| `tags` | 3 | UUID PK |
| `product_tags` | 3 | **Composite PK** (product_id, tag_id) |
| `price_history` | 7 | UUID PK, Audit trail |

**RPC Functions:**
- `rpc__count_active_products()` → Parametresiz
- `rpc__products_by_category(p_category_id UUID)` → Parametreli
- `rpc__avg_product_rating(p_product_id UUID)` → NUMERIC dönüş

### Orders (Order Processing)

| Tablo | Sütun Sayısı | Özellikler |
|-------|-------------|------------|
| `customers` | 11 | UUID PK, JSONB (address), Cross-schema FK |
| `coupons` | 12 | UUID PK, Business logic |
| `orders` | 20 | UUID PK, Multiple ENUMs, JSONB, Soft Delete |
| `order_items` | 10 | UUID PK, Child of orders, Multi-FK |
| `payments` | 11 | UUID PK, ENUM (method, status), JSONB |
| `refunds` | 10 | UUID PK, Cross-schema FK |
| `order_status_history` | 7 | UUID PK, State machine tracking |
| `carts` | 6 | UUID PK |
| `cart_items` | 7 | UUID PK, Child of carts |

**RPC Functions:**
- `rpc__total_revenue()` → NUMERIC dönüş
- `rpc__orders_by_status(p_status TEXT)` → Parametreli
- `rpc__customer_total_spent(p_customer_id UUID)` → Parametreli NUMERIC

### Logistics (Warehouse & Shipping)

| Tablo | Sütun Sayısı | Özellikler |
|-------|-------------|------------|
| `warehouses` | 9 | UUID PK, Cross-schema FK |
| `storage_zones` | 8 | UUID PK, Composite Unique |
| `storage_bins` | 7 | UUID PK, 3-level hierarchy |
| `inventory` | 12 | UUID PK, CHECK constraint, Composite Unique |
| `stock_movements` | 9 | UUID PK, Audit-style |
| `suppliers` | 10 | UUID PK, NUMERIC |
| `purchase_orders` | 13 | UUID PK, Dual approver FK |
| `purchase_order_items` | 8 | UUID PK, Child of PO |
| `shipments` | 13 | UUID PK, ENUM, Multi-FK |
| `shipment_items` | 5 | UUID PK, Multi-FK |
| `shipment_tracking` | 7 | UUID PK, Event-sourcing style |

**RPC Functions:**
- `rpc__low_stock_count(p_warehouse_id UUID)` → Parametreli
- `rpc__warehouse_utilization(p_warehouse_id UUID)` → NUMERIC %

### Analytics (Audit & Reporting)

| Tablo | Sütun Sayısı | Özellikler |
|-------|-------------|------------|
| `audit_logs` | 12 | UUID PK, ENUM, JSONB, INET |
| `events` | 11 | UUID PK, ENUM, Event-sourcing |
| `dashboards` | 8 | UUID PK, JSONB layout |
| `dashboard_widgets` | 9 | UUID PK, Child of dashboard |
| `reports` | 11 | UUID PK, ENUM (report_type), JSONB schedule |
| `report_executions` | 11 | UUID PK |
| `notifications` | 10 | UUID PK, Multi-channel |
| `alert_rules` | 10 | UUID PK, JSONB condition/action |
| `alert_history` | 7 | UUID PK |
| `metrics` | 6 | UUID PK, Time-series, JSONB dimensions |

**Views:**
- `recent_events` — Son 24 saatin olayları
- `unread_notifications` — Okunmamış bildirimler

**Materialized Views:**
- `daily_event_counts` — Günlük olay istatistikleri
- `user_activity_summary` — Kullanıcı aktivite özeti

**RPC Functions:**
- `rpc__event_count_by_severity(p_severity TEXT)`
- `rpc__unread_notification_count(p_user_id UUID)`
- `rpc__dashboard_count()`

---

## Konfigürasyon

### Config Alanları Açıklaması

```yaml
# ─── Global Defaults ─────────────────────────────────────
defaultHelperLibBasePath:         # Helper library import path
defaultPaginationPackagePath:     # Pagination engine import path
defaultImportRepo:                # Base import repo
defaultEventManagerIntegration:   # Event bus on/off
defaultGenerateTests:             # Test generation on/off
defaultRBACEnabled:               # RBAC middleware on/off
defaultSSEEnabled:                # Server-Sent Events on/off
defaultSSEPackagePath:            # SSE package import path
defaultOutboxEnabled:             # Outbox pattern on/off

# ─── Project Config ──────────────────────────────────────
projects:
  - name:                         # Proje adı
    outputFolder:                 # Çıktı dizini (api_v2)
    projectPath:                  # Go project path
    moduleName:                   # go.mod module adı
    importRepo:                   # Import path override

    # Feature Flag Overrides (per-project)
    generateTests:                # Proje bazlı test override
    rbacEnabled:                  # Proje bazlı RBAC override
    sseEnabled:                   # Proje bazlı SSE override
    outboxEnabled:                # Proje bazlı Outbox override

    # ─── Schema Config ─────────────────────────────────
    schemas:
      - name:                     # PostgreSQL şema adı
        skipSchemaPrefixInOps:    # Swagger operationId'den prefix kaldır
        ignoreTables: []          # Atlanacak tablolar
        relaxGrouping:            # Düz rota yapısı (nested yerine)
        eventHandlers:            # Event handler stub generation
          generate:               # Üret/üretme
          outputDir:              # Çıktı dizini
          packageName:            # Go package adı
```

### Bu Demo'daki Konfigürasyon Seçimleri

| Şema | `skipSchemaPrefixInOps` | `relaxGrouping` | `eventHandlers` |
|------|:-:|:-:|:-:|
| iam | ✅ | ❌ | ✅ |
| catalog | ✅ | ❌ | ✅ |
| orders | ✅ | ❌ | ✅ |
| logistics | ❌ | ✅ | ✅ |
| analytics | ✅ | ❌ | ✅ |

> **Not:** `logistics` şemasında farklı config değerleri kullanılarak farkın gösterilmesi amaçlanmıştır.

---

## Üretilen Çıktılar

DataBridge V2 çalıştırıldığında her şema için aşağıdaki dosya yapısı oluşturulur:

```
api_v2/
├── external_api.go                    # Ana API dosyası — tüm şemaları birleştirir
├── shared/
│   └── types/
│       └── shared_types.go            # Cross-schema shared structs
│
├── iam/
│   ├── iam.go                         # Route registration (endpoint file)
│   ├── structures/
│   │   ├── organizations.go           # Form, Model, Edit, Filter, Identity, Page structs
│   │   ├── users.go
│   │   ├── roles.go
│   │   ├── permissions.go
│   │   ├── role_permissions.go
│   │   ├── user_roles.go
│   │   ├── teams.go
│   │   ├── team_members.go
│   │   ├── api_keys.go
│   │   ├── sessions.go
│   │   ├── invitations.go
│   │   └── iam_enums.go               # user_status, auth_provider ENUMs
│   ├── services/
│   │   ├── organizations.go           # GORM CRUD + business logic
│   │   ├── users.go
│   │   ├── ...
│   │   ├── rpc_service.go             # RPC function implementations
│   │   └── outbox_event.go            # Outbox event struct
│   ├── controllers/
│   │   ├── organizations.go           # Fiber handlers + Swagger annotations
│   │   ├── users.go
│   │   ├── ...
│   │   └── rpc_controller.go          # RPC REST endpoints
│   └── tests/
│       ├── main_test.go               # Test setup
│       ├── test_helpers.go            # Shared helpers
│       ├── organizations_api_test.go  # API tests
│       └── ...
│
├── catalog/
│   ├── catalog.go
│   ├── structures/ ...
│   ├── services/ ...
│   ├── controllers/ ...
│   └── tests/ ...
│
├── orders/
│   ├── orders.go
│   ├── structures/ ...
│   ├── services/ ...
│   ├── controllers/ ...
│   └── tests/ ...
│
├── logistics/                         # relaxGrouping: true — düz rota yapısı
│   ├── logistics.go
│   ├── structures/ ...
│   ├── services/ ...
│   ├── controllers/ ...
│   └── tests/ ...
│
└── analytics/
    ├── analytics.go
    ├── structures/ ...
    ├── services/ ...
    ├── controllers/ ...
    └── tests/ ...

pkg/
├── iam_events/
│   └── event_handlers.go             # IAM event handler stubs
├── catalog_events/
│   └── event_handlers.go
├── orders_events/
│   └── event_handlers.go
├── logistics_events/
│   └── event_handlers.go
└── analytics_events/
    └── event_handlers.go
```

### Struct Varyantları (Her Tablo İçin)

DataBridge her tablo için **6 farklı struct** üretir:

| Struct | Kullanım | Örnek |
|--------|----------|-------|
| `{Table}Form` | Create request body | `UsersForm` |
| `{Table}Model` | Full DB model (GORM tags) | `UsersModel` |
| `{Table}Edit` | Update request body (pointer fields) | `UsersEdit` |
| `{Table}Filter` | Query filter parameters | `UsersFilter` |
| `{Table}Identity` | Minimal identifier | `UsersIdentity` |
| `{Table}Page` | Paginated response wrapper | `UsersPage` |

### Endpoint Tipleri (Her Tablo İçin)

| HTTP Method | Path | Açıklama |
|-------------|------|----------|
| `GET` | `/{resource}` | List (paginated, filterable, sortable) |
| `GET` | `/{resource}/:id` | Get by ID |
| `POST` | `/{resource}` | Create single |
| `POST` | `/{resource}/bulk` | Bulk create |
| `PUT` | `/{resource}/:id` | Update |
| `PUT` | `/{resource}/bulk` | Bulk update |
| `DELETE` | `/{resource}/:id` | Delete (or soft delete) |
| `DELETE` | `/{resource}/bulk` | Bulk delete |
| `GET` | `/{resource}/events` | SSE stream (if `sseEnabled`) |

---

## API Örnekleri

Detaylı API kullanım örnekleri için [API_EXAMPLES.md](./docs/API_EXAMPLES.md) dosyasına bakın.

### Temel CRUD

```bash
# List users (paginated)
curl "http://localhost:3000/api/v2/iam/users?page=1&size=10"

# Get single user
curl "http://localhost:3000/api/v2/iam/users/b0000000-0000-0000-0000-000000000001"

# Create user
curl -X POST "http://localhost:3000/api/v2/iam/users" \
  -H "Content-Type: application/json" \
  -d '{"email":"new@maple.dev","name":"New User","status":"active"}'

# Update user
curl -X PUT "http://localhost:3000/api/v2/iam/users/b0000000-0000-0000-0000-000000000001" \
  -H "Content-Type: application/json" \
  -d '{"display_name":"Updated Name"}'

# Delete (soft delete)
curl -X DELETE "http://localhost:3000/api/v2/iam/users/b0000000-0000-0000-0000-000000000001"
```

### Gelişmiş Filtreleme

```bash
# Multi-field filter
curl "http://localhost:3000/api/v2/catalog/products?filters=[{\"field\":\"status\",\"op\":\"EQ\",\"value\":\"active\"},{\"field\":\"base_price\",\"op\":\"GTE\",\"value\":100}]"

# Full-text search
curl "http://localhost:3000/api/v2/catalog/products?filters=[{\"field\":\"name\",\"op\":\"SEARCH\",\"value\":\"laptop developer\",\"options\":{\"language\":\"english\"}}]"

# Nested AND/OR logic
curl "http://localhost:3000/api/v2/catalog/products?filters=[{\"logic\":\"OR\",\"filters\":[{\"field\":\"status\",\"op\":\"EQ\",\"value\":\"active\"},{\"field\":\"is_featured\",\"op\":\"EQ\",\"value\":true}]}]"
```

### Aggregation & Analytics

```bash
# Count by status
curl "http://localhost:3000/api/v2/orders/orders?aggregations=[{\"field\":\"status\",\"function\":\"COUNT\"}]&group_by=status"

# Revenue statistics
curl "http://localhost:3000/api/v2/orders/orders?aggregations=[{\"field\":\"total\",\"function\":\"SUM\"},{\"field\":\"total\",\"function\":\"AVG\"},{\"field\":\"total\",\"function\":\"MAX\"}]"

# Moving average
curl "http://localhost:3000/api/v2/analytics/metrics?aggregations=[{\"field\":\"value\",\"function\":\"MOVING_AVG\",\"options\":{\"window\":7}}]"
```

### RPC Çağrıları

```bash
# Parametresiz
curl "http://localhost:3000/api/v2/iam/rpc/count-active-users"

# Parametreli
curl "http://localhost:3000/api/v2/catalog/rpc/avg-product-rating?p_product_id=20000000-0000-0000-0000-000000000001"

# Revenue sorgusu
curl "http://localhost:3000/api/v2/orders/rpc/total-revenue"
```

### SSE (Real-time Events)

```bash
# User events stream
curl -N "http://localhost:3000/api/v2/iam/users/events"

# Product updates stream
curl -N "http://localhost:3000/api/v2/catalog/products/events"

# Order status changes
curl -N "http://localhost:3000/api/v2/orders/orders/events"
```

### Preloading (Eager Load Relations)

```bash
# Products with brand and category
curl "http://localhost:3000/api/v2/catalog/products?preloads=Brand,Category"

# Orders with customer and items
curl "http://localhost:3000/api/v2/orders/orders?preloads=Customer,OrderItems&preload_size[OrderItems]=5"

# Filtered preloads
curl "http://localhost:3000/api/v2/orders/orders?preloads=OrderItems&preload_filters[OrderItems]=[{\"field\":\"quantity\",\"op\":\"GT\",\"value\":1}]"
```

---

## Özellik Detayları

### 1. Pagination Engine

838ns/operasyon, 3.4M+ eşzamanlı işlem/saniye performansıyla enterprise-grade pagination:

**25+ Filter Operatörü:**
- Text: `SEARCH`, `LIKE`, `ILIKE`, `REGEXP`
- Karşılaştırma: `EQ`, `NE`, `GT`, `GTE`, `LT`, `LTE`, `BETWEEN`
- Array: `IN`, `NOT_IN`, `CONTAINS`, `OVERLAPS`
- Vector: `VECTOR_SIMILARITY`, `VECTOR_DISTANCE`
- Geo: `GEO_DISTANCE`, `GEO_WITHIN`, `GEO_NEAREST`

**22 Aggregation Fonksiyonu:**
- Temel: `COUNT`, `SUM`, `AVG`, `MIN`, `MAX`
- İstatistik: `MEDIAN`, `MODE`, `STDDEV`, `VARIANCE`, `PERCENTILE`
- İleri: `GROWTH_RATE`, `MOVING_AVG`, `CUMULATIVE_SUM`, `CORRELATION`
- Window: `ROW_NUMBER`, `RANK`, `DENSE_RANK`, `PERCENT_RANK`

**16+ Dil Desteği:**
Turkish, English, Spanish, French, German, Italian, Portuguese, Russian, Chinese, Japanese, Korean, Arabic, Dutch, Swedish, Norwegian, Danish

### 2. Event Manager

Thread-safe, centralized event bus:

```go
// Request handler (gate — işlemi engelleyebilir)
em.OnRequest("/api/v2/orders/orders", events.CreationRequest, func(e events.Event) (bool, error) {
    // Sipariş oluşturmadan önce stok kontrolü
    return checkInventory(e.Data)
})

// Event handler (observe — yan etkiler)
em.OnEvent("/api/v2/orders/orders", events.CreationSuccess, func(e events.Event) {
    // Sipariş oluşturulduktan sonra bildirim gönder
    sendNotification(e.Data)
})

// Async handler (fire-and-forget)
em.OnEventAsync("/api/v2/catalog/products", events.UpdateSuccess, func(e events.Event) {
    // Ürün güncellendiğinde search index'i güncelle
    updateSearchIndex(e.Data)
})
```

**Event Tipleri:**
| Grup | Tipler |
|------|--------|
| Creation | Request, Success, Error, BatchRequest, BatchSuccess, BatchError |
| Update | Request, Success, Error, BatchRequest, BatchSuccess, BatchError |
| Deletion | Request, Success, Error, BatchRequest, BatchSuccess, BatchError |
| Query | Request, Success, Error, ListRequest, ListSuccess, ListError, PaginateRequest, PaginateSuccess, PaginateError |
| Auth | AuthorizationError, ParseError, ValidationError |

### 3. SSE (Server-Sent Events)

Non-blocking, topic-based real-time event streaming:

```go
hub := sse.NewHub(
    sse.WithBufferSize(256),
    sse.WithPingInterval(30 * time.Second),
    sse.WithBackplane(redisBackplane), // Multi-pod dağıtım
)

// Publish events
hub.Publish("catalog.products", sse.Event{
    Type:    "product.updated",
    Payload: productJSON,
})
```

**JavaScript Client:**
```javascript
const es = new EventSource('/api/v2/catalog/products/events');
es.onmessage = (event) => {
    const data = JSON.parse(event.data);
    console.log('Product update:', data);
};
```

### 4. Outbox Pattern

Güvenilir olay dağıtımı için transactional outbox:

```
┌──────────────┐     ┌───────────────┐     ┌──────────────┐
│   Service    │────►│  outbox_events│────►│  Event Bus   │
│  (DB Write)  │     │  (same TX)    │     │  (Consumer)  │
└──────────────┘     └───────────────┘     └──────────────┘
```

`outbox_events` tablosu yapısı:
```sql
CREATE TABLE outbox_events (
    id              BIGSERIAL PRIMARY KEY,
    aggregate_type  TEXT NOT NULL,        -- "catalog.products"
    aggregate_id    TEXT NOT NULL,        -- product UUID
    event_type      TEXT NOT NULL,        -- "product.created"
    payload         JSONB NOT NULL,       -- full event data
    created_at      TIMESTAMPTZ,
    published_at    TIMESTAMPTZ,          -- NULL = unpublished
    retry_count     INT DEFAULT 0
);
```

### 5. RBAC (Role-Based Access Control)

```go
// Üretilen endpoint'e otomatik eklenen middleware:
group.Get("/users", rbac.RequirePermission("iam.users", "read"), controller.ListUsers)
group.Post("/users", rbac.RequirePermission("iam.users", "create"), controller.CreateUser)
group.Put("/users/:id", rbac.RequirePermission("iam.users", "update"), controller.UpdateUser)
group.Delete("/users/:id", rbac.RequirePermission("iam.users", "delete"), controller.DeleteUser)
```

### 6. Cross-Schema Shared Types

DataBridge FK analizi yaparak şemalar arası referansları otomatik tespit eder:

```
catalog.products.created_by → iam.users.id
orders.orders.customer_id → orders.customers.id → iam.users.id
logistics.inventory.product_id → catalog.products.id
analytics.audit_logs.user_id → iam.users.id
```

Bu referanslar `shared/types/shared_types.go` dosyasında birleştirilir.

---

## Performans

### Generator Performansı

| Metrik | Değer |
|--------|-------|
| Template yükleme | Tek seferlik, cache'lenir |
| Paralel proje işleme | 2 concurrent proje |
| Dosya yazma | `sync.Pool` + buffered I/O |
| String deduplication | String pool ile bellek optimizasyonu |
| Bellek izleme | Otomatik memory snapshot profiling |

### Pagination Engine Performansı

| Metrik | Değer |
|--------|-------|
| Operasyon süresi | ~838ns/op |
| Concurrent throughput | 3.4M+ ops/sec |
| Filter parsing | O(n) tekli geçiş |
| Aggregation | SQL-native, DB-side hesaplama |

---

## SSS

### DataBridge V2 hangi veritabanlarını destekler?
Şu an yalnızca **PostgreSQL** desteklenmektedir.

### Üretilen kod hangi framework'ü kullanır?
- **HTTP:** [Fiber](https://gofiber.io/) (Express-inspired Go framework)
- **ORM:** [GORM](https://gorm.io/) (Go ORM)
- **Swagger:** [Swaggo](https://github.com/swaggo/swag) annotations

### Üretilen kodu düzenleyebilir miyim?
Evet, ancak kodu yeniden üretirseniz değişiklikleriniz kaybolur. Event handler stubs böyle tasarlanmıştır — ayrı bir `pkg/` dizininde üretilir ve üzerine yazılmaz.

### Yeni bir şema eklemek için ne yapmalıyım?
1. PostgreSQL'de şemayı oluşturun
2. `config.yaml`'a şema adını ekleyin
3. DataBridge'i yeniden çalıştırın

### Belirli tabloları atlamak mümkün mü?
Evet, `ignoreTables` listesine ekleyin:
```yaml
schemas:
  - name: "iam"
    ignoreTables: ["sessions", "api_keys"]
```

### Testler nasıl çalıştırılır?
```bash
# Tüm testler
go test ./api_v2/...

# Belirli bir şema
go test ./api_v2/iam/tests/...

# Coverage ile
go test -cover ./api_v2/...
```

---

## Lisans

Bu demo projesi [Maple Technologies](https://maple.dev) tarafından geliştirilmiştir.  
DataBridge V2, propietary bir code generator'dır.
