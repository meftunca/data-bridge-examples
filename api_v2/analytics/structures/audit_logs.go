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
// • AuditLogsForm     - Data input validation and creation
// • AuditLogs        - Main database model with relationships
// • AuditLogsEdit    - Partial update operations
// • AuditLogsIdentity - Bulk operation identifiers
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
package analytics_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"
	shared_types "data-bridge-examples/api_v2/shared/types"
	"net"

	"github.com/maple-tech/baseline/types"
)

// AuditLogsForm handles data input validation and creation operations.
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
type AuditLogsForm struct {
	Name         string                 `gorm:"column:name" json:"name" example:"Rerum et."`
	UserId       *types.URID            `gorm:"column:user_id" json:"userId,omitempty" example:"C6XVJY6HPFDKHJO464RRTLCEZA"`
	Action       string                 `gorm:"column:action" json:"action" example:"Nesciunt deserunt."`
	ResourceType string                 `gorm:"column:resource_type" json:"resourceType" example:"Praesentium culpa."`
	ResourceId   string                 `gorm:"column:resource_id" json:"resourceId" example:"Sunt quod."`
	Severity     AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	IpAddress    *net.IP                `gorm:"column:ip_address" json:"ipAddress,omitempty"`
	UserAgent    string                 `gorm:"column:user_agent" json:"userAgent" example:"Non molestiae."`
	OldValues    types.JSON             `gorm:"column:old_values" json:"oldValues,omitempty"`
	NewValues    types.JSON             `gorm:"column:new_values" json:"newValues,omitempty"`
	Metadata     types.JSON             `gorm:"column:metadata" json:"metadata"`
}

func (p *AuditLogsForm) TableName() string {
	return "analytics.audit_logs"
}

// AuditLogs represents the main database model for analytics.audit_logs table.
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
// • Table: analytics.audit_logs
// • Type: BASE TABLE
// • Schema: analytics
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
type AuditLogs struct {
	Id           types.URID             `gorm:"column:id;primary_key" json:"id" example:"TMN7UBTCPFDZLL5E3U7ET3ZTGA"`
	Name         string                 `gorm:"column:name" json:"name" example:"Aliquid voluptas."`
	UserId       *types.URID            `gorm:"column:user_id" json:"userId,omitempty" example:"DXII73YHCFDSFAZCQWL2RZNIHM"`
	UserIdDetail *shared_types.IamUsers `gorm:"foreignkey:UserId" json:"userDetail,omitempty"`
	Action       string                 `gorm:"column:action" json:"action" example:"Accusamus eum."`
	ResourceType string                 `gorm:"column:resource_type" json:"resourceType" example:"Molestiae vel."`
	ResourceId   string                 `gorm:"column:resource_id" json:"resourceId" example:"Dolores consectetur."`
	Severity     AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	IpAddress    *net.IP                `gorm:"column:ip_address" json:"ipAddress,omitempty"`
	UserAgent    string                 `gorm:"column:user_agent" json:"userAgent" example:"Ex accusantium."`
	OldValues    types.JSON             `gorm:"column:old_values" json:"oldValues,omitempty"`
	NewValues    types.JSON             `gorm:"column:new_values" json:"newValues,omitempty"`
	Metadata     types.JSON             `gorm:"column:metadata" json:"metadata"`
	CreatedAt    types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *AuditLogs) TableName() string {
	return "analytics.audit_logs"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type AuditLogsPage struct {
	paginationRuntime.DefaultPageResponse
	Items []AuditLogs `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type AuditLogsEdit struct {
	Name         *string                 `gorm:"column:name" json:"name" example:"Assumenda quasi."`
	UserId       *types.URID             `gorm:"column:user_id" json:"userId,omitempty" example:"5WINP7GM3NEK7NDASOSGGIA6HM"`
	Action       *string                 `gorm:"column:action" json:"action" example:"Corporis qui."`
	ResourceType *string                 `gorm:"column:resource_type" json:"resourceType" example:"Dolorum et."`
	ResourceId   *string                 `gorm:"column:resource_id" json:"resourceId" example:"Et eaque."`
	Severity     *AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	IpAddress    *net.IP                 `gorm:"column:ip_address" json:"ipAddress,omitempty"`
	UserAgent    *string                 `gorm:"column:user_agent" json:"userAgent" example:"Aliquid ut."`
	OldValues    *types.JSON             `gorm:"column:old_values" json:"oldValues,omitempty"`
	NewValues    *types.JSON             `gorm:"column:new_values" json:"newValues,omitempty"`
	Metadata     *types.JSON             `gorm:"column:metadata" json:"metadata"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type AuditLogsFilter struct {
	Id           *types.URID             `gorm:"column:id;primary_key" json:"id" example:"YYUZKAMKOREZVMXGWR5MXUNNTQ"`
	Name         *string                 `gorm:"column:name" json:"name" example:"Tenetur dolore."`
	UserId       *types.URID             `gorm:"column:user_id" json:"userId,omitempty" example:"A6TGFWA36VD3XBBMEJX3BYNKAA"`
	Action       *string                 `gorm:"column:action" json:"action" example:"Et pariatur."`
	ResourceType *string                 `gorm:"column:resource_type" json:"resourceType" example:"Maiores iusto."`
	ResourceId   *string                 `gorm:"column:resource_id" json:"resourceId" example:"Omnis nihil."`
	Severity     *AnalyticsEventSeverity `gorm:"column:severity" json:"severity"`
	IpAddress    *net.IP                 `gorm:"column:ip_address" json:"ipAddress,omitempty"`
	UserAgent    *string                 `gorm:"column:user_agent" json:"userAgent" example:"Harum similique."`
	OldValues    *types.JSON             `gorm:"column:old_values" json:"oldValues,omitempty"`
	NewValues    *types.JSON             `gorm:"column:new_values" json:"newValues,omitempty"`
	Metadata     *types.JSON             `gorm:"column:metadata" json:"metadata"`
	CreatedAt    *types.NullTime         `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *AuditLogsFilter) TableName() string {
	return "analytics.audit_logs"
}

// --- Batch Update Struct ---
type AuditLogsBatchUpdate struct {
	Data       AuditLogsEdit     `json:"data"`
	PathParams AuditLogsIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type AuditLogsIdentity struct {
	Id types.URID `json:"id"`
}
