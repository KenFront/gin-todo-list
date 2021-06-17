package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddTodoProps struct {
	Db               *gorm.DB
	GetUserIdByToken func(c *gin.Context) (uuid.UUID, error)
	GetNewTodoId     func() uuid.UUID
}

func AddTodo(p AddTodoProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.AddTodo
		if err := c.ShouldBindJSON(&payload); err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_CREATE_TODO_PAYLOAD_IS_INVALID,
				Error:      err,
			})
		}

		id := p.GetNewTodoId()

		userId, err := p.GetUserIdByToken(c)
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
