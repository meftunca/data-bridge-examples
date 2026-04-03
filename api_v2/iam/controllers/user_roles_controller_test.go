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
// MockUserRolesService, IUserRolesService arayüzünü taklit eder.
type MockUserRolesService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockUserRolesService) CreateUserRoles(data structures.UserRolesForm) (structures.UserRolesForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.UserRolesForm{}, args.Error(1)
	}
	return args.Get(0).(structures.UserRolesForm), args.Error(1)
}

func (m *MockUserRolesService) CreateUserRolesMultiple(data []structures.UserRolesForm) ([]structures.UserRolesForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.UserRolesForm), args.Error(1)
}

func (m *MockUserRolesService) UpdateUserRoles(id types.URID, data structures.UserRolesEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockUserRolesService) UpdateUserRolesMultiple(data []structures.UserRolesBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockUserRolesService) DeleteUserRoles(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockUserRolesService) DeleteUserRolesMultiple(identities []structures.UserRolesIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateUserRoles_Success
func TestCreateUserRoles_Success(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_userRoles", controller.CreateUserRoles)

	inputData := structures.UserRolesForm{}
	expectedResult := inputData
	mockService.On("CreateUserRoles", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_userRoles", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.UserRolesForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateUserRoles_ServiceError
func TestCreateUserRoles_ServiceError(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_userRoles", controller.CreateUserRoles)

	inputData := structures.UserRolesForm{}
	mockService.On("CreateUserRoles", inputData).Return(structures.UserRolesForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_userRoles", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteUserRolesMultiple_Success(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_userRoles/bulk", controller.DeleteUserRolesMultiple)

	identities := []structures.UserRolesIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteUserRolesMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_userRoles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteUserRolesMultiple_ServiceError(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_userRoles/bulk", controller.DeleteUserRolesMultiple)

	identities := []structures.UserRolesIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteUserRolesMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_userRoles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteUserRolesMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_userRoles/bulk", controller.DeleteUserRolesMultiple)

	var emptyIdentities []structures.UserRolesIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_userRoles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteUserRolesMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateUserRolesMultiple_Success(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_userRoles/bulk", controller.UpdateUserRolesMultiple)

	batchUpdatePayload := []structures.UserRolesBatchUpdate{
		{
			PathParams: structures.UserRolesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.UserRolesEdit{Name: userRolesStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.UserRolesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.UserRolesEdit{Name: userRolesStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateUserRolesMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_userRoles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateUserRolesMultiple_ServiceError(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_userRoles/bulk", controller.UpdateUserRolesMultiple)

	batchUpdatePayload := []structures.UserRolesBatchUpdate{
		{PathParams: structures.UserRolesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.UserRolesEdit{Name: userRolesStringPtr("Fail Update")}},
	}
	mockService.On("UpdateUserRolesMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_userRoles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateUserRolesMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_userRoles/bulk", controller.UpdateUserRolesMultiple)

	var emptyPayload []structures.UserRolesBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_userRoles/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateUserRolesMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde UserRolesEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func userRolesStringPtr(s string) *string { return testhelper.StringPtr(s) }
func userRolesIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteUserRolesMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockUserRolesService)
	controller := UserRolesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_userRoles/bulk", controller.DeleteUserRolesMultiple)

	req := httptest.NewRequest("DELETE", "/test_userRoles/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteUserRolesMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateUserRoles(pk, editData) mock'lanır; Delete için c.Svc.DeleteUserRoles(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetUserRolesWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetUserRolesWithPagination_Success(t *testing.T) {
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
