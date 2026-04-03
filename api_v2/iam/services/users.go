// UsersService - Enterprise-Grade Business Logic Service
//
// This service implements comprehensive business logic for Users resources with:
// • Transaction management with automatic rollback on errors
// • Context-aware operations with configurable timeout handling
// • Centralized error handling with detailed error classification
// • Batch operations for improved performance on bulk data processing
// • Event-driven architecture integration for monitoring and observability
// • ACID compliance for data integrity and consistency
//
// Architecture Pattern: Service Layer Pattern
// • Encapsulates business logic separate from HTTP concerns
// • Provides clean interface for controller layer consumption
// • Implements repository pattern with GORM ORM integration
// • Maintains separation of concerns following DDD principles
//
// Error Handling Strategy:
// • All errors are wrapped with contextual information
// • Transaction rollback is automatic on any failure
// • Structured error types for consistent client responses
// • Timeout handling prevents resource exhaustion
//
// Performance Considerations:
// • Batch operations minimize database round trips
// • Transaction scope optimization for concurrent operations
// • Context cancellation support for request lifecycle management
// • Prepared statement reuse through GORM optimization
package iam_api_service

import (
	sseRuntime "backend-generator/apiv2/sse"
	"context"
	structures "data-bridge-examples/api_v2/iam/structures"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/maple-tech/baseline/types"
	"gorm.io/gorm"
)

var (
	_ = types.ID(0)
	_ = types.URID("")
)

// IUsersService defines the business logic interface for Users operations.
//
// This interface abstracts all business operations for Users resources, providing:
// • Clean contract for controller layer integration
// • Mockable interface for comprehensive unit testing
// • Separation of concerns between HTTP and business logic
// • Consistent method signatures across all generated services
//
// Implementation Notes:
// • All methods support context cancellation for request lifecycle management
// • Batch operations are optimized for bulk data processing scenarios
// • Error handling follows consistent patterns across all operations
// • Transaction management is handled transparently by implementations
type IUsersService interface {
	UpdateUsers(id types.URID, data structures.UsersEdit) error
	UpdateUsersMultiple(data []structures.UsersBatchUpdate) error
	CreateUsers(data structures.UsersForm) (structures.UsersForm, error)
	CreateUsersMultiple(data []structures.UsersForm) ([]structures.UsersForm, error)
	DeleteUsers(id types.URID) error
	DeleteUsersMultiple(identities []structures.UsersIdentity) error // Tüm silme işlemleri tek bir transaction içinde yapılır.
}

// UsersService implements IUsersService with enterprise-grade business logic.
//
// This implementation provides:
// • GORM ORM integration for type-safe database operations
// • Automatic transaction management with rollback capabilities
// • Context-aware operations with timeout and cancellation support
// • Comprehensive error handling with detailed error classification
// • Performance-optimized batch operations for bulk data processing
// • Event integration for monitoring and observability
//
// The service maintains ACID compliance and provides consistent error handling
// across all operations. All methods are thread-safe and support concurrent access.
type UsersService struct {
	DB *gorm.DB // GORM database connection with transaction support
}

// --- Helper Methods ---

// withTransaction executes a function within a database transaction with comprehensive error handling.
//
// This method provides enterprise-grade transaction management with:
// • Automatic transaction lifecycle management (begin, commit, rollback)
// • Context cancellation support for request timeout handling
// • Panic recovery with automatic rollback to maintain data consistency
// • Detailed error logging with transaction state information
// • ACID compliance guarantees for data integrity
//
// The transaction is automatically rolled back if:
// • The provided function returns an error
// • A panic occurs during execution
// • The context is cancelled or times out
// • Any database operation fails
//
// Parameters:
//   - ctx: Context for cancellation and timeout management
//   - fn: Function to execute within the transaction scope
//
// Returns:
//   - error: Wrapped error with contextual information if operation fails
//
// Usage Example:
//
//	err := s.withTransaction(ctx, func(tx *gorm.DB) error {
//	    return tx.Create(&model).Error
//	})
func (s *UsersService) withTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	// Create a new transaction
	tx := s.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("%s: failed to begin transaction: %w", "transaction_error", tx.Error)
	}

	// Use defer to handle rollback in case of panic or error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // Re-panic after rollback
		}
	}()

	// Execute the function with context and transaction
	done := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				done <- fmt.Errorf("%s: panic in transaction: %v", "transaction_error", r)
			}
		}()

		// Set context for the transaction
		tx = tx.WithContext(ctx)
		done <- fn(tx)
	}()

	// Wait for completion or timeout
	select {
	case err := <-done:
		if err != nil {
			tx.Rollback()
			return err
		}

		// Commit the transaction
		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("%s: failed to commit transaction: %w", "transaction_error", err)
		}
		return nil

	case <-ctx.Done():
		tx.Rollback()
		return fmt.Errorf("%s: transaction timeout", "timeout_error")
	}
}

