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
// MockInvitationsService, IInvitationsService arayüzünü taklit eder.
type MockInvitationsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockInvitationsService) CreateInvitations(data structures.InvitationsForm) (structures.InvitationsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.InvitationsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.InvitationsForm), args.Error(1)
}

func (m *MockInvitationsService) CreateInvitationsMultiple(data []structures.InvitationsForm) ([]structures.InvitationsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.InvitationsForm), args.Error(1)
}

func (m *MockInvitationsService) UpdateInvitations(id types.URID, data structures.InvitationsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockInvitationsService) UpdateInvitationsMultiple(data []structures.InvitationsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockInvitationsService) DeleteInvitations(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockInvitationsService) DeleteInvitationsMultiple(identities []structures.InvitationsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateInvitations_Success
func TestCreateInvitations_Success(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_invitations", controller.CreateInvitations)

	inputData := structures.InvitationsForm{}
	expectedResult := inputData
	mockService.On("CreateInvitations", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_invitations", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.InvitationsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateInvitations_ServiceError
func TestCreateInvitations_ServiceError(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_invitations", controller.CreateInvitations)

	inputData := structures.InvitationsForm{}
	mockService.On("CreateInvitations", inputData).Return(structures.InvitationsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_invitations", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteInvitationsMultiple_Success(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_invitations/bulk", controller.DeleteInvitationsMultiple)

	identities := []structures.InvitationsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteInvitationsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_invitations/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteInvitationsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_invitations/bulk", controller.DeleteInvitationsMultiple)

	identities := []structures.InvitationsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteInvitationsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_invitations/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteInvitationsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_invitations/bulk", controller.DeleteInvitationsMultiple)

	var emptyIdentities []structures.InvitationsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_invitations/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteInvitationsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateInvitationsMultiple_Success(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_invitations/bulk", controller.UpdateInvitationsMultiple)

	batchUpdatePayload := []structures.InvitationsBatchUpdate{
		{
			PathParams: structures.InvitationsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.InvitationsEdit{Name: invitationsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.InvitationsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.InvitationsEdit{Name: invitationsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateInvitationsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_invitations/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateInvitationsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_invitations/bulk", controller.UpdateInvitationsMultiple)

	batchUpdatePayload := []structures.InvitationsBatchUpdate{
		{PathParams: structures.InvitationsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.InvitationsEdit{Name: invitationsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateInvitationsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_invitations/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateInvitationsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_invitations/bulk", controller.UpdateInvitationsMultiple)

	var emptyPayload []structures.InvitationsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_invitations/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateInvitationsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde InvitationsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func invitationsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func invitationsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteInvitationsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockInvitationsService)
	controller := InvitationsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_invitations/bulk", controller.DeleteInvitationsMultiple)

	req := httptest.NewRequest("DELETE", "/test_invitations/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteInvitationsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateInvitations(pk, editData) mock'lanır; Delete için c.Svc.DeleteInvitations(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetInvitationsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetInvitationsWithPagination_Success(t *testing.T) {
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
