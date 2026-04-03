package shared_types

import (
	"net"
	"time"
)

// AnalyticsAlertHistory represents the shared type for analytics.alert_history
type AnalyticsAlertHistory struct {
	Id             string     `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name"`
	AlertRuleId    string     `json:"alert_rule_id"`
	TriggeredValue string     `json:"triggered_value"`
	Resolved       bool       `json:"resolved"`
	ResolvedAt     *time.Time `json:"resolved_at"`
	ResolvedBy     string     `json:"resolved_by"`
	CreatedAt      time.Time  `json:"created_at"`
}

func (p *AnalyticsAlertHistory) TableName() string {
	return "analytics.alert_history"
}

// AnalyticsDashboardWidgets represents the shared type for analytics.dashboard_widgets
type AnalyticsDashboardWidgets struct {
	Id              string    `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	DashboardId     string    `json:"dashboard_id"`
	WidgetType      string    `json:"widget_type"`
	Config          string    `json:"config"`
	Position        string    `json:"position"`
	DataSource      string    `json:"data_source"`
	RefreshInterval int       `json:"refresh_interval"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (p *AnalyticsDashboardWidgets) TableName() string {
	return "analytics.dashboard_widgets"
}

// IamInvitations represents the shared type for iam.invitations
type IamInvitations struct {
	Id             string     `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	OrganizationId string     `json:"organization_id"`
	InvitedBy      string     `json:"invited_by"`
	RoleId         string     `json:"role_id"`
	Token          string     `json:"token"`
	AcceptedAt     *time.Time `json:"accepted_at"`
	ExpiresAt      time.Time  `json:"expires_at"`
	CreatedAt      time.Time  `json:"created_at"`
}

func (p *IamInvitations) TableName() string {
	return "iam.invitations"
}

// IamRolePermissions represents the shared type for iam.role_permissions
type IamRolePermissions struct {
	Id           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	RoleId       string    `json:"role_id"`
	PermissionId string    `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func (p *IamRolePermissions) TableName() string {
	return "iam.role_permissions"
}

// CatalogCategories represents the shared type for catalog.categories
type CatalogCategories struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	ParentId    string    `json:"parent_id"`
	Icon        string    `json:"icon"`
	SortOrder   int       `json:"sort_order"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *CatalogCategories) TableName() string {
	return "catalog.categories"
}

// CatalogCollections represents the shared type for catalog.collections
type CatalogCollections struct {
	Id            string     `json:"id" gorm:"primaryKey"`
	Name          string     `json:"name"`
	Slug          string     `json:"slug"`
	Description   string     `json:"description"`
	CoverImageUrl string     `json:"cover_image_url"`
	IsActive      bool       `json:"is_active"`
	StartDate     *time.Time `json:"start_date"`
	EndDate       *time.Time `json:"end_date"`
	CreatedBy     string     `json:"created_by"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (p *CatalogCollections) TableName() string {
	return "catalog.collections"
}

// CatalogProductReviews represents the shared type for catalog.product_reviews
type CatalogProductReviews struct {
	Id           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	ProductId    string    `json:"product_id"`
	UserId       string    `json:"user_id"`
	Rating       int16     `json:"rating"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	IsVerified   bool      `json:"is_verified"`
	HelpfulCount int       `json:"helpful_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (p *CatalogProductReviews) TableName() string {
	return "catalog.product_reviews"
}

// OrdersCoupons represents the shared type for orders.coupons
type OrdersCoupons struct {
	Id            string     `json:"id" gorm:"primaryKey"`
	Name          string     `json:"name"`
	Code          string     `json:"code"`
	DiscountType  string     `json:"discount_type"`
	DiscountValue string     `json:"discount_value"`
	MinOrderValue string     `json:"min_order_value"`
	MaxUses       int        `json:"max_uses"`
	UsedCount     int        `json:"used_count"`
	IsActive      bool       `json:"is_active"`
	ValidFrom     time.Time  `json:"valid_from"`
	ValidUntil    *time.Time `json:"valid_until"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (p *OrdersCoupons) TableName() string {
	return "orders.coupons"
}

// LogisticsInventory represents the shared type for logistics.inventory
type LogisticsInventory struct {
	Id              string     `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name"`
	ProductId       string     `json:"product_id"`
	VariantId       string     `json:"variant_id"`
	WarehouseId     string     `json:"warehouse_id"`
	BinId           string     `json:"bin_id"`
	Quantity        int        `json:"quantity"`
	Reserved        int        `json:"reserved"`
	ReorderLevel    int        `json:"reorder_level"`
	ReorderQuantity int        `json:"reorder_quantity"`
	LastCountedAt   *time.Time `json:"last_counted_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (p *LogisticsInventory) TableName() string {
	return "logistics.inventory"
}

// LogisticsPurchaseOrders represents the shared type for logistics.purchase_orders
type LogisticsPurchaseOrders struct {
	Id           string     `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	PoNumber     string     `json:"po_number"`
	SupplierId   string     `json:"supplier_id"`
	WarehouseId  string     `json:"warehouse_id"`
	Status       string     `json:"status"`
	TotalAmount  string     `json:"total_amount"`
	ExpectedDate *time.Time `json:"expected_date"`
	ReceivedDate *time.Time `json:"received_date"`
	CreatedBy    string     `json:"created_by"`
	ApprovedBy   string     `json:"approved_by"`
	Notes        string     `json:"notes"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (p *LogisticsPurchaseOrders) TableName() string {
	return "logistics.purchase_orders"
}

// IamOrganizations represents the shared type for iam.organizations
type IamOrganizations struct {
	Id          string     `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	LogoUrl     string     `json:"logo_url"`
	ParentId    string     `json:"parent_id"`
	Settings    string     `json:"settings"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (p *IamOrganizations) TableName() string {
	return "iam.organizations"
}

// OrdersCartItems represents the shared type for orders.cart_items
type OrdersCartItems struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CartId    string    `json:"cart_id"`
	ProductId string    `json:"product_id"`
	VariantId string    `json:"variant_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *OrdersCartItems) TableName() string {
	return "orders.cart_items"
}

// OrdersOrderStatusHistory represents the shared type for orders.order_status_history
type OrdersOrderStatusHistory struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	OrderId    string    `json:"order_id"`
	FromStatus string    `json:"from_status"`
	ToStatus   string    `json:"to_status"`
	ChangedBy  string    `json:"changed_by"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"created_at"`
}

func (p *OrdersOrderStatusHistory) TableName() string {
	return "orders.order_status_history"
}

// LogisticsShipmentTracking represents the shared type for logistics.shipment_tracking
type LogisticsShipmentTracking struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	ShipmentId  string    `json:"shipment_id"`
	Status      string    `json:"status"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	EventTime   time.Time `json:"event_time"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *LogisticsShipmentTracking) TableName() string {
	return "logistics.shipment_tracking"
}

