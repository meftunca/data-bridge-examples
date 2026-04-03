package orders_api_controller

import (
	services "data-bridge-examples/api_v2/orders/services"
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

// CustomerTotalSpentRequest, CustomerTotalSpent RPC fonksiyonunun girdi parametrelerini temsil eder.
type CustomerTotalSpentRequest struct {
	PCustomerId types.URID `json:"p_customer_id"`
}

// CustomerTotalSpent, `POST /rpc/customer-total-spent` endpoint'ini yönetir.
func (c *RPCController) CustomerTotalSpent(ctx *fiber.Ctx) error {
	req := new(CustomerTotalSpentRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.CustomerTotalSpent(req.PCustomerId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// OrdersByStatusRequest, OrdersByStatus RPC fonksiyonunun girdi parametrelerini temsil eder.
type OrdersByStatusRequest struct {
	PStatus string `json:"p_status"`
}

// OrdersByStatus, `POST /rpc/orders-by-status` endpoint'ini yönetir.
func (c *RPCController) OrdersByStatus(ctx *fiber.Ctx) error {
	req := new(OrdersByStatusRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.OrdersByStatus(req.PStatus)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// TotalRevenueRequest, TotalRevenue RPC fonksiyonunun girdi parametrelerini temsil eder.
type TotalRevenueRequest struct {
}

// TotalRevenue, `POST /rpc/total-revenue` endpoint'ini yönetir.
func (c *RPCController) TotalRevenue(ctx *fiber.Ctx) error {
	req := new(TotalRevenueRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.TotalRevenue()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
