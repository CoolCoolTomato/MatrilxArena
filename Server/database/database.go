package database

import (
	"gorm.io/gorm"
)

var Database *gorm.DB

func GetDatabase() *gorm.DB {
	return Database
}
