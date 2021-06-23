package controller_auth

import (
	"errors"
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var payload model.SignIn

	if err := c.ShouldBindJSON(&payload); err != nil {
		controller.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_SIGN_IN_PAYLOAD_IS_INVALID,
			Error:      err,
		})
	}

	user := model.User{
		Account:  payload.Account,
		Password: payload.Password,
	}

	if err := config.GetDB().First(&user, "account = ?", payload.Account).Error; err != nil {
		controller.ApiOnError(&model.ApiError{
			StatusCode: http.StatusUnauthorized,
			ErrorType:  model.ERROR_SIGN_IN_FAILED,
			Error:      err,
		})
	}

	if !util.CheckPasswordHash(payload.Password, user.Password) {
		controller.ApiOnError(&model.ApiError{
			StatusCode: http.StatusUnauthorized,
			ErrorType:  model.ERROR_SIGN_IN_FAILED,
			Error:      errors.New(string(model.ERROR_SIGN_IN_FAILED)),
		})
	}

	if user.Status != model.USER_ACTIVE {
		controller.ApiOnError(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_USER_IS_NOT_ACTIVE,
			Error:      errors.New(string(model.ERROR_USER_IS_NOT_ACTIVE)),
		})
	}

	err := util.SetAuth(c, user.ID.String())
	if err != nil {
		controller.ApiOnError(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_SIGN_IN_FAILED,
			Error:      err,
		})
	}

	controller.ApiOnSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       "Sign in successfully",
	})
}
