package dto

import (
	"backend_template/models"
)

type UserDTO struct {
	User      models.User       `json:"user"`
	Roles     []models.Role     `json:"roles"`
	Resources []models.Resource `json:"resources"`
}
