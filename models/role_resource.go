package models

type RoleResource struct {
	ID         uint `json:"id"`
	ResourceID int  `json:"resource_id"`
	RoleID     int  `json:"role_id"`
}

func (p *RoleResource) TableName() string {
	return "role_resource"
}

var roleResources []RoleResource
