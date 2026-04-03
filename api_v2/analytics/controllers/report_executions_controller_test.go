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
// MockReportExecutionsService, IReportExecutionsService arayüzünü taklit eder.
type MockReportExecutionsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockReportExecutionsService) CreateReportExecutions(data structures.ReportExecutionsForm) (structures.ReportExecutionsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ReportExecutionsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ReportExecutionsForm), args.Error(1)
}

func (m *MockReportExecutionsService) CreateReportExecutionsMultiple(data []structures.ReportExecutionsForm) ([]structures.ReportExecutionsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ReportExecutionsForm), args.Error(1)
}

func (m *MockReportExecutionsService) UpdateReportExecutions(id types.URID, data structures.ReportExecutionsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockReportExecutionsService) UpdateReportExecutionsMultiple(data []structures.ReportExecutionsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockReportExecutionsService) DeleteReportExecutions(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockReportExecutionsService) DeleteReportExecutionsMultiple(identities []structures.ReportExecutionsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateReportExecutions_Success
func TestCreateReportExecutions_Success(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_reportExecutions", controller.CreateReportExecutions)

	inputData := structures.ReportExecutionsForm{}
	expectedResult := inputData
	mockService.On("CreateReportExecutions", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_reportExecutions", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ReportExecutionsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateReportExecutions_ServiceError
func TestCreateReportExecutions_ServiceError(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_reportExecutions", controller.CreateReportExecutions)

	inputData := structures.ReportExecutionsForm{}
	mockService.On("CreateReportExecutions", inputData).Return(structures.ReportExecutionsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_reportExecutions", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteReportExecutionsMultiple_Success(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_reportExecutions/bulk", controller.DeleteReportExecutionsMultiple)

	identities := []structures.ReportExecutionsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteReportExecutionsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_reportExecutions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteReportExecutionsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_reportExecutions/bulk", controller.DeleteReportExecutionsMultiple)

	identities := []structures.ReportExecutionsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteReportExecutionsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_reportExecutions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteReportExecutionsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_reportExecutions/bulk", controller.DeleteReportExecutionsMultiple)

	var emptyIdentities []structures.ReportExecutionsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_reportExecutions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteReportExecutionsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateReportExecutionsMultiple_Success(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_reportExecutions/bulk", controller.UpdateReportExecutionsMultiple)

	batchUpdatePayload := []structures.ReportExecutionsBatchUpdate{
		{
			PathParams: structures.ReportExecutionsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ReportExecutionsEdit{Name: reportExecutionsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ReportExecutionsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ReportExecutionsEdit{Name: reportExecutionsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateReportExecutionsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_reportExecutions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateReportExecutionsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_reportExecutions/bulk", controller.UpdateReportExecutionsMultiple)

	batchUpdatePayload := []structures.ReportExecutionsBatchUpdate{
		{PathParams: structures.ReportExecutionsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ReportExecutionsEdit{Name: reportExecutionsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateReportExecutionsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_reportExecutions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateReportExecutionsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_reportExecutions/bulk", controller.UpdateReportExecutionsMultiple)

	var emptyPayload []structures.ReportExecutionsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_reportExecutions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateReportExecutionsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ReportExecutionsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func reportExecutionsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func reportExecutionsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteReportExecutionsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockReportExecutionsService)
	controller := ReportExecutionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_reportExecutions/bulk", controller.DeleteReportExecutionsMultiple)

	req := httptest.NewRequest("DELETE", "/test_reportExecutions/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteReportExecutionsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateReportExecutions(pk, editData) mock'lanır; Delete için c.Svc.DeleteReportExecutions(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetReportExecutionsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetReportExecutionsWithPagination_Success(t *testing.T) {
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
