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
	createdDashboardsID string

	// Batch Update için PK'lar
	createdDashboardsIDForBatchUpdateA string
	createdDashboardsIDForBatchUpdateB string
	createdDashboardsIDNotBatched      string

	// Batch Delete için PK'lar
	createdDashboardsIDForBatchDelete1    string
	createdDashboardsIDForBatchDelete2    string
	createdDashboardsIDToKeepForBatchTest string
)

// extractDashboardsPKPath, create yanıtından veya pagination'dan tüm PK değerlerini
// alarak "/" ile birleştirilmiş bir path segment döndürür.
func extractDashboardsPKPath(t *testing.T, routePath, createBody string, pkFields []string) string {
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

// buildDashboardsBatchPKMap, "/" ile birleştirilmiş PK path'inden batch işlem
// için ayrı PK alanları içeren map oluşturur.
func buildDashboardsBatchPKMap(pkPath string, pkFields []string) map[string]interface{} {
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
// Test_A_CreateDashboards: Yeni bir kaynak oluşturur.
func Test_A_CreateDashboards(t *testing.T) {
	pkFields := []string{"id"}
	createData := map[string]interface{}{"name": "API Test Create - Dashboards"}
	bodyBytes, _ := json.Marshal(createData)
	resp, bodyResp := makeRequest("POST", "/analytics/dashboards", string(bodyBytes))
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode, "API Create failed, body: %s", string(bodyResp))
	createdDashboardsID = extractDashboardsPKPath(t, "/analytics/dashboards", bodyResp, pkFields)
}

// Test_B_ReadDashboards: Oluşturulan kaynağı PK ile okur.
func Test_B_ReadDashboards(t *testing.T) {
	if !assert.NotEmpty(t, createdDashboardsID, "Create test must run first and set createdDashboardsID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	url := fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", createdDashboardsID)
	resp, bodyResp := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Read failed, body: %s", string(bodyResp))
	assert.Contains(t, string(bodyResp), "API Test Create - Dashboards")
}

// Test_C_UpdateDashboards: Kaynağı günceller.
func Test_C_UpdateDashboards(t *testing.T) {
	if !assert.NotEmpty(t, createdDashboardsID, "Create test must run first and set createdDashboardsID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	updateData := map[string]interface{}{"name": "API Test Update - Dashboards (Updated)"}
	bodyBytes, _ := json.Marshal(updateData)
	url := fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", createdDashboardsID)
	resp, bodyResp := makeRequest("PUT", url, string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Update failed, body: %s", string(bodyResp))

	respVerify, bodyVerify := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, respVerify.StatusCode)
	assert.Contains(t, string(bodyVerify), "API Test Update - Dashboards (Updated)")
}

// Test_D_ListWithPaginationDashboards: Kaynakları paginated olarak listeler.
func Test_D_ListWithPaginationDashboards(t *testing.T) {
	url := fmt.Sprintf("%s/pagination", "/analytics/dashboards")
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
func Test_E_SetupForBatchUpdateDashboards(t *testing.T) {
	pkFields := []string{"id"}

	createDataA := map[string]interface{}{"name": "BatchUpdate Item A - Original"}

	bodyA, _ := json.Marshal(createDataA)
	resp1, body1 := makeRequest("POST", "/analytics/dashboards", string(bodyA))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdDashboardsIDForBatchUpdateA = extractDashboardsPKPath(t, "/analytics/dashboards", body1, pkFields)

	createDataB := map[string]interface{}{"name": "BatchUpdate Item B - Original"}

	bodyB, _ := json.Marshal(createDataB)
	resp2, body2 := makeRequest("POST", "/analytics/dashboards", string(bodyB))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdDashboardsIDForBatchUpdateB = extractDashboardsPKPath(t, "/analytics/dashboards", body2, pkFields)

	createDataC := map[string]interface{}{"name": "Item Not In Batch Update"}

	bodyC, _ := json.Marshal(createDataC)
	resp3, body3 := makeRequest("POST", "/analytics/dashboards", string(bodyC))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdDashboardsIDNotBatched = extractDashboardsPKPath(t, "/analytics/dashboards", body3, pkFields)
}

func Test_F_BatchUpdateDashboards(t *testing.T) {
	if !assert.NotEmpty(t, createdDashboardsIDForBatchUpdateA, "SetupForBatchUpdate test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"id"}
	batchUpdatePayload := []map[string]interface{}{
		{"pathParams": buildDashboardsBatchPKMap(createdDashboardsIDForBatchUpdateA, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item A - Updated"}},
		{"pathParams": buildDashboardsBatchPKMap(createdDashboardsIDForBatchUpdateB, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item B - Updated"}},
	}
	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	resp, _ := makeRequest("PUT", "/analytics/dashboards/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchUpdate failed")

	_, bodyA := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", createdDashboardsIDForBatchUpdateA), "")
	assert.Contains(t, string(bodyA), "BatchUpdate Item A - Updated")

	_, bodyB := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", createdDashboardsIDForBatchUpdateB), "")
	assert.Contains(t, string(bodyB), "BatchUpdate Item B - Updated")

	_, bodyKeep := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", createdDashboardsIDNotBatched), "")
	assert.Contains(t, string(bodyKeep), "Item Not In Batch Update")
}

// --- Batch Delete API Tests ---
func Test_G_SetupForBatchDeleteDashboards(t *testing.T) {
	pkFields := []string{"id"}

	createData1 := map[string]interface{}{"name": "Batch Delete Item 1"}

	body1Bytes, _ := json.Marshal(createData1)
	resp1, body1 := makeRequest("POST", "/analytics/dashboards", string(body1Bytes))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdDashboardsIDForBatchDelete1 = extractDashboardsPKPath(t, "/analytics/dashboards", body1, pkFields)

	createData2 := map[string]interface{}{"name": "Batch Delete Item 2"}

	body2Bytes, _ := json.Marshal(createData2)
	resp2, body2 := makeRequest("POST", "/analytics/dashboards", string(body2Bytes))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdDashboardsIDForBatchDelete2 = extractDashboardsPKPath(t, "/analytics/dashboards", body2, pkFields)

	createData3 := map[string]interface{}{"name": "Item To Keep For Batch Test"}

	body3Bytes, _ := json.Marshal(createData3)
	resp3, body3 := makeRequest("POST", "/analytics/dashboards", string(body3Bytes))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdDashboardsIDToKeepForBatchTest = extractDashboardsPKPath(t, "/analytics/dashboards", body3, pkFields)
}

func Test_H_BatchDeleteDashboards(t *testing.T) {
	if !assert.NotEmpty(t, createdDashboardsIDForBatchDelete1, "SetupForBatchDelete test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"id"}
	identitiesToDelete := []map[string]interface{}{
		buildDashboardsBatchPKMap(createdDashboardsIDForBatchDelete1, pkFields),
		buildDashboardsBatchPKMap(createdDashboardsIDForBatchDelete2, pkFields),
	}
	bodyBytes, _ := json.Marshal(identitiesToDelete)
	resp, bodyStr := makeRequest("DELETE", "/analytics/dashboards/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchDelete failed")
	assert.Contains(t, string(bodyStr), "records deleted successfully")

	respCheck1, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", createdDashboardsIDForBatchDelete1), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck1.StatusCode)

	respCheck2, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", createdDashboardsIDForBatchDelete2), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck2.StatusCode)

	respCheckKeep, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", createdDashboardsIDToKeepForBatchTest), "")
	assert.Equal(t, fiber.StatusOK, respCheckKeep.StatusCode)
}

// Test_Z_CleanupDashboards: Kalan test verilerini temizler.
func Test_Z_CleanupDashboards(t *testing.T) {
	idList := []string{createdDashboardsID, createdDashboardsIDNotBatched, createdDashboardsIDToKeepForBatchTest}
	for _, pkPath := range idList {
		if pkPath != "" {
			url := fmt.Sprintf("%s/with-id/%s", "/analytics/dashboards", pkPath)
			makeRequest("DELETE", url, "")
		}
	}
	t.Log("Cleanup test ran.")
}
