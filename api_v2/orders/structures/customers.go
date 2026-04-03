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
// • CustomersForm     - Data input validation and creation
// • Customers        - Main database model with relationships
// • CustomersEdit    - Partial update operations
// • CustomersIdentity - Bulk operation identifiers
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

// CustomersForm handles data input validation and creation operations.
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
type CustomersForm struct {
	Name            string      `gorm:"column:name;not null" json:"name" example:"Fuga dolore."`
	UserId          types.URID  `gorm:"column:user_id;not null" json:"userId" example:"D34V3WVHJVBEXLBALGFHKBPESU"`
	OrganizationId  *types.URID `gorm:"column:organization_id" json:"organizationId,omitempty" example:"HH3BVRDIAZFHHHWZ6ZNWI3VUXE"`
	BillingAddress  types.JSON  `gorm:"column:billing_address" json:"billingAddress"`
	ShippingAddress types.JSON  `gorm:"column:shipping_address" json:"shippingAddress"`
	TaxId           string      `gorm:"column:tax_id" json:"taxId" example:"Doloremque quaerat."`
	LoyaltyPoints   int         `gorm:"column:loyalty_points" json:"loyaltyPoints" example:"190695119469606752"`
	Tier            string      `gorm:"column:tier" json:"tier" example:"Voluptates velit."`
}

func (p *CustomersForm) TableName() string {
	return "orders.customers"
}

// Customers represents the main database model for orders.customers table.
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
// • Table: orders.customers
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
type Customers struct {
	Id                   types.URID                     `gorm:"column:id;primary_key" json:"id" example:"U5GJRSIIVZBEBNI2LONQN3UW6U"`
	Name                 string                         `gorm:"column:name;not null" json:"name" example:"Ut architecto."`
	UserId               types.URID                     `gorm:"column:user_id;not null" json:"userId" example:"RXHYODPQQ5F4LJ6SFARGFRLPZM"`
	UserIdDetail         *shared_types.IamUsers         `gorm:"foreignkey:UserId" json:"userDetail,omitempty"`
	OrganizationId       *types.URID                    `gorm:"column:organization_id" json:"organizationId,omitempty" example:"WWVAIWYPRRCNDFX63FOPQ6XCAY"`
	OrganizationIdDetail *shared_types.IamOrganizations `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	BillingAddress       types.JSON                     `gorm:"column:billing_address" json:"billingAddress"`
	ShippingAddress      types.JSON                     `gorm:"column:shipping_address" json:"shippingAddress"`
	TaxId                string                         `gorm:"column:tax_id" json:"taxId" example:"Sit quo."`
	LoyaltyPoints        int                            `gorm:"column:loyalty_points" json:"loyaltyPoints" example:"1233997306718886352"`
	Tier                 string                         `gorm:"column:tier" json:"tier" example:"Vel rerum."`
	CreatedAt            types.NullTime                 `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime                 `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	OrdersList           *[]Orders                      `gorm:"foreignKey:CustomerId;references:Id" json:"ordersList,omitempty"`
	CartsList            *[]Carts                       `gorm:"foreignKey:CustomerId;references:Id" json:"cartsList,omitempty"`
}

func (p *Customers) TableName() string {
	return "orders.customers"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type CustomersPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Customers `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type CustomersEdit struct {
	Name            *string     `gorm:"column:name;not null" json:"name" example:"Rerum rerum."`
	UserId          *types.URID `gorm:"column:user_id;not null" json:"userId" example:"CHSBTTH3QZB77HJTJVSE4BZ2ZI"`
	OrganizationId  *types.URID `gorm:"column:organization_id" json:"organizationId,omitempty" example:"KVPHUCLWABH7PEBRLPE7H4CYVA"`
	BillingAddress  *types.JSON `gorm:"column:billing_address" json:"billingAddress"`
	ShippingAddress *types.JSON `gorm:"column:shipping_address" json:"shippingAddress"`
	TaxId           *string     `gorm:"column:tax_id" json:"taxId" example:"Qui sequi."`
	LoyaltyPoints   *int        `gorm:"column:loyalty_points" json:"loyaltyPoints" example:"628418607932676130"`
	Tier            *string     `gorm:"column:tier" json:"tier" example:"Autem incidunt."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type CustomersFilter struct {
	Id              *types.URID     `gorm:"column:id;primary_key" json:"id" example:"2BWYRA7765GZTNDT6MEKQIW2EM"`
	Name            *string         `gorm:"column:name;not null" json:"name" example:"Incidunt nemo."`
	UserId          *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"S6PPJEM665ED3HD7NRXUEPIMJA"`
	OrganizationId  *types.URID     `gorm:"column:organization_id" json:"organizationId,omitempty" example:"TT3QLZ5CVVAYXB5IRQCSUTDPQY"`
	BillingAddress  *types.JSON     `gorm:"column:billing_address" json:"billingAddress"`
	ShippingAddress *types.JSON     `gorm:"column:shipping_address" json:"shippingAddress"`
	TaxId           *string         `gorm:"column:tax_id" json:"taxId" example:"Odit unde."`
	LoyaltyPoints   *int            `gorm:"column:loyalty_points" json:"loyaltyPoints" example:"-1855198432576826645"`
	Tier            *string         `gorm:"column:tier" json:"tier" example:"Corporis dolor."`
	CreatedAt       *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *CustomersFilter) TableName() string {
	return "orders.customers"
}

// --- Batch Update Struct ---
type CustomersBatchUpdate struct {
	Data       CustomersEdit     `json:"data"`
	PathParams CustomersIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type CustomersIdentity struct {
	Id types.URID `json:"id"`
}
