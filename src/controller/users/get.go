package controller_users

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetProps struct {
	Db *gorm.DB
}

func Get(p GetProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		id := controller.GetUserId(c)

		if err := p.Db.First(&user, "id = ?", id).Error; err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GET_USER_FAILED,
				Error:      err,
			})
		}

		controller.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       user,
		})
	}
}
