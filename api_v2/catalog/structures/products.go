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
// • ProductsForm     - Data input validation and creation
// • Products        - Main database model with relationships
// • ProductsEdit    - Partial update operations
// • ProductsIdentity - Bulk operation identifiers
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

// ProductsForm handles data input validation and creation operations.
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
type ProductsForm struct {
	Name             string               `gorm:"column:name;not null" json:"name" example:"Et voluptatem."`
	Slug             string               `gorm:"column:slug;not null" json:"slug" example:"Voluptatem quod."`
	Sku              string               `gorm:"column:sku;not null" json:"sku" example:"Numquam ut."`
	Description      string               `gorm:"column:description" json:"description" example:"Quas dolores."`
	ShortDescription string               `gorm:"column:short_description" json:"shortDescription" example:"Tenetur voluptates."`
	Status           CatalogProductStatus `gorm:"column:status" json:"status"`
	BrandId          *types.URID          `gorm:"column:brand_id" json:"brandId,omitempty" example:"FJNXCCVASZCA3LB5XVCKXDH7KU"`
	CategoryId       *types.URID          `gorm:"column:category_id" json:"categoryId,omitempty" example:"CG4ISXREOVFCVA4YEALJM4PSH4"`
	BasePrice        float64              `gorm:"column:base_price" json:"basePrice"`
	Currency         string               `gorm:"column:currency" json:"currency" example:"Quis molestiae."`
	WeightKg         *float64             `gorm:"column:weight_kg" json:"weightKg,omitempty"`
	DimensionsCm     types.JSON           `gorm:"column:dimensions_cm" json:"dimensionsCm"`
	Attributes       types.JSON           `gorm:"column:attributes" json:"attributes"`
	Tags             types.StringArray    `gorm:"column:tags" json:"tags"`
	IsFeatured       bool                 `gorm:"column:is_featured" json:"isFeatured" example:"true"`
	CreatedBy        *types.URID          `gorm:"column:created_by" json:"createdBy,omitempty" example:"J6ODFKJQZRG6BLLBGJF2XNBVHM"`
	UpdatedBy        *types.URID          `gorm:"column:updated_by" json:"updatedBy,omitempty" example:"ZLSV56J3KJC63FCL4IUJHM4NHM"`
	SearchVector     *string              `gorm:"column:search_vector" json:"searchVector,omitempty" example:"Delectus culpa."`
}

func (p *ProductsForm) TableName() string {
	return "catalog.products"
}

