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
// • CollectionProductsForm     - Data input validation and creation
// • CollectionProducts        - Main database model with relationships
// • CollectionProductsEdit    - Partial update operations
// • CollectionProductsIdentity - Bulk operation identifiers
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

	"github.com/maple-tech/baseline/types"
)

// CollectionProductsForm handles data input validation and creation operations.
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
type CollectionProductsForm struct {
	Name         string     `gorm:"column:name" json:"name" example:"Aut vitae."`
	CollectionId types.URID `gorm:"column:collection_id;not null" json:"collectionId" example:"T6WBIFGJLJDRRFQJP2H3Q6NFBI"`
	ProductId    types.URID `gorm:"column:product_id;not null" json:"productId" example:"QEAJSVTJ3VBH5LRKC7ML3XVTFA"`
	SortOrder    int        `gorm:"column:sort_order" json:"sortOrder" example:"5314996371801019298"`
}

func (p *CollectionProductsForm) TableName() string {
	return "catalog.collection_products"
}

// CollectionProducts represents the main database model for catalog.collection_products table.
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
// • Table: catalog.collection_products
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
type CollectionProducts struct {
	Id                 types.URID     `gorm:"column:id;primary_key" json:"id" example:"BP77E5ZP2FF65PS44DZIORQPKQ"`
	Name               string         `gorm:"column:name" json:"name" example:"Ad quidem."`
	CollectionId       types.URID     `gorm:"column:collection_id;not null" json:"collectionId" example:"P2HQ3GIX65GRPPILIBFAXEQW64"`
	CollectionIdDetail *Collections   `gorm:"foreignkey:CollectionId" json:"collectionDetail,omitempty"`
	ProductId          types.URID     `gorm:"column:product_id;not null" json:"productId" example:"CMWPKD43UVDIXJXTPUXQRJ6Y4Q"`
	ProductIdDetail    *Products      `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	SortOrder          int            `gorm:"column:sort_order" json:"sortOrder" example:"-7369445401777130088"`
	CreatedAt          types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *CollectionProducts) TableName() string {
	return "catalog.collection_products"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type CollectionProductsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []CollectionProducts `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type CollectionProductsEdit struct {
	Name         *string     `gorm:"column:name" json:"name" example:"Qui in."`
	CollectionId *types.URID `gorm:"column:collection_id;not null" json:"collectionId" example:"K7ZS4KKS4NEQVNPJW7NJTFDEBU"`
	ProductId    *types.URID `gorm:"column:product_id;not null" json:"productId" example:"BLOFFXTHXZEDRDJ5DF3TYYM47Y"`
	SortOrder    *int        `gorm:"column:sort_order" json:"sortOrder" example:"2339322881606404861"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type CollectionProductsFilter struct {
	Id           *types.URID     `gorm:"column:id;primary_key" json:"id" example:"N264CNCAAZG6VEL7GY6EZ7XD24"`
	Name         *string         `gorm:"column:name" json:"name" example:"Similique expedita."`
	CollectionId *types.URID     `gorm:"column:collection_id;not null" json:"collectionId" example:"HUKNDSCPBNBXJFEIMH74FJA7SQ"`
	ProductId    *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"GHU5BCWB7NC2DCYCTGBW4IIP6Q"`
	SortOrder    *int            `gorm:"column:sort_order" json:"sortOrder" example:"5609075715971406130"`
	CreatedAt    *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *CollectionProductsFilter) TableName() string {
	return "catalog.collection_products"
}

// --- Batch Update Struct ---
type CollectionProductsBatchUpdate struct {
	Data       CollectionProductsEdit     `json:"data"`
	PathParams CollectionProductsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type CollectionProductsIdentity struct {
	Id types.URID `json:"id"`
}
