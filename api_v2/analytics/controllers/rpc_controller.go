package analytics_api_controller

import (
	services "data-bridge-examples/api_v2/analytics/services"
	"time" // Zaman tipleri için

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/types"
	"github.com/maple-tech/baseline/web"
)

var (
	_ = types.ID(0)
	_ = types.URID("")
	_ = time.Time{}
)

// RPCController, RPC endpoint'lerini yönetir.
type RPCController struct {
	Svc *services.RPCService
}

// DashboardCountRequest, DashboardCount RPC fonksiyonunun girdi parametrelerini temsil eder.
type DashboardCountRequest struct {
}

// DashboardCount, `POST /rpc/dashboard-count` endpoint'ini yönetir.
func (c *RPCController) DashboardCount(ctx *fiber.Ctx) error {
	req := new(DashboardCountRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.DashboardCount()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// EventCountBySeverityRequest, EventCountBySeverity RPC fonksiyonunun girdi parametrelerini temsil eder.
type EventCountBySeverityRequest struct {
	PSeverity string `json:"p_severity"`
}

// EventCountBySeverity, `POST /rpc/event-count-by-severity` endpoint'ini yönetir.
func (c *RPCController) EventCountBySeverity(ctx *fiber.Ctx) error {
	req := new(EventCountBySeverityRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.EventCountBySeverity(req.PSeverity)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// UnreadNotificationCountRequest, UnreadNotificationCount RPC fonksiyonunun girdi parametrelerini temsil eder.
type UnreadNotificationCountRequest struct {
	PUserId types.URID `json:"p_user_id"`
}

// UnreadNotificationCount, `POST /rpc/unread-notification-count` endpoint'ini yönetir.
func (c *RPCController) UnreadNotificationCount(ctx *fiber.Ctx) error {
	req := new(UnreadNotificationCountRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.UnreadNotificationCount(req.PUserId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
