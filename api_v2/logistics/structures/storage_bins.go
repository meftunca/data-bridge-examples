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
// • StorageBinsForm     - Data input validation and creation
// • StorageBins        - Main database model with relationships
// • StorageBinsEdit    - Partial update operations
// • StorageBinsIdentity - Bulk operation identifiers
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

// StorageBinsForm handles data input validation and creation operations.
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
type StorageBinsForm struct {
	Name         string     `gorm:"column:name;not null" json:"name" example:"Qui eum."`
	ZoneId       types.URID `gorm:"column:zone_id;not null" json:"zoneId" example:"2O737TQZ5VGT3EFVTFIOE2KXPQ"`
	BinCode      string     `gorm:"column:bin_code" json:"binCode" example:"Totam distinctio."`
	MaxCapacity  int        `gorm:"column:max_capacity" json:"maxCapacity" example:"4414053521326819654"`
	CurrentCount int        `gorm:"column:current_count" json:"currentCount" example:"-7664039807535286808"`
}

func (p *StorageBinsForm) TableName() string {
	return "logistics.storage_bins"
}

// StorageBins represents the main database model for logistics.storage_bins table.
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
// • Table: logistics.storage_bins
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
type StorageBins struct {
	Id            types.URID     `gorm:"column:id;primary_key" json:"id" example:"OGYKZFI4NFHNZF7SQZIMOVWZH4"`
	Name          string         `gorm:"column:name;not null" json:"name" example:"Qui expedita."`
	ZoneId        types.URID     `gorm:"column:zone_id;not null" json:"zoneId" example:"E7J4PMKSSVBNNN37HGHHLQTNAM"`
	ZoneIdDetail  *StorageZones  `gorm:"foreignkey:ZoneId" json:"zoneDetail,omitempty"`
	BinCode       string         `gorm:"column:bin_code" json:"binCode" example:"Dolorem est."`
	MaxCapacity   int            `gorm:"column:max_capacity" json:"maxCapacity" example:"-746189675918043600"`
	CurrentCount  int            `gorm:"column:current_count" json:"currentCount" example:"1622156784627261289"`
	CreatedAt     types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt     types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	InventoryList *[]Inventory   `gorm:"foreignKey:BinId;references:Id" json:"inventoryList,omitempty"`
}

func (p *StorageBins) TableName() string {
	return "logistics.storage_bins"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type StorageBinsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []StorageBins `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type StorageBinsEdit struct {
	Name         *string     `gorm:"column:name;not null" json:"name" example:"Assumenda molestiae."`
	ZoneId       *types.URID `gorm:"column:zone_id;not null" json:"zoneId" example:"X3EBBSEQ55GGVLHGFMKHSXLFOM"`
	BinCode      *string     `gorm:"column:bin_code" json:"binCode" example:"Temporibus aut."`
	MaxCapacity  *int        `gorm:"column:max_capacity" json:"maxCapacity" example:"1767210354452775193"`
	CurrentCount *int        `gorm:"column:current_count" json:"currentCount" example:"1701630289175280297"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type StorageBinsFilter struct {
	Id           *types.URID     `gorm:"column:id;primary_key" json:"id" example:"HEX6YXGE4ZHVPPV7WKTG5CAZKY"`
	Name         *string         `gorm:"column:name;not null" json:"name" example:"Distinctio rerum."`
	ZoneId       *types.URID     `gorm:"column:zone_id;not null" json:"zoneId" example:"K2JROWXMSBAKZIES4TBRPXIAFE"`
	BinCode      *string         `gorm:"column:bin_code" json:"binCode" example:"Incidunt repellendus."`
	MaxCapacity  *int            `gorm:"column:max_capacity" json:"maxCapacity" example:"8568309562517370871"`
	CurrentCount *int            `gorm:"column:current_count" json:"currentCount" example:"-4186547535944263194"`
	CreatedAt    *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt    *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *StorageBinsFilter) TableName() string {
	return "logistics.storage_bins"
}

// --- Batch Update Struct ---
type StorageBinsBatchUpdate struct {
	Data       StorageBinsEdit     `json:"data"`
	PathParams StorageBinsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type StorageBinsIdentity struct {
	Id types.URID `json:"id"`
}
