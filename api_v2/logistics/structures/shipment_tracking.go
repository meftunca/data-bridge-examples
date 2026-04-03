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
// • ShipmentTrackingForm     - Data input validation and creation
// • ShipmentTracking        - Main database model with relationships
// • ShipmentTrackingEdit    - Partial update operations
// • ShipmentTrackingIdentity - Bulk operation identifiers
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
package logistics_api_structure

import (
	paginationRuntime "backend-generator/apiv2/pagination"

	"github.com/maple-tech/baseline/types"
)

// ShipmentTrackingForm handles data input validation and creation operations.
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
type ShipmentTrackingForm struct {
	Name        string         `gorm:"column:name" json:"name" example:"Fugit fugiat."`
	ShipmentId  types.URID     `gorm:"column:shipment_id;not null" json:"shipmentId" example:"FLZX5E7UKNF4RHBW2FJDNICBCU"`
	Status      string         `gorm:"column:status" json:"status" example:"Veniam consequuntur."`
	Location    string         `gorm:"column:location" json:"location" example:"A excepturi."`
	Description string         `gorm:"column:description" json:"description" example:"Alias minima."`
	EventTime   types.NullTime `gorm:"column:event_time" json:"eventTime"`
}

func (p *ShipmentTrackingForm) TableName() string {
	return "logistics.shipment_tracking"
}

// ShipmentTracking represents the main database model for logistics.shipment_tracking table.
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
// • Table: logistics.shipment_tracking
// • Type: BASE TABLE
// • Schema: logistics
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
type ShipmentTracking struct {
	Id               types.URID     `gorm:"column:id;primary_key" json:"id" example:"BJ6QYUXF75DV5EWI4I3H3ZUT2I"`
	Name             string         `gorm:"column:name" json:"name" example:"Quas fugit."`
	ShipmentId       types.URID     `gorm:"column:shipment_id;not null" json:"shipmentId" example:"N33LWJHDQREBXCBUFZZKXGON34"`
	ShipmentIdDetail *Shipments     `gorm:"foreignkey:ShipmentId" json:"shipmentDetail,omitempty"`
	Status           string         `gorm:"column:status" json:"status" example:"Voluptatibus veritatis."`
	Location         string         `gorm:"column:location" json:"location" example:"Commodi quia."`
	Description      string         `gorm:"column:description" json:"description" example:"Pariatur eaque."`
	EventTime        types.NullTime `gorm:"column:event_time" json:"eventTime"`
	CreatedAt        types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *ShipmentTracking) TableName() string {
	return "logistics.shipment_tracking"
}

// --- Sayfalama (Pagination) için Yardımcı Struct ---
type ShipmentTrackingPage struct {
	paginationRuntime.DefaultPageResponse
	Items []ShipmentTracking `json:"items"`
}

// --- Edit Struct (Güncelleme için - tüm alanlar pointer) ---
type ShipmentTrackingEdit struct {
	Name        *string         `gorm:"column:name" json:"name" example:"Qui asperiores."`
	ShipmentId  *types.URID     `gorm:"column:shipment_id;not null" json:"shipmentId" example:"PKJBIVPEPZHL7HRQDHPBRTNG7M"`
	Status      *string         `gorm:"column:status" json:"status" example:"Nihil alias."`
	Location    *string         `gorm:"column:location" json:"location" example:"Nihil consequatur."`
	Description *string         `gorm:"column:description" json:"description" example:"Atque tempora."`
	EventTime   *types.NullTime `gorm:"column:event_time" json:"eventTime"`
}

// --- Filter Struct (Filtreleme için - tüm alanlar pointer) ---
type ShipmentTrackingFilter struct {
	Id          *types.URID     `gorm:"column:id;primary_key" json:"id" example:"LPJRXJUQRVFMRP5ZGD5KQ5JBEQ"`
	Name        *string         `gorm:"column:name" json:"name" example:"Possimus et."`
	ShipmentId  *types.URID     `gorm:"column:shipment_id;not null" json:"shipmentId" example:"MEJ5M5MVPJFV5I3JKFDYZLTYRI"`
	Status      *string         `gorm:"column:status" json:"status" example:"Atque facere."`
	Location    *string         `gorm:"column:location" json:"location" example:"Unde sed."`
	Description *string         `gorm:"column:description" json:"description" example:"Neque atque."`
	EventTime   *types.NullTime `gorm:"column:event_time" json:"eventTime"`
	CreatedAt   *types.NullTime `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (p *ShipmentTrackingFilter) TableName() string {
	return "logistics.shipment_tracking"
}

// --- Batch Update Struct ---
type ShipmentTrackingBatchUpdate struct {
	Data       ShipmentTrackingEdit     `json:"data"`
	PathParams ShipmentTrackingIdentity `json:"pathParams"`
}

// --- Kimlik Struct'ı (Toplu İşlemler için) ---
type ShipmentTrackingIdentity struct {
	Id types.URID `json:"id"`
}
