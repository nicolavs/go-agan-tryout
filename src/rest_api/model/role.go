package model

// Role model
type Role struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}

// UserRole relation
type UserRole struct {
	UserID uint `gorm:"primaryKey;index" json:"user_id"`
	RoleID uint `gorm:"primaryKey;index" json:"role_id"`
	User   User
	Role   Role
}
