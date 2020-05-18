package models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	_, exists := os.LookupEnv("DEBUG_SQL")

	db.LogMode(exists)

	db.AutoMigrate(&User{})
	return db
}
