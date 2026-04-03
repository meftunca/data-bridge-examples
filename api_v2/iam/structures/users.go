// Database Structure Templates - Enterprise-Grade Data Models
//
// This template generates comprehensive data structures for database entities with:
// • Complete GORM model definitions with proper tags and relationships
// • Type-safe form structures for data input validation
// • Edit structures for partial update operations
// • Identity structures for efficient bulk operations
// • Cross-schema relationship support with proper imports
// • Enterprise-grade validation and serialization tags
//
// Generated Structure Types:
// • UsersForm     - Data input validation and creation
// • Users        - Main database model with relationships
// • UsersEdit    - Partial update operations
// • UsersIdentity - Bulk operation identifiers
//
// Features:
// • Automatic field validation through struct tags
// • JSON serialization with proper naming conventions
// • Database relationship mapping with foreign keys
// • Timestamp management (created_at, updated_at, deleted_at)
// • Cross-schema reference support for complex database designs
// • GORM compatibility with optimized query generation
//
// Architecture Benefits:
// • Type safety across all data operations
// • Consistent validation rules throughout the application
// • Clean separation between input, storage, and output concerns
// • Efficient bulk operations with identity-based processing
// • Maintainable code with generated documentation
package iam_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"
	shared_types "data-bridge-examples/api_v2/shared/types"

	"github.com/maple-tech/baseline/types"
)

// UsersForm handles data input validation and creation operations.
//
// This structure is specifically designed for:
// • HTTP request body parsing and validation
// • Data sanitization before database operations
// • Input validation with comprehensive error messages
// • Clean separation from database concerns (excludes auto-generated fields)
//
// Excluded Fields:
// • id: Auto-generated primary key
// • created_at, updated_at: Automatic timestamp management
// • deleted_at: Soft delete timestamp (managed by GORM)
//
// Validation Features:
// • Required field validation through struct tags
// • Data type validation and conversion
// • Custom validation rules via form tags
// • JSON unmarshaling with proper error handling
type UsersForm struct {
	Email          string          `gorm:"column:email;not null" json:"email" example:"Eos optio."`
	Name           string          `gorm:"column:name;not null" json:"name" example:"Assumenda et."`
	DisplayName    string          `gorm:"column:display_name" json:"displayName" example:"Officiis modi."`
	AvatarUrl      string          `gorm:"column:avatar_url" json:"avatarUrl" example:"Dignissimos hic."`
	Phone          string          `gorm:"column:phone" json:"phone" example:"Minima omnis."`
	Status         IamUserStatus   `gorm:"column:status" json:"status"`
	AuthProvider   IamAuthProvider `gorm:"column:auth_provider" json:"authProvider"`
	OrganizationId *types.URID     `gorm:"column:organization_id" json:"organizationId,omitempty" example:"MYAABQ6VAVAPJJID2EOUEOJULM"`
	Metadata       types.JSON      `gorm:"column:metadata" json:"metadata"`
	LastLoginAt    types.NullTime  `gorm:"column:last_login_at" json:"lastLoginAt,omitempty"`
	EmailVerified  bool            `gorm:"column:email_verified" json:"emailVerified" example:"true"`
}

func (p *UsersForm) TableName() string {
	return "iam.users"
}

