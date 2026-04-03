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
// • PurchaseOrdersForm     - Data input validation and creation
// • PurchaseOrders        - Main database model with relationships
// • PurchaseOrdersEdit    - Partial update operations
// • PurchaseOrdersIdentity - Bulk operation identifiers
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

// PurchaseOrdersForm handles data input validation and creation operations.
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
type PurchaseOrdersForm struct {
	Name         string      `gorm:"column:name" json:"name" example:"Aspernatur qui."`
	PoNumber     string      `gorm:"column:po_number;not null" json:"poNumber" example:"Eius distinctio."`
	SupplierId   types.URID  `gorm:"column:supplier_id;not null" json:"supplierId" example:"OAGO2K4NXJHSPJRE6XBA6YY7RM"`
	WarehouseId  types.URID  `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"FESOA2OL2RC4JNPTOW3BA5E2OY"`
	Status       string      `gorm:"column:status" json:"status" example:"Quisquam maxime."`
	TotalAmount  float64     `gorm:"column:total_amount" json:"totalAmount"`
	ExpectedDate *types.Date `gorm:"column:expected_date" json:"expectedDate,omitempty"`
	ReceivedDate *types.Date `gorm:"column:received_date" json:"receivedDate,omitempty"`
	CreatedBy    *types.URID `gorm:"column:created_by" json:"createdBy,omitempty" example:"OXWM2R5UZNAVZMDDWSLMCLLDMU"`
	ApprovedBy   *types.URID `gorm:"column:approved_by" json:"approvedBy,omitempty" example:"D5AB36LROFGLZFJVT2AMVGHTU4"`
	Notes        string      `gorm:"column:notes" json:"notes" example:"Molestiae inventore."`
}

func (p *PurchaseOrdersForm) TableName() string {
	return "logistics.purchase_orders"
}

