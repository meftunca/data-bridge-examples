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
// • RolePermissionsForm     - Data input validation and creation
// • RolePermissions        - Main database model with relationships
// • RolePermissionsEdit    - Partial update operations
// • RolePermissionsIdentity - Bulk operation identifiers
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

// RolePermissionsForm handles data input validation and creation operations.
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
type RolePermissionsForm struct {
	Name         string     `gorm:"column:name" json:"name" example:"Magnam ut."`
	RoleId       types.URID `gorm:"column:role_id;not null" json:"roleId" example:"YE7TNPXXNVFNTNYI42I7P2H7PQ"`
	PermissionId types.URID `gorm:"column:permission_id;not null" json:"permissionId" example:"G4N3YBKQ3JCXBGBLREMS7EXUH4"`
}

func (p *RolePermissionsForm) TableName() string {
	return "iam.role_permissions"
}

// RolePermissions represents the main database model for iam.role_permissions table.
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
// • Table: iam.role_permissions
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
type RolePermissions struct {
	Id                 types.URID     `gorm:"column:id;primary_key" json:"id" example:"5KNO5DEHJNE5FFQHASTFOK7MOM"`
	Name               string         `gorm:"column:name" json:"name" example:"Occaecati aut."`
	RoleId             types.URID     `gorm:"column:role_id;not null" json:"roleId" example:"73O47IGRERH4JK2HLI3SCBK3TY"`
	RoleIdDetail       *Roles         `gorm:"foreignkey:RoleId" json:"roleDetail,omitempty"`
	PermissionId       types.URID     `gorm:"column:permission_id;not null" json:"permissionId" example:"XOIYYTSVORBALN3CGL2OYON2FI"`
	PermissionIdDetail *Permissions   `gorm:"foreignkey:PermissionId" json:"permissionDetail,omitempty"`
	CreatedAt          types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *RolePermissions) TableName() string {
	return "iam.role_permissions"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type RolePermissionsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []RolePermissions `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type RolePermissionsEdit struct {
	Name         *string     `gorm:"column:name" json:"name" example:"Omnis ut."`
	RoleId       *types.URID `gorm:"column:role_id;not null" json:"roleId" example:"2YQ75SI64BELDGOL2R3RO2JSRU"`
	PermissionId *types.URID `gorm:"column:permission_id;not null" json:"permissionId" example:"EFJE7PG3HVEWLLYS7CB6DHASAM"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type RolePermissionsFilter struct {
	Id           *types.URID     `gorm:"column:id;primary_key" json:"id" example:"4EDGSECGGNALFKT2TY43NKCSV4"`
	Name         *string         `gorm:"column:name" json:"name" example:"Omnis qui."`
	RoleId       *types.URID     `gorm:"column:role_id;not null" json:"roleId" example:"FJI2OSAZJBC3XBO6UGELLU5OA4"`
	PermissionId *types.URID     `gorm:"column:permission_id;not null" json:"permissionId" example:"KFZX5VRE5VBALF72JO2T32ZTE4"`
	CreatedAt    *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *RolePermissionsFilter) TableName() string {
	return "iam.role_permissions"
}

// --- Batch Update Struct ---
type RolePermissionsBatchUpdate struct {
	Data       RolePermissionsEdit     `json:"data"`
	PathParams RolePermissionsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type RolePermissionsIdentity struct {
	Id types.URID `json:"id"`
}
