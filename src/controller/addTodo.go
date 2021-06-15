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
		if c.ShouldBindJSON(&payload) != nil {
			panic(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_CREATE_TODO_PAYLOAD_IS_INVALID,
			})
		}

		id, err := uuid.NewUUID()
		if err != nil {
			panic(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GENERATE_ID_FAILED,
			})
		}

		userId, err := p.GetUserId(c)
		if err != nil {
			panic(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_NOT_FOUNT_THIS_USER,
			})
		}

		todo := model.Todo{
			ID:          id,
			Title:       payload.Title,
			Description: payload.Description,
			UserId:      userId,
		}

		if err := p.Db.Create(&todo).Error; err != nil {
			panic(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_CREATE_TODO_FAILED,
			})
		}

		if err := p.Db.First(&todo, "id = ?", id).Error; err != nil {
			panic(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GET_CREATED_TODO_FAILED,
			})
		}

		util.ApiSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       todo,
		})
	}
}
