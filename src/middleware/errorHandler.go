package middleware

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func catchError(c *gin.Context) {
	r := recover()
	switch {
	case util.IsSameType(r, &model.ApiError{}):
		err := r.(*model.ApiError)
		c.AbortWithStatusJSON(err.StatusCode, gin.H{
			"error": err.ErrorType,
		})
	case r != nil:
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer catchError(c)
		c.Next()
	}
}

func UseErrorHandler(r *gin.Engine) gin.IRoutes {
	return r.Use(errorHandler())
}
