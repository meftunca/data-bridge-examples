package tests

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	paginationRuntime "backend-generator/apiv2/pagination"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Suppress unused import warnings
var _ = strings.Join
var _ = time.Now

// Testler arasında oluşturulan kaynakların PK yollarını taşımak için global değişkenler.
// Composite PK'larda "/" ile birleştirilmiş path segment saklanır (ör: "1/test-key").
var (
	createdOrderStatusHistoryID string

	// Batch Update için PK'lar
	createdOrderStatusHistoryIDForBatchUpdateA string
	createdOrderStatusHistoryIDForBatchUpdateB string
	createdOrderStatusHistoryIDNotBatched      string

	// Batch Delete için PK'lar
	createdOrderStatusHistoryIDForBatchDelete1    string
	createdOrderStatusHistoryIDForBatchDelete2    string
	createdOrderStatusHistoryIDToKeepForBatchTest string
)

// extractOrderStatusHistoryPKPath, create yanıtından veya pagination'dan tüm PK değerlerini
// alarak "/" ile birleştirilmiş bir path segment döndürür.
func extractOrderStatusHistoryPKPath(t *testing.T, routePath, createBody string, pkFields []string) string {
	var direct map[string]interface{}
	if err := json.Unmarshal([]byte(createBody), &direct); err == nil {
		parts := make([]string, 0, len(pkFields))
		allFound := true
		for _, name := range pkFields {
			if val, ok := direct[name]; ok {
				s := fmt.Sprintf("%v", val)
				if s != "" && s != "<nil>" {
					parts = append(parts, s)
				} else {
					allFound = false
					break
				}
			} else {
				allFound = false
				break
			}
		}
		if allFound && len(parts) == len(pkFields) {
			return strings.Join(parts, "/")
		}
	}

	resp, bodyResp := makeRequest("GET", fmt.Sprintf("%s/pagination?sort=-%s", routePath, pkFields[0]), "")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "failed to retrieve paginated data for PK extraction, body: %s", bodyResp)

	var pageResponse paginationRuntime.Page[map[string]interface{}]
	err := json.Unmarshal([]byte(bodyResp), &pageResponse)
	assert.NoError(t, err, "failed to unmarshal paginated response for PK extraction")
	if !assert.NotEmpty(t, pageResponse.Items, "paginated response should include at least one item for PK extraction") {
		return ""
	}

	item := pageResponse.Items[0]
	parts := make([]string, 0, len(pkFields))
	for _, name := range pkFields {
		val, ok := item[name]
		assert.True(t, ok, "primary key '%s' not found in paginated response", name)
		parts = append(parts, fmt.Sprintf("%v", val))
	}
	return strings.Join(parts, "/")
}

// buildOrderStatusHistoryBatchPKMap, "/" ile birleştirilmiş PK path'inden batch işlem
// için ayrı PK alanları içeren map oluşturur.
func buildOrderStatusHistoryBatchPKMap(pkPath string, pkFields []string) map[string]interface{} {
	result := make(map[string]interface{})
	parts := strings.Split(pkPath, "/")
	for i, name := range pkFields {
		if i < len(parts) {
			result[name] = parts[i]
		}
	}
	return result
}

// VIEW'lar için Create, Update, Delete testleri olmaz
// Test_A_CreateOrderStatusHistory: Yeni bir kaynak oluşturur.
func Test_A_CreateOrderStatusHistory(t *testing.T) {
	pkFields := []string{"id"}
	createData := map[string]interface{}{"name": "API Test Create - OrderStatusHistory"}
	bodyBytes, _ := json.Marshal(createData)
	resp, bodyResp := makeRequest("POST", "/orders/order-status-history", string(bodyBytes))
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode, "API Create failed, body: %s", string(bodyResp))
	createdOrderStatusHistoryID = extractOrderStatusHistoryPKPath(t, "/orders/order-status-history", bodyResp, pkFields)
}

