package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"rest_api/handler"
	"rest_api/middleware"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.Handler) // default

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api") // /api

	v1 := api.Group("/v1") // /api/v1

	// Auth
	auth := v1.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := v1.Group("/user")
	user.Get("/", middleware.JWTProtected(), handler.GetAllUsers)
	user.Post("/", middleware.JWTProtected(), handler.CreateUser)

	// Role

	// UserRole

	// Category

	// Question

	// Test

}
