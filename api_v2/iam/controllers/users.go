package iam_api_controller

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	paginationRuntime "backend-generator/apiv2/pagination"
	services "data-bridge-examples/api_v2/iam/services"
	structures "data-bridge-examples/api_v2/iam/structures"

	"github.com/maple-tech/baseline/events"
	"github.com/maple-tech/baseline/types"
	"github.com/maple-tech/baseline/web"
)

// paginationUsers provides advanced filtering, sorting, and pagination
// capabilities for Users resources. Pre-configured with schema analysis
// for optimal query performance and type-safe operations.
var paginationUsers = paginationRuntime.NewPagination[structures.Users](structures.Users{})

var (
	_ = types.ID(0)
	_ = types.URID("")
)

func handleUsersControllerError(ctx *fiber.Ctx, status int, err error, errType string) error {
	if err == nil {
		err = errors.New(strings.ToLower(errType))
	}
	return ctx.Status(status).JSON(web.Fiber400Response{Message: err.Error(), Type: errType})
}

// UsersController handles HTTP requests for Users resources.
//
// This controller implements enterprise-grade patterns including:
// • Request/response lifecycle events for comprehensive monitoring
// • Centralized error handling with appropriate HTTP status codes
// • Context-aware operations with configurable timeout management
// • Service layer delegation maintaining clean separation of concerns
// • Input validation and security measures for production deployment
//
// All methods follow RESTful conventions and return consistent JSON responses
// with proper error handling and status codes for client applications.
type UsersController struct {
	Svc services.IUsersService // Business logic and data operations interface
	EM  *events.EventManager   // Request lifecycle and monitoring events
}

// @Summary      Search Users
// @Description  Searches and filters Users resources without pagination. Returns a simple array of results.
// @Tags         Users
// @Id           IamUsersSearch
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        size           query   int     false "Maximum number of results (1-1000)" default(100) example(50)
// @Param        sort           query   string  false "Sort fields with direction: field1,-field2 (- for DESC)" example(-created_at,name)
// @Param        filters        query   string  false "JSON array: [[\"field\",\"operator\",\"value\"]] - Operators: =,!=,>,>=,<,<=,like,in" example([["name","like","%test%"],["status","=","active"]])
// @Param        preloads       query   string  false "Comma-separated relation names to include" example(user,category,tags)
// @Success      200            {array}   structures.Users "List of Users resources"
// @Failure      400            {object}  web.Fiber400Response "Invalid request parameters"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /users [get]
func (c *UsersController) SearchUsers(ctx *fiber.Ctx) error {
	// Context with timeout for this operation
	timeoutCtx, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	_, err := c.EM.Emit(events.QueryRequest, ctx, nil, "Query request for Users")
	if err != nil {
		return handleUsersControllerError(ctx, fiber.StatusBadRequest, err, "ValidationError")
	}

	dbForSearch := c.Svc.(*services.UsersService).DB

	// Create a channel to handle timeout
	type searchResult struct {
		response []structures.Users
		err      error
	}

	resultChan := make(chan searchResult, 1)

	go func() {
		response, searchErr := paginationUsers.With(dbForSearch.Model(&structures.Users{})).Request(ctx).ResponseTiny()
		resultChan <- searchResult{response: response, err: searchErr}
	}()

	select {
	case result := <-resultChan:
		if result.err != nil {
			c.EM.Emit(events.QueryError, ctx, result.err, "Failed to query Users")
			// Enhanced error handling - check for specific database errors
			if errors.Is(result.err, gorm.ErrRecordNotFound) {
				return handleUsersControllerError(ctx, fiber.StatusNotFound, result.err, "NotFoundError")
			}
			return handleUsersControllerError(ctx, fiber.StatusInternalServerError, result.err, "DatabaseError")
		}
		c.EM.EmitWithData(events.QuerySuccess, ctx, nil, fmt.Sprintf("Successfully queried Users, count: %d", len(result.response)), result.response)
		return ctx.JSON(result.response)

	case <-timeoutCtx.Done():
		timeoutErr := timeoutCtx.Err()
		c.EM.Emit(events.QueryError, ctx, timeoutErr, "Query timeout for Users")
		return handleUsersControllerError(ctx, fiber.StatusRequestTimeout, timeoutErr, "TimeoutError")
	}
}

