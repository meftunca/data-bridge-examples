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
// • ReportsForm     - Data input validation and creation
// • Reports        - Main database model with relationships
// • ReportsEdit    - Partial update operations
// • ReportsIdentity - Bulk operation identifiers
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

// ReportsForm handles data input validation and creation operations.
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
type ReportsForm struct {
	Name           string              `gorm:"column:name;not null" json:"name" example:"Dolorum dolores."`
	Description    string              `gorm:"column:description" json:"description" example:"Alias sint."`
	ReportType     AnalyticsReportType `gorm:"column:report_type" json:"reportType"`
	OwnerId        types.URID          `gorm:"column:owner_id;not null" json:"ownerId" example:"RVWTWOLNKFA4VJWIGRGBRDGLJM"`
	OrganizationId *types.URID         `gorm:"column:organization_id" json:"organizationId,omitempty" example:"UOVT7OPHOJHBZIZ55GQSRHQKE4"`
	QueryConfig    types.JSON          `gorm:"column:query_config" json:"queryConfig"`
	Schedule       types.JSON          `gorm:"column:schedule" json:"schedule"`
	IsActive       bool                `gorm:"column:is_active" json:"isActive" example:"true"`
	LastRunAt      types.NullTime      `gorm:"column:last_run_at" json:"lastRunAt,omitempty"`
}

func (p *ReportsForm) TableName() string {
	return "analytics.reports"
}

// Reports represents the main database model for analytics.reports table.
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
// • Table: analytics.reports
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
type Reports struct {
	Id                   types.URID                     `gorm:"column:id;primary_key" json:"id" example:"6LQA36VGRNF4RJ27FFA2EAPNCY"`
	Name                 string                         `gorm:"column:name;not null" json:"name" example:"Ab voluptatum."`
	Description          string                         `gorm:"column:description" json:"description" example:"Quibusdam occaecati."`
	ReportType           AnalyticsReportType            `gorm:"column:report_type" json:"reportType"`
	OwnerId              types.URID                     `gorm:"column:owner_id;not null" json:"ownerId" example:"3VZGBEMCB5A5TIDUYEWGFUOJNI"`
	OwnerIdDetail        *shared_types.IamUsers         `gorm:"foreignkey:OwnerId" json:"ownerDetail,omitempty"`
	OrganizationId       *types.URID                    `gorm:"column:organization_id" json:"organizationId,omitempty" example:"GGH7LUP46BFJXCDEUUAMI5QKQY"`
	OrganizationIdDetail *shared_types.IamOrganizations `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	QueryConfig          types.JSON                     `gorm:"column:query_config" json:"queryConfig"`
	Schedule             types.JSON                     `gorm:"column:schedule" json:"schedule"`
	IsActive             bool                           `gorm:"column:is_active" json:"isActive" example:"true"`
	LastRunAt            types.NullTime                 `gorm:"column:last_run_at" json:"lastRunAt,omitempty"`
	CreatedAt            types.NullTime                 `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime                 `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	ReportExecutionsList *[]ReportExecutions            `gorm:"foreignKey:ReportId;references:Id" json:"reportExecutionsList,omitempty"`
}

func (p *Reports) TableName() string {
	return "analytics.reports"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ReportsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Reports `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ReportsEdit struct {
	Name           *string              `gorm:"column:name;not null" json:"name" example:"Inventore porro."`
	Description    *string              `gorm:"column:description" json:"description" example:"Eius voluptas."`
	ReportType     *AnalyticsReportType `gorm:"column:report_type" json:"reportType"`
	OwnerId        *types.URID          `gorm:"column:owner_id;not null" json:"ownerId" example:"AHKJY7SNDVAAPGY6NOS2GG6ANY"`
	OrganizationId *types.URID          `gorm:"column:organization_id" json:"organizationId,omitempty" example:"NC2KXXJ4HFAJPNQAI3W7QTUPU4"`
	QueryConfig    *types.JSON          `gorm:"column:query_config" json:"queryConfig"`
	Schedule       *types.JSON          `gorm:"column:schedule" json:"schedule"`
	IsActive       *bool                `gorm:"column:is_active" json:"isActive" example:"true"`
	LastRunAt      *types.NullTime      `gorm:"column:last_run_at" json:"lastRunAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ReportsFilter struct {
	Id             *types.URID          `gorm:"column:id;primary_key" json:"id" example:"U5T7OFUKJ5DT3ASKOYNYDQ7JC4"`
	Name           *string              `gorm:"column:name;not null" json:"name" example:"Sunt nulla."`
	Description    *string              `gorm:"column:description" json:"description" example:"Ipsa quasi."`
	ReportType     *AnalyticsReportType `gorm:"column:report_type" json:"reportType"`
	OwnerId        *types.URID          `gorm:"column:owner_id;not null" json:"ownerId" example:"NCKIWK6N6JFGRHEP5VHVWDGDFA"`
	OrganizationId *types.URID          `gorm:"column:organization_id" json:"organizationId,omitempty" example:"TYQLYLNX6NB55NU4UHDSBVFHBQ"`
	QueryConfig    *types.JSON          `gorm:"column:query_config" json:"queryConfig"`
	Schedule       *types.JSON          `gorm:"column:schedule" json:"schedule"`
	IsActive       *bool                `gorm:"column:is_active" json:"isActive" example:"true"`
	LastRunAt      *types.NullTime      `gorm:"column:last_run_at" json:"lastRunAt,omitempty"`
	CreatedAt      *types.NullTime      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ReportsFilter) TableName() string {
	return "analytics.reports"
}

// --- Batch Update Struct ---
type ReportsBatchUpdate struct {
	Data       ReportsEdit     `json:"data"`
	PathParams ReportsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ReportsIdentity struct {
	Id types.URID `json:"id"`
}
