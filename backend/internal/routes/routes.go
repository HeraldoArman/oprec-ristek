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
	tryoutGroup.Get("/user/:userid",controllers.GetTryoutByUser)
	tryoutGroup.Delete("/:id",controllers.DeleteTryout)
	tryoutGroup.Put("/:id",controllers.UpdateTryout)
	tryoutGroup.Post("/", controllers.AddTryout)

	userGroup := api.Group("/user")
	userGroup.Get("/",controllers.ListUser)
	userGroup.Get("/:id",controllers.GetUserById)
	// userGroup.Get("/tryout/:userId",controllers.GetTryoutByUser)
	userGroup.Delete("/:id",controllers.DeleteUser)
	userGroup.Post("/", controllers.AddUser)
}