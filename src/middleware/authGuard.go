package middleware

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func authGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := util.GetUserId(c)

		user := model.User{
			ID: id,
		}

		result := config.GetDB().First(&user, "id = ?", id)

		if result.Error != nil {
			panic(&util.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  result.Error.Error(),
			})
		}
		c.Next()
	}
}

func UseAuthGuard(r *gin.RouterGroup) gin.IRoutes {
	return r.Use(authGuard())
}
