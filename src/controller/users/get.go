package controller_users

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GetProps struct {
	Db               *gorm.DB
	GetUserIdByToken func(c *gin.Context) (uuid.UUID, error)
}

func Get(p GetProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		id, err := p.GetUserIdByToken(c)
		if err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      err,
			})
		}

		if err := p.Db.First(&user, "id = ?", id).Error; err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GET_USER_FAILED,
				Error:      err,
			})
		}

		util.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       user,
		})
	}
}
