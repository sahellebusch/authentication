package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Username: "jimcarey", Password: "legend"})
	db.Create(&User{Username: "santaclause", Password: "northpole"})

	return db
}
