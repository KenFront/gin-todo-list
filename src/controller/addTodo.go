package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var payload model.AddTodo

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	todo := model.Todo{Title: payload.Title, Description: payload.Description}

	createActionResult := config.GetDB().Create(&todo)
	createdDataResult := config.GetDB().First(&todo)

	switch {
	case createActionResult.Error != nil:
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": createActionResult.Error,
		})
	case createdDataResult.Error != nil:
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": createdDataResult.Error,
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
		})
	}
}
