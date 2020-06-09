package router

import (
	"backend_template/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")

	// Pass on to the next-in-chain
	c.Next()
}
func InitRouter() *gin.Engine {

	PersonController := controller.PersonController{}
	router := gin.Default()
	router.Use(DummyMiddleware)
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
