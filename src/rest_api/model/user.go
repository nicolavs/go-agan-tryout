package model

// User model
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not_null" json:"id"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
}

type CreateUser struct {
	Username string `json:"username"`
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
