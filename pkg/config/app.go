package config

//main purpose of this file is to return DB  whenever necessary

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB //main db object
)

func Connect() {
	d, err := gorm.Open("mysql", "root:ARNAV19/simplerest?charset=utf8&parseTime=True&loc=Local") //mysql connection
	if err != nil {
		panic(err)

	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
