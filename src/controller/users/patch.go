package controller_users

import (
	"errors"
	"net/http"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PatchProps struct {
	Db *gorm.DB
}

func Patch(p PatchProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.PatchUser
		if err := c.ShouldBindJSON(&payload); err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_PATCH_USER_PAYLOAD_IS_INVALID,
				Error:      err,
			})
		}

		if (model.PatchUser{} == payload) {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_NO_VALUE_IN_PATCH_USER_PAYLOAD,
				Error:      errors.New(string(model.ERROR_NO_VALUE_IN_PATCH_USER_PAYLOAD)),
			})
		}

		id := controller.GetUserId(c)

		var user model.User

		if err := p.Db.Model(&user).Where("id = ?", id).Updates(model.User{
			Name:     payload.Name,
			Account:  payload.Account,
			Password: payload.Password,
			Email:    payload.Email,
			Status:   payload.Status,
		}).Error; err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_PATCH_USER_FAILED,
				Error:      err,
			})
		}

		if err := p.Db.First(&user, "id = ?", id).Error; err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GET_PATCHED_USER_FAILED,
				Error:      err,
			})
		}

		controller.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       user,
		})
	}
}