// @Summary      List Users with Pagination
// @Description  Fetches a paginated list of Users resources with metadata including total count, current page, and pagination links.
// @Tags         Users
// @Id           IamUsersPagination
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        page           query   int     false "Page number (starts from 1)" default(1) example(2)
// @Param        size           query   int     false "Number of items per page (1-1000)" default(10) example(25)
// @Param        sort           query   string  false "Sort fields with direction: field1,-field2 (- for DESC)" example(-created_at,name)
// @Param        filters        query   string  false "JSON array: [[\"field\",\"operator\",\"value\"]] - Operators: =,!=,>,>=,<,<=,like,in" example([["status","=","active"],["created_at",">=","2024-01-01"]])
// @Param        preloads       query   string  false "Comma-separated relation names to include" example(user,category)
// @Success      200            {object}  structures.UsersPage "Paginated Users resources with metadata"
// @Failure      400            {object}  web.Fiber400Response "Invalid pagination parameters"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /users/pagination [get]
func (c *UsersController) GetUsersWithPagination(ctx *fiber.Ctx) error {
	_, err := c.EM.Emit(events.PaginationQueryRequest, ctx, nil, "Pagination query request for Users")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: err.Error(), Type: "RequestRejectedByHandler"})
	}
	dbForPagination := c.Svc.(*services.UsersService).DB
	request := paginationUsers.With(dbForPagination.Model(&structures.Users{})).Request(ctx)

	if ctx.Query("collectionBy") != "" {
		response, pagErr := request.ResponsePaginatedCollection()
		if pagErr != nil {
			c.EM.Emit(events.PaginationQueryError, ctx, pagErr, "Failed to retrieve grouped Users collection")
			return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: pagErr.Error(), Type: "PaginationError"})
		}
		c.EM.EmitWithData(events.PaginationQuerySuccess, ctx, nil, fmt.Sprintf("Successfully retrieved grouped Users collection, count: %d", len(response.Items)), response)
		return ctx.JSON(response)
	}

	if ctx.Query("groupby") != "" || ctx.Query("aggregations") != "" {
		response, pagErr := request.ResponseRaw()
		if pagErr != nil {
			c.EM.Emit(events.PaginationQueryError, ctx, pagErr, "Failed to retrieve aggregated Users response")
			return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: pagErr.Error(), Type: "PaginationError"})
		}
		c.EM.EmitWithData(events.PaginationQuerySuccess, ctx, nil, fmt.Sprintf("Successfully retrieved aggregated Users response, count: %d", len(response.Items)), response)
		return ctx.JSON(response)
	}

	response, pagErr := request.Response()
	if pagErr != nil {
		c.EM.Emit(events.PaginationQueryError, ctx, pagErr, "Failed to retrieve paginated Users")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: pagErr.Error(), Type: "PaginationError"})
	}
	c.EM.EmitWithData(events.PaginationQuerySuccess, ctx, nil, fmt.Sprintf("Successfully retrieved paginated Users, count: %d", len(response.Items)), response)
	return ctx.JSON(response)
}

