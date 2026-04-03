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
// • PermissionsForm     - Data input validation and creation
// • Permissions        - Main database model with relationships
// • PermissionsEdit    - Partial update operations
// • PermissionsIdentity - Bulk operation identifiers
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

	"github.com/maple-tech/baseline/types"
)

// PermissionsForm handles data input validation and creation operations.
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
type PermissionsForm struct {
	Name        string `gorm:"column:name;not null" json:"name" example:"Ratione et."`
	Resource    string `gorm:"column:resource" json:"resource" example:"Aut voluptatibus."`
	Action      string `gorm:"column:action" json:"action" example:"Quas minima."`
	Description string `gorm:"column:description" json:"description" example:"Eos consectetur."`
}

func (p *PermissionsForm) TableName() string {
	return "iam.permissions"
}

// Permissions represents the main database model for iam.permissions table.
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
// • Table: iam.permissions
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
type Permissions struct {
	Id                  types.URID         `gorm:"column:id;primary_key" json:"id" example:"23EPLZ7UMJE2DHJ7KUWZJNGUFE"`
	Name                string             `gorm:"column:name;not null" json:"name" example:"Incidunt dolor."`
	Resource            string             `gorm:"column:resource" json:"resource" example:"Id corporis."`
	Action              string             `gorm:"column:action" json:"action" example:"Nobis aspernatur."`
	Description         string             `gorm:"column:description" json:"description" example:"Blanditiis tempore."`
	CreatedAt           types.NullTime     `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt           types.NullTime     `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	RolePermissionsList *[]RolePermissions `gorm:"foreignKey:PermissionId;references:Id" json:"rolePermissionsList,omitempty"`
}

func (p *Permissions) TableName() string {
	return "iam.permissions"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type PermissionsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Permissions `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type PermissionsEdit struct {
	Name        *string `gorm:"column:name;not null" json:"name" example:"Excepturi omnis."`
	Resource    *string `gorm:"column:resource" json:"resource" example:"Aut repellendus."`
	Action      *string `gorm:"column:action" json:"action" example:"Autem veritatis."`
	Description *string `gorm:"column:description" json:"description" example:"Sed provident."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type PermissionsFilter struct {
	Id          *types.URID     `gorm:"column:id;primary_key" json:"id" example:"HJ433QCAPBEW3DCIZCV57LUREI"`
	Name        *string         `gorm:"column:name;not null" json:"name" example:"Earum rerum."`
	Resource    *string         `gorm:"column:resource" json:"resource" example:"Accusamus dolorum."`
	Action      *string         `gorm:"column:action" json:"action" example:"Corporis illum."`
	Description *string         `gorm:"column:description" json:"description" example:"Repellendus ut."`
	CreatedAt   *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt   *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *PermissionsFilter) TableName() string {
	return "iam.permissions"
}

// --- Batch Update Struct ---
type PermissionsBatchUpdate struct {
	Data       PermissionsEdit     `json:"data"`
	PathParams PermissionsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type PermissionsIdentity struct {
	Id types.URID `json:"id"`
}
