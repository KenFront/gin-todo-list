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
			panic(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  err.Error(),
			})
		}

		id, err := uuid.NewUUID()
		if err != nil {
			panic(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  err.Error(),
			})
		}

		todo := model.Todo{
			ID:          id,
			Title:       payload.Title,
			Description: payload.Description,
			UserId:      p.GetUserId(c),
		}

		createActionResult := p.Db.Create(&todo)
		createdDataResult := p.Db.First(&todo, "id = ?", id)

		switch {
		case createActionResult.Error != nil:
			panic(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  createActionResult.Error.Error(),
			})
		case createdDataResult.Error != nil:
			panic(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  createdDataResult.Error.Error(),
			})
		default:
			util.ApiSuccess(c, &model.ApiSuccess{
				StatusCode: http.StatusOK,
				Data:       todo,
			})
		}
	}
}
