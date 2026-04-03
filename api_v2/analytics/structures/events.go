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
// • EventsForm     - Data input validation and creation
// • Events        - Main database model with relationships
// • EventsEdit    - Partial update operations
// • EventsIdentity - Bulk operation identifiers
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

// EventsForm handles data input validation and creation operations.
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
type EventsForm struct {
	Name         string                 `gorm:"column:name" json:"name" example:"Dolores qui."`
	EventType    string                 `gorm:"column:event_type" json:"eventType" example:"Error temporibus."`
	SourceSchema string                 `gorm:"column:source_schema" json:"sourceSchema" example:"Esse sit."`
	SourceTable  string                 `gorm:"column:source_table" json:"sourceTable" example:"Illum blanditiis."`
	SourceId     string                 `gorm:"column:source_id" json:"sourceId" example:"Vitae quasi."`
	ActorId      *types.URID            `gorm:"column:actor_id" json:"actorId,omitempty" example:"26373JQ4TJG27J2IZ6BU57NFVE"`
	Payload      types.JSON             `gorm:"column:payload" json:"payload"`
	Severity     AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	Processed    bool                   `gorm:"column:processed" json:"processed" example:"false"`
	ProcessedAt  types.NullTime         `gorm:"column:processed_at" json:"processedAt,omitempty"`
}

func (p *EventsForm) TableName() string {
	return "analytics.events"
}

// Events represents the main database model for analytics.events table.
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
// • Table: analytics.events
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
type Events struct {
	Id            types.URID             `gorm:"column:id;primary_key" json:"id" example:"Y4C74RDBKFFEXJ37R4RQDHYX5A"`
	Name          string                 `gorm:"column:name" json:"name" example:"Ipsam dolores."`
	EventType     string                 `gorm:"column:event_type" json:"eventType" example:"Et voluptatibus."`
	SourceSchema  string                 `gorm:"column:source_schema" json:"sourceSchema" example:"Quae iusto."`
	SourceTable   string                 `gorm:"column:source_table" json:"sourceTable" example:"Cupiditate itaque."`
	SourceId      string                 `gorm:"column:source_id" json:"sourceId" example:"Nostrum voluptas."`
	ActorId       *types.URID            `gorm:"column:actor_id" json:"actorId,omitempty" example:"COY37FBTHFHN7DL3XCRX7CTO6U"`
	ActorIdDetail *shared_types.IamUsers `gorm:"foreignkey:ActorId" json:"actorDetail,omitempty"`
	Payload       types.JSON             `gorm:"column:payload" json:"payload"`
	Severity      AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	Processed     bool                   `gorm:"column:processed" json:"processed" example:"false"`
	ProcessedAt   types.NullTime         `gorm:"column:processed_at" json:"processedAt,omitempty"`
	CreatedAt     types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *Events) TableName() string {
	return "analytics.events"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type EventsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Events `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type EventsEdit struct {
	Name         *string                 `gorm:"column:name" json:"name" example:"Et in."`
	EventType    *string                 `gorm:"column:event_type" json:"eventType" example:"Suscipit soluta."`
	SourceSchema *string                 `gorm:"column:source_schema" json:"sourceSchema" example:"Beatae ut."`
	SourceTable  *string                 `gorm:"column:source_table" json:"sourceTable" example:"Officiis necessitatibus."`
	SourceId     *string                 `gorm:"column:source_id" json:"sourceId" example:"Ullam aut."`
	ActorId      *types.URID             `gorm:"column:actor_id" json:"actorId,omitempty" example:"LZIPKWYKAZFJPB67AQZWDFRXWU"`
	Payload      *types.JSON             `gorm:"column:payload" json:"payload"`
	Severity     *AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	Processed    *bool                   `gorm:"column:processed" json:"processed" example:"true"`
	ProcessedAt  *types.NullTime         `gorm:"column:processed_at" json:"processedAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type EventsFilter struct {
	Id           *types.URID             `gorm:"column:id;primary_key" json:"id" example:"6LYHDHCNI5H6ND3M2MACUOWQEU"`
	Name         *string                 `gorm:"column:name" json:"name" example:"Aperiam aliquid."`
	EventType    *string                 `gorm:"column:event_type" json:"eventType" example:"At autem."`
	SourceSchema *string                 `gorm:"column:source_schema" json:"sourceSchema" example:"Et nisi."`
	SourceTable  *string                 `gorm:"column:source_table" json:"sourceTable" example:"Sit voluptate."`
	SourceId     *string                 `gorm:"column:source_id" json:"sourceId" example:"In aut."`
	ActorId      *types.URID             `gorm:"column:actor_id" json:"actorId,omitempty" example:"QT37MFMMJBDVVEGMA3H6RFJV2U"`
	Payload      *types.JSON             `gorm:"column:payload" json:"payload"`
	Severity     *AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	Processed    *bool                   `gorm:"column:processed" json:"processed" example:"true"`
	ProcessedAt  *types.NullTime         `gorm:"column:processed_at" json:"processedAt,omitempty"`
	CreatedAt    *types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *EventsFilter) TableName() string {
	return "analytics.events"
}

// --- Batch Update Struct ---
type EventsBatchUpdate struct {
	Data       EventsEdit     `json:"data"`
	PathParams EventsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type EventsIdentity struct {
	Id types.URID `json:"id"`
}
