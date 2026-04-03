# Missing Comments Report

> Generated — 2026-04-03 12:39:58

**Total: 448 columns without comments**

## Schema: `iam`

### `api_keys` (9 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.api_keys.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.api_keys.name IS '...';` |
| 3 | `key_hash` | `COMMENT ON COLUMN iam.api_keys.key_hash IS '...';` |
| 4 | `user_id` | `COMMENT ON COLUMN iam.api_keys.user_id IS '...';` |
| 5 | `organization_id` | `COMMENT ON COLUMN iam.api_keys.organization_id IS '...';` |
| 6 | `scopes` | `COMMENT ON COLUMN iam.api_keys.scopes IS '...';` |
| 7 | `is_active` | `COMMENT ON COLUMN iam.api_keys.is_active IS '...';` |
| 8 | `last_used_at` | `COMMENT ON COLUMN iam.api_keys.last_used_at IS '...';` |
| 9 | `expires_at` | `COMMENT ON COLUMN iam.api_keys.expires_at IS '...';` |

### `invitations` (9 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.invitations.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.invitations.name IS '...';` |
| 3 | `email` | `COMMENT ON COLUMN iam.invitations.email IS '...';` |
| 4 | `organization_id` | `COMMENT ON COLUMN iam.invitations.organization_id IS '...';` |
| 5 | `invited_by` | `COMMENT ON COLUMN iam.invitations.invited_by IS '...';` |
| 6 | `role_id` | `COMMENT ON COLUMN iam.invitations.role_id IS '...';` |
| 7 | `token` | `COMMENT ON COLUMN iam.invitations.token IS '...';` |
| 8 | `accepted_at` | `COMMENT ON COLUMN iam.invitations.accepted_at IS '...';` |
| 9 | `expires_at` | `COMMENT ON COLUMN iam.invitations.expires_at IS '...';` |

### `organizations` (8 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.organizations.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.organizations.name IS '...';` |
| 3 | `slug` | `COMMENT ON COLUMN iam.organizations.slug IS '...';` |
| 4 | `description` | `COMMENT ON COLUMN iam.organizations.description IS '...';` |
| 5 | `logo_url` | `COMMENT ON COLUMN iam.organizations.logo_url IS '...';` |
| 6 | `parent_id` | `COMMENT ON COLUMN iam.organizations.parent_id IS '...';` |
| 7 | `settings` | `COMMENT ON COLUMN iam.organizations.settings IS '...';` |
| 8 | `is_active` | `COMMENT ON COLUMN iam.organizations.is_active IS '...';` |

### `permissions` (5 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.permissions.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.permissions.name IS '...';` |
| 3 | `resource` | `COMMENT ON COLUMN iam.permissions.resource IS '...';` |
| 4 | `action` | `COMMENT ON COLUMN iam.permissions.action IS '...';` |
| 5 | `description` | `COMMENT ON COLUMN iam.permissions.description IS '...';` |

### `role_permissions` (4 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.role_permissions.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.role_permissions.name IS '...';` |
| 3 | `role_id` | `COMMENT ON COLUMN iam.role_permissions.role_id IS '...';` |
| 4 | `permission_id` | `COMMENT ON COLUMN iam.role_permissions.permission_id IS '...';` |

### `roles` (6 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.roles.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.roles.name IS '...';` |
| 3 | `slug` | `COMMENT ON COLUMN iam.roles.slug IS '...';` |
| 4 | `description` | `COMMENT ON COLUMN iam.roles.description IS '...';` |
| 5 | `organization_id` | `COMMENT ON COLUMN iam.roles.organization_id IS '...';` |
| 6 | `is_system` | `COMMENT ON COLUMN iam.roles.is_system IS '...';` |

### `sessions` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.sessions.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.sessions.name IS '...';` |
| 3 | `user_id` | `COMMENT ON COLUMN iam.sessions.user_id IS '...';` |
| 4 | `ip_address` | `COMMENT ON COLUMN iam.sessions.ip_address IS '...';` |
| 5 | `user_agent` | `COMMENT ON COLUMN iam.sessions.user_agent IS '...';` |
| 6 | `is_active` | `COMMENT ON COLUMN iam.sessions.is_active IS '...';` |
| 7 | `expires_at` | `COMMENT ON COLUMN iam.sessions.expires_at IS '...';` |

### `team_members` (6 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.team_members.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.team_members.name IS '...';` |
| 3 | `team_id` | `COMMENT ON COLUMN iam.team_members.team_id IS '...';` |
| 4 | `user_id` | `COMMENT ON COLUMN iam.team_members.user_id IS '...';` |
| 5 | `role` | `COMMENT ON COLUMN iam.team_members.role IS '...';` |
| 6 | `joined_at` | `COMMENT ON COLUMN iam.team_members.joined_at IS '...';` |

### `teams` (6 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.teams.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.teams.name IS '...';` |
| 3 | `description` | `COMMENT ON COLUMN iam.teams.description IS '...';` |
| 4 | `organization_id` | `COMMENT ON COLUMN iam.teams.organization_id IS '...';` |
| 5 | `lead_id` | `COMMENT ON COLUMN iam.teams.lead_id IS '...';` |
| 6 | `tags` | `COMMENT ON COLUMN iam.teams.tags IS '...';` |

