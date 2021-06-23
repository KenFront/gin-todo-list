package controller_todos

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetListProps struct {
	Db *gorm.DB
}

func GetList(p GetListProps) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos []model.Todo

		userId := controller.GetUserId(c)

		if err := p.Db.Find(&todos, "user_id = ?", userId).Error; err != nil {
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
}
