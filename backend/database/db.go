package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./inventory.db")

	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
