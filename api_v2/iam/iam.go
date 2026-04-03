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
package iam

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Dinamik olarak ilgili paketleri import et
	controller "data-bridge-examples/api_v2/iam/controllers"
	services "data-bridge-examples/api_v2/iam/services"

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

	// --- invitations için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		InvitationsController := &controller.InvitationsController{
			Svc: &services.InvitationsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		InvitationsGroup := app.Group("/invitations")

		// Rotaları tanımla
		InvitationsGroup.Get("/pagination", InvitationsController.GetInvitationsWithPagination)

		InvitationsGroup.Get("/with-id/:id", InvitationsController.GetInvitationsById)
		InvitationsGroup.Post("/", InvitationsController.CreateInvitations)
		InvitationsGroup.Post("/bulk", InvitationsController.CreateInvitationsMultiple)
		InvitationsGroup.Put("/with-id/:id", InvitationsController.UpdateInvitations)
		InvitationsGroup.Delete("/with-id/:id", InvitationsController.DeleteInvitations)
		InvitationsGroup.Delete("/bulk", InvitationsController.DeleteInvitationsMultiple)
		InvitationsGroup.Put("/bulk", InvitationsController.UpdateInvitationsMultiple)

		InvitationsGroup.Get("/events", sse.StreamResource("iam.invitations"))
	}

	// --- organizations için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		OrganizationsController := &controller.OrganizationsController{
			Svc: &services.OrganizationsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		OrganizationsGroup := app.Group("/organizations")

		// Rotaları tanımla
		OrganizationsGroup.Get("/pagination", OrganizationsController.GetOrganizationsWithPagination)

		OrganizationsGroup.Get("/with-id/:id", OrganizationsController.GetOrganizationsById)
		OrganizationsGroup.Post("/", OrganizationsController.CreateOrganizations)
		OrganizationsGroup.Post("/bulk", OrganizationsController.CreateOrganizationsMultiple)
		OrganizationsGroup.Put("/with-id/:id", OrganizationsController.UpdateOrganizations)
		OrganizationsGroup.Delete("/with-id/:id", OrganizationsController.DeleteOrganizations)
		OrganizationsGroup.Delete("/bulk", OrganizationsController.DeleteOrganizationsMultiple)
		OrganizationsGroup.Put("/bulk", OrganizationsController.UpdateOrganizationsMultiple)

		OrganizationsGroup.Get("/events", sse.StreamResource("iam.organizations"))
	}

	// --- permissions için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		PermissionsController := &controller.PermissionsController{
			Svc: &services.PermissionsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		PermissionsGroup := app.Group("/permissions")

		// Rotaları tanımla
		PermissionsGroup.Get("/pagination", PermissionsController.GetPermissionsWithPagination)

		PermissionsGroup.Get("/with-id/:id", PermissionsController.GetPermissionsById)
		PermissionsGroup.Post("/", PermissionsController.CreatePermissions)
		PermissionsGroup.Post("/bulk", PermissionsController.CreatePermissionsMultiple)
		PermissionsGroup.Put("/with-id/:id", PermissionsController.UpdatePermissions)
		PermissionsGroup.Delete("/with-id/:id", PermissionsController.DeletePermissions)
		PermissionsGroup.Delete("/bulk", PermissionsController.DeletePermissionsMultiple)
		PermissionsGroup.Put("/bulk", PermissionsController.UpdatePermissionsMultiple)

		PermissionsGroup.Get("/events", sse.StreamResource("iam.permissions"))
	}

	// --- roles için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		RolesController := &controller.RolesController{
			Svc: &services.RolesService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		RolesGroup := app.Group("/roles")

		// Rotaları tanımla
		RolesGroup.Get("/pagination", RolesController.GetRolesWithPagination)

		RolesGroup.Get("/with-id/:id", RolesController.GetRolesById)
		RolesGroup.Post("/", RolesController.CreateRoles)
		RolesGroup.Post("/bulk", RolesController.CreateRolesMultiple)
		RolesGroup.Put("/with-id/:id", RolesController.UpdateRoles)
		RolesGroup.Delete("/with-id/:id", RolesController.DeleteRoles)
		RolesGroup.Delete("/bulk", RolesController.DeleteRolesMultiple)
		RolesGroup.Put("/bulk", RolesController.UpdateRolesMultiple)

		RolesGroup.Get("/events", sse.StreamResource("iam.roles"))
	}

	// --- sessions için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		SessionsController := &controller.SessionsController{
			Svc: &services.SessionsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		SessionsGroup := app.Group("/sessions")

		// Rotaları tanımla
		SessionsGroup.Get("/pagination", SessionsController.GetSessionsWithPagination)

		SessionsGroup.Get("/with-id/:id", SessionsController.GetSessionsById)
		SessionsGroup.Post("/", SessionsController.CreateSessions)
		SessionsGroup.Post("/bulk", SessionsController.CreateSessionsMultiple)
		SessionsGroup.Put("/with-id/:id", SessionsController.UpdateSessions)
		SessionsGroup.Delete("/with-id/:id", SessionsController.DeleteSessions)
		SessionsGroup.Delete("/bulk", SessionsController.DeleteSessionsMultiple)
		SessionsGroup.Put("/bulk", SessionsController.UpdateSessionsMultiple)

		SessionsGroup.Get("/events", sse.StreamResource("iam.sessions"))
	}

	// --- teams için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		TeamsController := &controller.TeamsController{
			Svc: &services.TeamsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		TeamsGroup := app.Group("/teams")

		// Rotaları tanımla
		TeamsGroup.Get("/pagination", TeamsController.GetTeamsWithPagination)

		TeamsGroup.Get("/with-id/:id", TeamsController.GetTeamsById)
		TeamsGroup.Post("/", TeamsController.CreateTeams)
		TeamsGroup.Post("/bulk", TeamsController.CreateTeamsMultiple)
		TeamsGroup.Put("/with-id/:id", TeamsController.UpdateTeams)
		TeamsGroup.Delete("/with-id/:id", TeamsController.DeleteTeams)
		TeamsGroup.Delete("/bulk", TeamsController.DeleteTeamsMultiple)
		TeamsGroup.Put("/bulk", TeamsController.UpdateTeamsMultiple)

		TeamsGroup.Get("/events", sse.StreamResource("iam.teams"))
	}

	// --- users için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		UsersController := &controller.UsersController{
			Svc: &services.UsersService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		UsersGroup := app.Group("/users")

		// Rotaları tanımla
		UsersGroup.Get("/pagination", UsersController.GetUsersWithPagination)

		UsersGroup.Get("/with-id/:id", UsersController.GetUsersById)
		UsersGroup.Post("/", UsersController.CreateUsers)
		UsersGroup.Post("/bulk", UsersController.CreateUsersMultiple)
		UsersGroup.Put("/with-id/:id", UsersController.UpdateUsers)
		UsersGroup.Delete("/with-id/:id", UsersController.DeleteUsers)
		UsersGroup.Delete("/bulk", UsersController.DeleteUsersMultiple)
		UsersGroup.Put("/bulk", UsersController.UpdateUsersMultiple)

		UsersGroup.Get("/events", sse.StreamResource("iam.users"))
	}

	// --- api_keys için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ApiKeysController := &controller.ApiKeysController{
			Svc: &services.ApiKeysService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ApiKeysGroup := app.Group("/api-keys")

		// Rotaları tanımla
		ApiKeysGroup.Get("/pagination", ApiKeysController.GetApiKeysWithPagination)

		ApiKeysGroup.Get("/with-id/:id", ApiKeysController.GetApiKeysById)
		ApiKeysGroup.Post("/", ApiKeysController.CreateApiKeys)
		ApiKeysGroup.Post("/bulk", ApiKeysController.CreateApiKeysMultiple)
		ApiKeysGroup.Put("/with-id/:id", ApiKeysController.UpdateApiKeys)
		ApiKeysGroup.Delete("/with-id/:id", ApiKeysController.DeleteApiKeys)
		ApiKeysGroup.Delete("/bulk", ApiKeysController.DeleteApiKeysMultiple)
		ApiKeysGroup.Put("/bulk", ApiKeysController.UpdateApiKeysMultiple)

		ApiKeysGroup.Get("/events", sse.StreamResource("iam.api-keys"))
	}

	// --- role_permissions için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		RolePermissionsController := &controller.RolePermissionsController{
			Svc: &services.RolePermissionsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		RolePermissionsGroup := app.Group("/role-permissions")

		// Rotaları tanımla
		RolePermissionsGroup.Get("/pagination", RolePermissionsController.GetRolePermissionsWithPagination)

		RolePermissionsGroup.Get("/with-id/:id", RolePermissionsController.GetRolePermissionsById)
		RolePermissionsGroup.Post("/", RolePermissionsController.CreateRolePermissions)
		RolePermissionsGroup.Post("/bulk", RolePermissionsController.CreateRolePermissionsMultiple)
		RolePermissionsGroup.Put("/with-id/:id", RolePermissionsController.UpdateRolePermissions)
		RolePermissionsGroup.Delete("/with-id/:id", RolePermissionsController.DeleteRolePermissions)
		RolePermissionsGroup.Delete("/bulk", RolePermissionsController.DeleteRolePermissionsMultiple)
		RolePermissionsGroup.Put("/bulk", RolePermissionsController.UpdateRolePermissionsMultiple)

		RolePermissionsGroup.Get("/events", sse.StreamResource("iam.role-permissions"))
	}

	// --- team_members için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		TeamMembersController := &controller.TeamMembersController{
			Svc: &services.TeamMembersService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		TeamMembersGroup := app.Group("/team-members")

		// Rotaları tanımla
		TeamMembersGroup.Get("/pagination", TeamMembersController.GetTeamMembersWithPagination)

		TeamMembersGroup.Get("/with-id/:id", TeamMembersController.GetTeamMembersById)
		TeamMembersGroup.Post("/", TeamMembersController.CreateTeamMembers)
		TeamMembersGroup.Post("/bulk", TeamMembersController.CreateTeamMembersMultiple)
		TeamMembersGroup.Put("/with-id/:id", TeamMembersController.UpdateTeamMembers)
		TeamMembersGroup.Delete("/with-id/:id", TeamMembersController.DeleteTeamMembers)
		TeamMembersGroup.Delete("/bulk", TeamMembersController.DeleteTeamMembersMultiple)
		TeamMembersGroup.Put("/bulk", TeamMembersController.UpdateTeamMembersMultiple)

		TeamMembersGroup.Get("/events", sse.StreamResource("iam.team-members"))
	}

	// --- user_roles için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		UserRolesController := &controller.UserRolesController{
			Svc: &services.UserRolesService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		UserRolesGroup := app.Group("/user-roles")

		// Rotaları tanımla
		UserRolesGroup.Get("/pagination", UserRolesController.GetUserRolesWithPagination)

		UserRolesGroup.Get("/with-id/:id", UserRolesController.GetUserRolesById)
		UserRolesGroup.Post("/", UserRolesController.CreateUserRoles)
		UserRolesGroup.Post("/bulk", UserRolesController.CreateUserRolesMultiple)
		UserRolesGroup.Put("/with-id/:id", UserRolesController.UpdateUserRoles)
		UserRolesGroup.Delete("/with-id/:id", UserRolesController.DeleteUserRoles)
		UserRolesGroup.Delete("/bulk", UserRolesController.DeleteUserRolesMultiple)
		UserRolesGroup.Put("/bulk", UserRolesController.UpdateUserRolesMultiple)

		UserRolesGroup.Get("/events", sse.StreamResource("iam.user-roles"))
	}

	rpcSvc := &services.RPCService{DB: db}
	rpcController := &controller.RPCController{Svc: rpcSvc}
	rpcGroup := app.Group("/rpc")
	{
		rpcGroup.Post("/count-active-users", rpcController.CountActiveUsers)
		rpcGroup.Post("/user-permissions", rpcController.UserPermissions)
		rpcGroup.Post("/users-by-organization", rpcController.UsersByOrganization)
	}
}
