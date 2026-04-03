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
// MockStorageZonesService, IStorageZonesService arayüzünü taklit eder.
type MockStorageZonesService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockStorageZonesService) CreateStorageZones(data structures.StorageZonesForm) (structures.StorageZonesForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.StorageZonesForm{}, args.Error(1)
	}
	return args.Get(0).(structures.StorageZonesForm), args.Error(1)
}

func (m *MockStorageZonesService) CreateStorageZonesMultiple(data []structures.StorageZonesForm) ([]structures.StorageZonesForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.StorageZonesForm), args.Error(1)
}

func (m *MockStorageZonesService) UpdateStorageZones(id types.URID, data structures.StorageZonesEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockStorageZonesService) UpdateStorageZonesMultiple(data []structures.StorageZonesBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockStorageZonesService) DeleteStorageZones(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockStorageZonesService) DeleteStorageZonesMultiple(identities []structures.StorageZonesIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateStorageZones_Success
func TestCreateStorageZones_Success(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_storageZones", controller.CreateStorageZones)

	inputData := structures.StorageZonesForm{}
	expectedResult := inputData
	mockService.On("CreateStorageZones", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_storageZones", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.StorageZonesForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateStorageZones_ServiceError
func TestCreateStorageZones_ServiceError(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_storageZones", controller.CreateStorageZones)

	inputData := structures.StorageZonesForm{}
	mockService.On("CreateStorageZones", inputData).Return(structures.StorageZonesForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_storageZones", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteStorageZonesMultiple_Success(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_storageZones/bulk", controller.DeleteStorageZonesMultiple)

	identities := []structures.StorageZonesIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteStorageZonesMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_storageZones/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteStorageZonesMultiple_ServiceError(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_storageZones/bulk", controller.DeleteStorageZonesMultiple)

	identities := []structures.StorageZonesIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteStorageZonesMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_storageZones/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteStorageZonesMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_storageZones/bulk", controller.DeleteStorageZonesMultiple)

	var emptyIdentities []structures.StorageZonesIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_storageZones/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteStorageZonesMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateStorageZonesMultiple_Success(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_storageZones/bulk", controller.UpdateStorageZonesMultiple)

	batchUpdatePayload := []structures.StorageZonesBatchUpdate{
		{
			PathParams: structures.StorageZonesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.StorageZonesEdit{Name: storageZonesStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.StorageZonesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.StorageZonesEdit{Name: storageZonesStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateStorageZonesMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_storageZones/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateStorageZonesMultiple_ServiceError(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_storageZones/bulk", controller.UpdateStorageZonesMultiple)

	batchUpdatePayload := []structures.StorageZonesBatchUpdate{
		{PathParams: structures.StorageZonesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.StorageZonesEdit{Name: storageZonesStringPtr("Fail Update")}},
	}
	mockService.On("UpdateStorageZonesMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_storageZones/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateStorageZonesMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_storageZones/bulk", controller.UpdateStorageZonesMultiple)

	var emptyPayload []structures.StorageZonesBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_storageZones/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateStorageZonesMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde StorageZonesEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func storageZonesStringPtr(s string) *string { return testhelper.StringPtr(s) }
func storageZonesIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteStorageZonesMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockStorageZonesService)
	controller := StorageZonesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_storageZones/bulk", controller.DeleteStorageZonesMultiple)

	req := httptest.NewRequest("DELETE", "/test_storageZones/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteStorageZonesMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateStorageZones(pk, editData) mock'lanır; Delete için c.Svc.DeleteStorageZones(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetStorageZonesWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetStorageZonesWithPagination_Success(t *testing.T) {
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
