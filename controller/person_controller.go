package controller

import (
	"backend_template/models"
	"backend_template/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type PersonController struct {
	BaseController
	service.PersonService
}

// NewPersonController PersonController的构造函数

func NewPersonController(p service.PersonService) PersonController {
	return PersonController{PersonService: p}
}

func (con *PersonController) FindAll(c *gin.Context) {

	persons := con.PersonService.FindAll()
	c.JSON(200, gin.H{
		"data": persons})

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
func (con *PersonController) FindByID(c *gin.Context) {
	var person models.Person
	id, _ := strconv.Atoi(c.Param("id"))
	person = con.PersonService.FindByID(uint(id))
	if person != (models.Person{}) {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "person get successfully!",
			"data": person})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "can not find!",
			"data": person})
	}
}
func (p *PersonController) Create(c *gin.Context) {
	var person models.Person
	err := c.BindJSON(&person)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdPerson := p.PersonService.Save(person)

	c.JSON(http.StatusOK, gin.H{"person": createdPerson})
}

func (p *PersonController) Update(c *gin.Context) {
	var personModel models.Person
	err := c.BindJSON(&personModel)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	person := p.PersonService.FindByID(uint(id))
	if person == (models.Person{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	person.Name = personModel.Name

	p.PersonService.Save(person)

	c.Status(http.StatusOK)
}

func (p *PersonController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	person := p.PersonService.FindByID(uint(id))
	if person == (models.Person{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.PersonService.Delete(person)

	c.Status(http.StatusOK)
}
