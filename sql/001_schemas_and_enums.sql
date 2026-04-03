-- ============================================================
-- DataBridge V2 — Innovation Hub Demo
-- 001_schemas_and_enums.sql
-- 
-- 5 Schema Multi-Domain Architecture:
--   iam      → Identity & Access Management (RBAC, Auth)
--   catalog  → Product Catalog & Content Management
--   orders   → Order Processing & Payments
--   logistics → Warehouse, Inventory & Shipping
--   analytics → Audit, Reporting & Real-time Events
-- ============================================================

-- ─── Schemas ────────────────────────────────────────────────
CREATE SCHEMA IF NOT EXISTS iam;
CREATE SCHEMA IF NOT EXISTS catalog;
CREATE SCHEMA IF NOT EXISTS orders;
CREATE SCHEMA IF NOT EXISTS logistics;
CREATE SCHEMA IF NOT EXISTS analytics;

-- ─── ENUM Types ─────────────────────────────────────────────
-- DataBridge auto-detects and generates Go types for ENUMs

CREATE TYPE iam.user_status AS ENUM (
    'active', 'inactive', 'suspended', 'pending_verification'
);

CREATE TYPE iam.auth_provider AS ENUM (
    'local', 'google', 'github', 'azure_ad', 'saml'
);

CREATE TYPE catalog.product_status AS ENUM (
    'draft', 'active', 'discontinued', 'archived'
);

CREATE TYPE catalog.media_type AS ENUM (
    'image', 'video', 'document', 'audio', '3d_model'
);

CREATE TYPE orders.order_status AS ENUM (
    'pending', 'confirmed', 'processing', 'shipped', 
    'delivered', 'cancelled', 'refunded'
);

CREATE TYPE orders.payment_status AS ENUM (
    'pending', 'authorized', 'captured', 'failed', 'refunded'
);

CREATE TYPE orders.payment_method AS ENUM (
    'credit_card', 'debit_card', 'bank_transfer', 'wallet', 'crypto'
);

CREATE TYPE logistics.shipment_status AS ENUM (
    'preparing', 'picked_up', 'in_transit', 'out_for_delivery',
    'delivered', 'returned', 'lost'
);

CREATE TYPE logistics.stock_level AS ENUM (
    'out_of_stock', 'low', 'normal', 'high', 'overstock'
);

CREATE TYPE analytics.event_severity AS ENUM (
    'debug', 'info', 'warning', 'error', 'critical'
);

CREATE TYPE analytics.report_type AS ENUM (
    'daily', 'weekly', 'monthly', 'quarterly', 'annual', 'custom'
);
