-- ============================================================
-- DataBridge V2 — Innovation Hub Demo
-- 004_orders_schema.sql
--
-- Order Processing & Payment Schema
-- Features demonstrated:
--   ✓ Complex multi-table relationships
--   ✓ ENUM usage (order_status, payment_status, payment_method)
--   ✓ Cross-schema FKs (→ iam.users, catalog.products, catalog.product_variants)
--   ✓ NUMERIC for financial precision
--   ✓ JSONB for flexible shipping/billing info
--   ✓ Soft delete support
--   ✓ Parent-child: orders → order_items
--   ✓ Outbox pattern compatible
-- ============================================================

-- ─── Customers (extends iam.users with commerce data) ───────
CREATE TABLE orders.customers (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    user_id         UUID NOT NULL REFERENCES iam.users(id),
    organization_id UUID REFERENCES iam.organizations(id),
    billing_address JSONB NOT NULL DEFAULT '{}',
    shipping_address JSONB NOT NULL DEFAULT '{}',
    tax_id          TEXT NOT NULL DEFAULT '',
    loyalty_points  INTEGER NOT NULL DEFAULT 0,
    tier            TEXT NOT NULL DEFAULT 'standard',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Coupons ────────────────────────────────────────────────
CREATE TABLE orders.coupons (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    code            TEXT NOT NULL UNIQUE,
    discount_type   TEXT NOT NULL DEFAULT 'percentage',
    discount_value  NUMERIC(12,2) NOT NULL DEFAULT 0,
    min_order_value NUMERIC(12,2) NOT NULL DEFAULT 0,
    max_uses        INTEGER NOT NULL DEFAULT 0,
    used_count      INTEGER NOT NULL DEFAULT 0,
    is_active       BOOLEAN NOT NULL DEFAULT true,
    valid_from      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    valid_until     TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Orders ─────────────────────────────────────────────────
CREATE TABLE orders.orders (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    order_number    TEXT NOT NULL UNIQUE,
    customer_id     UUID NOT NULL REFERENCES orders.customers(id),
    status          orders.order_status NOT NULL DEFAULT 'pending',
    subtotal        NUMERIC(14,2) NOT NULL DEFAULT 0.00,
    tax_amount      NUMERIC(14,2) NOT NULL DEFAULT 0.00,
    shipping_amount NUMERIC(14,2) NOT NULL DEFAULT 0.00,
    discount_amount NUMERIC(14,2) NOT NULL DEFAULT 0.00,
    total           NUMERIC(14,2) NOT NULL DEFAULT 0.00,
    currency        TEXT NOT NULL DEFAULT 'USD',
    coupon_id       UUID REFERENCES orders.coupons(id),
    shipping_address JSONB NOT NULL DEFAULT '{}',
    billing_address  JSONB NOT NULL DEFAULT '{}',
    notes           TEXT NOT NULL DEFAULT '',
    placed_at       TIMESTAMPTZ,
    confirmed_at    TIMESTAMPTZ,
    shipped_at      TIMESTAMPTZ,
    delivered_at    TIMESTAMPTZ,
    cancelled_at    TIMESTAMPTZ,
    created_by      UUID REFERENCES iam.users(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

-- ─── Order Items (child of orders) ──────────────────────────
CREATE TABLE orders.order_items (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    order_id        UUID NOT NULL REFERENCES orders.orders(id),
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    variant_id      UUID REFERENCES catalog.product_variants(id),
    quantity        INTEGER NOT NULL DEFAULT 1,
    unit_price      NUMERIC(12,2) NOT NULL DEFAULT 0.00,
    total_price     NUMERIC(12,2) NOT NULL DEFAULT 0.00,
    discount        NUMERIC(12,2) NOT NULL DEFAULT 0.00,
    metadata        JSONB NOT NULL DEFAULT '{}',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Payments ───────────────────────────────────────────────
CREATE TABLE orders.payments (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    order_id        UUID NOT NULL REFERENCES orders.orders(id),
    amount          NUMERIC(14,2) NOT NULL,
    currency        TEXT NOT NULL DEFAULT 'USD',
    method          orders.payment_method NOT NULL DEFAULT 'credit_card',
    status          orders.payment_status NOT NULL DEFAULT 'pending',
    provider_ref    TEXT NOT NULL DEFAULT '',
    provider_data   JSONB NOT NULL DEFAULT '{}',
    paid_at         TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Refunds ────────────────────────────────────────────────
CREATE TABLE orders.refunds (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    order_id        UUID NOT NULL REFERENCES orders.orders(id),
    payment_id      UUID NOT NULL REFERENCES orders.payments(id),
    amount          NUMERIC(14,2) NOT NULL,
    reason          TEXT NOT NULL DEFAULT '',
    status          TEXT NOT NULL DEFAULT 'pending',
    processed_by    UUID REFERENCES iam.users(id),
    processed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Order Status History ───────────────────────────────────
CREATE TABLE orders.order_status_history (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    order_id        UUID NOT NULL REFERENCES orders.orders(id),
    from_status     TEXT NOT NULL DEFAULT '',
    to_status       TEXT NOT NULL DEFAULT '',
    changed_by      UUID REFERENCES iam.users(id),
    note            TEXT NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Shopping Carts ─────────────────────────────────────────
CREATE TABLE orders.carts (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    customer_id     UUID NOT NULL REFERENCES orders.customers(id),
    is_active       BOOLEAN NOT NULL DEFAULT true,
    expires_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Cart Items ─────────────────────────────────────────────
CREATE TABLE orders.cart_items (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    cart_id         UUID NOT NULL REFERENCES orders.carts(id),
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    variant_id      UUID REFERENCES catalog.product_variants(id),
    quantity        INTEGER NOT NULL DEFAULT 1,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── RPC Functions ──────────────────────────────────────────
CREATE OR REPLACE FUNCTION orders.rpc__total_revenue()
RETURNS NUMERIC LANGUAGE SQL STABLE AS $$
    SELECT COALESCE(SUM(total), 0)::NUMERIC(14,2)
    FROM orders.orders 
    WHERE status NOT IN ('cancelled', 'refunded');
$$;

CREATE OR REPLACE FUNCTION orders.rpc__orders_by_status(p_status TEXT)
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM orders.orders WHERE status = p_status::orders.order_status;
$$;

CREATE OR REPLACE FUNCTION orders.rpc__customer_total_spent(p_customer_id UUID)
RETURNS NUMERIC LANGUAGE SQL STABLE AS $$
    SELECT COALESCE(SUM(total), 0)::NUMERIC(14,2)
    FROM orders.orders 
    WHERE customer_id = p_customer_id AND status NOT IN ('cancelled', 'refunded');
$$;
