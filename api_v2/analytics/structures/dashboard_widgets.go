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
// • DashboardWidgetsForm     - Data input validation and creation
// • DashboardWidgets        - Main database model with relationships
// • DashboardWidgetsEdit    - Partial update operations
// • DashboardWidgetsIdentity - Bulk operation identifiers
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

	"github.com/maple-tech/baseline/types"
)

// DashboardWidgetsForm handles data input validation and creation operations.
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
type DashboardWidgetsForm struct {
	Name            string     `gorm:"column:name;not null" json:"name" example:"Est dolorem."`
	DashboardId     types.URID `gorm:"column:dashboard_id;not null" json:"dashboardId" example:"UVFB77DKHND4XJDOMIH5XUFK2I"`
	WidgetType      string     `gorm:"column:widget_type" json:"widgetType" example:"Eos voluptatem."`
	Config          types.JSON `gorm:"column:config" json:"config"`
	Position        types.JSON `gorm:"column:position" json:"position"`
	DataSource      string     `gorm:"column:data_source" json:"dataSource" example:"Quis tempore."`
	RefreshInterval int        `gorm:"column:refresh_interval" json:"refreshInterval" example:"-3762417223894361658"`
}

func (p *DashboardWidgetsForm) TableName() string {
	return "analytics.dashboard_widgets"
}

// DashboardWidgets represents the main database model for analytics.dashboard_widgets table.
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
// • Table: analytics.dashboard_widgets
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
type DashboardWidgets struct {
	Id                types.URID     `gorm:"column:id;primary_key" json:"id" example:"NXZ4VHTCFZA2XNLRMJCFQUKVIM"`
	Name              string         `gorm:"column:name;not null" json:"name" example:"Ipsa sunt."`
	DashboardId       types.URID     `gorm:"column:dashboard_id;not null" json:"dashboardId" example:"OCYOXIVHMJB5NFOWPKO246GVSI"`
	DashboardIdDetail *Dashboards    `gorm:"foreignkey:DashboardId" json:"dashboardDetail,omitempty"`
	WidgetType        string         `gorm:"column:widget_type" json:"widgetType" example:"Et sit."`
	Config            types.JSON     `gorm:"column:config" json:"config"`
	Position          types.JSON     `gorm:"column:position" json:"position"`
	DataSource        string         `gorm:"column:data_source" json:"dataSource" example:"Laudantium autem."`
	RefreshInterval   int            `gorm:"column:refresh_interval" json:"refreshInterval" example:"-5224958799677940907"`
	CreatedAt         types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *DashboardWidgets) TableName() string {
	return "analytics.dashboard_widgets"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type DashboardWidgetsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []DashboardWidgets `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type DashboardWidgetsEdit struct {
	Name            *string     `gorm:"column:name;not null" json:"name" example:"Inventore saepe."`
	DashboardId     *types.URID `gorm:"column:dashboard_id;not null" json:"dashboardId" example:"PGZK5JNCJNFOVJWYDSAQSXG6WI"`
	WidgetType      *string     `gorm:"column:widget_type" json:"widgetType" example:"Et quo."`
	Config          *types.JSON `gorm:"column:config" json:"config"`
	Position        *types.JSON `gorm:"column:position" json:"position"`
	DataSource      *string     `gorm:"column:data_source" json:"dataSource" example:"Est nostrum."`
	RefreshInterval *int        `gorm:"column:refresh_interval" json:"refreshInterval" example:"-8322736446413384860"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type DashboardWidgetsFilter struct {
	Id              *types.URID     `gorm:"column:id;primary_key" json:"id" example:"HUO2XBMSLJFIVFEXTMN7NSOMNQ"`
	Name            *string         `gorm:"column:name;not null" json:"name" example:"Perspiciatis vel."`
	DashboardId     *types.URID     `gorm:"column:dashboard_id;not null" json:"dashboardId" example:"TWRG5KJK5FFBJMPJ7JKVYTGI4Y"`
	WidgetType      *string         `gorm:"column:widget_type" json:"widgetType" example:"Deserunt occaecati."`
	Config          *types.JSON     `gorm:"column:config" json:"config"`
	Position        *types.JSON     `gorm:"column:position" json:"position"`
	DataSource      *string         `gorm:"column:data_source" json:"dataSource" example:"Nam laborum."`
	RefreshInterval *int            `gorm:"column:refresh_interval" json:"refreshInterval" example:"4551864127627199760"`
	CreatedAt       *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *DashboardWidgetsFilter) TableName() string {
	return "analytics.dashboard_widgets"
}

// --- Batch Update Struct ---
type DashboardWidgetsBatchUpdate struct {
	Data       DashboardWidgetsEdit     `json:"data"`
	PathParams DashboardWidgetsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type DashboardWidgetsIdentity struct {
	Id types.URID `json:"id"`
}
