package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
