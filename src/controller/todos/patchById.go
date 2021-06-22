package controller_todos

import (
	"errors"
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PatchProps struct {
	Db *gorm.DB
}

func PatchById(p PatchProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var uri model.TodoUri

		if err := c.ShouldBindUri(&uri); err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_PATCH_TODO_PATH_FAILED,
				Error:      errors.New(string(model.ERROR_DELETE_TODO_PATH_FAILED)),
			})
			return
		}

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

		userId, isExist := c.Get("userId")
		if !isExist {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      errors.New(string(model.ERROR_SIGN_IN_FAILED)),
			})
		}

		var todo model.Todo

		if err := p.Db.Model(&todo).Where("id = ? AND user_id = ?", uri.TodoId, userId).Updates(model.Todo{
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

		if err := p.Db.First(&todo, "id = ? AND user_id = ?", uri.TodoId, userId).Error; err != nil {
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
}
