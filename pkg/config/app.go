package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "root:root@123@/go_books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