### `user_roles` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.user_roles.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN iam.user_roles.name IS '...';` |
| 3 | `user_id` | `COMMENT ON COLUMN iam.user_roles.user_id IS '...';` |
| 4 | `role_id` | `COMMENT ON COLUMN iam.user_roles.role_id IS '...';` |
| 5 | `granted_by` | `COMMENT ON COLUMN iam.user_roles.granted_by IS '...';` |
| 6 | `granted_at` | `COMMENT ON COLUMN iam.user_roles.granted_at IS '...';` |
| 7 | `expires_at` | `COMMENT ON COLUMN iam.user_roles.expires_at IS '...';` |

### `users` (12 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN iam.users.id IS '...';` |
| 2 | `email` | `COMMENT ON COLUMN iam.users.email IS '...';` |
| 3 | `name` | `COMMENT ON COLUMN iam.users.name IS '...';` |
| 4 | `display_name` | `COMMENT ON COLUMN iam.users.display_name IS '...';` |
| 5 | `avatar_url` | `COMMENT ON COLUMN iam.users.avatar_url IS '...';` |
| 6 | `phone` | `COMMENT ON COLUMN iam.users.phone IS '...';` |
| 7 | `status` | `COMMENT ON COLUMN iam.users.status IS '...';` |
| 8 | `auth_provider` | `COMMENT ON COLUMN iam.users.auth_provider IS '...';` |
| 9 | `organization_id` | `COMMENT ON COLUMN iam.users.organization_id IS '...';` |
| 10 | `metadata` | `COMMENT ON COLUMN iam.users.metadata IS '...';` |
| 11 | `last_login_at` | `COMMENT ON COLUMN iam.users.last_login_at IS '...';` |
| 12 | `email_verified` | `COMMENT ON COLUMN iam.users.email_verified IS '...';` |

## Schema: `catalog`

### `brands` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.brands.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.brands.name IS '...';` |
| 3 | `slug` | `COMMENT ON COLUMN catalog.brands.slug IS '...';` |
| 4 | `logo_url` | `COMMENT ON COLUMN catalog.brands.logo_url IS '...';` |
| 5 | `website` | `COMMENT ON COLUMN catalog.brands.website IS '...';` |
| 6 | `description` | `COMMENT ON COLUMN catalog.brands.description IS '...';` |
| 7 | `organization_id` | `COMMENT ON COLUMN catalog.brands.organization_id IS '...';` |

### `categories` (8 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.categories.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.categories.name IS '...';` |
| 3 | `slug` | `COMMENT ON COLUMN catalog.categories.slug IS '...';` |
| 4 | `description` | `COMMENT ON COLUMN catalog.categories.description IS '...';` |
| 5 | `parent_id` | `COMMENT ON COLUMN catalog.categories.parent_id IS '...';` |
| 6 | `icon` | `COMMENT ON COLUMN catalog.categories.icon IS '...';` |
| 7 | `sort_order` | `COMMENT ON COLUMN catalog.categories.sort_order IS '...';` |
| 8 | `is_active` | `COMMENT ON COLUMN catalog.categories.is_active IS '...';` |

### `collection_products` (5 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.collection_products.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.collection_products.name IS '...';` |
| 3 | `collection_id` | `COMMENT ON COLUMN catalog.collection_products.collection_id IS '...';` |
| 4 | `product_id` | `COMMENT ON COLUMN catalog.collection_products.product_id IS '...';` |
| 5 | `sort_order` | `COMMENT ON COLUMN catalog.collection_products.sort_order IS '...';` |

### `collections` (8 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.collections.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.collections.name IS '...';` |
| 3 | `slug` | `COMMENT ON COLUMN catalog.collections.slug IS '...';` |
| 4 | `description` | `COMMENT ON COLUMN catalog.collections.description IS '...';` |
| 5 | `cover_image_url` | `COMMENT ON COLUMN catalog.collections.cover_image_url IS '...';` |
| 6 | `is_active` | `COMMENT ON COLUMN catalog.collections.is_active IS '...';` |
| 7 | `start_date` | `COMMENT ON COLUMN catalog.collections.start_date IS '...';` |
| 8 | `end_date` | `COMMENT ON COLUMN catalog.collections.end_date IS '...';` |

### `price_history` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.price_history.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.price_history.name IS '...';` |
| 3 | `product_id` | `COMMENT ON COLUMN catalog.price_history.product_id IS '...';` |
| 4 | `old_price` | `COMMENT ON COLUMN catalog.price_history.old_price IS '...';` |
| 5 | `new_price` | `COMMENT ON COLUMN catalog.price_history.new_price IS '...';` |
| 6 | `changed_by` | `COMMENT ON COLUMN catalog.price_history.changed_by IS '...';` |
| 7 | `reason` | `COMMENT ON COLUMN catalog.price_history.reason IS '...';` |

### `product_media` (9 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.product_media.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.product_media.name IS '...';` |
| 3 | `product_id` | `COMMENT ON COLUMN catalog.product_media.product_id IS '...';` |
| 4 | `media_type` | `COMMENT ON COLUMN catalog.product_media.media_type IS '...';` |
| 5 | `url` | `COMMENT ON COLUMN catalog.product_media.url IS '...';` |
| 6 | `alt_text` | `COMMENT ON COLUMN catalog.product_media.alt_text IS '...';` |
| 7 | `sort_order` | `COMMENT ON COLUMN catalog.product_media.sort_order IS '...';` |
| 8 | `is_primary` | `COMMENT ON COLUMN catalog.product_media.is_primary IS '...';` |
| 9 | `metadata` | `COMMENT ON COLUMN catalog.product_media.metadata IS '...';` |

