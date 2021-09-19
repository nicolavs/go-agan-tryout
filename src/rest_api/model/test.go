package model

import (
	"database/sql"
	"time"
)

// Test godoc
type Test struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime `gorm:"index"`
	Remarks   string       `gorm:"not null"`
}

// TestQuestion godoc
type TestQuestion struct {
	TestID     uint `gorm:"primaryKey" json:"test_id"`
	Test       Test
	QuestionID uint `gorm:"primaryKey" json:"question_id"`
	Question   Question
}
