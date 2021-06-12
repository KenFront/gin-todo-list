package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
)

func PatchTodoById(c *gin.Context) {
	var payload model.PatchTodo
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if (model.PatchTodo{} == payload) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "No changed",
		})
	}

	var todo model.Todo
	id := c.Params.ByName("todoId")

	if err := config.GetDB().Where("id = ?", id).First(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	result := config.GetDB().Model(&todo).Updates(model.Todo{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
	})

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
