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

		if config.GetDB().First(&user, "id = ?", id).Error != nil {
			panic(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
			})
		}
		c.Next()
	}
}

func UseAuthGuard(r *gin.RouterGroup) gin.IRoutes {
	return r.Use(authGuard())
}
