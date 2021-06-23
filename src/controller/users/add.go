package controller_users

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddProps struct {
	Db           *gorm.DB
	GetNewUserId func() uuid.UUID
}

func Add(p AddProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.AddUser

		if err := c.ShouldBindJSON(&payload); err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_CREATE_USER_PAYLOAD_IS_INVALID,
				Error:      err,
			})
		}

		hashedPassword, err := util.HashPassword(payload.Password)
		if err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_HASH_PASSWORD_FAILD,
				Error:      err,
			})
		}

		id := p.GetNewUserId()

		user := model.User{
			ID:       id,
			Name:     payload.Name,
			Account:  payload.Account,
			Password: hashedPassword,
			Email:    payload.Email,
		}

		if err := p.Db.Create(&user).Error; err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_CREATE_USER_FAILED,
				Error:      err,
			})
		}

		if err := p.Db.First(&user, "id = ?", id).Error; err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GET_CREATED_USER_FAILED,
				Error:      err,
			})
		}

		controller.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       user,
		})
	}
}
