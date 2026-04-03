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
// • AlertHistoryForm     - Data input validation and creation
// • AlertHistory        - Main database model with relationships
// • AlertHistoryEdit    - Partial update operations
// • AlertHistoryIdentity - Bulk operation identifiers
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

// AlertHistoryForm handles data input validation and creation operations.
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
type AlertHistoryForm struct {
	Name           string         `gorm:"column:name" json:"name" example:"Ex possimus."`
	AlertRuleId    types.URID     `gorm:"column:alert_rule_id;not null" json:"alertRuleId" example:"JWITNCECYZABTE32JJ6VO2ZLGM"`
	TriggeredValue types.JSON     `gorm:"column:triggered_value" json:"triggeredValue"`
	Resolved       bool           `gorm:"column:resolved" json:"resolved" example:"true"`
	ResolvedAt     types.NullTime `gorm:"column:resolved_at" json:"resolvedAt,omitempty"`
	ResolvedBy     *types.URID    `gorm:"column:resolved_by" json:"resolvedBy,omitempty" example:"KD2RGSQYTJHXBC3BRNAFBWGSIM"`
}

func (p *AlertHistoryForm) TableName() string {
	return "analytics.alert_history"
}

// AlertHistory represents the main database model for analytics.alert_history table.
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
// • Table: analytics.alert_history
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
type AlertHistory struct {
	Id                types.URID             `gorm:"column:id;primary_key" json:"id" example:"DQ6OPT3ZW5FEXG5LHBPSVOW7CI"`
	Name              string                 `gorm:"column:name" json:"name" example:"Illum qui."`
	AlertRuleId       types.URID             `gorm:"column:alert_rule_id;not null" json:"alertRuleId" example:"UVYZJU6CINBPVOUSGVP5ZGS46I"`
	AlertRuleIdDetail *AlertRules            `gorm:"foreignkey:AlertRuleId" json:"alertRuleDetail,omitempty"`
	TriggeredValue    types.JSON             `gorm:"column:triggered_value" json:"triggeredValue"`
	Resolved          bool                   `gorm:"column:resolved" json:"resolved" example:"true"`
	ResolvedAt        types.NullTime         `gorm:"column:resolved_at" json:"resolvedAt,omitempty"`
	ResolvedBy        *types.URID            `gorm:"column:resolved_by" json:"resolvedBy,omitempty" example:"YKQ6KBH3SFBJ5MG6PBTOST5BZM"`
	ResolvedByDetail  *shared_types.IamUsers `gorm:"foreignkey:ResolvedBy" json:"resolvedByDetail,omitempty"`
	CreatedAt         types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *AlertHistory) TableName() string {
	return "analytics.alert_history"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type AlertHistoryPage struct {
	paginationRuntime.DefaultPageResponse
	Items []AlertHistory `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type AlertHistoryEdit struct {
	Name           *string         `gorm:"column:name" json:"name" example:"Beatae eum."`
	AlertRuleId    *types.URID     `gorm:"column:alert_rule_id;not null" json:"alertRuleId" example:"77DOUGGULNDLLNPM5VQAO4FGTY"`
	TriggeredValue *types.JSON     `gorm:"column:triggered_value" json:"triggeredValue"`
	Resolved       *bool           `gorm:"column:resolved" json:"resolved" example:"false"`
	ResolvedAt     *types.NullTime `gorm:"column:resolved_at" json:"resolvedAt,omitempty"`
	ResolvedBy     *types.URID     `gorm:"column:resolved_by" json:"resolvedBy,omitempty" example:"GA6WKK2F5BEMRKPOS5FZLSXQ44"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type AlertHistoryFilter struct {
	Id             *types.URID     `gorm:"column:id;primary_key" json:"id" example:"254EQQA3JRHE3MHA7DV27PPSEU"`
	Name           *string         `gorm:"column:name" json:"name" example:"Rem eligendi."`
	AlertRuleId    *types.URID     `gorm:"column:alert_rule_id;not null" json:"alertRuleId" example:"HG7NE6KSYRFP3HLFA33E2Z5ICE"`
	TriggeredValue *types.JSON     `gorm:"column:triggered_value" json:"triggeredValue"`
	Resolved       *bool           `gorm:"column:resolved" json:"resolved" example:"true"`
	ResolvedAt     *types.NullTime `gorm:"column:resolved_at" json:"resolvedAt,omitempty"`
	ResolvedBy     *types.URID     `gorm:"column:resolved_by" json:"resolvedBy,omitempty" example:"VRDHXOAAJBF2BF6CMTEK3GJYXI"`
	CreatedAt      *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *AlertHistoryFilter) TableName() string {
	return "analytics.alert_history"
}

// --- Batch Update Struct ---
type AlertHistoryBatchUpdate struct {
	Data       AlertHistoryEdit     `json:"data"`
	PathParams AlertHistoryIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type AlertHistoryIdentity struct {
	Id types.URID `json:"id"`
}
