// model database untuk submission
package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	TryoutID     uint   `gorm:"not null;index" json:"tryout_id"`
	UserUsername string `gorm:"index;not null" json:"user_username"`
	QuestionID   uint   `gorm:"not null;index" json:"question_id"`
	Answer       bool   `gorm:"not null" json:"answer"`
	IsCorrect    bool   `gorm:"not null" json:"isCorrect"`
	User         User   `gorm:"foreignKey:UserUsername;references:Username;constraint:OnDelete:CASCADE;" json:"-"`
	Tryout       Tryout `gorm:"foreignKey:TryoutID;constraint:OnDelete:CASCADE;" json:"-"`
}

func (s *Submission) CreateSubmission() (*Submission, error) {
	if err := Db.Create(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func GetAllSubmission() ([]Submission, error) {
	var allSubmission []Submission

	if err := Db.Order("created_at DESC").Find(&allSubmission).Error; err != nil {
		return nil, err
	}
	return allSubmission, nil
}

func GetSubmissionByTryoutID(id string) ([]Submission, error) {
	var submission []Submission

	if err := Db.Where("tryout_id=?", id).Find(&submission).Error; err != nil {
		return nil, err
	}
	return submission, nil
}

func GetSubmissionByUserUsername(username string) ([]Submission, error) {
	var submission []Submission

	if err := Db.Where("user_username=?", username).Find(&submission).Error; err != nil {
		return nil, err
	}
	return submission, nil
}

func GetSubmissionByID(id string) (Submission, error) {
	var submission Submission
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return Submission{}, err
	}

	if err := Db.Where("id=?", uint(uid)).First(&submission).Error; err != nil {
		return Submission{}, err
	}
	return submission, nil
}
func DeleteSubmission(id string) (Submission, error) {
	var submission Submission
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return Submission{}, err
	}

	if err := Db.Where("id=?", uint(uid)).Delete(&submission).Error; err != nil {
		return Submission{}, err
	}
	return submission, nil
}

func EvaluateSubmission(question Question, submission Submission) bool {
	return question.CorrectAnswer == submission.Answer
}

func GetTotalScore(username string, tryutid string) (int, int, error) {
	var submission []Submission
	var totalCorrect int
	var totalWrong int
	if err := Db.Where("user_username=? AND tryout_id=?", username, tryutid).Find(&submission).Error; err != nil {
		return 0, 0, err
	}
	for _, sub := range submission {
		if sub.IsCorrect {
			totalCorrect++
		} else {
			totalWrong++
		}
	}
	return totalCorrect, totalWrong, nil
}

func GetSubmissionByTryoutIDAndUser(username string, tryoutID string) ([]Submission, error) {
	var submission []Submission

	if err := Db.Where("user_username=? AND tryout_id=?", username, tryoutID).Find(&submission).Error; err != nil {
		return nil, err
	}
	return submission, nil
}
