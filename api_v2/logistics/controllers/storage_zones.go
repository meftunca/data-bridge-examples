package logistics_api_controller

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	paginationRuntime "backend-generator/apiv2/pagination"
	services "data-bridge-examples/api_v2/logistics/services"
	structures "data-bridge-examples/api_v2/logistics/structures"

	"github.com/maple-tech/baseline/events"
	"github.com/maple-tech/baseline/types"
	"github.com/maple-tech/baseline/web"
)

// paginationStorageZones provides advanced filtering, sorting, and pagination
// capabilities for StorageZones resources. Pre-configured with schema analysis
// for optimal query performance and type-safe operations.
var paginationStorageZones = paginationRuntime.NewPagination[structures.StorageZones](structures.StorageZones{})

var (
	_ = types.ID(0)
	_ = types.URID("")
)

func handleStorageZonesControllerError(ctx *fiber.Ctx, status int, err error, errType string) error {
	if err == nil {
		err = errors.New(strings.ToLower(errType))
	}
	return ctx.Status(status).JSON(web.Fiber400Response{Message: err.Error(), Type: errType})
}

// StorageZonesController handles HTTP requests for StorageZones resources.
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
type StorageZonesController struct {
	Svc services.IStorageZonesService // Business logic and data operations interface
	EM  *events.EventManager          // Request lifecycle and monitoring events
}

// @Summary      Search StorageZones
// @Description  Searches and filters StorageZones resources without pagination. Returns a simple array of results.
// @Tags         logistics
// @Id           LogisticsStorageZonesSearch
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        size           query   int     false "Maximum number of results (1-1000)" default(100) example(50)
// @Param        sort           query   string  false "Sort fields with direction: field1,-field2 (- for DESC)" example(-created_at,name)
// @Param        filters        query   string  false "JSON array: [[\"field\",\"operator\",\"value\"]] - Operators: =,!=,>,>=,<,<=,like,in" example([["name","like","%test%"],["status","=","active"]])
// @Param        preloads       query   string  false "Comma-separated relation names to include" example(user,category,tags)
// @Success      200            {array}   structures.StorageZones "List of StorageZones resources"
// @Failure      400            {object}  web.Fiber400Response "Invalid request parameters"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /logistics/storage-zones [get]
func (c *StorageZonesController) SearchStorageZones(ctx *fiber.Ctx) error {
	// Context with timeout for this operation
	timeoutCtx, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	_, err := c.EM.Emit(events.QueryRequest, ctx, nil, "Query request for StorageZones")
	if err != nil {
		return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, err, "ValidationError")
	}

	dbForSearch := c.Svc.(*services.StorageZonesService).DB

	// Create a channel to handle timeout
	type searchResult struct {
		response []structures.StorageZones
		err      error
	}

	resultChan := make(chan searchResult, 1)

	go func() {
		response, searchErr := paginationStorageZones.With(dbForSearch.Model(&structures.StorageZones{})).Request(ctx).ResponseTiny()
		resultChan <- searchResult{response: response, err: searchErr}
	}()

	select {
	case result := <-resultChan:
		if result.err != nil {
			c.EM.Emit(events.QueryError, ctx, result.err, "Failed to query StorageZones")
			// Enhanced error handling - check for specific database errors
			if errors.Is(result.err, gorm.ErrRecordNotFound) {
				return handleStorageZonesControllerError(ctx, fiber.StatusNotFound, result.err, "NotFoundError")
			}
			return handleStorageZonesControllerError(ctx, fiber.StatusInternalServerError, result.err, "DatabaseError")
		}
		c.EM.EmitWithData(events.QuerySuccess, ctx, nil, fmt.Sprintf("Successfully queried StorageZones, count: %d", len(result.response)), result.response)
		return ctx.JSON(result.response)

	case <-timeoutCtx.Done():
		timeoutErr := timeoutCtx.Err()
		c.EM.Emit(events.QueryError, ctx, timeoutErr, "Query timeout for StorageZones")
		return handleStorageZonesControllerError(ctx, fiber.StatusRequestTimeout, timeoutErr, "TimeoutError")
	}
}

