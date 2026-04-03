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
// • TeamsForm     - Data input validation and creation
// • Teams        - Main database model with relationships
// • TeamsEdit    - Partial update operations
// • TeamsIdentity - Bulk operation identifiers
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

// TeamsForm handles data input validation and creation operations.
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
type TeamsForm struct {
	Name           string            `gorm:"column:name;not null" json:"name" example:"Dolorum unde."`
	Description    string            `gorm:"column:description" json:"description" example:"Fugiat soluta."`
	OrganizationId types.URID        `gorm:"column:organization_id;not null" json:"organizationId" example:"JPWXGCUDL5BNBEHCSREB5YK7IY"`
	LeadId         *types.URID       `gorm:"column:lead_id" json:"leadId,omitempty" example:"SGRX4VOB5NAM3PHZN7VZRIEOBY"`
	Tags           types.StringArray `gorm:"column:tags" json:"tags"`
}

func (p *TeamsForm) TableName() string {
	return "iam.teams"
}

// Teams represents the main database model for iam.teams table.
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
// • Table: iam.teams
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
type Teams struct {
	Id                   types.URID        `gorm:"column:id;primary_key" json:"id" example:"KWH2ILWQHRAE5IP6QZRSWAU7GY"`
	Name                 string            `gorm:"column:name;not null" json:"name" example:"Aperiam est."`
	Description          string            `gorm:"column:description" json:"description" example:"Et voluptate."`
	OrganizationId       types.URID        `gorm:"column:organization_id;not null" json:"organizationId" example:"Y2S2QKKYE5CMLIH7XFXBUVYQAQ"`
	OrganizationIdDetail *Organizations    `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	LeadId               *types.URID       `gorm:"column:lead_id" json:"leadId,omitempty" example:"D7GW3PHSS5EERMPCRBBWKEQ25Y"`
	LeadIdDetail         *Users            `gorm:"foreignkey:LeadId" json:"leadDetail,omitempty"`
	Tags                 types.StringArray `gorm:"column:tags" json:"tags"`
	CreatedAt            types.NullTime    `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime    `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	TeamMembersList      *[]TeamMembers    `gorm:"foreignKey:TeamId;references:Id" json:"teamMembersList,omitempty"`
}

func (p *Teams) TableName() string {
	return "iam.teams"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type TeamsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Teams `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type TeamsEdit struct {
	Name           *string            `gorm:"column:name;not null" json:"name" example:"Sed esse."`
	Description    *string            `gorm:"column:description" json:"description" example:"Aut voluptatem."`
	OrganizationId *types.URID        `gorm:"column:organization_id;not null" json:"organizationId" example:"EEQZR4JCVRCD7FZOD7F7T6NH4Q"`
	LeadId         *types.URID        `gorm:"column:lead_id" json:"leadId,omitempty" example:"Z3XAZHPN3BCE3CEVYTPFC7FUDQ"`
	Tags           *types.StringArray `gorm:"column:tags" json:"tags"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type TeamsFilter struct {
	Id             *types.URID        `gorm:"column:id;primary_key" json:"id" example:"L7MVAKLWBNGIRLOFUPAHEZPK2Y"`
	Name           *string            `gorm:"column:name;not null" json:"name" example:"Quos est."`
	Description    *string            `gorm:"column:description" json:"description" example:"Sed necessitatibus."`
	OrganizationId *types.URID        `gorm:"column:organization_id;not null" json:"organizationId" example:"2WGBEYCBCZDJVCFQALR552ROP4"`
	LeadId         *types.URID        `gorm:"column:lead_id" json:"leadId,omitempty" example:"6IXR2SZURJGUPJUTAOWEYWUNNE"`
	Tags           *types.StringArray `gorm:"column:tags" json:"tags"`
	CreatedAt      *types.NullTime    `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime    `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *TeamsFilter) TableName() string {
	return "iam.teams"
}

// --- Batch Update Struct ---
type TeamsBatchUpdate struct {
	Data       TeamsEdit     `json:"data"`
	PathParams TeamsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type TeamsIdentity struct {
	Id types.URID `json:"id"`
}
