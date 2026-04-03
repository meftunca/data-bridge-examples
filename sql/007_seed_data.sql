-- ============================================================
-- DataBridge V2 — Innovation Hub Demo
-- 007_seed_data.sql
--
-- Sample data for demonstration & testing
-- ============================================================

-- ─── IAM: Organizations ─────────────────────────────────────
INSERT INTO iam.organizations (id, name, slug, description) VALUES
    ('a0000000-0000-0000-0000-000000000001', 'Maple Technologies', 'maple-tech', 'Parent holding company'),
    ('a0000000-0000-0000-0000-000000000002', 'Maple Commerce', 'maple-commerce', 'E-commerce division'),
    ('a0000000-0000-0000-0000-000000000003', 'Maple Logistics', 'maple-logistics', 'Logistics & supply chain');

UPDATE iam.organizations 
SET parent_id = 'a0000000-0000-0000-0000-000000000001' 
WHERE id IN ('a0000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000003');

-- ─── IAM: Users ─────────────────────────────────────────────
INSERT INTO iam.users (id, email, name, display_name, status, auth_provider, organization_id) VALUES
    ('b0000000-0000-0000-0000-000000000001', 'admin@maple.dev', 'Ada Admin', 'Ada', 'active', 'local', 'a0000000-0000-0000-0000-000000000001'),
    ('b0000000-0000-0000-0000-000000000002', 'commerce@maple.dev', 'Carlos Commerce', 'Carlos', 'active', 'google', 'a0000000-0000-0000-0000-000000000002'),
    ('b0000000-0000-0000-0000-000000000003', 'logistics@maple.dev', 'Lara Logistics', 'Lara', 'active', 'local', 'a0000000-0000-0000-0000-000000000003'),
    ('b0000000-0000-0000-0000-000000000004', 'analyst@maple.dev', 'Nadia Analyst', 'Nadia', 'active', 'github', 'a0000000-0000-0000-0000-000000000001'),
    ('b0000000-0000-0000-0000-000000000005', 'viewer@maple.dev', 'Viktor Viewer', 'Viktor', 'pending_verification', 'local', 'a0000000-0000-0000-0000-000000000002');

-- ─── IAM: Roles & Permissions ───────────────────────────────
INSERT INTO iam.roles (id, name, slug, description, organization_id, is_system) VALUES
    ('c0000000-0000-0000-0000-000000000001', 'Super Admin', 'super-admin', 'Full system access', NULL, true),
    ('c0000000-0000-0000-0000-000000000002', 'Org Admin', 'org-admin', 'Organization administrator', NULL, true),
    ('c0000000-0000-0000-0000-000000000003', 'Product Manager', 'product-manager', 'Manage products catalog', 'a0000000-0000-0000-0000-000000000002', false),
    ('c0000000-0000-0000-0000-000000000004', 'Warehouse Manager', 'warehouse-manager', 'Manage warehouses', 'a0000000-0000-0000-0000-000000000003', false),
    ('c0000000-0000-0000-0000-000000000005', 'Analyst', 'analyst', 'View reports and dashboards', NULL, false);

INSERT INTO iam.permissions (id, name, resource, action) VALUES
    ('d0000000-0000-0000-0000-000000000001', 'Full Access', '*', '*'),
    ('d0000000-0000-0000-0000-000000000002', 'Read All', '*', 'read'),
    ('d0000000-0000-0000-0000-000000000003', 'Manage Products', 'catalog.products', '*'),
    ('d0000000-0000-0000-0000-000000000004', 'Manage Orders', 'orders.orders', '*'),
    ('d0000000-0000-0000-0000-000000000005', 'Manage Inventory', 'logistics.inventory', '*'),
    ('d0000000-0000-0000-0000-000000000006', 'View Analytics', 'analytics.*', 'read');

INSERT INTO iam.role_permissions (name, role_id, permission_id) VALUES
    ('Super Admin → Full Access', 'c0000000-0000-0000-0000-000000000001', 'd0000000-0000-0000-0000-000000000001'),
    ('Org Admin → Read All', 'c0000000-0000-0000-0000-000000000002', 'd0000000-0000-0000-0000-000000000002'),
    ('Product Manager → Manage Products', 'c0000000-0000-0000-0000-000000000003', 'd0000000-0000-0000-0000-000000000003'),
    ('Warehouse Manager → Manage Inventory', 'c0000000-0000-0000-0000-000000000004', 'd0000000-0000-0000-0000-000000000005'),
    ('Analyst → View Analytics', 'c0000000-0000-0000-0000-000000000005', 'd0000000-0000-0000-0000-000000000006');

