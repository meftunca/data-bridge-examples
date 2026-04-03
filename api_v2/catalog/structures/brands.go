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
// • BrandsForm     - Data input validation and creation
// • Brands        - Main database model with relationships
// • BrandsEdit    - Partial update operations
// • BrandsIdentity - Bulk operation identifiers
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
package catalog_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"
	shared_types "data-bridge-examples/api_v2/shared/types"

	"github.com/maple-tech/baseline/types"
)

// BrandsForm handles data input validation and creation operations.
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
type BrandsForm struct {
	Name           string      `gorm:"column:name;not null" json:"name" example:"Architecto quisquam."`
	Slug           string      `gorm:"column:slug;not null" json:"slug" example:"Nihil aut."`
	LogoUrl        string      `gorm:"column:logo_url" json:"logoUrl" example:"Et modi."`
	Website        string      `gorm:"column:website" json:"website" example:"Libero eum."`
	Description    string      `gorm:"column:description" json:"description" example:"Itaque odit."`
	OrganizationId *types.URID `gorm:"column:organization_id" json:"organizationId,omitempty" example:"QBVNSOUD6NBWZHIWMY4R2L3M2E"`
}

func (p *BrandsForm) TableName() string {
	return "catalog.brands"
}

// Brands represents the main database model for catalog.brands table.
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
// • Table: catalog.brands
// • Type: BASE TABLE
// • Schema: catalog
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
type Brands struct {
	Id                   types.URID                     `gorm:"column:id;primary_key" json:"id" example:"RUDUXMMS7JAP7NOTIN55KRDQXQ"`
	Name                 string                         `gorm:"column:name;not null" json:"name" example:"Minus quia."`
	Slug                 string                         `gorm:"column:slug;not null" json:"slug" example:"Atque blanditiis."`
	LogoUrl              string                         `gorm:"column:logo_url" json:"logoUrl" example:"Quisquam quae."`
	Website              string                         `gorm:"column:website" json:"website" example:"Sunt ab."`
	Description          string                         `gorm:"column:description" json:"description" example:"Et dolorem."`
	OrganizationId       *types.URID                    `gorm:"column:organization_id" json:"organizationId,omitempty" example:"P74DRKSZMFDHBPJQDY7Z3D2KCE"`
	OrganizationIdDetail *shared_types.IamOrganizations `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	CreatedAt            types.NullTime                 `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime                 `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	ProductsList         *[]Products                    `gorm:"foreignKey:BrandId;references:Id" json:"productsList,omitempty"`
}

func (p *Brands) TableName() string {
	return "catalog.brands"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type BrandsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Brands `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type BrandsEdit struct {
	Name           *string     `gorm:"column:name;not null" json:"name" example:"Et aut."`
	Slug           *string     `gorm:"column:slug;not null" json:"slug" example:"Excepturi quasi."`
	LogoUrl        *string     `gorm:"column:logo_url" json:"logoUrl" example:"Iste quo."`
	Website        *string     `gorm:"column:website" json:"website" example:"Quo et."`
	Description    *string     `gorm:"column:description" json:"description" example:"Sequi aut."`
	OrganizationId *types.URID `gorm:"column:organization_id" json:"organizationId,omitempty" example:"JCQ2MZZUGNAEZP2BHZ4DNROSLI"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type BrandsFilter struct {
	Id             *types.URID     `gorm:"column:id;primary_key" json:"id" example:"ALWHCYGO3BDPLOCM5CKJYOGFYI"`
	Name           *string         `gorm:"column:name;not null" json:"name" example:"Sunt voluptate."`
	Slug           *string         `gorm:"column:slug;not null" json:"slug" example:"Molestias rerum."`
	LogoUrl        *string         `gorm:"column:logo_url" json:"logoUrl" example:"Eum voluptatum."`
	Website        *string         `gorm:"column:website" json:"website" example:"Omnis vero."`
	Description    *string         `gorm:"column:description" json:"description" example:"Aliquid numquam."`
	OrganizationId *types.URID     `gorm:"column:organization_id" json:"organizationId,omitempty" example:"CYXARZAC35B4JMEMLPJO7UDISY"`
	CreatedAt      *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *BrandsFilter) TableName() string {
	return "catalog.brands"
}

// --- Batch Update Struct ---
type BrandsBatchUpdate struct {
	Data       BrandsEdit     `json:"data"`
	PathParams BrandsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type BrandsIdentity struct {
	Id types.URID `json:"id"`
}
