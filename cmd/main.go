package main

import (
	"log"
	"os"

	"attendance-backend/internal/config"
	"attendance-backend/internal/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — using system environment variables")
	}

	config.ConnectDB()

	app := fiber.New(fiber.Config{
		// Allow up to 10MB body for base64 image payloads
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api")
	attendance := api.Group("/attendance")
	attendance.Post("/clock-in", controllers.ClockIn)
	attendance.Get("/logs", controllers.GetLogs)

	settings := api.Group("/settings")
	settings.Get("/office", controllers.GetOfficeSettings)
	settings.Post("/office", controllers.SaveOfficeSettings)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
