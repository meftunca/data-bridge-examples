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
// • OrderStatusHistoryForm     - Data input validation and creation
// • OrderStatusHistory        - Main database model with relationships
// • OrderStatusHistoryEdit    - Partial update operations
// • OrderStatusHistoryIdentity - Bulk operation identifiers
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
package orders_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"
	shared_types "data-bridge-examples/api_v2/shared/types"

	"github.com/maple-tech/baseline/types"
)

// OrderStatusHistoryForm handles data input validation and creation operations.
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
type OrderStatusHistoryForm struct {
	Name       string      `gorm:"column:name" json:"name" example:"Nobis quis."`
	OrderId    types.URID  `gorm:"column:order_id;not null" json:"orderId" example:"ONGMYNKQXVF3PLERHR626EESMU"`
	FromStatus string      `gorm:"column:from_status" json:"fromStatus" example:"Omnis quasi."`
	ToStatus   string      `gorm:"column:to_status" json:"toStatus" example:"Aut error."`
	ChangedBy  *types.URID `gorm:"column:changed_by" json:"changedBy,omitempty" example:"T4WXZPZPDFA7HDMWIL2YATUNQM"`
	Note       string      `gorm:"column:note" json:"note" example:"Repellat rerum."`
}

func (p *OrderStatusHistoryForm) TableName() string {
	return "orders.order_status_history"
}

// OrderStatusHistory represents the main database model for orders.order_status_history table.
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
// • Table: orders.order_status_history
// • Type: BASE TABLE
// • Schema: orders
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
type OrderStatusHistory struct {
	Id              types.URID             `gorm:"column:id;primary_key" json:"id" example:"4GOULX53O5CJDKYXJTBGG4WVIY"`
	Name            string                 `gorm:"column:name" json:"name" example:"Quasi non."`
	OrderId         types.URID             `gorm:"column:order_id;not null" json:"orderId" example:"PRJ6T6BCS5BBDO35E6CXKI7EDI"`
	OrderIdDetail   *Orders                `gorm:"foreignkey:OrderId" json:"orderDetail,omitempty"`
	FromStatus      string                 `gorm:"column:from_status" json:"fromStatus" example:"Et et."`
	ToStatus        string                 `gorm:"column:to_status" json:"toStatus" example:"Sunt veritatis."`
	ChangedBy       *types.URID            `gorm:"column:changed_by" json:"changedBy,omitempty" example:"IQRASKMIJRFMTBNJLQJXTGJQOY"`
	ChangedByDetail *shared_types.IamUsers `gorm:"foreignkey:ChangedBy" json:"changedByDetail,omitempty"`
	Note            string                 `gorm:"column:note" json:"note" example:"Omnis rerum."`
	CreatedAt       types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *OrderStatusHistory) TableName() string {
	return "orders.order_status_history"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type OrderStatusHistoryPage struct {
	paginationRuntime.DefaultPageResponse
	Items []OrderStatusHistory `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type OrderStatusHistoryEdit struct {
	Name       *string     `gorm:"column:name" json:"name" example:"Fugiat et."`
	OrderId    *types.URID `gorm:"column:order_id;not null" json:"orderId" example:"GUVS3UQZ2JCKZP5N4KTMC2FLLM"`
	FromStatus *string     `gorm:"column:from_status" json:"fromStatus" example:"Numquam ab."`
	ToStatus   *string     `gorm:"column:to_status" json:"toStatus" example:"Recusandae nihil."`
	ChangedBy  *types.URID `gorm:"column:changed_by" json:"changedBy,omitempty" example:"RD5DLUWB5JEXZP2JBJ7CWYIT2Q"`
	Note       *string     `gorm:"column:note" json:"note" example:"Itaque aut."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type OrderStatusHistoryFilter struct {
	Id         *types.URID     `gorm:"column:id;primary_key" json:"id" example:"UVMMR4E7VZBD3JLQGRMVE2IXKY"`
	Name       *string         `gorm:"column:name" json:"name" example:"Sunt accusantium."`
	OrderId    *types.URID     `gorm:"column:order_id;not null" json:"orderId" example:"JHN6MRUW7ZCOVBADTPFWEIB6OM"`
	FromStatus *string         `gorm:"column:from_status" json:"fromStatus" example:"Harum perspiciatis."`
	ToStatus   *string         `gorm:"column:to_status" json:"toStatus" example:"Et praesentium."`
	ChangedBy  *types.URID     `gorm:"column:changed_by" json:"changedBy,omitempty" example:"A3IYFU2RYBDODM76NEJDTTUYQI"`
	Note       *string         `gorm:"column:note" json:"note" example:"Culpa dolores."`
	CreatedAt  *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *OrderStatusHistoryFilter) TableName() string {
	return "orders.order_status_history"
}

// --- Batch Update Struct ---
type OrderStatusHistoryBatchUpdate struct {
	Data       OrderStatusHistoryEdit     `json:"data"`
	PathParams OrderStatusHistoryIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type OrderStatusHistoryIdentity struct {
	Id types.URID `json:"id"`
}
