package controller_todos

import (
	"errors"
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteProps struct {
	Db *gorm.DB
}

func DeleteById(p DeleteProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var uri model.TodoUri

		if err := c.ShouldBindUri(&uri); err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_DELETE_TODO_PATH_FAILED,
				Error:      errors.New(string(model.ERROR_DELETE_TODO_PATH_FAILED)),
			})
			return
		}
		userId, isExist := c.Get("userId")

		if !isExist {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      errors.New(string(model.ERROR_SIGN_IN_FAILED)),
			})
		}

		var todo model.Todo

		if err := p.Db.First(&todo, "id = ? AND user_id = ?", uri.TodoId, userId).Error; err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_DELETE_TODO_NOT_EXIST,
				Error:      err,
			})
		}

		if err := p.Db.Delete(&todo, "id = ? AND user_id = ?", uri.TodoId, userId).Error; err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_DELETE_TODO_FAILED,
				Error:      err,
			})
		}

		util.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       todo,
		})
	}
}
