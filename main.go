package main

import (
	"fmt"
	"log"
	"os"

	api "data-bridge-examples/api_v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/maple-tech/baseline/events"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost port=55433 dbname=innovation_hub user=databridge password=databridge_demo_2026 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB bağlantısı başarısız: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	eventManager := events.NewEventManager()

	app := fiber.New(fiber.Config{
		AppName:      "Innovation Hub API",
		ErrorHandler: defaultErrorHandler,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	apiGroup := app.Group("/api/v1")
	api.Setup(apiGroup, db, eventManager)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("🚀 Innovation Hub API: http://localhost:%s\n", port)
	fmt.Printf("📚 Swagger UI: http://localhost:%s/swagger/index.html\n", port)
	log.Fatal(app.Listen(":" + port))
}

func defaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