// @Summary      List StorageZones with Pagination
// @Description  Fetches a paginated list of StorageZones resources with metadata including total count, current page, and pagination links.
// @Tags         logistics
// @Id           LogisticsStorageZonesPagination
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        page           query   int     false "Page number (starts from 1)" default(1) example(2)
// @Param        size           query   int     false "Number of items per page (1-1000)" default(10) example(25)
// @Param        sort           query   string  false "Sort fields with direction: field1,-field2 (- for DESC)" example(-created_at,name)
// @Param        filters        query   string  false "JSON array: [[\"field\",\"operator\",\"value\"]] - Operators: =,!=,>,>=,<,<=,like,in" example([["status","=","active"],["created_at",">=","2024-01-01"]])
// @Param        preloads       query   string  false "Comma-separated relation names to include" example(user,category)
// @Success      200            {object}  structures.StorageZonesPage "Paginated StorageZones resources with metadata"
// @Failure      400            {object}  web.Fiber400Response "Invalid pagination parameters"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /logistics/storage-zones/pagination [get]
func (c *StorageZonesController) GetStorageZonesWithPagination(ctx *fiber.Ctx) error {
	_, err := c.EM.Emit(events.PaginationQueryRequest, ctx, nil, "Pagination query request for StorageZones")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: err.Error(), Type: "RequestRejectedByHandler"})
	}
	dbForPagination := c.Svc.(*services.StorageZonesService).DB
	request := paginationStorageZones.With(dbForPagination.Model(&structures.StorageZones{})).Request(ctx)

	if ctx.Query("collectionBy") != "" {
		response, pagErr := request.ResponsePaginatedCollection()
		if pagErr != nil {
			c.EM.Emit(events.PaginationQueryError, ctx, pagErr, "Failed to retrieve grouped StorageZones collection")
			return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: pagErr.Error(), Type: "PaginationError"})
		}
		c.EM.EmitWithData(events.PaginationQuerySuccess, ctx, nil, fmt.Sprintf("Successfully retrieved grouped StorageZones collection, count: %d", len(response.Items)), response)
		return ctx.JSON(response)
	}

	if ctx.Query("groupby") != "" || ctx.Query("aggregations") != "" {
		response, pagErr := request.ResponseRaw()
		if pagErr != nil {
			c.EM.Emit(events.PaginationQueryError, ctx, pagErr, "Failed to retrieve aggregated StorageZones response")
			return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: pagErr.Error(), Type: "PaginationError"})
		}
		c.EM.EmitWithData(events.PaginationQuerySuccess, ctx, nil, fmt.Sprintf("Successfully retrieved aggregated StorageZones response, count: %d", len(response.Items)), response)
		return ctx.JSON(response)
	}

	response, pagErr := request.Response()
	if pagErr != nil {
		c.EM.Emit(events.PaginationQueryError, ctx, pagErr, "Failed to retrieve paginated StorageZones")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: pagErr.Error(), Type: "PaginationError"})
	}
	c.EM.EmitWithData(events.PaginationQuerySuccess, ctx, nil, fmt.Sprintf("Successfully retrieved paginated StorageZones, count: %d", len(response.Items)), response)
	return ctx.JSON(response)
}

