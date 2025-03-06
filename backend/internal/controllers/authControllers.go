package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/models"
)

func ListUser(c *fiber.Ctx) error {
	allUser, err := models.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	        "error": err.Error(),
	    })
	}
	return c.JSON(allUser)	
}

func GetUserById(c *fiber.Ctx) error {
	username := c.Params("username")
	tryout, err := models.GetUser(username)
	if err != nil {
	    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	        "error": "Username not found",
	    })
	}

    return c.JSON(tryout)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("username")
	tryout, err := models.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	        "error": "Username not found",
	    })
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Username deleted",
		"detail" : tryout,
	})
}

func AddUser(c *fiber.Ctx) error {
	user := models.User{}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "error processing data",
			"detail" : err.Error(),
		})
	}

	newUser, err := user.CreateUser()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "error processing data",
			"detail" : err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "new user Added",
		"data" : newUser,
	})
}
