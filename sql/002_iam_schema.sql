-- ============================================================
-- DataBridge V2 — Innovation Hub Demo
-- 002_iam_schema.sql
--
-- Identity & Access Management Schema
-- Features demonstrated:
--   ✓ Self-referencing FK (organizations → parent)
--   ✓ Cross-table FK relationships
--   ✓ ENUM usage (user_status, auth_provider)
--   ✓ UUID primary keys
--   ✓ Composite unique constraints
--   ✓ Soft delete (deleted_at)
--   ✓ RBAC tables (roles, permissions)
--   ✓ Audit columns (created_by, updated_by)
--   ✓ JSONB columns
--   ✓ Array columns (text[])
-- ============================================================

-- ─── Organizations (self-referencing hierarchy) ─────────────
CREATE TABLE iam.organizations (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    slug            TEXT NOT NULL UNIQUE,
    description     TEXT NOT NULL DEFAULT '',
    logo_url        TEXT NOT NULL DEFAULT '',
    parent_id       UUID REFERENCES iam.organizations(id),
    settings        JSONB NOT NULL DEFAULT '{}',
    is_active       BOOLEAN NOT NULL DEFAULT true,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

-- ─── Users ──────────────────────────────────────────────────
CREATE TABLE iam.users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email           TEXT NOT NULL UNIQUE,
    name            TEXT NOT NULL,
    display_name    TEXT NOT NULL DEFAULT '',
    avatar_url      TEXT NOT NULL DEFAULT '',
    phone           TEXT NOT NULL DEFAULT '',
    status          iam.user_status NOT NULL DEFAULT 'pending_verification',
    auth_provider   iam.auth_provider NOT NULL DEFAULT 'local',
    organization_id UUID REFERENCES iam.organizations(id),
    metadata        JSONB NOT NULL DEFAULT '{}',
    last_login_at   TIMESTAMPTZ,
    email_verified  BOOLEAN NOT NULL DEFAULT false,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

-- ─── Roles ──────────────────────────────────────────────────
CREATE TABLE iam.roles (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    slug            TEXT NOT NULL UNIQUE,
    description     TEXT NOT NULL DEFAULT '',
    organization_id UUID REFERENCES iam.organizations(id),
    is_system       BOOLEAN NOT NULL DEFAULT false,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Permissions ────────────────────────────────────────────
CREATE TABLE iam.permissions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    resource        TEXT NOT NULL DEFAULT '*',
    action          TEXT NOT NULL DEFAULT 'read',
    description     TEXT NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Role Permissions (M2M) ────────────────────────────────
CREATE TABLE iam.role_permissions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    role_id         UUID NOT NULL REFERENCES iam.roles(id),
    permission_id   UUID NOT NULL REFERENCES iam.permissions(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (role_id, permission_id)
);

-- ─── User Roles (M2M) ──────────────────────────────────────
CREATE TABLE iam.user_roles (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    user_id         UUID NOT NULL REFERENCES iam.users(id),
    role_id         UUID NOT NULL REFERENCES iam.roles(id),
    granted_by      UUID REFERENCES iam.users(id),
    granted_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, role_id)
);

-- ─── Teams ──────────────────────────────────────────────────
CREATE TABLE iam.teams (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    description     TEXT NOT NULL DEFAULT '',
    organization_id UUID NOT NULL REFERENCES iam.organizations(id),
    lead_id         UUID REFERENCES iam.users(id),
    tags            TEXT[] NOT NULL DEFAULT '{}',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Team Members ───────────────────────────────────────────
CREATE TABLE iam.team_members (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    team_id         UUID NOT NULL REFERENCES iam.teams(id),
    user_id         UUID NOT NULL REFERENCES iam.users(id),
    role            TEXT NOT NULL DEFAULT 'member',
    joined_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (team_id, user_id)
);

-- ─── API Keys ───────────────────────────────────────────────
CREATE TABLE iam.api_keys (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    key_hash        TEXT NOT NULL,
    user_id         UUID NOT NULL REFERENCES iam.users(id),
    organization_id UUID NOT NULL REFERENCES iam.organizations(id),
    scopes          TEXT[] NOT NULL DEFAULT '{}',
    is_active       BOOLEAN NOT NULL DEFAULT true,
    last_used_at    TIMESTAMPTZ,
    expires_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Audit Sessions ────────────────────────────────────────
CREATE TABLE iam.sessions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    user_id         UUID NOT NULL REFERENCES iam.users(id),
    ip_address      INET,
    user_agent      TEXT NOT NULL DEFAULT '',
    is_active       BOOLEAN NOT NULL DEFAULT true,
    expires_at      TIMESTAMPTZ NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Organization Invitations ───────────────────────────────
CREATE TABLE iam.invitations (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    email           TEXT NOT NULL,
    organization_id UUID NOT NULL REFERENCES iam.organizations(id),
    invited_by      UUID NOT NULL REFERENCES iam.users(id),
    role_id         UUID REFERENCES iam.roles(id),
    token           TEXT NOT NULL UNIQUE,
    accepted_at     TIMESTAMPTZ,
    expires_at      TIMESTAMPTZ NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── RPC Functions ──────────────────────────────────────────
CREATE OR REPLACE FUNCTION iam.rpc__count_active_users()
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM iam.users WHERE status = 'active';
$$;

CREATE OR REPLACE FUNCTION iam.rpc__users_by_organization(p_org_id UUID)
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM iam.users WHERE organization_id = p_org_id;
$$;

CREATE OR REPLACE FUNCTION iam.rpc__user_permissions(p_user_id UUID)
RETURNS TABLE(resource TEXT, action TEXT) LANGUAGE SQL STABLE AS $$
    SELECT DISTINCT p.resource, p.action
    FROM iam.user_roles ur
    JOIN iam.role_permissions rp ON rp.role_id = ur.role_id
    JOIN iam.permissions p ON p.id = rp.permission_id
    WHERE ur.user_id = p_user_id
      AND (ur.expires_at IS NULL OR ur.expires_at > NOW());
$$;
