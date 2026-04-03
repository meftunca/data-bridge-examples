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
// MockUsersService, IUsersService arayüzünü taklit eder.
type MockUsersService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockUsersService) CreateUsers(data structures.UsersForm) (structures.UsersForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.UsersForm{}, args.Error(1)
	}
	return args.Get(0).(structures.UsersForm), args.Error(1)
}

func (m *MockUsersService) CreateUsersMultiple(data []structures.UsersForm) ([]structures.UsersForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.UsersForm), args.Error(1)
}

func (m *MockUsersService) UpdateUsers(id types.URID, data structures.UsersEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockUsersService) UpdateUsersMultiple(data []structures.UsersBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockUsersService) DeleteUsers(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockUsersService) DeleteUsersMultiple(identities []structures.UsersIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateUsers_Success
func TestCreateUsers_Success(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_users", controller.CreateUsers)

	inputData := structures.UsersForm{}
	expectedResult := inputData
	mockService.On("CreateUsers", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.UsersForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateUsers_ServiceError
func TestCreateUsers_ServiceError(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_users", controller.CreateUsers)

	inputData := structures.UsersForm{}
	mockService.On("CreateUsers", inputData).Return(structures.UsersForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteUsersMultiple_Success(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_users/bulk", controller.DeleteUsersMultiple)

	identities := []structures.UsersIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteUsersMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_users/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteUsersMultiple_ServiceError(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_users/bulk", controller.DeleteUsersMultiple)

	identities := []structures.UsersIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteUsersMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_users/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteUsersMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_users/bulk", controller.DeleteUsersMultiple)

	var emptyIdentities []structures.UsersIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_users/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteUsersMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateUsersMultiple_Success(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_users/bulk", controller.UpdateUsersMultiple)

	batchUpdatePayload := []structures.UsersBatchUpdate{
		{
			PathParams: structures.UsersIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.UsersEdit{Name: usersStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.UsersIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.UsersEdit{Name: usersStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateUsersMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_users/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateUsersMultiple_ServiceError(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_users/bulk", controller.UpdateUsersMultiple)

	batchUpdatePayload := []structures.UsersBatchUpdate{
		{PathParams: structures.UsersIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.UsersEdit{Name: usersStringPtr("Fail Update")}},
	}
	mockService.On("UpdateUsersMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_users/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateUsersMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_users/bulk", controller.UpdateUsersMultiple)

	var emptyPayload []structures.UsersBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_users/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateUsersMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde UsersEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func usersStringPtr(s string) *string { return testhelper.StringPtr(s) }
func usersIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteUsersMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockUsersService)
	controller := UsersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_users/bulk", controller.DeleteUsersMultiple)

	req := httptest.NewRequest("DELETE", "/test_users/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteUsersMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateUsers(pk, editData) mock'lanır; Delete için c.Svc.DeleteUsers(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetUsersWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetUsersWithPagination_Success(t *testing.T) {
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
