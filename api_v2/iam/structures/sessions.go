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
// • SessionsForm     - Data input validation and creation
// • Sessions        - Main database model with relationships
// • SessionsEdit    - Partial update operations
// • SessionsIdentity - Bulk operation identifiers
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
package iam_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"
	"net"

	"github.com/maple-tech/baseline/types"
)

// SessionsForm handles data input validation and creation operations.
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
type SessionsForm struct {
	Name      string         `gorm:"column:name" json:"name" example:"Ea libero."`
	UserId    types.URID     `gorm:"column:user_id;not null" json:"userId" example:"O6KW7AMFZ5G2TEU3L5TSEZ3PFQ"`
	IpAddress *net.IP        `gorm:"column:ip_address" json:"ipAddress,omitempty"`
	UserAgent string         `gorm:"column:user_agent" json:"userAgent" example:"Laborum quia."`
	IsActive  bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	ExpiresAt types.NullTime `gorm:"column:expires_at;not null" json:"expiresAt"`
}

func (p *SessionsForm) TableName() string {
	return "iam.sessions"
}

// Sessions represents the main database model for iam.sessions table.
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
// • Table: iam.sessions
// • Type: BASE TABLE
// • Schema: iam
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
type Sessions struct {
	Id           types.URID     `gorm:"column:id;primary_key" json:"id" example:"R256C4PUSVFMFF4AIT4ZYXBGIM"`
	Name         string         `gorm:"column:name" json:"name" example:"Nisi doloribus."`
	UserId       types.URID     `gorm:"column:user_id;not null" json:"userId" example:"VJC4U4W6Z5BRJN6BIO7ANWVGPU"`
	UserIdDetail *Users         `gorm:"foreignkey:UserId" json:"userDetail,omitempty"`
	IpAddress    *net.IP        `gorm:"column:ip_address" json:"ipAddress,omitempty"`
	UserAgent    string         `gorm:"column:user_agent" json:"userAgent" example:"A praesentium."`
	IsActive     bool           `gorm:"column:is_active" json:"isActive" example:"true"`
	ExpiresAt    types.NullTime `gorm:"column:expires_at;not null" json:"expiresAt"`
	CreatedAt    types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *Sessions) TableName() string {
	return "iam.sessions"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type SessionsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Sessions `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type SessionsEdit struct {
	Name      *string         `gorm:"column:name" json:"name" example:"Dicta minima."`
	UserId    *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"7SUHFGMODVB43EZ5NC6TCUARZ4"`
	IpAddress *net.IP         `gorm:"column:ip_address" json:"ipAddress,omitempty"`
	UserAgent *string         `gorm:"column:user_agent" json:"userAgent" example:"Quia blanditiis."`
	IsActive  *bool           `gorm:"column:is_active" json:"isActive" example:"true"`
	ExpiresAt *types.NullTime `gorm:"column:expires_at;not null" json:"expiresAt"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type SessionsFilter struct {
	Id        *types.URID     `gorm:"column:id;primary_key" json:"id" example:"DFQFDOBUUNDZFJT5EU2SMXT5VI"`
	Name      *string         `gorm:"column:name" json:"name" example:"Consequuntur unde."`
	UserId    *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"TVHFRWSJTFFUNBDQ7JT5OZID4A"`
	IpAddress *net.IP         `gorm:"column:ip_address" json:"ipAddress,omitempty"`
	UserAgent *string         `gorm:"column:user_agent" json:"userAgent" example:"Eos ut."`
	IsActive  *bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	ExpiresAt *types.NullTime `gorm:"column:expires_at;not null" json:"expiresAt"`
	CreatedAt *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *SessionsFilter) TableName() string {
	return "iam.sessions"
}

// --- Batch Update Struct ---
type SessionsBatchUpdate struct {
	Data       SessionsEdit     `json:"data"`
	PathParams SessionsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type SessionsIdentity struct {
	Id types.URID `json:"id"`
}
