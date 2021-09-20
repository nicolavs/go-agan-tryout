package model

import (
	"database/sql/driver"
	"time"
)

type answer string

const (
	OptionA answer = "a"
	OptionB answer = "b"
	OptionC answer = "c"
	OptionD answer = "d"
	OptionE answer = "e"
)

func (p *answer) Scan(value interface{}) error {
	*p = answer(value.([]byte))
	return nil
}

func (p answer) Value() (driver.Value, error) {
	return string(p), nil
}

// UserTestQuestionSolution godoc
type UserTestQuestionSolution struct {
	UserID         int32 `gorm:"primaryKey;index" json:"user_id"`
	TestQuestionID uint  `gorm:"primaryKey;index" json:"test_question_id"`
	User           User
	TestQuestion   TestQuestion
	Solution       answer    `gorm:"not null;type:answer_enum" json:"solution"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

// UserQuestionSolution godoc
type UserQuestionSolution struct {
	UserID     int32 `gorm:"primaryKey;index" json:"user_id"`
	QuestionID uint  `gorm:"primaryKey;index" json:"question_id"`
	User       User
	Question   Question
	Solution   answer    `gorm:"not null;type:answer_enum" json:"solution"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
