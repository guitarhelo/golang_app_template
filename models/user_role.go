package models

type UserRole struct {
	ID     uint `json:"id"`
	UserID int  `json:"user_id"`
	RoleID int  `json:"role_id"`
}

func (p *UserRole) TableName() string {
	return "user_role"
}

var userRoles []UserRole
