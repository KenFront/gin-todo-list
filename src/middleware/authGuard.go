package middleware

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func authGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := util.CheckAuth(c)
		if err != nil {
			util.DeleteAuth(c)
			panic(&util.ApiError{
				StatusCode: http.StatusUnauthorized,
				ErrorType:  err.Error(),
			})
		}

		id, err := uuid.Parse(claim.UserId)
		if err != nil {
			panic(&util.ApiError{
				StatusCode: http.StatusUnauthorized,
				ErrorType:  err.Error(),
			})
		}

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
