package controllers

import (
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
func GetQuestionByTryout(c *fiber.Ctx) error {
	tryoutId := c.Params("tryoutId")
	questionList, err := models.GetQuestionByTryoutID(tryoutId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(questionList)
}

func GetQuestionById(c *fiber.Ctx) error {
	id := c.Params("id")
	question, err := models.GetQuestionByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(question)
}

func DeleteQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	question, err := models.DeleteQuestion(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(question)
}

func UpdateQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	question, err := models.GetQuestionByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	updateData := models.Question{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Error parsing data",
			"detail": err.Error(),
		})
	}
	updateFields := map[string]interface{}{}
	if updateData.Question != "" {
		updateFields["question"] = updateData.Question
	}
	updateFields["correct_answer"] = updateData.CorrectAnswer
	err = models.Db.Model(&question).Updates(updateFields).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to update database",
			"detail": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tryout Updated",
		"detail":  updateFields,
	})
}

func AddQuestion(c *fiber.Ctx) error {
	question := models.Question{}
	err := c.BodyParser(&question)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Error parsing data",
			"detail": err.Error(),
		})
	}

	newQuestion, err := question.CreateQuestion()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to create new question",
			"detail": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Question Added",
		"detail":  newQuestion,
	})
}
