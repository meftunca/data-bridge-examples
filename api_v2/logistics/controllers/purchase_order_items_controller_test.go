package logistics_api_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	testhelper "data-bridge-examples/api_v2/logistics/testhelper"

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// Projenizin doğru import path'lerini kullandığınızdan emin olun
	structures "data-bridge-examples/api_v2/logistics/structures"

	"github.com/maple-tech/baseline/types" // types.ID, types.URID için
	"github.com/maple-tech/baseline/web"   // web.Fiber200Response, web.Fiber400Response için
)

var (
	_ = types.ID(0)
	_ = types.URID("")
)

// --- Mock Service Definition ---
// MockPurchaseOrderItemsService, IPurchaseOrderItemsService arayüzünü taklit eder.
type MockPurchaseOrderItemsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockPurchaseOrderItemsService) CreatePurchaseOrderItems(data structures.PurchaseOrderItemsForm) (structures.PurchaseOrderItemsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.PurchaseOrderItemsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.PurchaseOrderItemsForm), args.Error(1)
}

func (m *MockPurchaseOrderItemsService) CreatePurchaseOrderItemsMultiple(data []structures.PurchaseOrderItemsForm) ([]structures.PurchaseOrderItemsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.PurchaseOrderItemsForm), args.Error(1)
}

func (m *MockPurchaseOrderItemsService) UpdatePurchaseOrderItems(id types.URID, data structures.PurchaseOrderItemsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockPurchaseOrderItemsService) UpdatePurchaseOrderItemsMultiple(data []structures.PurchaseOrderItemsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockPurchaseOrderItemsService) DeletePurchaseOrderItems(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockPurchaseOrderItemsService) DeletePurchaseOrderItemsMultiple(identities []structures.PurchaseOrderItemsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreatePurchaseOrderItems_Success
func TestCreatePurchaseOrderItems_Success(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_purchaseOrderItems", controller.CreatePurchaseOrderItems)

	inputData := structures.PurchaseOrderItemsForm{}
	expectedResult := inputData
	mockService.On("CreatePurchaseOrderItems", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_purchaseOrderItems", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.PurchaseOrderItemsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreatePurchaseOrderItems_ServiceError
func TestCreatePurchaseOrderItems_ServiceError(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_purchaseOrderItems", controller.CreatePurchaseOrderItems)

	inputData := structures.PurchaseOrderItemsForm{}
	mockService.On("CreatePurchaseOrderItems", inputData).Return(structures.PurchaseOrderItemsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_purchaseOrderItems", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeletePurchaseOrderItemsMultiple_Success(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_purchaseOrderItems/bulk", controller.DeletePurchaseOrderItemsMultiple)

	identities := []structures.PurchaseOrderItemsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeletePurchaseOrderItemsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_purchaseOrderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeletePurchaseOrderItemsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_purchaseOrderItems/bulk", controller.DeletePurchaseOrderItemsMultiple)

	identities := []structures.PurchaseOrderItemsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeletePurchaseOrderItemsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_purchaseOrderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeletePurchaseOrderItemsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_purchaseOrderItems/bulk", controller.DeletePurchaseOrderItemsMultiple)

	var emptyIdentities []structures.PurchaseOrderItemsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_purchaseOrderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeletePurchaseOrderItemsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdatePurchaseOrderItemsMultiple_Success(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_purchaseOrderItems/bulk", controller.UpdatePurchaseOrderItemsMultiple)

	batchUpdatePayload := []structures.PurchaseOrderItemsBatchUpdate{
		{
			PathParams: structures.PurchaseOrderItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.PurchaseOrderItemsEdit{Name: purchaseOrderItemsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.PurchaseOrderItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.PurchaseOrderItemsEdit{Name: purchaseOrderItemsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdatePurchaseOrderItemsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_purchaseOrderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdatePurchaseOrderItemsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_purchaseOrderItems/bulk", controller.UpdatePurchaseOrderItemsMultiple)

	batchUpdatePayload := []structures.PurchaseOrderItemsBatchUpdate{
		{PathParams: structures.PurchaseOrderItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.PurchaseOrderItemsEdit{Name: purchaseOrderItemsStringPtr("Fail Update")}},
	}
	mockService.On("UpdatePurchaseOrderItemsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_purchaseOrderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdatePurchaseOrderItemsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_purchaseOrderItems/bulk", controller.UpdatePurchaseOrderItemsMultiple)

	var emptyPayload []structures.PurchaseOrderItemsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_purchaseOrderItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdatePurchaseOrderItemsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde PurchaseOrderItemsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func purchaseOrderItemsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func purchaseOrderItemsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeletePurchaseOrderItemsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockPurchaseOrderItemsService)
	controller := PurchaseOrderItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_purchaseOrderItems/bulk", controller.DeletePurchaseOrderItemsMultiple)

	req := httptest.NewRequest("DELETE", "/test_purchaseOrderItems/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeletePurchaseOrderItemsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdatePurchaseOrderItems(pk, editData) mock'lanır; Delete için c.Svc.DeletePurchaseOrderItems(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetPurchaseOrderItemsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetPurchaseOrderItemsWithPagination_Success(t *testing.T) {
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
