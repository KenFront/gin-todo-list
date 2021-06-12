package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddTodo(c *gin.Context) {
	var payload model.AddTodo

	if err := c.ShouldBindJSON(&payload); err != nil {
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

	todo := model.Todo{ID: id, Title: payload.Title, Description: payload.Description}

	createActionResult := config.GetDB().Create(&todo)
	createdDataResult := config.GetDB().First(&todo, "id = ?", id)

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
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
		})
	}
}