// @Summary      Get a StorageZones by ID
// @Description  Fetches a single StorageZones resource by its unique identifier(s). Returns the complete resource data.
// @Tags         logistics
// @Id           LogisticsStorageZonesWithID
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        id   path    types.URID   true  "Id" example(123)
// @Param        preloads       query   string  false "Comma-separated relation names to include" example(user,category)
// @Success      200            {object}  structures.StorageZones "StorageZones resource found"
// @Failure      400            {object}  web.Fiber400Response "Invalid ID format"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      404            {object}  web.Fiber400Response "StorageZones not found"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /logistics/storage-zones/with-id/:id [get]
func (c *StorageZonesController) GetStorageZonesById(ctx *fiber.Ctx) error {
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
	proceed, reqErr := c.EM.Emit(events.FindQueryRequest, ctx, nil, fmt.Sprintf("FindQuery request for StorageZones with ID(s) from path"))
	if !proceed {
		c.EM.Emit(events.FindQueryError, ctx, reqErr, fmt.Sprintf("StorageZones find request rejected by handler"))
		return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, reqErr, "ValidationError")
	}

	// Execute find with timeout context
	dbForFind := c.Svc.(*services.StorageZonesService).DB.WithContext(timeoutCtx)
	queryBuilder := dbForFind.Model(&structures.StorageZones{})
	queryBuilder = queryBuilder.Where("id = ?", id)
	var result *structures.StorageZones
	findErr := queryBuilder.First(&result).Error

	if findErr != nil {
		// Check if it's a timeout error
		if errors.Is(findErr, context.DeadlineExceeded) {
			c.EM.Emit(events.FindQueryError, ctx, findErr, fmt.Sprintf("StorageZones find timeout with ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusRequestTimeout, findErr, "TimeoutError")
		}

		// Check for record not found
		if errors.Is(findErr, gorm.ErrRecordNotFound) {
			c.EM.Emit(events.FindQueryError, ctx, findErr, fmt.Sprintf("StorageZones not found with provided ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusNotFound, findErr, "NotFoundError")
		}

		c.EM.Emit(events.FindQueryError, ctx, findErr, fmt.Sprintf("Error finding StorageZones with ID(s)"))
		return handleStorageZonesControllerError(ctx, fiber.StatusInternalServerError, findErr, "DatabaseError")
	}

	c.EM.EmitWithData(events.FindQuerySuccess, ctx, nil, "StorageZones found successfully", result)
	return ctx.JSON(result)
}

// @Summary      Create a new StorageZones
// @Description  Adds a new StorageZones to the database. Auto-generated fields like ID, created_at will be populated automatically.
// @Tags         logistics
// @Id           LogisticsStorageZonesCreate
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        request        body    structures.StorageZonesForm true "The StorageZones data to create (exclude auto-generated fields)"
// @Success      201            {object}  structures.StorageZonesForm "StorageZones created successfully"
// @Failure      400            {object}  web.Fiber400Response "Invalid input data or validation error"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      409            {object}  web.Fiber400Response "StorageZones already exists (duplicate key)"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /logistics/storage-zones [post]
func (c *StorageZonesController) CreateStorageZones(ctx *fiber.Ctx) error {

	// Idempotency key check — prevents duplicate creates across pod retries.
	// If X-Idempotency-Key header is present, the caller guarantees uniqueness.
	// Downstream (cache/db) can deduplicate on this key.
	_ = ctx.Get("X-Idempotency-Key") // read; enforcement is done at middleware/cache layer

	data := new(structures.StorageZonesForm)
	if err := ctx.BodyParser(data); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "CREATE StorageZones: Failed to parse request body")
		return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, err, "ValidationError")
	}

	// Emit creation request event
	proceed, reqErr := c.EM.Emit(events.CreationRequest, ctx, nil, fmt.Sprintf("CREATE StorageZones request: %+v", data))
	if !proceed {
		c.EM.Emit(events.CreationError, ctx, reqErr, fmt.Sprintf("CREATE StorageZones: Request rejected by handler"))
		return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, reqErr, "ValidationError")
	}

	// No system-required fields to set automatically.

	// Execute creation with service layer (service handles timeout)
	result, serviceErr := c.Svc.CreateStorageZones(*data)

	if serviceErr != nil {
		c.EM.Emit(events.CreationError, ctx, serviceErr, fmt.Sprintf("CREATE StorageZones: Service error - %+v", data))

		// Check if it's a timeout error
		if errors.Is(serviceErr, context.DeadlineExceeded) {
			return handleStorageZonesControllerError(ctx, fiber.StatusRequestTimeout, serviceErr, "TimeoutError")
		}

		// Check for database constraint errors
		if strings.Contains(serviceErr.Error(), "duplicate key") {
			return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		if strings.Contains(serviceErr.Error(), "foreign key constraint") {
			return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		return handleStorageZonesControllerError(ctx, fiber.StatusInternalServerError, serviceErr, "DatabaseError")
	}

	c.EM.EmitWithData(events.CreationSuccess, ctx, nil, fmt.Sprintf("CREATE StorageZones success"), result)
	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// @Summary      Create multiple StorageZones
// @Description  Adds multiple StorageZones to the database in a single request.
// @Tags         logistics
// @Id           LogisticsStorageZonesCreateMultiple
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]])
// @Param        x-company      header  string  true  "Company ID" default([[companyID]])
// @Param        request        body    []structures.StorageZonesForm true "A list of StorageZones to create"
// @Success      200            {object}  []structures.StorageZonesForm
// @Failure      400            {object}  web.Fiber400Response
// @Router       /logistics/storage-zones/bulk [post]
func (c *StorageZonesController) CreateStorageZonesMultiple(ctx *fiber.Ctx) error {
	var bulkData []structures.StorageZonesForm
	if err := ctx.BodyParser(&bulkData); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "BATCH CREATE StorageZones: Failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: err.Error(), Type: "BodyParseError"})
	}
	if len(bulkData) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Request body is an empty list.", Type: "ValidationError"})
	}
	proceed, reqErr := c.EM.Emit(events.BatchCreationRequest, ctx, nil, fmt.Sprintf("BATCH CREATE StorageZones request with %d items", len(bulkData)))
	if !proceed {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: reqErr.Error(), Type: "RequestRejectedByHandler"})
	}
	// No system-required fields to set automatically.
	result, serviceErr := c.Svc.CreateStorageZonesMultiple(bulkData)
	if serviceErr != nil {
		c.EM.Emit(events.BatchCreationError, ctx, serviceErr, fmt.Sprintf("BATCH CREATE StorageZones: Service error - %d items", len(bulkData)))
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: serviceErr.Error(), Type: "BatchCreationError"})
	}
	c.EM.EmitWithData(events.BatchCreationSuccess, ctx, nil, fmt.Sprintf("BATCH CREATE StorageZones: %d items created successfully", len(result)), result)
	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// @Summary      Update a StorageZones
