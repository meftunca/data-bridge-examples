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
// MockBrandsService, IBrandsService arayüzünü taklit eder.
type MockBrandsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockBrandsService) CreateBrands(data structures.BrandsForm) (structures.BrandsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.BrandsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.BrandsForm), args.Error(1)
}

func (m *MockBrandsService) CreateBrandsMultiple(data []structures.BrandsForm) ([]structures.BrandsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.BrandsForm), args.Error(1)
}

func (m *MockBrandsService) UpdateBrands(id types.URID, data structures.BrandsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockBrandsService) UpdateBrandsMultiple(data []structures.BrandsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockBrandsService) DeleteBrands(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockBrandsService) DeleteBrandsMultiple(identities []structures.BrandsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateBrands_Success
func TestCreateBrands_Success(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_brands", controller.CreateBrands)

	inputData := structures.BrandsForm{}
	expectedResult := inputData
	mockService.On("CreateBrands", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_brands", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.BrandsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateBrands_ServiceError
func TestCreateBrands_ServiceError(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_brands", controller.CreateBrands)

	inputData := structures.BrandsForm{}
	mockService.On("CreateBrands", inputData).Return(structures.BrandsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_brands", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteBrandsMultiple_Success(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_brands/bulk", controller.DeleteBrandsMultiple)

	identities := []structures.BrandsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteBrandsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_brands/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteBrandsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_brands/bulk", controller.DeleteBrandsMultiple)

	identities := []structures.BrandsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteBrandsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_brands/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteBrandsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_brands/bulk", controller.DeleteBrandsMultiple)

	var emptyIdentities []structures.BrandsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_brands/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteBrandsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateBrandsMultiple_Success(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_brands/bulk", controller.UpdateBrandsMultiple)

	batchUpdatePayload := []structures.BrandsBatchUpdate{
		{
			PathParams: structures.BrandsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.BrandsEdit{Name: brandsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.BrandsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.BrandsEdit{Name: brandsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateBrandsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_brands/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateBrandsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_brands/bulk", controller.UpdateBrandsMultiple)

	batchUpdatePayload := []structures.BrandsBatchUpdate{
		{PathParams: structures.BrandsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.BrandsEdit{Name: brandsStringPtr("Fail Update")}},
	}
	mockService.On("UpdateBrandsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_brands/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateBrandsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_brands/bulk", controller.UpdateBrandsMultiple)

	var emptyPayload []structures.BrandsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_brands/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateBrandsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde BrandsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func brandsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func brandsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteBrandsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockBrandsService)
	controller := BrandsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_brands/bulk", controller.DeleteBrandsMultiple)

	req := httptest.NewRequest("DELETE", "/test_brands/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteBrandsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateBrands(pk, editData) mock'lanır; Delete için c.Svc.DeleteBrands(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetBrandsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetBrandsWithPagination_Success(t *testing.T) {
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
