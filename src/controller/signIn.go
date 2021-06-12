package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var payload model.SignIn

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user := model.User{
		Account:  payload.Account,
		Password: payload.Password,
	}

	result := config.GetDB().First(&user, "account = ?", payload.Account)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"error": result.Error.Error(),
		})
	}

	if util.CheckPasswordHash(payload.Password, user.Password) {
		err := util.SetAuth(c, user.ID.String())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Sign in successfully",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Sign in fail",
		})
	}

}
