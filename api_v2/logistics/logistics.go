// Endpoint Registration Template - Enterprise-Grade Route Management
//
// This template generates comprehensive HTTP endpoint registration for generated resources.
// It implements enterprise patterns including:
//
// • RESTful API conventions with consistent URL patterns
// • Dependency injection for controllers and services
// • Event-driven architecture integration for request monitoring
// • Route grouping for logical resource organization
// • Comprehensive CRUD operations with bulk processing support
// • Type-safe parameter binding and validation
// • RBAC endpoint-based permission middleware (when RBACEnabled=true)
// • SSE real-time subscription endpoint (when SSEEnabled=true)
//
// Generated Route Patterns:
// • GET    /{resource}/pagination     - Paginated listing with filtering
// • GET    /{resource}/with-id/{id}   - Single resource retrieval
// • POST   /{resource}                - Create new resource
// • POST   /{resource}/bulk           - Bulk resource creation
// • PUT    /{resource}/with-id/{id}   - Update existing resource
// • PUT    /{resource}/bulk           - Bulk resource updates
// • DELETE /{resource}/with-id/{id}   - Delete single resource
// • DELETE /{resource}/bulk           - Bulk resource deletion
// • GET    /{resource}/events         - SSE real-time subscription (SSEEnabled=true)
//
// Permission Key Format (RBACEnabled=true):
//
//	{schema}.{resource}.{action}
//	e.g.: core.users.read, core.users.write, core.users.delete
package logistics

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Dinamik olarak ilgili paketleri import et
	controller "data-bridge-examples/api_v2/logistics/controllers"
	services "data-bridge-examples/api_v2/logistics/services"

	"github.com/maple-tech/baseline/events" // EventManager için

	"backend-generator/apiv2/sse"
)

