package iam_api_controller

import (
	services "data-bridge-examples/api_v2/iam/services"
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

// CountActiveUsersRequest, CountActiveUsers RPC fonksiyonunun girdi parametrelerini temsil eder.
type CountActiveUsersRequest struct {
}

// CountActiveUsers, `POST /rpc/count-active-users` endpoint'ini yönetir.
func (c *RPCController) CountActiveUsers(ctx *fiber.Ctx) error {
	req := new(CountActiveUsersRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.CountActiveUsers()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// UserPermissionsRequest, UserPermissions RPC fonksiyonunun girdi parametrelerini temsil eder.
type UserPermissionsRequest struct {
	PUserId types.URID `json:"p_user_id"`
}

// UserPermissions, `POST /rpc/user-permissions` endpoint'ini yönetir.
func (c *RPCController) UserPermissions(ctx *fiber.Ctx) error {
	req := new(UserPermissionsRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.UserPermissions(req.PUserId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// UsersByOrganizationRequest, UsersByOrganization RPC fonksiyonunun girdi parametrelerini temsil eder.
type UsersByOrganizationRequest struct {
	POrgId types.URID `json:"p_org_id"`
}

// UsersByOrganization, `POST /rpc/users-by-organization` endpoint'ini yönetir.
func (c *RPCController) UsersByOrganization(ctx *fiber.Ctx) error {
	req := new(UsersByOrganizationRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body", Type: "BodyParseError"})
	}

	// TODO: Gelen istek üzerinde validasyon yapılabilir.

	result, err := c.Svc.UsersByOrganization(req.POrgId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "RPC execution failed", Type: "RPCError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
