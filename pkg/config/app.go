package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect make a connection with mysql database
func Connect() {
	d, err := gorm.Open("mysql", "learning:Password@9352@/testing?charset=utf8&parseTime=True&")
	if err != nil {
		panic(any(err))
	}
	db = d
}

// GetDB return db object
func GetDB() *gorm.DB {
	return db
}
