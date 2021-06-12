package middleware

import "github.com/gin-gonic/gin"

func UseRecovery(r *gin.Engine) {
	r.Use(gin.Recovery())
}
