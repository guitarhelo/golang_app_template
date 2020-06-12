package router

import (
	"backend_template/controller"
	"backend_template/docs"
	"backend_template/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "backend_template/docs"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	"net/http"
)

func InitRouter() *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080/api/v2"
	docs.SwaggerInfo.BasePath = "/v2"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
