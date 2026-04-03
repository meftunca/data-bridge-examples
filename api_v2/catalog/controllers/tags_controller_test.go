package catalog_api_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	testhelper "data-bridge-examples/api_v2/catalog/testhelper"

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// Projenizin doğru import path'lerini kullandığınızdan emin olun
	structures "data-bridge-examples/api_v2/catalog/structures"

	"github.com/maple-tech/baseline/types" // types.ID, types.URID için
	"github.com/maple-tech/baseline/web"   // web.Fiber200Response, web.Fiber400Response için
)

var (
	_ = types.ID(0)
	_ = types.URID("")
)

// --- Mock Service Definition ---
// MockTagsService, ITagsService arayüzünü taklit eder.
type MockTagsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockTagsService) CreateTags(data structures.TagsForm) (structures.TagsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.TagsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.TagsForm), args.Error(1)
}

func (m *MockTagsService) CreateTagsMultiple(data []structures.TagsForm) ([]structures.TagsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.TagsForm), args.Error(1)
}

func (m *MockTagsService) UpdateTags(id types.URID, data structures.TagsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockTagsService) UpdateTagsMultiple(data []structures.TagsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockTagsService) DeleteTags(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockTagsService) DeleteTagsMultiple(identities []structures.TagsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateTags_Success
func TestCreateTags_Success(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_tags", controller.CreateTags)

	inputData := structures.TagsForm{}
	expectedResult := inputData
	mockService.On("CreateTags", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_tags", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.TagsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateTags_ServiceError
func TestCreateTags_ServiceError(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_tags", controller.CreateTags)

	inputData := structures.TagsForm{}
	mockService.On("CreateTags", inputData).Return(structures.TagsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_tags", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteTagsMultiple_Success(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_tags/bulk", controller.DeleteTagsMultiple)

	identities := []structures.TagsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteTagsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_tags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteTagsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_tags/bulk", controller.DeleteTagsMultiple)

	identities := []structures.TagsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteTagsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_tags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteTagsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_tags/bulk", controller.DeleteTagsMultiple)

	var emptyIdentities []structures.TagsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_tags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteTagsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateTagsMultiple_Success(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_tags/bulk", controller.UpdateTagsMultiple)

	batchUpdatePayload := []structures.TagsBatchUpdate{
		{
			PathParams: structures.TagsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.TagsEdit{Name: tagsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.TagsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.TagsEdit{Name: tagsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateTagsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_tags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateTagsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_tags/bulk", controller.UpdateTagsMultiple)

	batchUpdatePayload := []structures.TagsBatchUpdate{
		{PathParams: structures.TagsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.TagsEdit{Name: tagsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateTagsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_tags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateTagsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_tags/bulk", controller.UpdateTagsMultiple)

	var emptyPayload []structures.TagsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_tags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateTagsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde TagsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func tagsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func tagsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteTagsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockTagsService)
	controller := TagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_tags/bulk", controller.DeleteTagsMultiple)

	req := httptest.NewRequest("DELETE", "/test_tags/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteTagsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateTags(pk, editData) mock'lanır; Delete için c.Svc.DeleteTags(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetTagsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetTagsWithPagination_Success(t *testing.T) {
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
