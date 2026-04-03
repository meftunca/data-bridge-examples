-- ============================================================
-- DataBridge V2 — Innovation Hub Demo
-- 003_catalog_schema.sql
--
-- Product Catalog & Content Management Schema
-- Features demonstrated:
--   ✓ Self-referencing categories (tree structure)
--   ✓ Cross-schema FK references (→ iam.users, iam.organizations)
--   ✓ ENUM usage (product_status, media_type)
--   ✓ JSONB for flexible attributes
--   ✓ NUMERIC for precise pricing
--   ✓ Full-text search columns (tsvector)
--   ✓ Array columns (text[])
--   ✓ Composite PK table
--   ✓ Parent-child hierarchies (product → variants, media)
--   ✓ Many-to-many through join tables
-- ============================================================

-- ─── Brands ─────────────────────────────────────────────────
CREATE TABLE catalog.brands (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    slug            TEXT NOT NULL UNIQUE,
    logo_url        TEXT NOT NULL DEFAULT '',
    website         TEXT NOT NULL DEFAULT '',
    description     TEXT NOT NULL DEFAULT '',
    organization_id UUID REFERENCES iam.organizations(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Categories (self-referencing tree) ─────────────────────
CREATE TABLE catalog.categories (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    slug            TEXT NOT NULL UNIQUE,
    description     TEXT NOT NULL DEFAULT '',
    parent_id       UUID REFERENCES catalog.categories(id),
    icon            TEXT NOT NULL DEFAULT '',
    sort_order      INTEGER NOT NULL DEFAULT 0,
    is_active       BOOLEAN NOT NULL DEFAULT true,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Products ───────────────────────────────────────────────
CREATE TABLE catalog.products (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    slug            TEXT NOT NULL UNIQUE,
    sku             TEXT NOT NULL UNIQUE,
    description     TEXT NOT NULL DEFAULT '',
    short_description TEXT NOT NULL DEFAULT '',
    status          catalog.product_status NOT NULL DEFAULT 'draft',
    brand_id        UUID REFERENCES catalog.brands(id),
    category_id     UUID REFERENCES catalog.categories(id),
    base_price      NUMERIC(12,2) NOT NULL DEFAULT 0.00,
    currency        TEXT NOT NULL DEFAULT 'USD',
    weight_kg       NUMERIC(8,3),
    dimensions_cm   JSONB NOT NULL DEFAULT '{}',
    attributes      JSONB NOT NULL DEFAULT '{}',
    tags            TEXT[] NOT NULL DEFAULT '{}',
    is_featured     BOOLEAN NOT NULL DEFAULT false,
    created_by      UUID REFERENCES iam.users(id),
    updated_by      UUID REFERENCES iam.users(id),
    search_vector   TSVECTOR,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

-- ─── Product Variants ───────────────────────────────────────
CREATE TABLE catalog.product_variants (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    sku             TEXT NOT NULL UNIQUE,
    price_override  NUMERIC(12,2),
    attributes      JSONB NOT NULL DEFAULT '{}',
    is_active       BOOLEAN NOT NULL DEFAULT true,
    stock_quantity  INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Product Media ──────────────────────────────────────────
CREATE TABLE catalog.product_media (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    media_type      catalog.media_type NOT NULL DEFAULT 'image',
    url             TEXT NOT NULL,
    alt_text        TEXT NOT NULL DEFAULT '',
    sort_order      INTEGER NOT NULL DEFAULT 0,
    is_primary      BOOLEAN NOT NULL DEFAULT false,
    metadata        JSONB NOT NULL DEFAULT '{}',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Product Reviews ────────────────────────────────────────
CREATE TABLE catalog.product_reviews (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    user_id         UUID NOT NULL REFERENCES iam.users(id),
    rating          SMALLINT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    title           TEXT NOT NULL DEFAULT '',
    body            TEXT NOT NULL DEFAULT '',
    is_verified     BOOLEAN NOT NULL DEFAULT false,
    helpful_count   INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Collections ────────────────────────────────────────────
CREATE TABLE catalog.collections (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    slug            TEXT NOT NULL UNIQUE,
    description     TEXT NOT NULL DEFAULT '',
    cover_image_url TEXT NOT NULL DEFAULT '',
    is_active       BOOLEAN NOT NULL DEFAULT true,
    start_date      DATE,
    end_date        DATE,
    created_by      UUID REFERENCES iam.users(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Collection Products (M2M) ─────────────────────────────
CREATE TABLE catalog.collection_products (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    collection_id   UUID NOT NULL REFERENCES catalog.collections(id),
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    sort_order      INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (collection_id, product_id)
);

-- ─── Tags (standalone) ─────────────────────────────────────
CREATE TABLE catalog.tags (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL UNIQUE,
    slug            TEXT NOT NULL UNIQUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Product Tags (M2M - Composite PK) ─────────────────────
CREATE TABLE catalog.product_tags (
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    tag_id          UUID NOT NULL REFERENCES catalog.tags(id),
    name            TEXT NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (product_id, tag_id)
);

-- ─── Price History ──────────────────────────────────────────
CREATE TABLE catalog.price_history (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    product_id      UUID NOT NULL REFERENCES catalog.products(id),
    old_price       NUMERIC(12,2) NOT NULL,
    new_price       NUMERIC(12,2) NOT NULL,
    changed_by      UUID REFERENCES iam.users(id),
    reason          TEXT NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── RPC Functions ──────────────────────────────────────────
CREATE OR REPLACE FUNCTION catalog.rpc__count_active_products()
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM catalog.products WHERE status = 'active';
$$;

CREATE OR REPLACE FUNCTION catalog.rpc__products_by_category(p_category_id UUID)
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM catalog.products WHERE category_id = p_category_id;
$$;

CREATE OR REPLACE FUNCTION catalog.rpc__avg_product_rating(p_product_id UUID)
RETURNS NUMERIC LANGUAGE SQL STABLE AS $$
    SELECT COALESCE(AVG(rating), 0)::NUMERIC(3,2)
    FROM catalog.product_reviews 
    WHERE product_id = p_product_id;
$$;
