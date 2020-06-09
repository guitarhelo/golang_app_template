package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	ID       uint   `json:"id"`
	Rolename string `json:"rolename"`
}

func (p *Role) TableName() string {
	return "role"
}

var roles []Role
