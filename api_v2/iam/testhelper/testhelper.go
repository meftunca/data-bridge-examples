package testhelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// --- Common Test Helper Functions ---

// MakeRequest, testlerde tekrar eden HTTP istek yapma işlemini basitleştiren bir yardımcı fonksiyondur.
func MakeRequest(app *fiber.App, method, url, body string) (*http.Response, []byte) {
	req := httptest.NewRequest(method, url, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "test-token")  // Test için sabit bir token
	req.Header.Set("X-Company", "test-company-id") // Test için sabit bir company id

	resp, err := app.Test(req, -1) // -1 timeout'u devre dışı bırakır
	if err != nil {
		panic(fmt.Sprintf("Test request failed: %v", err))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.Bytes()
	resp.Body.Close()

	return resp, responseBody
}

// ExtractIDFromResponse, JSON response'dan ID veya URID çıkarmak için kullanılır
func ExtractIDFromResponse(t *testing.T, body []byte, pkFieldJsonName string) string {
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	assert.NoError(t, err, "Failed to unmarshal response body for ID extraction")

	idVal, ok := result[pkFieldJsonName]
	assert.True(t, ok, fmt.Sprintf("Primary key '%s' not found in response", pkFieldJsonName))

	idStr := fmt.Sprintf("%v", idVal)
	assert.NotEmpty(t, idStr, "Extracted ID should not be empty")
	return idStr
}

// --- Pointer Helper Functions ---

// StringPtr, string değerler için pointer döndüren yardımcı fonksiyondur.
func StringPtr(s string) *string {
	return &s
}

// IntPtr, int değerler için pointer döndüren yardımcı fonksiyondur.
func IntPtr(i int) *int {
	return &i
}

// Int64Ptr, int64 değerler için pointer döndüren yardımcı fonksiyondur.
func Int64Ptr(i int64) *int64 {
	return &i
}

// Float64Ptr, float64 değerler için pointer döndüren yardımcı fonksiyondur.
func Float64Ptr(f float64) *float64 {
	return &f
}

// BoolPtr, bool değerler için pointer döndüren yardımcı fonksiyondur.
func BoolPtr(b bool) *bool {
	return &b
}

// --- Database Helper Functions ---

// SetupTestDB, test veritabanı bağlantısını kurar
func SetupTestDB(dsn string) (*gorm.DB, error) {
	// Bu fonksiyon baseline/db package'ından çağrılabilir
	// Şimdilik placeholder olarak bırakıyorum
	return nil, fmt.Errorf("SetupTestDB should be implemented using baseline/db")
}

// CleanupTestDB, test sırasında oluşturulan verileri temizler
func CleanupTestDB(db *gorm.DB, tableNames ...string) error {
	for _, tableName := range tableNames {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName)).Error; err != nil {
			return fmt.Errorf("failed to truncate table %s: %w", tableName, err)
		}
	}
	return nil
}

// --- Response Validation Helper Functions ---

// AssertSuccessResponse, başarılı response'ları validate eder
func AssertSuccessResponse(t *testing.T, resp *http.Response, expectedStatus int, body []byte) {
	assert.Equal(t, expectedStatus, resp.StatusCode, "Expected status code mismatch, body: %s", string(body))

	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	assert.NoError(t, err, "Response should be valid JSON")
}

// AssertErrorResponse, hata response'larını validate eder
func AssertErrorResponse(t *testing.T, resp *http.Response, expectedStatus int, body []byte) {
	assert.Equal(t, expectedStatus, resp.StatusCode, "Expected error status code mismatch, body: %s", string(body))

	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	assert.NoError(t, err, "Error response should be valid JSON")

	// Error response'da genellikle "error" veya "message" field'ı olur
	_, hasError := result["error"]
	_, hasMessage := result["message"]
	assert.True(t, hasError || hasMessage, "Error response should contain error or message field")
}

// --- JSON Helper Functions ---

// MarshalToJSON, struct'ı JSON string'e çevirir
func MarshalToJSON(t *testing.T, data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	assert.NoError(t, err, "Failed to marshal data to JSON")
	return string(jsonBytes)
}

// UnmarshalFromJSON, JSON string'i struct'a çevirir
func UnmarshalFromJSON(t *testing.T, jsonStr string, target interface{}) {
	err := json.Unmarshal([]byte(jsonStr), target)
	assert.NoError(t, err, "Failed to unmarshal JSON to target struct")
}