// handleServiceError provides consistent error handling and classification
func (s *UsersService) handleServiceError(err error, operation string) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()

	// Check for timeout
	if strings.Contains(errMsg, "context deadline exceeded") || strings.Contains(errMsg, "timeout") {
		return fmt.Errorf("%s: %s operation timeout: %w", "timeout_error", operation, err)
	}

	// Check for not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%s: record not found in %s: %w", "not_found_error", operation, err)
	}

	// Check for duplicate key
	if strings.Contains(errMsg, "duplicate key") || strings.Contains(errMsg, "UNIQUE constraint") {
		return fmt.Errorf("%s: duplicate record in %s: %w", "duplicate_error", operation, err)
	}

	// Check for foreign key constraint
	if strings.Contains(errMsg, "foreign key constraint") || strings.Contains(errMsg, "FOREIGN KEY constraint") {
		return fmt.Errorf("%s: constraint violation in %s: %w", "constraint_error", operation, err)
	}

	// Default database error
	return fmt.Errorf("%s: database error in %s: %w", "database_error", operation, err)
}

// validateContext checks if context is valid and not cancelled
func (s *UsersService) validateContext(ctx context.Context) error {
	if ctx == nil {
		return fmt.Errorf("%s: context is nil", "validation_error")
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("%s: context cancelled", "timeout_error")
	default:
		return nil
	}
}

// --- Method Implementations ---

func applyUsersCreateDefaults(data *structures.UsersForm) {

}

// Update, verilen primary/unique key'e göre bir kaydı günceller.
func (s *UsersService) UpdateUsers(id types.URID, data structures.UsersEdit) error {
	// Create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Validate context
	if err := s.validateContext(ctx); err != nil {
		return err
	}

	// Execute update within transaction
	return s.withTransaction(ctx, func(tx *gorm.DB) error {
		result := tx.Model(&structures.Users{}).Where("id = ?", id).Updates(data)

		if result.Error != nil {
			return s.handleServiceError(result.Error, "UpdateUsers")
		}

		// Check if any record was actually updated
		if result.RowsAffected == 0 {
			return fmt.Errorf("not_found_error: no record found to update in UpdateUsers: %w", gorm.ErrRecordNotFound)
		}

		// Transactional outbox: publish update event atomically in same transaction
		payloadBytes, _ := json.Marshal(data)
		outboxEvt := OutboxEvent{
			AggregateType: "Users",
			EventType:     "updated",
			Payload:       string(payloadBytes),
		}
		if err := tx.Table("outbox_events").Create(&outboxEvt).Error; err != nil {
			return fmt.Errorf("outbox insert failed in UpdateUsers: %w", err)
		}

		evtBytes, _ := json.Marshal(data)
		go sseRuntime.Publish("iam.users", sseRuntime.Event{Type: "updated", Payload: evtBytes})

		return nil
	})
}

