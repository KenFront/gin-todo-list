package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddTodo(p model.AddTodoProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.AddTodo
		if err := c.ShouldBindJSON(&payload); err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_CREATE_TODO_PAYLOAD_IS_INVALID,
				Error:      err,
			})
		}

		id, err := uuid.NewUUID()
		if err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GENERATE_ID_FAILED,
				Error:      err,
			})
		}

		userId, err := p.GetUserId(c)
		if err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_NOT_FOUNT_THIS_USER,
				Error:      err,
			})
		}

		todo := model.Todo{
			ID:          id,
			Title:       payload.Title,
			Description: payload.Description,
			UserId:      userId,
		}

		if err := p.Db.Create(&todo).Error; err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_CREATE_TODO_FAILED,
				Error:      err,
			})
		}

		if err := p.Db.First(&todo, "id = ?", id).Error; err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GET_CREATED_TODO_FAILED,
				Error:      err,
			})
		}

		util.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       todo,
		})
	}
}
