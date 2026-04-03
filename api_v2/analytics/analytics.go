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
package analytics

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Dinamik olarak ilgili paketleri import et
	controller "data-bridge-examples/api_v2/analytics/controllers"
	services "data-bridge-examples/api_v2/analytics/services"

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

	// --- dashboards için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		DashboardsController := &controller.DashboardsController{
			Svc: &services.DashboardsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		DashboardsGroup := app.Group("/dashboards")

		// Rotaları tanımla
		DashboardsGroup.Get("/pagination", DashboardsController.GetDashboardsWithPagination)

		DashboardsGroup.Get("/with-id/:id", DashboardsController.GetDashboardsById)
		DashboardsGroup.Post("/", DashboardsController.CreateDashboards)
		DashboardsGroup.Post("/bulk", DashboardsController.CreateDashboardsMultiple)
		DashboardsGroup.Put("/with-id/:id", DashboardsController.UpdateDashboards)
		DashboardsGroup.Delete("/with-id/:id", DashboardsController.DeleteDashboards)
		DashboardsGroup.Delete("/bulk", DashboardsController.DeleteDashboardsMultiple)
		DashboardsGroup.Put("/bulk", DashboardsController.UpdateDashboardsMultiple)

		DashboardsGroup.Get("/events", sse.StreamResource("analytics.dashboards"))
	}

	// --- events için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		EventsController := &controller.EventsController{
			Svc: &services.EventsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		EventsGroup := app.Group("/events")

		// Rotaları tanımla
		EventsGroup.Get("/pagination", EventsController.GetEventsWithPagination)

		EventsGroup.Get("/with-id/:id", EventsController.GetEventsById)
		EventsGroup.Post("/", EventsController.CreateEvents)
		EventsGroup.Post("/bulk", EventsController.CreateEventsMultiple)
		EventsGroup.Put("/with-id/:id", EventsController.UpdateEvents)
		EventsGroup.Delete("/with-id/:id", EventsController.DeleteEvents)
		EventsGroup.Delete("/bulk", EventsController.DeleteEventsMultiple)
		EventsGroup.Put("/bulk", EventsController.UpdateEventsMultiple)

		EventsGroup.Get("/events", sse.StreamResource("analytics.events"))
	}

	// --- metrics için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		MetricsController := &controller.MetricsController{
			Svc: &services.MetricsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		MetricsGroup := app.Group("/metrics")

		// Rotaları tanımla
		MetricsGroup.Get("/pagination", MetricsController.GetMetricsWithPagination)

		MetricsGroup.Get("/with-id/:id", MetricsController.GetMetricsById)
		MetricsGroup.Post("/", MetricsController.CreateMetrics)
		MetricsGroup.Post("/bulk", MetricsController.CreateMetricsMultiple)
		MetricsGroup.Put("/with-id/:id", MetricsController.UpdateMetrics)
		MetricsGroup.Delete("/with-id/:id", MetricsController.DeleteMetrics)
		MetricsGroup.Delete("/bulk", MetricsController.DeleteMetricsMultiple)
		MetricsGroup.Put("/bulk", MetricsController.UpdateMetricsMultiple)

		MetricsGroup.Get("/events", sse.StreamResource("analytics.metrics"))
	}

	// --- notifications için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		NotificationsController := &controller.NotificationsController{
			Svc: &services.NotificationsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		NotificationsGroup := app.Group("/notifications")

		// Rotaları tanımla
		NotificationsGroup.Get("/pagination", NotificationsController.GetNotificationsWithPagination)

		NotificationsGroup.Get("/with-id/:id", NotificationsController.GetNotificationsById)
		NotificationsGroup.Post("/", NotificationsController.CreateNotifications)
		NotificationsGroup.Post("/bulk", NotificationsController.CreateNotificationsMultiple)
		NotificationsGroup.Put("/with-id/:id", NotificationsController.UpdateNotifications)
		NotificationsGroup.Delete("/with-id/:id", NotificationsController.DeleteNotifications)
		NotificationsGroup.Delete("/bulk", NotificationsController.DeleteNotificationsMultiple)
		NotificationsGroup.Put("/bulk", NotificationsController.UpdateNotificationsMultiple)

		NotificationsGroup.Get("/events", sse.StreamResource("analytics.notifications"))
	}

	// --- reports için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ReportsController := &controller.ReportsController{
			Svc: &services.ReportsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ReportsGroup := app.Group("/reports")

		// Rotaları tanımla
		ReportsGroup.Get("/pagination", ReportsController.GetReportsWithPagination)

		ReportsGroup.Get("/with-id/:id", ReportsController.GetReportsById)
		ReportsGroup.Post("/", ReportsController.CreateReports)
		ReportsGroup.Post("/bulk", ReportsController.CreateReportsMultiple)
		ReportsGroup.Put("/with-id/:id", ReportsController.UpdateReports)
		ReportsGroup.Delete("/with-id/:id", ReportsController.DeleteReports)
		ReportsGroup.Delete("/bulk", ReportsController.DeleteReportsMultiple)
		ReportsGroup.Put("/bulk", ReportsController.UpdateReportsMultiple)

		ReportsGroup.Get("/events", sse.StreamResource("analytics.reports"))
	}

	// --- alert_history için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		AlertHistoryController := &controller.AlertHistoryController{
			Svc: &services.AlertHistoryService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		AlertHistoryGroup := app.Group("/alert-history")

		// Rotaları tanımla
		AlertHistoryGroup.Get("/pagination", AlertHistoryController.GetAlertHistoryWithPagination)

		AlertHistoryGroup.Get("/with-id/:id", AlertHistoryController.GetAlertHistoryById)
		AlertHistoryGroup.Post("/", AlertHistoryController.CreateAlertHistory)
		AlertHistoryGroup.Post("/bulk", AlertHistoryController.CreateAlertHistoryMultiple)
		AlertHistoryGroup.Put("/with-id/:id", AlertHistoryController.UpdateAlertHistory)
		AlertHistoryGroup.Delete("/with-id/:id", AlertHistoryController.DeleteAlertHistory)
		AlertHistoryGroup.Delete("/bulk", AlertHistoryController.DeleteAlertHistoryMultiple)
		AlertHistoryGroup.Put("/bulk", AlertHistoryController.UpdateAlertHistoryMultiple)

		AlertHistoryGroup.Get("/events", sse.StreamResource("analytics.alert-history"))
	}

	// --- alert_rules için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		AlertRulesController := &controller.AlertRulesController{
			Svc: &services.AlertRulesService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		AlertRulesGroup := app.Group("/alert-rules")

		// Rotaları tanımla
		AlertRulesGroup.Get("/pagination", AlertRulesController.GetAlertRulesWithPagination)

		AlertRulesGroup.Get("/with-id/:id", AlertRulesController.GetAlertRulesById)
		AlertRulesGroup.Post("/", AlertRulesController.CreateAlertRules)
		AlertRulesGroup.Post("/bulk", AlertRulesController.CreateAlertRulesMultiple)
		AlertRulesGroup.Put("/with-id/:id", AlertRulesController.UpdateAlertRules)
		AlertRulesGroup.Delete("/with-id/:id", AlertRulesController.DeleteAlertRules)
		AlertRulesGroup.Delete("/bulk", AlertRulesController.DeleteAlertRulesMultiple)
		AlertRulesGroup.Put("/bulk", AlertRulesController.UpdateAlertRulesMultiple)

		AlertRulesGroup.Get("/events", sse.StreamResource("analytics.alert-rules"))
	}

	// --- audit_logs için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		AuditLogsController := &controller.AuditLogsController{
			Svc: &services.AuditLogsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		AuditLogsGroup := app.Group("/audit-logs")

		// Rotaları tanımla
		AuditLogsGroup.Get("/pagination", AuditLogsController.GetAuditLogsWithPagination)

		AuditLogsGroup.Get("/with-id/:id", AuditLogsController.GetAuditLogsById)
		AuditLogsGroup.Post("/", AuditLogsController.CreateAuditLogs)
		AuditLogsGroup.Post("/bulk", AuditLogsController.CreateAuditLogsMultiple)
		AuditLogsGroup.Put("/with-id/:id", AuditLogsController.UpdateAuditLogs)
		AuditLogsGroup.Delete("/with-id/:id", AuditLogsController.DeleteAuditLogs)
		AuditLogsGroup.Delete("/bulk", AuditLogsController.DeleteAuditLogsMultiple)
		AuditLogsGroup.Put("/bulk", AuditLogsController.UpdateAuditLogsMultiple)

		AuditLogsGroup.Get("/events", sse.StreamResource("analytics.audit-logs"))
	}

	// --- dashboard_widgets için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		DashboardWidgetsController := &controller.DashboardWidgetsController{
			Svc: &services.DashboardWidgetsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		DashboardWidgetsGroup := app.Group("/dashboard-widgets")

		// Rotaları tanımla
		DashboardWidgetsGroup.Get("/pagination", DashboardWidgetsController.GetDashboardWidgetsWithPagination)

		DashboardWidgetsGroup.Get("/with-id/:id", DashboardWidgetsController.GetDashboardWidgetsById)
		DashboardWidgetsGroup.Post("/", DashboardWidgetsController.CreateDashboardWidgets)
		DashboardWidgetsGroup.Post("/bulk", DashboardWidgetsController.CreateDashboardWidgetsMultiple)
		DashboardWidgetsGroup.Put("/with-id/:id", DashboardWidgetsController.UpdateDashboardWidgets)
		DashboardWidgetsGroup.Delete("/with-id/:id", DashboardWidgetsController.DeleteDashboardWidgets)
		DashboardWidgetsGroup.Delete("/bulk", DashboardWidgetsController.DeleteDashboardWidgetsMultiple)
		DashboardWidgetsGroup.Put("/bulk", DashboardWidgetsController.UpdateDashboardWidgetsMultiple)

		DashboardWidgetsGroup.Get("/events", sse.StreamResource("analytics.dashboard-widgets"))
	}

	// --- recent_events için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		RecentEventsController := &controller.RecentEventsController{
			Svc: &services.RecentEventsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		RecentEventsGroup := app.Group("/recent-events")

		// Rotaları tanımla
		RecentEventsGroup.Get("/pagination", RecentEventsController.GetRecentEventsWithPagination)

		RecentEventsGroup.Get("/events", sse.StreamResource("analytics.recent-events"))
	}

	// --- report_executions için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ReportExecutionsController := &controller.ReportExecutionsController{
			Svc: &services.ReportExecutionsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ReportExecutionsGroup := app.Group("/report-executions")

		// Rotaları tanımla
		ReportExecutionsGroup.Get("/pagination", ReportExecutionsController.GetReportExecutionsWithPagination)

		ReportExecutionsGroup.Get("/with-id/:id", ReportExecutionsController.GetReportExecutionsById)
		ReportExecutionsGroup.Post("/", ReportExecutionsController.CreateReportExecutions)
		ReportExecutionsGroup.Post("/bulk", ReportExecutionsController.CreateReportExecutionsMultiple)
		ReportExecutionsGroup.Put("/with-id/:id", ReportExecutionsController.UpdateReportExecutions)
		ReportExecutionsGroup.Delete("/with-id/:id", ReportExecutionsController.DeleteReportExecutions)
		ReportExecutionsGroup.Delete("/bulk", ReportExecutionsController.DeleteReportExecutionsMultiple)
		ReportExecutionsGroup.Put("/bulk", ReportExecutionsController.UpdateReportExecutionsMultiple)

		ReportExecutionsGroup.Get("/events", sse.StreamResource("analytics.report-executions"))
	}

	// --- unread_notifications için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		UnreadNotificationsController := &controller.UnreadNotificationsController{
			Svc: &services.UnreadNotificationsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		UnreadNotificationsGroup := app.Group("/unread-notifications")

		// Rotaları tanımla
		UnreadNotificationsGroup.Get("/pagination", UnreadNotificationsController.GetUnreadNotificationsWithPagination)

		UnreadNotificationsGroup.Get("/events", sse.StreamResource("analytics.unread-notifications"))
	}

	rpcSvc := &services.RPCService{DB: db}
	rpcController := &controller.RPCController{Svc: rpcSvc}
	rpcGroup := app.Group("/rpc")
	{
		rpcGroup.Post("/dashboard-count", rpcController.DashboardCount)
		rpcGroup.Post("/event-count-by-severity", rpcController.EventCountBySeverity)
		rpcGroup.Post("/unread-notification-count", rpcController.UnreadNotificationCount)
	}
}