// LogisticsStorageBins represents the shared type for logistics.storage_bins
type LogisticsStorageBins struct {
	Id           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	ZoneId       string    `json:"zone_id"`
	BinCode      string    `json:"bin_code"`
	MaxCapacity  int       `json:"max_capacity"`
	CurrentCount int       `json:"current_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (p *LogisticsStorageBins) TableName() string {
	return "logistics.storage_bins"
}

// AnalyticsAlertRules represents the shared type for analytics.alert_rules
type AnalyticsAlertRules struct {
	Id              string     `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	OwnerId         string     `json:"owner_id"`
	Condition       string     `json:"condition"`
	Action          string     `json:"action"`
	Severity        string     `json:"severity"`
	IsActive        bool       `json:"is_active"`
	LastTriggered   *time.Time `json:"last_triggered"`
	CooldownMinutes int        `json:"cooldown_minutes"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (p *AnalyticsAlertRules) TableName() string {
	return "analytics.alert_rules"
}

// AnalyticsAuditLogs represents the shared type for analytics.audit_logs
type AnalyticsAuditLogs struct {
	Id           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	UserId       string    `json:"user_id"`
	Action       string    `json:"action"`
	ResourceType string    `json:"resource_type"`
	ResourceId   string    `json:"resource_id"`
	Severity     string    `json:"severity"`
	IpAddress    *net.IP   `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	OldValues    string    `json:"old_values"`
	NewValues    string    `json:"new_values"`
	Metadata     string    `json:"metadata"`
	CreatedAt    time.Time `json:"created_at"`
}

func (p *AnalyticsAuditLogs) TableName() string {
	return "analytics.audit_logs"
}

