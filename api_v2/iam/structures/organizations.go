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
// • OrganizationsForm     - Data input validation and creation
// • Organizations        - Main database model with relationships
// • OrganizationsEdit    - Partial update operations
// • OrganizationsIdentity - Bulk operation identifiers
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

// OrganizationsForm handles data input validation and creation operations.
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
type OrganizationsForm struct {
	Name        string      `gorm:"column:name;not null" json:"name" example:"Ut non."`
	Slug        string      `gorm:"column:slug;not null" json:"slug" example:"Nihil excepturi."`
	Description string      `gorm:"column:description" json:"description" example:"Repudiandae laudantium."`
	LogoUrl     string      `gorm:"column:logo_url" json:"logoUrl" example:"Aliquam expedita."`
	ParentId    *types.URID `gorm:"column:parent_id" json:"parentId,omitempty" example:"GAGPSS3A4BBUPE7KSU43YFEOUM"`
	Settings    types.JSON  `gorm:"column:settings" json:"settings"`
	IsActive    bool        `gorm:"column:is_active" json:"isActive" example:"false"`
}

func (p *OrganizationsForm) TableName() string {
	return "iam.organizations"
}

// Organizations represents the main database model for iam.organizations table.
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
// • Table: iam.organizations
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
type Organizations struct {
	Id                types.URID                          `gorm:"column:id;primary_key" json:"id" example:"DZZQX6JZT5GEPA4RKLGIAOTXXA"`
	Name              string                              `gorm:"column:name;not null" json:"name" example:"Ut debitis."`
	Slug              string                              `gorm:"column:slug;not null" json:"slug" example:"Odit est."`
	Description       string                              `gorm:"column:description" json:"description" example:"Libero iste."`
	LogoUrl           string                              `gorm:"column:logo_url" json:"logoUrl" example:"Consequatur ipsam."`
	ParentId          *types.URID                         `gorm:"column:parent_id" json:"parentId,omitempty" example:"2SWA2VWE4FBCJML3SRBHVYCTEE"`
	ParentIdDetail    *Organizations                      `gorm:"foreignkey:ParentId" json:"parentDetail,omitempty"`
	Settings          types.JSON                          `gorm:"column:settings" json:"settings"`
	IsActive          bool                                `gorm:"column:is_active" json:"isActive" example:"true"`
	CreatedAt         types.NullTime                      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         types.NullTime                      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt         types.NullTime                      `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
	OrganizationsList *[]Organizations                    `gorm:"foreignKey:ParentId;references:Id" json:"organizationsList,omitempty"`
	UsersList         *[]Users                            `gorm:"foreignKey:OrganizationId;references:Id" json:"usersList,omitempty"`
	RolesList         *[]Roles                            `gorm:"foreignKey:OrganizationId;references:Id" json:"rolesList,omitempty"`
	TeamsList         *[]Teams                            `gorm:"foreignKey:OrganizationId;references:Id" json:"teamsList,omitempty"`
	ApiKeysList       *[]ApiKeys                          `gorm:"foreignKey:OrganizationId;references:Id" json:"apiKeysList,omitempty"`
	InvitationsList   *[]Invitations                      `gorm:"foreignKey:OrganizationId;references:Id" json:"invitationsList,omitempty"`
	BrandsList        *[]shared_types.CatalogBrands       `gorm:"foreignKey:OrganizationId;references:Id" json:"brandsList,omitempty"`
	CustomersList     *[]shared_types.OrdersCustomers     `gorm:"foreignKey:OrganizationId;references:Id" json:"customersList,omitempty"`
	WarehousesList    *[]shared_types.LogisticsWarehouses `gorm:"foreignKey:OrganizationId;references:Id" json:"warehousesList,omitempty"`
	DashboardsList    *[]shared_types.AnalyticsDashboards `gorm:"foreignKey:OrganizationId;references:Id" json:"dashboardsList,omitempty"`
	ReportsList       *[]shared_types.AnalyticsReports    `gorm:"foreignKey:OrganizationId;references:Id" json:"reportsList,omitempty"`
}

func (p *Organizations) TableName() string {
	return "iam.organizations"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type OrganizationsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Organizations `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type OrganizationsEdit struct {
	Name        *string     `gorm:"column:name;not null" json:"name" example:"Inventore illo."`
	Slug        *string     `gorm:"column:slug;not null" json:"slug" example:"Illo et."`
	Description *string     `gorm:"column:description" json:"description" example:"Et aut."`
	LogoUrl     *string     `gorm:"column:logo_url" json:"logoUrl" example:"Quam reiciendis."`
	ParentId    *types.URID `gorm:"column:parent_id" json:"parentId,omitempty" example:"DAIGDR2DDZCJZOSIOCPYHK7QKM"`
	Settings    *types.JSON `gorm:"column:settings" json:"settings"`
	IsActive    *bool       `gorm:"column:is_active" json:"isActive" example:"true"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type OrganizationsFilter struct {
	Id          *types.URID     `gorm:"column:id;primary_key" json:"id" example:"PIKVZK3FURDIZOZHXWQYLGR7NI"`
	Name        *string         `gorm:"column:name;not null" json:"name" example:"Molestiae qui."`
	Slug        *string         `gorm:"column:slug;not null" json:"slug" example:"Provident sequi."`
	Description *string         `gorm:"column:description" json:"description" example:"Perspiciatis unde."`
	LogoUrl     *string         `gorm:"column:logo_url" json:"logoUrl" example:"Laudantium ipsam."`
	ParentId    *types.URID     `gorm:"column:parent_id" json:"parentId,omitempty" example:"VRTDJ6KOFJAQHL2SDJCF3476WM"`
	Settings    *types.JSON     `gorm:"column:settings" json:"settings"`
	IsActive    *bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	CreatedAt   *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt   *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt   *types.NullTime `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}

func (p *OrganizationsFilter) TableName() string {
	return "iam.organizations"
}

// --- Batch Update Struct ---
type OrganizationsBatchUpdate struct {
	Data       OrganizationsEdit     `json:"data"`
	PathParams OrganizationsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type OrganizationsIdentity struct {
	Id types.URID `json:"id"`
}
