package handler

import (
	"github.com/gofiber/fiber/v2"
	"rest_api/database"
	"rest_api/model"
	"strconv"
)

// GetAllRoles godoc
// @Summary Get all roles
// @Description Get all roles
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param skip query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} model.ResponseHTTP{data=[]model.Role}
// @Failure 503 {object} model.ResponseHTTP
// @Router /user [get]
func GetAllRoles(c *fiber.Ctx) error {
	limitStr := c.Query("limit", "20")
	limit, _ := strconv.Atoi(limitStr)
	skipStr := c.Query("skip", "0")
	skip, _ := strconv.Atoi(skipStr)

	var result []model.Role
	var count int64
	database.DB.Model(&model.Role{}).Count(&count)
	database.DB.Find(&result).Limit(limit).Offset(skip)

	return c.Status(200).JSON(model.ResponseHTTP{
		Status:  "success",
		Message: "Success get",
		Data:    result,
		Count:   count,
	})
}
