package route

import (
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/gin-gonic/gin"
)

func UseAuth(r *gin.Engine) {
	r.POST("/signin", controller.SignIn)
	r.POST("/signout", controller.SignOut)
	r.POST("/register", controller.Regiser)
}
