---
layout: home

hero:
  name: Innovation Hub API
  text: DataBridge V2 Demo
  tagline: 5 schema · 54 tablo · 14 RPC · Tam otomatik üretilmiş API katmanı
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
    details: Organizasyonlar, kullanıcılar, roller, izinler, takımlar, API anahtarları, oturumlar, davetler
    link: /iam/
  - icon: 📦
    title: Catalog — Ürün Kataloğu
    details: Markalar, kategoriler, ürünler, varyantlar, medya, değerlendirmeler, koleksiyonlar, etiketler, fiyat geçmişi
    link: /catalog/
  - icon: 🛒
    title: Orders — Sipariş Yönetimi
    details: Müşteriler, kuponlar, siparişler, ödemeler, iadeler, sepetler, sipariş durumu geçmişi
    link: /orders/
  - icon: 🚚
    title: Logistics — Lojistik & Depo
    details: Depolar, depolama bölgeleri, envanter, stok hareketleri, tedarikçiler, satın alma siparişleri, sevkiyat
    link: /logistics/
  - icon: 📊
    title: Analytics — Analitik & Raporlama
    details: Denetim logları, olaylar, metrikler, panolar, uyarılar, bildirimler, raporlar
    link: /analytics/
  - icon: ⚡
    title: Teknik Özellikler
    details: "EventManager · SSE · Outbox Pattern · 25+ filtre operatörü · 22 aggregation · Bulk CRUD · Swagger · TypeScript client"
---

## Hızlı Başlangıç

```bash
# Veritabanını başlat
docker compose up -d

# SQL migration'ları çalıştır
for f in sql/*.sql; do
  PGPASSWORD=databridge_demo_2026 psql -h localhost -p 55433 \
    -U databridge -d innovation_hub -f "$f"
done

# API'yi başlat
go run main.go
```

## Üretilen Yapı

```
api_v2/
├── api.go                    # Setup() — tüm schema'ları bağlar
├── iam/                      # 11 tablo
│   ├── structures/           # Go struct'lar + GORM tag'ler
│   ├── services/             # CRUD + RPC iş mantığı
│   ├── controllers/          # HTTP handler'lar + Swagger annotation
│   ├── tests/                # API test'leri
│   └── pkg/iam_events/       # Event handler scaffold
├── catalog/                  # 11 tablo
├── orders/                   # 9 tablo
├── logistics/                # 11 tablo
├── analytics/                # 12 tablo
└── shared/types/             # Cross-schema paylaşımlı tipler
```
