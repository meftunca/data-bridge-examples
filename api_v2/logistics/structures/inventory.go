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
// • InventoryForm     - Data input validation and creation
// • Inventory        - Main database model with relationships
// • InventoryEdit    - Partial update operations
// • InventoryIdentity - Bulk operation identifiers
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

// InventoryForm handles data input validation and creation operations.
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
type InventoryForm struct {
	Name            string         `gorm:"column:name" json:"name" example:"Unde cupiditate."`
	ProductId       types.URID     `gorm:"column:product_id;not null" json:"productId" example:"JACJAVTWMRHYPLFZTV4LC7S3QQ"`
	VariantId       *types.URID    `gorm:"column:variant_id" json:"variantId,omitempty" example:"S4PIPWESM5CATCEEPBHKWBYCOM"`
	WarehouseId     types.URID     `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"SXUBQKQCTVCI5BH42FNXODHGUA"`
	BinId           *types.URID    `gorm:"column:bin_id" json:"binId,omitempty" example:"SET7RH6CHJB27I3C2ITLCDG6TE"`
	Quantity        int            `gorm:"column:quantity" json:"quantity" example:"-3762619710832071451"`
	Reserved        int            `gorm:"column:reserved" json:"reserved" example:"6518810292333631750"`
	ReorderLevel    int            `gorm:"column:reorder_level" json:"reorderLevel" example:"237072413050753137"`
	ReorderQuantity int            `gorm:"column:reorder_quantity" json:"reorderQuantity" example:"6582485565975305081"`
	LastCountedAt   types.NullTime `gorm:"column:last_counted_at" json:"lastCountedAt,omitempty"`
}

func (p *InventoryForm) TableName() string {
	return "logistics.inventory"
}

// Inventory represents the main database model for logistics.inventory table.
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
// • Table: logistics.inventory
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
type Inventory struct {
	Id                 types.URID                           `gorm:"column:id;primary_key" json:"id" example:"LEVM45TTMRG3PPKINNFJQLBFQM"`
	Name               string                               `gorm:"column:name" json:"name" example:"Voluptas reiciendis."`
	ProductId          types.URID                           `gorm:"column:product_id;not null" json:"productId" example:"YGYN7FYACRET3HQJIMKSKLEIBQ"`
	ProductIdDetail    *shared_types.CatalogProducts        `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	VariantId          *types.URID                          `gorm:"column:variant_id" json:"variantId,omitempty" example:"GW57QGCGXZAUFCGXUBD7YREMMA"`
	VariantIdDetail    *shared_types.CatalogProductVariants `gorm:"foreignkey:VariantId" json:"variantDetail,omitempty"`
	WarehouseId        types.URID                           `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"CCBL46UFJJFW7FBNXEYC652SXM"`
	WarehouseIdDetail  *Warehouses                          `gorm:"foreignkey:WarehouseId" json:"warehouseDetail,omitempty"`
	BinId              *types.URID                          `gorm:"column:bin_id" json:"binId,omitempty" example:"6E5Y6KCXMVES7CKEY5AR4LYE3U"`
	BinIdDetail        *StorageBins                         `gorm:"foreignkey:BinId" json:"binDetail,omitempty"`
	Quantity           int                                  `gorm:"column:quantity" json:"quantity" example:"-5836463249071313332"`
	Reserved           int                                  `gorm:"column:reserved" json:"reserved" example:"-4677192765293119322"`
	ReorderLevel       int                                  `gorm:"column:reorder_level" json:"reorderLevel" example:"-8601164409900196864"`
	ReorderQuantity    int                                  `gorm:"column:reorder_quantity" json:"reorderQuantity" example:"2603447199116609237"`
	LastCountedAt      types.NullTime                       `gorm:"column:last_counted_at" json:"lastCountedAt,omitempty"`
	CreatedAt          types.NullTime                       `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt          types.NullTime                       `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	StockMovementsList *[]StockMovements                    `gorm:"foreignKey:InventoryId;references:Id" json:"stockMovementsList,omitempty"`
}

func (p *Inventory) TableName() string {
	return "logistics.inventory"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type InventoryPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Inventory `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type InventoryEdit struct {
	Name            *string         `gorm:"column:name" json:"name" example:"Iure eius."`
	ProductId       *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"FVEFDHLRUZG2VCPCDUWWWS5DNA"`
	VariantId       *types.URID     `gorm:"column:variant_id" json:"variantId,omitempty" example:"X5O6Z5VYHVG7VPKNYCSKBXRSUQ"`
	WarehouseId     *types.URID     `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"KUI7V5TPZNDVFIFW5SX5WPGBVU"`
	BinId           *types.URID     `gorm:"column:bin_id" json:"binId,omitempty" example:"VEJGA4HOQFAUVIY7ZZEMIOIOZA"`
	Quantity        *int            `gorm:"column:quantity" json:"quantity" example:"2005962869672212292"`
	Reserved        *int            `gorm:"column:reserved" json:"reserved" example:"-93434794599045479"`
	ReorderLevel    *int            `gorm:"column:reorder_level" json:"reorderLevel" example:"-3701361372202100695"`
	ReorderQuantity *int            `gorm:"column:reorder_quantity" json:"reorderQuantity" example:"4115957765108922253"`
	LastCountedAt   *types.NullTime `gorm:"column:last_counted_at" json:"lastCountedAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type InventoryFilter struct {
	Id              *types.URID     `gorm:"column:id;primary_key" json:"id" example:"B5PTHOCDQ5FQ3JQXN5FO5TZMBU"`
	Name            *string         `gorm:"column:name" json:"name" example:"Magnam unde."`
	ProductId       *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"WZY2TF7H7JHQTKHLTMIXKYFHTM"`
	VariantId       *types.URID     `gorm:"column:variant_id" json:"variantId,omitempty" example:"B6HG7EJ4VVGHRHTHELYRY2LVXI"`
	WarehouseId     *types.URID     `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"M6FVQZOHKBC2TFSGZJK5WVLOBQ"`
	BinId           *types.URID     `gorm:"column:bin_id" json:"binId,omitempty" example:"BCSXEI44UBFI3G6CTLJVAFTPDU"`
	Quantity        *int            `gorm:"column:quantity" json:"quantity" example:"-2957348061269271086"`
	Reserved        *int            `gorm:"column:reserved" json:"reserved" example:"-2446585716923887196"`
	ReorderLevel    *int            `gorm:"column:reorder_level" json:"reorderLevel" example:"-4638513742711508047"`
	ReorderQuantity *int            `gorm:"column:reorder_quantity" json:"reorderQuantity" example:"6555744133248877369"`
	LastCountedAt   *types.NullTime `gorm:"column:last_counted_at" json:"lastCountedAt,omitempty"`
	CreatedAt       *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *InventoryFilter) TableName() string {
	return "logistics.inventory"
}

// --- Batch Update Struct ---
type InventoryBatchUpdate struct {
	Data       InventoryEdit     `json:"data"`
	PathParams InventoryIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type InventoryIdentity struct {
	Id types.URID `json:"id"`
}
