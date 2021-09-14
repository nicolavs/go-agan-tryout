package model

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model

	ID       uint   `gorm:"primarykey;auto_increment;not_null" json:"id"`
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
}

type CreateUser struct {
	Username string ` json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Identity string `json:"username"`
	Password string `json:"password"`
}

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
