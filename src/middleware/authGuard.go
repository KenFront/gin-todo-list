package middleware

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func authGuard(c *gin.Context) {
	id, err := util.GetUserIdByToken(c)
	if err != nil {
		util.DeleteAuth(c)
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ErrorType(err.Error()),
			Error:      err,
		})
	}

	user := model.User{
		ID: id,
	}

	if err := config.GetDB().First(&user, "id = ?", id).Error; err != nil {
		util.DeleteAuth(c)
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_SIGN_IN_FAILED,
			Error:      err,
		})
	}

	controller.SetUserId(c, id)

	c.Next()
}

func UseAuthGuard(r *gin.RouterGroup) gin.IRoutes {
	return r.Use(authGuard)
}
