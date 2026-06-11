package controllers

import (
	"log"

	"attendance-backend/internal/models"
	"attendance-backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetOfficeSettings(c *fiber.Ctx) error {
	settings, err := services.GetOfficeSettings()
	if err != nil {
		log.Printf("[GetOfficeSettings] error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch settings",
		})
	}
	return c.JSON(settings)
}

func SaveOfficeSettings(c *fiber.Ctx) error {
	var settings models.OfficeSettings
	if err := c.BodyParser(&settings); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}
	if settings.RadiusMeters <= 0 {
		settings.RadiusMeters = 200.0
	}
	if err := services.SaveOfficeSettings(settings); err != nil {
		log.Printf("[SaveOfficeSettings] error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to save settings",
		})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Settings saved"})
}
