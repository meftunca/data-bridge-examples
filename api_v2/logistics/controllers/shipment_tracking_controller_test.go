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
// MockShipmentTrackingService, IShipmentTrackingService arayüzünü taklit eder.
type MockShipmentTrackingService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockShipmentTrackingService) CreateShipmentTracking(data structures.ShipmentTrackingForm) (structures.ShipmentTrackingForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ShipmentTrackingForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ShipmentTrackingForm), args.Error(1)
}

func (m *MockShipmentTrackingService) CreateShipmentTrackingMultiple(data []structures.ShipmentTrackingForm) ([]structures.ShipmentTrackingForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ShipmentTrackingForm), args.Error(1)
}

func (m *MockShipmentTrackingService) UpdateShipmentTracking(id types.URID, data structures.ShipmentTrackingEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockShipmentTrackingService) UpdateShipmentTrackingMultiple(data []structures.ShipmentTrackingBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockShipmentTrackingService) DeleteShipmentTracking(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockShipmentTrackingService) DeleteShipmentTrackingMultiple(identities []structures.ShipmentTrackingIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateShipmentTracking_Success
func TestCreateShipmentTracking_Success(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_shipmentTracking", controller.CreateShipmentTracking)

	inputData := structures.ShipmentTrackingForm{}
	expectedResult := inputData
	mockService.On("CreateShipmentTracking", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_shipmentTracking", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ShipmentTrackingForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateShipmentTracking_ServiceError
func TestCreateShipmentTracking_ServiceError(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_shipmentTracking", controller.CreateShipmentTracking)

	inputData := structures.ShipmentTrackingForm{}
	mockService.On("CreateShipmentTracking", inputData).Return(structures.ShipmentTrackingForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_shipmentTracking", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteShipmentTrackingMultiple_Success(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipmentTracking/bulk", controller.DeleteShipmentTrackingMultiple)

	identities := []structures.ShipmentTrackingIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteShipmentTrackingMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_shipmentTracking/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteShipmentTrackingMultiple_ServiceError(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipmentTracking/bulk", controller.DeleteShipmentTrackingMultiple)

	identities := []structures.ShipmentTrackingIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteShipmentTrackingMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_shipmentTracking/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteShipmentTrackingMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipmentTracking/bulk", controller.DeleteShipmentTrackingMultiple)

	var emptyIdentities []structures.ShipmentTrackingIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_shipmentTracking/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteShipmentTrackingMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateShipmentTrackingMultiple_Success(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipmentTracking/bulk", controller.UpdateShipmentTrackingMultiple)

	batchUpdatePayload := []structures.ShipmentTrackingBatchUpdate{
		{
			PathParams: structures.ShipmentTrackingIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ShipmentTrackingEdit{Name: shipmentTrackingStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ShipmentTrackingIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ShipmentTrackingEdit{Name: shipmentTrackingStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateShipmentTrackingMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_shipmentTracking/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateShipmentTrackingMultiple_ServiceError(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipmentTracking/bulk", controller.UpdateShipmentTrackingMultiple)

	batchUpdatePayload := []structures.ShipmentTrackingBatchUpdate{
		{PathParams: structures.ShipmentTrackingIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ShipmentTrackingEdit{Name: shipmentTrackingStringPtr("Fail Update")}},
	}
	mockService.On("UpdateShipmentTrackingMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_shipmentTracking/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateShipmentTrackingMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipmentTracking/bulk", controller.UpdateShipmentTrackingMultiple)

	var emptyPayload []structures.ShipmentTrackingBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_shipmentTracking/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateShipmentTrackingMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ShipmentTrackingEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func shipmentTrackingStringPtr(s string) *string { return testhelper.StringPtr(s) }
func shipmentTrackingIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteShipmentTrackingMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockShipmentTrackingService)
	controller := ShipmentTrackingController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipmentTracking/bulk", controller.DeleteShipmentTrackingMultiple)

	req := httptest.NewRequest("DELETE", "/test_shipmentTracking/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteShipmentTrackingMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateShipmentTracking(pk, editData) mock'lanır; Delete için c.Svc.DeleteShipmentTracking(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetShipmentTrackingWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetShipmentTrackingWithPagination_Success(t *testing.T) {
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
