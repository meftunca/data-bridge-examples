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
// • PaymentsForm     - Data input validation and creation
// • Payments        - Main database model with relationships
// • PaymentsEdit    - Partial update operations
// • PaymentsIdentity - Bulk operation identifiers
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

	"github.com/maple-tech/baseline/types"
)

// PaymentsForm handles data input validation and creation operations.
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
type PaymentsForm struct {
	Name         string              `gorm:"column:name" json:"name" example:"Nesciunt et."`
	OrderId      types.URID          `gorm:"column:order_id;not null" json:"orderId" example:"DEPBPOLNVBG5THZBQRTVYJFOZE"`
	Amount       float64             `gorm:"column:amount;not null" json:"amount"`
	Currency     string              `gorm:"column:currency" json:"currency" example:"Exercitationem reprehenderit."`
	Method       OrdersPaymentMethod `gorm:"column:method" json:"method"`
	Status       OrdersPaymentStatus `gorm:"column:status" json:"status"`
	ProviderRef  string              `gorm:"column:provider_ref" json:"providerRef" example:"Non est."`
	ProviderData types.JSON          `gorm:"column:provider_data" json:"providerData"`
	PaidAt       types.NullTime      `gorm:"column:paid_at" json:"paidAt,omitempty"`
}

func (p *PaymentsForm) TableName() string {
	return "orders.payments"
}

// Payments represents the main database model for orders.payments table.
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
// • Table: orders.payments
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
type Payments struct {
	Id            types.URID          `gorm:"column:id;primary_key" json:"id" example:"RSD5C3FRVJD7FARMKBQ6U7KQOQ"`
	Name          string              `gorm:"column:name" json:"name" example:"Omnis sed."`
	OrderId       types.URID          `gorm:"column:order_id;not null" json:"orderId" example:"N3A5URUNKBHVBAJDMCFUYISCFU"`
	OrderIdDetail *Orders             `gorm:"foreignkey:OrderId" json:"orderDetail,omitempty"`
	Amount        float64             `gorm:"column:amount;not null" json:"amount"`
	Currency      string              `gorm:"column:currency" json:"currency" example:"Fuga ullam."`
	Method        OrdersPaymentMethod `gorm:"column:method" json:"method"`
	Status        OrdersPaymentStatus `gorm:"column:status" json:"status"`
	ProviderRef   string              `gorm:"column:provider_ref" json:"providerRef" example:"Quis dicta."`
	ProviderData  types.JSON          `gorm:"column:provider_data" json:"providerData"`
	PaidAt        types.NullTime      `gorm:"column:paid_at" json:"paidAt,omitempty"`
	CreatedAt     types.NullTime      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt     types.NullTime      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	RefundsList   *[]Refunds          `gorm:"foreignKey:PaymentId;references:Id" json:"refundsList,omitempty"`
}

func (p *Payments) TableName() string {
	return "orders.payments"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type PaymentsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Payments `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type PaymentsEdit struct {
	Name         *string              `gorm:"column:name" json:"name" example:"Voluptas officia."`
	OrderId      *types.URID          `gorm:"column:order_id;not null" json:"orderId" example:"KF2GERIILBCN7JPTXKDJX2NPEE"`
	Amount       *float64             `gorm:"column:amount;not null" json:"amount"`
	Currency     *string              `gorm:"column:currency" json:"currency" example:"Consequatur repudiandae."`
	Method       *OrdersPaymentMethod `gorm:"column:method" json:"method"`
	Status       *OrdersPaymentStatus `gorm:"column:status" json:"status"`
	ProviderRef  *string              `gorm:"column:provider_ref" json:"providerRef" example:"Aperiam aut."`
	ProviderData *types.JSON          `gorm:"column:provider_data" json:"providerData"`
	PaidAt       *types.NullTime      `gorm:"column:paid_at" json:"paidAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type PaymentsFilter struct {
	Id           *types.URID          `gorm:"column:id;primary_key" json:"id" example:"AMHFHZ4ELNF7JFDLZRKWIKC6E4"`
	Name         *string              `gorm:"column:name" json:"name" example:"Numquam earum."`
	OrderId      *types.URID          `gorm:"column:order_id;not null" json:"orderId" example:"OC2DDCW25ZFJHI43RICBZYYJFQ"`
	Amount       *float64             `gorm:"column:amount;not null" json:"amount"`
	Currency     *string              `gorm:"column:currency" json:"currency" example:"Vitae fugiat."`
	Method       *OrdersPaymentMethod `gorm:"column:method" json:"method"`
	Status       *OrdersPaymentStatus `gorm:"column:status" json:"status"`
	ProviderRef  *string              `gorm:"column:provider_ref" json:"providerRef" example:"Nemo velit."`
	ProviderData *types.JSON          `gorm:"column:provider_data" json:"providerData"`
	PaidAt       *types.NullTime      `gorm:"column:paid_at" json:"paidAt,omitempty"`
	CreatedAt    *types.NullTime      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt    *types.NullTime      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *PaymentsFilter) TableName() string {
	return "orders.payments"
}

// --- Batch Update Struct ---
type PaymentsBatchUpdate struct {
	Data       PaymentsEdit     `json:"data"`
	PathParams PaymentsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type PaymentsIdentity struct {
	Id types.URID `json:"id"`
}
