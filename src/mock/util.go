package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UtilGetUserId(c *gin.Context) (uuid.UUID, error) {
	return uuid.Nil, nil
}
