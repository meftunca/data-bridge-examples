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
// MockTeamsService, ITeamsService arayüzünü taklit eder.
type MockTeamsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockTeamsService) CreateTeams(data structures.TeamsForm) (structures.TeamsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.TeamsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.TeamsForm), args.Error(1)
}

func (m *MockTeamsService) CreateTeamsMultiple(data []structures.TeamsForm) ([]structures.TeamsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.TeamsForm), args.Error(1)
}

func (m *MockTeamsService) UpdateTeams(id types.URID, data structures.TeamsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockTeamsService) UpdateTeamsMultiple(data []structures.TeamsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockTeamsService) DeleteTeams(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockTeamsService) DeleteTeamsMultiple(identities []structures.TeamsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateTeams_Success
func TestCreateTeams_Success(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_teams", controller.CreateTeams)

	inputData := structures.TeamsForm{}
	expectedResult := inputData
	mockService.On("CreateTeams", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_teams", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.TeamsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateTeams_ServiceError
func TestCreateTeams_ServiceError(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_teams", controller.CreateTeams)

	inputData := structures.TeamsForm{}
	mockService.On("CreateTeams", inputData).Return(structures.TeamsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_teams", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteTeamsMultiple_Success(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_teams/bulk", controller.DeleteTeamsMultiple)

	identities := []structures.TeamsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteTeamsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_teams/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteTeamsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_teams/bulk", controller.DeleteTeamsMultiple)

	identities := []structures.TeamsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteTeamsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_teams/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteTeamsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_teams/bulk", controller.DeleteTeamsMultiple)

	var emptyIdentities []structures.TeamsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_teams/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteTeamsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateTeamsMultiple_Success(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_teams/bulk", controller.UpdateTeamsMultiple)

	batchUpdatePayload := []structures.TeamsBatchUpdate{
		{
			PathParams: structures.TeamsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.TeamsEdit{Name: teamsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.TeamsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.TeamsEdit{Name: teamsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateTeamsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_teams/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateTeamsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_teams/bulk", controller.UpdateTeamsMultiple)

	batchUpdatePayload := []structures.TeamsBatchUpdate{
		{PathParams: structures.TeamsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.TeamsEdit{Name: teamsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateTeamsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_teams/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateTeamsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_teams/bulk", controller.UpdateTeamsMultiple)

	var emptyPayload []structures.TeamsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_teams/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateTeamsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde TeamsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func teamsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func teamsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteTeamsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockTeamsService)
	controller := TeamsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_teams/bulk", controller.DeleteTeamsMultiple)

	req := httptest.NewRequest("DELETE", "/test_teams/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteTeamsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateTeams(pk, editData) mock'lanır; Delete için c.Svc.DeleteTeams(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetTeamsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetTeamsWithPagination_Success(t *testing.T) {
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
