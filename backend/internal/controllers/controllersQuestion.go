package controllers

import (
	// "gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/models"
)

func ListAllQuestion(c *fiber.Ctx) error {

	allQuestion, err := models.GetAllQuestion()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(allQuestion)
}
