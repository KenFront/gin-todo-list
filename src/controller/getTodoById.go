package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func GetTodoById(c *gin.Context) {
	var todo model.Todo
	id := c.Param("todoId")
	userId := util.GetUserId(c)

	result := config.GetDB().First(&todo, "id = ? AND user_id = ?", id, userId)

	if result.Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  result.Error.Error(),
		})
	}
	util.ApiSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       todo,
	})
}