// @Summary      Get a Users by ID
// @Description  Fetches a single Users resource by its unique identifier(s). Returns the complete resource data.
// @Tags         Users
// @Id           IamUsersWithID
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        id   path    types.URID   true  "Id" example(123)
// @Param        preloads       query   string  false "Comma-separated relation names to include" example(user,category)
// @Success      200            {object}  structures.Users "Users resource found"
// @Failure      400            {object}  web.Fiber400Response "Invalid ID format"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      404            {object}  web.Fiber400Response "Users not found"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /users/with-id/:id [get]
func (c *UsersController) GetUsersById(ctx *fiber.Ctx) error {
	// Create timeout context for this operation
	timeoutCtx, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	idStr := ctx.Params("id")
	var id types.URID
	if err := id.FromBase32(idStr); err != nil { // Veya FromUUID, duruma göre değişebilir
		c.EM.Emit(events.QueryParamsParsingError, ctx, err, "Invalid 'id' format in path parameter")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid " + "id" + " format", Type: "QueryParamsParsingError"})
	}

	// Emit find query request event
	proceed, reqErr := c.EM.Emit(events.FindQueryRequest, ctx, nil, fmt.Sprintf("FindQuery request for Users with ID(s) from path"))
	if !proceed {
		c.EM.Emit(events.FindQueryError, ctx, reqErr, fmt.Sprintf("Users find request rejected by handler"))
		return handleUsersControllerError(ctx, fiber.StatusBadRequest, reqErr, "ValidationError")
	}

	// Execute find with timeout context
	dbForFind := c.Svc.(*services.UsersService).DB.WithContext(timeoutCtx)
	queryBuilder := dbForFind.Model(&structures.Users{})
	queryBuilder = queryBuilder.Where("id = ?", id)
	var result *structures.Users
	findErr := queryBuilder.First(&result).Error

	if findErr != nil {
		// Check if it's a timeout error
		if errors.Is(findErr, context.DeadlineExceeded) {
			c.EM.Emit(events.FindQueryError, ctx, findErr, fmt.Sprintf("Users find timeout with ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusRequestTimeout, findErr, "TimeoutError")
		}

		// Check for record not found
		if errors.Is(findErr, gorm.ErrRecordNotFound) {
			c.EM.Emit(events.FindQueryError, ctx, findErr, fmt.Sprintf("Users not found with provided ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusNotFound, findErr, "NotFoundError")
		}

		c.EM.Emit(events.FindQueryError, ctx, findErr, fmt.Sprintf("Error finding Users with ID(s)"))
		return handleUsersControllerError(ctx, fiber.StatusInternalServerError, findErr, "DatabaseError")
	}

	c.EM.EmitWithData(events.FindQuerySuccess, ctx, nil, "Users found successfully", result)
	return ctx.JSON(result)
}

// @Summary      Create a new Users
// @Description  Adds a new Users to the database. Auto-generated fields like ID, created_at will be populated automatically.
// @Tags         Users
// @Id           IamUsersCreate
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        request        body    structures.UsersForm true "The Users data to create (exclude auto-generated fields)"
// @Success      201            {object}  structures.UsersForm "Users created successfully"
// @Failure      400            {object}  web.Fiber400Response "Invalid input data or validation error"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      409            {object}  web.Fiber400Response "Users already exists (duplicate key)"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /users [post]
func (c *UsersController) CreateUsers(ctx *fiber.Ctx) error {

	// Idempotency key check — prevents duplicate creates across pod retries.
	// If X-Idempotency-Key header is present, the caller guarantees uniqueness.
	// Downstream (cache/db) can deduplicate on this key.
	_ = ctx.Get("X-Idempotency-Key") // read; enforcement is done at middleware/cache layer

	data := new(structures.UsersForm)
	if err := ctx.BodyParser(data); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "CREATE Users: Failed to parse request body")
		return handleUsersControllerError(ctx, fiber.StatusBadRequest, err, "ValidationError")
	}

	// Emit creation request event
	proceed, reqErr := c.EM.Emit(events.CreationRequest, ctx, nil, fmt.Sprintf("CREATE Users request: %+v", data))
	if !proceed {
		c.EM.Emit(events.CreationError, ctx, reqErr, fmt.Sprintf("CREATE Users: Request rejected by handler"))
		return handleUsersControllerError(ctx, fiber.StatusBadRequest, reqErr, "ValidationError")
	}

	// No system-required fields to set automatically.

	// Execute creation with service layer (service handles timeout)
	result, serviceErr := c.Svc.CreateUsers(*data)

	if serviceErr != nil {
		c.EM.Emit(events.CreationError, ctx, serviceErr, fmt.Sprintf("CREATE Users: Service error - %+v", data))

		// Check if it's a timeout error
		if errors.Is(serviceErr, context.DeadlineExceeded) {
			return handleUsersControllerError(ctx, fiber.StatusRequestTimeout, serviceErr, "TimeoutError")
		}

		// Check for database constraint errors
		if strings.Contains(serviceErr.Error(), "duplicate key") {
			return handleUsersControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		if strings.Contains(serviceErr.Error(), "foreign key constraint") {
			return handleUsersControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		return handleUsersControllerError(ctx, fiber.StatusInternalServerError, serviceErr, "DatabaseError")
	}

	c.EM.EmitWithData(events.CreationSuccess, ctx, nil, fmt.Sprintf("CREATE Users success"), result)
	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// @Summary      Create multiple Users
// @Description  Adds multiple Users to the database in a single request.
// @Tags         Users
// @Id           IamUsersCreateMultiple
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]])
// @Param        x-company      header  string  true  "Company ID" default([[companyID]])
// @Param        request        body    []structures.UsersForm true "A list of Users to create"
// @Success      200            {object}  []structures.UsersForm
// @Failure      400            {object}  web.Fiber400Response
// @Router       /users/bulk [post]
func (c *UsersController) CreateUsersMultiple(ctx *fiber.Ctx) error {
	var bulkData []structures.UsersForm
	if err := ctx.BodyParser(&bulkData); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "BATCH CREATE Users: Failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: err.Error(), Type: "BodyParseError"})
	}
	if len(bulkData) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Request body is an empty list.", Type: "ValidationError"})
	}
	proceed, reqErr := c.EM.Emit(events.BatchCreationRequest, ctx, nil, fmt.Sprintf("BATCH CREATE Users request with %d items", len(bulkData)))
	if !proceed {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: reqErr.Error(), Type: "RequestRejectedByHandler"})
	}
	// No system-required fields to set automatically.
	result, serviceErr := c.Svc.CreateUsersMultiple(bulkData)
	if serviceErr != nil {
		c.EM.Emit(events.BatchCreationError, ctx, serviceErr, fmt.Sprintf("BATCH CREATE Users: Service error - %d items", len(bulkData)))
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: serviceErr.Error(), Type: "BatchCreationError"})
	}
	c.EM.EmitWithData(events.BatchCreationSuccess, ctx, nil, fmt.Sprintf("BATCH CREATE Users: %d items created successfully", len(result)), result)
	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// @Summary      Update a Users
