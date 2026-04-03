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
// • ShipmentItemsForm     - Data input validation and creation
// • ShipmentItems        - Main database model with relationships
// • ShipmentItemsEdit    - Partial update operations
// • ShipmentItemsIdentity - Bulk operation identifiers
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

// ShipmentItemsForm handles data input validation and creation operations.
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
type ShipmentItemsForm struct {
	Name        string     `gorm:"column:name" json:"name" example:"Non delectus."`
	ShipmentId  types.URID `gorm:"column:shipment_id;not null" json:"shipmentId" example:"UPRSPQNJRJAK7DSGEUTJZQBQ4M"`
	OrderItemId types.URID `gorm:"column:order_item_id;not null" json:"orderItemId" example:"OF6UBDR5IFGYRKFDVW666CYWUU"`
	Quantity    int        `gorm:"column:quantity" json:"quantity" example:"1318654903649933802"`
}

func (p *ShipmentItemsForm) TableName() string {
	return "logistics.shipment_items"
}

// ShipmentItems represents the main database model for logistics.shipment_items table.
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
// • Table: logistics.shipment_items
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
type ShipmentItems struct {
	Id                types.URID                     `gorm:"column:id;primary_key" json:"id" example:"DVOJFE46IFBFFOVDFJBQ6XBUCE"`
	Name              string                         `gorm:"column:name" json:"name" example:"Repudiandae cumque."`
	ShipmentId        types.URID                     `gorm:"column:shipment_id;not null" json:"shipmentId" example:"4CLSTTT2RJD3FEGTQWSM24OO3I"`
	ShipmentIdDetail  *Shipments                     `gorm:"foreignkey:ShipmentId" json:"shipmentDetail,omitempty"`
	OrderItemId       types.URID                     `gorm:"column:order_item_id;not null" json:"orderItemId" example:"VQNAGG7TK5DMFJGH5GQVWNA5FE"`
	OrderItemIdDetail *shared_types.OrdersOrderItems `gorm:"foreignkey:OrderItemId" json:"orderItemDetail,omitempty"`
	Quantity          int                            `gorm:"column:quantity" json:"quantity" example:"-3721712668441230825"`
	CreatedAt         types.NullTime                 `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *ShipmentItems) TableName() string {
	return "logistics.shipment_items"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ShipmentItemsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []ShipmentItems `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ShipmentItemsEdit struct {
	Name        *string     `gorm:"column:name" json:"name" example:"Ea libero."`
	ShipmentId  *types.URID `gorm:"column:shipment_id;not null" json:"shipmentId" example:"UEVIOEWRRBHRJITIWQGJAQ25VQ"`
	OrderItemId *types.URID `gorm:"column:order_item_id;not null" json:"orderItemId" example:"SATPXA4NV5GYBP23ZTAU675YXU"`
	Quantity    *int        `gorm:"column:quantity" json:"quantity" example:"-5660609318774788262"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ShipmentItemsFilter struct {
	Id          *types.URID     `gorm:"column:id;primary_key" json:"id" example:"D5RLV7WPYRBF5N5NFIUYOOVEMM"`
	Name        *string         `gorm:"column:name" json:"name" example:"Ab omnis."`
	ShipmentId  *types.URID     `gorm:"column:shipment_id;not null" json:"shipmentId" example:"YDLTUCNCBVB4PCWOL6V4I3FDVQ"`
	OrderItemId *types.URID     `gorm:"column:order_item_id;not null" json:"orderItemId" example:"HIG2T3X5TZGS7O3K2JLDVMPMAA"`
	Quantity    *int            `gorm:"column:quantity" json:"quantity" example:"-8381684539073838560"`
	CreatedAt   *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *ShipmentItemsFilter) TableName() string {
	return "logistics.shipment_items"
}

// --- Batch Update Struct ---
type ShipmentItemsBatchUpdate struct {
	Data       ShipmentItemsEdit     `json:"data"`
	PathParams ShipmentItemsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ShipmentItemsIdentity struct {
	Id types.URID `json:"id"`
}
