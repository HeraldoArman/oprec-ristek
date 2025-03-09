package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	TryoutID      uint   `gorm:"not null;index" json:"tryout_id"`
	Question      string `gorm:"type:text;not null" json:"question"`
	CorrectAnswer bool   `gorm:"not null" json:"correct_answer"`
	Tryout        Tryout `gorm:"foreignKey:TryoutID;constraint:OnDelete:CASCADE;" json:"tryout"`
}

func (q *Question) CreateQuestion() (*Question, error) {
	if err := Db.Create(&q).Error; err != nil {
		return nil, err
	}
	return q, nil
}

func GetAllQuestion() ([]Question, error) {
	var allQuestion []Question

	if err := Db.Order("created_at DESC").Find(&allQuestion).Error; err != nil {
		return nil, err
	}
	return allQuestion, nil
}

func GetQuestionByTryoutID(id string) ([]Question, error) {
	var question []Question

	if err := Db.Where("tryout_id=?", id).Find(&question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func GetQuestionByID(id string) (Question, error) {
	var question Question
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return Question{}, err
	}

	if err := Db.Where("id=?", uint(uid)).First(&question).Error; err != nil {
		return Question{}, err
	}
	return question, nil
}

func DeleteQuestion(id string) (Question, error) {
	var question Question
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return Question{}, err
	}

	if err := Db.Where("id = ?", uint(uid)).Delete(&question).Error; err != nil {
		return Question{}, err
	}
	return question, nil
}
