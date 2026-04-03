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
// • RefundsForm     - Data input validation and creation
// • Refunds        - Main database model with relationships
// • RefundsEdit    - Partial update operations
// • RefundsIdentity - Bulk operation identifiers
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

// RefundsForm handles data input validation and creation operations.
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
type RefundsForm struct {
	Name        string         `gorm:"column:name" json:"name" example:"Et et."`
	OrderId     types.URID     `gorm:"column:order_id;not null" json:"orderId" example:"EB7INAC4R5C65JGXYJKHGTM6HI"`
	PaymentId   types.URID     `gorm:"column:payment_id;not null" json:"paymentId" example:"4S2UULLWNRCWXH3UHA7AOZTKWY"`
	Amount      float64        `gorm:"column:amount;not null" json:"amount"`
	Reason      string         `gorm:"column:reason" json:"reason" example:"Est unde."`
	Status      string         `gorm:"column:status" json:"status" example:"Cupiditate et."`
	ProcessedBy *types.URID    `gorm:"column:processed_by" json:"processedBy,omitempty" example:"KVPIVKKG4NE7ZIQBYWCTQTDXHQ"`
	ProcessedAt types.NullTime `gorm:"column:processed_at" json:"processedAt,omitempty"`
}

func (p *RefundsForm) TableName() string {
	return "orders.refunds"
}

// Refunds represents the main database model for orders.refunds table.
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
// • Table: orders.refunds
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
type Refunds struct {
	Id                types.URID             `gorm:"column:id;primary_key" json:"id" example:"LTMAUVPJQVA27LYKN2ERFGDMLE"`
	Name              string                 `gorm:"column:name" json:"name" example:"Sunt dolor."`
	OrderId           types.URID             `gorm:"column:order_id;not null" json:"orderId" example:"HFOHGKRHLVGEBNGSSWT5HMSKMY"`
	OrderIdDetail     *Orders                `gorm:"foreignkey:OrderId" json:"orderDetail,omitempty"`
	PaymentId         types.URID             `gorm:"column:payment_id;not null" json:"paymentId" example:"GLV3BIIRVVDITFAMUFSBVGIMHI"`
	PaymentIdDetail   *Payments              `gorm:"foreignkey:PaymentId" json:"paymentDetail,omitempty"`
	Amount            float64                `gorm:"column:amount;not null" json:"amount"`
	Reason            string                 `gorm:"column:reason" json:"reason" example:"Nam aut."`
	Status            string                 `gorm:"column:status" json:"status" example:"Est ratione."`
	ProcessedBy       *types.URID            `gorm:"column:processed_by" json:"processedBy,omitempty" example:"LO7AY6CWABCDBDBDKAGHH6JMPU"`
	ProcessedByDetail *shared_types.IamUsers `gorm:"foreignkey:ProcessedBy" json:"processedByDetail,omitempty"`
	ProcessedAt       types.NullTime         `gorm:"column:processed_at" json:"processedAt,omitempty"`
	CreatedAt         types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         types.NullTime         `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *Refunds) TableName() string {
	return "orders.refunds"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type RefundsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Refunds `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type RefundsEdit struct {
	Name        *string         `gorm:"column:name" json:"name" example:"Quos delectus."`
	OrderId     *types.URID     `gorm:"column:order_id;not null" json:"orderId" example:"WOTUL6XH2JCFZLURLXDDMU2ZBY"`
	PaymentId   *types.URID     `gorm:"column:payment_id;not null" json:"paymentId" example:"VIIPNFSDOJAMJMHUXLVZE4N6AY"`
	Amount      *float64        `gorm:"column:amount;not null" json:"amount"`
	Reason      *string         `gorm:"column:reason" json:"reason" example:"Libero et."`
	Status      *string         `gorm:"column:status" json:"status" example:"Quis non."`
	ProcessedBy *types.URID     `gorm:"column:processed_by" json:"processedBy,omitempty" example:"FRCN6F5XIVC2XNKAIT3GZULAXY"`
	ProcessedAt *types.NullTime `gorm:"column:processed_at" json:"processedAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type RefundsFilter struct {
	Id          *types.URID     `gorm:"column:id;primary_key" json:"id" example:"ROYEZDHAEZHHBEMGLGGDSCANF4"`
	Name        *string         `gorm:"column:name" json:"name" example:"Eveniet nulla."`
	OrderId     *types.URID     `gorm:"column:order_id;not null" json:"orderId" example:"MZZGHB7EO5HUHFHZL262ZY2TF4"`
	PaymentId   *types.URID     `gorm:"column:payment_id;not null" json:"paymentId" example:"GJJODLVU5ZDSNKDMNT5WKE7SRQ"`
	Amount      *float64        `gorm:"column:amount;not null" json:"amount"`
	Reason      *string         `gorm:"column:reason" json:"reason" example:"Aliquid minima."`
	Status      *string         `gorm:"column:status" json:"status" example:"Rerum adipisci."`
	ProcessedBy *types.URID     `gorm:"column:processed_by" json:"processedBy,omitempty" example:"MWWTT67L5BBKLCR2OBIYU3L6TA"`
	ProcessedAt *types.NullTime `gorm:"column:processed_at" json:"processedAt,omitempty"`
	CreatedAt   *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt   *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *RefundsFilter) TableName() string {
	return "orders.refunds"
}

// --- Batch Update Struct ---
type RefundsBatchUpdate struct {
	Data       RefundsEdit     `json:"data"`
	PathParams RefundsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type RefundsIdentity struct {
	Id types.URID `json:"id"`
}
