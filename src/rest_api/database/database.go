package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"rest_api/config"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"rest_api/model"
)

var DB *gorm.DB

var tables []interface{}

// ConnectDb function
func ConnectDb() {
	var err error

	p := config.Config("POSTGRES_PORT")
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	// TODO turn off logging at production

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("POSTGRES_HOST"),
		port, config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// create enum
	db.Exec("CREATE TYPE answer_enum AS ENUM ('a','b','c','d','e')")

	// append list of table here for table auto migration
	tables = append(tables, &model.User{})
	tables = append(tables, &model.Role{})
	tables = append(tables, &model.UserRole{})
	tables = append(tables, &model.QuestionCategory{})
	tables = append(tables, &model.QuestionType{})
	tables = append(tables, &model.Question{})
	tables = append(tables, &model.Test{})
	tables = append(tables, &model.TestQuestion{})
	tables = append(tables, &model.TestQuestion{})
	tables = append(tables, &model.UserQuestionSolution{})
	tables = append(tables, &model.UserTestQuestionSolution{})
	log.Println("running migrations")
	for _, table := range tables {
		err := db.AutoMigrate(table)
		if err != nil {
			log.Fatal("Failed to create table")
		}
	}
	createAdminUser(db)
	DB = db
}

func createAdminUser(db *gorm.DB) {
	var user model.User
	var role model.Role
	var userRole model.UserRole

	result := db.Where(&model.Role{Name: "administrator"}).Limit(1).Find(&role)
	if result.RowsAffected == 0 {
		role.Name = "administrator"
		role.Description = "administrator role"
		if err := db.Create(&role).Error; err != nil {
			log.Fatal("Failed to create admin role. \n", err)
		}
	}

	result = db.Where(&model.User{Username: "admin"}).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		hash, err := config.HashPassword(config.Config("ADMIN_DEFAULT_PASSWORD"))
		if err != nil {
			log.Fatal("Failed to create admin user. \n", err)
		}
		user.Username = "admin"
		user.Email = "aganpro@gmail.com"
		user.Password = hash
		user.Disable = false
		if err := db.Create(&user).Error; err != nil {
			log.Fatal("Failed to create admin user. \n", err)
		}

	}

	result = db.Where(&model.UserRole{
		UserID: user.ID,
		RoleID: role.ID,
	}).Limit(1).Find(&userRole)

	if result.RowsAffected == 0 {
		userRole.UserID = user.ID
		userRole.RoleID = role.ID
		if err := db.Create(&userRole).Error; err != nil {
			log.Fatal("Failed to create admin user role. \n", err)
		}
	}

}
