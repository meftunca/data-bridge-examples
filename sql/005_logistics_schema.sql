-- ============================================================
-- DataBridge V2 — Innovation Hub Demo
-- 005_logistics_schema.sql
--
-- Warehouse, Inventory & Shipping Schema
-- Features demonstrated:
--   ✓ Deep cross-schema FKs (→ iam, catalog, orders)
--   ✓ Parent-child hierarchies (warehouse → zones → bins)
--   ✓ Tracking & state machines (shipment statuses)
--   ✓ Inventory management patterns
--   ✓ Composite unique constraints
--   ✓ Check constraints (quantity >= 0)
-- ============================================================

-- ─── Warehouses ─────────────────────────────────────────────
CREATE TABLE logistics.warehouses (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    code            TEXT NOT NULL UNIQUE,
    address         JSONB NOT NULL DEFAULT '{}',
    organization_id UUID NOT NULL REFERENCES iam.organizations(id),
    manager_id      UUID REFERENCES iam.users(id),
    is_active       BOOLEAN NOT NULL DEFAULT true,
    capacity        INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Storage Zones ──────────────────────────────────────────
CREATE TABLE logistics.storage_zones (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    warehouse_id    UUID NOT NULL REFERENCES logistics.warehouses(id),
    zone_code       TEXT NOT NULL DEFAULT 'A',
    zone_type       TEXT NOT NULL DEFAULT 'general',
    temperature_min NUMERIC(5,2),
    temperature_max NUMERIC(5,2),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (warehouse_id, zone_code)
);

-- ─── Storage Bins ───────────────────────────────────────────
CREATE TABLE logistics.storage_bins (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    zone_id         UUID NOT NULL REFERENCES logistics.storage_zones(id),
    bin_code        TEXT NOT NULL DEFAULT 'B001',
    max_capacity    INTEGER NOT NULL DEFAULT 100,
    current_count   INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Inventory ──────────────────────────────────────────────
CREATE TABLE logistics.inventory (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    variant_id      UUID REFERENCES catalog.product_variants(id),
    warehouse_id    UUID NOT NULL REFERENCES logistics.warehouses(id),
    bin_id          UUID REFERENCES logistics.storage_bins(id),
    quantity        INTEGER NOT NULL DEFAULT 0 CHECK (quantity >= 0),
    reserved        INTEGER NOT NULL DEFAULT 0,
    reorder_level   INTEGER NOT NULL DEFAULT 10,
    reorder_quantity INTEGER NOT NULL DEFAULT 50,
    last_counted_at TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (product_id, variant_id, warehouse_id)
);

-- ─── Stock Movements ────────────────────────────────────────
CREATE TABLE logistics.stock_movements (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    inventory_id    UUID NOT NULL REFERENCES logistics.inventory(id),
    movement_type   TEXT NOT NULL DEFAULT 'adjustment',
    quantity        INTEGER NOT NULL DEFAULT 0,
    reference_type  TEXT NOT NULL DEFAULT '',
    reference_id    TEXT NOT NULL DEFAULT '',
    performed_by    UUID REFERENCES iam.users(id),
    notes           TEXT NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Suppliers ──────────────────────────────────────────────
CREATE TABLE logistics.suppliers (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    code            TEXT NOT NULL UNIQUE,
    contact_name    TEXT NOT NULL DEFAULT '',
    contact_email   TEXT NOT NULL DEFAULT '',
    contact_phone   TEXT NOT NULL DEFAULT '',
    address         JSONB NOT NULL DEFAULT '{}',
    is_active       BOOLEAN NOT NULL DEFAULT true,
    rating          NUMERIC(3,2) NOT NULL DEFAULT 0.00,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Purchase Orders ────────────────────────────────────────
CREATE TABLE logistics.purchase_orders (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    po_number       TEXT NOT NULL UNIQUE,
    supplier_id     UUID NOT NULL REFERENCES logistics.suppliers(id),
    warehouse_id    UUID NOT NULL REFERENCES logistics.warehouses(id),
    status          TEXT NOT NULL DEFAULT 'draft',
    total_amount    NUMERIC(14,2) NOT NULL DEFAULT 0.00,
    expected_date   DATE,
    received_date   DATE,
    created_by      UUID REFERENCES iam.users(id),
    approved_by     UUID REFERENCES iam.users(id),
    notes           TEXT NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Purchase Order Items ───────────────────────────────────
CREATE TABLE logistics.purchase_order_items (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    purchase_order_id UUID NOT NULL REFERENCES logistics.purchase_orders(id),
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    quantity        INTEGER NOT NULL DEFAULT 1,
    unit_cost       NUMERIC(12,2) NOT NULL DEFAULT 0.00,
    received_qty    INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Shipments ──────────────────────────────────────────────
CREATE TABLE logistics.shipments (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    tracking_number TEXT NOT NULL DEFAULT '',
    order_id        UUID NOT NULL REFERENCES orders.orders(id),
    warehouse_id    UUID NOT NULL REFERENCES logistics.warehouses(id),
    carrier         TEXT NOT NULL DEFAULT '',
    status          logistics.shipment_status NOT NULL DEFAULT 'preparing',
    weight_kg       NUMERIC(8,3),
    shipped_at      TIMESTAMPTZ,
    delivered_at    TIMESTAMPTZ,
    estimated_delivery DATE,
    shipping_cost   NUMERIC(12,2) NOT NULL DEFAULT 0.00,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Shipment Items ─────────────────────────────────────────
CREATE TABLE logistics.shipment_items (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    shipment_id     UUID NOT NULL REFERENCES logistics.shipments(id),
    order_item_id   UUID NOT NULL REFERENCES orders.order_items(id),
    quantity        INTEGER NOT NULL DEFAULT 1,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Shipment Tracking Events ───────────────────────────────
CREATE TABLE logistics.shipment_tracking (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    shipment_id     UUID NOT NULL REFERENCES logistics.shipments(id),
    status          TEXT NOT NULL DEFAULT '',
    location        TEXT NOT NULL DEFAULT '',
    description     TEXT NOT NULL DEFAULT '',
    event_time      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── RPC Functions ──────────────────────────────────────────
CREATE OR REPLACE FUNCTION logistics.rpc__low_stock_count(p_warehouse_id UUID)
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM logistics.inventory 
    WHERE warehouse_id = p_warehouse_id AND quantity <= reorder_level;
$$;

CREATE OR REPLACE FUNCTION logistics.rpc__warehouse_utilization(p_warehouse_id UUID)
RETURNS NUMERIC LANGUAGE SQL STABLE AS $$
    SELECT COALESCE(
        (SUM(quantity)::NUMERIC / NULLIF(w.capacity, 0)) * 100, 0
    )::NUMERIC(5,2)
    FROM logistics.inventory i
    JOIN logistics.warehouses w ON w.id = i.warehouse_id
    WHERE i.warehouse_id = p_warehouse_id
    GROUP BY w.capacity;
$$;
