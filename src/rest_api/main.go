package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"rest_api/router"

	"github.com/gofiber/fiber/v2"
	"rest_api/config"
	"rest_api/database"
	_ "rest_api/docs"
)

// @title Agan Tryout App
// @version 1.0
// @description This is an API for Agan Tryout Application
func main() {
	database.ConnectDb()
	app := fiber.New(config.FiberConfig())

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
