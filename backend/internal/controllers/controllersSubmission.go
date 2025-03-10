// fungsi fungsi controller untuk submission
package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/models"
)

func ListAllSubmission(c *fiber.Ctx) error {
	allSubmission, err := models.GetAllSubmission()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(allSubmission)
}

func GetSubmissionByTryout(c *fiber.Ctx) error {
	tryoutId := c.Params("tryoutId")
	submissionList, err := models.GetSubmissionByTryoutID(tryoutId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(submissionList)
}

func GetSubmissionByUser(c *fiber.Ctx) error {
	username := c.Params("username")
	submissionList, err := models.GetSubmissionByUserUsername(username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(submissionList)
}

func GetSubmissionById(c *fiber.Ctx) error {
	id := c.Params("id")
	submission, err := models.GetSubmissionByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(submission)
}

func DeleteSubmission(c *fiber.Ctx) error {
	id := c.Params("id")
	submission, err := models.DeleteSubmission(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(submission)
}
func UpdateSubmission(c *fiber.Ctx) error {
	id := c.Params("id")
	submission, err := models.GetSubmissionByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "Submission not found",
			"detail": err.Error(),
		})
	}

	updateData := models.Submission{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Error parsing data",
			"detail": err.Error(),
		})
	}
	updateFields := map[string]interface{}{}

	updateFields["answer"] = updateData.Answer
	question, err := models.GetQuestionByID(fmt.Sprintf("%d", submission.QuestionID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to get question",
			"detail": err.Error(),
		})
	}
	updateFields["correct"] = models.EvaluateSubmission(question, updateData)

	if err := models.Db.Model(&submission).Updates(updateFields).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to update submission",
			"detail": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Submission Updated",
		"detail":  updateFields,
	})
}

func AddSubmission(c *fiber.Ctx) error {
	submission := models.Submission{}
	if err := c.BodyParser(&submission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Error parsing data",
			"detail": err.Error(),
		})
	}

	question, err := models.GetQuestionByID(fmt.Sprintf("%d", submission.QuestionID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to get question",
			"detail": err.Error(),
		})
	}
	submission.Correct = models.EvaluateSubmission(question, submission)

	if _, err := submission.CreateSubmission(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to create submission",
			"detail": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Submission Created",
		"detail":  submission,
	})
}

func GetSubmissionByTryoutIdAndUser(c *fiber.Ctx) error {
	tryoutId := c.Query("tryoutId")
	username := c.Query("username")
	if tryoutId == "" || username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Tryout ID and username must be provided",
		})
	}

	submission, err := models.GetSubmissionByTryoutIDAndUser(username, tryoutId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(submission)
}

func GetTotalScore(c *fiber.Ctx) error {
	tryoutId := c.Query("tryoutId")
	username := c.Query("username")
	if tryoutId == "" || username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Tryout ID and username must be provided",
		})
	}

	totalCorrect, totalWrong, err := models.GetTotalScore(username, tryoutId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"total_correct": totalCorrect,
		"total_wrong":   totalWrong,
		"score":         totalCorrect * 100 / (totalCorrect + totalWrong),
	})
}
