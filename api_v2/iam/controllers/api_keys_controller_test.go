package iam_api_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	testhelper "data-bridge-examples/api_v2/iam/testhelper"

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// Projenizin doğru import path'lerini kullandığınızdan emin olun
	structures "data-bridge-examples/api_v2/iam/structures"

	"github.com/maple-tech/baseline/types" // types.ID, types.URID için
	"github.com/maple-tech/baseline/web"   // web.Fiber200Response, web.Fiber400Response için
)

var (
	_ = types.ID(0)
	_ = types.URID("")
)

// --- Mock Service Definition ---
// MockApiKeysService, IApiKeysService arayüzünü taklit eder.
type MockApiKeysService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockApiKeysService) CreateApiKeys(data structures.ApiKeysForm) (structures.ApiKeysForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ApiKeysForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ApiKeysForm), args.Error(1)
}

func (m *MockApiKeysService) CreateApiKeysMultiple(data []structures.ApiKeysForm) ([]structures.ApiKeysForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ApiKeysForm), args.Error(1)
}

func (m *MockApiKeysService) UpdateApiKeys(id types.URID, data structures.ApiKeysEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockApiKeysService) UpdateApiKeysMultiple(data []structures.ApiKeysBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockApiKeysService) DeleteApiKeys(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockApiKeysService) DeleteApiKeysMultiple(identities []structures.ApiKeysIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateApiKeys_Success
func TestCreateApiKeys_Success(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_apiKeys", controller.CreateApiKeys)

	inputData := structures.ApiKeysForm{}
	expectedResult := inputData
	mockService.On("CreateApiKeys", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_apiKeys", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ApiKeysForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateApiKeys_ServiceError
func TestCreateApiKeys_ServiceError(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_apiKeys", controller.CreateApiKeys)

	inputData := structures.ApiKeysForm{}
	mockService.On("CreateApiKeys", inputData).Return(structures.ApiKeysForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_apiKeys", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteApiKeysMultiple_Success(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_apiKeys/bulk", controller.DeleteApiKeysMultiple)

	identities := []structures.ApiKeysIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteApiKeysMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_apiKeys/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteApiKeysMultiple_ServiceError(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_apiKeys/bulk", controller.DeleteApiKeysMultiple)

	identities := []structures.ApiKeysIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteApiKeysMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_apiKeys/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteApiKeysMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_apiKeys/bulk", controller.DeleteApiKeysMultiple)

	var emptyIdentities []structures.ApiKeysIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_apiKeys/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteApiKeysMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateApiKeysMultiple_Success(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_apiKeys/bulk", controller.UpdateApiKeysMultiple)

	batchUpdatePayload := []structures.ApiKeysBatchUpdate{
		{
			PathParams: structures.ApiKeysIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ApiKeysEdit{Name: apiKeysStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ApiKeysIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ApiKeysEdit{Name: apiKeysStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateApiKeysMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_apiKeys/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateApiKeysMultiple_ServiceError(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_apiKeys/bulk", controller.UpdateApiKeysMultiple)

	batchUpdatePayload := []structures.ApiKeysBatchUpdate{
		{PathParams: structures.ApiKeysIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ApiKeysEdit{Name: apiKeysStringPtr("Fail Update")}},
	}
	mockService.On("UpdateApiKeysMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_apiKeys/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateApiKeysMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_apiKeys/bulk", controller.UpdateApiKeysMultiple)

	var emptyPayload []structures.ApiKeysBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_apiKeys/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateApiKeysMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ApiKeysEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func apiKeysStringPtr(s string) *string { return testhelper.StringPtr(s) }
func apiKeysIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteApiKeysMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockApiKeysService)
	controller := ApiKeysController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_apiKeys/bulk", controller.DeleteApiKeysMultiple)

	req := httptest.NewRequest("DELETE", "/test_apiKeys/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteApiKeysMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateApiKeys(pk, editData) mock'lanır; Delete için c.Svc.DeleteApiKeys(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetApiKeysWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetApiKeysWithPagination_Success(t *testing.T) {
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
