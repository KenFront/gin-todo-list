package middleware

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthGuardProps struct {
	Db               *gorm.DB
	GetUserIdByToken func(c *gin.Context) (uuid.UUID, error)
}

func AuthGuard(p AuthGuardProps) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := p.GetUserIdByToken(c)
		if err != nil {
			util.DeleteAuth(c)
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ErrorType(err.Error()),
				Error:      err,
			})
		}

		if err := p.Db.First(&model.User{}, "id = ?", id).Error; err != nil {
			util.DeleteAuth(c)
			controller.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      err,
			})
		}

		controller.SetUserId(c, id)

		c.Next()
	}
}

func UseAuthGuard(r *gin.RouterGroup) gin.IRoutes {
	db := config.GetDB()
	return r.Use(AuthGuard(AuthGuardProps{
		Db:               db,
		GetUserIdByToken: util.GetUserIdByToken,
	}))
}
