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
// MockProductsService, IProductsService arayüzünü taklit eder.
type MockProductsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockProductsService) CreateProducts(data structures.ProductsForm) (structures.ProductsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ProductsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ProductsForm), args.Error(1)
}

func (m *MockProductsService) CreateProductsMultiple(data []structures.ProductsForm) ([]structures.ProductsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ProductsForm), args.Error(1)
}

func (m *MockProductsService) UpdateProducts(id types.URID, data structures.ProductsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockProductsService) UpdateProductsMultiple(data []structures.ProductsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockProductsService) DeleteProducts(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockProductsService) DeleteProductsMultiple(identities []structures.ProductsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateProducts_Success
func TestCreateProducts_Success(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_products", controller.CreateProducts)

	inputData := structures.ProductsForm{}
	expectedResult := inputData
	mockService.On("CreateProducts", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_products", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ProductsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateProducts_ServiceError
func TestCreateProducts_ServiceError(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_products", controller.CreateProducts)

	inputData := structures.ProductsForm{}
	mockService.On("CreateProducts", inputData).Return(structures.ProductsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_products", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteProductsMultiple_Success(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_products/bulk", controller.DeleteProductsMultiple)

	identities := []structures.ProductsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteProductsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_products/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteProductsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_products/bulk", controller.DeleteProductsMultiple)

	identities := []structures.ProductsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteProductsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_products/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteProductsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_products/bulk", controller.DeleteProductsMultiple)

	var emptyIdentities []structures.ProductsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_products/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteProductsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateProductsMultiple_Success(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_products/bulk", controller.UpdateProductsMultiple)

	batchUpdatePayload := []structures.ProductsBatchUpdate{
		{
			PathParams: structures.ProductsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ProductsEdit{Name: productsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ProductsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ProductsEdit{Name: productsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateProductsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_products/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateProductsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_products/bulk", controller.UpdateProductsMultiple)

	batchUpdatePayload := []structures.ProductsBatchUpdate{
		{PathParams: structures.ProductsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ProductsEdit{Name: productsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateProductsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_products/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateProductsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_products/bulk", controller.UpdateProductsMultiple)

	var emptyPayload []structures.ProductsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_products/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateProductsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ProductsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func productsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func productsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteProductsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockProductsService)
	controller := ProductsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_products/bulk", controller.DeleteProductsMultiple)

	req := httptest.NewRequest("DELETE", "/test_products/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteProductsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateProducts(pk, editData) mock'lanır; Delete için c.Svc.DeleteProducts(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetProductsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetProductsWithPagination_Success(t *testing.T) {
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
