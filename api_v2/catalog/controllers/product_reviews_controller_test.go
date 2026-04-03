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
// MockProductReviewsService, IProductReviewsService arayüzünü taklit eder.
type MockProductReviewsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockProductReviewsService) CreateProductReviews(data structures.ProductReviewsForm) (structures.ProductReviewsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.ProductReviewsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.ProductReviewsForm), args.Error(1)
}

func (m *MockProductReviewsService) CreateProductReviewsMultiple(data []structures.ProductReviewsForm) ([]structures.ProductReviewsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.ProductReviewsForm), args.Error(1)
}

func (m *MockProductReviewsService) UpdateProductReviews(id types.URID, data structures.ProductReviewsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockProductReviewsService) UpdateProductReviewsMultiple(data []structures.ProductReviewsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockProductReviewsService) DeleteProductReviews(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockProductReviewsService) DeleteProductReviewsMultiple(identities []structures.ProductReviewsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateProductReviews_Success
func TestCreateProductReviews_Success(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_productReviews", controller.CreateProductReviews)

	inputData := structures.ProductReviewsForm{}
	expectedResult := inputData
	mockService.On("CreateProductReviews", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_productReviews", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.ProductReviewsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateProductReviews_ServiceError
func TestCreateProductReviews_ServiceError(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_productReviews", controller.CreateProductReviews)

	inputData := structures.ProductReviewsForm{}
	mockService.On("CreateProductReviews", inputData).Return(structures.ProductReviewsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_productReviews", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteProductReviewsMultiple_Success(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_productReviews/bulk", controller.DeleteProductReviewsMultiple)

	identities := []structures.ProductReviewsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteProductReviewsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_productReviews/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteProductReviewsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_productReviews/bulk", controller.DeleteProductReviewsMultiple)

	identities := []structures.ProductReviewsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteProductReviewsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_productReviews/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteProductReviewsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_productReviews/bulk", controller.DeleteProductReviewsMultiple)

	var emptyIdentities []structures.ProductReviewsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_productReviews/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteProductReviewsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateProductReviewsMultiple_Success(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_productReviews/bulk", controller.UpdateProductReviewsMultiple)

	batchUpdatePayload := []structures.ProductReviewsBatchUpdate{
		{
			PathParams: structures.ProductReviewsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.ProductReviewsEdit{Name: productReviewsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.ProductReviewsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.ProductReviewsEdit{Name: productReviewsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateProductReviewsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_productReviews/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateProductReviewsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_productReviews/bulk", controller.UpdateProductReviewsMultiple)

	batchUpdatePayload := []structures.ProductReviewsBatchUpdate{
		{PathParams: structures.ProductReviewsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.ProductReviewsEdit{Name: productReviewsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateProductReviewsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_productReviews/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateProductReviewsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_productReviews/bulk", controller.UpdateProductReviewsMultiple)

	var emptyPayload []structures.ProductReviewsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_productReviews/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateProductReviewsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde ProductReviewsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func productReviewsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func productReviewsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteProductReviewsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockProductReviewsService)
	controller := ProductReviewsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_productReviews/bulk", controller.DeleteProductReviewsMultiple)

	req := httptest.NewRequest("DELETE", "/test_productReviews/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteProductReviewsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateProductReviews(pk, editData) mock'lanır; Delete için c.Svc.DeleteProductReviews(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetProductReviewsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetProductReviewsWithPagination_Success(t *testing.T) {
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
