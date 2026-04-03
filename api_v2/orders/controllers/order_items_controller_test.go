package orders_api_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	testhelper "data-bridge-examples/api_v2/orders/testhelper"

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// Projenizin doğru import path'lerini kullandığınızdan emin olun
	structures "data-bridge-examples/api_v2/orders/structures"

	"github.com/maple-tech/baseline/types" // types.ID, types.URID için
	"github.com/maple-tech/baseline/web"   // web.Fiber200Response, web.Fiber400Response için
)

var (
	_ = types.ID(0)
	_ = types.URID("")
)

// --- Mock Service Definition ---
// MockOrderItemsService, IOrderItemsService arayüzünü taklit eder.
type MockOrderItemsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockOrderItemsService) CreateOrderItems(data structures.OrderItemsForm) (structures.OrderItemsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.OrderItemsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.OrderItemsForm), args.Error(1)
}

func (m *MockOrderItemsService) CreateOrderItemsMultiple(data []structures.OrderItemsForm) ([]structures.OrderItemsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.OrderItemsForm), args.Error(1)
}

func (m *MockOrderItemsService) UpdateOrderItems(id types.URID, data structures.OrderItemsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockOrderItemsService) UpdateOrderItemsMultiple(data []structures.OrderItemsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockOrderItemsService) DeleteOrderItems(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockOrderItemsService) DeleteOrderItemsMultiple(identities []structures.OrderItemsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateOrderItems_Success
func TestCreateOrderItems_Success(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_orderItems", controller.CreateOrderItems)

	inputData := structures.OrderItemsForm{}
	expectedResult := inputData
	mockService.On("CreateOrderItems", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_orderItems", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.OrderItemsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateOrderItems_ServiceError
func TestCreateOrderItems_ServiceError(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_orderItems", controller.CreateOrderItems)

	inputData := structures.OrderItemsForm{}
	mockService.On("CreateOrderItems", inputData).Return(structures.OrderItemsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_orderItems", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteOrderItemsMultiple_Success(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_orderItems/bulk", controller.DeleteOrderItemsMultiple)

	identities := []structures.OrderItemsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteOrderItemsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_orderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteOrderItemsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_orderItems/bulk", controller.DeleteOrderItemsMultiple)

	identities := []structures.OrderItemsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteOrderItemsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_orderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteOrderItemsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_orderItems/bulk", controller.DeleteOrderItemsMultiple)

	var emptyIdentities []structures.OrderItemsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_orderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteOrderItemsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateOrderItemsMultiple_Success(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_orderItems/bulk", controller.UpdateOrderItemsMultiple)

	batchUpdatePayload := []structures.OrderItemsBatchUpdate{
		{
			PathParams: structures.OrderItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.OrderItemsEdit{Name: orderItemsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.OrderItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.OrderItemsEdit{Name: orderItemsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateOrderItemsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_orderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateOrderItemsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_orderItems/bulk", controller.UpdateOrderItemsMultiple)

	batchUpdatePayload := []structures.OrderItemsBatchUpdate{
		{PathParams: structures.OrderItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.OrderItemsEdit{Name: orderItemsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateOrderItemsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_orderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateOrderItemsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_orderItems/bulk", controller.UpdateOrderItemsMultiple)

	var emptyPayload []structures.OrderItemsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_orderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateOrderItemsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde OrderItemsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func orderItemsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func orderItemsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteOrderItemsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockOrderItemsService)
	controller := OrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_orderItems/bulk", controller.DeleteOrderItemsMultiple)

	req := httptest.NewRequest("DELETE", "/test_orderItems/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteOrderItemsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateOrderItems(pk, editData) mock'lanır; Delete için c.Svc.DeleteOrderItems(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetOrderItemsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetOrderItemsWithPagination_Success(t *testing.T) {
	// Bu test, pagination'ın kendisini değil, controller'ın pagination'ı doğru çağırıp
	// sonucunu doğru döndürdüğünü test eder. Pagination'ın iç mantığı kendi unit testlerine sahip olmalı.
	// Bu test için servis mock'lamaya gerek yok çünkü pagination DB'den okuyor.
	// Ancak, DB'yi mocklamak çok karmaşık olacağından, bu test entegrasyon testine daha yakındır
	// veya controller'ın doğrudan pagination objesini çağırdığı varsayılır.
	// Şimdilik, controller'ın pagination sonucunu doğru formatta döndürdüğünü test edelim.

	// ARRANGE
	// Gerçek bir servis ve DB mock'u olmadan bu testi tam unit test yapmak zor.
	// Bu nedenle, bu genellikle API/entegrasyon testi olarak daha anlamlıdır.
	// VEYA, pagination'ı da mock'layabiliriz (ileride düşünülebilir).
	// Şimdilik bu testi kavramsal olarak bırakıyorum, çünkü gerçek DB veya karmaşık mock gerektirir.
	t.Skip("Skipping GetWithPagination unit test as it requires DB interaction or complex pagination mocking.")
}
