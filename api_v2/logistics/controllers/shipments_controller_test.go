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
// MockShipmentsService, IShipmentsService arayüzünü taklit eder.
type MockShipmentsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockShipmentsService) CreateShipments(data structures.ShipmentsForm) (structures.ShipmentsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ShipmentsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ShipmentsForm), args.Error(1)
}

func (m *MockShipmentsService) CreateShipmentsMultiple(data []structures.ShipmentsForm) ([]structures.ShipmentsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ShipmentsForm), args.Error(1)
}

func (m *MockShipmentsService) UpdateShipments(id types.URID, data structures.ShipmentsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockShipmentsService) UpdateShipmentsMultiple(data []structures.ShipmentsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockShipmentsService) DeleteShipments(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockShipmentsService) DeleteShipmentsMultiple(identities []structures.ShipmentsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateShipments_Success
func TestCreateShipments_Success(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_shipments", controller.CreateShipments)

	inputData := structures.ShipmentsForm{}
	expectedResult := inputData
	mockService.On("CreateShipments", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_shipments", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ShipmentsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateShipments_ServiceError
func TestCreateShipments_ServiceError(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_shipments", controller.CreateShipments)

	inputData := structures.ShipmentsForm{}
	mockService.On("CreateShipments", inputData).Return(structures.ShipmentsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_shipments", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteShipmentsMultiple_Success(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipments/bulk", controller.DeleteShipmentsMultiple)

	identities := []structures.ShipmentsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteShipmentsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_shipments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteShipmentsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipments/bulk", controller.DeleteShipmentsMultiple)

	identities := []structures.ShipmentsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteShipmentsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_shipments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteShipmentsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipments/bulk", controller.DeleteShipmentsMultiple)

	var emptyIdentities []structures.ShipmentsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_shipments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteShipmentsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateShipmentsMultiple_Success(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipments/bulk", controller.UpdateShipmentsMultiple)

	batchUpdatePayload := []structures.ShipmentsBatchUpdate{
		{
			PathParams: structures.ShipmentsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ShipmentsEdit{Name: shipmentsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ShipmentsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ShipmentsEdit{Name: shipmentsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateShipmentsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_shipments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateShipmentsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipments/bulk", controller.UpdateShipmentsMultiple)

	batchUpdatePayload := []structures.ShipmentsBatchUpdate{
		{PathParams: structures.ShipmentsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ShipmentsEdit{Name: shipmentsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateShipmentsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_shipments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateShipmentsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_shipments/bulk", controller.UpdateShipmentsMultiple)

	var emptyPayload []structures.ShipmentsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_shipments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateShipmentsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ShipmentsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func shipmentsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func shipmentsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteShipmentsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockShipmentsService)
	controller := ShipmentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_shipments/bulk", controller.DeleteShipmentsMultiple)

	req := httptest.NewRequest("DELETE", "/test_shipments/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteShipmentsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateShipments(pk, editData) mock'lanır; Delete için c.Svc.DeleteShipments(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetShipmentsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetShipmentsWithPagination_Success(t *testing.T) {
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