### `product_reviews` (9 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.product_reviews.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.product_reviews.name IS '...';` |
| 3 | `product_id` | `COMMENT ON COLUMN catalog.product_reviews.product_id IS '...';` |
| 4 | `user_id` | `COMMENT ON COLUMN catalog.product_reviews.user_id IS '...';` |
| 5 | `rating` | `COMMENT ON COLUMN catalog.product_reviews.rating IS '...';` |
| 6 | `title` | `COMMENT ON COLUMN catalog.product_reviews.title IS '...';` |
| 7 | `body` | `COMMENT ON COLUMN catalog.product_reviews.body IS '...';` |
| 8 | `is_verified` | `COMMENT ON COLUMN catalog.product_reviews.is_verified IS '...';` |
| 9 | `helpful_count` | `COMMENT ON COLUMN catalog.product_reviews.helpful_count IS '...';` |

### `product_tags` (3 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `product_id` | `COMMENT ON COLUMN catalog.product_tags.product_id IS '...';` |
| 2 | `tag_id` | `COMMENT ON COLUMN catalog.product_tags.tag_id IS '...';` |
| 3 | `name` | `COMMENT ON COLUMN catalog.product_tags.name IS '...';` |

### `product_variants` (8 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.product_variants.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.product_variants.name IS '...';` |
| 3 | `product_id` | `COMMENT ON COLUMN catalog.product_variants.product_id IS '...';` |
| 4 | `sku` | `COMMENT ON COLUMN catalog.product_variants.sku IS '...';` |
| 5 | `price_override` | `COMMENT ON COLUMN catalog.product_variants.price_override IS '...';` |
| 6 | `attributes` | `COMMENT ON COLUMN catalog.product_variants.attributes IS '...';` |
| 7 | `is_active` | `COMMENT ON COLUMN catalog.product_variants.is_active IS '...';` |
| 8 | `stock_quantity` | `COMMENT ON COLUMN catalog.product_variants.stock_quantity IS '...';` |

### `products` (17 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.products.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.products.name IS '...';` |
| 3 | `slug` | `COMMENT ON COLUMN catalog.products.slug IS '...';` |
| 4 | `sku` | `COMMENT ON COLUMN catalog.products.sku IS '...';` |
| 5 | `description` | `COMMENT ON COLUMN catalog.products.description IS '...';` |
| 6 | `short_description` | `COMMENT ON COLUMN catalog.products.short_description IS '...';` |
| 7 | `status` | `COMMENT ON COLUMN catalog.products.status IS '...';` |
| 8 | `brand_id` | `COMMENT ON COLUMN catalog.products.brand_id IS '...';` |
| 9 | `category_id` | `COMMENT ON COLUMN catalog.products.category_id IS '...';` |
| 10 | `base_price` | `COMMENT ON COLUMN catalog.products.base_price IS '...';` |
| 11 | `currency` | `COMMENT ON COLUMN catalog.products.currency IS '...';` |
| 12 | `weight_kg` | `COMMENT ON COLUMN catalog.products.weight_kg IS '...';` |
| 13 | `dimensions_cm` | `COMMENT ON COLUMN catalog.products.dimensions_cm IS '...';` |
| 14 | `attributes` | `COMMENT ON COLUMN catalog.products.attributes IS '...';` |
| 15 | `tags` | `COMMENT ON COLUMN catalog.products.tags IS '...';` |
| 16 | `is_featured` | `COMMENT ON COLUMN catalog.products.is_featured IS '...';` |
| 17 | `search_vector` | `COMMENT ON COLUMN catalog.products.search_vector IS '...';` |

### `tags` (3 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN catalog.tags.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN catalog.tags.name IS '...';` |
| 3 | `slug` | `COMMENT ON COLUMN catalog.tags.slug IS '...';` |

## Schema: `orders`

### `cart_items` (6 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.cart_items.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.cart_items.name IS '...';` |
| 3 | `cart_id` | `COMMENT ON COLUMN orders.cart_items.cart_id IS '...';` |
| 4 | `product_id` | `COMMENT ON COLUMN orders.cart_items.product_id IS '...';` |
| 5 | `variant_id` | `COMMENT ON COLUMN orders.cart_items.variant_id IS '...';` |
| 6 | `quantity` | `COMMENT ON COLUMN orders.cart_items.quantity IS '...';` |

### `carts` (5 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.carts.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.carts.name IS '...';` |
| 3 | `customer_id` | `COMMENT ON COLUMN orders.carts.customer_id IS '...';` |
| 4 | `is_active` | `COMMENT ON COLUMN orders.carts.is_active IS '...';` |
| 5 | `expires_at` | `COMMENT ON COLUMN orders.carts.expires_at IS '...';` |

