package controller

import (
	"backend_template/models"
	"backend_template/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	BaseController
	service.UserService
}

// NewPersonController PersonController的构造函数

func NewUserController(p service.UserService) UserController {
	return UserController{UserService: p}
}

func (con *UserController) FindAll(c *gin.Context) {

	users := con.UserService.FindAll()
	c.JSON(200, gin.H{
		"data": users})

}
func (con *UserController) LogOut(c *gin.Context) {

	c.JSON(200, gin.H{
		"Message": "user logout"})

}

func (con *UserController) Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"Message": "User Login"})

}

/*
func (con *PersonController) CreatePerson(c *gin.Context) {

	c.JSON(200, gin.H{
		"Message": "create a person"})

}
func (con *PersonController) UpdatePerson(c *gin.Context) {

	c.JSON(200, gin.H{
		"Message": "update a person"})

}
func (con *PersonController) DeletePerson(c *gin.Context) {

	c.JSON(200, gin.H{
		"Message": "delete a person"})

}

*/
func (con *UserController) FindByID(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	user = con.UserService.FindByID(uint(id))
	if user != (models.User{}) {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "user get successfully!",
			"data": user})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "can not find any user!",
			"data": user})
	}
}
func (p *UserController) Create(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdUser := p.UserService.Save(user)

	c.JSON(http.StatusOK, gin.H{"person": createdUser})
}

func (p *UserController) Update(c *gin.Context) {
	var userModel models.User
	err := c.BindJSON(&userModel)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))
	if user == (models.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	user.Username = userModel.Username

	p.UserService.Save(user)

	c.Status(http.StatusOK)
}

func (p *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))
	if user == (models.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.UserService.Delete(user)

	c.Status(http.StatusOK)
}
