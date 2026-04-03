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
// • RolesForm     - Data input validation and creation
// • Roles        - Main database model with relationships
// • RolesEdit    - Partial update operations
// • RolesIdentity - Bulk operation identifiers
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

// RolesForm handles data input validation and creation operations.
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
type RolesForm struct {
	Name           string      `gorm:"column:name;not null" json:"name" example:"Consectetur ea."`
	Slug           string      `gorm:"column:slug;not null" json:"slug" example:"Dolor dolorem."`
	Description    string      `gorm:"column:description" json:"description" example:"Aspernatur nesciunt."`
	OrganizationId *types.URID `gorm:"column:organization_id" json:"organizationId,omitempty" example:"3SZIZ55GVNCALBCLHPSR3HXGZU"`
	IsSystem       bool        `gorm:"column:is_system" json:"isSystem" example:"false"`
}

func (p *RolesForm) TableName() string {
	return "iam.roles"
}

// Roles represents the main database model for iam.roles table.
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
// • Table: iam.roles
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
type Roles struct {
	Id                   types.URID         `gorm:"column:id;primary_key" json:"id" example:"QSKQZ6MLP5CTPEHCNAVQRX5QK4"`
	Name                 string             `gorm:"column:name;not null" json:"name" example:"Repellendus non."`
	Slug                 string             `gorm:"column:slug;not null" json:"slug" example:"Velit quisquam."`
	Description          string             `gorm:"column:description" json:"description" example:"Possimus et."`
	OrganizationId       *types.URID        `gorm:"column:organization_id" json:"organizationId,omitempty" example:"2KR4RIB4EVGMVKB463L5VVOLZY"`
	OrganizationIdDetail *Organizations     `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	IsSystem             bool               `gorm:"column:is_system" json:"isSystem" example:"false"`
	CreatedAt            types.NullTime     `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime     `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	RolePermissionsList  *[]RolePermissions `gorm:"foreignKey:RoleId;references:Id" json:"rolePermissionsList,omitempty"`
	UserRolesList        *[]UserRoles       `gorm:"foreignKey:RoleId;references:Id" json:"userRolesList,omitempty"`
	InvitationsList      *[]Invitations     `gorm:"foreignKey:RoleId;references:Id" json:"invitationsList,omitempty"`
}

func (p *Roles) TableName() string {
	return "iam.roles"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type RolesPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Roles `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type RolesEdit struct {
	Name           *string     `gorm:"column:name;not null" json:"name" example:"Qui omnis."`
	Slug           *string     `gorm:"column:slug;not null" json:"slug" example:"Aliquam placeat."`
	Description    *string     `gorm:"column:description" json:"description" example:"Accusantium dicta."`
	OrganizationId *types.URID `gorm:"column:organization_id" json:"organizationId,omitempty" example:"CHM5YFJ2VZCNNKOKAL2KUS7QUA"`
	IsSystem       *bool       `gorm:"column:is_system" json:"isSystem" example:"true"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type RolesFilter struct {
	Id             *types.URID     `gorm:"column:id;primary_key" json:"id" example:"NQDRHI7ATJEV3HVRCQABVYP34A"`
	Name           *string         `gorm:"column:name;not null" json:"name" example:"Voluptas est."`
	Slug           *string         `gorm:"column:slug;not null" json:"slug" example:"Accusantium aliquid."`
	Description    *string         `gorm:"column:description" json:"description" example:"Quia eum."`
	OrganizationId *types.URID     `gorm:"column:organization_id" json:"organizationId,omitempty" example:"TUWNF5TEDVBVPECSTZ5YLUIP2U"`
	IsSystem       *bool           `gorm:"column:is_system" json:"isSystem" example:"false"`
	CreatedAt      *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *RolesFilter) TableName() string {
	return "iam.roles"
}

// --- Batch Update Struct ---
type RolesBatchUpdate struct {
	Data       RolesEdit     `json:"data"`
	PathParams RolesIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type RolesIdentity struct {
	Id types.URID `json:"id"`
}