// Users represents the main database model for iam.users table.
//
// This structure provides:
// • Complete GORM model definition with proper field mapping
// • Automatic relationship resolution through struct tags
// • Cross-schema reference support for complex database designs
// • Timestamp management with created_at, updated_at, deleted_at
// • Type-safe field definitions matching database schema
// • Optimized query generation through GORM integration
//
// Database Mapping:
// • Table: iam.users
// • Type: BASE TABLE
// • Schema: iam
//
// Relationship Features:
// • Automatic foreign key resolution
// • Cross-schema relationship support
// • Lazy loading for performance optimization
// • Proper join field generation for complex queries
//
// GORM Integration:
// • Automatic primary key detection
// • Soft delete support (deleted_at field)
// • Timestamp management (created_at, updated_at)
// • Index optimization for query performance
type Users struct {
	Id                     types.URID                                `gorm:"column:id;primary_key" json:"id" example:"BLFW5KF5SNCRVGFS66SPKY7X6U"`
	Email                  string                                    `gorm:"column:email;not null" json:"email" example:"Ipsam est."`
	Name                   string                                    `gorm:"column:name;not null" json:"name" example:"Ullam autem."`
	DisplayName            string                                    `gorm:"column:display_name" json:"displayName" example:"Ea eos."`
	AvatarUrl              string                                    `gorm:"column:avatar_url" json:"avatarUrl" example:"Illo odio."`
	Phone                  string                                    `gorm:"column:phone" json:"phone" example:"Fuga enim."`
	Status                 IamUserStatus                             `gorm:"column:status" json:"status"`
	AuthProvider           IamAuthProvider                           `gorm:"column:auth_provider" json:"authProvider"`
	OrganizationId         *types.URID                               `gorm:"column:organization_id" json:"organizationId,omitempty" example:"OUVMIXY2F5HYPEXQTH3SMITQTM"`
	OrganizationIdDetail   *Organizations                            `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	Metadata               types.JSON                                `gorm:"column:metadata" json:"metadata"`
	LastLoginAt            types.NullTime                            `gorm:"column:last_login_at" json:"lastLoginAt,omitempty"`
	EmailVerified          bool                                      `gorm:"column:email_verified" json:"emailVerified" example:"true"`
	CreatedAt              types.NullTime                            `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt              types.NullTime                            `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt              types.NullTime                            `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
	UserRolesList          *[]UserRoles                              `gorm:"foreignKey:UserId;references:Id" json:"userRolesList,omitempty"`
	TeamsList              *[]Teams                                  `gorm:"foreignKey:LeadId;references:Id" json:"teamsList,omitempty"`
	TeamMembersList        *[]TeamMembers                            `gorm:"foreignKey:UserId;references:Id" json:"teamMembersList,omitempty"`
	ApiKeysList            *[]ApiKeys                                `gorm:"foreignKey:UserId;references:Id" json:"apiKeysList,omitempty"`
	SessionsList           *[]Sessions                               `gorm:"foreignKey:UserId;references:Id" json:"sessionsList,omitempty"`
	InvitationsList        *[]Invitations                            `gorm:"foreignKey:InvitedBy;references:Id" json:"invitationsList,omitempty"`
	ProductsList           *[]shared_types.CatalogProducts           `gorm:"foreignKey:CreatedBy;references:Id" json:"productsList,omitempty"`
	ProductReviewsList     *[]shared_types.CatalogProductReviews     `gorm:"foreignKey:UserId;references:Id" json:"productReviewsList,omitempty"`
	CollectionsList        *[]shared_types.CatalogCollections        `gorm:"foreignKey:CreatedBy;references:Id" json:"collectionsList,omitempty"`
	PriceHistoryList       *[]shared_types.CatalogPriceHistory       `gorm:"foreignKey:ChangedBy;references:Id" json:"priceHistoryList,omitempty"`
	CustomersList          *[]shared_types.OrdersCustomers           `gorm:"foreignKey:UserId;references:Id" json:"customersList,omitempty"`
	OrdersList             *[]shared_types.OrdersOrders              `gorm:"foreignKey:CreatedBy;references:Id" json:"ordersList,omitempty"`
	RefundsList            *[]shared_types.OrdersRefunds             `gorm:"foreignKey:ProcessedBy;references:Id" json:"refundsList,omitempty"`
	OrderStatusHistoryList *[]shared_types.OrdersOrderStatusHistory  `gorm:"foreignKey:ChangedBy;references:Id" json:"orderStatusHistoryList,omitempty"`
	WarehousesList         *[]shared_types.LogisticsWarehouses       `gorm:"foreignKey:ManagerId;references:Id" json:"warehousesList,omitempty"`
	StockMovementsList     *[]shared_types.LogisticsStockMovements   `gorm:"foreignKey:PerformedBy;references:Id" json:"stockMovementsList,omitempty"`
	PurchaseOrdersList     *[]shared_types.LogisticsPurchaseOrders   `gorm:"foreignKey:CreatedBy;references:Id" json:"purchaseOrdersList,omitempty"`
	AuditLogsList          *[]shared_types.AnalyticsAuditLogs        `gorm:"foreignKey:UserId;references:Id" json:"auditLogsList,omitempty"`
	EventsList             *[]shared_types.AnalyticsEvents           `gorm:"foreignKey:ActorId;references:Id" json:"eventsList,omitempty"`
	DashboardsList         *[]shared_types.AnalyticsDashboards       `gorm:"foreignKey:OwnerId;references:Id" json:"dashboardsList,omitempty"`
	ReportsList            *[]shared_types.AnalyticsReports          `gorm:"foreignKey:OwnerId;references:Id" json:"reportsList,omitempty"`
	ReportExecutionsList   *[]shared_types.AnalyticsReportExecutions `gorm:"foreignKey:ExecutedBy;references:Id" json:"reportExecutionsList,omitempty"`
	NotificationsList      *[]shared_types.AnalyticsNotifications    `gorm:"foreignKey:UserId;references:Id" json:"notificationsList,omitempty"`
	AlertRulesList         *[]shared_types.AnalyticsAlertRules       `gorm:"foreignKey:OwnerId;references:Id" json:"alertRulesList,omitempty"`
	AlertHistoryList       *[]shared_types.AnalyticsAlertHistory     `gorm:"foreignKey:ResolvedBy;references:Id" json:"alertHistoryList,omitempty"`
}

