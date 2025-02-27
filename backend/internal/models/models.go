package models

import (
	"github.com/heraldoarman/oprec-ristek/internal/database"
	"gorm.io/gorm"
)


var Db *gorm.DB

func init() {
	database.ConnectDB()
	Db = database.GetDB()
	Db.AutoMigrate(&Tryout{})
	Db.AutoMigrate(&User{})

	// db.AutoMigrate(&IzinAkses{})
	// db.AutoMigrate(&Penyihir{})

}
