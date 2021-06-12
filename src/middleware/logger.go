package middleware

import "github.com/gin-gonic/gin"

func UseLogger(r *gin.Engine) gin.IRoutes {
	return r.Use(gin.Logger())
}
