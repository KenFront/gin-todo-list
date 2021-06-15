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

	if config.GetDB().Find(&todos, "user_id = ?", userId).Error != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_GET_TODOS_FAILED,
		})
	}

	util.ApiSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       todos,
	})
}
