package route

import (
	controller_auth "github.com/KenFront/gin-todo-list/src/controller/auth"
	"github.com/gin-gonic/gin"
)

func UseAuth(r *gin.Engine) {
	r.POST("/signin", controller_auth.SignIn)
	r.POST("/signout", controller_auth.SignOut)
}
