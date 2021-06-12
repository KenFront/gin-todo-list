package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	util.DeleteAuth(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign out successfully",
	})
}
