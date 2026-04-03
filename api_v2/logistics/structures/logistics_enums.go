package logistics_api_structure

// Bu dosya, veritabanındaki ENUM tiplerinden otomatik olarak üretilmiştir.

// LogisticsShipmentStatus represents the 'shipment_status' enum type from the 'logistics' schema.
type LogisticsShipmentStatus string

const (
	LogisticsShipmentStatusPreparing      LogisticsShipmentStatus = "preparing"
	LogisticsShipmentStatusPickedUp       LogisticsShipmentStatus = "picked_up"
	LogisticsShipmentStatusInTransit      LogisticsShipmentStatus = "in_transit"
	LogisticsShipmentStatusOutForDelivery LogisticsShipmentStatus = "out_for_delivery"
	LogisticsShipmentStatusDelivered      LogisticsShipmentStatus = "delivered"
	LogisticsShipmentStatusReturned       LogisticsShipmentStatus = "returned"
	LogisticsShipmentStatusLost           LogisticsShipmentStatus = "lost"
)

// LogisticsStockLevel represents the 'stock_level' enum type from the 'logistics' schema.
type LogisticsStockLevel string

const (
	LogisticsStockLevelOutOfStock LogisticsStockLevel = "out_of_stock"
	LogisticsStockLevelLow        LogisticsStockLevel = "low"
	LogisticsStockLevelNormal     LogisticsStockLevel = "normal"
	LogisticsStockLevelHigh       LogisticsStockLevel = "high"
	LogisticsStockLevelOverstock  LogisticsStockLevel = "overstock"
)
