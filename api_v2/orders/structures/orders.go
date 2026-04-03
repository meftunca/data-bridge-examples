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
// • OrdersForm     - Data input validation and creation
// • Orders        - Main database model with relationships
// • OrdersEdit    - Partial update operations
// • OrdersIdentity - Bulk operation identifiers
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

// OrdersForm handles data input validation and creation operations.
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
type OrdersForm struct {
	Name            string            `gorm:"column:name" json:"name" example:"Eveniet voluptatem."`
	OrderNumber     string            `gorm:"column:order_number;not null" json:"orderNumber" example:"Est commodi."`
	CustomerId      types.URID        `gorm:"column:customer_id;not null" json:"customerId" example:"RBQCEYRQFFH5XHKFCTIVWXWS2E"`
	Status          OrdersOrderStatus `gorm:"column:status" json:"status"`
	Subtotal        float64           `gorm:"column:subtotal" json:"subtotal"`
	TaxAmount       float64           `gorm:"column:tax_amount" json:"taxAmount"`
	ShippingAmount  float64           `gorm:"column:shipping_amount" json:"shippingAmount"`
	DiscountAmount  float64           `gorm:"column:discount_amount" json:"discountAmount"`
	Total           float64           `gorm:"column:total" json:"total"`
	Currency        string            `gorm:"column:currency" json:"currency" example:"Et facilis."`
	CouponId        *types.URID       `gorm:"column:coupon_id" json:"couponId,omitempty" example:"JQQ6ON3RFRAKLKKPOFAZO2UDBE"`
	ShippingAddress types.JSON        `gorm:"column:shipping_address" json:"shippingAddress"`
	BillingAddress  types.JSON        `gorm:"column:billing_address" json:"billingAddress"`
	Notes           string            `gorm:"column:notes" json:"notes" example:"Error odio."`
	PlacedAt        types.NullTime    `gorm:"column:placed_at" json:"placedAt,omitempty"`
	ConfirmedAt     types.NullTime    `gorm:"column:confirmed_at" json:"confirmedAt,omitempty"`
	ShippedAt       types.NullTime    `gorm:"column:shipped_at" json:"shippedAt,omitempty"`
	DeliveredAt     types.NullTime    `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`
	CancelledAt     types.NullTime    `gorm:"column:cancelled_at" json:"cancelledAt,omitempty"`
	CreatedBy       *types.URID       `gorm:"column:created_by" json:"createdBy,omitempty" example:"ITVV7BKFEVBOVNVJ3ZDJLPJRZM"`
}

func (p *OrdersForm) TableName() string {
	return "orders.orders"
}