INSERT INTO iam.user_roles (name, user_id, role_id, granted_by) VALUES
    ('Ada as Super Admin', 'b0000000-0000-0000-0000-000000000001', 'c0000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000001'),
    ('Carlos as Product Manager', 'b0000000-0000-0000-0000-000000000002', 'c0000000-0000-0000-0000-000000000003', 'b0000000-0000-0000-0000-000000000001'),
    ('Lara as Warehouse Manager', 'b0000000-0000-0000-0000-000000000003', 'c0000000-0000-0000-0000-000000000004', 'b0000000-0000-0000-0000-000000000001'),
    ('Nadia as Analyst', 'b0000000-0000-0000-0000-000000000004', 'c0000000-0000-0000-0000-000000000005', 'b0000000-0000-0000-0000-000000000001');

-- ─── IAM: Teams ─────────────────────────────────────────────
INSERT INTO iam.teams (id, name, description, organization_id, lead_id, tags) VALUES
    ('e0000000-0000-0000-0000-000000000001', 'Platform Team', 'Core platform development', 'a0000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000001', ARRAY['engineering', 'platform']),
    ('e0000000-0000-0000-0000-000000000002', 'Commerce Team', 'Product & order management', 'a0000000-0000-0000-0000-000000000002', 'b0000000-0000-0000-0000-000000000002', ARRAY['commerce', 'frontend']),
    ('e0000000-0000-0000-0000-000000000003', 'Supply Chain Team', 'Warehouse & logistics operations', 'a0000000-0000-0000-0000-000000000003', 'b0000000-0000-0000-0000-000000000003', ARRAY['logistics', 'ops']);

INSERT INTO iam.team_members (name, team_id, user_id, role) VALUES
    ('Ada in Platform', 'e0000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000001', 'lead'),
    ('Nadia in Platform', 'e0000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000004', 'member'),
    ('Carlos in Commerce', 'e0000000-0000-0000-0000-000000000002', 'b0000000-0000-0000-0000-000000000002', 'lead'),
    ('Lara in Supply Chain', 'e0000000-0000-0000-0000-000000000003', 'b0000000-0000-0000-0000-000000000003', 'lead');

-- ─── IAM: API Keys ──────────────────────────────────────────
INSERT INTO iam.api_keys (name, key_hash, user_id, organization_id, scopes) VALUES
    ('Admin CLI Key', 'hash_placeholder_admin', 'b0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', ARRAY['admin', 'read', 'write']),
    ('Commerce CI Key', 'hash_placeholder_commerce', 'b0000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000002', ARRAY['catalog.read', 'catalog.write', 'orders.read']);

-- ─── Catalog: Brands ────────────────────────────────────────
INSERT INTO catalog.brands (id, name, slug, logo_url, website, organization_id) VALUES
    ('f0000000-0000-0000-0000-000000000001', 'Maple Originals', 'maple-originals', '/logos/maple.svg', 'https://maple.dev', 'a0000000-0000-0000-0000-000000000002'),
    ('f0000000-0000-0000-0000-000000000002', 'TechVault', 'techvault', '/logos/techvault.svg', 'https://techvault.io', 'a0000000-0000-0000-0000-000000000002');

-- ─── Catalog: Categories ────────────────────────────────────
INSERT INTO catalog.categories (id, name, slug, description, parent_id, sort_order) VALUES
    ('10000000-0000-0000-0000-000000000001', 'Electronics', 'electronics', 'Electronic devices', NULL, 1),
    ('10000000-0000-0000-0000-000000000002', 'Laptops', 'laptops', 'Laptop computers', '10000000-0000-0000-0000-000000000001', 1),
    ('10000000-0000-0000-0000-000000000003', 'Smartphones', 'smartphones', 'Mobile phones', '10000000-0000-0000-0000-000000000001', 2),
    ('10000000-0000-0000-0000-000000000004', 'Accessories', 'accessories', 'Device accessories', NULL, 2),
    ('10000000-0000-0000-0000-000000000005', 'Cases', 'cases', 'Protective cases', '10000000-0000-0000-0000-000000000004', 1);

