package util

import (
	"errors"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserIdByToken(c *gin.Context) (uuid.UUID, error) {
	claim, err := GetToken(c)

	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(claim.UserId)
	if err != nil {
		return uuid.Nil, errors.New(string(model.ERROR_TOKEN_PARSE_FAILED))
	}

	return id, nil
}
