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
// • ProductTagsForm     - Data input validation and creation
// • ProductTags        - Main database model with relationships
// • ProductTagsEdit    - Partial update operations
// • ProductTagsIdentity - Bulk operation identifiers
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

// ProductTagsForm handles data input validation and creation operations.
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
type ProductTagsForm struct {
	ProductId types.URID `gorm:"column:product_id;primary_key;not null" json:"productId" example:"XDMNQT4X7VB5DCOWQTM2VRRBI4"`
	TagId     types.URID `gorm:"column:tag_id;primary_key;not null" json:"tagId" example:"BB6EMZRWY5ETXPCJZKLUTHK4ZA"`
	Name      string     `gorm:"column:name" json:"name" example:"Officiis eligendi."`
}

func (p *ProductTagsForm) TableName() string {
	return "catalog.product_tags"
}

// ProductTags represents the main database model for catalog.product_tags table.
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
// • Table: catalog.product_tags
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
type ProductTags struct {
	ProductId       types.URID     `gorm:"column:product_id;primary_key;not null" json:"productId" example:"B6NRD6RTGFE7TPIA3GCMDBVMUE"`
	ProductIdDetail *Products      `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	TagId           types.URID     `gorm:"column:tag_id;primary_key;not null" json:"tagId" example:"MGW44ARKSFGWJELUYPIHZBCWA4"`
	TagIdDetail     *Tags          `gorm:"foreignkey:TagId" json:"tagDetail,omitempty"`
	Name            string         `gorm:"column:name" json:"name" example:"Qui exercitationem."`
	CreatedAt       types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *ProductTags) TableName() string {
	return "catalog.product_tags"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ProductTagsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []ProductTags `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ProductTagsEdit struct {
	ProductId *types.URID `gorm:"column:product_id;primary_key;not null" json:"productId" example:"YVA2UTS7TNBJBOSNF5FCB3ATVQ"`
	TagId     *types.URID `gorm:"column:tag_id;primary_key;not null" json:"tagId" example:"4X2626WXIRFBTBE5ZIZOC4O4DM"`
	Name      *string     `gorm:"column:name" json:"name" example:"Autem fugiat."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ProductTagsFilter struct {
	ProductId *types.URID     `gorm:"column:product_id;primary_key;not null" json:"productId" example:"PXJEGHNFAVH3NOIDR2SQJR2OP4"`
	TagId     *types.URID     `gorm:"column:tag_id;primary_key;not null" json:"tagId" example:"XBVOEAAHU5ATBC4D4HL56FNZKU"`
	Name      *string         `gorm:"column:name" json:"name" example:"Nesciunt et."`
	CreatedAt *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *ProductTagsFilter) TableName() string {
	return "catalog.product_tags"
}

// --- Batch Update Struct ---
type ProductTagsBatchUpdate struct {
	Data       ProductTagsEdit     `json:"data"`
	PathParams ProductTagsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ProductTagsIdentity struct {
	ProductId types.URID `json:"productId"`
	TagId     types.URID `json:"tagId"`
}
