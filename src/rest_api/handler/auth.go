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

func getUserByUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: u}).Find(&user).Limit(1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
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
	var ud model.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	username := input.Username
	pass := input.Password

	user, err := getUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}

	ud = model.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if !CheckPasswordHash(pass, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	//claims["roles"] =
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
