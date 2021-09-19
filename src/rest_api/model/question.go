package model

// Question godoc
type Question struct {
	ID             uint `gorm:"primaryKey;autoIncrement" json:"id"`
	QuestionTypeID uint `gorm:"not null: index" json:"question_type_id"`
	QuestionType   QuestionType
}

// QuestionType godoc
type QuestionType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}
