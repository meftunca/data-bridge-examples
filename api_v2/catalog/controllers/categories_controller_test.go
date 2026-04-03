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
// MockCategoriesService, ICategoriesService arayüzünü taklit eder.
type MockCategoriesService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockCategoriesService) CreateCategories(data structures.CategoriesForm) (structures.CategoriesForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.CategoriesForm{}, args.Error(1)
	}
	return args.Get(0).(structures.CategoriesForm), args.Error(1)
}

func (m *MockCategoriesService) CreateCategoriesMultiple(data []structures.CategoriesForm) ([]structures.CategoriesForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.CategoriesForm), args.Error(1)
}

func (m *MockCategoriesService) UpdateCategories(id types.URID, data structures.CategoriesEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockCategoriesService) UpdateCategoriesMultiple(data []structures.CategoriesBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockCategoriesService) DeleteCategories(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockCategoriesService) DeleteCategoriesMultiple(identities []structures.CategoriesIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateCategories_Success
func TestCreateCategories_Success(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_categories", controller.CreateCategories)

	inputData := structures.CategoriesForm{}
	expectedResult := inputData
	mockService.On("CreateCategories", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_categories", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.CategoriesForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateCategories_ServiceError
func TestCreateCategories_ServiceError(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_categories", controller.CreateCategories)

	inputData := structures.CategoriesForm{}
	mockService.On("CreateCategories", inputData).Return(structures.CategoriesForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_categories", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteCategoriesMultiple_Success(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_categories/bulk", controller.DeleteCategoriesMultiple)

	identities := []structures.CategoriesIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteCategoriesMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_categories/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteCategoriesMultiple_ServiceError(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_categories/bulk", controller.DeleteCategoriesMultiple)

	identities := []structures.CategoriesIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteCategoriesMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_categories/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteCategoriesMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_categories/bulk", controller.DeleteCategoriesMultiple)

	var emptyIdentities []structures.CategoriesIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_categories/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteCategoriesMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateCategoriesMultiple_Success(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_categories/bulk", controller.UpdateCategoriesMultiple)

	batchUpdatePayload := []structures.CategoriesBatchUpdate{
		{
			PathParams: structures.CategoriesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.CategoriesEdit{Name: categoriesStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.CategoriesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.CategoriesEdit{Name: categoriesStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateCategoriesMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_categories/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateCategoriesMultiple_ServiceError(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_categories/bulk", controller.UpdateCategoriesMultiple)

	batchUpdatePayload := []structures.CategoriesBatchUpdate{
		{PathParams: structures.CategoriesIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.CategoriesEdit{Name: categoriesStringPtr("Fail Update")}},
	}
	mockService.On("UpdateCategoriesMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_categories/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateCategoriesMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_categories/bulk", controller.UpdateCategoriesMultiple)

	var emptyPayload []structures.CategoriesBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_categories/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateCategoriesMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde CategoriesEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func categoriesStringPtr(s string) *string { return testhelper.StringPtr(s) }
func categoriesIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteCategoriesMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockCategoriesService)
	controller := CategoriesController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_categories/bulk", controller.DeleteCategoriesMultiple)

	req := httptest.NewRequest("DELETE", "/test_categories/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteCategoriesMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateCategories(pk, editData) mock'lanır; Delete için c.Svc.DeleteCategories(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetCategoriesWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetCategoriesWithPagination_Success(t *testing.T) {
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
