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
// MockPermissionsService, IPermissionsService arayüzünü taklit eder.
type MockPermissionsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockPermissionsService) CreatePermissions(data structures.PermissionsForm) (structures.PermissionsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.PermissionsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.PermissionsForm), args.Error(1)
}

func (m *MockPermissionsService) CreatePermissionsMultiple(data []structures.PermissionsForm) ([]structures.PermissionsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.PermissionsForm), args.Error(1)
}

func (m *MockPermissionsService) UpdatePermissions(id types.URID, data structures.PermissionsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockPermissionsService) UpdatePermissionsMultiple(data []structures.PermissionsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockPermissionsService) DeletePermissions(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockPermissionsService) DeletePermissionsMultiple(identities []structures.PermissionsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreatePermissions_Success
func TestCreatePermissions_Success(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_permissions", controller.CreatePermissions)

	inputData := structures.PermissionsForm{}
	expectedResult := inputData
	mockService.On("CreatePermissions", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_permissions", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.PermissionsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreatePermissions_ServiceError
func TestCreatePermissions_ServiceError(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_permissions", controller.CreatePermissions)

	inputData := structures.PermissionsForm{}
	mockService.On("CreatePermissions", inputData).Return(structures.PermissionsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_permissions", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeletePermissionsMultiple_Success(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_permissions/bulk", controller.DeletePermissionsMultiple)

	identities := []structures.PermissionsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeletePermissionsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_permissions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeletePermissionsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_permissions/bulk", controller.DeletePermissionsMultiple)

	identities := []structures.PermissionsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeletePermissionsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_permissions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeletePermissionsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_permissions/bulk", controller.DeletePermissionsMultiple)

	var emptyIdentities []structures.PermissionsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_permissions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeletePermissionsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdatePermissionsMultiple_Success(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_permissions/bulk", controller.UpdatePermissionsMultiple)

	batchUpdatePayload := []structures.PermissionsBatchUpdate{
		{
			PathParams: structures.PermissionsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.PermissionsEdit{Name: permissionsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.PermissionsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.PermissionsEdit{Name: permissionsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdatePermissionsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_permissions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdatePermissionsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_permissions/bulk", controller.UpdatePermissionsMultiple)

	batchUpdatePayload := []structures.PermissionsBatchUpdate{
		{PathParams: structures.PermissionsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.PermissionsEdit{Name: permissionsStringPtr("Fail Update")}},
	}
	mockService.On("UpdatePermissionsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_permissions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdatePermissionsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_permissions/bulk", controller.UpdatePermissionsMultiple)

	var emptyPayload []structures.PermissionsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_permissions/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdatePermissionsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde PermissionsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func permissionsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func permissionsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeletePermissionsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockPermissionsService)
	controller := PermissionsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_permissions/bulk", controller.DeletePermissionsMultiple)

	req := httptest.NewRequest("DELETE", "/test_permissions/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeletePermissionsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdatePermissions(pk, editData) mock'lanır; Delete için c.Svc.DeletePermissions(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetPermissionsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetPermissionsWithPagination_Success(t *testing.T) {
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
