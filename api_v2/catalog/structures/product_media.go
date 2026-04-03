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
// • ProductMediaForm     - Data input validation and creation
// • ProductMedia        - Main database model with relationships
// • ProductMediaEdit    - Partial update operations
// • ProductMediaIdentity - Bulk operation identifiers
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

	"github.com/maple-tech/baseline/types"
)

// ProductMediaForm handles data input validation and creation operations.
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
type ProductMediaForm struct {
	Name      string           `gorm:"column:name" json:"name" example:"Vero laudantium."`
	ProductId types.URID       `gorm:"column:product_id;not null" json:"productId" example:"UIJLEWWKS5DTRABWHQJUNBS5IE"`
	MediaType CatalogMediaType `gorm:"column:media_type" json:"mediaType"`
	Url       string           `gorm:"column:url;not null" json:"url" example:"Molestias vitae."`
	AltText   string           `gorm:"column:alt_text" json:"altText" example:"Perferendis rem."`
	SortOrder int              `gorm:"column:sort_order" json:"sortOrder" example:"7732681043845844854"`
	IsPrimary bool             `gorm:"column:is_primary" json:"isPrimary" example:"false"`
	Metadata  types.JSON       `gorm:"column:metadata" json:"metadata"`
}

func (p *ProductMediaForm) TableName() string {
	return "catalog.product_media"
}

// ProductMedia represents the main database model for catalog.product_media table.
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
// • Table: catalog.product_media
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
type ProductMedia struct {
	Id              types.URID       `gorm:"column:id;primary_key" json:"id" example:"4MNWGKYKHZHCTBELYWTG4YZBGM"`
	Name            string           `gorm:"column:name" json:"name" example:"Sed aut."`
	ProductId       types.URID       `gorm:"column:product_id;not null" json:"productId" example:"6TKDMHSG6FAD3ODUMVVLSUXKC4"`
	ProductIdDetail *Products        `gorm:"foreignkey:ProductId" json:"productDetail,omitempty"`
	MediaType       CatalogMediaType `gorm:"column:media_type" json:"mediaType"`
	Url             string           `gorm:"column:url;not null" json:"url" example:"Et voluptatibus."`
	AltText         string           `gorm:"column:alt_text" json:"altText" example:"Nobis fugit."`
	SortOrder       int              `gorm:"column:sort_order" json:"sortOrder" example:"-4615935244453350065"`
	IsPrimary       bool             `gorm:"column:is_primary" json:"isPrimary" example:"false"`
	Metadata        types.JSON       `gorm:"column:metadata" json:"metadata"`
	CreatedAt       types.NullTime   `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       types.NullTime   `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ProductMedia) TableName() string {
	return "catalog.product_media"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ProductMediaPage struct {
	paginationRuntime.DefaultPageResponse
	Items []ProductMedia `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ProductMediaEdit struct {
	Name      *string           `gorm:"column:name" json:"name" example:"Culpa qui."`
	ProductId *types.URID       `gorm:"column:product_id;not null" json:"productId" example:"U6HEU5WXNZGQXD5LMM4YEJHALU"`
	MediaType *CatalogMediaType `gorm:"column:media_type" json:"mediaType"`
	Url       *string           `gorm:"column:url;not null" json:"url" example:"Nam officia."`
	AltText   *string           `gorm:"column:alt_text" json:"altText" example:"Doloremque sint."`
	SortOrder *int              `gorm:"column:sort_order" json:"sortOrder" example:"6167842641739538854"`
	IsPrimary *bool             `gorm:"column:is_primary" json:"isPrimary" example:"false"`
	Metadata  *types.JSON       `gorm:"column:metadata" json:"metadata"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ProductMediaFilter struct {
	Id        *types.URID       `gorm:"column:id;primary_key" json:"id" example:"K7TNRTWVJ5ERDG7LRDQN472VIU"`
	Name      *string           `gorm:"column:name" json:"name" example:"Ut quo."`
	ProductId *types.URID       `gorm:"column:product_id;not null" json:"productId" example:"YIDWWNU7AJDIRGRZ5AGH7RDR4U"`
	MediaType *CatalogMediaType `gorm:"column:media_type" json:"mediaType"`
	Url       *string           `gorm:"column:url;not null" json:"url" example:"Qui cupiditate."`
	AltText   *string           `gorm:"column:alt_text" json:"altText" example:"Est nisi."`
	SortOrder *int              `gorm:"column:sort_order" json:"sortOrder" example:"4430970398398264411"`
	IsPrimary *bool             `gorm:"column:is_primary" json:"isPrimary" example:"true"`
	Metadata  *types.JSON       `gorm:"column:metadata" json:"metadata"`
	CreatedAt *types.NullTime   `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt *types.NullTime   `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ProductMediaFilter) TableName() string {
	return "catalog.product_media"
}

// --- Batch Update Struct ---
type ProductMediaBatchUpdate struct {
	Data       ProductMediaEdit     `json:"data"`
	PathParams ProductMediaIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ProductMediaIdentity struct {
	Id types.URID `json:"id"`
}
