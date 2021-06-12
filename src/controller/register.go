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
		panic(&util.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  err.Error(),
		})
	}

	hashedPassword, err := util.HashPassword(payload.Password)
	if err != nil {
		panic(&util.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  err.Error(),
		})
	}

	id, err := uuid.NewUUID()
	if err != nil {
		panic(&util.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  err.Error(),
		})
	}

	user := model.User{
		ID:       id,
		Name:     payload.Name,
		Account:  payload.Account,
		Password: hashedPassword,
		Email:    payload.Email,
	}

	createActionResult := config.GetDB().Create(&user)
	createdDataResult := config.GetDB().First(&user, "id = ?", id)

	switch {
	case createActionResult.Error != nil:
		panic(&util.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  createActionResult.Error.Error(),
		})
	case createdDataResult.Error != nil:
		panic(&util.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  createdDataResult.Error.Error(),
		})
	default:
		user.Password = "******"
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}
