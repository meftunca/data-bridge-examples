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
	createdProductTagsID string

	// Batch Update için PK'lar
	createdProductTagsIDForBatchUpdateA string
	createdProductTagsIDForBatchUpdateB string
	createdProductTagsIDNotBatched      string

	// Batch Delete için PK'lar
	createdProductTagsIDForBatchDelete1    string
	createdProductTagsIDForBatchDelete2    string
	createdProductTagsIDToKeepForBatchTest string
)

// extractProductTagsPKPath, create yanıtından veya pagination'dan tüm PK değerlerini
// alarak "/" ile birleştirilmiş bir path segment döndürür.
func extractProductTagsPKPath(t *testing.T, routePath, createBody string, pkFields []string) string {
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

// buildProductTagsBatchPKMap, "/" ile birleştirilmiş PK path'inden batch işlem
// için ayrı PK alanları içeren map oluşturur.
func buildProductTagsBatchPKMap(pkPath string, pkFields []string) map[string]interface{} {
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
// Test_A_CreateProductTags: Yeni bir kaynak oluşturur.
func Test_A_CreateProductTags(t *testing.T) {
	pkFields := []string{"productId", "tagId"}
	createData := map[string]interface{}{"name": "API Test Create - ProductTags", "productId": 1, "tagId": 1}
	bodyBytes, _ := json.Marshal(createData)
	resp, bodyResp := makeRequest("POST", "/catalog/product-tags", string(bodyBytes))
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode, "API Create failed, body: %s", string(bodyResp))
	createdProductTagsID = extractProductTagsPKPath(t, "/catalog/product-tags", bodyResp, pkFields)
}

// Test_B_ReadProductTags: Oluşturulan kaynağı PK ile okur.
func Test_B_ReadProductTags(t *testing.T) {
	if !assert.NotEmpty(t, createdProductTagsID, "Create test must run first and set createdProductTagsID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	url := fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", createdProductTagsID)
	resp, bodyResp := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Read failed, body: %s", string(bodyResp))
	assert.Contains(t, string(bodyResp), "API Test Create - ProductTags")
}

// Test_C_UpdateProductTags: Kaynağı günceller.
func Test_C_UpdateProductTags(t *testing.T) {
	if !assert.NotEmpty(t, createdProductTagsID, "Create test must run first and set createdProductTagsID") {
		t.Skip("Skipping test, dependency failed.")
		return
	}
	updateData := map[string]interface{}{"name": "API Test Update - ProductTags (Updated)"}
	bodyBytes, _ := json.Marshal(updateData)
	url := fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", createdProductTagsID)
	resp, bodyResp := makeRequest("PUT", url, string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API Update failed, body: %s", string(bodyResp))

	respVerify, bodyVerify := makeRequest("GET", url, "")
	assert.Equal(t, fiber.StatusOK, respVerify.StatusCode)
	assert.Contains(t, string(bodyVerify), "API Test Update - ProductTags (Updated)")
}

// Test_D_ListWithPaginationProductTags: Kaynakları paginated olarak listeler.
func Test_D_ListWithPaginationProductTags(t *testing.T) {
	url := fmt.Sprintf("%s/pagination", "/catalog/product-tags")
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
func Test_E_SetupForBatchUpdateProductTags(t *testing.T) {
	pkFields := []string{"productId", "tagId"}

	createDataA := map[string]interface{}{"name": "BatchUpdate Item A - Original", "productId": 1, "tagId": 1}

	bodyA, _ := json.Marshal(createDataA)
	resp1, body1 := makeRequest("POST", "/catalog/product-tags", string(bodyA))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdProductTagsIDForBatchUpdateA = extractProductTagsPKPath(t, "/catalog/product-tags", body1, pkFields)

	createDataB := map[string]interface{}{"name": "BatchUpdate Item B - Original", "productId": 1, "tagId": 1}

	bodyB, _ := json.Marshal(createDataB)
	resp2, body2 := makeRequest("POST", "/catalog/product-tags", string(bodyB))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdProductTagsIDForBatchUpdateB = extractProductTagsPKPath(t, "/catalog/product-tags", body2, pkFields)

	createDataC := map[string]interface{}{"name": "Item Not In Batch Update", "productId": 1, "tagId": 1}

	bodyC, _ := json.Marshal(createDataC)
	resp3, body3 := makeRequest("POST", "/catalog/product-tags", string(bodyC))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdProductTagsIDNotBatched = extractProductTagsPKPath(t, "/catalog/product-tags", body3, pkFields)
}

func Test_F_BatchUpdateProductTags(t *testing.T) {
	if !assert.NotEmpty(t, createdProductTagsIDForBatchUpdateA, "SetupForBatchUpdate test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"productId", "tagId"}
	batchUpdatePayload := []map[string]interface{}{
		{"pathParams": buildProductTagsBatchPKMap(createdProductTagsIDForBatchUpdateA, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item A - Updated"}},
		{"pathParams": buildProductTagsBatchPKMap(createdProductTagsIDForBatchUpdateB, pkFields), "data": map[string]interface{}{"name": "BatchUpdate Item B - Updated"}},
	}
	bodyBytes, _ := json.Marshal(batchUpdatePayload)
	resp, _ := makeRequest("PUT", "/catalog/product-tags/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchUpdate failed")

	_, bodyA := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", createdProductTagsIDForBatchUpdateA), "")
	assert.Contains(t, string(bodyA), "BatchUpdate Item A - Updated")

	_, bodyB := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", createdProductTagsIDForBatchUpdateB), "")
	assert.Contains(t, string(bodyB), "BatchUpdate Item B - Updated")

	_, bodyKeep := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", createdProductTagsIDNotBatched), "")
	assert.Contains(t, string(bodyKeep), "Item Not In Batch Update")
}

// --- Batch Delete API Tests ---
func Test_G_SetupForBatchDeleteProductTags(t *testing.T) {
	pkFields := []string{"productId", "tagId"}

	createData1 := map[string]interface{}{"name": "Batch Delete Item 1", "productId": 1, "tagId": 1}

	body1Bytes, _ := json.Marshal(createData1)
	resp1, body1 := makeRequest("POST", "/catalog/product-tags", string(body1Bytes))
	assert.Equal(t, fiber.StatusCreated, resp1.StatusCode)
	createdProductTagsIDForBatchDelete1 = extractProductTagsPKPath(t, "/catalog/product-tags", body1, pkFields)

	createData2 := map[string]interface{}{"name": "Batch Delete Item 2", "productId": 1, "tagId": 1}

	body2Bytes, _ := json.Marshal(createData2)
	resp2, body2 := makeRequest("POST", "/catalog/product-tags", string(body2Bytes))
	assert.Equal(t, fiber.StatusCreated, resp2.StatusCode)
	createdProductTagsIDForBatchDelete2 = extractProductTagsPKPath(t, "/catalog/product-tags", body2, pkFields)

	createData3 := map[string]interface{}{"name": "Item To Keep For Batch Test", "productId": 1, "tagId": 1}

	body3Bytes, _ := json.Marshal(createData3)
	resp3, body3 := makeRequest("POST", "/catalog/product-tags", string(body3Bytes))
	assert.Equal(t, fiber.StatusCreated, resp3.StatusCode)
	createdProductTagsIDToKeepForBatchTest = extractProductTagsPKPath(t, "/catalog/product-tags", body3, pkFields)
}

func Test_H_BatchDeleteProductTags(t *testing.T) {
	if !assert.NotEmpty(t, createdProductTagsIDForBatchDelete1, "SetupForBatchDelete test must run first") {
		t.Skip("Skipping test, dependency failed.")
	}
	pkFields := []string{"productId", "tagId"}
	identitiesToDelete := []map[string]interface{}{
		buildProductTagsBatchPKMap(createdProductTagsIDForBatchDelete1, pkFields),
		buildProductTagsBatchPKMap(createdProductTagsIDForBatchDelete2, pkFields),
	}
	bodyBytes, _ := json.Marshal(identitiesToDelete)
	resp, bodyStr := makeRequest("DELETE", "/catalog/product-tags/bulk", string(bodyBytes))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "API BatchDelete failed")
	assert.Contains(t, string(bodyStr), "records deleted successfully")

	respCheck1, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", createdProductTagsIDForBatchDelete1), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck1.StatusCode)

	respCheck2, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", createdProductTagsIDForBatchDelete2), "")
	assert.Equal(t, fiber.StatusNotFound, respCheck2.StatusCode)

	respCheckKeep, _ := makeRequest("GET", fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", createdProductTagsIDToKeepForBatchTest), "")
	assert.Equal(t, fiber.StatusOK, respCheckKeep.StatusCode)
}

// Test_Z_CleanupProductTags: Kalan test verilerini temizler.
func Test_Z_CleanupProductTags(t *testing.T) {
	idList := []string{createdProductTagsID, createdProductTagsIDNotBatched, createdProductTagsIDToKeepForBatchTest}
	for _, pkPath := range idList {
		if pkPath != "" {
			url := fmt.Sprintf("%s/with-id/%s", "/catalog/product-tags", pkPath)
			makeRequest("DELETE", url, "")
		}
	}
	t.Log("Cleanup test ran.")
}
