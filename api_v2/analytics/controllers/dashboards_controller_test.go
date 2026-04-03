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
// MockDashboardsService, IDashboardsService arayüzünü taklit eder.
type MockDashboardsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockDashboardsService) CreateDashboards(data structures.DashboardsForm) (structures.DashboardsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.DashboardsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.DashboardsForm), args.Error(1)
}

func (m *MockDashboardsService) CreateDashboardsMultiple(data []structures.DashboardsForm) ([]structures.DashboardsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.DashboardsForm), args.Error(1)
}

func (m *MockDashboardsService) UpdateDashboards(id types.URID, data structures.DashboardsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockDashboardsService) UpdateDashboardsMultiple(data []structures.DashboardsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockDashboardsService) DeleteDashboards(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockDashboardsService) DeleteDashboardsMultiple(identities []structures.DashboardsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateDashboards_Success
func TestCreateDashboards_Success(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_dashboards", controller.CreateDashboards)

	inputData := structures.DashboardsForm{}
	expectedResult := inputData
	mockService.On("CreateDashboards", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_dashboards", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.DashboardsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateDashboards_ServiceError
func TestCreateDashboards_ServiceError(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_dashboards", controller.CreateDashboards)

	inputData := structures.DashboardsForm{}
	mockService.On("CreateDashboards", inputData).Return(structures.DashboardsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_dashboards", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteDashboardsMultiple_Success(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_dashboards/bulk", controller.DeleteDashboardsMultiple)

	identities := []structures.DashboardsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteDashboardsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_dashboards/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteDashboardsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_dashboards/bulk", controller.DeleteDashboardsMultiple)

	identities := []structures.DashboardsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteDashboardsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_dashboards/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteDashboardsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_dashboards/bulk", controller.DeleteDashboardsMultiple)

	var emptyIdentities []structures.DashboardsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_dashboards/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteDashboardsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateDashboardsMultiple_Success(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_dashboards/bulk", controller.UpdateDashboardsMultiple)

	batchUpdatePayload := []structures.DashboardsBatchUpdate{
		{
			PathParams: structures.DashboardsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.DashboardsEdit{Name: dashboardsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.DashboardsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.DashboardsEdit{Name: dashboardsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateDashboardsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_dashboards/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateDashboardsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_dashboards/bulk", controller.UpdateDashboardsMultiple)

	batchUpdatePayload := []structures.DashboardsBatchUpdate{
		{PathParams: structures.DashboardsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.DashboardsEdit{Name: dashboardsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateDashboardsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_dashboards/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateDashboardsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_dashboards/bulk", controller.UpdateDashboardsMultiple)

	var emptyPayload []structures.DashboardsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_dashboards/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateDashboardsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde DashboardsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func dashboardsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func dashboardsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteDashboardsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockDashboardsService)
	controller := DashboardsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_dashboards/bulk", controller.DeleteDashboardsMultiple)

	req := httptest.NewRequest("DELETE", "/test_dashboards/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteDashboardsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateDashboards(pk, editData) mock'lanır; Delete için c.Svc.DeleteDashboards(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetDashboardsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetDashboardsWithPagination_Success(t *testing.T) {
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
