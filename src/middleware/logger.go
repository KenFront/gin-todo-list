package middleware

import "github.com/gin-gonic/gin"

func UseLogger(r *gin.Engine) {
	r.Use(gin.Logger())
}
