package controllers

import (
	// "strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/models"
)

func ListTryout(c *fiber.Ctx) error {
	allTryout, err := models.GetAllTryout()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	        "error": err.Error(),
	    })
	}
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
	tryout, err := models.GetTryoutsByUsername(id)
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
    id := c.Params("id")
    tryout, err := models.GetTryoutByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Tryout not found",
			"detail" : err.Error(),
        })
    }
    updateData := new(models.Tryout)
    if err := c.BodyParser(updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":  "Error parsing data",
            "detail": err.Error(),
        })
    }

    // Membuat map untuk update
    updateFields := map[string]interface{}{}
    if updateData.Title != "" {
        updateFields["title"] = updateData.Title
    }
    if updateData.Detail != "" {
        updateFields["detail"] = updateData.Detail
    }
    if updateData.UserUsername != nil {
        updateFields["user_username"] = updateData.UserUsername
    }

    // Update hanya field yang diubah
    if len(updateFields) > 0 {
        err := models.Db.Model(tryout).Updates(updateFields).Error
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error":  "Failed to update database",
                "detail": err.Error(),
            })
        }
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Tryout Updated",
        "detail":  updateFields,
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

	newTryout, err := tryout.CreateTryout()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "error processing data",
			"detail" : err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Tryout Added",
		"data" : newTryout,
	})
}
