package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UtilGetUserIdByToken(id uuid.UUID) func(c *gin.Context) (uuid.UUID, error) {
	return func(c *gin.Context) (uuid.UUID, error) {
		return id, nil
	}
}

func UtilGetNewTodoId(id uuid.UUID) func() uuid.UUID {
	return func() uuid.UUID {
		return id
	}
}
