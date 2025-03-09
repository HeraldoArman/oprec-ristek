// connect ke database dan migrate database
package models

import (
	"github.com/heraldoarman/oprec-ristek/internal/database"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	database.ConnectDB()
	Db = database.GetDB()
	// Db.Exec(`
	// 	DO $$ BEGIN
	// 		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_tryout') THEN
	// 			CREATE TYPE kategori_tryout AS ENUM ('Saintek', 'Soshum', 'Bahasa', 'Pemrograman', 'Lainnya');
	// 		END IF;
	// 	END $$;
	// `)
	Db.AutoMigrate(&User{}, &Tryout{}, &Question{}, &Submission{})
	// Db.AutoMigrate()
}
