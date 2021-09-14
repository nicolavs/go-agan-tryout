package config

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"time"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	var preFork bool
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	switch environmentStr := getenv("ENVIRONMENT", "development"); environmentStr {
	case "production":
		preFork = true
	default:
		preFork = false
	}

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		Prefork:     preFork,
	}
}

// Config func to get env value from key ---
func Config(key string) string {
	return os.Getenv(key)

}
