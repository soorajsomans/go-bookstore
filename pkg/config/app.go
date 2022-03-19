package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	connectionString := ""
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
