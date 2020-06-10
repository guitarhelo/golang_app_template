package middleware

import (
	"backend_template/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Rbac struct {
	service.UserService
}

func UserPermissionMiddleware(c *gin.Context) {
	fmt.Println("Im a Rbac middleware!")

	// Pass on to the next-in-chain
	c.Next()
}
