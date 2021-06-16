package middleware

import (
	"errors"
	"fmt"
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
		if e := c.Error(err.Error); e.Err != nil {
			fmt.Println(e)
		}
		c.AbortWithStatusJSON(err.StatusCode, gin.H{
			"error": err.ErrorType,
		})
	case util.IsSameType(r, errors.New("")):
		err := r.(error)
		fmt.Println(r)
		if e := c.Error(err); e.Err != nil {
			fmt.Println(e)
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
