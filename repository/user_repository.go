package repository

import (
	"backend_template/config"
	"backend_template/models"
	"backend_template/models/dto"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	BaseRepository
	DB *gorm.DB
}

func NewUserRepostiory(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

//FindAll Fetch all user data

func (p *UserRepository) FindAll() []models.User {
	var users []models.User
	config.DB.Find(&users)

	return users
}
func (p *UserRepository) FindByID(id uint) models.User {
	var user models.User
	config.DB.First(&user, id)

	return user
}

func (p *UserRepository) Save(user models.User) models.User {
	config.DB.Save(&user)

	return user
}

func (p *UserRepository) Delete(user models.User) {
	config.DB.Delete(&user)
}
func (p *UserRepository) LoadByName(name string) dto.UserDTO {
	var user models.User
	var roles []models.Role
	var user_dto dto.UserDTO
	config.DB.Where("name = ?", name).First(&user)
	config.DB.Where("id = ?", user.ID).Find(&roles)
	user_dto.User = user
	user_dto.Roles = roles
	user_dto.Resources = nil

	return user_dto
}
