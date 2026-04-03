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
// • ReportExecutionsForm     - Data input validation and creation
// • ReportExecutions        - Main database model with relationships
// • ReportExecutionsEdit    - Partial update operations
// • ReportExecutionsIdentity - Bulk operation identifiers
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

// ReportExecutionsForm handles data input validation and creation operations.
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
type ReportExecutionsForm struct {
	Name         string         `gorm:"column:name" json:"name" example:"Minima reprehenderit."`
	ReportId     types.URID     `gorm:"column:report_id;not null" json:"reportId" example:"CRZ4KDNLWJFLJCIPCIYBI2XJWE"`
	Status       string         `gorm:"column:status" json:"status" example:"Illo laborum."`
	ResultData   types.JSON     `gorm:"column:result_data" json:"resultData,omitempty"`
	RowCount     int            `gorm:"column:row_count" json:"rowCount" example:"-5909382339582231835"`
	DurationMs   int            `gorm:"column:duration_ms" json:"durationMs" example:"6640058050001509240"`
	ErrorMessage string         `gorm:"column:error_message" json:"errorMessage" example:"Vel id."`
	ExecutedBy   *types.URID    `gorm:"column:executed_by" json:"executedBy,omitempty" example:"POUITDCIBFCWBM7GV4WS4PAQCE"`
	StartedAt    types.NullTime `gorm:"column:started_at" json:"startedAt,omitempty"`
	CompletedAt  types.NullTime `gorm:"column:completed_at" json:"completedAt,omitempty"`
}

func (p *ReportExecutionsForm) TableName() string {
	return "analytics.report_executions"
}

// ReportExecutions represents the main database model for analytics.report_executions table.
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
// • Table: analytics.report_executions
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
type ReportExecutions struct {
	Id               types.URID             `gorm:"column:id;primary_key" json:"id" example:"BW4PCH2GM5FGNFZQBWPQDAXE6Y"`
	Name             string                 `gorm:"column:name" json:"name" example:"Omnis omnis."`
	ReportId         types.URID             `gorm:"column:report_id;not null" json:"reportId" example:"ES2E5NDROJARHGZBIXOINT6QVI"`
	ReportIdDetail   *Reports               `gorm:"foreignkey:ReportId" json:"reportDetail,omitempty"`
	Status           string                 `gorm:"column:status" json:"status" example:"Explicabo ea."`
	ResultData       types.JSON             `gorm:"column:result_data" json:"resultData,omitempty"`
	RowCount         int                    `gorm:"column:row_count" json:"rowCount" example:"-2912827573989132739"`
	DurationMs       int                    `gorm:"column:duration_ms" json:"durationMs" example:"-3890900890048804286"`
	ErrorMessage     string                 `gorm:"column:error_message" json:"errorMessage" example:"Id repellat."`
	ExecutedBy       *types.URID            `gorm:"column:executed_by" json:"executedBy,omitempty" example:"T7YGBHVTTFBUHOP3YSVGJRWWL4"`
	ExecutedByDetail *shared_types.IamUsers `gorm:"foreignkey:ExecutedBy" json:"executedByDetail,omitempty"`
	StartedAt        types.NullTime         `gorm:"column:started_at" json:"startedAt,omitempty"`
	CompletedAt      types.NullTime         `gorm:"column:completed_at" json:"completedAt,omitempty"`
	CreatedAt        types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *ReportExecutions) TableName() string {
	return "analytics.report_executions"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ReportExecutionsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []ReportExecutions `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ReportExecutionsEdit struct {
	Name         *string         `gorm:"column:name" json:"name" example:"Necessitatibus eligendi."`
	ReportId     *types.URID     `gorm:"column:report_id;not null" json:"reportId" example:"EV3H62EINNCAHIXT52DJDY7RHE"`
	Status       *string         `gorm:"column:status" json:"status" example:"Harum occaecati."`
	ResultData   *types.JSON     `gorm:"column:result_data" json:"resultData,omitempty"`
	RowCount     *int            `gorm:"column:row_count" json:"rowCount" example:"4930781955944835665"`
	DurationMs   *int            `gorm:"column:duration_ms" json:"durationMs" example:"212122403209527124"`
	ErrorMessage *string         `gorm:"column:error_message" json:"errorMessage" example:"Quis sint."`
	ExecutedBy   *types.URID     `gorm:"column:executed_by" json:"executedBy,omitempty" example:"PUHFLI6ECRB33MQKQBOK6OIE4I"`
	StartedAt    *types.NullTime `gorm:"column:started_at" json:"startedAt,omitempty"`
	CompletedAt  *types.NullTime `gorm:"column:completed_at" json:"completedAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ReportExecutionsFilter struct {
	Id           *types.URID     `gorm:"column:id;primary_key" json:"id" example:"TGQDCY33KJFYLBSFVG7T63WB3U"`
	Name         *string         `gorm:"column:name" json:"name" example:"Sit ea."`
	ReportId     *types.URID     `gorm:"column:report_id;not null" json:"reportId" example:"IU4J2BBYUVEHHJEBIUDNQIXZTM"`
	Status       *string         `gorm:"column:status" json:"status" example:"Aperiam qui."`
	ResultData   *types.JSON     `gorm:"column:result_data" json:"resultData,omitempty"`
	RowCount     *int            `gorm:"column:row_count" json:"rowCount" example:"3196969024270793404"`
	DurationMs   *int            `gorm:"column:duration_ms" json:"durationMs" example:"-8664627165296447059"`
	ErrorMessage *string         `gorm:"column:error_message" json:"errorMessage" example:"Ea reiciendis."`
	ExecutedBy   *types.URID     `gorm:"column:executed_by" json:"executedBy,omitempty" example:"HTCJKS7DUNHWVCYVP6SVQUQT4I"`
	StartedAt    *types.NullTime `gorm:"column:started_at" json:"startedAt,omitempty"`
	CompletedAt  *types.NullTime `gorm:"column:completed_at" json:"completedAt,omitempty"`
	CreatedAt    *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *ReportExecutionsFilter) TableName() string {
	return "analytics.report_executions"
}

// --- Batch Update Struct ---
type ReportExecutionsBatchUpdate struct {
	Data       ReportExecutionsEdit     `json:"data"`
	PathParams ReportExecutionsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ReportExecutionsIdentity struct {
	Id types.URID `json:"id"`
}
