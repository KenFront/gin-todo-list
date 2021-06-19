package controller_users

import (
	"errors"
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeleteProps struct {
	Db               *gorm.DB
	GetUserIdByToken func(c *gin.Context) (uuid.UUID, error)
}

func DeleteById(p DeleteProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		id := c.Param("userId")

		userIdByContext, isExist := c.Get("userId")
		if !isExist {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      errors.New(string(model.ERROR_SIGN_IN_FAILED)),
			})
		}

		userId := userIdByContext.(uuid.UUID)

		if id != userId.String() {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_NOT_DELETE_SELF,
				Error:      errors.New(string(model.ERROR_NOT_DELETE_SELF)),
			})
		}

		if err := p.Db.First(&user, "id = ?", id).Error; err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_DELETE_USER_NOT_EXIST,
				Error:      err,
			})
		}

		if err := p.Db.Delete(&user).Error; err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_DELETE_USER_FAILED,
				Error:      err,
			})
		}

		util.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       user,
		})
	}
}
