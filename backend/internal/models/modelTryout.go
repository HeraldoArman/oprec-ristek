// model database untuk tryout
package models

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type KategoriTryout string

// enum type
const (
	Saintek     KategoriTryout = "Saintek"
	Soshum      KategoriTryout = "Soshum"
	Bahasa      KategoriTryout = "Bahasa"
	Pemrograman KategoriTryout = "Pemrograman"
	Lainnya     KategoriTryout = "Lainnya"
)

func (kt *KategoriTryout) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan KategoriTryout: %v", value)
	}
	*kt = KategoriTryout(strValue)
	return nil
}

func (kt KategoriTryout) Value() (interface{}, error) {
	return string(kt), nil
}

type Tryout struct {
	gorm.Model
	Title        string         `gorm:"size:255;not null" json:"title"`
	Detail       string         `gorm:"not null" json:"detail"`
	ImageLink    string         `gorm:"size:255" json:"image"`
	Kategori     KategoriTryout `gorm:"type:kategori_tryout" json:"kategori"`
	Questions    []Question     `gorm:"foreignKey:TryoutID;constraint:OnDelete:CASCADE;" json:"questions"`
	Submission   []Submission   `gorm:"foreignKey:TryoutID;constraint:OnDelete:CASCADE;" json:"submission"`
	UserUsername *string        `gorm:"index;constraint:OnDelete:SET NULL;" json:"username"`
	User         *User          `gorm:"foreignKey:UserUsername;references:Username;constraint:OnDelete:SET NULL;" json:"user"`
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

	if err := Db.Preload("User").Order("created_at DESC").Where("user_username = ?", username).Find(&tryouts).Error; err != nil {
		return nil, err
	}

	return tryouts, nil
}
func GetTryoutsByUsernameAndTitle(username string, query string) ([]Tryout, error) {
	var tryouts []Tryout

	db := Db.Preload("User").Order("created_at DESC").Where("user_username = ?", username)

	if query != "" {
		db = db.Where("LOWER(title) LIKE LOWER(?)", "%"+query+"%")
	}

	if err := db.Find(&tryouts).Error; err != nil {
		return nil, err
	}

	return tryouts, nil
}

func GetAllTryout() ([]Tryout, error) {
	var allTryout []Tryout

	if err := Db.Order("created_at DESC").Find(&allTryout).Error; err != nil {
		return nil, err
	}
	return allTryout, nil
}

func GetTryoutByTitle(query string) ([]Tryout, error) {
	var tryouts []Tryout

	if err := Db.Where("LOWER(title) LIKE LOWER(?)", "%"+query+"%").Find(&tryouts).Error; err != nil {
		return nil, err
	}

	return tryouts, nil
}

func DeleteTryout(id string) (Tryout, error) {
	var tryout Tryout
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return Tryout{}, err
	}

	if err := Db.Where("id = ?", uint(uid)).Delete(&tryout).Error; err != nil {
		return Tryout{}, err
	}
	return tryout, nil
}

func GetTryoutByCategory(category string) ([]Tryout, error) {
	var tryouts []Tryout

	if err := Db.Where("kategori = ?", category).Find(&tryouts).Error; err != nil {
		return nil, err
	}

	return tryouts, nil
}

func IsThereSubmission(id uint) bool {
	var submission Submission

	if err := Db.Where("tryout_id = ?", id).First(&submission).Error; err != nil {
		return false
	}
	return true
}

func GetTryoutByCategoryAndTitle(category string, query string) ([]Tryout, error) {
	var tryouts []Tryout

	db := Db.Where("kategori = ?", category)

	if query != "" {
		db = db.Where("title LIKE ?", "%"+query+"%")
	}

	if err := db.Find(&tryouts).Error; err != nil {
		return nil, err
	}

	return tryouts, nil
}