-- ─── Catalog: Products ──────────────────────────────────────
INSERT INTO catalog.products (id, name, slug, sku, description, status, brand_id, category_id, base_price, tags, is_featured, created_by) VALUES
    ('20000000-0000-0000-0000-000000000001', 'ProBook X1', 'probook-x1', 'PBX1-001', 'High-performance laptop for developers', 'active', 'f0000000-0000-0000-0000-000000000001', '10000000-0000-0000-0000-000000000002', 1299.99, ARRAY['featured', 'developer', 'laptop'], true, 'b0000000-0000-0000-0000-000000000002'),
    ('20000000-0000-0000-0000-000000000002', 'SmartPhone Z', 'smartphone-z', 'SPZ-001', 'Flagship smartphone with AI features', 'active', 'f0000000-0000-0000-0000-000000000001', '10000000-0000-0000-0000-000000000003', 899.99, ARRAY['flagship', 'ai', '5g'], true, 'b0000000-0000-0000-0000-000000000002'),
    ('20000000-0000-0000-0000-000000000003', 'USB-C Hub Pro', 'usbc-hub-pro', 'HUB-001', '7-in-1 USB-C hub', 'active', 'f0000000-0000-0000-0000-000000000002', '10000000-0000-0000-0000-000000000004', 79.99, ARRAY['accessory', 'usb-c'], false, 'b0000000-0000-0000-0000-000000000002'),
    ('20000000-0000-0000-0000-000000000004', 'ProBook Case', 'probook-case', 'CASE-PB1', 'Protective case for ProBook X1', 'active', 'f0000000-0000-0000-0000-000000000002', '10000000-0000-0000-0000-000000000005', 49.99, ARRAY['case', 'protection'], false, 'b0000000-0000-0000-0000-000000000002'),
    ('20000000-0000-0000-0000-000000000005', 'Budget Laptop', 'budget-laptop', 'BL-001', 'Affordable everyday laptop', 'draft', 'f0000000-0000-0000-0000-000000000002', '10000000-0000-0000-0000-000000000002', 499.99, ARRAY['budget', 'everyday'], false, 'b0000000-0000-0000-0000-000000000002');

-- ─── Catalog: Product Variants ──────────────────────────────
INSERT INTO catalog.product_variants (id, name, product_id, sku, price_override, attributes, stock_quantity) VALUES
    ('30000000-0000-0000-0000-000000000001', 'ProBook X1 - 16GB/512GB', '20000000-0000-0000-0000-000000000001', 'PBX1-16-512', 1299.99, '{"ram": "16GB", "storage": "512GB", "color": "silver"}', 50),
    ('30000000-0000-0000-0000-000000000002', 'ProBook X1 - 32GB/1TB', '20000000-0000-0000-0000-000000000001', 'PBX1-32-1T', 1599.99, '{"ram": "32GB", "storage": "1TB", "color": "space gray"}', 25),
    ('30000000-0000-0000-0000-000000000003', 'SmartPhone Z - 128GB Black', '20000000-0000-0000-0000-000000000002', 'SPZ-128-BLK', NULL, '{"storage": "128GB", "color": "midnight black"}', 100),
    ('30000000-0000-0000-0000-000000000004', 'SmartPhone Z - 256GB Blue', '20000000-0000-0000-0000-000000000002', 'SPZ-256-BLU', 949.99, '{"storage": "256GB", "color": "ocean blue"}', 75);

-- ─── Catalog: Collections ───────────────────────────────────
INSERT INTO catalog.collections (id, name, slug, description, created_by) VALUES
    ('40000000-0000-0000-0000-000000000001', 'Summer 2026 Collection', 'summer-2026', 'Hot picks for summer', 'b0000000-0000-0000-0000-000000000002'),
    ('40000000-0000-0000-0000-000000000002', 'Developer Essentials', 'dev-essentials', 'Must-have tools for developers', 'b0000000-0000-0000-0000-000000000002');

INSERT INTO catalog.collection_products (name, collection_id, product_id, sort_order) VALUES
    ('ProBook in Summer', '40000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000001', 1),
    ('SmartPhone in Summer', '40000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000002', 2),
    ('ProBook for Devs', '40000000-0000-0000-0000-000000000002', '20000000-0000-0000-0000-000000000001', 1),
    ('USB-C Hub for Devs', '40000000-0000-0000-0000-000000000002', '20000000-0000-0000-0000-000000000003', 2);

