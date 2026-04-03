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
// • ShipmentsForm     - Data input validation and creation
// • Shipments        - Main database model with relationships
// • ShipmentsEdit    - Partial update operations
// • ShipmentsIdentity - Bulk operation identifiers
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

// ShipmentsForm handles data input validation and creation operations.
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
type ShipmentsForm struct {
	Name              string                  `gorm:"column:name" json:"name" example:"Quo quia."`
	TrackingNumber    string                  `gorm:"column:tracking_number" json:"trackingNumber" example:"Non officia."`
	OrderId           types.URID              `gorm:"column:order_id;not null" json:"orderId" example:"ND7SVN3EVRFXVH3MBZC6VMH6SM"`
	WarehouseId       types.URID              `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"VTTKKFJ7MVA7LO5LKBSPIC6GU4"`
	Carrier           string                  `gorm:"column:carrier" json:"carrier" example:"Dolor qui."`
	Status            LogisticsShipmentStatus `gorm:"column:status" json:"status"`
	WeightKg          *float64                `gorm:"column:weight_kg" json:"weightKg,omitempty"`
	ShippedAt         types.NullTime          `gorm:"column:shipped_at" json:"shippedAt,omitempty"`
	DeliveredAt       types.NullTime          `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`
	EstimatedDelivery *types.Date             `gorm:"column:estimated_delivery" json:"estimatedDelivery,omitempty"`
	ShippingCost      float64                 `gorm:"column:shipping_cost" json:"shippingCost"`
}

func (p *ShipmentsForm) TableName() string {
	return "logistics.shipments"
}

// Shipments represents the main database model for logistics.shipments table.
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
// • Table: logistics.shipments
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
type Shipments struct {
	Id                   types.URID                 `gorm:"column:id;primary_key" json:"id" example:"PTQYOPQFJ5GUVC2TYDYDGXP5DI"`
	Name                 string                     `gorm:"column:name" json:"name" example:"Et soluta."`
	TrackingNumber       string                     `gorm:"column:tracking_number" json:"trackingNumber" example:"Hic quae."`
	OrderId              types.URID                 `gorm:"column:order_id;not null" json:"orderId" example:"BBOXGCBZ6RFEHMC6MBHEN7CBLQ"`
	OrderIdDetail        *shared_types.OrdersOrders `gorm:"foreignkey:OrderId" json:"orderDetail,omitempty"`
	WarehouseId          types.URID                 `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"VHMAGMM3ZND43PKCC6JUQCERIE"`
	WarehouseIdDetail    *Warehouses                `gorm:"foreignkey:WarehouseId" json:"warehouseDetail,omitempty"`
	Carrier              string                     `gorm:"column:carrier" json:"carrier" example:"Consequuntur voluptates."`
	Status               LogisticsShipmentStatus    `gorm:"column:status" json:"status"`
	WeightKg             *float64                   `gorm:"column:weight_kg" json:"weightKg,omitempty"`
	ShippedAt            types.NullTime             `gorm:"column:shipped_at" json:"shippedAt,omitempty"`
	DeliveredAt          types.NullTime             `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`
	EstimatedDelivery    *types.Date                `gorm:"column:estimated_delivery" json:"estimatedDelivery,omitempty"`
	ShippingCost         float64                    `gorm:"column:shipping_cost" json:"shippingCost"`
	CreatedAt            types.NullTime             `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime             `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	ShipmentItemsList    *[]ShipmentItems           `gorm:"foreignKey:ShipmentId;references:Id" json:"shipmentItemsList,omitempty"`
	ShipmentTrackingList *[]ShipmentTracking        `gorm:"foreignKey:ShipmentId;references:Id" json:"shipmentTrackingList,omitempty"`
}

func (p *Shipments) TableName() string {
	return "logistics.shipments"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ShipmentsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Shipments `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ShipmentsEdit struct {
	Name              *string                  `gorm:"column:name" json:"name" example:"Repudiandae officiis."`
	TrackingNumber    *string                  `gorm:"column:tracking_number" json:"trackingNumber" example:"At eius."`
	OrderId           *types.URID              `gorm:"column:order_id;not null" json:"orderId" example:"S3NKBCMCXRCDZPTZOOEGK3WV7M"`
	WarehouseId       *types.URID              `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"DYXSVBSTENBP5MUW53JYPKEQ3U"`
	Carrier           *string                  `gorm:"column:carrier" json:"carrier" example:"Velit inventore."`
	Status            *LogisticsShipmentStatus `gorm:"column:status" json:"status"`
	WeightKg          *float64                 `gorm:"column:weight_kg" json:"weightKg,omitempty"`
	ShippedAt         *types.NullTime          `gorm:"column:shipped_at" json:"shippedAt,omitempty"`
	DeliveredAt       *types.NullTime          `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`
	EstimatedDelivery *types.Date              `gorm:"column:estimated_delivery" json:"estimatedDelivery,omitempty"`
	ShippingCost      *float64                 `gorm:"column:shipping_cost" json:"shippingCost"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ShipmentsFilter struct {
	Id                *types.URID              `gorm:"column:id;primary_key" json:"id" example:"QSUTCWPT5NAIFL7WWDOZJC34JU"`
	Name              *string                  `gorm:"column:name" json:"name" example:"Et necessitatibus."`
	TrackingNumber    *string                  `gorm:"column:tracking_number" json:"trackingNumber" example:"Est accusantium."`
	OrderId           *types.URID              `gorm:"column:order_id;not null" json:"orderId" example:"X5XERV5D6JE7TCISR4BHGIFJD4"`
	WarehouseId       *types.URID              `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"UVBZDBMDM5GVLFJDQX57322ZXU"`
	Carrier           *string                  `gorm:"column:carrier" json:"carrier" example:"Quia sapiente."`
	Status            *LogisticsShipmentStatus `gorm:"column:status" json:"status"`
	WeightKg          *float64                 `gorm:"column:weight_kg" json:"weightKg,omitempty"`
	ShippedAt         *types.NullTime          `gorm:"column:shipped_at" json:"shippedAt,omitempty"`
	DeliveredAt       *types.NullTime          `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`
	EstimatedDelivery *types.Date              `gorm:"column:estimated_delivery" json:"estimatedDelivery,omitempty"`
	ShippingCost      *float64                 `gorm:"column:shipping_cost" json:"shippingCost"`
	CreatedAt         *types.NullTime          `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         *types.NullTime          `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ShipmentsFilter) TableName() string {
	return "logistics.shipments"
}

// --- Batch Update Struct ---
type ShipmentsBatchUpdate struct {
	Data       ShipmentsEdit     `json:"data"`
	PathParams ShipmentsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ShipmentsIdentity struct {
	Id types.URID `json:"id"`
}