// Orders represents the main database model for orders.orders table.
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
// • Table: orders.orders
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
type Orders struct {
	Id                     types.URID                         `gorm:"column:id;primary_key" json:"id" example:"OO3WZ5QRL5G5TPYBZHVILCHS64"`
	Name                   string                             `gorm:"column:name" json:"name" example:"Libero officia."`
	OrderNumber            string                             `gorm:"column:order_number;not null" json:"orderNumber" example:"Optio error."`
	CustomerId             types.URID                         `gorm:"column:customer_id;not null" json:"customerId" example:"MMBYVO3XQ5BXFM2NWTGFWZIHM4"`
	CustomerIdDetail       *Customers                         `gorm:"foreignkey:CustomerId" json:"customerDetail,omitempty"`
	Status                 OrdersOrderStatus                  `gorm:"column:status" json:"status"`
	Subtotal               float64                            `gorm:"column:subtotal" json:"subtotal"`
	TaxAmount              float64                            `gorm:"column:tax_amount" json:"taxAmount"`
	ShippingAmount         float64                            `gorm:"column:shipping_amount" json:"shippingAmount"`
	DiscountAmount         float64                            `gorm:"column:discount_amount" json:"discountAmount"`
	Total                  float64                            `gorm:"column:total" json:"total"`
	Currency               string                             `gorm:"column:currency" json:"currency" example:"Eos voluptatem."`
	CouponId               *types.URID                        `gorm:"column:coupon_id" json:"couponId,omitempty" example:"L2B6QOD75REL7C56LADW6FTQCA"`
	CouponIdDetail         *Coupons                           `gorm:"foreignkey:CouponId" json:"couponDetail,omitempty"`
	ShippingAddress        types.JSON                         `gorm:"column:shipping_address" json:"shippingAddress"`
	BillingAddress         types.JSON                         `gorm:"column:billing_address" json:"billingAddress"`
	Notes                  string                             `gorm:"column:notes" json:"notes" example:"Minima natus."`
	PlacedAt               types.NullTime                     `gorm:"column:placed_at" json:"placedAt,omitempty"`
	ConfirmedAt            types.NullTime                     `gorm:"column:confirmed_at" json:"confirmedAt,omitempty"`
	ShippedAt              types.NullTime                     `gorm:"column:shipped_at" json:"shippedAt,omitempty"`
	DeliveredAt            types.NullTime                     `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`
	CancelledAt            types.NullTime                     `gorm:"column:cancelled_at" json:"cancelledAt,omitempty"`
	CreatedBy              *types.URID                        `gorm:"column:created_by" json:"createdBy,omitempty" example:"6NEMVB7TTNFH5MDRK4DFCCNJ4Y"`
	CreatedByDetail        *shared_types.IamUsers             `gorm:"foreignkey:CreatedBy" json:"createdByDetail,omitempty"`
	CreatedAt              types.NullTime                     `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt              types.NullTime                     `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt              types.NullTime                     `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
	OrderItemsList         *[]OrderItems                      `gorm:"foreignKey:OrderId;references:Id" json:"orderItemsList,omitempty"`
	PaymentsList           *[]Payments                        `gorm:"foreignKey:OrderId;references:Id" json:"paymentsList,omitempty"`
	RefundsList            *[]Refunds                         `gorm:"foreignKey:OrderId;references:Id" json:"refundsList,omitempty"`
	OrderStatusHistoryList *[]OrderStatusHistory              `gorm:"foreignKey:OrderId;references:Id" json:"orderStatusHistoryList,omitempty"`
	ShipmentsList          *[]shared_types.LogisticsShipments `gorm:"foreignKey:OrderId;references:Id" json:"shipmentsList,omitempty"`
}

func (p *Orders) TableName() string {
	return "orders.orders"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type OrdersPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Orders `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type OrdersEdit struct {
	Name            *string            `gorm:"column:name" json:"name" example:"Ullam earum."`
	OrderNumber     *string            `gorm:"column:order_number;not null" json:"orderNumber" example:"Explicabo ad."`
	CustomerId      *types.URID        `gorm:"column:customer_id;not null" json:"customerId" example:"HD4KFM2VTVBO5AHEUTSNBOH6NM"`
	Status          *OrdersOrderStatus `gorm:"column:status" json:"status"`
	Subtotal        *float64           `gorm:"column:subtotal" json:"subtotal"`
	TaxAmount       *float64           `gorm:"column:tax_amount" json:"taxAmount"`
	ShippingAmount  *float64           `gorm:"column:shipping_amount" json:"shippingAmount"`
	DiscountAmount  *float64           `gorm:"column:discount_amount" json:"discountAmount"`
	Total           *float64           `gorm:"column:total" json:"total"`
	Currency        *string            `gorm:"column:currency" json:"currency" example:"Ut dolorem."`
	CouponId        *types.URID        `gorm:"column:coupon_id" json:"couponId,omitempty" example:"KZMZYEJYYBC4HM5IYZXSMBQ3UE"`
	ShippingAddress *types.JSON        `gorm:"column:shipping_address" json:"shippingAddress"`
	BillingAddress  *types.JSON        `gorm:"column:billing_address" json:"billingAddress"`
	Notes           *string            `gorm:"column:notes" json:"notes" example:"Sunt est."`
	PlacedAt        *types.NullTime    `gorm:"column:placed_at" json:"placedAt,omitempty"`
	ConfirmedAt     *types.NullTime    `gorm:"column:confirmed_at" json:"confirmedAt,omitempty"`
	ShippedAt       *types.NullTime    `gorm:"column:shipped_at" json:"shippedAt,omitempty"`
	DeliveredAt     *types.NullTime    `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`
	CancelledAt     *types.NullTime    `gorm:"column:cancelled_at" json:"cancelledAt,omitempty"`
	CreatedBy       *types.URID        `gorm:"column:created_by" json:"createdBy,omitempty" example:"MWKHHIRRL5CVDECMM7Q2EHR3SQ"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type OrdersFilter struct {
	Id              *types.URID        `gorm:"column:id;primary_key" json:"id" example:"XCBETESIKVC4ZD7Q7QBHHOQOWU"`
	Name            *string            `gorm:"column:name" json:"name" example:"Sit saepe."`
	OrderNumber     *string            `gorm:"column:order_number;not null" json:"orderNumber" example:"Consequatur ullam."`
	CustomerId      *types.URID        `gorm:"column:customer_id;not null" json:"customerId" example:"4Y5GCVREJJFBRH7SBBHS75Y6VM"`
	Status          *OrdersOrderStatus `gorm:"column:status" json:"status"`
	Subtotal        *float64           `gorm:"column:subtotal" json:"subtotal"`
	TaxAmount       *float64           `gorm:"column:tax_amount" json:"taxAmount"`
	ShippingAmount  *float64           `gorm:"column:shipping_amount" json:"shippingAmount"`
	DiscountAmount  *float64           `gorm:"column:discount_amount" json:"discountAmount"`
	Total           *float64           `gorm:"column:total" json:"total"`
	Currency        *string            `gorm:"column:currency" json:"currency" example:"Officia voluptatem."`
	CouponId        *types.URID        `gorm:"column:coupon_id" json:"couponId,omitempty" example:"PKJZERIE6FBWJAL5KKWEGHP52A"`
	ShippingAddress *types.JSON        `gorm:"column:shipping_address" json:"shippingAddress"`
	BillingAddress  *types.JSON        `gorm:"column:billing_address" json:"billingAddress"`
	Notes           *string            `gorm:"column:notes" json:"notes" example:"Modi ut."`
	PlacedAt        *types.NullTime    `gorm:"column:placed_at" json:"placedAt,omitempty"`
	ConfirmedAt     *types.NullTime    `gorm:"column:confirmed_at" json:"confirmedAt,omitempty"`
	ShippedAt       *types.NullTime    `gorm:"column:shipped_at" json:"shippedAt,omitempty"`
	DeliveredAt     *types.NullTime    `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`
	CancelledAt     *types.NullTime    `gorm:"column:cancelled_at" json:"cancelledAt,omitempty"`
	CreatedBy       *types.URID        `gorm:"column:created_by" json:"createdBy,omitempty" example:"ND2PONIFWVGONKUDVLUUH3GGQE"`
	CreatedAt       *types.NullTime    `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       *types.NullTime    `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt       *types.NullTime    `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}

func (p *OrdersFilter) TableName() string {
	return "orders.orders"
}

// --- Batch Update Struct ---
type OrdersBatchUpdate struct {
	Data       OrdersEdit     `json:"data"`
	PathParams OrdersIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type OrdersIdentity struct {
	Id types.URID `json:"id"`
}
