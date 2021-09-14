package handler

import (
	"github.com/gofiber/fiber/v2"
	"rest_api/database"
	"rest_api/model"
)

// GetAllUsers from db
func GetAllUsers(c *fiber.Ctx) error {
	// query user table in the database
	rows, err := database.DB.Query("SELECT id, username, email, password FROM users order by name")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := model.Users{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		// Exit if we get an error
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		result.Users = append(result.Users, user)
	}
	return c.JSON(result)
}
