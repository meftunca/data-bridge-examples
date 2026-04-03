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
// MockPaymentsService, IPaymentsService arayüzünü taklit eder.
type MockPaymentsService struct {
	mock.Mock
}

// --- Mock Metod Implementasyonları (DÜZELTİLMİŞ) ---

func (m *MockPaymentsService) CreatePayments(data structures.PaymentsForm) (structures.PaymentsForm, error) {
	args := m.Called(data) // Artık doğru ve basit
	if args.Get(0) == nil {
		return structures.PaymentsForm{}, args.Error(1)
	}
	return args.Get(0).(structures.PaymentsForm), args.Error(1)
}

func (m *MockPaymentsService) CreatePaymentsMultiple(data []structures.PaymentsForm) ([]structures.PaymentsForm, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]structures.PaymentsForm), args.Error(1)
}

func (m *MockPaymentsService) UpdatePayments(id types.URID, data structures.PaymentsEdit) error {
	args := m.Called(id, data) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockPaymentsService) UpdatePaymentsMultiple(data []structures.PaymentsBatchUpdate) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockPaymentsService) DeletePayments(id types.URID) error {
	args := m.Called(id) // Artık doğru ve basit
	return args.Error(0)
}

func (m *MockPaymentsService) DeletePaymentsMultiple(identities []structures.PaymentsIdentity) error {
	args := m.Called(identities)
	return args.Error(0)
}

// --- Test Fonksiyonları ---

// TestCreatePayments_Success
func TestCreatePayments_Success(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_payments", controller.CreatePayments)

	inputData := structures.PaymentsForm{}
	expectedResult := inputData
	mockService.On("CreatePayments", inputData).Return(expectedResult, nil).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_payments", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var actualResult structures.PaymentsForm
	json.NewDecoder(resp.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult, actualResult)
	mockService.AssertExpectations(t)
}

// TestCreatePayments_ServiceError
func TestCreatePayments_ServiceError(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Post("/test_payments", controller.CreatePayments)

	inputData := structures.PaymentsForm{}
	mockService.On("CreatePayments", inputData).Return(structures.PaymentsForm{}, errors.New("service error")).Once()

	bodyBytes, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/test_payments", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

// --- Batch Delete Unit Tests ---
func TestDeletePaymentsMultiple_Success(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_payments/bulk", controller.DeletePaymentsMultiple)

	identities := []structures.PaymentsIdentity{
		{Id: types.URID("00000000-0000-0000-0000-000000000001")},
		{Id: types.URID("00000000-0000-0000-0000-000000000002")},
	}

	mockService.On("DeletePaymentsMultiple", identities).Return(nil).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_payments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "deleted successfully")
	mockService.AssertExpectations(t)
}

func TestDeletePaymentsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_payments/bulk", controller.DeletePaymentsMultiple)

	identities := []structures.PaymentsIdentity{{Id: types.URID("00000000-0000-0000-0000-000000000001")}}
	mockService.On("DeletePaymentsMultiple", identities).Return(errors.New("internal service failure")).Once()

	bodyBytes, _ := json.Marshal(identities)
	req := httptest.NewRequest("DELETE", "/test_payments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeletePaymentsMultiple_EmptyIdentities(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_payments/bulk", controller.DeletePaymentsMultiple)

	var emptyIdentities []structures.PaymentsIdentity
	bodyBytes, _ := json.Marshal(emptyIdentities)
	req := httptest.NewRequest("DELETE", "/test_payments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeletePaymentsMultiple", mock.Anything)
}

// --- Batch Update Unit Tests ---
func TestUpdatePaymentsMultiple_Success(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_payments/bulk", controller.UpdatePaymentsMultiple)

	batchUpdatePayload := []structures.PaymentsBatchUpdate{
		{
			PathParams: structures.PaymentsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")},
			Data:       structures.PaymentsEdit{Name: paymentsStringPtr("Updated Name 1")},
		},
		{
			PathParams: structures.PaymentsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000002")},
			Data:       structures.PaymentsEdit{Name: paymentsStringPtr("Updated Name 2")},
		},
	}

	mockService.On("UpdatePaymentsMultiple", batchUpdatePayload).Return(nil).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_payments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var responseMessage web.Fiber200Response
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.Contains(t, responseMessage.Message, "records processed for batch update")
	mockService.AssertExpectations(t)
}

func TestUpdatePaymentsMultiple_ServiceError(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_payments/bulk", controller.UpdatePaymentsMultiple)

	batchUpdatePayload := []structures.PaymentsBatchUpdate{
		{PathParams: structures.PaymentsIdentity{Id: types.URID("00000000-0000-0000-0000-000000000001")}, Data: structures.PaymentsEdit{Name: paymentsStringPtr("Fail Update")}},
	}
	mockService.On("UpdatePaymentsMultiple", batchUpdatePayload).Return(errors.New("service layer batch update failed")).Once()

	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	req := httptest.NewRequest("PUT", "/test_payments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdatePaymentsMultiple_EmptyPayload(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Put("/test_payments/bulk", controller.UpdatePaymentsMultiple)

	var emptyPayload []structures.PaymentsBatchUpdate
	bodyBytes, _ := json.Marshal(emptyPayload)
	req := httptest.NewRequest("PUT", "/test_payments/bulk", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "UpdatePaymentsMultiple", mock.Anything)
}

// stringPtr, string değerler için pointer döndüren bir yardımcı fonksiyondur.
// Testlerde PaymentsEdit içindeki pointer alanları doldurmak için kullanılır.
// Bu fonksiyonu test dosyasının kendisine veya paylaşılan bir test helper'ına ekleyebilirsiniz.
func paymentsStringPtr(s string) *string { return testhelper.StringPtr(s) }
func paymentsIntPtr(i int) *int          { return testhelper.IntPtr(i) } // Örnek, diğer tipler için de eklenebilir
func TestDeletePaymentsMultiple_InvalidBody(t *testing.T) {
	mockService := new(MockPaymentsService)
	controller := PaymentsController{Svc: mockService, EM: events.NewEventManager()}
	app := fiber.New()
	app.Delete("/test_payments/bulk", controller.DeletePaymentsMultiple)

	req := httptest.NewRequest("DELETE", "/test_payments/bulk", bytes.NewBufferString(`[{"id":1}, invalid_json]`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockService.AssertNotCalled(t, "DeletePaymentsMultiple", mock.Anything)
}

// Update (tekil) ve Delete (tekil) testleri TestCreate yukarıdaki desenle yapılabilir:
// Update için c.Svc.UpdatePayments(pk, editData) mock'lanır; Delete için c.Svc.DeletePayments(pk) mock'lanır.
// GetById: servis DB'ye doğrudan eriştiğinden entegrasyon testi gerektirir (bkz. TestGetPaymentsWithPagination_Success).

// GetWithPagination için test (VIEW'lar için de geçerli)
func TestGetPaymentsWithPagination_Success(t *testing.T) {
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