// AnalyticsEvents represents the shared type for analytics.events
type AnalyticsEvents struct {
	Id           string     `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	EventType    string     `json:"event_type"`
	SourceSchema string     `json:"source_schema"`
	SourceTable  string     `json:"source_table"`
	SourceId     string     `json:"source_id"`
	ActorId      string     `json:"actor_id"`
	Payload      string     `json:"payload"`
	Severity     string     `json:"severity"`
	Processed    bool       `json:"processed"`
	ProcessedAt  *time.Time `json:"processed_at"`
	CreatedAt    time.Time  `json:"created_at"`
}

func (p *AnalyticsEvents) TableName() string {
	return "analytics.events"
}

// IamRoles represents the shared type for iam.roles
type IamRoles struct {
	Id             string    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	Description    string    `json:"description"`
	OrganizationId string    `json:"organization_id"`
	IsSystem       bool      `json:"is_system"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *IamRoles) TableName() string {
	return "iam.roles"
}

// IamTeamMembers represents the shared type for iam.team_members
type IamTeamMembers struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	TeamId    string    `json:"team_id"`
	UserId    string    `json:"user_id"`
	Role      string    `json:"role"`
	JoinedAt  time.Time `json:"joined_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *IamTeamMembers) TableName() string {
	return "iam.team_members"
}

// CatalogBrands represents the shared type for catalog.brands
type CatalogBrands struct {
	Id             string    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	LogoUrl        string    `json:"logo_url"`
	Website        string    `json:"website"`
	Description    string    `json:"description"`
	OrganizationId string    `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *CatalogBrands) TableName() string {
	return "catalog.brands"
}

// AnalyticsMetrics represents the shared type for analytics.metrics
type AnalyticsMetrics struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	MetricKey  string    `json:"metric_key"`
	Value      string    `json:"value"`
	Dimensions string    `json:"dimensions"`
	RecordedAt time.Time `json:"recorded_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func (p *AnalyticsMetrics) TableName() string {
	return "analytics.metrics"
}

// AnalyticsReportExecutions represents the shared type for analytics.report_executions
type AnalyticsReportExecutions struct {
	Id           string     `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	ReportId     string     `json:"report_id"`
	Status       string     `json:"status"`
	ResultData   string     `json:"result_data"`
	RowCount     int        `json:"row_count"`
	DurationMs   int        `json:"duration_ms"`
	ErrorMessage string     `json:"error_message"`
	ExecutedBy   string     `json:"executed_by"`
	StartedAt    *time.Time `json:"started_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	CreatedAt    time.Time  `json:"created_at"`
}

func (p *AnalyticsReportExecutions) TableName() string {
	return "analytics.report_executions"
}

// AnalyticsReports represents the shared type for analytics.reports
type AnalyticsReports struct {
	Id             string     `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	ReportType     string     `json:"report_type"`
	OwnerId        string     `json:"owner_id"`
	OrganizationId string     `json:"organization_id"`
	QueryConfig    string     `json:"query_config"`
	Schedule       string     `json:"schedule"`
	IsActive       bool       `json:"is_active"`
	LastRunAt      *time.Time `json:"last_run_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (p *AnalyticsReports) TableName() string {
	return "analytics.reports"
}

// CatalogProductVariants represents the shared type for catalog.product_variants
type CatalogProductVariants struct {
	Id            string    `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	ProductId     string    `json:"product_id"`
	Sku           string    `json:"sku"`
	PriceOverride string    `json:"price_override"`
	Attributes    string    `json:"attributes"`
	IsActive      bool      `json:"is_active"`
	StockQuantity int       `json:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (p *CatalogProductVariants) TableName() string {
	return "catalog.product_variants"
}

// IamUserRoles represents the shared type for iam.user_roles
type IamUserRoles struct {
	Id        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	UserId    string     `json:"user_id"`
	RoleId    string     `json:"role_id"`
	GrantedBy string     `json:"granted_by"`
	GrantedAt time.Time  `json:"granted_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedAt time.Time  `json:"created_at"`
}

func (p *IamUserRoles) TableName() string {
	return "iam.user_roles"
}