// @Description  Updates an existing StorageZones by its unique identifier(s). Only provided fields will be updated (partial update).
// @Tags         logistics
// @Id           LogisticsStorageZonesUpdateWithID
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        id   path    types.URID   true  "Id" example(123)
// @Param        request        body    structures.StorageZonesEdit true "Fields to update (only include fields you want to change)"
// @Success      200            {object}  web.Fiber200Response "StorageZones updated successfully"
// @Failure      400            {object}  web.Fiber400Response "Invalid input data or validation error"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      404            {object}  web.Fiber400Response "StorageZones not found"
// @Failure      409            {object}  web.Fiber400Response "Constraint violation (duplicate key, etc.)"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /logistics/storage-zones/with-id/:id [put]
func (c *StorageZonesController) UpdateStorageZones(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")
	var id types.URID
	if err := id.FromBase32(idStr); err != nil { // Veya FromUUID, duruma göre değişebilir
		c.EM.Emit(events.QueryParamsParsingError, ctx, err, "Invalid 'id' format in path parameter")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid " + "id" + " format", Type: "QueryParamsParsingError"})
	}
	editData := new(structures.StorageZonesEdit)
	if err := ctx.BodyParser(editData); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "UPDATE StorageZones: Failed to parse request body")
		return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, err, "ValidationError")
	}

	// Emit update request event
	proceed, reqErr := c.EM.Emit(events.UpdateRequest, ctx, nil, fmt.Sprintf("UPDATE StorageZones request for ID(s) with data: %+v", editData))
	if !proceed {
		c.EM.Emit(events.UpdateError, ctx, reqErr, fmt.Sprintf("UPDATE StorageZones: Request rejected by handler"))
		return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, reqErr, "ValidationError")
	}

	// No system-required fields to set automatically.

	// Execute update with service layer (service handles timeout)
	serviceErr := c.Svc.UpdateStorageZones(id, *editData)

	if serviceErr != nil {
		// Check if it's a timeout error
		if errors.Is(serviceErr, context.DeadlineExceeded) {
			c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE StorageZones: Timeout error for ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusRequestTimeout, serviceErr, "TimeoutError")
		}

		// Check for record not found
		if errors.Is(serviceErr, gorm.ErrRecordNotFound) {
			c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE StorageZones: Record not found for ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusNotFound, serviceErr, "NotFoundError")
		}

		// Check for database constraint errors
		if strings.Contains(serviceErr.Error(), "duplicate key") {
			c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE StorageZones: Duplicate key error for ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		if strings.Contains(serviceErr.Error(), "foreign key constraint") {
			c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE StorageZones: Foreign key constraint error for ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		c.EM.Emit(events.UpdateError, ctx, serviceErr, fmt.Sprintf("UPDATE StorageZones: Service error for ID(s) - Data: %+v", editData))
		return handleStorageZonesControllerError(ctx, fiber.StatusInternalServerError, serviceErr, "DatabaseError")
	}

	c.EM.Emit(events.UpdateSuccess, ctx, nil, fmt.Sprintf("UPDATE StorageZones success for ID(s)"))
	return ctx.Status(fiber.StatusOK).JSON(web.Fiber200Response{Message: "StorageZones updated successfully", Type: "UpdateSuccess"})
}

