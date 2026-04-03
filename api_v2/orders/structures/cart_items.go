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
// • CartItemsForm     - Data input validation and creation
// • CartItems        - Main database model with relationships
// • CartItemsEdit    - Partial update operations
// • CartItemsIdentity - Bulk operation identifiers
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
	shared_types "data-bridge-examples/api_v2/shared/types"

	"github.com/maple-tech/baseline/types"
)

// CartItemsForm handles data input validation and creation operations.
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
type CartItemsForm struct {
	Name      string      `gorm:"column:name" json:"name" example:"Vitae itaque."`
	CartId    types.URID  `gorm:"column:cart_id;not null" json:"cartId" example:"WG247SU42RCJHISZ4YDJJAO7QI"`
	ProductId types.URID  `gorm:"column:product_id;not null" json:"productId" example:"U2B6UUKUB5HTJADLKZR6JAM4VA"`
	VariantId *types.URID `gorm:"column:variant_id" json:"variantId,omitempty" example:"JWRLH3ARZJG3NOCJ3SVK7OYITM"`
	Quantity  int         `gorm:"column:quantity" json:"quantity" example:"-8629194097258483787"`
}

func (p *CartItemsForm) TableName() string {
	return "orders.cart_items"
}

// CartItems represents the main database model for orders.cart_items table.
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
// • Table: orders.cart_items
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
type CartItems struct {
	Id              types.URID                           `gorm:"column:id;primary_key" json:"id" example:"4PQDXGXSR5EZBIYCNVDO3YFJII"`
	Name            string                               `gorm:"column:name" json:"name" example:"Dolorum sint."`
	CartId          types.URID                           `gorm:"column:cart_id;not null" json:"cartId" example:"2FOOD37V3JHNLANRMAUTTAOKSU"`
	CartIdDetail    *Carts                               `gorm:"foreignkey:CartId" json:"cartDetail,omitempty"`
	ProductId       types.URID                           `gorm:"column:product_id;not null" json:"productId" example:"Z2APT6AAJFCJ7JZTHYLMX5L4VU"`
	ProductIdDetail *shared_types.CatalogProducts        `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	VariantId       *types.URID                          `gorm:"column:variant_id" json:"variantId,omitempty" example:"NZ3IZTUN7NFPPFRSWIBVACD3VY"`
	VariantIdDetail *shared_types.CatalogProductVariants `gorm:"foreignkey:VariantId" json:"variantDetail,omitempty"`
	Quantity        int                                  `gorm:"column:quantity" json:"quantity" example:"-8104438715647472042"`
	CreatedAt       types.NullTime                       `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       types.NullTime                       `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *CartItems) TableName() string {
	return "orders.cart_items"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type CartItemsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []CartItems `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type CartItemsEdit struct {
	Name      *string     `gorm:"column:name" json:"name" example:"Beatae beatae."`
	CartId    *types.URID `gorm:"column:cart_id;not null" json:"cartId" example:"3KXII7ZDJZA75FA6BEEUUC2CEQ"`
	ProductId *types.URID `gorm:"column:product_id;not null" json:"productId" example:"4B2UNEHRYFEKTDM3PFNQVIYRVA"`
	VariantId *types.URID `gorm:"column:variant_id" json:"variantId,omitempty" example:"ST7PCX64FVDBTIJUIVLABZ3PZY"`
	Quantity  *int        `gorm:"column:quantity" json:"quantity" example:"1051997233511381692"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type CartItemsFilter struct {
	Id        *types.URID     `gorm:"column:id;primary_key" json:"id" example:"6IQSGKGBWFBIFBDAAAJHKOHTCM"`
	Name      *string         `gorm:"column:name" json:"name" example:"Omnis quam."`
	CartId    *types.URID     `gorm:"column:cart_id;not null" json:"cartId" example:"IGDLLYIAH5EQTANR4BK7ORJREI"`
	ProductId *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"XVEJCOLZGFCSNIO6LS7ZMZ4O7A"`
	VariantId *types.URID     `gorm:"column:variant_id" json:"variantId,omitempty" example:"7NLQAGAARJAAFDFP2OVUCIETZE"`
	Quantity  *int            `gorm:"column:quantity" json:"quantity" example:"7453097854211319625"`
	CreatedAt *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *CartItemsFilter) TableName() string {
	return "orders.cart_items"
}

// --- Batch Update Struct ---
type CartItemsBatchUpdate struct {
	Data       CartItemsEdit     `json:"data"`
	PathParams CartItemsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type CartItemsIdentity struct {
	Id types.URID `json:"id"`
}
