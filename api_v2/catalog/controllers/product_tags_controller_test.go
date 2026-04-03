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
// MockProductTagsService, IProductTagsService arayüzünü taklit eder.
type MockProductTagsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockProductTagsService) CreateProductTags(data structures.ProductTagsForm) (structures.ProductTagsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ProductTagsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ProductTagsForm), args.Error(1)
}

func (m *MockProductTagsService) CreateProductTagsMultiple(data []structures.ProductTagsForm) ([]structures.ProductTagsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ProductTagsForm), args.Error(1)
}

func (m *MockProductTagsService) UpdateProductTags(productId types.URID, tagId types.URID, data structures.ProductTagsEdit) error {
	args := m.Called(productId, tagId, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockProductTagsService) UpdateProductTagsMultiple(data []structures.ProductTagsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockProductTagsService) DeleteProductTags(productId types.URID, tagId types.URID) error {
	args := m.Called(productId, tagId) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockProductTagsService) DeleteProductTagsMultiple(identities []structures.ProductTagsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateProductTags_Success
func TestCreateProductTags_Success(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_productTags", controller.CreateProductTags)

	inputData := structures.ProductTagsForm{}
	expectedResult := inputData
	mockService.On("CreateProductTags", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_productTags", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ProductTagsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateProductTags_ServiceError
func TestCreateProductTags_ServiceError(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_productTags", controller.CreateProductTags)

	inputData := structures.ProductTagsForm{}
	mockService.On("CreateProductTags", inputData).Return(structures.ProductTagsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_productTags", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteProductTagsMultiple_Success(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_productTags/bulk", controller.DeleteProductTagsMultiple)

	identities := []structures.ProductTagsIdentity{
		{ProductId: types.URID("00000000-0000-0000-0000-000000000001")},
		{ProductId: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteProductTagsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_productTags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteProductTagsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_productTags/bulk", controller.DeleteProductTagsMultiple)

	identities := []structures.ProductTagsIdentity{{ProductId: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteProductTagsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_productTags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteProductTagsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_productTags/bulk", controller.DeleteProductTagsMultiple)

	var emptyIdentities []structures.ProductTagsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_productTags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteProductTagsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateProductTagsMultiple_Success(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_productTags/bulk", controller.UpdateProductTagsMultiple)

	batchUpdatePayload := []structures.ProductTagsBatchUpdate{
		{
			PathParams: structures.ProductTagsIdentity{ProductId: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ProductTagsEdit{Name: productTagsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ProductTagsIdentity{ProductId: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ProductTagsEdit{Name: productTagsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateProductTagsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_productTags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateProductTagsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_productTags/bulk", controller.UpdateProductTagsMultiple)

	batchUpdatePayload := []structures.ProductTagsBatchUpdate{
		{PathParams: structures.ProductTagsIdentity{ProductId: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ProductTagsEdit{Name: productTagsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateProductTagsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_productTags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateProductTagsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_productTags/bulk", controller.UpdateProductTagsMultiple)

	var emptyPayload []structures.ProductTagsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_productTags/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateProductTagsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ProductTagsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func productTagsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func productTagsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteProductTagsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockProductTagsService)
	controller := ProductTagsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_productTags/bulk", controller.DeleteProductTagsMultiple)

	req := httptest.NewRequest("DELETE", "/test_productTags/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteProductTagsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateProductTags(pk, editData) mock'lanır; Delete için c.Svc.DeleteProductTags(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetProductTagsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetProductTagsWithPagination_Success(t *testing.T) {
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