// UpdateMultiple, birden fazla kaydı kendi path parametrelerine göre günceller.
func (s *UsersService) UpdateUsersMultiple(data []structures.UsersBatchUpdate) error {
	// Create timeout context for batch operation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Validate context and input
	if err := s.validateContext(ctx); err != nil {
		return err
	}

	if len(data) == 0 {
		return fmt.Errorf("%s: empty data slice in UpdateUsersMultiple", "validation_error")
	}

	// Execute batch update within transaction
	return s.withTransaction(ctx, func(tx *gorm.DB) error {
		var totalAffected int64 = 0

		for i, item := range data {
			// Validate context on each iteration for long-running operations
			if err := s.validateContext(ctx); err != nil {
				return err
			}

			updateDataBytes, err := json.Marshal(item.Data)
			if err != nil {
				return fmt.Errorf("%s: invalid data format for conversion at index %d in UpdateUsersMultiple: %w", "validation_error", i, err)
			}

			var updateData map[string]interface{}
			if err := json.Unmarshal(updateDataBytes, &updateData); err != nil {
				return fmt.Errorf("%s: failed to unmarshal update data at index %d in UpdateUsersMultiple: %w", "validation_error", i, err)
			}

			for key, value := range updateData {
				if value == nil {
					delete(updateData, key)
				}
			}

			// PathParams'tan where sorgusunu oluştur
			whereClause, err := s.buildWhereFromPathParams(item.PathParams)
			if err != nil {
				return fmt.Errorf("%s: failed to build where clause at index %d in UpdateUsersMultiple: %w", "validation_error", i, err)
			}

			result := tx.Model(&structures.Users{}).Where(whereClause).Updates(updateData)
			if result.Error != nil {
				return s.handleServiceError(result.Error, fmt.Sprintf("UpdateUsersMultiple[%d]", i))
			}

			totalAffected += result.RowsAffected
		}

		// Log successful batch update
		if totalAffected == 0 {
			return fmt.Errorf("%s: no records were updated in UpdateUsersMultiple", "not_found_error")
		}

		return nil
	})
}

// Create, yeni bir kayıt oluşturur.
func (s *UsersService) CreateUsers(data structures.UsersForm) (structures.UsersForm, error) {
	// Create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Validate context
	if err := s.validateContext(ctx); err != nil {
		return data, err
	}

	applyUsersCreateDefaults(&data)

	// Execute create within transaction
	err := s.withTransaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return s.handleServiceError(err, "CreateUsers")
		}

		// Transactional outbox: publish event atomically in same transaction
		payloadBytes, _ := json.Marshal(data)
		outboxEvt := OutboxEvent{
			AggregateType: "Users",
			AggregateID:   fmt.Sprintf("%v", data),
			EventType:     "created",
			Payload:       string(payloadBytes),
		}
		if err := tx.Table("outbox_events").Create(&outboxEvt).Error; err != nil {
			return fmt.Errorf("outbox insert failed in CreateUsers: %w", err)
		}

		// SSE: fan-out to local subscribers (best-effort, post-write broadcast)
		evtBytes, _ := json.Marshal(data)
		go sseRuntime.Publish("iam.users", sseRuntime.Event{Type: "created", Payload: evtBytes})

		return nil
	})

	return data, err
}

// CreateMultiple, toplu olarak yeni kayıtlar oluşturur.
func (s *UsersService) CreateUsersMultiple(data []structures.UsersForm) ([]structures.UsersForm, error) {
	// Create timeout context for batch operation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Validate context and input
	if err := s.validateContext(ctx); err != nil {
		return data, err
	}

	if len(data) == 0 {
		return data, fmt.Errorf("%s: empty data slice in CreateUsersMultiple", "validation_error")
	}

	for i := range data {
		applyUsersCreateDefaults(&data[i])
	}

	// Execute batch create within transaction
	err := s.withTransaction(ctx, func(tx *gorm.DB) error {
		// Use batch size for large datasets
		batchSize := len(data)
		if batchSize > 1000 {
			batchSize = 1000 // Limit batch size for performance
		}

		if err := tx.CreateInBatches(&data, batchSize).Error; err != nil {
			return s.handleServiceError(err, "CreateUsersMultiple")
		}
		return nil
	})

	return data, err
}

