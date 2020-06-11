package middleware

import (
	"backend_template/models/dto"
	"backend_template/service"
	"backend_template/util"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/vibrantbyte/go-antpath/antpath"
)

type Rbac struct {
	service.UserService
}

var userDto dto.UserDTO

func GetUserPermissionList() []string {
	var userPermissionList []string
	permissionList := userDto.Resources

	for _, v := range permissionList {
		//fmt.Println(k, v)
		for i := 0; i < len(v); i++ {
			userPermissionList = append(userPermissionList, v[i].Url)
		}
	}
	//fmt.Print(userPermissionList)

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
	userPermissionList = util.RemoveRepeatedElement(userPermissionList)
	fmt.Print(userPermissionList)
	if UseHavePermission(url, userPermissionList) != "" {
		// Pass on to the next-in-chain
		c.Next()
	} else {
		c.Abort()
		c.JSON(401, gin.H{
			"code": "401",

			"Message": "you have not permission to access this resources"})

	}

}
func UseHavePermission(requestURL string, userPermissionList []string) string {
	var matcher PathMatcher
	matcher = New()

	for i := 0; i < len(userPermissionList); i++ {
		if matcher.Match(userPermissionList[i], requestURL) {
			return userPermissionList[i]
		}

	}

	return ""
}
