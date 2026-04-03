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
	createdShipmentTrackingID string

	// Batch Update için PK'lar
	createdShipmentTrackingIDForBatchUpdateA string
	createdShipmentTrackingIDForBatchUpdateB string
	createdShipmentTrackingIDNotBatched      string

	// Batch Delete için PK'lar
	createdShipmentTrackingIDForBatchDelete1    string
	createdShipmentTrackingIDForBatchDelete2    string
	createdShipmentTrackingIDToKeepForBatchTest string
)

// extractShipmentTrackingPKPath, create yanıtından veya pagination'dan tüm PK değerlerini
// alarak "/" ile birleştirilmiş bir path segment döndürür.
func extractShipmentTrackingPKPath(t *testing.T, routePath, createBody string, pkFields []string) string {
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

// buildShipmentTrackingBatchPKMap, "/" ile birleştirilmiş PK path'inden batch işlem
// için ayrı PK alanları içeren map oluşturur.
func buildShipmentTrackingBatchPKMap(pkPath string, pkFields []string) map[string]interface{} {
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
// Test_A_CreateShipmentTracking: Yeni bir kaynak oluşturur.
func Test_A_CreateShipmentTracking(t *testing.T) {
	pkFields := []string{"id"}
	createData := map[string]interface{}{"name": "API Test Create - ShipmentTracking"}
	bodyBytes, _ := json.Marshal(createData)
	resp, bodyResp := makeRequest("POST", "/logistics/shipment-tracking", string(bodyBytes))
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode, "API Create failed, body: %s", string(bodyResp))
	createdShipmentTrackingID = extractShipmentTrackingPKPath(t, "/logistics/shipment-tracking", bodyResp, pkFields)
}

// Test_B_ReadShipmentTracking: Oluşturulan kaynağı PK ile okur.
func Test_B_ReadShipmentTracking(t *testing.T) {
	if !assert.NotEmpty(t, createdShipmentTrackingID, "Create test must run first and set createdShipmentTrackingID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	url := fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", createdShipmentTrackingID)
	resp, bodyResp := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Read failed, body: %s", string(bodyResp))
	assert.Contains(t, string(bodyResp), "API Test Create - ShipmentTracking")
}

// Test_C_UpdateShipmentTracking: Kaynağı günceller.
func Test_C_UpdateShipmentTracking(t *testing.T) {
	if !assert.NotEmpty(t, createdShipmentTrackingID, "Create test must run first and set createdShipmentTrackingID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	updateData := map[string]interface{}{"name": "API Test Update - ShipmentTracking (Updated)"}
	bodyBytes, _ := json.Marshal(updateData)
	url := fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", createdShipmentTrackingID)
	resp, bodyResp := makeRequest("PUT", url, string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Update failed, body: %s", string(bodyResp))

	respVerify, bodyVerify := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, respVerify.StatusCode)
	assert.Contains(t, string(bodyVerify), "API Test Update - ShipmentTracking (Updated)")
}

// Test_D_ListWithPaginationShipmentTracking: Kaynakları paginated olarak listeler.
func Test_D_ListWithPaginationShipmentTracking(t *testing.T) {
	url := fmt.Sprintf("%s/pagination", "/logistics/shipment-tracking")
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
func Test_E_SetupForBatchUpdateShipmentTracking(t *testing.T) {
	pkFields := []string{"id"}

	createDataA := map[string]interface{}{"name": "BatchUpdate Item A - Original"}

	bodyA, _ := json.Marshal(createDataA)
	resp1, body1 := makeRequest("POST", "/logistics/shipment-tracking", string(bodyA))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdShipmentTrackingIDForBatchUpdateA = extractShipmentTrackingPKPath(t, "/logistics/shipment-tracking", body1, pkFields)

	createDataB := map[string]interface{}{"name": "BatchUpdate Item B - Original"}

	bodyB, _ := json.Marshal(createDataB)
	resp2, body2 := makeRequest("POST", "/logistics/shipment-tracking", string(bodyB))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdShipmentTrackingIDForBatchUpdateB = extractShipmentTrackingPKPath(t, "/logistics/shipment-tracking", body2, pkFields)

	createDataC := map[string]interface{}{"name": "Item Not In Batch Update"}

	bodyC, _ := json.Marshal(createDataC)
	resp3, body3 := makeRequest("POST", "/logistics/shipment-tracking", string(bodyC))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdShipmentTrackingIDNotBatched = extractShipmentTrackingPKPath(t, "/logistics/shipment-tracking", body3, pkFields)
}

func Test_F_BatchUpdateShipmentTracking(t *testing.T) {
	if !assert.NotEmpty(t, createdShipmentTrackingIDForBatchUpdateA, "SetupForBatchUpdate test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"id"}
	batchUpdatePayload := []map[string]interface{}{
		{"pathParams": buildShipmentTrackingBatchPKMap(createdShipmentTrackingIDForBatchUpdateA, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item A - Updated"}},
		{"pathParams": buildShipmentTrackingBatchPKMap(createdShipmentTrackingIDForBatchUpdateB, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item B - Updated"}},
	}
	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	resp, _ := makeRequest("PUT", "/logistics/shipment-tracking/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchUpdate failed")

	_, bodyA := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", createdShipmentTrackingIDForBatchUpdateA), "")
	assert.Contains(t, string(bodyA), "BatchUpdate Item A - Updated")

	_, bodyB := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", createdShipmentTrackingIDForBatchUpdateB), "")
	assert.Contains(t, string(bodyB), "BatchUpdate Item B - Updated")

	_, bodyKeep := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", createdShipmentTrackingIDNotBatched), "")
	assert.Contains(t, string(bodyKeep), "Item Not In Batch Update")
}

// --- Batch Delete API Tests ---
func Test_G_SetupForBatchDeleteShipmentTracking(t *testing.T) {
	pkFields := []string{"id"}

	createData1 := map[string]interface{}{"name": "Batch Delete Item 1"}

	body1Bytes, _ := json.Marshal(createData1)
	resp1, body1 := makeRequest("POST", "/logistics/shipment-tracking", string(body1Bytes))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdShipmentTrackingIDForBatchDelete1 = extractShipmentTrackingPKPath(t, "/logistics/shipment-tracking", body1, pkFields)

	createData2 := map[string]interface{}{"name": "Batch Delete Item 2"}

	body2Bytes, _ := json.Marshal(createData2)
	resp2, body2 := makeRequest("POST", "/logistics/shipment-tracking", string(body2Bytes))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdShipmentTrackingIDForBatchDelete2 = extractShipmentTrackingPKPath(t, "/logistics/shipment-tracking", body2, pkFields)

	createData3 := map[string]interface{}{"name": "Item To Keep For Batch Test"}

	body3Bytes, _ := json.Marshal(createData3)
	resp3, body3 := makeRequest("POST", "/logistics/shipment-tracking", string(body3Bytes))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdShipmentTrackingIDToKeepForBatchTest = extractShipmentTrackingPKPath(t, "/logistics/shipment-tracking", body3, pkFields)
}

func Test_H_BatchDeleteShipmentTracking(t *testing.T) {
	if !assert.NotEmpty(t, createdShipmentTrackingIDForBatchDelete1, "SetupForBatchDelete test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"id"}
	identitiesToDelete := []map[string]interface{}{
		buildShipmentTrackingBatchPKMap(createdShipmentTrackingIDForBatchDelete1, pkFields),
		buildShipmentTrackingBatchPKMap(createdShipmentTrackingIDForBatchDelete2, pkFields),
	}
	bodyBytes, _ := json.Marshal(identitiesToDelete)
	resp, bodyStr := makeRequest("DELETE", "/logistics/shipment-tracking/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchDelete failed")
	assert.Contains(t, string(bodyStr), "records deleted successfully")

	respCheck1, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", createdShipmentTrackingIDForBatchDelete1), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck1.StatusCode)

	respCheck2, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", createdShipmentTrackingIDForBatchDelete2), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck2.StatusCode)

	respCheckKeep, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", createdShipmentTrackingIDToKeepForBatchTest), "")
	assert.Equal(t, fiber.StatusOK, respCheckKeep.StatusCode)
}

// Test_Z_CleanupShipmentTracking: Kalan test verilerini temizler.
func Test_Z_CleanupShipmentTracking(t *testing.T) {
	idList := []string{createdShipmentTrackingID, createdShipmentTrackingIDNotBatched, createdShipmentTrackingIDToKeepForBatchTest}
	for _, pkPath := range idList {
		if pkPath != "" {
			url := fmt.Sprintf("%s/with-id/%s", "/logistics/shipment-tracking", pkPath)
			makeRequest("DELETE", url, "")
		}
	}
	t.Log("Cleanup test ran.")
}
