package catalog_api_structure

// Bu dosya, veritabanındaki ENUM tiplerinden otomatik olarak üretilmiştir.

// CatalogMediaType represents the 'media_type' enum type from the 'catalog' schema.
type CatalogMediaType string

const (
	CatalogMediaTypeImage    CatalogMediaType = "image"
	CatalogMediaTypeVideo    CatalogMediaType = "video"
	CatalogMediaTypeDocument CatalogMediaType = "document"
	CatalogMediaTypeAudio    CatalogMediaType = "audio"
	CatalogMediaType3DModel  CatalogMediaType = "3d_model"
)

// CatalogProductStatus represents the 'product_status' enum type from the 'catalog' schema.
type CatalogProductStatus string

const (
	CatalogProductStatusDraft        CatalogProductStatus = "draft"
	CatalogProductStatusActive       CatalogProductStatus = "active"
	CatalogProductStatusDiscontinued CatalogProductStatus = "discontinued"
	CatalogProductStatusArchived     CatalogProductStatus = "archived"
)
