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
// • WarehousesForm     - Data input validation and creation
// • Warehouses        - Main database model with relationships
// • WarehousesEdit    - Partial update operations
// • WarehousesIdentity - Bulk operation identifiers
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

// WarehousesForm handles data input validation and creation operations.
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
type WarehousesForm struct {
	Name           string      `gorm:"column:name;not null" json:"name" example:"Non architecto."`
	Code           string      `gorm:"column:code;not null" json:"code" example:"Ut consectetur."`
	Address        types.JSON  `gorm:"column:address" json:"address"`
	OrganizationId types.URID  `gorm:"column:organization_id;not null" json:"organizationId" example:"UAIFCUXM7NABJCUUMETR72VOHM"`
	ManagerId      *types.URID `gorm:"column:manager_id" json:"managerId,omitempty" example:"RUB46ZB4OFAQ7HJHMTA7CG5BAM"`
	IsActive       bool        `gorm:"column:is_active" json:"isActive" example:"true"`
	Capacity       int         `gorm:"column:capacity" json:"capacity" example:"-2319357886384468185"`
}

func (p *WarehousesForm) TableName() string {
	return "logistics.warehouses"
}

// Warehouses represents the main database model for logistics.warehouses table.
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
// • Table: logistics.warehouses
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
type Warehouses struct {
	Id                   types.URID                     `gorm:"column:id;primary_key" json:"id" example:"JVXBHHGJKBGA3IZQVAZNHPPIZM"`
	Name                 string                         `gorm:"column:name;not null" json:"name" example:"Quaerat ad."`
	Code                 string                         `gorm:"column:code;not null" json:"code" example:"Officia omnis."`
	Address              types.JSON                     `gorm:"column:address" json:"address"`
	OrganizationId       types.URID                     `gorm:"column:organization_id;not null" json:"organizationId" example:"BCO7BE5GORFVPLSYC7MIMABUSA"`
	OrganizationIdDetail *shared_types.IamOrganizations `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	ManagerId            *types.URID                    `gorm:"column:manager_id" json:"managerId,omitempty" example:"R5ACXQ3GJJF3FOLLVDTFM23NYI"`
	ManagerIdDetail      *shared_types.IamUsers         `gorm:"foreignkey:ManagerId" json:"managerDetail,omitempty"`
	IsActive             bool                           `gorm:"column:is_active" json:"isActive" example:"true"`
	Capacity             int                            `gorm:"column:capacity" json:"capacity" example:"-6943998576667707901"`
	CreatedAt            types.NullTime                 `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime                 `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	StorageZonesList     *[]StorageZones                `gorm:"foreignKey:WarehouseId;references:Id" json:"storageZonesList,omitempty"`
	InventoryList        *[]Inventory                   `gorm:"foreignKey:WarehouseId;references:Id" json:"inventoryList,omitempty"`
	PurchaseOrdersList   *[]PurchaseOrders              `gorm:"foreignKey:WarehouseId;references:Id" json:"purchaseOrdersList,omitempty"`
	ShipmentsList        *[]Shipments                   `gorm:"foreignKey:WarehouseId;references:Id" json:"shipmentsList,omitempty"`
}

func (p *Warehouses) TableName() string {
	return "logistics.warehouses"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type WarehousesPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Warehouses `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type WarehousesEdit struct {
	Name           *string     `gorm:"column:name;not null" json:"name" example:"Ipsam perferendis."`
	Code           *string     `gorm:"column:code;not null" json:"code" example:"Eum nihil."`
	Address        *types.JSON `gorm:"column:address" json:"address"`
	OrganizationId *types.URID `gorm:"column:organization_id;not null" json:"organizationId" example:"B3ZENHFATJHNNNW63EWOVE26F4"`
	ManagerId      *types.URID `gorm:"column:manager_id" json:"managerId,omitempty" example:"C5IGPCKKYBBTPGT4BODJRSV7KQ"`
	IsActive       *bool       `gorm:"column:is_active" json:"isActive" example:"true"`
	Capacity       *int        `gorm:"column:capacity" json:"capacity" example:"3125545686245912686"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type WarehousesFilter struct {
	Id             *types.URID     `gorm:"column:id;primary_key" json:"id" example:"QRXX5D76SRFXJML6BOWZIKDONQ"`
	Name           *string         `gorm:"column:name;not null" json:"name" example:"Optio esse."`
	Code           *string         `gorm:"column:code;not null" json:"code" example:"Sequi consequatur."`
	Address        *types.JSON     `gorm:"column:address" json:"address"`
	OrganizationId *types.URID     `gorm:"column:organization_id;not null" json:"organizationId" example:"C7HWGWLVQNHZZFHPC7JN2OX4P4"`
	ManagerId      *types.URID     `gorm:"column:manager_id" json:"managerId,omitempty" example:"YPZOVBD22JF5LFVW5FT3OI7TEY"`
	IsActive       *bool           `gorm:"column:is_active" json:"isActive" example:"true"`
	Capacity       *int            `gorm:"column:capacity" json:"capacity" example:"3045249127750539697"`
	CreatedAt      *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *WarehousesFilter) TableName() string {
	return "logistics.warehouses"
}

// --- Batch Update Struct ---
type WarehousesBatchUpdate struct {
	Data       WarehousesEdit     `json:"data"`
	PathParams WarehousesIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type WarehousesIdentity struct {
	Id types.URID `json:"id"`
}
