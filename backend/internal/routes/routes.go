// routing untuk semua fungsi controller
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/controllers"
)

func Routing(app *fiber.App) {
	api := app.Group("/")

	tryoutGroup := api.Group("/tryout")
	tryoutGroup.Get("/", controllers.ListTryout)
	tryoutGroup.Get("/:id", controllers.GetTryoutById)
	tryoutGroup.Get("/user/:userid", controllers.GetTryoutByUser)
	tryoutGroup.Delete("/:id", controllers.DeleteTryout)
	tryoutGroup.Put("/:id", controllers.UpdateTryout)
	tryoutGroup.Post("/", controllers.AddTryout)
	// tryoutGroup.Get("category/:category", controllers.ListTryoutByCategory)

	userGroup := api.Group("/user")
	userGroup.Get("/", controllers.ListUser)
	userGroup.Get("/:id", controllers.GetUserById)
	userGroup.Delete("/:id", controllers.DeleteUser)
	userGroup.Post("/", controllers.AddUser)

	questionGroup := api.Group("/question")
	questionGroup.Get("/", controllers.ListAllQuestion)
	questionGroup.Get("/tryout/:tryoutId", controllers.GetQuestionByTryout)
	questionGroup.Get("/:id", controllers.GetQuestionById)
	questionGroup.Delete("/:id", controllers.DeleteQuestion)
	questionGroup.Put("/:id", controllers.UpdateQuestion)
	questionGroup.Post("/", controllers.AddQuestion)

	submissionGroup := api.Group("/submission")
	submissionGroup.Get("/", controllers.ListAllSubmission)
	submissionGroup.Get("/tryout/user", controllers.GetSubmissionByTryoutIdAndUser)
	submissionGroup.Get("/tryout/:tryoutId", controllers.GetSubmissionByTryout)
	submissionGroup.Get("/user/:username", controllers.GetSubmissionByUser)
	submissionGroup.Get("/:id", controllers.GetSubmissionById)
	submissionGroup.Delete("/:id", controllers.DeleteSubmission)
	submissionGroup.Post("/", controllers.AddSubmission)
	submissionGroup.Put("/:id", controllers.UpdateSubmission)
	submissionGroup.Get("/totalScore", controllers.GetTotalScore)
	// submissionGroup.Post("/evaluate", controllers.EvaluateSubmission)

}
