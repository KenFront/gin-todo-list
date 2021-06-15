package util

import (
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
)

func ApiOnSuccess(c *gin.Context, res *model.ApiSuccess) {
	c.JSON(res.StatusCode, gin.H{
		"data": res.Data,
	})
}

func ApiOnError(res *model.ApiError) {
	panic(res)
}
