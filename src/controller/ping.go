package controller

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.Abort()
}
