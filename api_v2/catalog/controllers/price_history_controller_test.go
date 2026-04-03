package catalog_api_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	testhelper "data-bridge-examples/api_v2/catalog/testhelper"

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// Projenizin doğru import path'lerini kullandığınızdan emin olun
	structures "data-bridge-examples/api_v2/catalog/structures"

	"github.com/maple-tech/baseline/types" // types.ID, types.URID için
	"github.com/maple-tech/baseline/web"   // web.Fiber200Response, web.Fiber400Response için
)

var (
	_ = types.ID(0)
	_ = types.URID("")
)

// --- Mock Service Definition ---
// MockPriceHistoryService, IPriceHistoryService arayüzünü taklit eder.
type MockPriceHistoryService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockPriceHistoryService) CreatePriceHistory(data structures.PriceHistoryForm) (structures.PriceHistoryForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.PriceHistoryForm{}, args.Error(1)
	}
	return args.Get(0).(structures.PriceHistoryForm), args.Error(1)
}

func (m *MockPriceHistoryService) CreatePriceHistoryMultiple(data []structures.PriceHistoryForm) ([]structures.PriceHistoryForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.PriceHistoryForm), args.Error(1)
}

func (m *MockPriceHistoryService) UpdatePriceHistory(id types.URID, data structures.PriceHistoryEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockPriceHistoryService) UpdatePriceHistoryMultiple(data []structures.PriceHistoryBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockPriceHistoryService) DeletePriceHistory(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockPriceHistoryService) DeletePriceHistoryMultiple(identities []structures.PriceHistoryIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreatePriceHistory_Success
func TestCreatePriceHistory_Success(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_priceHistory", controller.CreatePriceHistory)

	inputData := structures.PriceHistoryForm{}
	expectedResult := inputData
	mockService.On("CreatePriceHistory", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_priceHistory", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.PriceHistoryForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreatePriceHistory_ServiceError
func TestCreatePriceHistory_ServiceError(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_priceHistory", controller.CreatePriceHistory)

	inputData := structures.PriceHistoryForm{}
	mockService.On("CreatePriceHistory", inputData).Return(structures.PriceHistoryForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_priceHistory", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeletePriceHistoryMultiple_Success(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_priceHistory/bulk", controller.DeletePriceHistoryMultiple)

	identities := []structures.PriceHistoryIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeletePriceHistoryMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_priceHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeletePriceHistoryMultiple_ServiceError(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_priceHistory/bulk", controller.DeletePriceHistoryMultiple)

	identities := []structures.PriceHistoryIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeletePriceHistoryMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_priceHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeletePriceHistoryMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_priceHistory/bulk", controller.DeletePriceHistoryMultiple)

	var emptyIdentities []structures.PriceHistoryIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_priceHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeletePriceHistoryMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdatePriceHistoryMultiple_Success(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_priceHistory/bulk", controller.UpdatePriceHistoryMultiple)

	batchUpdatePayload := []structures.PriceHistoryBatchUpdate{
		{
			PathParams: structures.PriceHistoryIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.PriceHistoryEdit{Name: priceHistoryStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.PriceHistoryIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.PriceHistoryEdit{Name: priceHistoryStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdatePriceHistoryMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_priceHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdatePriceHistoryMultiple_ServiceError(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_priceHistory/bulk", controller.UpdatePriceHistoryMultiple)

	batchUpdatePayload := []structures.PriceHistoryBatchUpdate{
		{PathParams: structures.PriceHistoryIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.PriceHistoryEdit{Name: priceHistoryStringPtr("Fail Update")}},
	}
	mockService.On("UpdatePriceHistoryMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_priceHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdatePriceHistoryMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_priceHistory/bulk", controller.UpdatePriceHistoryMultiple)

	var emptyPayload []structures.PriceHistoryBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_priceHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdatePriceHistoryMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde PriceHistoryEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func priceHistoryStringPtr(s string) *string { return testhelper.StringPtr(s) }
func priceHistoryIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeletePriceHistoryMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockPriceHistoryService)
	controller := PriceHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_priceHistory/bulk", controller.DeletePriceHistoryMultiple)

	req := httptest.NewRequest("DELETE", "/test_priceHistory/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeletePriceHistoryMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdatePriceHistory(pk, editData) mock'lanır; Delete için c.Svc.DeletePriceHistory(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetPriceHistoryWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetPriceHistoryWithPagination_Success(t *testing.T) {
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
