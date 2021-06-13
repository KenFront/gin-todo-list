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
			ErrorType:  err.Error(),
		})
	}

	if (model.PatchTodo{} == payload) {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  "No changed",
		})
	}

	var todo model.Todo
	id := c.Params.ByName("todoId")
	userId := util.GetUserId(c)

	if err := config.GetDB().First(&todo, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  err.Error(),
		})
	}

	result := config.GetDB().Model(&todo).Updates(model.Todo{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
	})

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
