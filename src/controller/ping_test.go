package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controller "github.com/KenFront/gin-todo-list/src/controller"

	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	controller.Ping(c)

	var jsonResponse gin.H
	mock.GetResponseBody(res.Body.Bytes(), &jsonResponse)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, gin.H{
		"data": "Server is working",
	}, jsonResponse)
}
