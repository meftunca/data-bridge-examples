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
// • InvitationsForm     - Data input validation and creation
// • Invitations        - Main database model with relationships
// • InvitationsEdit    - Partial update operations
// • InvitationsIdentity - Bulk operation identifiers
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

// InvitationsForm handles data input validation and creation operations.
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
type InvitationsForm struct {
	Name           string         `gorm:"column:name" json:"name" example:"Enim consequatur."`
	Email          string         `gorm:"column:email;not null" json:"email" example:"Harum veritatis."`
	OrganizationId types.URID     `gorm:"column:organization_id;not null" json:"organizationId" example:"N52IVCYF2VF55INV5T5QIRSAAQ"`
	InvitedBy      types.URID     `gorm:"column:invited_by;not null" json:"invitedBy" example:"HKBWZYDTINBSNIVGA4AUH7DYMI"`
	RoleId         *types.URID    `gorm:"column:role_id" json:"roleId,omitempty" example:"NUTWLKASOBFMPHCAWA4C3VVDMI"`
	Token          string         `gorm:"column:token;not null" json:"token" example:"Iusto a."`
	AcceptedAt     types.NullTime `gorm:"column:accepted_at" json:"acceptedAt,omitempty"`
	ExpiresAt      types.NullTime `gorm:"column:expires_at;not null" json:"expiresAt"`
}

func (p *InvitationsForm) TableName() string {
	return "iam.invitations"
}

// Invitations represents the main database model for iam.invitations table.
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
// • Table: iam.invitations
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
type Invitations struct {
	Id                   types.URID     `gorm:"column:id;primary_key" json:"id" example:"WPEC3BOAVBD5NBTWZWZ4IFKQEY"`
	Name                 string         `gorm:"column:name" json:"name" example:"Architecto accusamus."`
	Email                string         `gorm:"column:email;not null" json:"email" example:"Voluptas tempora."`
	OrganizationId       types.URID     `gorm:"column:organization_id;not null" json:"organizationId" example:"A5GDZOV5ORFVBPLNQXYK2XQ2PE"`
	OrganizationIdDetail *Organizations `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	InvitedBy            types.URID     `gorm:"column:invited_by;not null" json:"invitedBy" example:"53UITEOLC5FRBKA6DIUW23EEGU"`
	InvitedByDetail      *Users         `gorm:"foreignkey:InvitedBy" json:"invitedByDetail,omitempty"`
	RoleId               *types.URID    `gorm:"column:role_id" json:"roleId,omitempty" example:"PXVHGT6LDZCLRFMMQ5CFRX5PKQ"`
	RoleIdDetail         *Roles         `gorm:"foreignkey:RoleId" json:"roleDetail,omitempty"`
	Token                string         `gorm:"column:token;not null" json:"token" example:"Ipsa dolorem."`
	AcceptedAt           types.NullTime `gorm:"column:accepted_at" json:"acceptedAt,omitempty"`
	ExpiresAt            types.NullTime `gorm:"column:expires_at;not null" json:"expiresAt"`
	CreatedAt            types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *Invitations) TableName() string {
	return "iam.invitations"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type InvitationsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Invitations `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type InvitationsEdit struct {
	Name           *string         `gorm:"column:name" json:"name" example:"Eos iure."`
	Email          *string         `gorm:"column:email;not null" json:"email" example:"Quibusdam maiores."`
	OrganizationId *types.URID     `gorm:"column:organization_id;not null" json:"organizationId" example:"ZAJ6ZW6DAJHLNIR3UOW7XD7QBE"`
	InvitedBy      *types.URID     `gorm:"column:invited_by;not null" json:"invitedBy" example:"GL4M7OU6IRBOXNG4I6BA4F3CNE"`
	RoleId         *types.URID     `gorm:"column:role_id" json:"roleId,omitempty" example:"GFHVB5J2KNE3TM7VRBKX3BIYWE"`
	Token          *string         `gorm:"column:token;not null" json:"token" example:"Enim aut."`
	AcceptedAt     *types.NullTime `gorm:"column:accepted_at" json:"acceptedAt,omitempty"`
	ExpiresAt      *types.NullTime `gorm:"column:expires_at;not null" json:"expiresAt"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type InvitationsFilter struct {
	Id             *types.URID     `gorm:"column:id;primary_key" json:"id" example:"N7RARL2TIRACNAXQWTQMCQTAZ4"`
	Name           *string         `gorm:"column:name" json:"name" example:"Non fugiat."`
	Email          *string         `gorm:"column:email;not null" json:"email" example:"Aut aliquam."`
	OrganizationId *types.URID     `gorm:"column:organization_id;not null" json:"organizationId" example:"XRHX4LSIGNFRLMVJUAIS4W427E"`
	InvitedBy      *types.URID     `gorm:"column:invited_by;not null" json:"invitedBy" example:"BHYXLXRU3RCV5HFKAZKC4S3OIU"`
	RoleId         *types.URID     `gorm:"column:role_id" json:"roleId,omitempty" example:"2BXJTXVWPJB7JCUIAZKKSLWJBM"`
	Token          *string         `gorm:"column:token;not null" json:"token" example:"Necessitatibus eum."`
	AcceptedAt     *types.NullTime `gorm:"column:accepted_at" json:"acceptedAt,omitempty"`
	ExpiresAt      *types.NullTime `gorm:"column:expires_at;not null" json:"expiresAt"`
	CreatedAt      *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *InvitationsFilter) TableName() string {
	return "iam.invitations"
}

// --- Batch Update Struct ---
type InvitationsBatchUpdate struct {
	Data       InvitationsEdit     `json:"data"`
	PathParams InvitationsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type InvitationsIdentity struct {
	Id types.URID `json:"id"`
}
