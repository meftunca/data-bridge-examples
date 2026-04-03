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
// • MetricsForm     - Data input validation and creation
// • Metrics        - Main database model with relationships
// • MetricsEdit    - Partial update operations
// • MetricsIdentity - Bulk operation identifiers
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

// MetricsForm handles data input validation and creation operations.
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
type MetricsForm struct {
	Name       string         `gorm:"column:name;not null" json:"name" example:"Molestiae exercitationem."`
	MetricKey  string         `gorm:"column:metric_key;not null" json:"metricKey" example:"Ex nostrum."`
	Value      float64        `gorm:"column:value" json:"value"`
	Dimensions types.JSON     `gorm:"column:dimensions" json:"dimensions"`
	RecordedAt types.NullTime `gorm:"column:recorded_at" json:"recordedAt"`
}

func (p *MetricsForm) TableName() string {
	return "analytics.metrics"
}

// Metrics represents the main database model for analytics.metrics table.
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
// • Table: analytics.metrics
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
type Metrics struct {
	Id         types.URID     `gorm:"column:id;primary_key" json:"id" example:"FFIUPVR5H5HQZMG7ZNVHV6QRLQ"`
	Name       string         `gorm:"column:name;not null" json:"name" example:"Itaque dolorum."`
	MetricKey  string         `gorm:"column:metric_key;not null" json:"metricKey" example:"Eligendi laboriosam."`
	Value      float64        `gorm:"column:value" json:"value"`
	Dimensions types.JSON     `gorm:"column:dimensions" json:"dimensions"`
	RecordedAt types.NullTime `gorm:"column:recorded_at" json:"recordedAt"`
	CreatedAt  types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *Metrics) TableName() string {
	return "analytics.metrics"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type MetricsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Metrics `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type MetricsEdit struct {
	Name       *string         `gorm:"column:name;not null" json:"name" example:"Assumenda debitis."`
	MetricKey  *string         `gorm:"column:metric_key;not null" json:"metricKey" example:"A temporibus."`
	Value      *float64        `gorm:"column:value" json:"value"`
	Dimensions *types.JSON     `gorm:"column:dimensions" json:"dimensions"`
	RecordedAt *types.NullTime `gorm:"column:recorded_at" json:"recordedAt"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type MetricsFilter struct {
	Id         *types.URID     `gorm:"column:id;primary_key" json:"id" example:"MSWEPITGFFDFFEREUMYC57YIU4"`
	Name       *string         `gorm:"column:name;not null" json:"name" example:"Veniam assumenda."`
	MetricKey  *string         `gorm:"column:metric_key;not null" json:"metricKey" example:"Beatae quisquam."`
	Value      *float64        `gorm:"column:value" json:"value"`
	Dimensions *types.JSON     `gorm:"column:dimensions" json:"dimensions"`
	RecordedAt *types.NullTime `gorm:"column:recorded_at" json:"recordedAt"`
	CreatedAt  *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *MetricsFilter) TableName() string {
	return "analytics.metrics"
}

// --- Batch Update Struct ---
type MetricsBatchUpdate struct {
	Data       MetricsEdit     `json:"data"`
	PathParams MetricsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type MetricsIdentity struct {
	Id types.URID `json:"id"`
}