// @Description  Updates an existing Users by its unique identifier(s). Only provided fields will be updated (partial update).
// @Tags         Users
// @Id           IamUsersUpdateWithID
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        id   path    types.URID   true  "Id" example(123)
// @Param        request        body    structures.UsersEdit true "Fields to update (only include fields you want to change)"
// @Success      200            {object}  web.Fiber200Response "Users updated successfully"
// @Failure      400            {object}  web.Fiber400Response "Invalid input data or validation error"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      404            {object}  web.Fiber400Response "Users not found"
// @Failure      409            {object}  web.Fiber400Response "Constraint violation (duplicate key, etc.)"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /users/with-id/:id [put]
func (c *UsersController) UpdateUsers(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")
	var id types.URID
	if err := id.FromBase32(idStr); err != nil { // Veya FromUUID, duruma göre değişebilir
		c.EM.Emit(events.QueryParamsParsingError, ctx, err, "Invalid 'id' format in path parameter")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid " + "id" + " format", Type: "QueryParamsParsingError"})
	}
	editData := new(structures.UsersEdit)
	if err := ctx.BodyParser(editData); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "UPDATE Users: Failed to parse request body")
		return handleUsersControllerError(ctx, fiber.StatusBadRequest, err, "ValidationError")
	}

	// Emit update request event
	proceed, reqErr := c.EM.Emit(events.UpdateRequest, ctx, nil, fmt.Sprintf("UPDATE Users request for ID(s) with data: %+v", editData))
	if !proceed {
		c.EM.Emit(events.UpdateError, ctx, reqErr, fmt.Sprintf("UPDATE Users: Request rejected by handler"))
		return handleUsersControllerError(ctx, fiber.StatusBadRequest, reqErr, "ValidationError")
	}

	// No system-required fields to set automatically.

	// Execute update with service layer (service handles timeout)
	serviceErr := c.Svc.UpdateUsers(id, *editData)

	if serviceErr != nil {
		// Check if it's a timeout error
		if errors.Is(serviceErr, context.DeadlineExceeded) {
			c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE Users: Timeout error for ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusRequestTimeout, serviceErr, "TimeoutError")
		}

		// Check for record not found
		if errors.Is(serviceErr, gorm.ErrRecordNotFound) {
			c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE Users: Record not found for ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusNotFound, serviceErr, "NotFoundError")
		}

		// Check for database constraint errors
		if strings.Contains(serviceErr.Error(), "duplicate key") {
			c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE Users: Duplicate key error for ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		if strings.Contains(serviceErr.Error(), "foreign key constraint") {
			c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE Users: Foreign key constraint error for ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE Users: Service error for ID(s) - Data: %+v", editData))
		return handleUsersControllerError(ctx, fiber.StatusInternalServerError, serviceErr, "DatabaseError")
	}

	c.EM.Emit(events.UpdateSuccess, ctx, nil, fmt.Sprintf("UPDATE Users success for ID(s)"))
	return ctx.Status(fiber.StatusOK).JSON(web.Fiber200Response{Message: "Users updated successfully", Type: "UpdateSuccess"})
}

