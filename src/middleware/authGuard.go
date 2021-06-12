package middleware

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func authGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := util.CheckAuth(c)
		if err != nil {
			util.DeleteAuth(c)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}
		c.Next()
	}
}

func UseAuthGuard(r *gin.RouterGroup) gin.IRoutes {
	return r.Use(authGuard())
}
