package model

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddTodoProps struct {
	Db        *gorm.DB
	GetUserId func(c *gin.Context) uuid.UUID
}
