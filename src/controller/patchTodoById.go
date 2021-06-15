package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func PatchTodoById(c *gin.Context) {
	var payload model.PatchTodo
	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_PATCH_TODO_PAYLOAD_IS_INVALID,
		})
	}

	if (model.PatchTodo{} == payload) {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_NO_VALUE_IN_PATCH_TODO_PAYLOAD,
		})
	}

	var todo model.Todo
	id := c.Params.ByName("todoId")
	userId := util.GetUserId(c)

	if config.GetDB().Model(&todo).Updates(model.Todo{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
	}).Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_PATCH_TODO_FAILED,
		})
	}

	if config.GetDB().First(&todo, "id = ? AND user_id = ?", id, userId).Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_GET_PATCHED_TODO_FAILED,
		})
	}
	util.ApiSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       todo,
	})
}
