package router

import (
	"backend_template/controller"
	"backend_template/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	PersonController := controller.PersonController{}
	router := gin.Default()
	router.Use(middleware.UserPermissionMiddleware)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/persons", PersonController.FindAll)

		v1.GET("/persons/:id", PersonController.FindByID)
		v1.POST("/persons", PersonController.Create)
		v1.PUT("/persons/:id", PersonController.Update)
		v1.DELETE("/persons/:id", PersonController.Delete)

	}

	return router
}
