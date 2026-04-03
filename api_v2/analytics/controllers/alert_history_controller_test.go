package analytics_api_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	testhelper "data-bridge-examples/api_v2/analytics/testhelper"

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// Projenizin doğru import path'lerini kullandığınızdan emin olun
	structures "data-bridge-examples/api_v2/analytics/structures"

	"github.com/maple-tech/baseline/types" // types.ID, types.URID için
	"github.com/maple-tech/baseline/web"   // web.Fiber200Response, web.Fiber400Response için
)

var (
	_ = types.ID(0)
	_ = types.URID("")
)

// --- Mock Service Definition ---
// MockAlertHistoryService, IAlertHistoryService arayüzünü taklit eder.
type MockAlertHistoryService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockAlertHistoryService) CreateAlertHistory(data structures.AlertHistoryForm) (structures.AlertHistoryForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.AlertHistoryForm{}, args.Error(1)
	}
	return args.Get(0).(structures.AlertHistoryForm), args.Error(1)
}

func (m *MockAlertHistoryService) CreateAlertHistoryMultiple(data []structures.AlertHistoryForm) ([]structures.AlertHistoryForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.AlertHistoryForm), args.Error(1)
}

func (m *MockAlertHistoryService) UpdateAlertHistory(id types.URID, data structures.AlertHistoryEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockAlertHistoryService) UpdateAlertHistoryMultiple(data []structures.AlertHistoryBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockAlertHistoryService) DeleteAlertHistory(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockAlertHistoryService) DeleteAlertHistoryMultiple(identities []structures.AlertHistoryIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateAlertHistory_Success
func TestCreateAlertHistory_Success(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_alertHistory", controller.CreateAlertHistory)

	inputData := structures.AlertHistoryForm{}
	expectedResult := inputData
	mockService.On("CreateAlertHistory", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_alertHistory", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.AlertHistoryForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateAlertHistory_ServiceError
func TestCreateAlertHistory_ServiceError(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_alertHistory", controller.CreateAlertHistory)

	inputData := structures.AlertHistoryForm{}
	mockService.On("CreateAlertHistory", inputData).Return(structures.AlertHistoryForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_alertHistory", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteAlertHistoryMultiple_Success(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_alertHistory/bulk", controller.DeleteAlertHistoryMultiple)

	identities := []structures.AlertHistoryIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteAlertHistoryMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_alertHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteAlertHistoryMultiple_ServiceError(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_alertHistory/bulk", controller.DeleteAlertHistoryMultiple)

	identities := []structures.AlertHistoryIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteAlertHistoryMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_alertHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteAlertHistoryMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_alertHistory/bulk", controller.DeleteAlertHistoryMultiple)

	var emptyIdentities []structures.AlertHistoryIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_alertHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteAlertHistoryMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateAlertHistoryMultiple_Success(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_alertHistory/bulk", controller.UpdateAlertHistoryMultiple)

	batchUpdatePayload := []structures.AlertHistoryBatchUpdate{
		{
			PathParams: structures.AlertHistoryIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.AlertHistoryEdit{Name: alertHistoryStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.AlertHistoryIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.AlertHistoryEdit{Name: alertHistoryStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateAlertHistoryMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_alertHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateAlertHistoryMultiple_ServiceError(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_alertHistory/bulk", controller.UpdateAlertHistoryMultiple)

	batchUpdatePayload := []structures.AlertHistoryBatchUpdate{
		{PathParams: structures.AlertHistoryIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.AlertHistoryEdit{Name: alertHistoryStringPtr("Fail Update")}},
	}
	mockService.On("UpdateAlertHistoryMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_alertHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateAlertHistoryMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_alertHistory/bulk", controller.UpdateAlertHistoryMultiple)

	var emptyPayload []structures.AlertHistoryBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_alertHistory/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateAlertHistoryMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde AlertHistoryEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func alertHistoryStringPtr(s string) *string { return testhelper.StringPtr(s) }
func alertHistoryIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteAlertHistoryMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockAlertHistoryService)
	controller := AlertHistoryController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_alertHistory/bulk", controller.DeleteAlertHistoryMultiple)

	req := httptest.NewRequest("DELETE", "/test_alertHistory/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteAlertHistoryMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateAlertHistory(pk, editData) mock'lanır; Delete için c.Svc.DeleteAlertHistory(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetAlertHistoryWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetAlertHistoryWithPagination_Success(t *testing.T) {
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
