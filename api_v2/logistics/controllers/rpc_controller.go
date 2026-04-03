package logistics_api_controller

import (
	services "data-bridge-examples/api_v2/logistics/services"
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

// LowStockCountRequest, LowStockCount RPC fonksiyonunun girdi parametrelerini temsil eder.
type LowStockCountRequest struct {
	PWarehouseId types.URID `json:"p_warehouse_id"`
}

// LowStockCount, `POST /rpc/low-stock-count` endpoint'ini yönetir.
func (c *RPCController) LowStockCount(ctx *fiber.Ctx) error {
	req := new(LowStockCountRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.LowStockCount(req.PWarehouseId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// WarehouseUtilizationRequest, WarehouseUtilization RPC fonksiyonunun girdi parametrelerini temsil eder.
type WarehouseUtilizationRequest struct {
	PWarehouseId types.URID `json:"p_warehouse_id"`
}

// WarehouseUtilization, `POST /rpc/warehouse-utilization` endpoint'ini yönetir.
func (c *RPCController) WarehouseUtilization(ctx *fiber.Ctx) error {
	req := new(WarehouseUtilizationRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.WarehouseUtilization(req.PWarehouseId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
