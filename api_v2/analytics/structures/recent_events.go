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
// • RecentEventsForm     - Data input validation and creation
// • RecentEvents        - Main database model with relationships
// • RecentEventsEdit    - Partial update operations
// • RecentEventsIdentity - Bulk operation identifiers
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

// RecentEvents represents the main database model for analytics.recent_events table.
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
// • Table: analytics.recent_events
// • Type: VIEW
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
type RecentEvents struct {
	Id           *types.URID             `gorm:"column:id" json:"id,omitempty" example:"2VJICHMXBREZTOJSLXMWXQ3VL4"`
	Name         *string                 `gorm:"column:name" json:"name,omitempty" example:"Explicabo in."`
	EventType    *string                 `gorm:"column:event_type" json:"eventType,omitempty" example:"Quod soluta."`
	SourceSchema *string                 `gorm:"column:source_schema" json:"sourceSchema,omitempty" example:"Quisquam in."`
	SourceTable  *string                 `gorm:"column:source_table" json:"sourceTable,omitempty" example:"Ut expedita."`
	ActorId      *types.URID             `gorm:"column:actor_id" json:"actorId,omitempty" example:"HKSH2SWO6FDOVCAAFL67JLYMVU"`
	Severity     *AnalyticsEventSeverity `gorm:"column:severity" json:"severity,omitempty"`
	CreatedAt    types.NullTime          `gorm:"column:created_at;autoCreateTime" json:"createdAt,omitempty"`
}

func (p *RecentEvents) TableName() string {
	return "analytics.recent_events"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type RecentEventsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []RecentEvents `json:"items"`
}
