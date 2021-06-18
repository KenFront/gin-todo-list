package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var payload model.AddUser

	if err := c.ShouldBindJSON(&payload); err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_CREATE_USER_PAYLOAD_IS_INVALID,
			Error:      err,
		})
	}

	hashedPassword, err := util.HashPassword(payload.Password)
	if err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_HASH_PASSWORD_FAILD,
			Error:      err,
		})
	}

	id := util.GetNewUserId()

	user := model.User{
		ID:       id,
		Name:     payload.Name,
		Account:  payload.Account,
		Password: hashedPassword,
		Email:    payload.Email,
	}

	if err := config.GetDB().Create(&user).Error; err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_GET_CREATED_USER_FAILED,
			Error:      err,
		})
	}
	if err := config.GetDB().First(&user, "id = ?", id).Error; err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_CREATE_USER_FAILED,
			Error:      err,
		})
	}

	util.ApiOnSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       user,
	})
}
