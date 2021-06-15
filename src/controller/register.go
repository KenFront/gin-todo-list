package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Regiser(c *gin.Context) {
	var payload model.AddUser

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_CREATE_USER_PAYLOAD_IS_INVALID,
		})
	}

	hashedPassword, err := util.HashPassword(payload.Password)
	if err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_HASH_PASSWORD_FAILD,
		})
	}

	id, err := uuid.NewUUID()
	if err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_GENERATE_ID_FAILED,
		})
	}

	user := model.User{
		ID:       id,
		Name:     payload.Name,
		Account:  payload.Account,
		Password: hashedPassword,
		Email:    payload.Email,
	}

	if config.GetDB().Create(&user).Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_GET_CREATED_USER_FAILED,
		})
	}
	if config.GetDB().First(&user, "id = ?", id).Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_CREATE_USER_FAILED,
		})
	}

	util.ApiSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       user,
	})
}
