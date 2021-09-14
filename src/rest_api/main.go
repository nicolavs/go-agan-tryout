package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"rest_api/router"

	"github.com/gofiber/fiber/v2"
	"rest_api/config"
)

func main() {
	app := fiber.New(config.FiberConfig())

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
