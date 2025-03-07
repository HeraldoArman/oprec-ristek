package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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
	app.Use(limiter.New(limiter.Config{
		Max:        100,             
		Expiration: 1 * time.Minute, 
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Izinkan request dari frontend
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	
	routes.Routing(app)
	app.Listen(":3000")
}