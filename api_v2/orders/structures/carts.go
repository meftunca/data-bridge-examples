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
// • CartsForm     - Data input validation and creation
// • Carts        - Main database model with relationships
// • CartsEdit    - Partial update operations
// • CartsIdentity - Bulk operation identifiers
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
package orders_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"

	"github.com/maple-tech/baseline/types"
)

// CartsForm handles data input validation and creation operations.
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
type CartsForm struct {
	Name       string         `gorm:"column:name" json:"name" example:"Qui aliquam."`
	CustomerId types.URID     `gorm:"column:customer_id;not null" json:"customerId" example:"TMURL6BH2NHHTNDTPTOIDGSYSY"`
	IsActive   bool           `gorm:"column:is_active" json:"isActive" example:"true"`
	ExpiresAt  types.NullTime `gorm:"column:expires_at" json:"expiresAt,omitempty"`
}

func (p *CartsForm) TableName() string {
	return "orders.carts"
}

// Carts represents the main database model for orders.carts table.
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
// • Table: orders.carts
// • Type: BASE TABLE
// • Schema: orders
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
type Carts struct {
	Id               types.URID     `gorm:"column:id;primary_key" json:"id" example:"EHRQZAAAZZE4VHCXYY7TGZSO3U"`
	Name             string         `gorm:"column:name" json:"name" example:"Atque fuga."`
	CustomerId       types.URID     `gorm:"column:customer_id;not null" json:"customerId" example:"UZTSXZGQ5ZFPLLQPOQCJ72Z7WE"`
	CustomerIdDetail *Customers     `gorm:"foreignkey:CustomerId" json:"customerDetail,omitempty"`
	IsActive         bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	ExpiresAt        types.NullTime `gorm:"column:expires_at" json:"expiresAt,omitempty"`
	CreatedAt        types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	CartItemsList    *[]CartItems   `gorm:"foreignKey:CartId;references:Id" json:"cartItemsList,omitempty"`
}

func (p *Carts) TableName() string {
	return "orders.carts"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type CartsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Carts `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type CartsEdit struct {
	Name       *string         `gorm:"column:name" json:"name" example:"Vero autem."`
	CustomerId *types.URID     `gorm:"column:customer_id;not null" json:"customerId" example:"ARECSK5MXVFMTBUMI645NCWLR4"`
	IsActive   *bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	ExpiresAt  *types.NullTime `gorm:"column:expires_at" json:"expiresAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type CartsFilter struct {
	Id         *types.URID     `gorm:"column:id;primary_key" json:"id" example:"2UECW3FJPNBCDBXP7O7V7KPESY"`
	Name       *string         `gorm:"column:name" json:"name" example:"Dolor in."`
	CustomerId *types.URID     `gorm:"column:customer_id;not null" json:"customerId" example:"JLZGOF7ATFFHPIU7TRZ37TFWEM"`
	IsActive   *bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	ExpiresAt  *types.NullTime `gorm:"column:expires_at" json:"expiresAt,omitempty"`
	CreatedAt  *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt  *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *CartsFilter) TableName() string {
	return "orders.carts"
}

// --- Batch Update Struct ---
type CartsBatchUpdate struct {
	Data       CartsEdit     `json:"data"`
	PathParams CartsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type CartsIdentity struct {
	Id types.URID `json:"id"`
}
