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
// • TagsForm     - Data input validation and creation
// • Tags        - Main database model with relationships
// • TagsEdit    - Partial update operations
// • TagsIdentity - Bulk operation identifiers
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

// TagsForm handles data input validation and creation operations.
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
type TagsForm struct {
	Name string `gorm:"column:name;not null" json:"name" example:"Voluptatum corrupti."`
	Slug string `gorm:"column:slug;not null" json:"slug" example:"Itaque delectus."`
}

func (p *TagsForm) TableName() string {
	return "catalog.tags"
}

// Tags represents the main database model for catalog.tags table.
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
// • Table: catalog.tags
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
type Tags struct {
	Id              types.URID     `gorm:"column:id;primary_key" json:"id" example:"6ZMNP2UX7ZDPXGES2X5TIK7C3I"`
	Name            string         `gorm:"column:name;not null" json:"name" example:"Voluptatibus et."`
	Slug            string         `gorm:"column:slug;not null" json:"slug" example:"Dolor ratione."`
	CreatedAt       types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	ProductTagsList *[]ProductTags `gorm:"foreignKey:TagId;references:Id" json:"productTagsList,omitempty"`
}

func (p *Tags) TableName() string {
	return "catalog.tags"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type TagsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Tags `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type TagsEdit struct {
	Name *string `gorm:"column:name;not null" json:"name" example:"Quo et."`
	Slug *string `gorm:"column:slug;not null" json:"slug" example:"Explicabo error."`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type TagsFilter struct {
	Id        *types.URID     `gorm:"column:id;primary_key" json:"id" example:"WH63PTLZBJEXFPG4TWWFL24JYU"`
	Name      *string         `gorm:"column:name;not null" json:"name" example:"Occaecati sint."`
	Slug      *string         `gorm:"column:slug;not null" json:"slug" example:"Quaerat dolorem."`
	CreatedAt *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *TagsFilter) TableName() string {
	return "catalog.tags"
}

// --- Batch Update Struct ---
type TagsBatchUpdate struct {
	Data       TagsEdit     `json:"data"`
	PathParams TagsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type TagsIdentity struct {
	Id types.URID `json:"id"`
}
