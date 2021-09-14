package model

// User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct {
	Users []User `json:"users"`
}
