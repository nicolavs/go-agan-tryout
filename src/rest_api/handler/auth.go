package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"rest_api/config"
	"rest_api/database"
	"rest_api/model"
	"time"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUsername(u string, userRole *model.GetUserRoleModel) error {
	var user model.User
	db := database.DB
	var roles []string

	if err := db.Where(&model.User{Username: u}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if err := db.Model(&model.Role{}).Select("roles.name").
		Joins("join user_roles on user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", user.ID).
		Find(&roles).Error; err != nil {
		return err
	}

	userRole.User = user
	userRole.Role = roles

	return nil
}

// Login godoc
// @Summary Get jwt token
// @Description Get jwt token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body model.LoginInput true "Login Input"
// @Success 200 {object} model.ResponseHTTP{data=string}
// @Failure 503 {object} model.ResponseHTTP
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var input model.LoginInput
	var ud model.GetUserRoleModel

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	username := input.Username
	pass := input.Password

	err := getUserByUsername(username, &ud)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid username", "data": err})
	}

	if !CheckPasswordHash(pass, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["roles"] = ud.Role
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(model.ResponseHTTP{
		Status:  "success",
		Message: "Success login",
		Data:    t,
	})
}
