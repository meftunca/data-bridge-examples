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
// MockDashboardWidgetsService, IDashboardWidgetsService arayüzünü taklit eder.
type MockDashboardWidgetsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockDashboardWidgetsService) CreateDashboardWidgets(data structures.DashboardWidgetsForm) (structures.DashboardWidgetsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.DashboardWidgetsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.DashboardWidgetsForm), args.Error(1)
}

func (m *MockDashboardWidgetsService) CreateDashboardWidgetsMultiple(data []structures.DashboardWidgetsForm) ([]structures.DashboardWidgetsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.DashboardWidgetsForm), args.Error(1)
}

func (m *MockDashboardWidgetsService) UpdateDashboardWidgets(id types.URID, data structures.DashboardWidgetsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockDashboardWidgetsService) UpdateDashboardWidgetsMultiple(data []structures.DashboardWidgetsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockDashboardWidgetsService) DeleteDashboardWidgets(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockDashboardWidgetsService) DeleteDashboardWidgetsMultiple(identities []structures.DashboardWidgetsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateDashboardWidgets_Success
func TestCreateDashboardWidgets_Success(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_dashboardWidgets", controller.CreateDashboardWidgets)

	inputData := structures.DashboardWidgetsForm{}
	expectedResult := inputData
	mockService.On("CreateDashboardWidgets", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_dashboardWidgets", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.DashboardWidgetsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateDashboardWidgets_ServiceError
func TestCreateDashboardWidgets_ServiceError(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_dashboardWidgets", controller.CreateDashboardWidgets)

	inputData := structures.DashboardWidgetsForm{}
	mockService.On("CreateDashboardWidgets", inputData).Return(structures.DashboardWidgetsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_dashboardWidgets", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteDashboardWidgetsMultiple_Success(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_dashboardWidgets/bulk", controller.DeleteDashboardWidgetsMultiple)

	identities := []structures.DashboardWidgetsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteDashboardWidgetsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_dashboardWidgets/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteDashboardWidgetsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_dashboardWidgets/bulk", controller.DeleteDashboardWidgetsMultiple)

	identities := []structures.DashboardWidgetsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteDashboardWidgetsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_dashboardWidgets/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteDashboardWidgetsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_dashboardWidgets/bulk", controller.DeleteDashboardWidgetsMultiple)

	var emptyIdentities []structures.DashboardWidgetsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_dashboardWidgets/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteDashboardWidgetsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateDashboardWidgetsMultiple_Success(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_dashboardWidgets/bulk", controller.UpdateDashboardWidgetsMultiple)

	batchUpdatePayload := []structures.DashboardWidgetsBatchUpdate{
		{
			PathParams: structures.DashboardWidgetsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.DashboardWidgetsEdit{Name: dashboardWidgetsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.DashboardWidgetsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.DashboardWidgetsEdit{Name: dashboardWidgetsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateDashboardWidgetsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_dashboardWidgets/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateDashboardWidgetsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_dashboardWidgets/bulk", controller.UpdateDashboardWidgetsMultiple)

	batchUpdatePayload := []structures.DashboardWidgetsBatchUpdate{
		{PathParams: structures.DashboardWidgetsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.DashboardWidgetsEdit{Name: dashboardWidgetsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateDashboardWidgetsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_dashboardWidgets/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateDashboardWidgetsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_dashboardWidgets/bulk", controller.UpdateDashboardWidgetsMultiple)

	var emptyPayload []structures.DashboardWidgetsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_dashboardWidgets/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateDashboardWidgetsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde DashboardWidgetsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func dashboardWidgetsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func dashboardWidgetsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteDashboardWidgetsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockDashboardWidgetsService)
	controller := DashboardWidgetsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_dashboardWidgets/bulk", controller.DeleteDashboardWidgetsMultiple)

	req := httptest.NewRequest("DELETE", "/test_dashboardWidgets/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteDashboardWidgetsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateDashboardWidgets(pk, editData) mock'lanır; Delete için c.Svc.DeleteDashboardWidgets(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetDashboardWidgetsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetDashboardWidgetsWithPagination_Success(t *testing.T) {
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
