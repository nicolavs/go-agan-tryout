package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"rest_api/config"
	"rest_api/database"
	"rest_api/model"
)



func validToken(t *jwt.Token, rolesNeeded []string) bool {
	claims := t.Claims.(jwt.MapClaims)
	roles := claims["roles"].([]string)
	if config.Contains(roles, "administrator") {
		return true
	}

	for _, r := range rolesNeeded {
		if config.Contains(roles, r) {
			return true
		}
	}

	return false
}

// GetAllUsers from db
func GetAllUsers(c *fiber.Ctx) error {
	_ = c.Locals("user").(*jwt.Token)
	var users []model.User
	database.DB.Find(&users)

	return c.Status(200).JSON(model.ResponseHTTP{
		Status:  "success",
		Message: "Success login",
		Data:    users,
	})
}

// CreateUser into db
func CreateUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, []string{"administrator"}) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}
	// New Employee struct
	u := new(model.CreateUser)
	db := database.DB

	if err := c.BodyParser(u); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := config.HashPassword(u.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	newUser := model.User{
		Username: u.Username,
		Email:    u.Email,
		Password: hash,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.Status(201).JSON(model.ResponseHTTP{
		Status:  "success",
		Message: "Success login",
		Data:    u,
	})
}
