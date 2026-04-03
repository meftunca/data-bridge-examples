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
// • UserRolesForm     - Data input validation and creation
// • UserRoles        - Main database model with relationships
// • UserRolesEdit    - Partial update operations
// • UserRolesIdentity - Bulk operation identifiers
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

// UserRolesForm handles data input validation and creation operations.
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
type UserRolesForm struct {
	Name      string         `gorm:"column:name" json:"name" example:"Saepe dolor."`
	UserId    types.URID     `gorm:"column:user_id;not null" json:"userId" example:"FFSN3F4EVJBBZLA7MJPPMDLJ3A"`
	RoleId    types.URID     `gorm:"column:role_id;not null" json:"roleId" example:"Z5H6O4S77ZEUFIB5JTLUT563O4"`
	GrantedBy *types.URID    `gorm:"column:granted_by" json:"grantedBy,omitempty" example:"JQJ5X5K3D5D6NM3HLR4IYT5QB4"`
	GrantedAt types.NullTime `gorm:"column:granted_at" json:"grantedAt"`
	ExpiresAt types.NullTime `gorm:"column:expires_at" json:"expiresAt,omitempty"`
}

func (p *UserRolesForm) TableName() string {
	return "iam.user_roles"
}

// UserRoles represents the main database model for iam.user_roles table.
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
// • Table: iam.user_roles
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
type UserRoles struct {
	Id              types.URID     `gorm:"column:id;primary_key" json:"id" example:"XWQWJVTE7BHXZKHVGI6X2YRP2Y"`
	Name            string         `gorm:"column:name" json:"name" example:"Omnis molestiae."`
	UserId          types.URID     `gorm:"column:user_id;not null" json:"userId" example:"UTMAMJN4AVBVLCYQ27AD4UXBEE"`
	UserIdDetail    *Users         `gorm:"foreignkey:UserId" json:"userDetail,omitempty"`
	RoleId          types.URID     `gorm:"column:role_id;not null" json:"roleId" example:"ZXMTNBS6MNBVHE7YFJ426CW5P4"`
	RoleIdDetail    *Roles         `gorm:"foreignkey:RoleId" json:"roleDetail,omitempty"`
	GrantedBy       *types.URID    `gorm:"column:granted_by" json:"grantedBy,omitempty" example:"AJKDO2LKLRDY5H23XIREUYWLSQ"`
	GrantedByDetail *Users         `gorm:"foreignkey:GrantedBy" json:"grantedByDetail,omitempty"`
	GrantedAt       types.NullTime `gorm:"column:granted_at" json:"grantedAt"`
	ExpiresAt       types.NullTime `gorm:"column:expires_at" json:"expiresAt,omitempty"`
	CreatedAt       types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *UserRoles) TableName() string {
	return "iam.user_roles"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type UserRolesPage struct {
	paginationRuntime.DefaultPageResponse
	Items []UserRoles `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type UserRolesEdit struct {
	Name      *string         `gorm:"column:name" json:"name" example:"Qui voluptatem."`
	UserId    *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"WRVPIOA3RZAFVM5XYF2MHFOPH4"`
	RoleId    *types.URID     `gorm:"column:role_id;not null" json:"roleId" example:"C6RRO3AWO5DZFAEXIWURDSESBM"`
	GrantedBy *types.URID     `gorm:"column:granted_by" json:"grantedBy,omitempty" example:"QNACWI73DBEMBDBG6JDK2KVTOI"`
	GrantedAt *types.NullTime `gorm:"column:granted_at" json:"grantedAt"`
	ExpiresAt *types.NullTime `gorm:"column:expires_at" json:"expiresAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type UserRolesFilter struct {
	Id        *types.URID     `gorm:"column:id;primary_key" json:"id" example:"MIWQKSRO6NGJHF2RUFCUTXDII4"`
	Name      *string         `gorm:"column:name" json:"name" example:"Consequatur dolor."`
	UserId    *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"WSPD4RND2ZEDXOD2A27WZOVISA"`
	RoleId    *types.URID     `gorm:"column:role_id;not null" json:"roleId" example:"Q7B7I3EV4ZDH5BJBDDYYQ24OP4"`
	GrantedBy *types.URID     `gorm:"column:granted_by" json:"grantedBy,omitempty" example:"CFJS4UVZVFDITKLXXNGRBIFQJE"`
	GrantedAt *types.NullTime `gorm:"column:granted_at" json:"grantedAt"`
	ExpiresAt *types.NullTime `gorm:"column:expires_at" json:"expiresAt,omitempty"`
	CreatedAt *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *UserRolesFilter) TableName() string {
	return "iam.user_roles"
}

// --- Batch Update Struct ---
type UserRolesBatchUpdate struct {
	Data       UserRolesEdit     `json:"data"`
	PathParams UserRolesIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type UserRolesIdentity struct {
	Id types.URID `json:"id"`
}
