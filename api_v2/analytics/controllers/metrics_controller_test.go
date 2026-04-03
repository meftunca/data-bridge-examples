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
// MockMetricsService, IMetricsService arayüzünü taklit eder.
type MockMetricsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockMetricsService) CreateMetrics(data structures.MetricsForm) (structures.MetricsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.MetricsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.MetricsForm), args.Error(1)
}

func (m *MockMetricsService) CreateMetricsMultiple(data []structures.MetricsForm) ([]structures.MetricsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.MetricsForm), args.Error(1)
}

func (m *MockMetricsService) UpdateMetrics(id types.URID, data structures.MetricsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockMetricsService) UpdateMetricsMultiple(data []structures.MetricsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockMetricsService) DeleteMetrics(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockMetricsService) DeleteMetricsMultiple(identities []structures.MetricsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateMetrics_Success
func TestCreateMetrics_Success(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_metrics", controller.CreateMetrics)

	inputData := structures.MetricsForm{}
	expectedResult := inputData
	mockService.On("CreateMetrics", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_metrics", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.MetricsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateMetrics_ServiceError
func TestCreateMetrics_ServiceError(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_metrics", controller.CreateMetrics)

	inputData := structures.MetricsForm{}
	mockService.On("CreateMetrics", inputData).Return(structures.MetricsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_metrics", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteMetricsMultiple_Success(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_metrics/bulk", controller.DeleteMetricsMultiple)

	identities := []structures.MetricsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteMetricsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_metrics/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteMetricsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_metrics/bulk", controller.DeleteMetricsMultiple)

	identities := []structures.MetricsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteMetricsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_metrics/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteMetricsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_metrics/bulk", controller.DeleteMetricsMultiple)

	var emptyIdentities []structures.MetricsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_metrics/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteMetricsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateMetricsMultiple_Success(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_metrics/bulk", controller.UpdateMetricsMultiple)

	batchUpdatePayload := []structures.MetricsBatchUpdate{
		{
			PathParams: structures.MetricsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.MetricsEdit{Name: metricsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.MetricsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.MetricsEdit{Name: metricsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateMetricsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_metrics/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateMetricsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_metrics/bulk", controller.UpdateMetricsMultiple)

	batchUpdatePayload := []structures.MetricsBatchUpdate{
		{PathParams: structures.MetricsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.MetricsEdit{Name: metricsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateMetricsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_metrics/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateMetricsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_metrics/bulk", controller.UpdateMetricsMultiple)

	var emptyPayload []structures.MetricsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_metrics/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateMetricsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde MetricsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func metricsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func metricsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteMetricsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockMetricsService)
	controller := MetricsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_metrics/bulk", controller.DeleteMetricsMultiple)

	req := httptest.NewRequest("DELETE", "/test_metrics/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteMetricsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateMetrics(pk, editData) mock'lanır; Delete için c.Svc.DeleteMetrics(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetMetricsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetMetricsWithPagination_Success(t *testing.T) {
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
