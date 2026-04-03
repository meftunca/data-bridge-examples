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
// MockCollectionProductsService, ICollectionProductsService arayüzünü taklit eder.
type MockCollectionProductsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockCollectionProductsService) CreateCollectionProducts(data structures.CollectionProductsForm) (structures.CollectionProductsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.CollectionProductsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.CollectionProductsForm), args.Error(1)
}

func (m *MockCollectionProductsService) CreateCollectionProductsMultiple(data []structures.CollectionProductsForm) ([]structures.CollectionProductsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.CollectionProductsForm), args.Error(1)
}

func (m *MockCollectionProductsService) UpdateCollectionProducts(id types.URID, data structures.CollectionProductsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockCollectionProductsService) UpdateCollectionProductsMultiple(data []structures.CollectionProductsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockCollectionProductsService) DeleteCollectionProducts(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockCollectionProductsService) DeleteCollectionProductsMultiple(identities []structures.CollectionProductsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateCollectionProducts_Success
func TestCreateCollectionProducts_Success(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_collectionProducts", controller.CreateCollectionProducts)

	inputData := structures.CollectionProductsForm{}
	expectedResult := inputData
	mockService.On("CreateCollectionProducts", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_collectionProducts", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.CollectionProductsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateCollectionProducts_ServiceError
func TestCreateCollectionProducts_ServiceError(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_collectionProducts", controller.CreateCollectionProducts)

	inputData := structures.CollectionProductsForm{}
	mockService.On("CreateCollectionProducts", inputData).Return(structures.CollectionProductsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_collectionProducts", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteCollectionProductsMultiple_Success(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_collectionProducts/bulk", controller.DeleteCollectionProductsMultiple)

	identities := []structures.CollectionProductsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteCollectionProductsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_collectionProducts/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteCollectionProductsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_collectionProducts/bulk", controller.DeleteCollectionProductsMultiple)

	identities := []structures.CollectionProductsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteCollectionProductsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_collectionProducts/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteCollectionProductsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_collectionProducts/bulk", controller.DeleteCollectionProductsMultiple)

	var emptyIdentities []structures.CollectionProductsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_collectionProducts/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteCollectionProductsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateCollectionProductsMultiple_Success(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_collectionProducts/bulk", controller.UpdateCollectionProductsMultiple)

	batchUpdatePayload := []structures.CollectionProductsBatchUpdate{
		{
			PathParams: structures.CollectionProductsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.CollectionProductsEdit{Name: collectionProductsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.CollectionProductsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.CollectionProductsEdit{Name: collectionProductsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateCollectionProductsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_collectionProducts/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateCollectionProductsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_collectionProducts/bulk", controller.UpdateCollectionProductsMultiple)

	batchUpdatePayload := []structures.CollectionProductsBatchUpdate{
		{PathParams: structures.CollectionProductsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.CollectionProductsEdit{Name: collectionProductsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateCollectionProductsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_collectionProducts/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateCollectionProductsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_collectionProducts/bulk", controller.UpdateCollectionProductsMultiple)

	var emptyPayload []structures.CollectionProductsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_collectionProducts/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateCollectionProductsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde CollectionProductsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func collectionProductsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func collectionProductsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteCollectionProductsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockCollectionProductsService)
	controller := CollectionProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_collectionProducts/bulk", controller.DeleteCollectionProductsMultiple)

	req := httptest.NewRequest("DELETE", "/test_collectionProducts/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteCollectionProductsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateCollectionProducts(pk, editData) mock'lanır; Delete için c.Svc.DeleteCollectionProducts(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetCollectionProductsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetCollectionProductsWithPagination_Success(t *testing.T) {
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
