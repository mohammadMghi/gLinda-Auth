package datastore

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"log"


)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "admin:admin4321@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}