// PurchaseOrders represents the main database model for logistics.purchase_orders table.
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
// • Table: logistics.purchase_orders
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
type PurchaseOrders struct {
	Id                     types.URID             `gorm:"column:id;primary_key" json:"id" example:"BUMRLVJZPFC4DIKGED4LNHPECA"`
	Name                   string                 `gorm:"column:name" json:"name" example:"Non numquam."`
	PoNumber               string                 `gorm:"column:po_number;not null" json:"poNumber" example:"Ea sit."`
	SupplierId             types.URID             `gorm:"column:supplier_id;not null" json:"supplierId" example:"NRWCNEVZKZBGNAEYA5I3NITEVQ"`
	SupplierIdDetail       *Suppliers             `gorm:"foreignkey:SupplierId" json:"supplierDetail,omitempty"`
	WarehouseId            types.URID             `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"JFFMP6U22ZGMLPBIW5OHGCFY4I"`
	WarehouseIdDetail      *Warehouses            `gorm:"foreignkey:WarehouseId" json:"warehouseDetail,omitempty"`
	Status                 string                 `gorm:"column:status" json:"status" example:"Fuga molestiae."`
	TotalAmount            float64                `gorm:"column:total_amount" json:"totalAmount"`
	ExpectedDate           *types.Date            `gorm:"column:expected_date" json:"expectedDate,omitempty"`
	ReceivedDate           *types.Date            `gorm:"column:received_date" json:"receivedDate,omitempty"`
	CreatedBy              *types.URID            `gorm:"column:created_by" json:"createdBy,omitempty" example:"4OPBILS5ANHKHCPS2V6UDBZSDA"`
	CreatedByDetail        *shared_types.IamUsers `gorm:"foreignkey:CreatedBy" json:"createdByDetail,omitempty"`
	ApprovedBy             *types.URID            `gorm:"column:approved_by" json:"approvedBy,omitempty" example:"FCXX5CL66FEK7INDJAVIKVHAM4"`
	ApprovedByDetail       *shared_types.IamUsers `gorm:"foreignkey:ApprovedBy" json:"approvedByDetail,omitempty"`
	Notes                  string                 `gorm:"column:notes" json:"notes" example:"Ad vel."`
	CreatedAt              types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt              types.NullTime         `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	PurchaseOrderItemsList *[]PurchaseOrderItems  `gorm:"foreignKey:PurchaseOrderId;references:Id" json:"purchaseOrderItemsList,omitempty"`
}

func (p *PurchaseOrders) TableName() string {
	return "logistics.purchase_orders"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type PurchaseOrdersPage struct {
	paginationRuntime.DefaultPageResponse
	Items []PurchaseOrders `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type PurchaseOrdersEdit struct {
	Name         *string     `gorm:"column:name" json:"name" example:"Non alias."`
	PoNumber     *string     `gorm:"column:po_number;not null" json:"poNumber" example:"Saepe nihil."`
	SupplierId   *types.URID `gorm:"column:supplier_id;not null" json:"supplierId" example:"B7Q5L7QGPVEH3OANIV3VEQLIIY"`
	WarehouseId  *types.URID `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"PUU2RD46N5F2VKBUZVH4SGJ3KY"`
	Status       *string     `gorm:"column:status" json:"status" example:"Deleniti eum."`
	TotalAmount  *float64    `gorm:"column:total_amount" json:"totalAmount"`
	ExpectedDate *types.Date `gorm:"column:expected_date" json:"expectedDate,omitempty"`
	ReceivedDate *types.Date `gorm:"column:received_date" json:"receivedDate,omitempty"`
	CreatedBy    *types.URID `gorm:"column:created_by" json:"createdBy,omitempty" example:"GFMHY3XNT5EXLHHN2RBCUCJOEA"`
	ApprovedBy   *types.URID `gorm:"column:approved_by" json:"approvedBy,omitempty" example:"LA23AZBLPFGFJLF3ORI7GDXQQA"`
	Notes        *string     `gorm:"column:notes" json:"notes" example:"Voluptatem perferendis."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type PurchaseOrdersFilter struct {
	Id           *types.URID     `gorm:"column:id;primary_key" json:"id" example:"RER4UXMVTJECXOJLA6WQOPBPQI"`
	Name         *string         `gorm:"column:name" json:"name" example:"Aut sit."`
	PoNumber     *string         `gorm:"column:po_number;not null" json:"poNumber" example:"Magnam porro."`
	SupplierId   *types.URID     `gorm:"column:supplier_id;not null" json:"supplierId" example:"MSPTRLBAGVCELFV22QEBWCD6BU"`
	WarehouseId  *types.URID     `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"KKDGCLXOKBF7BCYCUBQWEOMLZI"`
	Status       *string         `gorm:"column:status" json:"status" example:"Dolor ea."`
	TotalAmount  *float64        `gorm:"column:total_amount" json:"totalAmount"`
	ExpectedDate *types.Date     `gorm:"column:expected_date" json:"expectedDate,omitempty"`
	ReceivedDate *types.Date     `gorm:"column:received_date" json:"receivedDate,omitempty"`
	CreatedBy    *types.URID     `gorm:"column:created_by" json:"createdBy,omitempty" example:"OJPRMYUHEFG57LHILWBCBLIUPE"`
	ApprovedBy   *types.URID     `gorm:"column:approved_by" json:"approvedBy,omitempty" example:"HC4GBJS2AVEK3AFUDDOPU7KX2U"`
	Notes        *string         `gorm:"column:notes" json:"notes" example:"Quo et."`
	CreatedAt    *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt    *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *PurchaseOrdersFilter) TableName() string {
	return "logistics.purchase_orders"
}

// --- Batch Update Struct ---
type PurchaseOrdersBatchUpdate struct {
	Data       PurchaseOrdersEdit     `json:"data"`
	PathParams PurchaseOrdersIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type PurchaseOrdersIdentity struct {
	Id types.URID `json:"id"`
}
