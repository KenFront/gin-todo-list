package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUtilGetUserId(c *gin.Context) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
