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
// • NotificationsForm     - Data input validation and creation
// • Notifications        - Main database model with relationships
// • NotificationsEdit    - Partial update operations
// • NotificationsIdentity - Bulk operation identifiers
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

// NotificationsForm handles data input validation and creation operations.
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
type NotificationsForm struct {
	Name      string         `gorm:"column:name" json:"name" example:"Voluptas et."`
	UserId    types.URID     `gorm:"column:user_id;not null" json:"userId" example:"JTXVJYLNLFEW7HRJBVA2SYAKDU"`
	Title     string         `gorm:"column:title" json:"title" example:"Est et."`
	Message   string         `gorm:"column:message" json:"message" example:"Natus quia."`
	Channel   string         `gorm:"column:channel" json:"channel" example:"Est porro."`
	IsRead    bool           `gorm:"column:is_read" json:"isRead" example:"false"`
	ActionUrl string         `gorm:"column:action_url" json:"actionUrl" example:"Quae animi."`
	Metadata  types.JSON     `gorm:"column:metadata" json:"metadata"`
	ReadAt    types.NullTime `gorm:"column:read_at" json:"readAt,omitempty"`
}

func (p *NotificationsForm) TableName() string {
	return "analytics.notifications"
}

// Notifications represents the main database model for analytics.notifications table.
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
// • Table: analytics.notifications
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
type Notifications struct {
	Id           types.URID             `gorm:"column:id;primary_key" json:"id" example:"QRFK3NF27FDMHE4D77TZMNHJAU"`
	Name         string                 `gorm:"column:name" json:"name" example:"Fugit ab."`
	UserId       types.URID             `gorm:"column:user_id;not null" json:"userId" example:"GKMX2NLTLZHLTKN2RLNRKPEVLQ"`
	UserIdDetail *shared_types.IamUsers `gorm:"foreignkey:UserId" json:"userDetail,omitempty"`
	Title        string                 `gorm:"column:title" json:"title" example:"Magnam omnis."`
	Message      string                 `gorm:"column:message" json:"message" example:"Soluta autem."`
	Channel      string                 `gorm:"column:channel" json:"channel" example:"Consequatur quisquam."`
	IsRead       bool                   `gorm:"column:is_read" json:"isRead" example:"false"`
	ActionUrl    string                 `gorm:"column:action_url" json:"actionUrl" example:"Vel dolorem."`
	Metadata     types.JSON             `gorm:"column:metadata" json:"metadata"`
	ReadAt       types.NullTime         `gorm:"column:read_at" json:"readAt,omitempty"`
	CreatedAt    types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *Notifications) TableName() string {
	return "analytics.notifications"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type NotificationsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Notifications `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type NotificationsEdit struct {
	Name      *string         `gorm:"column:name" json:"name" example:"Est consequuntur."`
	UserId    *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"LBP7MO646VESLACYUZJKPU7MGI"`
	Title     *string         `gorm:"column:title" json:"title" example:"Accusamus ipsum."`
	Message   *string         `gorm:"column:message" json:"message" example:"Ipsa quia."`
	Channel   *string         `gorm:"column:channel" json:"channel" example:"Quia et."`
	IsRead    *bool           `gorm:"column:is_read" json:"isRead" example:"false"`
	ActionUrl *string         `gorm:"column:action_url" json:"actionUrl" example:"Minima ullam."`
	Metadata  *types.JSON     `gorm:"column:metadata" json:"metadata"`
	ReadAt    *types.NullTime `gorm:"column:read_at" json:"readAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type NotificationsFilter struct {
	Id        *types.URID     `gorm:"column:id;primary_key" json:"id" example:"N6XZTOIZGBHNDKQ4RGKVMSFZWQ"`
	Name      *string         `gorm:"column:name" json:"name" example:"Et et."`
	UserId    *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"TA65NERU4BDRRD6QPK35BECZHQ"`
	Title     *string         `gorm:"column:title" json:"title" example:"Quia maiores."`
	Message   *string         `gorm:"column:message" json:"message" example:"Eum dicta."`
	Channel   *string         `gorm:"column:channel" json:"channel" example:"Libero nihil."`
	IsRead    *bool           `gorm:"column:is_read" json:"isRead" example:"true"`
	ActionUrl *string         `gorm:"column:action_url" json:"actionUrl" example:"Deleniti est."`
	Metadata  *types.JSON     `gorm:"column:metadata" json:"metadata"`
	ReadAt    *types.NullTime `gorm:"column:read_at" json:"readAt,omitempty"`
	CreatedAt *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *NotificationsFilter) TableName() string {
	return "analytics.notifications"
}

// --- Batch Update Struct ---
type NotificationsBatchUpdate struct {
	Data       NotificationsEdit     `json:"data"`
	PathParams NotificationsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type NotificationsIdentity struct {
	Id types.URID `json:"id"`
}
