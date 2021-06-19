package controller_todos

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GetByIdProps struct {
	Db               *gorm.DB
	GetUserIdByToken func(c *gin.Context) (uuid.UUID, error)
}

func GetById(p GetByIdProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo model.Todo
		id := c.Param("todoId")

		userId, err := p.GetUserIdByToken(c)
		if err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      err,
			})
		}

		if err := p.Db.First(&todo, "id = ? AND user_id = ?", id, userId).Error; err != nil {
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GET_TODO_BY_ID_FAILED,
				Error:      err,
			})
		}

		util.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       todo,
		})
	}
}