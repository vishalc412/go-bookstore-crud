package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/golangpoc?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
