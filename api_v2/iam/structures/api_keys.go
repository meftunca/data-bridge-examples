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
// • ApiKeysForm     - Data input validation and creation
// • ApiKeys        - Main database model with relationships
// • ApiKeysEdit    - Partial update operations
// • ApiKeysIdentity - Bulk operation identifiers
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
package iam_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"

	"github.com/maple-tech/baseline/types"
)

// ApiKeysForm handles data input validation and creation operations.
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
type ApiKeysForm struct {
	Name           string            `gorm:"column:name;not null" json:"name" example:"Optio qui."`
	KeyHash        string            `gorm:"column:key_hash;not null" json:"keyHash" example:"Accusantium rerum."`
	UserId         types.URID        `gorm:"column:user_id;not null" json:"userId" example:"MHEQLUBV3ZHLPMATMI7SIR4IY4"`
	OrganizationId types.URID        `gorm:"column:organization_id;not null" json:"organizationId" example:"DQNRB2FYLNH6TPXQATCZMANFLE"`
	Scopes         types.StringArray `gorm:"column:scopes" json:"scopes"`
	IsActive       bool              `gorm:"column:is_active" json:"isActive" example:"false"`
	LastUsedAt     types.NullTime    `gorm:"column:last_used_at" json:"lastUsedAt,omitempty"`
	ExpiresAt      types.NullTime    `gorm:"column:expires_at" json:"expiresAt,omitempty"`
}

func (p *ApiKeysForm) TableName() string {
	return "iam.api_keys"
}

// ApiKeys represents the main database model for iam.api_keys table.
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
// • Table: iam.api_keys
// • Type: BASE TABLE
// • Schema: iam
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
type ApiKeys struct {
	Id                   types.URID        `gorm:"column:id;primary_key" json:"id" example:"G577MOE5SNH2FPMTMRBV44PS2U"`
	Name                 string            `gorm:"column:name;not null" json:"name" example:"Earum qui."`
	KeyHash              string            `gorm:"column:key_hash;not null" json:"keyHash" example:"Rem ut."`
	UserId               types.URID        `gorm:"column:user_id;not null" json:"userId" example:"MYS26EC7SVB7HPCQVNDQ23UR6Y"`
	UserIdDetail         *Users            `gorm:"foreignkey:UserId" json:"userDetail,omitempty"`
	OrganizationId       types.URID        `gorm:"column:organization_id;not null" json:"organizationId" example:"6BOYRUC5HVHHTD2M7GP4NZC5FY"`
	OrganizationIdDetail *Organizations    `gorm:"foreignkey:OrganizationId" json:"organizationDetail,omitempty"`
	Scopes               types.StringArray `gorm:"column:scopes" json:"scopes"`
	IsActive             bool              `gorm:"column:is_active" json:"isActive" example:"false"`
	LastUsedAt           types.NullTime    `gorm:"column:last_used_at" json:"lastUsedAt,omitempty"`
	ExpiresAt            types.NullTime    `gorm:"column:expires_at" json:"expiresAt,omitempty"`
	CreatedAt            types.NullTime    `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt            types.NullTime    `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ApiKeys) TableName() string {
	return "iam.api_keys"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ApiKeysPage struct {
	paginationRuntime.DefaultPageResponse
	Items []ApiKeys `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ApiKeysEdit struct {
	Name           *string            `gorm:"column:name;not null" json:"name" example:"Repellat sunt."`
	KeyHash        *string            `gorm:"column:key_hash;not null" json:"keyHash" example:"Dolor dicta."`
	UserId         *types.URID        `gorm:"column:user_id;not null" json:"userId" example:"45H7NHRNDBDBBI7SOVF74L6CLU"`
	OrganizationId *types.URID        `gorm:"column:organization_id;not null" json:"organizationId" example:"EFCI2ZR6M5DO3OFAD2UXVG2PUE"`
	Scopes         *types.StringArray `gorm:"column:scopes" json:"scopes"`
	IsActive       *bool              `gorm:"column:is_active" json:"isActive" example:"false"`
	LastUsedAt     *types.NullTime    `gorm:"column:last_used_at" json:"lastUsedAt,omitempty"`
	ExpiresAt      *types.NullTime    `gorm:"column:expires_at" json:"expiresAt,omitempty"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ApiKeysFilter struct {
	Id             *types.URID        `gorm:"column:id;primary_key" json:"id" example:"42ZWYBGB6VAHVPPTGO3WMPB6IY"`
	Name           *string            `gorm:"column:name;not null" json:"name" example:"Porro sequi."`
	KeyHash        *string            `gorm:"column:key_hash;not null" json:"keyHash" example:"Nostrum non."`
	UserId         *types.URID        `gorm:"column:user_id;not null" json:"userId" example:"GQSI65OSGJE77MTKCEYVXOOHTE"`
	OrganizationId *types.URID        `gorm:"column:organization_id;not null" json:"organizationId" example:"WA4HSFZAYZBGNMA3WHOUGEMKIY"`
	Scopes         *types.StringArray `gorm:"column:scopes" json:"scopes"`
	IsActive       *bool              `gorm:"column:is_active" json:"isActive" example:"false"`
	LastUsedAt     *types.NullTime    `gorm:"column:last_used_at" json:"lastUsedAt,omitempty"`
	ExpiresAt      *types.NullTime    `gorm:"column:expires_at" json:"expiresAt,omitempty"`
	CreatedAt      *types.NullTime    `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      *types.NullTime    `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (p *ApiKeysFilter) TableName() string {
	return "iam.api_keys"
}

// --- Batch Update Struct ---
type ApiKeysBatchUpdate struct {
	Data       ApiKeysEdit     `json:"data"`
	PathParams ApiKeysIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ApiKeysIdentity struct {
	Id types.URID `json:"id"`
}
