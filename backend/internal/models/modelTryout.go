package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Tryout struct {
	gorm.Model
	Title    string `gorm:"size:255;not null" json:"title"`
	Detail   string `gorm:"not null" json:"detail"`
	UserUsername *string `gorm:"index;constraint:OnDelete:SET NULL;" json:"username"` // Boleh null agar tidak wajib punya Tryout
	User     *User   `gorm:"foreignKey:UserUsername;references:Username;constraint:OnDelete:SET NULL;" json:"user"`
	

}




func (t *Tryout) CreateTryout() (*Tryout, error) {
	if err := Db.Create(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func GetTryoutByID(id string) (*Tryout, error) {
    var tryout Tryout
    uid, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        return nil, err
    }
    if err := Db.Where("id=?", uint(uid)).First(&tryout).Error; err != nil {
        return nil, err
    }
    return &tryout, nil
}


func GetTryoutsByUsername(username string) ([]Tryout, error) {
	var tryouts []Tryout

	if err := Db.Preload("User").Where("username = ?", username).Find(&tryouts).Error; err != nil {
		return nil, err
	}
	return tryouts, nil
}

func GetAllTryout() ([]Tryout, error) {
	var allTryout []Tryout

	if err := Db.Find(&allTryout).Error; err != nil {
		return nil, err
	}
	return allTryout, nil
}

func DeleteTryout(id string) (Tryout, error) {
	var tryout Tryout
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return Tryout{},err
	}

	if err := Db.Where("id = ?", uint(uid)).Delete(&tryout).Error; err != nil {
		return Tryout{},err
	}
	return tryout, nil
}
