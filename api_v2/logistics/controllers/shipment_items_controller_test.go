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
// MockShipmentItemsService, IShipmentItemsService arayüzünü taklit eder.
type MockShipmentItemsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockShipmentItemsService) CreateShipmentItems(data structures.ShipmentItemsForm) (structures.ShipmentItemsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ShipmentItemsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ShipmentItemsForm), args.Error(1)
}

func (m *MockShipmentItemsService) CreateShipmentItemsMultiple(data []structures.ShipmentItemsForm) ([]structures.ShipmentItemsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ShipmentItemsForm), args.Error(1)
}

func (m *MockShipmentItemsService) UpdateShipmentItems(id types.URID, data structures.ShipmentItemsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockShipmentItemsService) UpdateShipmentItemsMultiple(data []structures.ShipmentItemsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockShipmentItemsService) DeleteShipmentItems(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockShipmentItemsService) DeleteShipmentItemsMultiple(identities []structures.ShipmentItemsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateShipmentItems_Success
func TestCreateShipmentItems_Success(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_shipmentItems", controller.CreateShipmentItems)

	inputData := structures.ShipmentItemsForm{}
	expectedResult := inputData
	mockService.On("CreateShipmentItems", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_shipmentItems", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ShipmentItemsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateShipmentItems_ServiceError
func TestCreateShipmentItems_ServiceError(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_shipmentItems", controller.CreateShipmentItems)

	inputData := structures.ShipmentItemsForm{}
	mockService.On("CreateShipmentItems", inputData).Return(structures.ShipmentItemsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_shipmentItems", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteShipmentItemsMultiple_Success(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipmentItems/bulk", controller.DeleteShipmentItemsMultiple)

	identities := []structures.ShipmentItemsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteShipmentItemsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_shipmentItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteShipmentItemsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipmentItems/bulk", controller.DeleteShipmentItemsMultiple)

	identities := []structures.ShipmentItemsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteShipmentItemsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_shipmentItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteShipmentItemsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipmentItems/bulk", controller.DeleteShipmentItemsMultiple)

	var emptyIdentities []structures.ShipmentItemsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_shipmentItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteShipmentItemsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateShipmentItemsMultiple_Success(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipmentItems/bulk", controller.UpdateShipmentItemsMultiple)

	batchUpdatePayload := []structures.ShipmentItemsBatchUpdate{
		{
			PathParams: structures.ShipmentItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ShipmentItemsEdit{Name: shipmentItemsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ShipmentItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ShipmentItemsEdit{Name: shipmentItemsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateShipmentItemsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_shipmentItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateShipmentItemsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipmentItems/bulk", controller.UpdateShipmentItemsMultiple)

	batchUpdatePayload := []structures.ShipmentItemsBatchUpdate{
		{PathParams: structures.ShipmentItemsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ShipmentItemsEdit{Name: shipmentItemsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateShipmentItemsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_shipmentItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateShipmentItemsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipmentItems/bulk", controller.UpdateShipmentItemsMultiple)

	var emptyPayload []structures.ShipmentItemsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_shipmentItems/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateShipmentItemsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ShipmentItemsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func shipmentItemsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func shipmentItemsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteShipmentItemsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockShipmentItemsService)
	controller := ShipmentItemsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipmentItems/bulk", controller.DeleteShipmentItemsMultiple)

	req := httptest.NewRequest("DELETE", "/test_shipmentItems/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteShipmentItemsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateShipmentItems(pk, editData) mock'lanır; Delete için c.Svc.DeleteShipmentItems(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetShipmentItemsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetShipmentItemsWithPagination_Success(t *testing.T) {
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
