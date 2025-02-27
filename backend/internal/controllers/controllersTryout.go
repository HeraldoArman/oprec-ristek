package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/models"
)

func ListTryout(c *fiber.Ctx) error {
	allTryout := models.GetAllTryout()
	return c.JSON(allTryout)	
}

func GetTryoutById(c *fiber.Ctx) error {
	id := c.Params("id")
	tryout, err := models.GetTryoutByID(id)
	if err != nil {
	    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	        "error": "Tryout not found",
	    })
	}

    return c.JSON(tryout)
}

func GetTryoutByUser(c *fiber.Ctx) error {
	id := c.Params("userId")
	tryout, err := models.GetTryoutsByUserID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	        "error": "Tryout not found",
	    })
	}
	return c.JSON(tryout)

}

func DeleteTryout(c *fiber.Ctx) error {
	id := c.Params("id")
	tryout, err := models.DeleteTryout(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	        "error": "Tryout not found",
	    })
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Tryout deleted",
		"detail" : tryout,
	})
}

func UpdateTryout(c *fiber.Ctx) error {
	updateTryout := models.Tryout{}
	err := c.BodyParser(&updateTryout)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "error parsing data",
			"detail" : err.Error(),
		})
	}
	id := c.Params("id")
	tryout, db := models.GetTryoutByID(id)
	if tryout == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "error, not found",
		})
	}
	tryout = &updateTryout
	err = db.Save(&tryout).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "error failed saving database",
			"detail" : err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Tryout Updated",
		"detail" : tryout,
	})

}

func AddTryout(c *fiber.Ctx) error {
	tryout := models.Tryout{}
	err := c.BodyParser(&tryout)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "error processing data",
			"detail" : err.Error(),
		})
	}

	newTryout := tryout.CreateTryout()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Tryout Added",
		"data" : newTryout,
	})
}
