package models

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name   string `json:"name"`
	Active bool
}

func (p *Person) TableName() string {
	return "person"
}

var Persons []Person
