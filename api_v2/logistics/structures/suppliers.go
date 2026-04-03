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
// • SuppliersForm     - Data input validation and creation
// • Suppliers        - Main database model with relationships
// • SuppliersEdit    - Partial update operations
// • SuppliersIdentity - Bulk operation identifiers
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

// SuppliersForm handles data input validation and creation operations.
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
type SuppliersForm struct {
	Name         string     `gorm:"column:name;not null" json:"name" example:"Est suscipit."`
	Code         string     `gorm:"column:code;not null" json:"code" example:"Aut aut."`
	ContactName  string     `gorm:"column:contact_name" json:"contactName" example:"Reprehenderit velit."`
	ContactEmail string     `gorm:"column:contact_email" json:"contactEmail" example:"Quo nobis."`
	ContactPhone string     `gorm:"column:contact_phone" json:"contactPhone" example:"Id amet."`
	Address      types.JSON `gorm:"column:address" json:"address"`
	IsActive     bool       `gorm:"column:is_active" json:"isActive" example:"true"`
	Rating       float64    `gorm:"column:rating" json:"rating"`
}

func (p *SuppliersForm) TableName() string {
	return "logistics.suppliers"
}

// Suppliers represents the main database model for logistics.suppliers table.
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
// • Table: logistics.suppliers
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
type Suppliers struct {
	Id                 types.URID        `gorm:"column:id;primary_key" json:"id" example:"JFDNB4AYMVBU7FE44UUKRFVFY4"`
	Name               string            `gorm:"column:name;not null" json:"name" example:"Est non."`
	Code               string            `gorm:"column:code;not null" json:"code" example:"Et aut."`
	ContactName        string            `gorm:"column:contact_name" json:"contactName" example:"Et dolore."`
	ContactEmail       string            `gorm:"column:contact_email" json:"contactEmail" example:"Qui dolorem."`
	ContactPhone       string            `gorm:"column:contact_phone" json:"contactPhone" example:"Voluptas voluptatem."`
	Address            types.JSON        `gorm:"column:address" json:"address"`
	IsActive           bool              `gorm:"column:is_active" json:"isActive" example:"true"`
	Rating             float64           `gorm:"column:rating" json:"rating"`
	CreatedAt          types.NullTime    `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt          types.NullTime    `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	PurchaseOrdersList *[]PurchaseOrders `gorm:"foreignKey:SupplierId;references:Id" json:"purchaseOrdersList,omitempty"`
}

func (p *Suppliers) TableName() string {
	return "logistics.suppliers"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type SuppliersPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Suppliers `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type SuppliersEdit struct {
	Name         *string     `gorm:"column:name;not null" json:"name" example:"Qui nihil."`
	Code         *string     `gorm:"column:code;not null" json:"code" example:"Vel quo."`
	ContactName  *string     `gorm:"column:contact_name" json:"contactName" example:"Velit itaque."`
	ContactEmail *string     `gorm:"column:contact_email" json:"contactEmail" example:"Fugit sit."`
	ContactPhone *string     `gorm:"column:contact_phone" json:"contactPhone" example:"Sapiente quaerat."`
	Address      *types.JSON `gorm:"column:address" json:"address"`
	IsActive     *bool       `gorm:"column:is_active" json:"isActive" example:"true"`
	Rating       *float64    `gorm:"column:rating" json:"rating"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type SuppliersFilter struct {
	Id           *types.URID     `gorm:"column:id;primary_key" json:"id" example:"YXZ6PNHOKBGW5AUDYHAI6NBZHI"`
	Name         *string         `gorm:"column:name;not null" json:"name" example:"Eius sequi."`
	Code         *string         `gorm:"column:code;not null" json:"code" example:"Quas et."`
	ContactName  *string         `gorm:"column:contact_name" json:"contactName" example:"Enim accusamus."`
	ContactEmail *string         `gorm:"column:contact_email" json:"contactEmail" example:"Voluptatem earum."`
	ContactPhone *string         `gorm:"column:contact_phone" json:"contactPhone" example:"Aut sunt."`
	Address      *types.JSON     `gorm:"column:address" json:"address"`
	IsActive     *bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	Rating       *float64        `gorm:"column:rating" json:"rating"`
	CreatedAt    *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt    *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *SuppliersFilter) TableName() string {
	return "logistics.suppliers"
}

// --- Batch Update Struct ---
type SuppliersBatchUpdate struct {
	Data       SuppliersEdit     `json:"data"`
	PathParams SuppliersIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type SuppliersIdentity struct {
	Id types.URID `json:"id"`
}
