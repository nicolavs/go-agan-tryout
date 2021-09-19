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
	TestID     int32 `gorm:"primaryKey;uniqueIndex:test_question_1" json:"test_id"`
	Test       Test
	QuestionID uint `gorm:"primaryKey" json:"question_id"`
	Question   Question
	Order      int32 `gorm:"not null;uniqueIndex:test_question_1" json:"order"`
}
