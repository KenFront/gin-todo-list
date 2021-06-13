package util

import (
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
)

func ApiSuccess(c *gin.Context, res *model.ApiSuccess) {
	c.JSON(res.StatusCode, gin.H{
		"data": res.Data,
	})

}
