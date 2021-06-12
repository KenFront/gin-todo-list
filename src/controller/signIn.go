package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	util.SetAuth(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign in successfully",
	})
}