// @Summary      Update multiple StorageZones
// @Description  Updates multiple StorageZones resources based on a list of their identifiers and new data.
// @Tags         logistics
// @Id           LogisticsStorageZonesUpdateMultiple
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]])
// @Param        x-company      header  string  true  "Company ID" default([[companyID]])
// @Param        request        body    []structures.StorageZonesBatchUpdate true "A list of StorageZones identifiers and their update data"
// @Success      200            {object}  web.Fiber200Response{message=string}
// @Failure      400            {object}  web.Fiber400Response
// @Failure      500            {object}  web.Fiber400Response
// @Router       /logistics/storage-zones/bulk [put]
func (c *StorageZonesController) UpdateStorageZonesMultiple(ctx *fiber.Ctx) error {
	var batchUpdateData []structures.StorageZonesBatchUpdate
	if err := ctx.BodyParser(&batchUpdateData); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "BATCH UPDATE StorageZones: Failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body: " + err.Error(), Type: "BodyParseError"})
	}
	if len(batchUpdateData) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Request body is an empty list for batch update.", Type: "ValidationError"})
	}
	proceed, reqErr := c.EM.Emit(events.BatchUpdateRequest, ctx, nil, fmt.Sprintf("BATCH UPDATE StorageZones request with %d items", len(batchUpdateData)))
	if !proceed {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: reqErr.Error(), Type: "RequestRejectedByHandler"})
	}
	serviceErr := c.Svc.UpdateStorageZonesMultiple(batchUpdateData)
	if serviceErr != nil {
		c.EM.Emit(events.BatchUpdateError, ctx, serviceErr, fmt.Sprintf("BATCH UPDATE StorageZones: Service error - %d items", len(batchUpdateData)))
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "Failed to batch update records: " + serviceErr.Error(), Type: "ServiceError"})
	}
	c.EM.EmitWithData(events.BatchUpdateSuccess, ctx, nil, fmt.Sprintf("BATCH UPDATE StorageZones: %d records processed successfully", len(batchUpdateData)), batchUpdateData)
	return ctx.Status(fiber.StatusOK).JSON(web.Fiber200Response{Message: fmt.Sprintf("%d records processed for batch update.", len(batchUpdateData)), Type: "BatchUpdateSuccess"})
}

