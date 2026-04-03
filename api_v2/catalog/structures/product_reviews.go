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
// • ProductReviewsForm     - Data input validation and creation
// • ProductReviews        - Main database model with relationships
// • ProductReviewsEdit    - Partial update operations
// • ProductReviewsIdentity - Bulk operation identifiers
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

// ProductReviewsForm handles data input validation and creation operations.
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
type ProductReviewsForm struct {
	Name         string     `gorm:"column:name" json:"name" example:"Vero dolor."`
	ProductId    types.URID `gorm:"column:product_id;not null" json:"productId" example:"6K6GA2NYTNHKNFJBRN7WNEQPLY"`
	UserId       types.URID `gorm:"column:user_id;not null" json:"userId" example:"4IDTZPI4YZAGFGDQBAJWUEAWZA"`
	Rating       int16      `gorm:"column:rating;not null" json:"rating" example:"-2798999126142893712"`
	Title        string     `gorm:"column:title" json:"title" example:"Voluptas dolor."`
	Body         string     `gorm:"column:body" json:"body" example:"Non pariatur."`
	IsVerified   bool       `gorm:"column:is_verified" json:"isVerified" example:"false"`
	HelpfulCount int        `gorm:"column:helpful_count" json:"helpfulCount" example:"-6706273900186366790"`
}

func (p *ProductReviewsForm) TableName() string {
	return "catalog.product_reviews"
}

// ProductReviews represents the main database model for catalog.product_reviews table.
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
// • Table: catalog.product_reviews
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
type ProductReviews struct {
	Id              types.URID             `gorm:"column:id;primary_key" json:"id" example:"G7YPCUH4MNE5XG4V4GMTXJTAIE"`
	Name            string                 `gorm:"column:name" json:"name" example:"Cumque et."`
	ProductId       types.URID             `gorm:"column:product_id;not null" json:"productId" example:"WSAF3MIC7BDQFERPXU3BZFVQMQ"`
	ProductIdDetail *Products              `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	UserId          types.URID             `gorm:"column:user_id;not null" json:"userId" example:"UGEFUROCEBHZLPQVTVADSZNO6E"`
	UserIdDetail    *shared_types.IamUsers `gorm:"foreignkey:UserId" json:"userDetail,omitempty"`
	Rating          int16                  `gorm:"column:rating;not null" json:"rating" example:"-8736733771519174919"`
	Title           string                 `gorm:"column:title" json:"title" example:"Aspernatur adipisci."`
	Body            string                 `gorm:"column:body" json:"body" example:"Ex illum."`
	IsVerified      bool                   `gorm:"column:is_verified" json:"isVerified" example:"false"`
	HelpfulCount    int                    `gorm:"column:helpful_count" json:"helpfulCount" example:"4284179964300972655"`
	CreatedAt       types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       types.NullTime         `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ProductReviews) TableName() string {
	return "catalog.product_reviews"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ProductReviewsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []ProductReviews `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ProductReviewsEdit struct {
	Name         *string     `gorm:"column:name" json:"name" example:"Numquam exercitationem."`
	ProductId    *types.URID `gorm:"column:product_id;not null" json:"productId" example:"FGAK3J5L7BCADDB4GIJM63U67Y"`
	UserId       *types.URID `gorm:"column:user_id;not null" json:"userId" example:"SITD2A2ON5FGRIZIM7N7RRCEOQ"`
	Rating       *int16      `gorm:"column:rating;not null" json:"rating" example:"-3922599630242125427"`
	Title        *string     `gorm:"column:title" json:"title" example:"Quaerat quaerat."`
	Body         *string     `gorm:"column:body" json:"body" example:"Ipsam laudantium."`
	IsVerified   *bool       `gorm:"column:is_verified" json:"isVerified" example:"false"`
	HelpfulCount *int        `gorm:"column:helpful_count" json:"helpfulCount" example:"-1037653363090840796"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ProductReviewsFilter struct {
	Id           *types.URID     `gorm:"column:id;primary_key" json:"id" example:"XCYGWFBKVZEQTPSI6YZCBE3VAY"`
	Name         *string         `gorm:"column:name" json:"name" example:"Cum quis."`
	ProductId    *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"5CMWUSUQSJHTBHU6E3O2IPUZ4Q"`
	UserId       *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"XLBI7EQWOZC7TLBKEJPZDDBKHE"`
	Rating       *int16          `gorm:"column:rating;not null" json:"rating" example:"533883341292978941"`
	Title        *string         `gorm:"column:title" json:"title" example:"Atque dolorum."`
	Body         *string         `gorm:"column:body" json:"body" example:"Aut eveniet."`
	IsVerified   *bool           `gorm:"column:is_verified" json:"isVerified" example:"true"`
	HelpfulCount *int            `gorm:"column:helpful_count" json:"helpfulCount" example:"-1790625025117786699"`
	CreatedAt    *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt    *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ProductReviewsFilter) TableName() string {
	return "catalog.product_reviews"
}

// --- Batch Update Struct ---
type ProductReviewsBatchUpdate struct {
	Data       ProductReviewsEdit     `json:"data"`
	PathParams ProductReviewsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ProductReviewsIdentity struct {
	Id types.URID `json:"id"`
}
