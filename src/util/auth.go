package util

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserId(c *gin.Context) uuid.UUID {
	claim, err := CheckAuth(c)
	if err != nil {
		DeleteAuth(c)
		panic(&model.ApiError{
			StatusCode: http.StatusUnauthorized,
			ErrorType:  err.Error(),
		})
	}

	id, err := uuid.Parse(claim.UserId)
	if err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusUnauthorized,
			ErrorType:  err.Error(),
		})
	}
	return id
}
