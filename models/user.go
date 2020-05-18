package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"varchar(40);unique_index"`
	Password string `gorm:"size:255"`
}
