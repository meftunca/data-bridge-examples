package api_v2

import (
	analytics "data-bridge-examples/api_v2/analytics"
	catalog "data-bridge-examples/api_v2/catalog"
	iam "data-bridge-examples/api_v2/iam"
	logistics "data-bridge-examples/api_v2/logistics"
	orders "data-bridge-examples/api_v2/orders"

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/events"
	"gorm.io/gorm"
)

// Setup configures all API routes for the application.
func Setup(app fiber.Router, db *gorm.DB, eventManager *events.EventManager) {
	iam.Run(app.Group("/iam"), db, eventManager)
	catalog.Run(app.Group("/catalog"), db, eventManager)
	orders.Run(app.Group("/orders"), db, eventManager)
	logistics.Run(app.Group("/logistics"), db, eventManager)
	analytics.Run(app.Group("/analytics"), db, eventManager)

}