// Products represents the main database model for catalog.products table.
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
// • Table: catalog.products
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
type Products struct {
	Id                     types.URID                                  `gorm:"column:id;primary_key" json:"id" example:"NTE7EPZEY5DLRD3BMWMW5FIRGQ"`
	Name                   string                                      `gorm:"column:name;not null" json:"name" example:"Assumenda cumque."`
	Slug                   string                                      `gorm:"column:slug;not null" json:"slug" example:"In quis."`
	Sku                    string                                      `gorm:"column:sku;not null" json:"sku" example:"Est omnis."`
	Description            string                                      `gorm:"column:description" json:"description" example:"Et qui."`
	ShortDescription       string                                      `gorm:"column:short_description" json:"shortDescription" example:"Laudantium porro."`
	Status                 CatalogProductStatus                        `gorm:"column:status" json:"status"`
	BrandId                *types.URID                                 `gorm:"column:brand_id" json:"brandId,omitempty" example:"4PPU3UXUINHVBHWNYDRNEGWIYI"`
	BrandIdDetail          *Brands                                     `gorm:"foreignkey:BrandId" json:"brandDetail,omitempty"`
	CategoryId             *types.URID                                 `gorm:"column:category_id" json:"categoryId,omitempty" example:"YYG3N7UDKJAQNKQX7236SEU6HQ"`
	CategoryIdDetail       *Categories                                 `gorm:"foreignkey:CategoryId" json:"categoryDetail,omitempty"`
	BasePrice              float64                                     `gorm:"column:base_price" json:"basePrice"`
	Currency               string                                      `gorm:"column:currency" json:"currency" example:"Laborum eius."`
	WeightKg               *float64                                    `gorm:"column:weight_kg" json:"weightKg,omitempty"`
	DimensionsCm           types.JSON                                  `gorm:"column:dimensions_cm" json:"dimensionsCm"`
	Attributes             types.JSON                                  `gorm:"column:attributes" json:"attributes"`
	Tags                   types.StringArray                           `gorm:"column:tags" json:"tags"`
	IsFeatured             bool                                        `gorm:"column:is_featured" json:"isFeatured" example:"false"`
	CreatedBy              *types.URID                                 `gorm:"column:created_by" json:"createdBy,omitempty" example:"PFNAUJ4XLJE47AWMIB47WA5UH4"`
	CreatedByDetail        *shared_types.IamUsers                      `gorm:"foreignkey:CreatedBy" json:"createdByDetail,omitempty"`
	UpdatedBy              *types.URID                                 `gorm:"column:updated_by" json:"updatedBy,omitempty" example:"2MC7CIZ5ERBMVFPD4KSACU7K5U"`
	UpdatedByDetail        *shared_types.IamUsers                      `gorm:"foreignkey:UpdatedBy" json:"updatedByDetail,omitempty"`
	SearchVector           *string                                     `gorm:"column:search_vector" json:"searchVector,omitempty" example:"Quia magnam."`
	CreatedAt              types.NullTime                              `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt              types.NullTime                              `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt              types.NullTime                              `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
	ProductVariantsList    *[]ProductVariants                          `gorm:"foreignKey:ProductId;references:Id" json:"productVariantsList,omitempty"`
	ProductMediaList       *[]ProductMedia                             `gorm:"foreignKey:ProductId;references:Id" json:"productMediaList,omitempty"`
	ProductReviewsList     *[]ProductReviews                           `gorm:"foreignKey:ProductId;references:Id" json:"productReviewsList,omitempty"`
	CollectionProductsList *[]CollectionProducts                       `gorm:"foreignKey:ProductId;references:Id" json:"collectionProductsList,omitempty"`
	ProductTagsList        *[]ProductTags                              `gorm:"foreignKey:ProductId;references:Id" json:"productTagsList,omitempty"`
	PriceHistoryList       *[]PriceHistory                             `gorm:"foreignKey:ProductId;references:Id" json:"priceHistoryList,omitempty"`
	OrderItemsList         *[]shared_types.OrdersOrderItems            `gorm:"foreignKey:ProductId;references:Id" json:"orderItemsList,omitempty"`
	CartItemsList          *[]shared_types.OrdersCartItems             `gorm:"foreignKey:ProductId;references:Id" json:"cartItemsList,omitempty"`
	InventoryList          *[]shared_types.LogisticsInventory          `gorm:"foreignKey:ProductId;references:Id" json:"inventoryList,omitempty"`
	PurchaseOrderItemsList *[]shared_types.LogisticsPurchaseOrderItems `gorm:"foreignKey:ProductId;references:Id" json:"purchaseOrderItemsList,omitempty"`
}

func (p *Products) TableName() string {
	return "catalog.products"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ProductsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Products `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ProductsEdit struct {
	Name             *string               `gorm:"column:name;not null" json:"name" example:"Accusamus ut."`
	Slug             *string               `gorm:"column:slug;not null" json:"slug" example:"Nihil enim."`
	Sku              *string               `gorm:"column:sku;not null" json:"sku" example:"Ea aperiam."`
	Description      *string               `gorm:"column:description" json:"description" example:"Deleniti sequi."`
	ShortDescription *string               `gorm:"column:short_description" json:"shortDescription" example:"Aperiam qui."`
	Status           *CatalogProductStatus `gorm:"column:status" json:"status"`
	BrandId          *types.URID           `gorm:"column:brand_id" json:"brandId,omitempty" example:"3WI4URLEVZBTDJGOXTRVSWV46Y"`
	CategoryId       *types.URID           `gorm:"column:category_id" json:"categoryId,omitempty" example:"DWWQIJQ47BCY3LFFSOHIAXK45E"`
	BasePrice        *float64              `gorm:"column:base_price" json:"basePrice"`
	Currency         *string               `gorm:"column:currency" json:"currency" example:"Quia omnis."`
	WeightKg         *float64              `gorm:"column:weight_kg" json:"weightKg,omitempty"`
	DimensionsCm     *types.JSON           `gorm:"column:dimensions_cm" json:"dimensionsCm"`
	Attributes       *types.JSON           `gorm:"column:attributes" json:"attributes"`
	Tags             *types.StringArray    `gorm:"column:tags" json:"tags"`
	IsFeatured       *bool                 `gorm:"column:is_featured" json:"isFeatured" example:"true"`
	CreatedBy        *types.URID           `gorm:"column:created_by" json:"createdBy,omitempty" example:"LBWWUDH6ONAJZNP55IQMK6ALDY"`
	UpdatedBy        *types.URID           `gorm:"column:updated_by" json:"updatedBy,omitempty" example:"VLA5UPGRXFGELAB7YBUAEQISAM"`
	SearchVector     *string               `gorm:"column:search_vector" json:"searchVector,omitempty" example:"Voluptatem veniam."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ProductsFilter struct {
	Id               *types.URID           `gorm:"column:id;primary_key" json:"id" example:"O5D2COK7KZAO7GMJ5T6TJWVINE"`
	Name             *string               `gorm:"column:name;not null" json:"name" example:"Sint molestiae."`
	Slug             *string               `gorm:"column:slug;not null" json:"slug" example:"Aut ut."`
	Sku              *string               `gorm:"column:sku;not null" json:"sku" example:"Unde labore."`
	Description      *string               `gorm:"column:description" json:"description" example:"Quae nesciunt."`
	ShortDescription *string               `gorm:"column:short_description" json:"shortDescription" example:"Est et."`
	Status           *CatalogProductStatus `gorm:"column:status" json:"status"`
	BrandId          *types.URID           `gorm:"column:brand_id" json:"brandId,omitempty" example:"TVKIJTNY7NHW7DWY5ZX4DCR5ZM"`
	CategoryId       *types.URID           `gorm:"column:category_id" json:"categoryId,omitempty" example:"INQSMIIMKND63NHINKUJTATSZU"`
	BasePrice        *float64              `gorm:"column:base_price" json:"basePrice"`
	Currency         *string               `gorm:"column:currency" json:"currency" example:"Assumenda vero."`
	WeightKg         *float64              `gorm:"column:weight_kg" json:"weightKg,omitempty"`
	DimensionsCm     *types.JSON           `gorm:"column:dimensions_cm" json:"dimensionsCm"`
	Attributes       *types.JSON           `gorm:"column:attributes" json:"attributes"`
	Tags             *types.StringArray    `gorm:"column:tags" json:"tags"`
	IsFeatured       *bool                 `gorm:"column:is_featured" json:"isFeatured" example:"false"`
	CreatedBy        *types.URID           `gorm:"column:created_by" json:"createdBy,omitempty" example:"WD2UHWFP7ZDN5GT6HGG6PP2FFE"`
	UpdatedBy        *types.URID           `gorm:"column:updated_by" json:"updatedBy,omitempty" example:"5BW2G5QAVJB6RGV5GBUMH75TOY"`
	SearchVector     *string               `gorm:"column:search_vector" json:"searchVector,omitempty" example:"Sed asperiores."`
	CreatedAt        *types.NullTime       `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        *types.NullTime       `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt        *types.NullTime       `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}

func (p *ProductsFilter) TableName() string {
	return "catalog.products"
}

// --- Batch Update Struct ---
type ProductsBatchUpdate struct {
	Data       ProductsEdit     `json:"data"`
	PathParams ProductsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ProductsIdentity struct {
	Id types.URID `json:"id"`
}