// @Summary      Update multiple Users
// @Description  Updates multiple Users resources based on a list of their identifiers and new data.
// @Tags         Users
// @Id           IamUsersUpdateMultiple
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]])
// @Param        x-company      header  string  true  "Company ID" default([[companyID]])
// @Param        request        body    []structures.UsersBatchUpdate true "A list of Users identifiers and their update data"
// @Success      200            {object}  web.Fiber200Response{message=string}
// @Failure      400            {object}  web.Fiber400Response
// @Failure      500            {object}  web.Fiber400Response
// @Router       /users/bulk [put]
func (c *UsersController) UpdateUsersMultiple(ctx *fiber.Ctx) error {
	var batchUpdateData []structures.UsersBatchUpdate
	if err := ctx.BodyParser(&batchUpdateData); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "BATCH UPDATE Users: Failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body: " + err.Error(), Type: "BodyParseError"})
	}
	if len(batchUpdateData) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Request body is an empty list for batch update.", Type: "ValidationError"})
	}
	proceed, reqErr := c.EM.Emit(events.BatchUpdateRequest, ctx, nil, fmt.Sprintf("BATCH UPDATE Users request with %d items", len(batchUpdateData)))
	if !proceed {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: reqErr.Error(), Type: "RequestRejectedByHandler"})
	}
	serviceErr := c.Svc.UpdateUsersMultiple(batchUpdateData)
	if serviceErr != nil {
		c.EM.Emit(events.BatchUpdateError, ctx, serviceErr, fmt.Sprintf("BATCH UPDATE Users: Service error - %d items", len(batchUpdateData)))
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "Failed to batch update records: " + serviceErr.Error(), Type: "ServiceError"})
	}
	c.EM.EmitWithData(events.BatchUpdateSuccess, ctx, nil, fmt.Sprintf("BATCH UPDATE Users: %d records processed successfully", len(batchUpdateData)), batchUpdateData)
	return ctx.Status(fiber.StatusOK).JSON(web.Fiber200Response{Message: fmt.Sprintf("%d records processed for batch update.", len(batchUpdateData)), Type: "BatchUpdateSuccess"})
}

