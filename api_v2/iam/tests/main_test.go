package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	api "data-bridge-examples/api_v2"
	testhelper "data-bridge-examples/api_v2/iam/testhelper"

	"github.com/maple-tech/baseline/events"
)

var (
	app *fiber.App
	db  *gorm.DB
)

// TestMain, tüm testlerden önce çalışır ve genel kurulumu yapar.
func TestMain(m *testing.M) {
	setup()
	// Testleri çalıştır
	code := m.Run()
	teardown()
	os.Exit(code)
}

// setup, veritabanı bağlantısını kurar, migration'ları çalıştırır ve Fiber uygulamasını hazırlar.
func setup() {
	// Test veritabanı bağlantı bilgilerini environment variables'dan al
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		// Örnek bir DSN, gerçek projede environment'tan okunmalı
		dsn = "host=localhost user=postgres password=postgres dbname=test_db port=5432 sslmode=disable"
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to test database: %v", err))
	}

	// Gerekirse burada veritabanı migration'ları çalıştırılır.
	// Örnek: db.AutoMigrate(&structures.User{}, &structures.Product{})

	app = fiber.New()
	// Ana API rotalarını kur
	api.Setup(app, db, events.NewEventManager())
}

// teardown, testler bittikten sonra veritabanını temizler.
func teardown() {
	// Burada test sırasında oluşturulan verileri temizleme işlemleri yapılır.
	// Örnek: db.Exec("TRUNCATE TABLE users, products RESTART IDENTITY")
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

// makeRequest, testlerde tekrar eden HTTP istek yapma işlemini basitleştiren bir yardımcı fonksiyondur.
func makeRequest(method, url, body string) (*http.Response, string) {
	resp, bodyBytes := testhelper.MakeRequest(app, method, url, body)
	return resp, string(bodyBytes)
}
