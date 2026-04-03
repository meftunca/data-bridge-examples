package orders_api_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	testhelper "data-bridge-examples/api_v2/orders/testhelper"

	"github.com/gofiber/fiber/v2"
	"github.com/maple-tech/baseline/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// Projenizin doğru import path'lerini kullandığınızdan emin olun
	structures "data-bridge-examples/api_v2/orders/structures"

	"github.com/maple-tech/baseline/types" // types.ID, types.URID için
	"github.com/maple-tech/baseline/web"   // web.Fiber200Response, web.Fiber400Response için
)

var (
	_ = types.ID(0)
	_ = types.URID("")
)

// --- Mock Service Definition ---
// MockCustomersService, ICustomersService arayüzünü taklit eder.
type MockCustomersService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockCustomersService) CreateCustomers(data structures.CustomersForm) (structures.CustomersForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.CustomersForm{}, args.Error(1)
	}
	return args.Get(0).(structures.CustomersForm), args.Error(1)
}

func (m *MockCustomersService) CreateCustomersMultiple(data []structures.CustomersForm) ([]structures.CustomersForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.CustomersForm), args.Error(1)
}

func (m *MockCustomersService) UpdateCustomers(id types.URID, data structures.CustomersEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockCustomersService) UpdateCustomersMultiple(data []structures.CustomersBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockCustomersService) DeleteCustomers(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockCustomersService) DeleteCustomersMultiple(identities []structures.CustomersIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreateCustomers_Success
func TestCreateCustomers_Success(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_customers", controller.CreateCustomers)

	inputData := structures.CustomersForm{}
	expectedResult := inputData
	mockService.On("CreateCustomers", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_customers", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.CustomersForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreateCustomers_ServiceError
func TestCreateCustomers_ServiceError(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_customers", controller.CreateCustomers)

	inputData := structures.CustomersForm{}
	mockService.On("CreateCustomers", inputData).Return(structures.CustomersForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_customers", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeleteCustomersMultiple_Success(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_customers/bulk", controller.DeleteCustomersMultiple)

	identities := []structures.CustomersIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeleteCustomersMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_customers/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeleteCustomersMultiple_ServiceError(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_customers/bulk", controller.DeleteCustomersMultiple)

	identities := []structures.CustomersIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeleteCustomersMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_customers/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteCustomersMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_customers/bulk", controller.DeleteCustomersMultiple)

	var emptyIdentities []structures.CustomersIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_customers/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteCustomersMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdateCustomersMultiple_Success(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_customers/bulk", controller.UpdateCustomersMultiple)

	batchUpdatePayload := []structures.CustomersBatchUpdate{
		{
			PathParams: structures.CustomersIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.CustomersEdit{Name: customersStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.CustomersIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.CustomersEdit{Name: customersStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdateCustomersMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_customers/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdateCustomersMultiple_ServiceError(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_customers/bulk", controller.UpdateCustomersMultiple)

	batchUpdatePayload := []structures.CustomersBatchUpdate{
		{PathParams: structures.CustomersIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.CustomersEdit{Name: customersStringPtr("Fail Update")}},
	}
	mockService.On("UpdateCustomersMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_customers/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdateCustomersMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_customers/bulk", controller.UpdateCustomersMultiple)

	var emptyPayload []structures.CustomersBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_customers/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdateCustomersMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde CustomersEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func customersStringPtr(s string) *string { return testhelper.StringPtr(s) }
func customersIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeleteCustomersMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockCustomersService)
	controller := CustomersController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_customers/bulk", controller.DeleteCustomersMultiple)

	req := httptest.NewRequest("DELETE", "/test_customers/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeleteCustomersMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdateCustomers(pk, editData) mock'lanır; Delete için c.Svc.DeleteCustomers(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetCustomersWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetCustomersWithPagination_Success(t *testing.T) {
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
