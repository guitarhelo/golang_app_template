package repository

import (
	"backend_template/config"
	"backend_template/models"
	"github.com/jinzhu/gorm"
)

type RoleRepository struct {
	BaseRepository
	DB *gorm.DB
}

func NewRoleRepostiory(DB *gorm.DB) RoleRepository {
	return RoleRepository{DB: DB}
}

//FindAll Fetch all Role data

func (p *RoleRepository) FindAll() []models.Role {
	var roles []models.Role
	config.DB.Find(&roles)

	return roles
}
func (p *RoleRepository) FindByID(id uint) models.Role {
	var role models.Role
	config.DB.First(&role, id)

	return role
}

func (p *RoleRepository) Save(role models.Role) models.Role {
	config.DB.Save(&role)

	return role
}

func (p *RoleRepository) Delete(role models.Role) {
	config.DB.Delete(&role)
}
