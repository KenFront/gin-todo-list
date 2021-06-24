package route

import (
	"github.com/KenFront/gin-todo-list/src/config"
	controller_auth "github.com/KenFront/gin-todo-list/src/controller/auth"
	"github.com/gin-gonic/gin"
)

func UseAuth(r *gin.Engine) {
	db := config.GetDB()

	r.POST("/signin", controller_auth.SignIn(controller_auth.SignInProps{
		Db: db,
	}))
	r.POST("/signout", controller_auth.SignOut)
}
