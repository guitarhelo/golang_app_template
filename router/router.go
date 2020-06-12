package router

import (
	"backend_template/controller"
	"backend_template/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	"backend_template/docs"
)

func InitRouter() *gin.Engine {

	PersonController := controller.PersonController{}
	router := gin.Default()
	router.Use(middleware.UserPermissionMiddleware)
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/css", "../templates/css")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	admin := router.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
