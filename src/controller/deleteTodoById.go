package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func DeleteTodoById(c *gin.Context) {
	var todo model.Todo
	id := c.Param("todoId")
	userId := util.GetUserId(c)
	result := config.GetDB().Delete(&todo, "id = ? AND user_id = ?", id, userId)

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted successfully.",
		})
	} else {
		panic(&util.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  result.Error.Error(),
		})
	}
}
