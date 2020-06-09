package repository

import (
	"backend_template/config"
	"backend_template/models"
	"github.com/jinzhu/gorm"
)

type ResourceRepository struct {
	BaseRepository
	DB *gorm.DB
}

func NewResourceRepostiory(DB *gorm.DB) PersonRepository {
	return PersonRepository{DB: DB}
}

//FindAll Fetch all resource data

func (p *ResourceRepository) FindAll() []models.Resource {
	var resources []models.Resource
	config.DB.Find(&resources)

	return resources
}
func (p *ResourceRepository) FindByID(id uint) models.Resource {
	var resource models.Resource
	config.DB.First(&resource, id)

	return resource
}

func (p *ResourceRepository) Save(resource models.Resource) models.Resource {
	config.DB.Save(&resource)

	return resource
}

func (p *ResourceRepository) Delete(resource models.Resource) {
	config.DB.Delete(&resource)
}
