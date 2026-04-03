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
// • UnreadNotificationsForm     - Data input validation and creation
// • UnreadNotifications        - Main database model with relationships
// • UnreadNotificationsEdit    - Partial update operations
// • UnreadNotificationsIdentity - Bulk operation identifiers
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

// UnreadNotifications represents the main database model for analytics.unread_notifications table.
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
// • Table: analytics.unread_notifications
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
type UnreadNotifications struct {
	Id        *types.URID    `gorm:"column:id" json:"id,omitempty" example:"3DDN4SC5CJHAHBBPWAKCU7BPNQ"`
	Name      *string        `gorm:"column:name" json:"name,omitempty" example:"Velit ducimus."`
	UserId    *types.URID    `gorm:"column:user_id" json:"userId,omitempty" example:"VB6R23QSTFH4BKJPMAJKY4R7PI"`
	Title     *string        `gorm:"column:title" json:"title,omitempty" example:"Ad hic."`
	Message   *string        `gorm:"column:message" json:"message,omitempty" example:"Atque officiis."`
	Channel   *string        `gorm:"column:channel" json:"channel,omitempty" example:"Alias consequatur."`
	ActionUrl *string        `gorm:"column:action_url" json:"actionUrl,omitempty" example:"Ut corporis."`
	CreatedAt types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt,omitempty"`
}

func (p *UnreadNotifications) TableName() string {
	return "analytics.unread_notifications"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type UnreadNotificationsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []UnreadNotifications `json:"items"`
}
