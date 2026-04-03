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
// • StorageZonesForm     - Data input validation and creation
// • StorageZones        - Main database model with relationships
// • StorageZonesEdit    - Partial update operations
// • StorageZonesIdentity - Bulk operation identifiers
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

	"github.com/maple-tech/baseline/types"
)

// StorageZonesForm handles data input validation and creation operations.
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
type StorageZonesForm struct {
	Name           string     `gorm:"column:name;not null" json:"name" example:"Nulla voluptatem."`
	WarehouseId    types.URID `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"5TRZU63KWVCSVOIE5CKYD7H6JE"`
	ZoneCode       string     `gorm:"column:zone_code" json:"zoneCode" example:"Sint ut."`
	ZoneType       string     `gorm:"column:zone_type" json:"zoneType" example:"Recusandae nisi."`
	TemperatureMin *float64   `gorm:"column:temperature_min" json:"temperatureMin,omitempty"`
	TemperatureMax *float64   `gorm:"column:temperature_max" json:"temperatureMax,omitempty"`
}

func (p *StorageZonesForm) TableName() string {
	return "logistics.storage_zones"
}

// StorageZones represents the main database model for logistics.storage_zones table.
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
// • Table: logistics.storage_zones
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
type StorageZones struct {
	Id                types.URID     `gorm:"column:id;primary_key" json:"id" example:"WIRLL4M2HBCNDI3BHTFBOHYXRA"`
	Name              string         `gorm:"column:name;not null" json:"name" example:"Nostrum quia."`
	WarehouseId       types.URID     `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"PVGXA7PQEFA4VIUFC6A6LS5QDI"`
	WarehouseIdDetail *Warehouses    `gorm:"foreignkey:WarehouseId" json:"warehouseDetail,omitempty"`
	ZoneCode          string         `gorm:"column:zone_code" json:"zoneCode" example:"Quia itaque."`
	ZoneType          string         `gorm:"column:zone_type" json:"zoneType" example:"Ut provident."`
	TemperatureMin    *float64       `gorm:"column:temperature_min" json:"temperatureMin,omitempty"`
	TemperatureMax    *float64       `gorm:"column:temperature_max" json:"temperatureMax,omitempty"`
	CreatedAt         types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	StorageBinsList   *[]StorageBins `gorm:"foreignKey:ZoneId;references:Id" json:"storageBinsList,omitempty"`
}

func (p *StorageZones) TableName() string {
	return "logistics.storage_zones"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type StorageZonesPage struct {
	paginationRuntime.DefaultPageResponse
	Items []StorageZones `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type StorageZonesEdit struct {
	Name           *string     `gorm:"column:name;not null" json:"name" example:"Praesentium dolor."`
	WarehouseId    *types.URID `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"24XLMDZ57BGB5IGXHAIBAAQVFU"`
	ZoneCode       *string     `gorm:"column:zone_code" json:"zoneCode" example:"Molestiae cupiditate."`
	ZoneType       *string     `gorm:"column:zone_type" json:"zoneType" example:"Omnis sit."`
	TemperatureMin *float64    `gorm:"column:temperature_min" json:"temperatureMin,omitempty"`
	TemperatureMax *float64    `gorm:"column:temperature_max" json:"temperatureMax,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type StorageZonesFilter struct {
	Id             *types.URID     `gorm:"column:id;primary_key" json:"id" example:"AVSVGTFDCZHPDI6D7KGYZJD6MM"`
	Name           *string         `gorm:"column:name;not null" json:"name" example:"Dicta voluptate."`
	WarehouseId    *types.URID     `gorm:"column:warehouse_id;not null" json:"warehouseId" example:"XDV23GO6MBH2JFDA7O25AY44E4"`
	ZoneCode       *string         `gorm:"column:zone_code" json:"zoneCode" example:"Rem corrupti."`
	ZoneType       *string         `gorm:"column:zone_type" json:"zoneType" example:"Aut quia."`
	TemperatureMin *float64        `gorm:"column:temperature_min" json:"temperatureMin,omitempty"`
	TemperatureMax *float64        `gorm:"column:temperature_max" json:"temperatureMax,omitempty"`
	CreatedAt      *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *StorageZonesFilter) TableName() string {
	return "logistics.storage_zones"
}

// --- Batch Update Struct ---
type StorageZonesBatchUpdate struct {
	Data       StorageZonesEdit     `json:"data"`
	PathParams StorageZonesIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type StorageZonesIdentity struct {
	Id types.URID `json:"id"`
}