// @Summary      Delete a Users
// @Description  Permanently deletes a Users by its unique identifier(s). This action cannot be undone.
// @Tags         Users
// @Id           IamUsersDeleteWithID
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        id   path    types.URID   true  "Id" example(123)
// @Success      200            {object}  web.Fiber200Response "Users deleted successfully"
// @Failure      400            {object}  web.Fiber400Response "Invalid ID format"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      404            {object}  web.Fiber400Response "Users not found"
// @Failure      409            {object}  web.Fiber400Response "Cannot delete: referenced by other records"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /users/with-id/:id [delete]
func (c *UsersController) DeleteUsers(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")
	var id types.URID
	if err := id.FromBase32(idStr); err != nil { // Veya FromUUID, duruma göre değişebilir
		c.EM.Emit(events.QueryParamsParsingError, ctx, err, "Invalid 'id' format in path parameter")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid " + "id" + " format", Type: "QueryParamsParsingError"})
	}

	// Emit deletion request event
	proceed, reqErr := c.EM.Emit(events.DeletionRequest, ctx, nil, fmt.Sprintf("DELETE Users request for ID(s)"))
	if !proceed {
		c.EM.Emit(events.DeletionError, ctx, reqErr, fmt.Sprintf("DELETE Users: Request rejected by handler"))
		return handleUsersControllerError(ctx, fiber.StatusBadRequest, reqErr, "ValidationError")
	}

	// Execute deletion with service layer (service handles timeout)
	serviceErr := c.Svc.DeleteUsers(id)

	if serviceErr != nil {
		// Check if it's a timeout error
		if errors.Is(serviceErr, context.DeadlineExceeded) {
			c.EM.Emit(events.DeletionError, ctx, serviceErr, fmt.Sprintf("DELETE Users: Timeout error for ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusRequestTimeout, serviceErr, "TimeoutError")
		}

		// Check for record not found
		if errors.Is(serviceErr, gorm.ErrRecordNotFound) {
			c.EM.Emit(events.DeletionError, ctx, serviceErr, fmt.Sprintf("DELETE Users: Record not found for ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusNotFound, serviceErr, "NotFoundError")
		}

		// Check for foreign key constraint errors (cascade delete issues)
		if strings.Contains(serviceErr.Error(), "foreign key constraint") {
			c.EM.Emit(events.DeletionError, ctx, serviceErr, fmt.Sprintf("DELETE Users: Foreign key constraint error for ID(s)"))
			return handleUsersControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		c.EM.Emit(events.DeletionError, ctx, serviceErr, fmt.Sprintf("DELETE Users: Service error for ID(s)"))
		return handleUsersControllerError(ctx, fiber.StatusInternalServerError, serviceErr, "DatabaseError")
	}

	c.EM.Emit(events.DeletionSuccess, ctx, nil, fmt.Sprintf("DELETE Users success for ID(s)"))
	return ctx.Status(fiber.StatusOK).JSON(web.Fiber200Response{Message: "Users deleted successfully", Type: "DeleteSuccess"})
}

// @Summary      Delete multiple Users
// @Description  Deletes multiple Users resources based on a list of their identifiers.
// @Tags         Users
// @Id           IamUsersDeleteMultiple
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]])
// @Param        x-company      header  string  true  "Company ID" default([[companyID]])
// @Param        request        body    []structures.UsersIdentity true "A list of Users identifiers to delete"
// @Success      200            {object}  web.Fiber200Response{message=string}
// @Failure      400            {object}  web.Fiber400Response
// @Failure      500            {object}  web.Fiber400Response
// @Router       /users/bulk [delete]
func (c *UsersController) DeleteUsersMultiple(ctx *fiber.Ctx) error {
	var identities []structures.UsersIdentity
	if err := ctx.BodyParser(&identities); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "BATCH DELETE Users: Failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body: " + err.Error(), Type: "BodyParseError"})
	}
	if len(identities) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "No identities provided for batch delete.", Type: "ValidationError"})
	}
	proceed, reqErr := c.EM.Emit(events.BatchDeletionRequest, ctx, nil, fmt.Sprintf("BATCH DELETE Users request for %d identities", len(identities)))
	if !proceed {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: reqErr.Error(), Type: "RequestRejectedByHandler"})
	}
	serviceErr := c.Svc.DeleteUsersMultiple(identities)
	if serviceErr != nil {
		c.EM.Emit(events.BatchDeletionError, ctx, serviceErr, fmt.Sprintf("BATCH DELETE Users: Service error - %d identities", len(identities)))
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "Failed to delete records: " + serviceErr.Error(), Type: "ServiceError"})
	}
	c.EM.Emit(events.BatchDeletionSuccess, ctx, nil, fmt.Sprintf("BATCH DELETE Users: %d records deleted successfully", len(identities)))
	return ctx.Status(fiber.StatusOK).JSON(web.Fiber200Response{Message: fmt.Sprintf("%d records deleted successfully", len(identities)), Type: "BatchDeleteSuccess"})
}
