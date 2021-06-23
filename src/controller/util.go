package controller

import (
	"errors"
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	userIdKey = "userId"
)

func SetUserId(c *gin.Context, id uuid.UUID) {
	c.Set(userIdKey, id)
}

func GetUserId(c *gin.Context) uuid.UUID {
	userId, isExist := c.Get(userIdKey)
	if !isExist {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_SIGN_IN_FAILED,
			Error:      errors.New(string(model.ERROR_SIGN_IN_FAILED)),
		})
	}

	id := userId.(uuid.UUID)

	return id
}