-- ─── Catalog: Reviews ───────────────────────────────────────
INSERT INTO catalog.product_reviews (name, product_id, user_id, rating, title, body, is_verified) VALUES
    ('ProBook Review by Ada', '20000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000001', 5, 'Best dev laptop', 'Perfect for Go & Docker workloads', true),
    ('SmartPhone Review by Carlos', '20000000-0000-0000-0000-000000000002', 'b0000000-0000-0000-0000-000000000002', 4, 'Great phone', 'Amazing camera and performance', true),
    ('USB-C Review by Nadia', '20000000-0000-0000-0000-000000000003', 'b0000000-0000-0000-0000-000000000004', 3, 'Decent hub', 'Works well but gets warm', false);

-- ─── Catalog: Tags ──────────────────────────────────────────
INSERT INTO catalog.tags (id, name, slug) VALUES
    ('50000000-0000-0000-0000-000000000001', 'New Arrival', 'new-arrival'),
    ('50000000-0000-0000-0000-000000000002', 'Best Seller', 'best-seller'),
    ('50000000-0000-0000-0000-000000000003', 'On Sale', 'on-sale');

INSERT INTO catalog.product_tags (product_id, tag_id, name) VALUES
    ('20000000-0000-0000-0000-000000000001', '50000000-0000-0000-0000-000000000001', 'ProBook New Arrival'),
    ('20000000-0000-0000-0000-000000000001', '50000000-0000-0000-0000-000000000002', 'ProBook Best Seller'),
    ('20000000-0000-0000-0000-000000000002', '50000000-0000-0000-0000-000000000001', 'SmartPhone New Arrival');

