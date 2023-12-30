package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=username dbname=mydb sslmode=disable password=mypassword")
	if err != nil {
		panic(err.Error())
	}

	DB = db

	return db
}
