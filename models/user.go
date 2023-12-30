package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username  string    `json:"username" binding:"required"`
    Email     string    `json:"email" binding:"required,email,unique"`
    Password  string    `json:"password" binding:"required,min=6"`
    Photos    []Photo   `json:"photos" gorm:"foreignkey:UserID"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
