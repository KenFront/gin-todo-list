package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []model.Todo

	result := config.GetDB().Find(&todos)

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": todos,
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": result.Error.Error(),
		})
	}
}
