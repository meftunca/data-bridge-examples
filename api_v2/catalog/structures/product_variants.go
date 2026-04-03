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
// • ProductVariantsForm     - Data input validation and creation
// • ProductVariants        - Main database model with relationships
// • ProductVariantsEdit    - Partial update operations
// • ProductVariantsIdentity - Bulk operation identifiers
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

// ProductVariantsForm handles data input validation and creation operations.
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
type ProductVariantsForm struct {
	Name          string     `gorm:"column:name;not null" json:"name" example:"Sint laborum."`
	ProductId     types.URID `gorm:"column:product_id;not null" json:"productId" example:"KAKNUG43JZHLPCC255NASGAYIY"`
	Sku           string     `gorm:"column:sku;not null" json:"sku" example:"Totam perspiciatis."`
	PriceOverride *float64   `gorm:"column:price_override" json:"priceOverride,omitempty"`
	Attributes    types.JSON `gorm:"column:attributes" json:"attributes"`
	IsActive      bool       `gorm:"column:is_active" json:"isActive" example:"false"`
	StockQuantity int        `gorm:"column:stock_quantity" json:"stockQuantity" example:"4251974792338410944"`
}

func (p *ProductVariantsForm) TableName() string {
	return "catalog.product_variants"
}

// ProductVariants represents the main database model for catalog.product_variants table.
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
// • Table: catalog.product_variants
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
type ProductVariants struct {
	Id              types.URID                         `gorm:"column:id;primary_key" json:"id" example:"DWANHWLW5FCPNCFXN5AL7G443A"`
	Name            string                             `gorm:"column:name;not null" json:"name" example:"Voluptatem voluptates."`
	ProductId       types.URID                         `gorm:"column:product_id;not null" json:"productId" example:"YPRQ4AZ5JRGBVFBSQAS4ECAU3E"`
	ProductIdDetail *Products                          `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	Sku             string                             `gorm:"column:sku;not null" json:"sku" example:"Tenetur et."`
	PriceOverride   *float64                           `gorm:"column:price_override" json:"priceOverride,omitempty"`
	Attributes      types.JSON                         `gorm:"column:attributes" json:"attributes"`
	IsActive        bool                               `gorm:"column:is_active" json:"isActive" example:"true"`
	StockQuantity   int                                `gorm:"column:stock_quantity" json:"stockQuantity" example:"1569207392136340129"`
	CreatedAt       types.NullTime                     `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       types.NullTime                     `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	OrderItemsList  *[]shared_types.OrdersOrderItems   `gorm:"foreignKey:VariantId;references:Id" json:"orderItemsList,omitempty"`
	CartItemsList   *[]shared_types.OrdersCartItems    `gorm:"foreignKey:VariantId;references:Id" json:"cartItemsList,omitempty"`
	InventoryList   *[]shared_types.LogisticsInventory `gorm:"foreignKey:VariantId;references:Id" json:"inventoryList,omitempty"`
}

func (p *ProductVariants) TableName() string {
	return "catalog.product_variants"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ProductVariantsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []ProductVariants `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ProductVariantsEdit struct {
	Name          *string     `gorm:"column:name;not null" json:"name" example:"Eveniet qui."`
	ProductId     *types.URID `gorm:"column:product_id;not null" json:"productId" example:"RZH42QJYR5HARMM46PB4NTZKXQ"`
	Sku           *string     `gorm:"column:sku;not null" json:"sku" example:"Repudiandae modi."`
	PriceOverride *float64    `gorm:"column:price_override" json:"priceOverride,omitempty"`
	Attributes    *types.JSON `gorm:"column:attributes" json:"attributes"`
	IsActive      *bool       `gorm:"column:is_active" json:"isActive" example:"false"`
	StockQuantity *int        `gorm:"column:stock_quantity" json:"stockQuantity" example:"2099114512460767131"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ProductVariantsFilter struct {
	Id            *types.URID     `gorm:"column:id;primary_key" json:"id" example:"WNCKD7NVDRG4RKBQFCODBJDW5A"`
	Name          *string         `gorm:"column:name;not null" json:"name" example:"Non quia."`
	ProductId     *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"JDWFSP25UVEWFFH6ZLEY2T3RXA"`
	Sku           *string         `gorm:"column:sku;not null" json:"sku" example:"Possimus consectetur."`
	PriceOverride *float64        `gorm:"column:price_override" json:"priceOverride,omitempty"`
	Attributes    *types.JSON     `gorm:"column:attributes" json:"attributes"`
	IsActive      *bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	StockQuantity *int            `gorm:"column:stock_quantity" json:"stockQuantity" example:"-6813889681419318423"`
	CreatedAt     *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt     *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ProductVariantsFilter) TableName() string {
	return "catalog.product_variants"
}

// --- Batch Update Struct ---
type ProductVariantsBatchUpdate struct {
	Data       ProductVariantsEdit     `json:"data"`
	PathParams ProductVariantsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ProductVariantsIdentity struct {
	Id types.URID `json:"id"`
}
