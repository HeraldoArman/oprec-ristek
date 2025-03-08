package models

import (
	// "strconv"

	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	QuestionID uint `gorm:"not null;index" json:"question_id"`
	Answer     bool `gorm:"not null" json:"answer"`
	Correct    bool `gorm:"not null" json:"correct"`
}
