package controllers

import (
	// "strconv"

	"fmt"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/models"
)

func ListTryout(c *fiber.Ctx) error {
	querySearch := c.Query("search")
	querySortByName := c.Query("sort_by_name")
	querySortByDate := c.Query("sort_by_date")

	var tryoutList []models.Tryout
	var err error

	if querySearch == "" {
		tryoutList, err = models.GetAllTryout()
	} else {
		tryoutList, err = models.GetTryoutByTitle(querySearch)
	}

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if querySortByName == "asc" {
		sort.Slice(tryoutList, func(i, j int) bool {
			return strings.ToLower(tryoutList[i].Title) < (tryoutList[j].Title)
		})
	} else if querySortByName == "desc" {
		sort.Slice(tryoutList, func(i, j int) bool {
			return strings.ToLower(tryoutList[i].Title) > strings.ToLower(tryoutList[j].Title)
		})
	}

	if querySortByDate == "asc" {
		sort.Slice(tryoutList, func(i, j int) bool {
			return tryoutList[i].CreatedAt.Before(tryoutList[j].CreatedAt)
		})
	} else if querySortByDate == "desc" {
		sort.Slice(tryoutList, func(i, j int) bool {
			return tryoutList[i].CreatedAt.After(tryoutList[j].CreatedAt)
		})
	}

	return c.JSON(tryoutList)
}
func GetTryoutByUser(c *fiber.Ctx) error {
	id := c.Params("userId")
	querySearch := c.Query("search")
	querySortByName := c.Query("sort_by_name")
	querySortByDate := c.Query("sort_by_date")

	var tryoutList []models.Tryout
	var err error

	if querySearch == "" {
		tryoutList, err = models.GetTryoutsByUsername(id)
	} else {
		tryoutList, err = models.GetTryoutsByUsernameAndTitle(id, querySearch)
		fmt.Println(tryoutList)
	}
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if querySortByName == "asc" {
		sort.Slice(tryoutList, func(i, j int) bool {
			return strings.ToLower(tryoutList[i].Title) < (tryoutList[j].Title)
		})
	} else if querySortByName == "desc" {
		sort.Slice(tryoutList, func(i, j int) bool {
			return strings.ToLower(tryoutList[i].Title) > strings.ToLower(tryoutList[j].Title)
		})
	}

	if querySortByDate == "asc" {
		sort.Slice(tryoutList, func(i, j int) bool {
			return tryoutList[i].CreatedAt.Before(tryoutList[j].CreatedAt)
		})
	} else if querySortByDate == "desc" {
		sort.Slice(tryoutList, func(i, j int) bool {
			return tryoutList[i].CreatedAt.After(tryoutList[j].CreatedAt)
		})
	}
	return c.JSON(tryoutList)

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

func DeleteTryout(c *fiber.Ctx) error {
	id := c.Params("id")
	tryout, err := models.DeleteTryout(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tryout not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tryout deleted",
		"detail":  tryout,
	})
}

func UpdateTryout(c *fiber.Ctx) error {
	id := c.Params("id")
	tryout, err := models.GetTryoutByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "Tryout not found",
			"detail": err.Error(),
		})
	}
	updateData := models.Tryout{}
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Error parsing data",
			"detail": err.Error(),
		})
	}

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
	if updateData.ImageLink != "" {
		updateFields["image_link"] = updateData.ImageLink
	}

	if updateData.ImageLink == "" {
		updateFields["image_link"] = ""
	}

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
			"error":  "Error parsing data",
			"detail": err.Error(),
		})
	}

	newTryout, err := tryout.CreateTryout()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Failed to create new question",
			"detail": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Tryout Added",
		"data":    newTryout,
	})
}
func ListTryoutByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	tryout := models.Tryout{}
	err := c.BodyParser(&tryout)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Error parsing data",
			"detail": err.Error(),
		})
	}
	tryoutList, err := models.GetTryoutByCategory(category)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(tryoutList)
}
