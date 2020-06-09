package repository

import (
	"backend_template/config"
	"backend_template/models"
	"backend_template/util"
	"github.com/jinzhu/gorm"
)

type PersonRepository struct {
	BaseRepository
	DB *gorm.DB
}

func NewPersonRepostiory(DB *gorm.DB) PersonRepository {
	return PersonRepository{DB: DB}
}

//FindAll Fetch all person data

func (p *PersonRepository) FindAll() []models.Person {
	var persons []models.Person
	config.DB.Find(&persons)

	return persons
}
func (p *PersonRepository) FindByID(id uint) models.Person {
	var person models.Person
	config.DB.First(&person, id)

	return person
}

func (p *PersonRepository) Save(person models.Person) models.Person {
	config.DB.Save(&person)

	return person
}

func (p *PersonRepository) Delete(person models.Person) {
	config.DB.Delete(&person)
}

func (p *PersonRepository) getResultSet(page int, rowsPerPage int) (*util.Response, error) {
	//your empty result set
	resultSet := make([]*models.Person, 0, 12)
	//prepare a handler to query
	handler := config.DB.
		Model(&models.Person{}).
		Where(&models.Person{Active: true})
	//use PageQuery to get data
	resp, err := util.PageQuery(page, rowsPerPage, handler, &resultSet)
	return resp, err
}
