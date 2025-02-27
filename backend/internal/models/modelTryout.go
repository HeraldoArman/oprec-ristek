package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Tryout struct {
	gorm.Model
	Title    	string `gorm:"size:255;not null" json:"title"`
	Detail		string `gorm:"not null" json:"detail"`
	// UserID   	uint   `gorm:"not null" json:"user_id"`
	// User     	User   `json:"user"`
}


func (t *Tryout) CreateTryout() *Tryout {
	Db.Create(&t)
	return t
}

func GetTryoutByID(id string) (*Tryout, *gorm.DB) {
	var tryout Tryout

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, nil
	}

	dbVar := Db.Preload("User").Where("id = ?", uint(uid)).First(&tryout)

	if dbVar.Error != nil {
		return nil, dbVar
	}

	return &tryout, dbVar
}


func GetTryoutsByUserID(userID string) ([]Tryout, *gorm.DB) {
	var tryouts []Tryout
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return nil, nil
	}
	dbVar := Db.Preload("User").Where("user_id = ?", uid).Find(&tryouts)
	return tryouts, dbVar
}


func GetAllTryout() []Tryout {
	var allTryout []Tryout
	Db.Find(&allTryout)
	return allTryout
}

func DeleteTryout(id string) (Tryout, error) {
	var tryout Tryout

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return Tryout{}, nil
	}

	err = Db.Where("id=?", uid).Delete(&tryout).Error
	if err != nil {
		return Tryout{}, err
	}
	return tryout, nil
}

