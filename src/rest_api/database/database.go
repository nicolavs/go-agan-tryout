package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"os"
	"rest_api/config"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"rest_api/model"
)

var DB *gorm.DB

// ConnectDb function
func ConnectDb() {
	var err error
	p := config.Config("DB_PORT")
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"),
		port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{})

	DB = db
}
