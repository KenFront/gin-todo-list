package route

import (
	controller_users "github.com/KenFront/gin-todo-list/src/controller/users"

	"github.com/gin-gonic/gin"
)

func UseUser(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", controller_users.Add)
	}
}
