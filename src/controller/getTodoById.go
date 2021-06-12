package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
)

func GetTodoById(c *gin.Context) {
	var todo model.Todo
	id := c.Param("todoId")

	result := config.GetDB().First(&todo, "id = ?", id)

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"error": result.Error.Error(),
		})
	}
}