-- ─── Orders: Customers ──────────────────────────────────────
INSERT INTO orders.customers (id, name, user_id, organization_id, tier, loyalty_points) VALUES
    ('60000000-0000-0000-0000-000000000001', 'Ada Customer Account', 'b0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'premium', 1500),
    ('60000000-0000-0000-0000-000000000002', 'Viktor Customer Account', 'b0000000-0000-0000-0000-000000000005', 'a0000000-0000-0000-0000-000000000002', 'standard', 200);

-- ─── Orders: Coupons ────────────────────────────────────────
INSERT INTO orders.coupons (id, name, code, discount_type, discount_value, min_order_value) VALUES
    ('65000000-0000-0000-0000-000000000001', 'Welcome 10%', 'WELCOME10', 'percentage', 10.00, 50.00),
    ('65000000-0000-0000-0000-000000000002', 'Flat $20 Off', 'FLAT20', 'fixed', 20.00, 100.00);

-- ─── Orders: Orders ─────────────────────────────────────────
INSERT INTO orders.orders (id, name, order_number, customer_id, status, subtotal, tax_amount, shipping_amount, total, coupon_id, created_by) VALUES
    ('70000000-0000-0000-0000-000000000001', 'Order #1001', 'ORD-2026-001', '60000000-0000-0000-0000-000000000001', 'delivered', 1379.98, 110.40, 15.00, 1505.38, '65000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000001'),
    ('70000000-0000-0000-0000-000000000002', 'Order #1002', 'ORD-2026-002', '60000000-0000-0000-0000-000000000002', 'processing', 899.99, 72.00, 0.00, 971.99, NULL, 'b0000000-0000-0000-0000-000000000005'),
    ('70000000-0000-0000-0000-000000000003', 'Order #1003', 'ORD-2026-003', '60000000-0000-0000-0000-000000000001', 'pending', 79.99, 6.40, 5.00, 91.39, NULL, 'b0000000-0000-0000-0000-000000000001');

-- ─── Orders: Order Items ────────────────────────────────────
INSERT INTO orders.order_items (name, order_id, product_id, variant_id, quantity, unit_price, total_price) VALUES
    ('ProBook X1 16GB', '70000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000001', '30000000-0000-0000-0000-000000000001', 1, 1299.99, 1299.99),
    ('USB-C Hub', '70000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000003', NULL, 1, 79.99, 79.99),
    ('SmartPhone Z 256GB', '70000000-0000-0000-0000-000000000002', '20000000-0000-0000-0000-000000000002', '30000000-0000-0000-0000-000000000004', 1, 949.99, 949.99),
    ('USB-C Hub Pro', '70000000-0000-0000-0000-000000000003', '20000000-0000-0000-0000-000000000003', NULL, 1, 79.99, 79.99);

-- ─── Orders: Payments ───────────────────────────────────────
INSERT INTO orders.payments (name, order_id, amount, method, status, provider_ref) VALUES
    ('Payment for ORD-001', '70000000-0000-0000-0000-000000000001', 1505.38, 'credit_card', 'captured', 'stripe_pi_001'),
    ('Payment for ORD-002', '70000000-0000-0000-0000-000000000002', 971.99, 'wallet', 'authorized', 'wallet_tx_002');

-- ─── Logistics: Warehouses ──────────────────────────────────
INSERT INTO logistics.warehouses (id, name, code, organization_id, manager_id, capacity) VALUES
    ('80000000-0000-0000-0000-000000000001', 'Istanbul Distribution Center', 'IST-DC', 'a0000000-0000-0000-0000-000000000003', 'b0000000-0000-0000-0000-000000000003', 10000),
    ('80000000-0000-0000-0000-000000000002', 'Berlin Fulfillment Hub', 'BER-FH', 'a0000000-0000-0000-0000-000000000003', 'b0000000-0000-0000-0000-000000000003', 5000);

-- ─── Logistics: Storage Zones & Bins ────────────────────────
INSERT INTO logistics.storage_zones (id, name, warehouse_id, zone_code, zone_type) VALUES
    ('81000000-0000-0000-0000-000000000001', 'Electronics Zone A', '80000000-0000-0000-0000-000000000001', 'EL-A', 'climate_controlled'),
    ('81000000-0000-0000-0000-000000000002', 'General Zone B', '80000000-0000-0000-0000-000000000001', 'GN-B', 'general'),
    ('81000000-0000-0000-0000-000000000003', 'Electronics Zone', '80000000-0000-0000-0000-000000000002', 'EL-A', 'climate_controlled');

INSERT INTO logistics.storage_bins (name, zone_id, bin_code, max_capacity) VALUES
    ('Bin A-001', '81000000-0000-0000-0000-000000000001', 'A-001', 200),
    ('Bin A-002', '81000000-0000-0000-0000-000000000001', 'A-002', 200),
    ('Bin B-001', '81000000-0000-0000-0000-000000000002', 'B-001', 500);

-- ─── Logistics: Inventory ───────────────────────────────────
INSERT INTO logistics.inventory (name, product_id, warehouse_id, quantity, reserved, reorder_level) VALUES
    ('ProBook X1 @ Istanbul', '20000000-0000-0000-0000-000000000001', '80000000-0000-0000-0000-000000000001', 45, 5, 20),
    ('SmartPhone Z @ Istanbul', '20000000-0000-0000-0000-000000000002', '80000000-0000-0000-0000-000000000001', 120, 10, 30),
    ('USB-C Hub @ Istanbul', '20000000-0000-0000-0000-000000000003', '80000000-0000-0000-0000-000000000001', 200, 0, 50),
    ('ProBook X1 @ Berlin', '20000000-0000-0000-0000-000000000001', '80000000-0000-0000-0000-000000000002', 30, 2, 15),
    ('SmartPhone Z @ Berlin', '20000000-0000-0000-0000-000000000002', '80000000-0000-0000-0000-000000000002', 80, 5, 25);

-- ─── Logistics: Suppliers ───────────────────────────────────
INSERT INTO logistics.suppliers (id, name, code, contact_name, contact_email) VALUES
    ('90000000-0000-0000-0000-000000000001', 'TechSource Global', 'TSG', 'John Tech', 'john@techsource.io'),
    ('90000000-0000-0000-0000-000000000002', 'AccessoryWorld', 'ACW', 'Sarah World', 'sarah@accworld.io');

-- ─── Analytics: Dashboards ──────────────────────────────────
INSERT INTO analytics.dashboards (id, name, description, owner_id, organization_id, is_public) VALUES
    ('a1000000-0000-0000-0000-000000000001', 'Executive Overview', 'Company-wide KPI dashboard', 'b0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', true),
    ('a1000000-0000-0000-0000-000000000002', 'Commerce Insights', 'Sales and product performance', 'b0000000-0000-0000-0000-000000000004', 'a0000000-0000-0000-0000-000000000002', false);

INSERT INTO analytics.dashboard_widgets (name, dashboard_id, widget_type, config, data_source) VALUES
    ('Revenue Chart', 'a1000000-0000-0000-0000-000000000001', 'line_chart', '{"metric": "revenue", "period": "monthly"}', 'orders.orders'),
    ('Inventory Status', 'a1000000-0000-0000-0000-000000000001', 'bar_chart', '{"metric": "stock_level", "group_by": "warehouse"}', 'logistics.inventory'),
    ('Top Products', 'a1000000-0000-0000-0000-000000000002', 'table', '{"limit": 10, "sort": "revenue_desc"}', 'catalog.products');

-- ─── Analytics: Reports ─────────────────────────────────────
INSERT INTO analytics.reports (id, name, report_type, owner_id, organization_id, query_config) VALUES
    ('a2000000-0000-0000-0000-000000000001', 'Monthly Revenue Report', 'monthly', 'b0000000-0000-0000-0000-000000000004', 'a0000000-0000-0000-0000-000000000001', '{"schema": "orders", "table": "orders", "aggregation": "SUM(total)", "group_by": "month"}'),
    ('a2000000-0000-0000-0000-000000000002', 'Inventory Health Report', 'weekly', 'b0000000-0000-0000-0000-000000000003', 'a0000000-0000-0000-0000-000000000003', '{"schema": "logistics", "table": "inventory", "filter": "quantity <= reorder_level"}');

-- ─── Analytics: Audit Logs ──────────────────────────────────
INSERT INTO analytics.audit_logs (name, user_id, action, resource_type, resource_id, severity) VALUES
    ('Product created', 'b0000000-0000-0000-0000-000000000002', 'create', 'catalog.products', '20000000-0000-0000-0000-000000000001', 'info'),
    ('Order placed', 'b0000000-0000-0000-0000-000000000001', 'create', 'orders.orders', '70000000-0000-0000-0000-000000000001', 'info'),
    ('Permission granted', 'b0000000-0000-0000-0000-000000000001', 'update', 'iam.user_roles', 'b0000000-0000-0000-0000-000000000002', 'warning');

-- ─── Analytics: Events ──────────────────────────────────────
INSERT INTO analytics.events (name, event_type, source_schema, source_table, source_id, actor_id, severity) VALUES
    ('Product Launch Event', 'product.created', 'catalog', 'products', '20000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000002', 'info'),
    ('Order Completion Event', 'order.delivered', 'orders', 'orders', '70000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000001', 'info'),
    ('Low Stock Alert Event', 'inventory.low_stock', 'logistics', 'inventory', '20000000-0000-0000-0000-000000000001', NULL, 'warning');

-- ─── Analytics: Notifications ───────────────────────────────
INSERT INTO analytics.notifications (name, user_id, title, message, channel) VALUES
    ('Low Stock Notification', 'b0000000-0000-0000-0000-000000000003', 'Low Stock Alert', 'ProBook X1 stock is below reorder level at Istanbul DC', 'in_app'),
    ('Order Completed', 'b0000000-0000-0000-0000-000000000001', 'Order Delivered', 'Your order ORD-2026-001 has been delivered', 'email');

-- ─── Analytics: Alert Rules ─────────────────────────────────
INSERT INTO analytics.alert_rules (name, description, owner_id, condition, action, severity) VALUES
    ('Low Stock Alert', 'Trigger when stock goes below reorder level', 'b0000000-0000-0000-0000-000000000003', '{"table": "logistics.inventory", "condition": "quantity <= reorder_level"}', '{"notify": ["warehouse-manager"], "channel": "slack"}', 'warning'),
    ('High Value Order', 'Trigger for orders above $1000', 'b0000000-0000-0000-0000-000000000004', '{"table": "orders.orders", "condition": "total > 1000"}', '{"notify": ["analyst"], "channel": "email"}', 'info');

-- ─── Analytics: Metrics ─────────────────────────────────────
INSERT INTO analytics.metrics (name, metric_key, value, dimensions) VALUES
    ('Daily Revenue', 'revenue.daily', 2477.37, '{"date": "2026-04-02", "currency": "USD"}'),
    ('Daily Orders', 'orders.daily.count', 3, '{"date": "2026-04-02"}'),
    ('Inventory Turnover', 'inventory.turnover', 2.35, '{"warehouse": "IST-DC", "period": "monthly"}');
