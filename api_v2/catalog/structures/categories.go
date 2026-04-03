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
// • CategoriesForm     - Data input validation and creation
// • Categories        - Main database model with relationships
// • CategoriesEdit    - Partial update operations
// • CategoriesIdentity - Bulk operation identifiers
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

// CategoriesForm handles data input validation and creation operations.
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
type CategoriesForm struct {
	Name        string      `gorm:"column:name;not null" json:"name" example:"Voluptas et."`
	Slug        string      `gorm:"column:slug;not null" json:"slug" example:"Veniam occaecati."`
	Description string      `gorm:"column:description" json:"description" example:"Corporis numquam."`
	ParentId    *types.URID `gorm:"column:parent_id" json:"parentId,omitempty" example:"OFNHCQSEMZFXDD3XAGSCZZKHSY"`
	Icon        string      `gorm:"column:icon" json:"icon" example:"Autem rerum."`
	SortOrder   int         `gorm:"column:sort_order" json:"sortOrder" example:"-4668801334172264041"`
	IsActive    bool        `gorm:"column:is_active" json:"isActive" example:"true"`
}

func (p *CategoriesForm) TableName() string {
	return "catalog.categories"
}

// Categories represents the main database model for catalog.categories table.
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
// • Table: catalog.categories
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
type Categories struct {
	Id             types.URID     `gorm:"column:id;primary_key" json:"id" example:"DLFCYRVDKVD4DJTUH4ZNY4HMAI"`
	Name           string         `gorm:"column:name;not null" json:"name" example:"Possimus fuga."`
	Slug           string         `gorm:"column:slug;not null" json:"slug" example:"Natus ullam."`
	Description    string         `gorm:"column:description" json:"description" example:"Aut maxime."`
	ParentId       *types.URID    `gorm:"column:parent_id" json:"parentId,omitempty" example:"FOY2LL4TIFCKTKQ2FXXQKVG6EM"`
	ParentIdDetail *Categories    `gorm:"foreignkey:ParentId" json:"parentDetail,omitempty"`
	Icon           string         `gorm:"column:icon" json:"icon" example:"Accusantium aperiam."`
	SortOrder      int            `gorm:"column:sort_order" json:"sortOrder" example:"-5176600687688600447"`
	IsActive       bool           `gorm:"column:is_active" json:"isActive" example:"true"`
	CreatedAt      types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	CategoriesList *[]Categories  `gorm:"foreignKey:ParentId;references:Id" json:"categoriesList,omitempty"`
	ProductsList   *[]Products    `gorm:"foreignKey:CategoryId;references:Id" json:"productsList,omitempty"`
}

func (p *Categories) TableName() string {
	return "catalog.categories"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type CategoriesPage struct {
	paginationRuntime.DefaultPageResponse
	Items []Categories `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type CategoriesEdit struct {
	Name        *string     `gorm:"column:name;not null" json:"name" example:"Ut sed."`
	Slug        *string     `gorm:"column:slug;not null" json:"slug" example:"Voluptatem excepturi."`
	Description *string     `gorm:"column:description" json:"description" example:"Commodi alias."`
	ParentId    *types.URID `gorm:"column:parent_id" json:"parentId,omitempty" example:"5APK5LEJJVHLFJTSTPRRHZIBYQ"`
	Icon        *string     `gorm:"column:icon" json:"icon" example:"Omnis ullam."`
	SortOrder   *int        `gorm:"column:sort_order" json:"sortOrder" example:"1297582316368301057"`
	IsActive    *bool       `gorm:"column:is_active" json:"isActive" example:"true"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type CategoriesFilter struct {
	Id          *types.URID     `gorm:"column:id;primary_key" json:"id" example:"V4Q2GTXQBZE2NHVUWWSGL4KDIQ"`
	Name        *string         `gorm:"column:name;not null" json:"name" example:"Sequi atque."`
	Slug        *string         `gorm:"column:slug;not null" json:"slug" example:"Aut numquam."`
	Description *string         `gorm:"column:description" json:"description" example:"Quaerat facilis."`
	ParentId    *types.URID     `gorm:"column:parent_id" json:"parentId,omitempty" example:"5WINAWZV5NC6NPOKOJOD6D44OA"`
	Icon        *string         `gorm:"column:icon" json:"icon" example:"Fugiat omnis."`
	SortOrder   *int            `gorm:"column:sort_order" json:"sortOrder" example:"876596958032762462"`
	IsActive    *bool           `gorm:"column:is_active" json:"isActive" example:"false"`
	CreatedAt   *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt   *types.NullTime `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *CategoriesFilter) TableName() string {
	return "catalog.categories"
}

// --- Batch Update Struct ---
type CategoriesBatchUpdate struct {
	Data       CategoriesEdit     `json:"data"`
	PathParams CategoriesIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type CategoriesIdentity struct {
	Id types.URID `json:"id"`
}