### `coupons` (11 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.coupons.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.coupons.name IS '...';` |
| 3 | `code` | `COMMENT ON COLUMN orders.coupons.code IS '...';` |
| 4 | `discount_type` | `COMMENT ON COLUMN orders.coupons.discount_type IS '...';` |
| 5 | `discount_value` | `COMMENT ON COLUMN orders.coupons.discount_value IS '...';` |
| 6 | `min_order_value` | `COMMENT ON COLUMN orders.coupons.min_order_value IS '...';` |
| 7 | `max_uses` | `COMMENT ON COLUMN orders.coupons.max_uses IS '...';` |
| 8 | `used_count` | `COMMENT ON COLUMN orders.coupons.used_count IS '...';` |
| 9 | `is_active` | `COMMENT ON COLUMN orders.coupons.is_active IS '...';` |
| 10 | `valid_from` | `COMMENT ON COLUMN orders.coupons.valid_from IS '...';` |
| 11 | `valid_until` | `COMMENT ON COLUMN orders.coupons.valid_until IS '...';` |

### `customers` (9 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.customers.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.customers.name IS '...';` |
| 3 | `user_id` | `COMMENT ON COLUMN orders.customers.user_id IS '...';` |
| 4 | `organization_id` | `COMMENT ON COLUMN orders.customers.organization_id IS '...';` |
| 5 | `billing_address` | `COMMENT ON COLUMN orders.customers.billing_address IS '...';` |
| 6 | `shipping_address` | `COMMENT ON COLUMN orders.customers.shipping_address IS '...';` |
| 7 | `tax_id` | `COMMENT ON COLUMN orders.customers.tax_id IS '...';` |
| 8 | `loyalty_points` | `COMMENT ON COLUMN orders.customers.loyalty_points IS '...';` |
| 9 | `tier` | `COMMENT ON COLUMN orders.customers.tier IS '...';` |

### `order_items` (10 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.order_items.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.order_items.name IS '...';` |
| 3 | `order_id` | `COMMENT ON COLUMN orders.order_items.order_id IS '...';` |
| 4 | `product_id` | `COMMENT ON COLUMN orders.order_items.product_id IS '...';` |
| 5 | `variant_id` | `COMMENT ON COLUMN orders.order_items.variant_id IS '...';` |
| 6 | `quantity` | `COMMENT ON COLUMN orders.order_items.quantity IS '...';` |
| 7 | `unit_price` | `COMMENT ON COLUMN orders.order_items.unit_price IS '...';` |
| 8 | `total_price` | `COMMENT ON COLUMN orders.order_items.total_price IS '...';` |
| 9 | `discount` | `COMMENT ON COLUMN orders.order_items.discount IS '...';` |
| 10 | `metadata` | `COMMENT ON COLUMN orders.order_items.metadata IS '...';` |

### `order_status_history` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.order_status_history.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.order_status_history.name IS '...';` |
| 3 | `order_id` | `COMMENT ON COLUMN orders.order_status_history.order_id IS '...';` |
| 4 | `from_status` | `COMMENT ON COLUMN orders.order_status_history.from_status IS '...';` |
| 5 | `to_status` | `COMMENT ON COLUMN orders.order_status_history.to_status IS '...';` |
| 6 | `changed_by` | `COMMENT ON COLUMN orders.order_status_history.changed_by IS '...';` |
| 7 | `note` | `COMMENT ON COLUMN orders.order_status_history.note IS '...';` |

### `orders` (20 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.orders.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.orders.name IS '...';` |
| 3 | `order_number` | `COMMENT ON COLUMN orders.orders.order_number IS '...';` |
| 4 | `customer_id` | `COMMENT ON COLUMN orders.orders.customer_id IS '...';` |
| 5 | `status` | `COMMENT ON COLUMN orders.orders.status IS '...';` |
| 6 | `subtotal` | `COMMENT ON COLUMN orders.orders.subtotal IS '...';` |
| 7 | `tax_amount` | `COMMENT ON COLUMN orders.orders.tax_amount IS '...';` |
| 8 | `shipping_amount` | `COMMENT ON COLUMN orders.orders.shipping_amount IS '...';` |
| 9 | `discount_amount` | `COMMENT ON COLUMN orders.orders.discount_amount IS '...';` |
| 10 | `total` | `COMMENT ON COLUMN orders.orders.total IS '...';` |
| 11 | `currency` | `COMMENT ON COLUMN orders.orders.currency IS '...';` |
| 12 | `coupon_id` | `COMMENT ON COLUMN orders.orders.coupon_id IS '...';` |
| 13 | `shipping_address` | `COMMENT ON COLUMN orders.orders.shipping_address IS '...';` |
| 14 | `billing_address` | `COMMENT ON COLUMN orders.orders.billing_address IS '...';` |
| 15 | `notes` | `COMMENT ON COLUMN orders.orders.notes IS '...';` |
| 16 | `placed_at` | `COMMENT ON COLUMN orders.orders.placed_at IS '...';` |
| 17 | `confirmed_at` | `COMMENT ON COLUMN orders.orders.confirmed_at IS '...';` |
| 18 | `shipped_at` | `COMMENT ON COLUMN orders.orders.shipped_at IS '...';` |
| 19 | `delivered_at` | `COMMENT ON COLUMN orders.orders.delivered_at IS '...';` |
| 20 | `cancelled_at` | `COMMENT ON COLUMN orders.orders.cancelled_at IS '...';` |

