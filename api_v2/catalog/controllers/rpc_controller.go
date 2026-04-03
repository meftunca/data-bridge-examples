package catalog_api_controller

import (
	services "data-bridge-examples/api_v2/catalog/services"
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

// AvgProductRatingRequest, AvgProductRating RPC fonksiyonunun girdi parametrelerini temsil eder.
type AvgProductRatingRequest struct {
	PProductId types.URID `json:"p_product_id"`
}

// AvgProductRating, `POST /rpc/avg-product-rating` endpoint'ini yönetir.
func (c *RPCController) AvgProductRating(ctx *fiber.Ctx) error {
	req := new(AvgProductRatingRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.AvgProductRating(req.PProductId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// CountActiveProductsRequest, CountActiveProducts RPC fonksiyonunun girdi parametrelerini temsil eder.
type CountActiveProductsRequest struct {
}

// CountActiveProducts, `POST /rpc/count-active-products` endpoint'ini yönetir.
func (c *RPCController) CountActiveProducts(ctx *fiber.Ctx) error {
	req := new(CountActiveProductsRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.CountActiveProducts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// ProductsByCategoryRequest, ProductsByCategory RPC fonksiyonunun girdi parametrelerini temsil eder.
type ProductsByCategoryRequest struct {
	PCategoryId types.URID `json:"p_category_id"`
}

// ProductsByCategory, `POST /rpc/products-by-category` endpoint'ini yönetir.
func (c *RPCController) ProductsByCategory(ctx *fiber.Ctx) error {
	req := new(ProductsByCategoryRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.ProductsByCategory(req.PCategoryId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
