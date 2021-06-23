package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	ApiOnSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       "Server is working",
	})
}
