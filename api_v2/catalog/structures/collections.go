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
// • CollectionsForm     - Data input validation and creation
// • Collections        - Main database model with relationships
// • CollectionsEdit    - Partial update operations
// • CollectionsIdentity - Bulk operation identifiers
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

// CollectionsForm handles data input validation and creation operations.
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
type CollectionsForm struct {
	Name          string      `gorm:"column:name;not null" json:"name" example:"Quia quod."`
	Slug          string      `gorm:"column:slug;not null" json:"slug" example:"Magni ratione."`
	Description   string      `gorm:"column:description" json:"description" example:"Vitae fugit."`
	CoverImageUrl string      `gorm:"column:cover_image_url" json:"coverImageUrl" example:"Non cumque."`
	IsActive      bool        `gorm:"column:is_active" json:"isActive" example:"false"`
	StartDate     *types.Date `gorm:"column:start_date" json:"startDate,omitempty"`
	EndDate       *types.Date `gorm:"column:end_date" json:"endDate,omitempty"`
	CreatedBy     *types.URID `gorm:"column:created_by" json:"createdBy,omitempty" example:"LFMFJIJKXZBXBFHE5KLTQE2LHY"`
}

func (p *CollectionsForm) TableName() string {
	return "catalog.collections"
}

// Collections represents the main database model for catalog.collections table.
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
// • Table: catalog.collections
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
type Collections struct {
	Id                     types.URID             `gorm:"column:id;primary_key" json:"id" example:"JW6IM34KMJBYVEWQJ7SOBVPYCM"`
	Name                   string                 `gorm:"column:name;not null" json:"name" example:"Ab nihil."`
	Slug                   string                 `gorm:"column:slug;not null" json:"slug" example:"Vel voluptate."`
	Description            string                 `gorm:"column:description" json:"description" example:"Quis voluptas."`
	CoverImageUrl          string                 `gorm:"column:cover_image_url" json:"coverImageUrl" example:"Perspiciatis est."`
	IsActive               bool                   `gorm:"column:is_active" json:"isActive" example:"true"`
	StartDate              *types.Date            `gorm:"column:start_date" json:"startDate,omitempty"`
	EndDate                *types.Date            `gorm:"column:end_date" json:"endDate,omitempty"`
	CreatedBy              *types.URID            `gorm:"column:created_by" json:"createdBy,omitempty" example:"OKNVVF3NM5BH7IFWUYGRFWMDKM"`
	CreatedByDetail        *shared_types.IamUsers `gorm:"foreignkey:CreatedBy" json:"createdByDetail,omitempty"`
	CreatedAt              types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt              types.NullTime         `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	CollectionProductsList *[]CollectionProducts  `gorm:"foreignKey:CollectionId;references:Id" json:"collectionProductsList,omitempty"`
}

func (p *Collections) TableName() string {
	return "catalog.collections"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type CollectionsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Collections `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type CollectionsEdit struct {
	Name          *string     `gorm:"column:name;not null" json:"name" example:"Blanditiis aut."`
	Slug          *string     `gorm:"column:slug;not null" json:"slug" example:"Quia fuga."`
	Description   *string     `gorm:"column:description" json:"description" example:"Rerum est."`
	CoverImageUrl *string     `gorm:"column:cover_image_url" json:"coverImageUrl" example:"Magnam voluptatum."`
	IsActive      *bool       `gorm:"column:is_active" json:"isActive" example:"false"`
	StartDate     *types.Date `gorm:"column:start_date" json:"startDate,omitempty"`
	EndDate       *types.Date `gorm:"column:end_date" json:"endDate,omitempty"`
	CreatedBy     *types.URID `gorm:"column:created_by" json:"createdBy,omitempty" example:"GISGERD22JEX3AZ2AB7F44TW2U"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type CollectionsFilter struct {
	Id            *types.URID     `gorm:"column:id;primary_key" json:"id" example:"KATWLINQPNAENFSTZKCFTLL2QY"`
	Name          *string         `gorm:"column:name;not null" json:"name" example:"Corporis neque."`
	Slug          *string         `gorm:"column:slug;not null" json:"slug" example:"Id exercitationem."`
	Description   *string         `gorm:"column:description" json:"description" example:"Dolores voluptatibus."`
	CoverImageUrl *string         `gorm:"column:cover_image_url" json:"coverImageUrl" example:"Id necessitatibus."`
	IsActive      *bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	StartDate     *types.Date     `gorm:"column:start_date" json:"startDate,omitempty"`
	EndDate       *types.Date     `gorm:"column:end_date" json:"endDate,omitempty"`
	CreatedBy     *types.URID     `gorm:"column:created_by" json:"createdBy,omitempty" example:"6OEYAPZIDFBRVG6S3VCNBLVM4A"`
	CreatedAt     *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt     *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *CollectionsFilter) TableName() string {
	return "catalog.collections"
}

// --- Batch Update Struct ---
type CollectionsBatchUpdate struct {
	Data       CollectionsEdit     `json:"data"`
	PathParams CollectionsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type CollectionsIdentity struct {
	Id types.URID `json:"id"`
}
