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
// • DashboardsForm     - Data input validation and creation
// • Dashboards        - Main database model with relationships
// • DashboardsEdit    - Partial update operations
// • DashboardsIdentity - Bulk operation identifiers
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
package analytics_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"
	shared_types "data-bridge-examples/api_v2/shared/types"

	"github.com/maple-tech/baseline/types"
)

// DashboardsForm handles data input validation and creation operations.
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
type DashboardsForm struct {
	Name           string      `gorm:"column:name;not null" json:"name" example:"Harum quo."`
	Description    string      `gorm:"column:description" json:"description" example:"Aut sunt."`
	OwnerId        types.URID  `gorm:"column:owner_id;not null" json:"ownerId" example:"E5JEHTI4YFCR3P7EARFZNURH4A"`
	OrganizationId *types.URID `gorm:"column:organization_id" json:"organizationId,omitempty" example:"W6VYW5RCJNBMBOQDCHCEBE24U4"`
	IsPublic       bool        `gorm:"column:is_public" json:"isPublic" example:"true"`
	Layout         types.JSON  `gorm:"column:layout" json:"layout"`
}

func (p *DashboardsForm) TableName() string {
	return "analytics.dashboards"
}

// Dashboards represents the main database model for analytics.dashboards table.
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
// • Table: analytics.dashboards
// • Type: BASE TABLE
// • Schema: analytics
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
type Dashboards struct {
	Id                   types.URID                     `gorm:"column:id;primary_key" json:"id" example:"NQW6EBDC6VFRBO7KQ4O44PJICE"`
	Name                 string                         `gorm:"column:name;not null" json:"name" example:"Cupiditate dolores."`
	Description          string                         `gorm:"column:description" json:"description" example:"Qui porro."`
	OwnerId              types.URID                     `gorm:"column:owner_id;not null" json:"ownerId" example:"JYV7B6HKBZBWLMIQIFZC5CIR6Y"`
	OwnerIdDetail        *shared_types.IamUsers         `gorm:"foreignkey:OwnerId" json:"ownerDetail,omitempty"`
	OrganizationId       *types.URID                    `gorm:"column:organization_id" json:"organizationId,omitempty" example:"BGG6YOJS3NH2VK7OHGPEV73AJY"`
	OrganizationIdDetail *shared_types.IamOrganizations `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	IsPublic             bool                           `gorm:"column:is_public" json:"isPublic" example:"true"`
	Layout               types.JSON                     `gorm:"column:layout" json:"layout"`
	CreatedAt            types.NullTime                 `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime                 `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DashboardWidgetsList *[]DashboardWidgets            `gorm:"foreignKey:DashboardId;references:Id" json:"dashboardWidgetsList,omitempty"`
}

func (p *Dashboards) TableName() string {
	return "analytics.dashboards"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type DashboardsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Dashboards `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type DashboardsEdit struct {
	Name           *string     `gorm:"column:name;not null" json:"name" example:"Sequi numquam."`
	Description    *string     `gorm:"column:description" json:"description" example:"Dignissimos qui."`
	OwnerId        *types.URID `gorm:"column:owner_id;not null" json:"ownerId" example:"MQWMAF5FMNFY3LSA2JPAFMLPNQ"`
	OrganizationId *types.URID `gorm:"column:organization_id" json:"organizationId,omitempty" example:"DI7XGABH7NCXVOLZKV3S6VKXQA"`
	IsPublic       *bool       `gorm:"column:is_public" json:"isPublic" example:"true"`
	Layout         *types.JSON `gorm:"column:layout" json:"layout"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type DashboardsFilter struct {
	Id             *types.URID     `gorm:"column:id;primary_key" json:"id" example:"A4VKQGUHLJGQZDGHCJA6IZBJN4"`
	Name           *string         `gorm:"column:name;not null" json:"name" example:"Officia voluptas."`
	Description    *string         `gorm:"column:description" json:"description" example:"Illum officia."`
	OwnerId        *types.URID     `gorm:"column:owner_id;not null" json:"ownerId" example:"A6XUQ7HGDZG5THIRPLGXX2UT2Y"`
	OrganizationId *types.URID     `gorm:"column:organization_id" json:"organizationId,omitempty" example:"DSZPBDOIFZGYXI7ABM6SNSIXFI"`
	IsPublic       *bool           `gorm:"column:is_public" json:"isPublic" example:"false"`
	Layout         *types.JSON     `gorm:"column:layout" json:"layout"`
	CreatedAt      *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *DashboardsFilter) TableName() string {
	return "analytics.dashboards"
}

// --- Batch Update Struct ---
type DashboardsBatchUpdate struct {
	Data       DashboardsEdit     `json:"data"`
	PathParams DashboardsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type DashboardsIdentity struct {
	Id types.URID `json:"id"`
}
