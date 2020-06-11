package dto

import (
	"backend_template/models"
)

type UserDTO struct {
	User      models.User              `json:"user"`
	Roles     []RoleDTO                `json:"roles"`
	Resources map[string][]ResourceDTO `json:"resources"`
}
