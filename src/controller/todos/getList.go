package controller_todos

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetListProps struct {
	Db *gorm.DB
}

func GetList(p GetListProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos []model.Todo

		userId, err := controller.GetUserId(c)
		if err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      err,
			})
		}

		if err := p.Db.Order("updated_at desc").Find(&todos, "user_id = ?", userId).Error; err != nil {
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_GET_TODOS_FAILED,
				Error:      err,
			})
		}

		controller.ApiOnSuccess(c, &model.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       todos,
		})
	}
}
