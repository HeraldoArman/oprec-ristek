package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/heraldoarman/oprec-ristek/internal/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	routes.Routing(app)
	app.Listen(":3000")
}