### `payments` (10 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.payments.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.payments.name IS '...';` |
| 3 | `order_id` | `COMMENT ON COLUMN orders.payments.order_id IS '...';` |
| 4 | `amount` | `COMMENT ON COLUMN orders.payments.amount IS '...';` |
| 5 | `currency` | `COMMENT ON COLUMN orders.payments.currency IS '...';` |
| 6 | `method` | `COMMENT ON COLUMN orders.payments.method IS '...';` |
| 7 | `status` | `COMMENT ON COLUMN orders.payments.status IS '...';` |
| 8 | `provider_ref` | `COMMENT ON COLUMN orders.payments.provider_ref IS '...';` |
| 9 | `provider_data` | `COMMENT ON COLUMN orders.payments.provider_data IS '...';` |
| 10 | `paid_at` | `COMMENT ON COLUMN orders.payments.paid_at IS '...';` |

### `refunds` (9 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN orders.refunds.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN orders.refunds.name IS '...';` |
| 3 | `order_id` | `COMMENT ON COLUMN orders.refunds.order_id IS '...';` |
| 4 | `payment_id` | `COMMENT ON COLUMN orders.refunds.payment_id IS '...';` |
| 5 | `amount` | `COMMENT ON COLUMN orders.refunds.amount IS '...';` |
| 6 | `reason` | `COMMENT ON COLUMN orders.refunds.reason IS '...';` |
| 7 | `status` | `COMMENT ON COLUMN orders.refunds.status IS '...';` |
| 8 | `processed_by` | `COMMENT ON COLUMN orders.refunds.processed_by IS '...';` |
| 9 | `processed_at` | `COMMENT ON COLUMN orders.refunds.processed_at IS '...';` |

## Schema: `logistics`

### `inventory` (11 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.inventory.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.inventory.name IS '...';` |
| 3 | `product_id` | `COMMENT ON COLUMN logistics.inventory.product_id IS '...';` |
| 4 | `variant_id` | `COMMENT ON COLUMN logistics.inventory.variant_id IS '...';` |
| 5 | `warehouse_id` | `COMMENT ON COLUMN logistics.inventory.warehouse_id IS '...';` |
| 6 | `bin_id` | `COMMENT ON COLUMN logistics.inventory.bin_id IS '...';` |
| 7 | `quantity` | `COMMENT ON COLUMN logistics.inventory.quantity IS '...';` |
| 8 | `reserved` | `COMMENT ON COLUMN logistics.inventory.reserved IS '...';` |
| 9 | `reorder_level` | `COMMENT ON COLUMN logistics.inventory.reorder_level IS '...';` |
| 10 | `reorder_quantity` | `COMMENT ON COLUMN logistics.inventory.reorder_quantity IS '...';` |
| 11 | `last_counted_at` | `COMMENT ON COLUMN logistics.inventory.last_counted_at IS '...';` |

### `purchase_order_items` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.purchase_order_items.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.purchase_order_items.name IS '...';` |
| 3 | `purchase_order_id` | `COMMENT ON COLUMN logistics.purchase_order_items.purchase_order_id IS '...';` |
| 4 | `product_id` | `COMMENT ON COLUMN logistics.purchase_order_items.product_id IS '...';` |
| 5 | `quantity` | `COMMENT ON COLUMN logistics.purchase_order_items.quantity IS '...';` |
| 6 | `unit_cost` | `COMMENT ON COLUMN logistics.purchase_order_items.unit_cost IS '...';` |
| 7 | `received_qty` | `COMMENT ON COLUMN logistics.purchase_order_items.received_qty IS '...';` |

### `purchase_orders` (11 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.purchase_orders.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.purchase_orders.name IS '...';` |
| 3 | `po_number` | `COMMENT ON COLUMN logistics.purchase_orders.po_number IS '...';` |
| 4 | `supplier_id` | `COMMENT ON COLUMN logistics.purchase_orders.supplier_id IS '...';` |
| 5 | `warehouse_id` | `COMMENT ON COLUMN logistics.purchase_orders.warehouse_id IS '...';` |
| 6 | `status` | `COMMENT ON COLUMN logistics.purchase_orders.status IS '...';` |
| 7 | `total_amount` | `COMMENT ON COLUMN logistics.purchase_orders.total_amount IS '...';` |
| 8 | `expected_date` | `COMMENT ON COLUMN logistics.purchase_orders.expected_date IS '...';` |
| 9 | `received_date` | `COMMENT ON COLUMN logistics.purchase_orders.received_date IS '...';` |
| 10 | `approved_by` | `COMMENT ON COLUMN logistics.purchase_orders.approved_by IS '...';` |
| 11 | `notes` | `COMMENT ON COLUMN logistics.purchase_orders.notes IS '...';` |

### `shipment_items` (5 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.shipment_items.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.shipment_items.name IS '...';` |
| 3 | `shipment_id` | `COMMENT ON COLUMN logistics.shipment_items.shipment_id IS '...';` |
| 4 | `order_item_id` | `COMMENT ON COLUMN logistics.shipment_items.order_item_id IS '...';` |
| 5 | `quantity` | `COMMENT ON COLUMN logistics.shipment_items.quantity IS '...';` |

