package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *User) TableName() string {
	return "user"
}

var users []User
