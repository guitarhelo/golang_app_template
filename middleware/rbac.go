package middleware

import (
	"backend_template/models/dto"
	"backend_template/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Rbac struct {
	service.UserService
}

func UserPermissionMiddleware(c *gin.Context) {
	var user_dto dto.UserDTO
	UserService := service.UserService{}
	fmt.Println("Im a Rbac middleware!")
	user_dto = UserService.LoadByName("john")

	fmt.Print(user_dto)
	// Pass on to the next-in-chain
	c.Next()
}
