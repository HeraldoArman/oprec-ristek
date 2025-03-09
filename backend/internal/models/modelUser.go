// model database untuk user (tidak selesai sampai level 4 untuk autentikasi)
package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Username string    `gorm:"size:100;primaryKey;unique;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Name     string    `gorm:"size:100;not null" json:"name"`
	Password string    `gorm:"not null" json:"-"`
	Tryouts  *[]Tryout `gorm:"foreignKey:UserUsername" json:"tryouts"`
}

func (u *User) CreateUser() (*User, error) {
	if err := Db.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUser(username string) (*User, error) {
	var user User
	err := Db.Preload("Tryouts").Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // user not found
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUser() ([]User, error) {
	var allUser []User
	if err := Db.Find(&allUser).Error; err != nil {
		return nil, err
	}
	return allUser, nil
}

func DeleteUser(username string) (User, error) {
	var user User

	err := Db.Where("username = ?", username).Delete(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
