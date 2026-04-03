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
// • TeamMembersForm     - Data input validation and creation
// • TeamMembers        - Main database model with relationships
// • TeamMembersEdit    - Partial update operations
// • TeamMembersIdentity - Bulk operation identifiers
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

// TeamMembersForm handles data input validation and creation operations.
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
type TeamMembersForm struct {
	Name     string         `gorm:"column:name" json:"name" example:"Voluptas dignissimos."`
	TeamId   types.URID     `gorm:"column:team_id;not null" json:"teamId" example:"GEUVGMZ22JAIRBDNXDOGXXO4CA"`
	UserId   types.URID     `gorm:"column:user_id;not null" json:"userId" example:"CLQ6U7XML5DTFCMONSSDK2DCRI"`
	Role     string         `gorm:"column:role" json:"role" example:"Et sit."`
	JoinedAt types.NullTime `gorm:"column:joined_at" json:"joinedAt"`
}

func (p *TeamMembersForm) TableName() string {
	return "iam.team_members"
}

// TeamMembers represents the main database model for iam.team_members table.
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
// • Table: iam.team_members
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
type TeamMembers struct {
	Id           types.URID     `gorm:"column:id;primary_key" json:"id" example:"KBK6MA5HBVBTLFNB6O5FDSP3CQ"`
	Name         string         `gorm:"column:name" json:"name" example:"Dolores modi."`
	TeamId       types.URID     `gorm:"column:team_id;not null" json:"teamId" example:"N52AN2L3ZVEJDOXZEQ4LQKF4ZI"`
	TeamIdDetail *Teams         `gorm:"foreignkey:TeamId" json:"teamDetail,omitempty"`
	UserId       types.URID     `gorm:"column:user_id;not null" json:"userId" example:"DORFXJGZI5BQVOA3XYIPZEXJGM"`
	UserIdDetail *Users         `gorm:"foreignkey:UserId" json:"userDetail,omitempty"`
	Role         string         `gorm:"column:role" json:"role" example:"Nemo dignissimos."`
	JoinedAt     types.NullTime `gorm:"column:joined_at" json:"joinedAt"`
	CreatedAt    types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *TeamMembers) TableName() string {
	return "iam.team_members"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type TeamMembersPage struct {
	paginationRuntime.DefaultPageResponse
	Items []TeamMembers `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type TeamMembersEdit struct {
	Name     *string         `gorm:"column:name" json:"name" example:"Dolorem aperiam."`
	TeamId   *types.URID     `gorm:"column:team_id;not null" json:"teamId" example:"LK6CXOIU45E5LJR7PQ64PAUKUQ"`
	UserId   *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"I34RUT3OVBEGRFKTYJXMFBUJ4A"`
	Role     *string         `gorm:"column:role" json:"role" example:"Quam praesentium."`
	JoinedAt *types.NullTime `gorm:"column:joined_at" json:"joinedAt"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type TeamMembersFilter struct {
	Id        *types.URID     `gorm:"column:id;primary_key" json:"id" example:"BXGEJZROPVFZZCQBOG36LFOGGI"`
	Name      *string         `gorm:"column:name" json:"name" example:"Dolores officiis."`
	TeamId    *types.URID     `gorm:"column:team_id;not null" json:"teamId" example:"MEKOG5M6SFCB7H5TRBM7C7NFGE"`
	UserId    *types.URID     `gorm:"column:user_id;not null" json:"userId" example:"EK74BIGDBRDTBNZ7GVRZ2GMSTU"`
	Role      *string         `gorm:"column:role" json:"role" example:"Consequatur quam."`
	JoinedAt  *types.NullTime `gorm:"column:joined_at" json:"joinedAt"`
	CreatedAt *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *TeamMembersFilter) TableName() string {
	return "iam.team_members"
}

// --- Batch Update Struct ---
type TeamMembersBatchUpdate struct {
	Data       TeamMembersEdit     `json:"data"`
	PathParams TeamMembersIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type TeamMembersIdentity struct {
	Id types.URID `json:"id"`
}
