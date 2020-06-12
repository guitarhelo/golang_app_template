package main

import (
	"backend_template/config"
	"backend_template/models"
	"backend_template/router"
	"github.com/gin-gonic/gin"
)

func welcomePathParam(c *gin.Context) {
	name := c.Param("name")
	welcomeMessage := models.Person{Name: name}

	c.JSON(200, gin.H{"message": welcomeMessage})
}
func main() {
	router := router.InitRouter()
	router.GET("/welcome/:name", welcomePathParam)
	config.ConnectDataBase()
	router.Run(":8080")
	defer config.DB.Close()
}
