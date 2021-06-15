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

	if config.GetDB().First(&todo, "id = ? AND user_id = ?", id, userId).Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_GET_TODO_BY_ID_FAILED,
		})
	}

	util.ApiSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       todo,
	})
}
