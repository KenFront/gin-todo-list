package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Regiser(c *gin.Context) {
	var payload model.AddUser

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	hashedPassword, err := util.HashPassword(payload.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	id, err := uuid.NewUUID()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user := model.User{
		ID:       id,
		Name:     payload.Name,
		Account:  payload.Account,
		Password: hashedPassword,
		Email:    payload.Email,
	}

	createActionResult := config.GetDB().Create(&user)
	createdDataResult := config.GetDB().First(&user, "id = ?", id)

	switch {
	case createActionResult.Error != nil:
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"error": createActionResult.Error,
		})
	case createdDataResult.Error != nil:
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"error": createdDataResult.Error,
		})
	default:
		user.Password = "******"
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}
