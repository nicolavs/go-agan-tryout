package model

// Question godoc
type Question struct {
	ID             uint  `gorm:"primaryKey;autoIncrement" json:"id"`
	QuestionTypeID int32 `gorm:"not null;index" json:"question_type_id"`
	QuestionType   QuestionType
	Text           string `gorm:"not null" json:"text"`
	Image          []byte `json:"image"`
	OptionA        string `gorm:"not null" json:"option_a"`
	OptionB        string `gorm:"not null" json:"option_b"`
	OptionC        string `gorm:"not null" json:"option_c"`
	OptionD        string `gorm:"not null" json:"option_d"`
	OptionE        string `gorm:"not null" json:"option_e"`
	Solution       answer `gorm:"not null;type:answer_enum" json:"solution"`
	SolutionImage  []byte `json:"solution_image"`
	WeightA        uint8  `json:"weight_a"`
	WeightB        uint8  `json:"weight_b"`
	WeightC        uint8  `json:"weight_c"`
	WeightD        uint8  `json:"weight_d"`
	WeightE        uint8  `json:"weight_e"`
}

// QuestionCategory  godoc
type QuestionCategory struct {
	ID   int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

// QuestionType godoc
type QuestionType struct {
	ID                 int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name               string `gorm:"uniqueIndex:question_type_category_1;not null;index" json:"name"`
	QuestionCategoryID int32  `gorm:"uniqueIndex:question_type_category_1;not null;index" json:"question_category_id"`
	QuestionCategory   QuestionCategory
}