// CatalogProductTags represents the shared type for catalog.product_tags
type CatalogProductTags struct {
	ProductId string    `json:"product_id"`
	TagId     string    `json:"tag_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *CatalogProductTags) TableName() string {
	return "catalog.product_tags"
}

// OrdersCarts represents the shared type for orders.carts
type OrdersCarts struct {
	Id         string     `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name"`
	CustomerId string     `json:"customer_id"`
	IsActive   bool       `json:"is_active"`
	ExpiresAt  *time.Time `json:"expires_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (p *OrdersCarts) TableName() string {
	return "orders.carts"
}

// OrdersCustomers represents the shared type for orders.customers
type OrdersCustomers struct {
	Id              string    `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	UserId          string    `json:"user_id"`
	OrganizationId  string    `json:"organization_id"`
	BillingAddress  string    `json:"billing_address"`
	ShippingAddress string    `json:"shipping_address"`
	TaxId           string    `json:"tax_id"`
	LoyaltyPoints   int       `json:"loyalty_points"`
	Tier            string    `json:"tier"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (p *OrdersCustomers) TableName() string {
	return "orders.customers"
}

// LogisticsSuppliers represents the shared type for logistics.suppliers
type LogisticsSuppliers struct {
	Id           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	ContactName  string    `json:"contact_name"`
	ContactEmail string    `json:"contact_email"`
	ContactPhone string    `json:"contact_phone"`
	Address      string    `json:"address"`
	IsActive     bool      `json:"is_active"`
	Rating       string    `json:"rating"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (p *LogisticsSuppliers) TableName() string {
	return "logistics.suppliers"
}

// AnalyticsDashboards represents the shared type for analytics.dashboards
type AnalyticsDashboards struct {
	Id             string    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	OwnerId        string    `json:"owner_id"`
	OrganizationId string    `json:"organization_id"`
	IsPublic       bool      `json:"is_public"`
	Layout         string    `json:"layout"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *AnalyticsDashboards) TableName() string {
	return "analytics.dashboards"
}

// AnalyticsNotifications represents the shared type for analytics.notifications
type AnalyticsNotifications struct {
	Id        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	UserId    string     `json:"user_id"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	Channel   string     `json:"channel"`
	IsRead    bool       `json:"is_read"`
	ActionUrl string     `json:"action_url"`
	Metadata  string     `json:"metadata"`
	ReadAt    *time.Time `json:"read_at"`
	CreatedAt time.Time  `json:"created_at"`
}

func (p *AnalyticsNotifications) TableName() string {
	return "analytics.notifications"
}

// IamTeams represents the shared type for iam.teams
type IamTeams struct {
	Id             string    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	OrganizationId string    `json:"organization_id"`
	LeadId         string    `json:"lead_id"`
	Tags           string    `json:"tags"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *IamTeams) TableName() string {
	return "iam.teams"
}

// CatalogCollectionProducts represents the shared type for catalog.collection_products
type CatalogCollectionProducts struct {
	Id           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	CollectionId string    `json:"collection_id"`
	ProductId    string    `json:"product_id"`
	SortOrder    int       `json:"sort_order"`
	CreatedAt    time.Time `json:"created_at"`
}

func (p *CatalogCollectionProducts) TableName() string {
	return "catalog.collection_products"
}

// CatalogTags represents the shared type for catalog.tags
type CatalogTags struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *CatalogTags) TableName() string {
	return "catalog.tags"
}

// OrdersRefunds represents the shared type for orders.refunds
type OrdersRefunds struct {
	Id          string     `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	OrderId     string     `json:"order_id"`
	PaymentId   string     `json:"payment_id"`
	Amount      string     `json:"amount"`
	Reason      string     `json:"reason"`
	Status      string     `json:"status"`
	ProcessedBy string     `json:"processed_by"`
	ProcessedAt *time.Time `json:"processed_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (p *OrdersRefunds) TableName() string {
	return "orders.refunds"
}

// LogisticsStockMovements represents the shared type for logistics.stock_movements
type LogisticsStockMovements struct {
	Id            string    `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	InventoryId   string    `json:"inventory_id"`
	MovementType  string    `json:"movement_type"`
	Quantity      int       `json:"quantity"`
	ReferenceType string    `json:"reference_type"`
	ReferenceId   string    `json:"reference_id"`
	PerformedBy   string    `json:"performed_by"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
}

func (p *LogisticsStockMovements) TableName() string {
	return "logistics.stock_movements"
}

// OrdersOrderItems represents the shared type for orders.order_items
type OrdersOrderItems struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	OrderId    string    `json:"order_id"`
	ProductId  string    `json:"product_id"`
	VariantId  string    `json:"variant_id"`
	Quantity   int       `json:"quantity"`
	UnitPrice  string    `json:"unit_price"`
	TotalPrice string    `json:"total_price"`
	Discount   string    `json:"discount"`
	Metadata   string    `json:"metadata"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (p *OrdersOrderItems) TableName() string {
	return "orders.order_items"
}

// CatalogProductMedia represents the shared type for catalog.product_media
type CatalogProductMedia struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	ProductId string    `json:"product_id"`
	MediaType string    `json:"media_type"`
	Url       string    `json:"url"`
	AltText   string    `json:"alt_text"`
	SortOrder int       `json:"sort_order"`
	IsPrimary bool      `json:"is_primary"`
	Metadata  string    `json:"metadata"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *CatalogProductMedia) TableName() string {
	return "catalog.product_media"
}

// LogisticsShipmentItems represents the shared type for logistics.shipment_items
type LogisticsShipmentItems struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	ShipmentId  string    `json:"shipment_id"`
	OrderItemId string    `json:"order_item_id"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *LogisticsShipmentItems) TableName() string {
	return "logistics.shipment_items"
}

// LogisticsShipments represents the shared type for logistics.shipments
type LogisticsShipments struct {
	Id                string     `json:"id" gorm:"primaryKey"`
	Name              string     `json:"name"`
	TrackingNumber    string     `json:"tracking_number"`
	OrderId           string     `json:"order_id"`
	WarehouseId       string     `json:"warehouse_id"`
	Carrier           string     `json:"carrier"`
	Status            string     `json:"status"`
	WeightKg          string     `json:"weight_kg"`
	ShippedAt         *time.Time `json:"shipped_at"`
	DeliveredAt       *time.Time `json:"delivered_at"`
	EstimatedDelivery *time.Time `json:"estimated_delivery"`
	ShippingCost      string     `json:"shipping_cost"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (p *LogisticsShipments) TableName() string {
	return "logistics.shipments"
}

// OrdersOrders represents the shared type for orders.orders
type OrdersOrders struct {
	Id              string     `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name"`
	OrderNumber     string     `json:"order_number"`
	CustomerId      string     `json:"customer_id"`
	Status          string     `json:"status"`
	Subtotal        string     `json:"subtotal"`
	TaxAmount       string     `json:"tax_amount"`
	ShippingAmount  string     `json:"shipping_amount"`
	DiscountAmount  string     `json:"discount_amount"`
	Total           string     `json:"total"`
	Currency        string     `json:"currency"`
	CouponId        string     `json:"coupon_id"`
	ShippingAddress string     `json:"shipping_address"`
	BillingAddress  string     `json:"billing_address"`
	Notes           string     `json:"notes"`
	PlacedAt        *time.Time `json:"placed_at"`
	ConfirmedAt     *time.Time `json:"confirmed_at"`
	ShippedAt       *time.Time `json:"shipped_at"`
	DeliveredAt     *time.Time `json:"delivered_at"`
	CancelledAt     *time.Time `json:"cancelled_at"`
	CreatedBy       string     `json:"created_by"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func (p *OrdersOrders) TableName() string {
	return "orders.orders"
}

// IamSessions represents the shared type for iam.sessions
type IamSessions struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	UserId    string    `json:"user_id"`
	IpAddress *net.IP   `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	IsActive  bool      `json:"is_active"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *IamSessions) TableName() string {
	return "iam.sessions"
}

// OrdersPayments represents the shared type for orders.payments
type OrdersPayments struct {
	Id           string     `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	OrderId      string     `json:"order_id"`
	Amount       string     `json:"amount"`
	Currency     string     `json:"currency"`
	Method       string     `json:"method"`
	Status       string     `json:"status"`
	ProviderRef  string     `json:"provider_ref"`
	ProviderData string     `json:"provider_data"`
	PaidAt       *time.Time `json:"paid_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (p *OrdersPayments) TableName() string {
	return "orders.payments"
}

// LogisticsPurchaseOrderItems represents the shared type for logistics.purchase_order_items
type LogisticsPurchaseOrderItems struct {
	Id              string    `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	PurchaseOrderId string    `json:"purchase_order_id"`
	ProductId       string    `json:"product_id"`
	Quantity        int       `json:"quantity"`
	UnitCost        string    `json:"unit_cost"`
	ReceivedQty     int       `json:"received_qty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (p *LogisticsPurchaseOrderItems) TableName() string {
	return "logistics.purchase_order_items"
}

// LogisticsStorageZones represents the shared type for logistics.storage_zones
type LogisticsStorageZones struct {
	Id             string    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	WarehouseId    string    `json:"warehouse_id"`
	ZoneCode       string    `json:"zone_code"`
	ZoneType       string    `json:"zone_type"`
	TemperatureMin string    `json:"temperature_min"`
	TemperatureMax string    `json:"temperature_max"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *LogisticsStorageZones) TableName() string {
	return "logistics.storage_zones"
}

// IamUsers represents the shared type for iam.users
type IamUsers struct {
	Id             string     `json:"id" gorm:"primaryKey"`
	Email          string     `json:"email"`
	Name           string     `json:"name"`
	DisplayName    string     `json:"display_name"`
	AvatarUrl      string     `json:"avatar_url"`
	Phone          string     `json:"phone"`
	Status         string     `json:"status"`
	AuthProvider   string     `json:"auth_provider"`
	OrganizationId string     `json:"organization_id"`
	Metadata       string     `json:"metadata"`
	LastLoginAt    *time.Time `json:"last_login_at"`
	EmailVerified  bool       `json:"email_verified"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

func (p *IamUsers) TableName() string {
	return "iam.users"
}

// CatalogProducts represents the shared type for catalog.products
type CatalogProducts struct {
	Id               string     `json:"id" gorm:"primaryKey"`
	Name             string     `json:"name"`
	Slug             string     `json:"slug"`
	Sku              string     `json:"sku"`
	Description      string     `json:"description"`
	ShortDescription string     `json:"short_description"`
	Status           string     `json:"status"`
	BrandId          string     `json:"brand_id"`
	CategoryId       string     `json:"category_id"`
	BasePrice        string     `json:"base_price"`
	Currency         string     `json:"currency"`
	WeightKg         string     `json:"weight_kg"`
	DimensionsCm     string     `json:"dimensions_cm"`
	Attributes       string     `json:"attributes"`
	Tags             string     `json:"tags"`
	IsFeatured       bool       `json:"is_featured"`
	CreatedBy        string     `json:"created_by"`
	UpdatedBy        string     `json:"updated_by"`
	SearchVector     string     `json:"search_vector"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

func (p *CatalogProducts) TableName() string {
	return "catalog.products"
}

// IamApiKeys represents the shared type for iam.api_keys
type IamApiKeys struct {
	Id             string     `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name"`
	KeyHash        string     `json:"key_hash"`
	UserId         string     `json:"user_id"`
	OrganizationId string     `json:"organization_id"`
	Scopes         string     `json:"scopes"`
	IsActive       bool       `json:"is_active"`
	LastUsedAt     *time.Time `json:"last_used_at"`
	ExpiresAt      *time.Time `json:"expires_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (p *IamApiKeys) TableName() string {
	return "iam.api_keys"
}

// IamPermissions represents the shared type for iam.permissions
type IamPermissions struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Resource    string    `json:"resource"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *IamPermissions) TableName() string {
	return "iam.permissions"
}

// CatalogPriceHistory represents the shared type for catalog.price_history
type CatalogPriceHistory struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	ProductId string    `json:"product_id"`
	OldPrice  string    `json:"old_price"`
	NewPrice  string    `json:"new_price"`
	ChangedBy string    `json:"changed_by"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *CatalogPriceHistory) TableName() string {
	return "catalog.price_history"
}

// LogisticsWarehouses represents the shared type for logistics.warehouses
type LogisticsWarehouses struct {
	Id             string    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Code           string    `json:"code"`
	Address        string    `json:"address"`
	OrganizationId string    `json:"organization_id"`
	ManagerId      string    `json:"manager_id"`
	IsActive       bool      `json:"is_active"`
	Capacity       int       `json:"capacity"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *LogisticsWarehouses) TableName() string {
	return "logistics.warehouses"
}
