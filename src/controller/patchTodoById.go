package controller

import (
	"errors"
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func PatchTodoById(c *gin.Context) {
	var payload model.PatchTodo
	if err := c.ShouldBindJSON(&payload); err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_PATCH_TODO_PAYLOAD_IS_INVALID,
			Error:      err,
		})
	}

	if (model.PatchTodo{} == payload) {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_NO_VALUE_IN_PATCH_TODO_PAYLOAD,
			Error:      errors.New(string(model.ERROR_NO_VALUE_IN_PATCH_TODO_PAYLOAD)),
		})
	}

	var todo model.Todo
	id := c.Params.ByName("todoId")

	userId, err := util.GetUserIdByToken(c)
	if err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_SIGN_IN_FAILED,
			Error:      err,
		})
	}

	if err := config.GetDB().Model(&todo).Updates(model.Todo{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
	}).Error; err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_PATCH_TODO_FAILED,
			Error:      err,
		})
	}

	if err := config.GetDB().First(&todo, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_GET_PATCHED_TODO_FAILED,
			Error:      err,
		})
	}
	util.ApiOnSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       todo,
	})
}