// Test_B_ReadOrderStatusHistory: Oluşturulan kaynağı PK ile okur.
func Test_B_ReadOrderStatusHistory(t *testing.T) {
	if !assert.NotEmpty(t, createdOrderStatusHistoryID, "Create test must run first and set createdOrderStatusHistoryID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	url := fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", createdOrderStatusHistoryID)
	resp, bodyResp := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Read failed, body: %s", string(bodyResp))
	assert.Contains(t, string(bodyResp), "API Test Create - OrderStatusHistory")
}

// Test_C_UpdateOrderStatusHistory: Kaynağı günceller.
func Test_C_UpdateOrderStatusHistory(t *testing.T) {
	if !assert.NotEmpty(t, createdOrderStatusHistoryID, "Create test must run first and set createdOrderStatusHistoryID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	updateData := map[string]interface{}{"name": "API Test Update - OrderStatusHistory (Updated)"}
	bodyBytes, _ := json.Marshal(updateData)
	url := fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", createdOrderStatusHistoryID)
	resp, bodyResp := makeRequest("PUT", url, string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Update failed, body: %s", string(bodyResp))

	respVerify, bodyVerify := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, respVerify.StatusCode)
	assert.Contains(t, string(bodyVerify), "API Test Update - OrderStatusHistory (Updated)")
}

// Test_D_ListWithPaginationOrderStatusHistory: Kaynakları paginated olarak listeler.
func Test_D_ListWithPaginationOrderStatusHistory(t *testing.T) {
	url := fmt.Sprintf("%s/pagination", "/orders/order-status-history")
	resp, bodyResp := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API ListWithPagination failed, body: %s", string(bodyResp))

	var pageResponse paginationRuntime.Page[map[string]interface{}]
	err := json.Unmarshal([]byte(bodyResp), &pageResponse)
	assert.NoError(t, err, "Failed to unmarshal paginated response")
	assert.GreaterOrEqual(t, pageResponse.Total, int64(0))

	if pageResponse.Total > 0 {
		assert.NotEmpty(t, pageResponse.Items, "Paginated items should not be empty when total > 0")
	}
}

// --- Batch Update API Tests ---
func Test_E_SetupForBatchUpdateOrderStatusHistory(t *testing.T) {
	pkFields := []string{"id"}

	createDataA := map[string]interface{}{"name": "BatchUpdate Item A - Original"}

	bodyA, _ := json.Marshal(createDataA)
	resp1, body1 := makeRequest("POST", "/orders/order-status-history", string(bodyA))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdOrderStatusHistoryIDForBatchUpdateA = extractOrderStatusHistoryPKPath(t, "/orders/order-status-history", body1, pkFields)

	createDataB := map[string]interface{}{"name": "BatchUpdate Item B - Original"}

	bodyB, _ := json.Marshal(createDataB)
	resp2, body2 := makeRequest("POST", "/orders/order-status-history", string(bodyB))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdOrderStatusHistoryIDForBatchUpdateB = extractOrderStatusHistoryPKPath(t, "/orders/order-status-history", body2, pkFields)

	createDataC := map[string]interface{}{"name": "Item Not In Batch Update"}

	bodyC, _ := json.Marshal(createDataC)
	resp3, body3 := makeRequest("POST", "/orders/order-status-history", string(bodyC))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdOrderStatusHistoryIDNotBatched = extractOrderStatusHistoryPKPath(t, "/orders/order-status-history", body3, pkFields)
}

func Test_F_BatchUpdateOrderStatusHistory(t *testing.T) {
	if !assert.NotEmpty(t, createdOrderStatusHistoryIDForBatchUpdateA, "SetupForBatchUpdate test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"id"}
	batchUpdatePayload := []map[string]interface{}{
		{"pathParams": buildOrderStatusHistoryBatchPKMap(createdOrderStatusHistoryIDForBatchUpdateA, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item A - Updated"}},
		{"pathParams": buildOrderStatusHistoryBatchPKMap(createdOrderStatusHistoryIDForBatchUpdateB, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item B - Updated"}},
	}
	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	resp, _ := makeRequest("PUT", "/orders/order-status-history/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchUpdate failed")

	_, bodyA := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", createdOrderStatusHistoryIDForBatchUpdateA), "")
	assert.Contains(t, string(bodyA), "BatchUpdate Item A - Updated")

	_, bodyB := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", createdOrderStatusHistoryIDForBatchUpdateB), "")
	assert.Contains(t, string(bodyB), "BatchUpdate Item B - Updated")

	_, bodyKeep := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", createdOrderStatusHistoryIDNotBatched), "")
	assert.Contains(t, string(bodyKeep), "Item Not In Batch Update")
}

// --- Batch Delete API Tests ---
func Test_G_SetupForBatchDeleteOrderStatusHistory(t *testing.T) {
	pkFields := []string{"id"}

	createData1 := map[string]interface{}{"name": "Batch Delete Item 1"}

	body1Bytes, _ := json.Marshal(createData1)
	resp1, body1 := makeRequest("POST", "/orders/order-status-history", string(body1Bytes))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdOrderStatusHistoryIDForBatchDelete1 = extractOrderStatusHistoryPKPath(t, "/orders/order-status-history", body1, pkFields)

	createData2 := map[string]interface{}{"name": "Batch Delete Item 2"}

	body2Bytes, _ := json.Marshal(createData2)
	resp2, body2 := makeRequest("POST", "/orders/order-status-history", string(body2Bytes))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdOrderStatusHistoryIDForBatchDelete2 = extractOrderStatusHistoryPKPath(t, "/orders/order-status-history", body2, pkFields)

	createData3 := map[string]interface{}{"name": "Item To Keep For Batch Test"}

	body3Bytes, _ := json.Marshal(createData3)
	resp3, body3 := makeRequest("POST", "/orders/order-status-history", string(body3Bytes))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdOrderStatusHistoryIDToKeepForBatchTest = extractOrderStatusHistoryPKPath(t, "/orders/order-status-history", body3, pkFields)
}

func Test_H_BatchDeleteOrderStatusHistory(t *testing.T) {
	if !assert.NotEmpty(t, createdOrderStatusHistoryIDForBatchDelete1, "SetupForBatchDelete test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"id"}
	identitiesToDelete := []map[string]interface{}{
		buildOrderStatusHistoryBatchPKMap(createdOrderStatusHistoryIDForBatchDelete1, pkFields),
		buildOrderStatusHistoryBatchPKMap(createdOrderStatusHistoryIDForBatchDelete2, pkFields),
	}
	bodyBytes, _ := json.Marshal(identitiesToDelete)
	resp, bodyStr := makeRequest("DELETE", "/orders/order-status-history/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchDelete failed")
	assert.Contains(t, string(bodyStr), "records deleted successfully")

	respCheck1, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", createdOrderStatusHistoryIDForBatchDelete1), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck1.StatusCode)

	respCheck2, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", createdOrderStatusHistoryIDForBatchDelete2), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck2.StatusCode)

	respCheckKeep, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", createdOrderStatusHistoryIDToKeepForBatchTest), "")
	assert.Equal(t, fiber.StatusOK, respCheckKeep.StatusCode)
}

// Test_Z_CleanupOrderStatusHistory: Kalan test verilerini temizler.
func Test_Z_CleanupOrderStatusHistory(t *testing.T) {
	idList := []string{createdOrderStatusHistoryID, createdOrderStatusHistoryIDNotBatched, createdOrderStatusHistoryIDToKeepForBatchTest}
	for _, pkPath := range idList {
		if pkPath != "" {
			url := fmt.Sprintf("%s/with-id/%s", "/orders/order-status-history", pkPath)
			makeRequest("DELETE", url, "")
		}
	}
	t.Log("Cleanup test ran.")
}
