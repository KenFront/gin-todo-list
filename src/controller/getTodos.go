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

	userId, err := util.GetUserId(c)
	if err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_NOT_FOUNT_THIS_USER,
			Error:      err,
		})
	}

	if err := config.GetDB().Find(&todos, "user_id = ?", userId).Error; err != nil {
		util.ApiOnError(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_GET_TODOS_FAILED,
			Error:      err,
		})
	}

	util.ApiOnSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       todos,
	})
}
