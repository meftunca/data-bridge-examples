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
// MockEventsService, IEventsService arayüzünü taklit eder.
type MockEventsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockEventsService) CreateEvents(data structures.EventsForm) (structures.EventsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.EventsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.EventsForm), args.Error(1)
}

func (m *MockEventsService) CreateEventsMultiple(data []structures.EventsForm) ([]structures.EventsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.EventsForm), args.Error(1)
}

func (m *MockEventsService) UpdateEvents(id types.URID, data structures.EventsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockEventsService) UpdateEventsMultiple(data []structures.EventsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockEventsService) DeleteEvents(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockEventsService) DeleteEventsMultiple(identities []structures.EventsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateEvents_Success
func TestCreateEvents_Success(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_events", controller.CreateEvents)

	inputData := structures.EventsForm{}
	expectedResult := inputData
	mockService.On("CreateEvents", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_events", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.EventsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateEvents_ServiceError
func TestCreateEvents_ServiceError(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_events", controller.CreateEvents)

	inputData := structures.EventsForm{}
	mockService.On("CreateEvents", inputData).Return(structures.EventsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_events", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteEventsMultiple_Success(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_events/bulk", controller.DeleteEventsMultiple)

	identities := []structures.EventsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteEventsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_events/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteEventsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_events/bulk", controller.DeleteEventsMultiple)

	identities := []structures.EventsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteEventsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_events/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteEventsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_events/bulk", controller.DeleteEventsMultiple)

	var emptyIdentities []structures.EventsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_events/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteEventsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateEventsMultiple_Success(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_events/bulk", controller.UpdateEventsMultiple)

	batchUpdatePayload := []structures.EventsBatchUpdate{
		{
			PathParams: structures.EventsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.EventsEdit{Name: eventsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.EventsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.EventsEdit{Name: eventsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateEventsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_events/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateEventsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_events/bulk", controller.UpdateEventsMultiple)

	batchUpdatePayload := []structures.EventsBatchUpdate{
		{PathParams: structures.EventsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.EventsEdit{Name: eventsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateEventsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_events/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateEventsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_events/bulk", controller.UpdateEventsMultiple)

	var emptyPayload []structures.EventsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_events/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateEventsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde EventsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func eventsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func eventsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteEventsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockEventsService)
	controller := EventsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_events/bulk", controller.DeleteEventsMultiple)

	req := httptest.NewRequest("DELETE", "/test_events/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteEventsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateEvents(pk, editData) mock'lanır; Delete için c.Svc.DeleteEvents(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetEventsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetEventsWithPagination_Success(t *testing.T) {
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
