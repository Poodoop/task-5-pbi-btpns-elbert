package database

import (
	"github.com/username/app/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})
}