### `shipment_tracking` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.shipment_tracking.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.shipment_tracking.name IS '...';` |
| 3 | `shipment_id` | `COMMENT ON COLUMN logistics.shipment_tracking.shipment_id IS '...';` |
| 4 | `status` | `COMMENT ON COLUMN logistics.shipment_tracking.status IS '...';` |
| 5 | `location` | `COMMENT ON COLUMN logistics.shipment_tracking.location IS '...';` |
| 6 | `description` | `COMMENT ON COLUMN logistics.shipment_tracking.description IS '...';` |
| 7 | `event_time` | `COMMENT ON COLUMN logistics.shipment_tracking.event_time IS '...';` |

### `shipments` (12 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.shipments.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.shipments.name IS '...';` |
| 3 | `tracking_number` | `COMMENT ON COLUMN logistics.shipments.tracking_number IS '...';` |
| 4 | `order_id` | `COMMENT ON COLUMN logistics.shipments.order_id IS '...';` |
| 5 | `warehouse_id` | `COMMENT ON COLUMN logistics.shipments.warehouse_id IS '...';` |
| 6 | `carrier` | `COMMENT ON COLUMN logistics.shipments.carrier IS '...';` |
| 7 | `status` | `COMMENT ON COLUMN logistics.shipments.status IS '...';` |
| 8 | `weight_kg` | `COMMENT ON COLUMN logistics.shipments.weight_kg IS '...';` |
| 9 | `shipped_at` | `COMMENT ON COLUMN logistics.shipments.shipped_at IS '...';` |
| 10 | `delivered_at` | `COMMENT ON COLUMN logistics.shipments.delivered_at IS '...';` |
| 11 | `estimated_delivery` | `COMMENT ON COLUMN logistics.shipments.estimated_delivery IS '...';` |
| 12 | `shipping_cost` | `COMMENT ON COLUMN logistics.shipments.shipping_cost IS '...';` |

### `stock_movements` (9 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.stock_movements.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.stock_movements.name IS '...';` |
| 3 | `inventory_id` | `COMMENT ON COLUMN logistics.stock_movements.inventory_id IS '...';` |
| 4 | `movement_type` | `COMMENT ON COLUMN logistics.stock_movements.movement_type IS '...';` |
| 5 | `quantity` | `COMMENT ON COLUMN logistics.stock_movements.quantity IS '...';` |
| 6 | `reference_type` | `COMMENT ON COLUMN logistics.stock_movements.reference_type IS '...';` |
| 7 | `reference_id` | `COMMENT ON COLUMN logistics.stock_movements.reference_id IS '...';` |
| 8 | `performed_by` | `COMMENT ON COLUMN logistics.stock_movements.performed_by IS '...';` |
| 9 | `notes` | `COMMENT ON COLUMN logistics.stock_movements.notes IS '...';` |

### `storage_bins` (6 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.storage_bins.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.storage_bins.name IS '...';` |
| 3 | `zone_id` | `COMMENT ON COLUMN logistics.storage_bins.zone_id IS '...';` |
| 4 | `bin_code` | `COMMENT ON COLUMN logistics.storage_bins.bin_code IS '...';` |
| 5 | `max_capacity` | `COMMENT ON COLUMN logistics.storage_bins.max_capacity IS '...';` |
| 6 | `current_count` | `COMMENT ON COLUMN logistics.storage_bins.current_count IS '...';` |

### `storage_zones` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.storage_zones.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.storage_zones.name IS '...';` |
| 3 | `warehouse_id` | `COMMENT ON COLUMN logistics.storage_zones.warehouse_id IS '...';` |
| 4 | `zone_code` | `COMMENT ON COLUMN logistics.storage_zones.zone_code IS '...';` |
| 5 | `zone_type` | `COMMENT ON COLUMN logistics.storage_zones.zone_type IS '...';` |
| 6 | `temperature_min` | `COMMENT ON COLUMN logistics.storage_zones.temperature_min IS '...';` |
| 7 | `temperature_max` | `COMMENT ON COLUMN logistics.storage_zones.temperature_max IS '...';` |

### `suppliers` (9 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.suppliers.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.suppliers.name IS '...';` |
| 3 | `code` | `COMMENT ON COLUMN logistics.suppliers.code IS '...';` |
| 4 | `contact_name` | `COMMENT ON COLUMN logistics.suppliers.contact_name IS '...';` |
| 5 | `contact_email` | `COMMENT ON COLUMN logistics.suppliers.contact_email IS '...';` |
| 6 | `contact_phone` | `COMMENT ON COLUMN logistics.suppliers.contact_phone IS '...';` |
| 7 | `address` | `COMMENT ON COLUMN logistics.suppliers.address IS '...';` |
| 8 | `is_active` | `COMMENT ON COLUMN logistics.suppliers.is_active IS '...';` |
| 9 | `rating` | `COMMENT ON COLUMN logistics.suppliers.rating IS '...';` |

### `warehouses` (8 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN logistics.warehouses.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN logistics.warehouses.name IS '...';` |
| 3 | `code` | `COMMENT ON COLUMN logistics.warehouses.code IS '...';` |
| 4 | `address` | `COMMENT ON COLUMN logistics.warehouses.address IS '...';` |
| 5 | `organization_id` | `COMMENT ON COLUMN logistics.warehouses.organization_id IS '...';` |
| 6 | `manager_id` | `COMMENT ON COLUMN logistics.warehouses.manager_id IS '...';` |
| 7 | `is_active` | `COMMENT ON COLUMN logistics.warehouses.is_active IS '...';` |
| 8 | `capacity` | `COMMENT ON COLUMN logistics.warehouses.capacity IS '...';` |

