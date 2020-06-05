package config
import (
	"backend_template/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("mysql", "root:111111@tcp(localhost:3306)/mydb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Person{})

	DB = database
}