// @Summary      Delete a StorageZones
// @Description  Permanently deletes a StorageZones by its unique identifier(s). This action cannot be undone.
// @Tags         logistics
// @Id           LogisticsStorageZonesDeleteWithID
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]]) example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...)
// @Param        x-company      header  string  true  "Company ID" default([[companyID]]) example(comp_123456789)
// @Param        id   path    types.URID   true  "Id" example(123)
// @Success      200            {object}  web.Fiber200Response "StorageZones deleted successfully"
// @Failure      400            {object}  web.Fiber400Response "Invalid ID format"
// @Failure      401            {object}  web.Fiber400Response "Authentication required"
// @Failure      403            {object}  web.Fiber400Response "Access denied"
// @Failure      404            {object}  web.Fiber400Response "StorageZones not found"
// @Failure      409            {object}  web.Fiber400Response "Cannot delete: referenced by other records"
// @Failure      408            {object}  web.Fiber400Response "Request timeout"
// @Failure      500            {object}  web.Fiber400Response "Internal server error"
// @Router       /logistics/storage-zones/with-id/:id [delete]
func (c *StorageZonesController) DeleteStorageZones(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")
	var id types.URID
	if err := id.FromBase32(idStr); err != nil { // Veya FromUUID, duruma göre değişebilir
		c.EM.Emit(events.QueryParamsParsingError, ctx, err, "Invalid 'id' format in path parameter")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid " + "id" + " format", Type: "QueryParamsParsingError"})
	}

	// Emit deletion request event
	proceed, reqErr := c.EM.Emit(events.DeletionRequest, ctx, nil, fmt.Sprintf("DELETE StorageZones request for ID(s)"))
	if !proceed {
		c.EM.Emit(events.DeletionError, ctx, reqErr, fmt.Sprintf("DELETE StorageZones: Request rejected by handler"))
		return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, reqErr, "ValidationError")
	}

	// Execute deletion with service layer (service handles timeout)
	serviceErr := c.Svc.DeleteStorageZones(id)

	if serviceErr != nil {
		// Check if it's a timeout error
		if errors.Is(serviceErr, context.DeadlineExceeded) {
			c.EM.Emit(events.DeletionError, ctx, serviceErr, fmt.Sprintf("DELETE StorageZones: Timeout error for ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusRequestTimeout, serviceErr, "TimeoutError")
		}

		// Check for record not found
		if errors.Is(serviceErr, gorm.ErrRecordNotFound) {
			c.EM.Emit(events.DeletionError, ctx, serviceErr, fmt.Sprintf("DELETE StorageZones: Record not found for ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusNotFound, serviceErr, "NotFoundError")
		}

		// Check for foreign key constraint errors (cascade delete issues)
		if strings.Contains(serviceErr.Error(), "foreign key constraint") {
			c.EM.Emit(events.DeletionError, ctx, serviceErr, fmt.Sprintf("DELETE StorageZones: Foreign key constraint error for ID(s)"))
			return handleStorageZonesControllerError(ctx, fiber.StatusBadRequest, serviceErr, "ValidationError")
		}

		c.EM.Emit(events.DeletionError, ctx, serviceErr, fmt.Sprintf("DELETE StorageZones: Service error for ID(s)"))
		return handleStorageZonesControllerError(ctx, fiber.StatusInternalServerError, serviceErr, "DatabaseError")
	}

	c.EM.Emit(events.DeletionSuccess, ctx, nil, fmt.Sprintf("DELETE StorageZones success for ID(s)"))
	return ctx.Status(fiber.StatusOK).JSON(web.Fiber200Response{Message: "StorageZones deleted successfully", Type: "DeleteSuccess"})
}

// @Summary      Delete multiple StorageZones
// @Description  Deletes multiple StorageZones resources based on a list of their identifiers.
// @Tags         logistics
// @Id           LogisticsStorageZonesDeleteMultiple
// @Param        authorization  header  string  true  "Authentication token" default([[sessionToken]])
// @Param        x-company      header  string  true  "Company ID" default([[companyID]])
// @Param        request        body    []structures.StorageZonesIdentity true "A list of StorageZones identifiers to delete"
// @Success      200            {object}  web.Fiber200Response{message=string}
// @Failure      400            {object}  web.Fiber400Response
// @Failure      500            {object}  web.Fiber400Response
// @Router       /logistics/storage-zones/bulk [delete]
func (c *StorageZonesController) DeleteStorageZonesMultiple(ctx *fiber.Ctx) error {
	var identities []structures.StorageZonesIdentity
	if err := ctx.BodyParser(&identities); err != nil {
		c.EM.Emit(events.BodyParsingError, ctx, err, "BATCH DELETE StorageZones: Failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "Invalid request body: " + err.Error(), Type: "BodyParseError"})
	}
	if len(identities) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: "No identities provided for batch delete.", Type: "ValidationError"})
	}
	proceed, reqErr := c.EM.Emit(events.BatchDeletionRequest, ctx, nil, fmt.Sprintf("BATCH DELETE StorageZones request for %d identities", len(identities)))
	if !proceed {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Fiber400Response{Message: reqErr.Error(), Type: "RequestRejectedByHandler"})
	}
	serviceErr := c.Svc.DeleteStorageZonesMultiple(identities)
	if serviceErr != nil {
		c.EM.Emit(events.BatchDeletionError, ctx, serviceErr, fmt.Sprintf("BATCH DELETE StorageZones: Service error - %d identities", len(identities)))
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Fiber400Response{Message: "Failed to delete records: " + serviceErr.Error(), Type: "ServiceError"})
	}
	c.EM.Emit(events.BatchDeletionSuccess, ctx, nil, fmt.Sprintf("BATCH DELETE StorageZones: %d records deleted successfully", len(identities)))
	return ctx.Status(fiber.StatusOK).JSON(web.Fiber200Response{Message: fmt.Sprintf("%d records deleted successfully", len(identities)), Type: "BatchDeleteSuccess"})
}