func (p *Users) TableName() string {
	return "iam.users"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type UsersPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Users `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type UsersEdit struct {
	Email          *string          `gorm:"column:email;not null" json:"email" example:"Eum ea."`
	Name           *string          `gorm:"column:name;not null" json:"name" example:"Sed unde."`
	DisplayName    *string          `gorm:"column:display_name" json:"displayName" example:"Necessitatibus voluptas."`
	AvatarUrl      *string          `gorm:"column:avatar_url" json:"avatarUrl" example:"Laborum qui."`
	Phone          *string          `gorm:"column:phone" json:"phone" example:"Voluptatem quo."`
	Status         *IamUserStatus   `gorm:"column:status" json:"status"`
	AuthProvider   *IamAuthProvider `gorm:"column:auth_provider" json:"authProvider"`
	OrganizationId *types.URID      `gorm:"column:organization_id" json:"organizationId,omitempty" example:"O5YNZ2BGEFA4XHYJFTP4R2ALPE"`
	Metadata       *types.JSON      `gorm:"column:metadata" json:"metadata"`
	LastLoginAt    *types.NullTime  `gorm:"column:last_login_at" json:"lastLoginAt,omitempty"`
	EmailVerified  *bool            `gorm:"column:email_verified" json:"emailVerified" example:"false"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type UsersFilter struct {
	Id             *types.URID      `gorm:"column:id;primary_key" json:"id" example:"4NCXGU65MVE5LGFEYJDUSRYFME"`
	Email          *string          `gorm:"column:email;not null" json:"email" example:"Ex voluptatem."`
	Name           *string          `gorm:"column:name;not null" json:"name" example:"Omnis doloribus."`
	DisplayName    *string          `gorm:"column:display_name" json:"displayName" example:"Vel incidunt."`
	AvatarUrl      *string          `gorm:"column:avatar_url" json:"avatarUrl" example:"Laudantium dolores."`
	Phone          *string          `gorm:"column:phone" json:"phone" example:"Odit expedita."`
	Status         *IamUserStatus   `gorm:"column:status" json:"status"`
	AuthProvider   *IamAuthProvider `gorm:"column:auth_provider" json:"authProvider"`
	OrganizationId *types.URID      `gorm:"column:organization_id" json:"organizationId,omitempty" example:"Y3YQAHCEBJB2LDRRZNNI3YXAYA"`
	Metadata       *types.JSON      `gorm:"column:metadata" json:"metadata"`
	LastLoginAt    *types.NullTime  `gorm:"column:last_login_at" json:"lastLoginAt,omitempty"`
	EmailVerified  *bool            `gorm:"column:email_verified" json:"emailVerified" example:"true"`
	CreatedAt      *types.NullTime  `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime  `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt      *types.NullTime  `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}

func (p *UsersFilter) TableName() string {
	return "iam.users"
}

// --- Batch Update Struct ---
type UsersBatchUpdate struct {
	Data       UsersEdit     `json:"data"`
	PathParams UsersIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type UsersIdentity struct {
	Id types.URID `json:"id"`
}
