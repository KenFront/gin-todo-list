package controller

import (
	"net/http"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	util.DeleteAuth(c)
	util.ApiOnSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       "Sign out successfully",
	})
}
