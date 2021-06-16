package middleware

import (
	"errors"
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
		if e := c.Error(err.Error); e != nil {
			panic(e)
		}
		c.AbortWithStatusJSON(err.StatusCode, gin.H{
			"error": err.ErrorType,
		})
	case r != nil && util.IsSameType(r, errors.New("")):
		err := r.(error)
		if e := c.Error(err); e != nil {
			panic(e)
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": model.ERROR_UNKNOWN,
		})
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
