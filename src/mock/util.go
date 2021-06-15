package mock

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUtilGetUserId(c *gin.Context) uuid.UUID {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(&model.ApiError{
			StatusCode: http.StatusServiceUnavailable,
			ErrorType:  model.ERROR_GENERATE_ID_FAILED,
		})
	}
	return id
}
