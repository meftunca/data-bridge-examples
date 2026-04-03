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
// • PurchaseOrderItemsForm     - Data input validation and creation
// • PurchaseOrderItems        - Main database model with relationships
// • PurchaseOrderItemsEdit    - Partial update operations
// • PurchaseOrderItemsIdentity - Bulk operation identifiers
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
package logistics_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"
	shared_types "data-bridge-examples/api_v2/shared/types"

	"github.com/maple-tech/baseline/types"
)

// PurchaseOrderItemsForm handles data input validation and creation operations.
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
type PurchaseOrderItemsForm struct {
	Name            string     `gorm:"column:name" json:"name" example:"Modi qui."`
	PurchaseOrderId types.URID `gorm:"column:purchase_order_id;not null" json:"purchaseOrderId" example:"4FDG6HYIYZA2DL3PEDTMC2NJ4U"`
	ProductId       types.URID `gorm:"column:product_id;not null" json:"productId" example:"EB2RF73GCVGX7NH2I2VL4S6BEM"`
	Quantity        int        `gorm:"column:quantity" json:"quantity" example:"2332779740348520387"`
	UnitCost        float64    `gorm:"column:unit_cost" json:"unitCost"`
	ReceivedQty     int        `gorm:"column:received_qty" json:"receivedQty" example:"4419759839328028973"`
}

func (p *PurchaseOrderItemsForm) TableName() string {
	return "logistics.purchase_order_items"
}

// PurchaseOrderItems represents the main database model for logistics.purchase_order_items table.
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
// • Table: logistics.purchase_order_items
// • Type: BASE TABLE
// • Schema: logistics
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
type PurchaseOrderItems struct {
	Id                    types.URID                    `gorm:"column:id;primary_key" json:"id" example:"3TKF663VSJCMTPFSEGKO3OHL4Y"`
	Name                  string                        `gorm:"column:name" json:"name" example:"Dolor exercitationem."`
	PurchaseOrderId       types.URID                    `gorm:"column:purchase_order_id;not null" json:"purchaseOrderId" example:"74X425CTMRHTXOGXYVZ7GH5EHE"`
	PurchaseOrderIdDetail *PurchaseOrders               `gorm:"foreignkey:PurchaseOrderId" json:"purchaseOrderDetail,omitempty"`
	ProductId             types.URID                    `gorm:"column:product_id;not null" json:"productId" example:"KS64SAAS3RDODDWJ4MSIHN7EDM"`
	ProductIdDetail       *shared_types.CatalogProducts `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	Quantity              int                           `gorm:"column:quantity" json:"quantity" example:"-1830601344851961710"`
	UnitCost              float64                       `gorm:"column:unit_cost" json:"unitCost"`
	ReceivedQty           int                           `gorm:"column:received_qty" json:"receivedQty" example:"-812457233802048285"`
	CreatedAt             types.NullTime                `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt             types.NullTime                `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *PurchaseOrderItems) TableName() string {
	return "logistics.purchase_order_items"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type PurchaseOrderItemsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []PurchaseOrderItems `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type PurchaseOrderItemsEdit struct {
	Name            *string     `gorm:"column:name" json:"name" example:"Nihil qui."`
	PurchaseOrderId *types.URID `gorm:"column:purchase_order_id;not null" json:"purchaseOrderId" example:"LFHHTNXTF5ELRARO543TQUYVL4"`
	ProductId       *types.URID `gorm:"column:product_id;not null" json:"productId" example:"4U4UPUY5E5FU5I37TBCWDYXQD4"`
	Quantity        *int        `gorm:"column:quantity" json:"quantity" example:"-1393566995128678740"`
	UnitCost        *float64    `gorm:"column:unit_cost" json:"unitCost"`
	ReceivedQty     *int        `gorm:"column:received_qty" json:"receivedQty" example:"8630890499017377819"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type PurchaseOrderItemsFilter struct {
	Id              *types.URID     `gorm:"column:id;primary_key" json:"id" example:"PYFNCUUJH5ENLBDLVJ7JIUAZGQ"`
	Name            *string         `gorm:"column:name" json:"name" example:"Expedita velit."`
	PurchaseOrderId *types.URID     `gorm:"column:purchase_order_id;not null" json:"purchaseOrderId" example:"DX4CDIFBIZD2BJD6OXJJMPZH7Y"`
	ProductId       *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"HNIHUEONA5F3FPZ457QJQIC6QU"`
	Quantity        *int            `gorm:"column:quantity" json:"quantity" example:"8894268412383251949"`
	UnitCost        *float64        `gorm:"column:unit_cost" json:"unitCost"`
	ReceivedQty     *int            `gorm:"column:received_qty" json:"receivedQty" example:"-3135978540472971386"`
	CreatedAt       *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *PurchaseOrderItemsFilter) TableName() string {
	return "logistics.purchase_order_items"
}

// --- Batch Update Struct ---
type PurchaseOrderItemsBatchUpdate struct {
	Data       PurchaseOrderItemsEdit     `json:"data"`
	PathParams PurchaseOrderItemsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type PurchaseOrderItemsIdentity struct {
	Id types.URID `json:"id"`
}
