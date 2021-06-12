package middleware

import "github.com/gin-gonic/gin"

func UseRecovery(r *gin.Engine) gin.IRoutes {
	return r.Use(gin.Recovery())
}