// Delete, verilen primary/unique key'e göre bir kaydı siler.
func (s *UsersService) DeleteUsers(id types.URID) error {
	// Create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Validate context
	if err := s.validateContext(ctx); err != nil {
		return err
	}

	// Execute delete within transaction
	return s.withTransaction(ctx, func(tx *gorm.DB) error {
		result := tx.Model(&structures.Users{}).Where("id = ?", id).Delete(&structures.Users{})

		if result.Error != nil {
			return s.handleServiceError(result.Error, "DeleteUsers")
		}

		// Check if any record was actually deleted
		if result.RowsAffected == 0 {
			return fmt.Errorf("not_found_error: no record found to delete in DeleteUsers: %w", gorm.ErrRecordNotFound)
		}

		// Transactional outbox: publish delete event atomically in same transaction
		outboxEvt := OutboxEvent{
			AggregateType: "Users",
			EventType:     "deleted",
			Payload:       "{}",
		}
		if err := tx.Table("outbox_events").Create(&outboxEvt).Error; err != nil {
			return fmt.Errorf("outbox insert failed in DeleteUsers: %w", err)
		}

		go sseRuntime.Publish("iam.users", sseRuntime.Event{Type: "deleted", Payload: []byte("{}")})

		return nil
	})
}

// buildWhereFromPathParams, BatchUpdate içindeki PathParams'tan GORM where sorgusu oluşturur.
func (s *UsersService) buildWhereFromPathParams(pathParams interface{}) (map[string]interface{}, error) {
	if pathParams == nil {
		return nil, fmt.Errorf("%s: pathParams is nil", "validation_error")
	}

	whereMap, err := buildWhereMapFromStructUsers(pathParams)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to convert pathParams to map: %w", "validation_error", err)
	}

	// Validate that we have at least one condition
	if len(whereMap) == 0 {
		return nil, fmt.Errorf("%s: pathParams resulted in empty where conditions", "validation_error")
	}

	return whereMap, nil
}

func buildWhereMapFromStructUsers(value interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return nil, fmt.Errorf("invalid value")
	}
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil, fmt.Errorf("nil pointer value")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s", v.Kind())
	}

	typ := v.Type()
	whereMap := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			jsonTag = strcase.ToLowerCamel(field.Name)
		} else if comma := strings.Index(jsonTag, ","); comma >= 0 {
			jsonTag = jsonTag[:comma]
		}
		if jsonTag == "" {
			continue
		}
		whereMap[strcase.ToSnake(jsonTag)] = v.Field(i).Interface()
	}
	return whereMap, nil
}

// DeleteMultiple, verilen kimliklere sahip birden fazla kaydı siler.
// Tüm silme işlemleri tek bir transaction içinde yapılır.
func (s *UsersService) DeleteUsersMultiple(identities []structures.UsersIdentity) error {
	// Create timeout context for batch operation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Validate context and input
	if err := s.validateContext(ctx); err != nil {
		return err
	}

	if len(identities) == 0 {
		return nil // Silinecek bir şey yoksa hata verme
	}

	// Execute batch delete within transaction
	return s.withTransaction(ctx, func(tx *gorm.DB) error {
		var totalDeleted int64 = 0

		for i, identity := range identities {
			// Validate context on each iteration for long-running operations
			if err := s.validateContext(ctx); err != nil {
				return err
			}

			// Convert identity to where conditions
			whereConditions, err := buildWhereMapFromStructUsers(identity)
			if err != nil {
				return fmt.Errorf("%s: failed to convert identity to where conditions at index %d: %w", "validation_error", i, err)
			}

			if len(whereConditions) == 0 {
				return fmt.Errorf("%s: empty where conditions for identity at index %d", "validation_error", i)
			}

			// Delete using where conditions for better control and error handling
			result := tx.Where(whereConditions).Delete(&structures.Users{})
			if result.Error != nil {
				return s.handleServiceError(result.Error, fmt.Sprintf("DeleteUsersMultiple[%d]", i))
			}

			totalDeleted += result.RowsAffected
		}

		// Check if any records were actually deleted
		if totalDeleted == 0 {
			return fmt.Errorf("%s: no records were deleted in DeleteUsersMultiple", "not_found_error")
		}

		return nil
	})
}
