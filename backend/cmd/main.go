package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heraldoarman/oprec-ristek/internal/routes"
)

func main() {
	app := fiber.New()
	routes.Routing(app)
	app.Listen(":3000")
}