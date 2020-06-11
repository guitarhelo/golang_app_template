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

var userDto dto.UserDTO

func GetUserPermissionList() []dto.ResourceDTO {
	var userPermissionList []dto.ResourceDTO
	permissionList := userDto.Resources

	for _, v := range permissionList {
		//fmt.Println(k, v)
		for i := 0; i < len(v); i++ {
			userPermissionList = append(userPermissionList, v[i])
		}
	}
	fmt.Print(userPermissionList)

	return userPermissionList
}
func UserPermissionMiddleware(c *gin.Context) {
	UserService := service.UserService{}
	fmt.Println("Im a Rbac middleware!")
	userDto = UserService.LoadByName("john")
	url := c.Request.RequestURI
	fmt.Print(url)
	fmt.Print(userDto)
	userPermissionList := GetUserPermissionList()
	fmt.Print(userPermissionList)
	// Pass on to the next-in-chain
	c.Next()
}
