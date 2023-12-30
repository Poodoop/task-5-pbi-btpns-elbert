package models

import (
	"github.com/jinzhu/gorm"
)

type Photo struct {
    gorm.Model
    Title    string    `json:"title"`
    Caption  string    `json:"caption"`
    PhotoURL string    `json:"photo_url"`
    UserID   uint      `json:"user_id"`
    User     User      `json:"user"`
}
