package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"rest_api/config"
	"rest_api/database"
	"rest_api/model"
	"strconv"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param skip query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} model.ResponseHTTP{data=[]model.User}
// @Failure 503 {object} model.ResponseHTTP
// @Router /user [get]
func GetAllUsers(c *fiber.Ctx) error {
	limitStr := c.Query("limit", "20")
	limit, _ := strconv.Atoi(limitStr)
	skipStr := c.Query("skip", "0")
	skip, _ := strconv.Atoi(skipStr)

	var users []model.User
	var count int64
	database.DB.Model(&model.User{}).Count(&count)
	database.DB.Find(&users).Limit(limit).Offset(skip)

	return c.Status(200).JSON(model.ResponseHTTP{
		Status:  "success",
		Message: "Success get",
		Data:    users,
		Count:   count,
	})
}

// CreateUser godoc
// @Summary Create new user
// @Description Create new user
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body model.CreateUser true "user"
// @Success 201 {object} model.ResponseHTTP{data=model.CreateUser}
// @Failure 503 {object} model.ResponseHTTP
// @Router /user [post]
func CreateUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	if !config.ValidToken(token, []string{"administrator"}) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
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
		Message: "Success create user",
		Data:    u,
	})
}
