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
	createdRolesID string

	// Batch Update için PK'lar
	createdRolesIDForBatchUpdateA string
	createdRolesIDForBatchUpdateB string
	createdRolesIDNotBatched      string

	// Batch Delete için PK'lar
	createdRolesIDForBatchDelete1    string
	createdRolesIDForBatchDelete2    string
	createdRolesIDToKeepForBatchTest string
)

// extractRolesPKPath, create yanıtından veya pagination'dan tüm PK değerlerini
// alarak "/" ile birleştirilmiş bir path segment döndürür.
func extractRolesPKPath(t *testing.T, routePath, createBody string, pkFields []string) string {
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

// buildRolesBatchPKMap, "/" ile birleştirilmiş PK path'inden batch işlem
// için ayrı PK alanları içeren map oluşturur.
func buildRolesBatchPKMap(pkPath string, pkFields []string) map[string]interface{} {
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
// Test_A_CreateRoles: Yeni bir kaynak oluşturur.
func Test_A_CreateRoles(t *testing.T) {
	pkFields := []string{"id"}
	createData := map[string]interface{}{"name": "API Test Create - Roles"}
	bodyBytes, _ := json.Marshal(createData)
	resp, bodyResp := makeRequest("POST", "/iam/roles", string(bodyBytes))
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode, "API Create failed, body: %s", string(bodyResp))
	createdRolesID = extractRolesPKPath(t, "/iam/roles", bodyResp, pkFields)
}

// Test_B_ReadRoles: Oluşturulan kaynağı PK ile okur.
func Test_B_ReadRoles(t *testing.T) {
	if !assert.NotEmpty(t, createdRolesID, "Create test must run first and set createdRolesID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	url := fmt.Sprintf("%s/with-id/%s", "/iam/roles", createdRolesID)
	resp, bodyResp := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Read failed, body: %s", string(bodyResp))
	assert.Contains(t, string(bodyResp), "API Test Create - Roles")
}

// Test_C_UpdateRoles: Kaynağı günceller.
func Test_C_UpdateRoles(t *testing.T) {
	if !assert.NotEmpty(t, createdRolesID, "Create test must run first and set createdRolesID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	updateData := map[string]interface{}{"name": "API Test Update - Roles (Updated)"}
	bodyBytes, _ := json.Marshal(updateData)
	url := fmt.Sprintf("%s/with-id/%s", "/iam/roles", createdRolesID)
	resp, bodyResp := makeRequest("PUT", url, string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Update failed, body: %s", string(bodyResp))

	respVerify, bodyVerify := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, respVerify.StatusCode)
	assert.Contains(t, string(bodyVerify), "API Test Update - Roles (Updated)")
}

// Test_D_ListWithPaginationRoles: Kaynakları paginated olarak listeler.
func Test_D_ListWithPaginationRoles(t *testing.T) {
	url := fmt.Sprintf("%s/pagination", "/iam/roles")
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
func Test_E_SetupForBatchUpdateRoles(t *testing.T) {
	pkFields := []string{"id"}

	createDataA := map[string]interface{}{"name": "BatchUpdate Item A - Original"}

	bodyA, _ := json.Marshal(createDataA)
	resp1, body1 := makeRequest("POST", "/iam/roles", string(bodyA))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdRolesIDForBatchUpdateA = extractRolesPKPath(t, "/iam/roles", body1, pkFields)

	createDataB := map[string]interface{}{"name": "BatchUpdate Item B - Original"}

	bodyB, _ := json.Marshal(createDataB)
	resp2, body2 := makeRequest("POST", "/iam/roles", string(bodyB))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdRolesIDForBatchUpdateB = extractRolesPKPath(t, "/iam/roles", body2, pkFields)

	createDataC := map[string]interface{}{"name": "Item Not In Batch Update"}

	bodyC, _ := json.Marshal(createDataC)
	resp3, body3 := makeRequest("POST", "/iam/roles", string(bodyC))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdRolesIDNotBatched = extractRolesPKPath(t, "/iam/roles", body3, pkFields)
}

func Test_F_BatchUpdateRoles(t *testing.T) {
	if !assert.NotEmpty(t, createdRolesIDForBatchUpdateA, "SetupForBatchUpdate test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"id"}
	batchUpdatePayload := []map[string]interface{}{
		{"pathParams": buildRolesBatchPKMap(createdRolesIDForBatchUpdateA, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item A - Updated"}},
		{"pathParams": buildRolesBatchPKMap(createdRolesIDForBatchUpdateB, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item B - Updated"}},
	}
	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	resp, _ := makeRequest("PUT", "/iam/roles/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchUpdate failed")

	_, bodyA := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/iam/roles", createdRolesIDForBatchUpdateA), "")
	assert.Contains(t, string(bodyA), "BatchUpdate Item A - Updated")

	_, bodyB := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/iam/roles", createdRolesIDForBatchUpdateB), "")
	assert.Contains(t, string(bodyB), "BatchUpdate Item B - Updated")

	_, bodyKeep := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/iam/roles", createdRolesIDNotBatched), "")
	assert.Contains(t, string(bodyKeep), "Item Not In Batch Update")
}

// --- Batch Delete API Tests ---
func Test_G_SetupForBatchDeleteRoles(t *testing.T) {
	pkFields := []string{"id"}

	createData1 := map[string]interface{}{"name": "Batch Delete Item 1"}

	body1Bytes, _ := json.Marshal(createData1)
	resp1, body1 := makeRequest("POST", "/iam/roles", string(body1Bytes))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdRolesIDForBatchDelete1 = extractRolesPKPath(t, "/iam/roles", body1, pkFields)

	createData2 := map[string]interface{}{"name": "Batch Delete Item 2"}

	body2Bytes, _ := json.Marshal(createData2)
	resp2, body2 := makeRequest("POST", "/iam/roles", string(body2Bytes))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdRolesIDForBatchDelete2 = extractRolesPKPath(t, "/iam/roles", body2, pkFields)

	createData3 := map[string]interface{}{"name": "Item To Keep For Batch Test"}

	body3Bytes, _ := json.Marshal(createData3)
	resp3, body3 := makeRequest("POST", "/iam/roles", string(body3Bytes))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdRolesIDToKeepForBatchTest = extractRolesPKPath(t, "/iam/roles", body3, pkFields)
}

func Test_H_BatchDeleteRoles(t *testing.T) {
	if !assert.NotEmpty(t, createdRolesIDForBatchDelete1, "SetupForBatchDelete test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"id"}
	identitiesToDelete := []map[string]interface{}{
		buildRolesBatchPKMap(createdRolesIDForBatchDelete1, pkFields),
		buildRolesBatchPKMap(createdRolesIDForBatchDelete2, pkFields),
	}
	bodyBytes, _ := json.Marshal(identitiesToDelete)
	resp, bodyStr := makeRequest("DELETE", "/iam/roles/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchDelete failed")
	assert.Contains(t, string(bodyStr), "records deleted successfully")

	respCheck1, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/iam/roles", createdRolesIDForBatchDelete1), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck1.StatusCode)

	respCheck2, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/iam/roles", createdRolesIDForBatchDelete2), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck2.StatusCode)

	respCheckKeep, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/iam/roles", createdRolesIDToKeepForBatchTest), "")
	assert.Equal(t, fiber.StatusOK, respCheckKeep.StatusCode)
}

// Test_Z_CleanupRoles: Kalan test verilerini temizler.
func Test_Z_CleanupRoles(t *testing.T) {
	idList := []string{createdRolesID, createdRolesIDNotBatched, createdRolesIDToKeepForBatchTest}
	for _, pkPath := range idList {
		if pkPath != "" {
			url := fmt.Sprintf("%s/with-id/%s", "/iam/roles", pkPath)
			makeRequest("DELETE", url, "")
		}
	}
	t.Log("Cleanup test ran.")
}
