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
// • OrderItemsForm     - Data input validation and creation
// • OrderItems        - Main database model with relationships
// • OrderItemsEdit    - Partial update operations
// • OrderItemsIdentity - Bulk operation identifiers
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

// OrderItemsForm handles data input validation and creation operations.
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
type OrderItemsForm struct {
	Name       string      `gorm:"column:name" json:"name" example:"Omnis doloremque."`
	OrderId    types.URID  `gorm:"column:order_id;not null" json:"orderId" example:"OTWA3WYQDVDLBOVHMUGCJI6Q7I"`
	ProductId  types.URID  `gorm:"column:product_id;not null" json:"productId" example:"Q4O4SAZTPJH4TJCXC3GIDT3VCU"`
	VariantId  *types.URID `gorm:"column:variant_id" json:"variantId,omitempty" example:"UO3PWJ733VFCLLKBRUIKCCI6I4"`
	Quantity   int         `gorm:"column:quantity" json:"quantity" example:"7897463541490618937"`
	UnitPrice  float64     `gorm:"column:unit_price" json:"unitPrice"`
	TotalPrice float64     `gorm:"column:total_price" json:"totalPrice"`
	Discount   float64     `gorm:"column:discount" json:"discount"`
	Metadata   types.JSON  `gorm:"column:metadata" json:"metadata"`
}

func (p *OrderItemsForm) TableName() string {
	return "orders.order_items"
}

// OrderItems represents the main database model for orders.order_items table.
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
// • Table: orders.order_items
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
type OrderItems struct {
	Id                types.URID                             `gorm:"column:id;primary_key" json:"id" example:"7YJLLE6OBRFTPKH4CIFLJDSVOI"`
	Name              string                                 `gorm:"column:name" json:"name" example:"Corrupti deleniti."`
	OrderId           types.URID                             `gorm:"column:order_id;not null" json:"orderId" example:"SGC6DKTDYZFQPNPOP37TCR65KI"`
	OrderIdDetail     *Orders                                `gorm:"foreignkey:OrderId" json:"orderDetail,omitempty"`
	ProductId         types.URID                             `gorm:"column:product_id;not null" json:"productId" example:"5PH6CSZWKNB5PCDZJWSVBZJT7M"`
	ProductIdDetail   *shared_types.CatalogProducts          `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	VariantId         *types.URID                            `gorm:"column:variant_id" json:"variantId,omitempty" example:"Z43E2NQ3FFC5BC7NXRQSL4Y6MU"`
	VariantIdDetail   *shared_types.CatalogProductVariants   `gorm:"foreignkey:VariantId" json:"variantDetail,omitempty"`
	Quantity          int                                    `gorm:"column:quantity" json:"quantity" example:"3427207634078680259"`
	UnitPrice         float64                                `gorm:"column:unit_price" json:"unitPrice"`
	TotalPrice        float64                                `gorm:"column:total_price" json:"totalPrice"`
	Discount          float64                                `gorm:"column:discount" json:"discount"`
	Metadata          types.JSON                             `gorm:"column:metadata" json:"metadata"`
	CreatedAt         types.NullTime                         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         types.NullTime                         `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	ShipmentItemsList *[]shared_types.LogisticsShipmentItems `gorm:"foreignKey:OrderItemId;references:Id" json:"shipmentItemsList,omitempty"`
}

func (p *OrderItems) TableName() string {
	return "orders.order_items"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type OrderItemsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []OrderItems `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type OrderItemsEdit struct {
	Name       *string     `gorm:"column:name" json:"name" example:"Nulla et."`
	OrderId    *types.URID `gorm:"column:order_id;not null" json:"orderId" example:"YHBY76MF5JFDPGX5CWAFYDVS5A"`
	ProductId  *types.URID `gorm:"column:product_id;not null" json:"productId" example:"ZBDU5GVSWRGVNB2SXP5O26BBNM"`
	VariantId  *types.URID `gorm:"column:variant_id" json:"variantId,omitempty" example:"EMLXAYMG2FA27D2GQE4MT4PMYI"`
	Quantity   *int        `gorm:"column:quantity" json:"quantity" example:"-2534001963477605030"`
	UnitPrice  *float64    `gorm:"column:unit_price" json:"unitPrice"`
	TotalPrice *float64    `gorm:"column:total_price" json:"totalPrice"`
	Discount   *float64    `gorm:"column:discount" json:"discount"`
	Metadata   *types.JSON `gorm:"column:metadata" json:"metadata"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type OrderItemsFilter struct {
	Id         *types.URID     `gorm:"column:id;primary_key" json:"id" example:"FNTKK5JAPRARTFZAK2EVFGR4HY"`
	Name       *string         `gorm:"column:name" json:"name" example:"Et minus."`
	OrderId    *types.URID     `gorm:"column:order_id;not null" json:"orderId" example:"V4CCV2COEJHHNGX4G2TTFQF7ZM"`
	ProductId  *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"SF4UVU3UMZHGBCYFZ47I2J6LNY"`
	VariantId  *types.URID     `gorm:"column:variant_id" json:"variantId,omitempty" example:"7US7GZO4YFEBLGARQJ5TYUR2O4"`
	Quantity   *int            `gorm:"column:quantity" json:"quantity" example:"4012108276520261309"`
	UnitPrice  *float64        `gorm:"column:unit_price" json:"unitPrice"`
	TotalPrice *float64        `gorm:"column:total_price" json:"totalPrice"`
	Discount   *float64        `gorm:"column:discount" json:"discount"`
	Metadata   *types.JSON     `gorm:"column:metadata" json:"metadata"`
	CreatedAt  *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt  *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *OrderItemsFilter) TableName() string {
	return "orders.order_items"
}

// --- Batch Update Struct ---
type OrderItemsBatchUpdate struct {
	Data       OrderItemsEdit     `json:"data"`
	PathParams OrderItemsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type OrderItemsIdentity struct {
	Id types.URID `json:"id"`
}
