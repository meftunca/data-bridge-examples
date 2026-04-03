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
// MockStorageBinsService, IStorageBinsService arayüzünü taklit eder.
type MockStorageBinsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockStorageBinsService) CreateStorageBins(data structures.StorageBinsForm) (structures.StorageBinsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.StorageBinsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.StorageBinsForm), args.Error(1)
}

func (m *MockStorageBinsService) CreateStorageBinsMultiple(data []structures.StorageBinsForm) ([]structures.StorageBinsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.StorageBinsForm), args.Error(1)
}

func (m *MockStorageBinsService) UpdateStorageBins(id types.URID, data structures.StorageBinsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockStorageBinsService) UpdateStorageBinsMultiple(data []structures.StorageBinsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockStorageBinsService) DeleteStorageBins(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockStorageBinsService) DeleteStorageBinsMultiple(identities []structures.StorageBinsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateStorageBins_Success
func TestCreateStorageBins_Success(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_storageBins", controller.CreateStorageBins)

	inputData := structures.StorageBinsForm{}
	expectedResult := inputData
	mockService.On("CreateStorageBins", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_storageBins", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.StorageBinsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateStorageBins_ServiceError
func TestCreateStorageBins_ServiceError(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_storageBins", controller.CreateStorageBins)

	inputData := structures.StorageBinsForm{}
	mockService.On("CreateStorageBins", inputData).Return(structures.StorageBinsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_storageBins", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteStorageBinsMultiple_Success(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_storageBins/bulk", controller.DeleteStorageBinsMultiple)

	identities := []structures.StorageBinsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteStorageBinsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_storageBins/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteStorageBinsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_storageBins/bulk", controller.DeleteStorageBinsMultiple)

	identities := []structures.StorageBinsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteStorageBinsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_storageBins/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteStorageBinsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_storageBins/bulk", controller.DeleteStorageBinsMultiple)

	var emptyIdentities []structures.StorageBinsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_storageBins/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteStorageBinsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateStorageBinsMultiple_Success(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_storageBins/bulk", controller.UpdateStorageBinsMultiple)

	batchUpdatePayload := []structures.StorageBinsBatchUpdate{
		{
			PathParams: structures.StorageBinsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.StorageBinsEdit{Name: storageBinsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.StorageBinsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.StorageBinsEdit{Name: storageBinsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateStorageBinsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_storageBins/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateStorageBinsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_storageBins/bulk", controller.UpdateStorageBinsMultiple)

	batchUpdatePayload := []structures.StorageBinsBatchUpdate{
		{PathParams: structures.StorageBinsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.StorageBinsEdit{Name: storageBinsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateStorageBinsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_storageBins/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateStorageBinsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_storageBins/bulk", controller.UpdateStorageBinsMultiple)

	var emptyPayload []structures.StorageBinsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_storageBins/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateStorageBinsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde StorageBinsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func storageBinsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func storageBinsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteStorageBinsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockStorageBinsService)
	controller := StorageBinsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_storageBins/bulk", controller.DeleteStorageBinsMultiple)

	req := httptest.NewRequest("DELETE", "/test_storageBins/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteStorageBinsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateStorageBins(pk, editData) mock'lanır; Delete için c.Svc.DeleteStorageBins(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetStorageBinsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetStorageBinsWithPagination_Success(t *testing.T) {
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