## Schema: `analytics`

### `alert_history` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.alert_history.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.alert_history.name IS '...';` |
| 3 | `alert_rule_id` | `COMMENT ON COLUMN analytics.alert_history.alert_rule_id IS '...';` |
| 4 | `triggered_value` | `COMMENT ON COLUMN analytics.alert_history.triggered_value IS '...';` |
| 5 | `resolved` | `COMMENT ON COLUMN analytics.alert_history.resolved IS '...';` |
| 6 | `resolved_at` | `COMMENT ON COLUMN analytics.alert_history.resolved_at IS '...';` |
| 7 | `resolved_by` | `COMMENT ON COLUMN analytics.alert_history.resolved_by IS '...';` |

### `alert_rules` (10 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.alert_rules.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.alert_rules.name IS '...';` |
| 3 | `description` | `COMMENT ON COLUMN analytics.alert_rules.description IS '...';` |
| 4 | `owner_id` | `COMMENT ON COLUMN analytics.alert_rules.owner_id IS '...';` |
| 5 | `condition` | `COMMENT ON COLUMN analytics.alert_rules.condition IS '...';` |
| 6 | `action` | `COMMENT ON COLUMN analytics.alert_rules.action IS '...';` |
| 7 | `severity` | `COMMENT ON COLUMN analytics.alert_rules.severity IS '...';` |
| 8 | `is_active` | `COMMENT ON COLUMN analytics.alert_rules.is_active IS '...';` |
| 9 | `last_triggered` | `COMMENT ON COLUMN analytics.alert_rules.last_triggered IS '...';` |
| 10 | `cooldown_minutes` | `COMMENT ON COLUMN analytics.alert_rules.cooldown_minutes IS '...';` |

### `audit_logs` (12 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.audit_logs.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.audit_logs.name IS '...';` |
| 3 | `user_id` | `COMMENT ON COLUMN analytics.audit_logs.user_id IS '...';` |
| 4 | `action` | `COMMENT ON COLUMN analytics.audit_logs.action IS '...';` |
| 5 | `resource_type` | `COMMENT ON COLUMN analytics.audit_logs.resource_type IS '...';` |
| 6 | `resource_id` | `COMMENT ON COLUMN analytics.audit_logs.resource_id IS '...';` |
| 7 | `severity` | `COMMENT ON COLUMN analytics.audit_logs.severity IS '...';` |
| 8 | `ip_address` | `COMMENT ON COLUMN analytics.audit_logs.ip_address IS '...';` |
| 9 | `user_agent` | `COMMENT ON COLUMN analytics.audit_logs.user_agent IS '...';` |
| 10 | `old_values` | `COMMENT ON COLUMN analytics.audit_logs.old_values IS '...';` |
| 11 | `new_values` | `COMMENT ON COLUMN analytics.audit_logs.new_values IS '...';` |
| 12 | `metadata` | `COMMENT ON COLUMN analytics.audit_logs.metadata IS '...';` |

### `dashboard_widgets` (8 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.dashboard_widgets.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.dashboard_widgets.name IS '...';` |
| 3 | `dashboard_id` | `COMMENT ON COLUMN analytics.dashboard_widgets.dashboard_id IS '...';` |
| 4 | `widget_type` | `COMMENT ON COLUMN analytics.dashboard_widgets.widget_type IS '...';` |
| 5 | `config` | `COMMENT ON COLUMN analytics.dashboard_widgets.config IS '...';` |
| 6 | `position` | `COMMENT ON COLUMN analytics.dashboard_widgets.position IS '...';` |
| 7 | `data_source` | `COMMENT ON COLUMN analytics.dashboard_widgets.data_source IS '...';` |
| 8 | `refresh_interval` | `COMMENT ON COLUMN analytics.dashboard_widgets.refresh_interval IS '...';` |

### `dashboards` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.dashboards.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.dashboards.name IS '...';` |
| 3 | `description` | `COMMENT ON COLUMN analytics.dashboards.description IS '...';` |
| 4 | `owner_id` | `COMMENT ON COLUMN analytics.dashboards.owner_id IS '...';` |
| 5 | `organization_id` | `COMMENT ON COLUMN analytics.dashboards.organization_id IS '...';` |
| 6 | `is_public` | `COMMENT ON COLUMN analytics.dashboards.is_public IS '...';` |
| 7 | `layout` | `COMMENT ON COLUMN analytics.dashboards.layout IS '...';` |

### `events` (11 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.events.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.events.name IS '...';` |
| 3 | `event_type` | `COMMENT ON COLUMN analytics.events.event_type IS '...';` |
| 4 | `source_schema` | `COMMENT ON COLUMN analytics.events.source_schema IS '...';` |
| 5 | `source_table` | `COMMENT ON COLUMN analytics.events.source_table IS '...';` |
| 6 | `source_id` | `COMMENT ON COLUMN analytics.events.source_id IS '...';` |
| 7 | `actor_id` | `COMMENT ON COLUMN analytics.events.actor_id IS '...';` |
| 8 | `payload` | `COMMENT ON COLUMN analytics.events.payload IS '...';` |
| 9 | `severity` | `COMMENT ON COLUMN analytics.events.severity IS '...';` |
| 10 | `processed` | `COMMENT ON COLUMN analytics.events.processed IS '...';` |
| 11 | `processed_at` | `COMMENT ON COLUMN analytics.events.processed_at IS '...';` |

