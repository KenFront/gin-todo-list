package route

import (
	"github.com/KenFront/gin-todo-list/src/config"
	controller_users "github.com/KenFront/gin-todo-list/src/controller/users"
	"github.com/KenFront/gin-todo-list/src/middleware"
	"github.com/KenFront/gin-todo-list/src/util"

	"github.com/gin-gonic/gin"
)

func UseUser(r *gin.Engine) {
	db := config.GetDB()

	users := r.Group("/users")
	{
		users.POST("", controller_users.Add)
	}

	usersWithAuth := r.Group("/users")
	middleware.UseAuthGuard(usersWithAuth)
	{
		usersWithAuth.GET("/self", controller_users.Get(controller_users.GetProps{
			Db:               db,
			GetUserIdByToken: util.GetUserIdByToken,
		}))
		usersWithAuth.DELETE("/self", controller_users.Delete(controller_users.DeleteProps{
			Db:               db,
			GetUserIdByToken: util.GetUserIdByToken,
		}))
	}
}
