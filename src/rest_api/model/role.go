package model

// Role model
type Role struct {
	ID          int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}

// UserRole relation
type UserRole struct {
	UserID int32 `gorm:"primaryKey;index" json:"user_id"`
	RoleID int32 `gorm:"primaryKey;index" json:"role_id"`
	User   User
	Role   Role
}

// GetUserRoleModel godoc
type GetUserRoleModel struct {
	User
	Role []string `json:"roles"`
}
