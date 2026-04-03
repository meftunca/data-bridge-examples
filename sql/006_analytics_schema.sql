-- ============================================================
-- DataBridge V2 — Innovation Hub Demo
-- 006_analytics_schema.sql
--
-- Audit, Reporting & Real-time Events Schema
-- Features demonstrated:
--   ✓ Cross-schema FKs from all schemas (iam, catalog, orders, logistics)
--   ✓ ENUM usage (event_severity, report_type)
--   ✓ JSONB for flexible payloads
--   ✓ INET type for IP tracking
--   ✓ Views (read-only aggregations)
--   ✓ Materialized Views
--   ✓ Outbox pattern table
--   ✓ Time-series data patterns
--   ✓ Parent-child: reports → report_widgets
-- ============================================================

-- ─── Audit Logs ─────────────────────────────────────────────
CREATE TABLE analytics.audit_logs (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    user_id         UUID REFERENCES iam.users(id),
    action          TEXT NOT NULL DEFAULT '',
    resource_type   TEXT NOT NULL DEFAULT '',
    resource_id     TEXT NOT NULL DEFAULT '',
    severity        analytics.event_severity NOT NULL DEFAULT 'info',
    ip_address      INET,
    user_agent      TEXT NOT NULL DEFAULT '',
    old_values      JSONB,
    new_values      JSONB,
    metadata        JSONB NOT NULL DEFAULT '{}',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Event Stream ───────────────────────────────────────────
CREATE TABLE analytics.events (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    event_type      TEXT NOT NULL DEFAULT '',
    source_schema   TEXT NOT NULL DEFAULT '',
    source_table    TEXT NOT NULL DEFAULT '',
    source_id       TEXT NOT NULL DEFAULT '',
    actor_id        UUID REFERENCES iam.users(id),
    payload         JSONB NOT NULL DEFAULT '{}',
    severity        analytics.event_severity NOT NULL DEFAULT 'info',
    processed       BOOLEAN NOT NULL DEFAULT false,
    processed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Dashboards ─────────────────────────────────────────────
CREATE TABLE analytics.dashboards (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    description     TEXT NOT NULL DEFAULT '',
    owner_id        UUID NOT NULL REFERENCES iam.users(id),
    organization_id UUID REFERENCES iam.organizations(id),
    is_public       BOOLEAN NOT NULL DEFAULT false,
    layout          JSONB NOT NULL DEFAULT '{}',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Dashboard Widgets ──────────────────────────────────────
CREATE TABLE analytics.dashboard_widgets (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    dashboard_id    UUID NOT NULL REFERENCES analytics.dashboards(id),
    widget_type     TEXT NOT NULL DEFAULT 'chart',
    config          JSONB NOT NULL DEFAULT '{}',
    position        JSONB NOT NULL DEFAULT '{"x": 0, "y": 0, "w": 6, "h": 4}',
    data_source     TEXT NOT NULL DEFAULT '',
    refresh_interval INTEGER NOT NULL DEFAULT 300,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Reports ────────────────────────────────────────────────
CREATE TABLE analytics.reports (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    description     TEXT NOT NULL DEFAULT '',
    report_type     analytics.report_type NOT NULL DEFAULT 'monthly',
    owner_id        UUID NOT NULL REFERENCES iam.users(id),
    organization_id UUID REFERENCES iam.organizations(id),
    query_config    JSONB NOT NULL DEFAULT '{}',
    schedule        JSONB NOT NULL DEFAULT '{}',
    is_active       BOOLEAN NOT NULL DEFAULT true,
    last_run_at     TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Report Executions ──────────────────────────────────────
CREATE TABLE analytics.report_executions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    report_id       UUID NOT NULL REFERENCES analytics.reports(id),
    status          TEXT NOT NULL DEFAULT 'pending',
    result_data     JSONB,
    row_count       INTEGER NOT NULL DEFAULT 0,
    duration_ms     INTEGER NOT NULL DEFAULT 0,
    error_message   TEXT NOT NULL DEFAULT '',
    executed_by     UUID REFERENCES iam.users(id),
    started_at      TIMESTAMPTZ,
    completed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Notifications ──────────────────────────────────────────
CREATE TABLE analytics.notifications (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    user_id         UUID NOT NULL REFERENCES iam.users(id),
    title           TEXT NOT NULL DEFAULT '',
    message         TEXT NOT NULL DEFAULT '',
    channel         TEXT NOT NULL DEFAULT 'in_app',
    is_read         BOOLEAN NOT NULL DEFAULT false,
    action_url      TEXT NOT NULL DEFAULT '',
    metadata        JSONB NOT NULL DEFAULT '{}',
    read_at         TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Alert Rules ────────────────────────────────────────────
CREATE TABLE analytics.alert_rules (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    description     TEXT NOT NULL DEFAULT '',
    owner_id        UUID NOT NULL REFERENCES iam.users(id),
    condition       JSONB NOT NULL DEFAULT '{}',
    action          JSONB NOT NULL DEFAULT '{}',
    severity        analytics.event_severity NOT NULL DEFAULT 'warning',
    is_active       BOOLEAN NOT NULL DEFAULT true,
    last_triggered  TIMESTAMPTZ,
    cooldown_minutes INTEGER NOT NULL DEFAULT 60,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Alert History ──────────────────────────────────────────
CREATE TABLE analytics.alert_history (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL DEFAULT '',
    alert_rule_id   UUID NOT NULL REFERENCES analytics.alert_rules(id),
    triggered_value JSONB NOT NULL DEFAULT '{}',
    resolved        BOOLEAN NOT NULL DEFAULT false,
    resolved_at     TIMESTAMPTZ,
    resolved_by     UUID REFERENCES iam.users(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Metrics (time-series) ──────────────────────────────────
CREATE TABLE analytics.metrics (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    metric_key      TEXT NOT NULL,
    value           NUMERIC(18,4) NOT NULL DEFAULT 0,
    dimensions      JSONB NOT NULL DEFAULT '{}',
    recorded_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Outbox Events (transactional outbox pattern) ───────────
CREATE TABLE IF NOT EXISTS outbox_events (
    id              BIGSERIAL PRIMARY KEY,
    aggregate_type  TEXT NOT NULL,
    aggregate_id    TEXT NOT NULL DEFAULT '',
    event_type      TEXT NOT NULL,
    payload         JSONB NOT NULL DEFAULT '{}',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    published_at    TIMESTAMPTZ,
    retry_count     INT NOT NULL DEFAULT 0
);

-- ─── Views ──────────────────────────────────────────────────
CREATE VIEW analytics.recent_events AS
SELECT id, name, event_type, source_schema, source_table, actor_id, severity, created_at
FROM analytics.events
WHERE created_at >= NOW() - INTERVAL '24 hours'
ORDER BY created_at DESC;

CREATE VIEW analytics.unread_notifications AS
SELECT id, name, user_id, title, message, channel, action_url, created_at
FROM analytics.notifications
WHERE is_read = false;

-- ─── Materialized Views ─────────────────────────────────────
CREATE MATERIALIZED VIEW analytics.daily_event_counts AS
SELECT 
    date_trunc('day', created_at)::DATE AS event_date,
    event_type,
    severity::TEXT AS severity,
    COUNT(*)::BIGINT AS event_count
FROM analytics.events
GROUP BY 1, 2, 3;

CREATE MATERIALIZED VIEW analytics.user_activity_summary AS
SELECT 
    user_id,
    COUNT(*)::BIGINT AS total_actions,
    MAX(created_at) AS last_activity,
    COUNT(DISTINCT action)::BIGINT AS unique_actions
FROM analytics.audit_logs
WHERE user_id IS NOT NULL
GROUP BY user_id;

-- ─── RPC Functions ──────────────────────────────────────────
CREATE OR REPLACE FUNCTION analytics.rpc__event_count_by_severity(p_severity TEXT)
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM analytics.events 
    WHERE severity = p_severity::analytics.event_severity;
$$;

CREATE OR REPLACE FUNCTION analytics.rpc__unread_notification_count(p_user_id UUID)
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM analytics.notifications 
    WHERE user_id = p_user_id AND is_read = false;
$$;

CREATE OR REPLACE FUNCTION analytics.rpc__dashboard_count()
RETURNS INTEGER LANGUAGE SQL STABLE AS $$
    SELECT COUNT(*)::INTEGER FROM analytics.dashboards;
$$;
