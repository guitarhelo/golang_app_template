package service

import (
	"backend_template/models"
	"backend_template/models/dto"
	"backend_template/repository"
)

type UserService struct {
	BaseService
	repository.UserRepository
}

// NewPersonService PersonService的构造函数
func NewUserService(p repository.UserRepository) UserService {
	return UserService{UserRepository: p}
}

//FindAll Fetch all person data

func (s *UserService) FindAll() []models.User {
	var users []models.User
	users = s.UserRepository.FindAll()

	return users
}

func (p *UserService) FindByID(id uint) models.User {
	return p.UserRepository.FindByID(id)
}
func (p *UserService) Save(user models.User) models.User {
	p.UserRepository.Save(user)

	return user
}

func (p *UserService) Delete(user models.User) {
	p.UserRepository.Delete(user)
}
func (p *UserService) LoadByName(name string) dto.UserDTO {
	var user_dto dto.UserDTO
	user_dto = p.UserRepository.LoadByName(name)
	return user_dto
}
