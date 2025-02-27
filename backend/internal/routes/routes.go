package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/controllers"
)

func Routing(app *fiber.App) {
	api := app.Group("/")

	tryoutGroup := api.Group("/tryout")
	tryoutGroup.Get("/",controllers.ListTryout)
	tryoutGroup.Get("/:id",controllers.GetTryoutById)
	tryoutGroup.Get("/:userId",controllers.GetTryoutByUser)
	tryoutGroup.Delete("/:id",controllers.DeleteTryout)
	tryoutGroup.Put("/:id",controllers.UpdateTryout)
	tryoutGroup.Post("/", controllers.AddTryout)
}