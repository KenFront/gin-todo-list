package controller_users

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteProps struct {
	Db *gorm.DB
}

func Delete(p DeleteProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		userId, err := controller.GetUserId(c)
		if err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      err,
			})
		}

		if err := p.Db.First(&user, "id = ?", userId).Error; err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_DELETE_USER_NOT_EXIST,
				Error:      err,
			})
		}

		if err := p.Db.Delete(&user, "id = ?", userId).Error; err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_DELETE_USER_FAILED,
				Error:      err,
			})
		}

		controller.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       user,
		})
	}
}
