package models

type Resource struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}

func (p *Resource) TableName() string {
	return "resource"
}

var resources []Resource
