package route

import (
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/gin-gonic/gin"
)

func UseUser(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", controller.AddUser)
	}
}
