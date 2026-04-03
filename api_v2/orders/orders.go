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
package orders

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Dinamik olarak ilgili paketleri import et
	controller "data-bridge-examples/api_v2/orders/controllers"
	services "data-bridge-examples/api_v2/orders/services"

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

	// --- carts için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		CartsController := &controller.CartsController{
			Svc: &services.CartsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		CartsGroup := app.Group("/carts")

		// Rotaları tanımla
		CartsGroup.Get("/pagination", CartsController.GetCartsWithPagination)

		CartsGroup.Get("/with-id/:id", CartsController.GetCartsById)
		CartsGroup.Post("/", CartsController.CreateCarts)
		CartsGroup.Post("/bulk", CartsController.CreateCartsMultiple)
		CartsGroup.Put("/with-id/:id", CartsController.UpdateCarts)
		CartsGroup.Delete("/with-id/:id", CartsController.DeleteCarts)
		CartsGroup.Delete("/bulk", CartsController.DeleteCartsMultiple)
		CartsGroup.Put("/bulk", CartsController.UpdateCartsMultiple)

		CartsGroup.Get("/events", sse.StreamResource("orders.carts"))
	}

	// --- coupons için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		CouponsController := &controller.CouponsController{
			Svc: &services.CouponsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		CouponsGroup := app.Group("/coupons")

		// Rotaları tanımla
		CouponsGroup.Get("/pagination", CouponsController.GetCouponsWithPagination)

		CouponsGroup.Get("/with-id/:id", CouponsController.GetCouponsById)
		CouponsGroup.Post("/", CouponsController.CreateCoupons)
		CouponsGroup.Post("/bulk", CouponsController.CreateCouponsMultiple)
		CouponsGroup.Put("/with-id/:id", CouponsController.UpdateCoupons)
		CouponsGroup.Delete("/with-id/:id", CouponsController.DeleteCoupons)
		CouponsGroup.Delete("/bulk", CouponsController.DeleteCouponsMultiple)
		CouponsGroup.Put("/bulk", CouponsController.UpdateCouponsMultiple)

		CouponsGroup.Get("/events", sse.StreamResource("orders.coupons"))
	}

	// --- customers için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		CustomersController := &controller.CustomersController{
			Svc: &services.CustomersService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		CustomersGroup := app.Group("/customers")

		// Rotaları tanımla
		CustomersGroup.Get("/pagination", CustomersController.GetCustomersWithPagination)

		CustomersGroup.Get("/with-id/:id", CustomersController.GetCustomersById)
		CustomersGroup.Post("/", CustomersController.CreateCustomers)
		CustomersGroup.Post("/bulk", CustomersController.CreateCustomersMultiple)
		CustomersGroup.Put("/with-id/:id", CustomersController.UpdateCustomers)
		CustomersGroup.Delete("/with-id/:id", CustomersController.DeleteCustomers)
		CustomersGroup.Delete("/bulk", CustomersController.DeleteCustomersMultiple)
		CustomersGroup.Put("/bulk", CustomersController.UpdateCustomersMultiple)

		CustomersGroup.Get("/events", sse.StreamResource("orders.customers"))
	}

	// --- orders için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		OrdersController := &controller.OrdersController{
			Svc: &services.OrdersService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		OrdersGroup := app.Group("/orders")

		// Rotaları tanımla
		OrdersGroup.Get("/pagination", OrdersController.GetOrdersWithPagination)

		OrdersGroup.Get("/with-id/:id", OrdersController.GetOrdersById)
		OrdersGroup.Post("/", OrdersController.CreateOrders)
		OrdersGroup.Post("/bulk", OrdersController.CreateOrdersMultiple)
		OrdersGroup.Put("/with-id/:id", OrdersController.UpdateOrders)
		OrdersGroup.Delete("/with-id/:id", OrdersController.DeleteOrders)
		OrdersGroup.Delete("/bulk", OrdersController.DeleteOrdersMultiple)
		OrdersGroup.Put("/bulk", OrdersController.UpdateOrdersMultiple)

		OrdersGroup.Get("/events", sse.StreamResource("orders.orders"))
	}

	// --- payments için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		PaymentsController := &controller.PaymentsController{
			Svc: &services.PaymentsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		PaymentsGroup := app.Group("/payments")

		// Rotaları tanımla
		PaymentsGroup.Get("/pagination", PaymentsController.GetPaymentsWithPagination)

		PaymentsGroup.Get("/with-id/:id", PaymentsController.GetPaymentsById)
		PaymentsGroup.Post("/", PaymentsController.CreatePayments)
		PaymentsGroup.Post("/bulk", PaymentsController.CreatePaymentsMultiple)
		PaymentsGroup.Put("/with-id/:id", PaymentsController.UpdatePayments)
		PaymentsGroup.Delete("/with-id/:id", PaymentsController.DeletePayments)
		PaymentsGroup.Delete("/bulk", PaymentsController.DeletePaymentsMultiple)
		PaymentsGroup.Put("/bulk", PaymentsController.UpdatePaymentsMultiple)

		PaymentsGroup.Get("/events", sse.StreamResource("orders.payments"))
	}

	// --- refunds için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		RefundsController := &controller.RefundsController{
			Svc: &services.RefundsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		RefundsGroup := app.Group("/refunds")

		// Rotaları tanımla
		RefundsGroup.Get("/pagination", RefundsController.GetRefundsWithPagination)

		RefundsGroup.Get("/with-id/:id", RefundsController.GetRefundsById)
		RefundsGroup.Post("/", RefundsController.CreateRefunds)
		RefundsGroup.Post("/bulk", RefundsController.CreateRefundsMultiple)
		RefundsGroup.Put("/with-id/:id", RefundsController.UpdateRefunds)
		RefundsGroup.Delete("/with-id/:id", RefundsController.DeleteRefunds)
		RefundsGroup.Delete("/bulk", RefundsController.DeleteRefundsMultiple)
		RefundsGroup.Put("/bulk", RefundsController.UpdateRefundsMultiple)

		RefundsGroup.Get("/events", sse.StreamResource("orders.refunds"))
	}

	// --- cart_items için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		CartItemsController := &controller.CartItemsController{
			Svc: &services.CartItemsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		CartItemsGroup := app.Group("/cart-items")

		// Rotaları tanımla
		CartItemsGroup.Get("/pagination", CartItemsController.GetCartItemsWithPagination)

		CartItemsGroup.Get("/with-id/:id", CartItemsController.GetCartItemsById)
		CartItemsGroup.Post("/", CartItemsController.CreateCartItems)
		CartItemsGroup.Post("/bulk", CartItemsController.CreateCartItemsMultiple)
		CartItemsGroup.Put("/with-id/:id", CartItemsController.UpdateCartItems)
		CartItemsGroup.Delete("/with-id/:id", CartItemsController.DeleteCartItems)
		CartItemsGroup.Delete("/bulk", CartItemsController.DeleteCartItemsMultiple)
		CartItemsGroup.Put("/bulk", CartItemsController.UpdateCartItemsMultiple)

		CartItemsGroup.Get("/events", sse.StreamResource("orders.cart-items"))
	}

	// --- order_items için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		OrderItemsController := &controller.OrderItemsController{
			Svc: &services.OrderItemsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		OrderItemsGroup := app.Group("/order-items")

		// Rotaları tanımla
		OrderItemsGroup.Get("/pagination", OrderItemsController.GetOrderItemsWithPagination)

		OrderItemsGroup.Get("/with-id/:id", OrderItemsController.GetOrderItemsById)
		OrderItemsGroup.Post("/", OrderItemsController.CreateOrderItems)
		OrderItemsGroup.Post("/bulk", OrderItemsController.CreateOrderItemsMultiple)
		OrderItemsGroup.Put("/with-id/:id", OrderItemsController.UpdateOrderItems)
		OrderItemsGroup.Delete("/with-id/:id", OrderItemsController.DeleteOrderItems)
		OrderItemsGroup.Delete("/bulk", OrderItemsController.DeleteOrderItemsMultiple)
		OrderItemsGroup.Put("/bulk", OrderItemsController.UpdateOrderItemsMultiple)

		OrderItemsGroup.Get("/events", sse.StreamResource("orders.order-items"))
	}

	// --- order_status_history için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		OrderStatusHistoryController := &controller.OrderStatusHistoryController{
			Svc: &services.OrderStatusHistoryService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		OrderStatusHistoryGroup := app.Group("/order-status-history")

		// Rotaları tanımla
		OrderStatusHistoryGroup.Get("/pagination", OrderStatusHistoryController.GetOrderStatusHistoryWithPagination)

		OrderStatusHistoryGroup.Get("/with-id/:id", OrderStatusHistoryController.GetOrderStatusHistoryById)
		OrderStatusHistoryGroup.Post("/", OrderStatusHistoryController.CreateOrderStatusHistory)
		OrderStatusHistoryGroup.Post("/bulk", OrderStatusHistoryController.CreateOrderStatusHistoryMultiple)
		OrderStatusHistoryGroup.Put("/with-id/:id", OrderStatusHistoryController.UpdateOrderStatusHistory)
		OrderStatusHistoryGroup.Delete("/with-id/:id", OrderStatusHistoryController.DeleteOrderStatusHistory)
		OrderStatusHistoryGroup.Delete("/bulk", OrderStatusHistoryController.DeleteOrderStatusHistoryMultiple)
		OrderStatusHistoryGroup.Put("/bulk", OrderStatusHistoryController.UpdateOrderStatusHistoryMultiple)

		OrderStatusHistoryGroup.Get("/events", sse.StreamResource("orders.order-status-history"))
	}

	rpcSvc := &services.RPCService{DB: db}
	rpcController := &controller.RPCController{Svc: rpcSvc}
	rpcGroup := app.Group("/rpc")
	{
		rpcGroup.Post("/customer-total-spent", rpcController.CustomerTotalSpent)
		rpcGroup.Post("/orders-by-status", rpcController.OrdersByStatus)
		rpcGroup.Post("/total-revenue", rpcController.TotalRevenue)
	}
}
