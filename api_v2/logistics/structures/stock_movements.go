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
// • StockMovementsForm     - Data input validation and creation
// • StockMovements        - Main database model with relationships
// • StockMovementsEdit    - Partial update operations
// • StockMovementsIdentity - Bulk operation identifiers
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

// StockMovementsForm handles data input validation and creation operations.
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
type StockMovementsForm struct {
	Name          string      `gorm:"column:name" json:"name" example:"Numquam et."`
	InventoryId   types.URID  `gorm:"column:inventory_id;not null" json:"inventoryId" example:"OVPCTTEUQJCYXNMVEGYHV5PJDY"`
	MovementType  string      `gorm:"column:movement_type" json:"movementType" example:"Corporis suscipit."`
	Quantity      int         `gorm:"column:quantity" json:"quantity" example:"-2659835845657338663"`
	ReferenceType string      `gorm:"column:reference_type" json:"referenceType" example:"Et nihil."`
	ReferenceId   string      `gorm:"column:reference_id" json:"referenceId" example:"Optio ea."`
	PerformedBy   *types.URID `gorm:"column:performed_by" json:"performedBy,omitempty" example:"OYGPWNZQPRAWLL3YLHLRLJLHCE"`
	Notes         string      `gorm:"column:notes" json:"notes" example:"Impedit accusantium."`
}

func (p *StockMovementsForm) TableName() string {
	return "logistics.stock_movements"
}

// StockMovements represents the main database model for logistics.stock_movements table.
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
// • Table: logistics.stock_movements
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
type StockMovements struct {
	Id                types.URID             `gorm:"column:id;primary_key" json:"id" example:"5QZAZAKTJNBZ7BMXFGPGSWDSSY"`
	Name              string                 `gorm:"column:name" json:"name" example:"Non laboriosam."`
	InventoryId       types.URID             `gorm:"column:inventory_id;not null" json:"inventoryId" example:"AENINDQK6FCEFAGOBMLFQA4NM4"`
	InventoryIdDetail *Inventory             `gorm:"foreignkey:InventoryId" json:"inventoryDetail,omitempty"`
	MovementType      string                 `gorm:"column:movement_type" json:"movementType" example:"Ut ut."`
	Quantity          int                    `gorm:"column:quantity" json:"quantity" example:"-6846136306775492352"`
	ReferenceType     string                 `gorm:"column:reference_type" json:"referenceType" example:"Ratione necessitatibus."`
	ReferenceId       string                 `gorm:"column:reference_id" json:"referenceId" example:"Aperiam qui."`
	PerformedBy       *types.URID            `gorm:"column:performed_by" json:"performedBy,omitempty" example:"VLBY6ECENZCHPLM5XBPCWHU24I"`
	PerformedByDetail *shared_types.IamUsers `gorm:"foreignkey:PerformedBy" json:"performedByDetail,omitempty"`
	Notes             string                 `gorm:"column:notes" json:"notes" example:"Quos aliquam."`
	CreatedAt         types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *StockMovements) TableName() string {
	return "logistics.stock_movements"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type StockMovementsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []StockMovements `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type StockMovementsEdit struct {
	Name          *string     `gorm:"column:name" json:"name" example:"Numquam magni."`
	InventoryId   *types.URID `gorm:"column:inventory_id;not null" json:"inventoryId" example:"2CZPJCHGIFGJJJZQBPOBSHILNQ"`
	MovementType  *string     `gorm:"column:movement_type" json:"movementType" example:"Quidem quas."`
	Quantity      *int        `gorm:"column:quantity" json:"quantity" example:"5095321399209733229"`
	ReferenceType *string     `gorm:"column:reference_type" json:"referenceType" example:"Voluptatem officiis."`
	ReferenceId   *string     `gorm:"column:reference_id" json:"referenceId" example:"Dicta id."`
	PerformedBy   *types.URID `gorm:"column:performed_by" json:"performedBy,omitempty" example:"ORWUWVCW3FG2LAQA52O4GH6Z74"`
	Notes         *string     `gorm:"column:notes" json:"notes" example:"Aliquid facilis."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type StockMovementsFilter struct {
	Id            *types.URID     `gorm:"column:id;primary_key" json:"id" example:"GQCOSJDUZFHJXNYD2E3YHXJNII"`
	Name          *string         `gorm:"column:name" json:"name" example:"Autem mollitia."`
	InventoryId   *types.URID     `gorm:"column:inventory_id;not null" json:"inventoryId" example:"52MZ2NWI35GIZJNRSLHPJVYCR4"`
	MovementType  *string         `gorm:"column:movement_type" json:"movementType" example:"Recusandae rerum."`
	Quantity      *int            `gorm:"column:quantity" json:"quantity" example:"-1054170702471037594"`
	ReferenceType *string         `gorm:"column:reference_type" json:"referenceType" example:"Odio consequuntur."`
	ReferenceId   *string         `gorm:"column:reference_id" json:"referenceId" example:"Aperiam quo."`
	PerformedBy   *types.URID     `gorm:"column:performed_by" json:"performedBy,omitempty" example:"DKJWC3FHQVARTBB6B6GU4X4JYA"`
	Notes         *string         `gorm:"column:notes" json:"notes" example:"Qui doloremque."`
	CreatedAt     *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *StockMovementsFilter) TableName() string {
	return "logistics.stock_movements"
}

// --- Batch Update Struct ---
type StockMovementsBatchUpdate struct {
	Data       StockMovementsEdit     `json:"data"`
	PathParams StockMovementsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type StockMovementsIdentity struct {
	Id types.URID `json:"id"`
}
