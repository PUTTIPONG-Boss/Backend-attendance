package controllers

import (
	"log"

	"attendance-backend/internal/models"
	"attendance-backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

func ClockIn(c *fiber.Ctx) error {
	var req models.ClockInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	entry, err := services.ClockIn(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	log.Printf("[ClockIn] employee_id=%s session=%s distance=%.2fm", entry.EmployeeID, entry.Session, entry.DistanceMeters)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":          "success",
		"message":         "Clock-in successful",
		"distance_meters": entry.DistanceMeters,
		"timestamp":       entry.Timestamp,
	})
}

func GetLogs(c *fiber.Ctx) error {
	logs, err := services.GetLogs()
	if err != nil {
		log.Printf("[GetLogs] error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch attendance logs",
		})
	}

	return c.JSON(logs)
}
