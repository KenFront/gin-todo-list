package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var payload model.SignIn

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  err.Error(),
		})
	}

	user := model.User{
		Account:  payload.Account,
		Password: payload.Password,
	}

	result := config.GetDB().First(&user, "account = ?", payload.Account)

	if result.Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  result.Error.Error(),
		})
	}

	if !util.CheckPasswordHash(payload.Password, user.Password) {
		panic(&model.ApiError{
			StatusCode: http.StatusUnauthorized,
			ErrorType:  "Sign in fail",
		})
	}

	err := util.SetAuth(c, user.ID.String())
	if err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  err.Error(),
		})
	}
	util.ApiSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       "Sign in successfully",
	})
}
