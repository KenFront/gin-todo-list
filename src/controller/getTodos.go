package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []model.Todo
	userId := util.GetUserId(c)

	result := config.GetDB().Find(&todos, "user_id = ?", userId)

	if result.Error != nil {
		panic(&util.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  result.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}
