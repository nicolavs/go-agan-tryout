package model

import (
	"database/sql"
	"time"
)

// Test godoc
type Test struct {
	ID        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime `gorm:"index"`
	Remarks   string       `gorm:"not null;size:256"`
}

// TestQuestion godoc
type TestQuestion struct {
	TestQuestionID uint  `gorm:"primaryKey;autoIncrement" json:"test_question_id"`
	TestID         int32 `gorm:"uniqueIndex:test_question_1;not null;index" json:"test_id"`
	Test           Test
	QuestionID     uint `gorm:"uniqueIndex:test_question_1;not null;index" json:"question_id"`
	Question       Question
}

// UserTest godoc
type UserTest struct {
	UserID int32 `gorm:"primaryKey;index" json:"user_id"`
	TestID int32 `gorm:"primaryKey;index" json:"test_id"`
	User   User
	Test   Test
}
