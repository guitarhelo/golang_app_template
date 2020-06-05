package main

import (
	"backend_template/controller"
	"backend_template/repository"
	"backend_template/service"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initPersonAPI(db *gorm.DB) controller.PersonController {
	wire.Build(repository.NewPersonRepostiory, service.NewPersonService, controller.NewPersonController)

	return controller.PersonController{}
}