// Run initializes and registers all HTTP endpoints for the generated API schema.
//
// This function sets up the complete routing infrastructure with:
// • Dependency injection for all controllers and services
// • Event manager integration for request lifecycle monitoring
// • Database connection propagation to service layer
// • Route group organization for clean URL structure
// • RESTful endpoint registration following HTTP conventions
//
// The registration process creates route groups for each resource with consistent
// patterns and comprehensive CRUD operations. All endpoints include proper error
// handling, validation, and observability through the event system.
//
// Parameters:
//   - app: Fiber router instance for route registration
//   - db: GORM database connection for service layer operations
//   - eventManager: Event manager for request monitoring and observability
//
// Generated Structure:
//
//	Each resource gets its own route group with standardized endpoints
//	supporting both single and bulk operations for optimal performance.
func Run(app fiber.Router, db *gorm.DB, eventManager *events.EventManager) {

	// --- inventory için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		InventoryController := &controller.InventoryController{
			Svc: &services.InventoryService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		InventoryGroup := app.Group("/inventory")

		// Rotaları tanımla
		InventoryGroup.Get("/pagination", InventoryController.GetInventoryWithPagination)

		InventoryGroup.Get("/with-id/:id", InventoryController.GetInventoryById)
		InventoryGroup.Post("/", InventoryController.CreateInventory)
		InventoryGroup.Post("/bulk", InventoryController.CreateInventoryMultiple)
		InventoryGroup.Put("/with-id/:id", InventoryController.UpdateInventory)
		InventoryGroup.Delete("/with-id/:id", InventoryController.DeleteInventory)
		InventoryGroup.Delete("/bulk", InventoryController.DeleteInventoryMultiple)
		InventoryGroup.Put("/bulk", InventoryController.UpdateInventoryMultiple)

		InventoryGroup.Get("/events", sse.StreamResource("logistics.inventory"))
	}

	// --- purchase_order_items için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		PurchaseOrderItemsController := &controller.PurchaseOrderItemsController{
			Svc: &services.PurchaseOrderItemsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		PurchaseOrderItemsGroup := app.Group("/purchase-order-items")

		// Rotaları tanımla
		PurchaseOrderItemsGroup.Get("/pagination", PurchaseOrderItemsController.GetPurchaseOrderItemsWithPagination)

		PurchaseOrderItemsGroup.Get("/with-id/:id", PurchaseOrderItemsController.GetPurchaseOrderItemsById)
		PurchaseOrderItemsGroup.Post("/", PurchaseOrderItemsController.CreatePurchaseOrderItems)
		PurchaseOrderItemsGroup.Post("/bulk", PurchaseOrderItemsController.CreatePurchaseOrderItemsMultiple)
		PurchaseOrderItemsGroup.Put("/with-id/:id", PurchaseOrderItemsController.UpdatePurchaseOrderItems)
		PurchaseOrderItemsGroup.Delete("/with-id/:id", PurchaseOrderItemsController.DeletePurchaseOrderItems)
		PurchaseOrderItemsGroup.Delete("/bulk", PurchaseOrderItemsController.DeletePurchaseOrderItemsMultiple)
		PurchaseOrderItemsGroup.Put("/bulk", PurchaseOrderItemsController.UpdatePurchaseOrderItemsMultiple)

		PurchaseOrderItemsGroup.Get("/events", sse.StreamResource("logistics.purchase-order-items"))
	}

	// --- purchase_orders için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		PurchaseOrdersController := &controller.PurchaseOrdersController{
			Svc: &services.PurchaseOrdersService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		PurchaseOrdersGroup := app.Group("/purchase-orders")

		// Rotaları tanımla
		PurchaseOrdersGroup.Get("/pagination", PurchaseOrdersController.GetPurchaseOrdersWithPagination)

		PurchaseOrdersGroup.Get("/with-id/:id", PurchaseOrdersController.GetPurchaseOrdersById)
		PurchaseOrdersGroup.Post("/", PurchaseOrdersController.CreatePurchaseOrders)
		PurchaseOrdersGroup.Post("/bulk", PurchaseOrdersController.CreatePurchaseOrdersMultiple)
		PurchaseOrdersGroup.Put("/with-id/:id", PurchaseOrdersController.UpdatePurchaseOrders)
		PurchaseOrdersGroup.Delete("/with-id/:id", PurchaseOrdersController.DeletePurchaseOrders)
		PurchaseOrdersGroup.Delete("/bulk", PurchaseOrdersController.DeletePurchaseOrdersMultiple)
		PurchaseOrdersGroup.Put("/bulk", PurchaseOrdersController.UpdatePurchaseOrdersMultiple)

		PurchaseOrdersGroup.Get("/events", sse.StreamResource("logistics.purchase-orders"))
	}

	// --- shipment_items için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ShipmentItemsController := &controller.ShipmentItemsController{
			Svc: &services.ShipmentItemsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ShipmentItemsGroup := app.Group("/shipment-items")

		// Rotaları tanımla
		ShipmentItemsGroup.Get("/pagination", ShipmentItemsController.GetShipmentItemsWithPagination)

		ShipmentItemsGroup.Get("/with-id/:id", ShipmentItemsController.GetShipmentItemsById)
		ShipmentItemsGroup.Post("/", ShipmentItemsController.CreateShipmentItems)
		ShipmentItemsGroup.Post("/bulk", ShipmentItemsController.CreateShipmentItemsMultiple)
		ShipmentItemsGroup.Put("/with-id/:id", ShipmentItemsController.UpdateShipmentItems)
		ShipmentItemsGroup.Delete("/with-id/:id", ShipmentItemsController.DeleteShipmentItems)
		ShipmentItemsGroup.Delete("/bulk", ShipmentItemsController.DeleteShipmentItemsMultiple)
		ShipmentItemsGroup.Put("/bulk", ShipmentItemsController.UpdateShipmentItemsMultiple)

		ShipmentItemsGroup.Get("/events", sse.StreamResource("logistics.shipment-items"))
	}

	// --- shipment_tracking için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ShipmentTrackingController := &controller.ShipmentTrackingController{
			Svc: &services.ShipmentTrackingService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ShipmentTrackingGroup := app.Group("/shipment-tracking")

		// Rotaları tanımla
		ShipmentTrackingGroup.Get("/pagination", ShipmentTrackingController.GetShipmentTrackingWithPagination)

		ShipmentTrackingGroup.Get("/with-id/:id", ShipmentTrackingController.GetShipmentTrackingById)
		ShipmentTrackingGroup.Post("/", ShipmentTrackingController.CreateShipmentTracking)
		ShipmentTrackingGroup.Post("/bulk", ShipmentTrackingController.CreateShipmentTrackingMultiple)
		ShipmentTrackingGroup.Put("/with-id/:id", ShipmentTrackingController.UpdateShipmentTracking)
		ShipmentTrackingGroup.Delete("/with-id/:id", ShipmentTrackingController.DeleteShipmentTracking)
		ShipmentTrackingGroup.Delete("/bulk", ShipmentTrackingController.DeleteShipmentTrackingMultiple)
		ShipmentTrackingGroup.Put("/bulk", ShipmentTrackingController.UpdateShipmentTrackingMultiple)

		ShipmentTrackingGroup.Get("/events", sse.StreamResource("logistics.shipment-tracking"))
	}

	// --- shipments için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ShipmentsController := &controller.ShipmentsController{
			Svc: &services.ShipmentsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ShipmentsGroup := app.Group("/shipments")

		// Rotaları tanımla
		ShipmentsGroup.Get("/pagination", ShipmentsController.GetShipmentsWithPagination)

		ShipmentsGroup.Get("/with-id/:id", ShipmentsController.GetShipmentsById)
		ShipmentsGroup.Post("/", ShipmentsController.CreateShipments)
		ShipmentsGroup.Post("/bulk", ShipmentsController.CreateShipmentsMultiple)
		ShipmentsGroup.Put("/with-id/:id", ShipmentsController.UpdateShipments)
		ShipmentsGroup.Delete("/with-id/:id", ShipmentsController.DeleteShipments)
		ShipmentsGroup.Delete("/bulk", ShipmentsController.DeleteShipmentsMultiple)
		ShipmentsGroup.Put("/bulk", ShipmentsController.UpdateShipmentsMultiple)

		ShipmentsGroup.Get("/events", sse.StreamResource("logistics.shipments"))
	}

	// --- stock_movements için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		StockMovementsController := &controller.StockMovementsController{
			Svc: &services.StockMovementsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		StockMovementsGroup := app.Group("/stock-movements")

		// Rotaları tanımla
		StockMovementsGroup.Get("/pagination", StockMovementsController.GetStockMovementsWithPagination)

		StockMovementsGroup.Get("/with-id/:id", StockMovementsController.GetStockMovementsById)
		StockMovementsGroup.Post("/", StockMovementsController.CreateStockMovements)
		StockMovementsGroup.Post("/bulk", StockMovementsController.CreateStockMovementsMultiple)
		StockMovementsGroup.Put("/with-id/:id", StockMovementsController.UpdateStockMovements)
		StockMovementsGroup.Delete("/with-id/:id", StockMovementsController.DeleteStockMovements)
		StockMovementsGroup.Delete("/bulk", StockMovementsController.DeleteStockMovementsMultiple)
		StockMovementsGroup.Put("/bulk", StockMovementsController.UpdateStockMovementsMultiple)

		StockMovementsGroup.Get("/events", sse.StreamResource("logistics.stock-movements"))
	}

	// --- storage_bins için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		StorageBinsController := &controller.StorageBinsController{
			Svc: &services.StorageBinsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		StorageBinsGroup := app.Group("/storage-bins")

		// Rotaları tanımla
		StorageBinsGroup.Get("/pagination", StorageBinsController.GetStorageBinsWithPagination)

		StorageBinsGroup.Get("/with-id/:id", StorageBinsController.GetStorageBinsById)
		StorageBinsGroup.Post("/", StorageBinsController.CreateStorageBins)
		StorageBinsGroup.Post("/bulk", StorageBinsController.CreateStorageBinsMultiple)
		StorageBinsGroup.Put("/with-id/:id", StorageBinsController.UpdateStorageBins)
		StorageBinsGroup.Delete("/with-id/:id", StorageBinsController.DeleteStorageBins)
		StorageBinsGroup.Delete("/bulk", StorageBinsController.DeleteStorageBinsMultiple)
		StorageBinsGroup.Put("/bulk", StorageBinsController.UpdateStorageBinsMultiple)

		StorageBinsGroup.Get("/events", sse.StreamResource("logistics.storage-bins"))
	}

	// --- storage_zones için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		StorageZonesController := &controller.StorageZonesController{
			Svc: &services.StorageZonesService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		StorageZonesGroup := app.Group("/storage-zones")

		// Rotaları tanımla
		StorageZonesGroup.Get("/pagination", StorageZonesController.GetStorageZonesWithPagination)

		StorageZonesGroup.Get("/with-id/:id", StorageZonesController.GetStorageZonesById)
		StorageZonesGroup.Post("/", StorageZonesController.CreateStorageZones)
		StorageZonesGroup.Post("/bulk", StorageZonesController.CreateStorageZonesMultiple)
		StorageZonesGroup.Put("/with-id/:id", StorageZonesController.UpdateStorageZones)
		StorageZonesGroup.Delete("/with-id/:id", StorageZonesController.DeleteStorageZones)
		StorageZonesGroup.Delete("/bulk", StorageZonesController.DeleteStorageZonesMultiple)
		StorageZonesGroup.Put("/bulk", StorageZonesController.UpdateStorageZonesMultiple)

		StorageZonesGroup.Get("/events", sse.StreamResource("logistics.storage-zones"))
	}

	// --- suppliers için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		SuppliersController := &controller.SuppliersController{
			Svc: &services.SuppliersService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		SuppliersGroup := app.Group("/suppliers")

		// Rotaları tanımla
		SuppliersGroup.Get("/pagination", SuppliersController.GetSuppliersWithPagination)

		SuppliersGroup.Get("/with-id/:id", SuppliersController.GetSuppliersById)
		SuppliersGroup.Post("/", SuppliersController.CreateSuppliers)
		SuppliersGroup.Post("/bulk", SuppliersController.CreateSuppliersMultiple)
		SuppliersGroup.Put("/with-id/:id", SuppliersController.UpdateSuppliers)
		SuppliersGroup.Delete("/with-id/:id", SuppliersController.DeleteSuppliers)
		SuppliersGroup.Delete("/bulk", SuppliersController.DeleteSuppliersMultiple)
		SuppliersGroup.Put("/bulk", SuppliersController.UpdateSuppliersMultiple)

		SuppliersGroup.Get("/events", sse.StreamResource("logistics.suppliers"))
	}

	// --- warehouses için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		WarehousesController := &controller.WarehousesController{
			Svc: &services.WarehousesService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		WarehousesGroup := app.Group("/warehouses")

		// Rotaları tanımla
		WarehousesGroup.Get("/pagination", WarehousesController.GetWarehousesWithPagination)

		WarehousesGroup.Get("/with-id/:id", WarehousesController.GetWarehousesById)
		WarehousesGroup.Post("/", WarehousesController.CreateWarehouses)
		WarehousesGroup.Post("/bulk", WarehousesController.CreateWarehousesMultiple)
		WarehousesGroup.Put("/with-id/:id", WarehousesController.UpdateWarehouses)
		WarehousesGroup.Delete("/with-id/:id", WarehousesController.DeleteWarehouses)
		WarehousesGroup.Delete("/bulk", WarehousesController.DeleteWarehousesMultiple)
		WarehousesGroup.Put("/bulk", WarehousesController.UpdateWarehousesMultiple)

		WarehousesGroup.Get("/events", sse.StreamResource("logistics.warehouses"))
	}

	rpcSvc := &services.RPCService{DB: db}
	rpcController := &controller.RPCController{Svc: rpcSvc}
	rpcGroup := app.Group("/rpc")
	{
		rpcGroup.Post("/low-stock-count", rpcController.LowStockCount)
		rpcGroup.Post("/warehouse-utilization", rpcController.WarehouseUtilization)
	}
}
