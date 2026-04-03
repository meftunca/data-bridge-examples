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
// • AlertRulesForm     - Data input validation and creation
// • AlertRules        - Main database model with relationships
// • AlertRulesEdit    - Partial update operations
// • AlertRulesIdentity - Bulk operation identifiers
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

// AlertRulesForm handles data input validation and creation operations.
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
type AlertRulesForm struct {
	Name            string                 `gorm:"column:name;not null" json:"name" example:"Recusandae blanditiis."`
	Description     string                 `gorm:"column:description" json:"description" example:"Laborum cum."`
	OwnerId         types.URID             `gorm:"column:owner_id;not null" json:"ownerId" example:"CZDJWXVKQRFRJLMEGCWZHEE6VE"`
	Condition       types.JSON             `gorm:"column:condition" json:"condition"`
	Action          types.JSON             `gorm:"column:action" json:"action"`
	Severity        AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	IsActive        bool                   `gorm:"column:is_active" json:"isActive" example:"true"`
	LastTriggered   types.NullTime         `gorm:"column:last_triggered" json:"lastTriggered,omitempty"`
	CooldownMinutes int                    `gorm:"column:cooldown_minutes" json:"cooldownMinutes" example:"-1582357493614568705"`
}

func (p *AlertRulesForm) TableName() string {
	return "analytics.alert_rules"
}

// AlertRules represents the main database model for analytics.alert_rules table.
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
// • Table: analytics.alert_rules
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
type AlertRules struct {
	Id               types.URID             `gorm:"column:id;primary_key" json:"id" example:"R2NJTQZ7TBAEHNQLGWDDZJCT3Y"`
	Name             string                 `gorm:"column:name;not null" json:"name" example:"Tenetur vero."`
	Description      string                 `gorm:"column:description" json:"description" example:"Facere enim."`
	OwnerId          types.URID             `gorm:"column:owner_id;not null" json:"ownerId" example:"EJ5YORMGW5DEXCWR6WFERP5X3I"`
	OwnerIdDetail    *shared_types.IamUsers `gorm:"foreignkey:OwnerId" json:"ownerDetail,omitempty"`
	Condition        types.JSON             `gorm:"column:condition" json:"condition"`
	Action           types.JSON             `gorm:"column:action" json:"action"`
	Severity         AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	IsActive         bool                   `gorm:"column:is_active" json:"isActive" example:"false"`
	LastTriggered    types.NullTime         `gorm:"column:last_triggered" json:"lastTriggered,omitempty"`
	CooldownMinutes  int                    `gorm:"column:cooldown_minutes" json:"cooldownMinutes" example:"7782872486427176825"`
	CreatedAt        types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        types.NullTime         `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	AlertHistoryList *[]AlertHistory        `gorm:"foreignKey:AlertRuleId;references:Id" json:"alertHistoryList,omitempty"`
}

func (p *AlertRules) TableName() string {
	return "analytics.alert_rules"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type AlertRulesPage struct {
	paginationRuntime.DefaultPageResponse
	Items []AlertRules `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type AlertRulesEdit struct {
	Name            *string                 `gorm:"column:name;not null" json:"name" example:"Ullam aut."`
	Description     *string                 `gorm:"column:description" json:"description" example:"Qui dolor."`
	OwnerId         *types.URID             `gorm:"column:owner_id;not null" json:"ownerId" example:"RBYRLW2V5RG7VLFT4KX6BKY76Y"`
	Condition       *types.JSON             `gorm:"column:condition" json:"condition"`
	Action          *types.JSON             `gorm:"column:action" json:"action"`
	Severity        *AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	IsActive        *bool                   `gorm:"column:is_active" json:"isActive" example:"false"`
	LastTriggered   *types.NullTime         `gorm:"column:last_triggered" json:"lastTriggered,omitempty"`
	CooldownMinutes *int                    `gorm:"column:cooldown_minutes" json:"cooldownMinutes" example:"7855657844030809506"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type AlertRulesFilter struct {
	Id              *types.URID             `gorm:"column:id;primary_key" json:"id" example:"V5TZB5MLGNDDZKWIEYZWYCP7NA"`
	Name            *string                 `gorm:"column:name;not null" json:"name" example:"Consequatur illo."`
	Description     *string                 `gorm:"column:description" json:"description" example:"Officia aut."`
	OwnerId         *types.URID             `gorm:"column:owner_id;not null" json:"ownerId" example:"C5VQ27ILGBDJBNTCNJKGMYFRNY"`
	Condition       *types.JSON             `gorm:"column:condition" json:"condition"`
	Action          *types.JSON             `gorm:"column:action" json:"action"`
	Severity        *AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	IsActive        *bool                   `gorm:"column:is_active" json:"isActive" example:"true"`
	LastTriggered   *types.NullTime         `gorm:"column:last_triggered" json:"lastTriggered,omitempty"`
	CooldownMinutes *int                    `gorm:"column:cooldown_minutes" json:"cooldownMinutes" example:"-6477421651643882415"`
	CreatedAt       *types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       *types.NullTime         `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *AlertRulesFilter) TableName() string {
	return "analytics.alert_rules"
}

// --- Batch Update Struct ---
type AlertRulesBatchUpdate struct {
	Data       AlertRulesEdit     `json:"data"`
	PathParams AlertRulesIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type AlertRulesIdentity struct {
	Id types.URID `json:"id"`
}