### `metrics` (6 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.metrics.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.metrics.name IS '...';` |
| 3 | `metric_key` | `COMMENT ON COLUMN analytics.metrics.metric_key IS '...';` |
| 4 | `value` | `COMMENT ON COLUMN analytics.metrics.value IS '...';` |
| 5 | `dimensions` | `COMMENT ON COLUMN analytics.metrics.dimensions IS '...';` |
| 6 | `recorded_at` | `COMMENT ON COLUMN analytics.metrics.recorded_at IS '...';` |

### `notifications` (10 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.notifications.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.notifications.name IS '...';` |
| 3 | `user_id` | `COMMENT ON COLUMN analytics.notifications.user_id IS '...';` |
| 4 | `title` | `COMMENT ON COLUMN analytics.notifications.title IS '...';` |
| 5 | `message` | `COMMENT ON COLUMN analytics.notifications.message IS '...';` |
| 6 | `channel` | `COMMENT ON COLUMN analytics.notifications.channel IS '...';` |
| 7 | `is_read` | `COMMENT ON COLUMN analytics.notifications.is_read IS '...';` |
| 8 | `action_url` | `COMMENT ON COLUMN analytics.notifications.action_url IS '...';` |
| 9 | `metadata` | `COMMENT ON COLUMN analytics.notifications.metadata IS '...';` |
| 10 | `read_at` | `COMMENT ON COLUMN analytics.notifications.read_at IS '...';` |

### `recent_events` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.recent_events.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.recent_events.name IS '...';` |
| 3 | `event_type` | `COMMENT ON COLUMN analytics.recent_events.event_type IS '...';` |
| 4 | `source_schema` | `COMMENT ON COLUMN analytics.recent_events.source_schema IS '...';` |
| 5 | `source_table` | `COMMENT ON COLUMN analytics.recent_events.source_table IS '...';` |
| 6 | `actor_id` | `COMMENT ON COLUMN analytics.recent_events.actor_id IS '...';` |
| 7 | `severity` | `COMMENT ON COLUMN analytics.recent_events.severity IS '...';` |

### `report_executions` (11 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.report_executions.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.report_executions.name IS '...';` |
| 3 | `report_id` | `COMMENT ON COLUMN analytics.report_executions.report_id IS '...';` |
| 4 | `status` | `COMMENT ON COLUMN analytics.report_executions.status IS '...';` |
| 5 | `result_data` | `COMMENT ON COLUMN analytics.report_executions.result_data IS '...';` |
| 6 | `row_count` | `COMMENT ON COLUMN analytics.report_executions.row_count IS '...';` |
| 7 | `duration_ms` | `COMMENT ON COLUMN analytics.report_executions.duration_ms IS '...';` |
| 8 | `error_message` | `COMMENT ON COLUMN analytics.report_executions.error_message IS '...';` |
| 9 | `executed_by` | `COMMENT ON COLUMN analytics.report_executions.executed_by IS '...';` |
| 10 | `started_at` | `COMMENT ON COLUMN analytics.report_executions.started_at IS '...';` |
| 11 | `completed_at` | `COMMENT ON COLUMN analytics.report_executions.completed_at IS '...';` |

### `reports` (10 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.reports.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.reports.name IS '...';` |
| 3 | `description` | `COMMENT ON COLUMN analytics.reports.description IS '...';` |
| 4 | `report_type` | `COMMENT ON COLUMN analytics.reports.report_type IS '...';` |
| 5 | `owner_id` | `COMMENT ON COLUMN analytics.reports.owner_id IS '...';` |
| 6 | `organization_id` | `COMMENT ON COLUMN analytics.reports.organization_id IS '...';` |
| 7 | `query_config` | `COMMENT ON COLUMN analytics.reports.query_config IS '...';` |
| 8 | `schedule` | `COMMENT ON COLUMN analytics.reports.schedule IS '...';` |
| 9 | `is_active` | `COMMENT ON COLUMN analytics.reports.is_active IS '...';` |
| 10 | `last_run_at` | `COMMENT ON COLUMN analytics.reports.last_run_at IS '...';` |

### `unread_notifications` (7 missing)

| # | Column | SQL fix |
|---|--------|--------|
| 1 | `id` | `COMMENT ON COLUMN analytics.unread_notifications.id IS '...';` |
| 2 | `name` | `COMMENT ON COLUMN analytics.unread_notifications.name IS '...';` |
| 3 | `user_id` | `COMMENT ON COLUMN analytics.unread_notifications.user_id IS '...';` |
| 4 | `title` | `COMMENT ON COLUMN analytics.unread_notifications.title IS '...';` |
| 5 | `message` | `COMMENT ON COLUMN analytics.unread_notifications.message IS '...';` |
| 6 | `channel` | `COMMENT ON COLUMN analytics.unread_notifications.channel IS '...';` |
| 7 | `action_url` | `COMMENT ON COLUMN analytics.unread_notifications.action_url IS '...';` |

