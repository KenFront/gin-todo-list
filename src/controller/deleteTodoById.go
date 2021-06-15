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

	userId, err := util.GetUserId(c)
	if err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_NOT_FOUNT_THIS_USER,
		})
	}

	if config.GetDB().Delete(&todo, "id = ? AND user_id = ?", id, userId).Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_DELETE_TODO_FAILED,
		})
	}

	util.ApiSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       "Deleted successfully.",
	})
}
