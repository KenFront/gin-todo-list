package middleware

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			r := recover()
			switch {
			case util.IsSameType(r, &util.ApiError{}):
				err := r.(*util.ApiError)
				c.AbortWithStatusJSON(err.StatusCode, gin.H{
					"error": err.ErrorType,
				})
			case r != nil:
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func UseErrorHandler(r *gin.Engine) gin.IRoutes {
	return r.Use(errorHandler())
}
