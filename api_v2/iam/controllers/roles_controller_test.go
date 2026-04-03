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
// MockRolesService, IRolesService arayüzünü taklit eder.
type MockRolesService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockRolesService) CreateRoles(data structures.RolesForm) (structures.RolesForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.RolesForm{}, args.Error(1)
	}
	return args.Get(0).(structures.RolesForm), args.Error(1)
}

func (m *MockRolesService) CreateRolesMultiple(data []structures.RolesForm) ([]structures.RolesForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.RolesForm), args.Error(1)
}

func (m *MockRolesService) UpdateRoles(id types.URID, data structures.RolesEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockRolesService) UpdateRolesMultiple(data []structures.RolesBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockRolesService) DeleteRoles(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockRolesService) DeleteRolesMultiple(identities []structures.RolesIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateRoles_Success
func TestCreateRoles_Success(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_roles", controller.CreateRoles)

	inputData := structures.RolesForm{}
	expectedResult := inputData
	mockService.On("CreateRoles", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_roles", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.RolesForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateRoles_ServiceError
func TestCreateRoles_ServiceError(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_roles", controller.CreateRoles)

	inputData := structures.RolesForm{}
	mockService.On("CreateRoles", inputData).Return(structures.RolesForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_roles", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteRolesMultiple_Success(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_roles/bulk", controller.DeleteRolesMultiple)

	identities := []structures.RolesIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteRolesMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_roles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteRolesMultiple_ServiceError(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_roles/bulk", controller.DeleteRolesMultiple)

	identities := []structures.RolesIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteRolesMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_roles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteRolesMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_roles/bulk", controller.DeleteRolesMultiple)

	var emptyIdentities []structures.RolesIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_roles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteRolesMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateRolesMultiple_Success(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_roles/bulk", controller.UpdateRolesMultiple)

	batchUpdatePayload := []structures.RolesBatchUpdate{
		{
			PathParams: structures.RolesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.RolesEdit{Name: rolesStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.RolesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.RolesEdit{Name: rolesStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateRolesMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_roles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateRolesMultiple_ServiceError(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_roles/bulk", controller.UpdateRolesMultiple)

	batchUpdatePayload := []structures.RolesBatchUpdate{
		{PathParams: structures.RolesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.RolesEdit{Name: rolesStringPtr("Fail Update")}},
	}
	mockService.On("UpdateRolesMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_roles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateRolesMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_roles/bulk", controller.UpdateRolesMultiple)

	var emptyPayload []structures.RolesBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_roles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateRolesMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde RolesEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func rolesStringPtr(s string) *string { return testhelper.StringPtr(s) }
func rolesIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteRolesMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockRolesService)
	controller := RolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_roles/bulk", controller.DeleteRolesMultiple)

	req := httptest.NewRequest("DELETE", "/test_roles/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteRolesMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateRoles(pk, editData) mock'lanır; Delete için c.Svc.DeleteRoles(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetRolesWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetRolesWithPagination_Success(t *testing.T) {
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
