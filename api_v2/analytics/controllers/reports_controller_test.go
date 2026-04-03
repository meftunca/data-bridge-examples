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
// MockReportsService, IReportsService arayüzünü taklit eder.
type MockReportsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockReportsService) CreateReports(data structures.ReportsForm) (structures.ReportsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ReportsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ReportsForm), args.Error(1)
}

func (m *MockReportsService) CreateReportsMultiple(data []structures.ReportsForm) ([]structures.ReportsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ReportsForm), args.Error(1)
}

func (m *MockReportsService) UpdateReports(id types.URID, data structures.ReportsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockReportsService) UpdateReportsMultiple(data []structures.ReportsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockReportsService) DeleteReports(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockReportsService) DeleteReportsMultiple(identities []structures.ReportsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateReports_Success
func TestCreateReports_Success(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_reports", controller.CreateReports)

	inputData := structures.ReportsForm{}
	expectedResult := inputData
	mockService.On("CreateReports", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_reports", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ReportsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateReports_ServiceError
func TestCreateReports_ServiceError(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_reports", controller.CreateReports)

	inputData := structures.ReportsForm{}
	mockService.On("CreateReports", inputData).Return(structures.ReportsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_reports", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteReportsMultiple_Success(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_reports/bulk", controller.DeleteReportsMultiple)

	identities := []structures.ReportsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteReportsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_reports/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteReportsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_reports/bulk", controller.DeleteReportsMultiple)

	identities := []structures.ReportsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteReportsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_reports/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteReportsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_reports/bulk", controller.DeleteReportsMultiple)

	var emptyIdentities []structures.ReportsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_reports/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteReportsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateReportsMultiple_Success(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_reports/bulk", controller.UpdateReportsMultiple)

	batchUpdatePayload := []structures.ReportsBatchUpdate{
		{
			PathParams: structures.ReportsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ReportsEdit{Name: reportsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ReportsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ReportsEdit{Name: reportsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateReportsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_reports/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateReportsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_reports/bulk", controller.UpdateReportsMultiple)

	batchUpdatePayload := []structures.ReportsBatchUpdate{
		{PathParams: structures.ReportsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ReportsEdit{Name: reportsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateReportsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_reports/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateReportsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_reports/bulk", controller.UpdateReportsMultiple)

	var emptyPayload []structures.ReportsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_reports/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateReportsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ReportsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func reportsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func reportsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteReportsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockReportsService)
	controller := ReportsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_reports/bulk", controller.DeleteReportsMultiple)

	req := httptest.NewRequest("DELETE", "/test_reports/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteReportsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateReports(pk, editData) mock'lanır; Delete için c.Svc.DeleteReports(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetReportsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetReportsWithPagination_Success(t *testing.T) {
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
