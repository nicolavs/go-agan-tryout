package router

import (
	"github.com/gofiber/fiber/v2"
	"rest_api/config"
	"rest_api/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api") // /api

	v1 := api.Group("/v1")               // /api/v1
	v1.Get("/user", handler.GetAllUsers) // /api/v1/user

	// Handle not founds
	app.Use(config.NotFound)
}
