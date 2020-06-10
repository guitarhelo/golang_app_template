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
func GetUserRoles(username string) []dto.RoleDTO {
	var roles []dto.RoleDTO
	var rolename string
	// Execute the query

	rows, err := config.DB.Raw("SELECT c.rolename from `user` a , user_role b ,role c where username=? and a.id=b.user_id and b.role_id=c.id", username).Rows() // (*sql.Rows, error)
	defer rows.Close()

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for rows.Next() {
		var temp_role_dto dto.RoleDTO
		rows.Scan(&rolename)

		temp_role_dto.RoleName = rolename
		roles = append(roles, temp_role_dto)

	}

	return roles
}
func (p *UserRepository) LoadByName(name string) dto.UserDTO {
	var user models.User
	var roles []dto.RoleDTO
	var user_dto dto.UserDTO
	//var rolename string
	//m1 = make(map[int]string)
	config.DB.Where("username = ?", name).First(&user)
	//config.DB.Where("id = ?", user.ID).Find(&roles)
	//config.DB.Raw("SELECT c.rolename from `user` a , user_role b ,role c where username=? and a.id=b.user_id and b.role_id=c.id", "john").Scan(&roles)

	roles = GetUserRoles(user.Username)

	// Execute the query

	/*
		rows, err := config.DB.Raw("SELECT c.rolename from `user` a , user_role b ,role c where username=? and a.id=b.user_id and b.role_id=c.id", user.Username).Rows() // (*sql.Rows, error)
		defer rows.Close()

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		for rows.Next() {
	        var temp_role_dto dto.RoleDTO
			rows.Scan(&rolename)

		    temp_role_dto.RoleName=rolename
			roles=append(roles, temp_role_dto)

		}
	*/

	user_dto.User = user
	user_dto.Roles = roles
	user_dto.Resources = nil

	return user_dto
}
