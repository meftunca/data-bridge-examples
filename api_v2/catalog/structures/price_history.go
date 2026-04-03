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
// • PriceHistoryForm     - Data input validation and creation
// • PriceHistory        - Main database model with relationships
// • PriceHistoryEdit    - Partial update operations
// • PriceHistoryIdentity - Bulk operation identifiers
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
package catalog_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"
	shared_types "data-bridge-examples/api_v2/shared/types"

	"github.com/maple-tech/baseline/types"
)

// PriceHistoryForm handles data input validation and creation operations.
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
type PriceHistoryForm struct {
	Name      string      `gorm:"column:name" json:"name" example:"Molestiae quasi."`
	ProductId types.URID  `gorm:"column:product_id;not null" json:"productId" example:"IKMPX5AG5VGRZMDDFXV3WEVCYA"`
	OldPrice  float64     `gorm:"column:old_price;not null" json:"oldPrice"`
	NewPrice  float64     `gorm:"column:new_price;not null" json:"newPrice"`
	ChangedBy *types.URID `gorm:"column:changed_by" json:"changedBy,omitempty" example:"3GAVLK66UJC7ZHWE65M3NDE4HQ"`
	Reason    string      `gorm:"column:reason" json:"reason" example:"Aut voluptatem."`
}

func (p *PriceHistoryForm) TableName() string {
	return "catalog.price_history"
}

// PriceHistory represents the main database model for catalog.price_history table.
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
// • Table: catalog.price_history
// • Type: BASE TABLE
// • Schema: catalog
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
type PriceHistory struct {
	Id              types.URID             `gorm:"column:id;primary_key" json:"id" example:"U4FSZUNBHJCIXLGILOAU4FUMEQ"`
	Name            string                 `gorm:"column:name" json:"name" example:"Eos quam."`
	ProductId       types.URID             `gorm:"column:product_id;not null" json:"productId" example:"6QBZ27LU3JH33LV3GD3TWEE3KY"`
	ProductIdDetail *Products              `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	OldPrice        float64                `gorm:"column:old_price;not null" json:"oldPrice"`
	NewPrice        float64                `gorm:"column:new_price;not null" json:"newPrice"`
	ChangedBy       *types.URID            `gorm:"column:changed_by" json:"changedBy,omitempty" example:"NYM725MRDNGJ3LFDZJCVPCLOCY"`
	ChangedByDetail *shared_types.IamUsers `gorm:"foreignkey:ChangedBy" json:"changedByDetail,omitempty"`
	Reason          string                 `gorm:"column:reason" json:"reason" example:"Illum non."`
	CreatedAt       types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *PriceHistory) TableName() string {
	return "catalog.price_history"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type PriceHistoryPage struct {
	paginationRuntime.DefaultPageResponse
	Items []PriceHistory `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type PriceHistoryEdit struct {
	Name      *string     `gorm:"column:name" json:"name" example:"Sit qui."`
	ProductId *types.URID `gorm:"column:product_id;not null" json:"productId" example:"42QJ26THY5EHNE2Z726OZDSK44"`
	OldPrice  *float64    `gorm:"column:old_price;not null" json:"oldPrice"`
	NewPrice  *float64    `gorm:"column:new_price;not null" json:"newPrice"`
	ChangedBy *types.URID `gorm:"column:changed_by" json:"changedBy,omitempty" example:"FHGW2PJKQBB3XEPGP2KUSMEINM"`
	Reason    *string     `gorm:"column:reason" json:"reason" example:"Iure dicta."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type PriceHistoryFilter struct {
	Id        *types.URID     `gorm:"column:id;primary_key" json:"id" example:"AM7NRP7AJVAQTCCL2VSOC6TL6I"`
	Name      *string         `gorm:"column:name" json:"name" example:"Reiciendis voluptates."`
	ProductId *types.URID     `gorm:"column:product_id;not null" json:"productId" example:"GN2VWKXI5NFL7CHBXLFNUTD6MI"`
	OldPrice  *float64        `gorm:"column:old_price;not null" json:"oldPrice"`
	NewPrice  *float64        `gorm:"column:new_price;not null" json:"newPrice"`
	ChangedBy *types.URID     `gorm:"column:changed_by" json:"changedBy,omitempty" example:"BJL2RRKYEVBMLGEF4UWNMO52LQ"`
	Reason    *string         `gorm:"column:reason" json:"reason" example:"Aut alias."`
	CreatedAt *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *PriceHistoryFilter) TableName() string {
	return "catalog.price_history"
}

// --- Batch Update Struct ---
type PriceHistoryBatchUpdate struct {
	Data       PriceHistoryEdit     `json:"data"`
	PathParams PriceHistoryIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type PriceHistoryIdentity struct {
	Id types.URID `json:"id"`
}
