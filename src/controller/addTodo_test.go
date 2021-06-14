package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddTodoHanlderSuccess(t *testing.T) {
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)

	c.Request = &http.Request{
		Header: make(http.Header),
		Body: mock.GetRequsetBody(map[string]interface{}{
			"title":       "123",
			"description": "456",
		}),
	}

	gormDB := mock.GetMockGorm(t)
	AddTodo(model.AddTodoProps{
		Db:        gormDB,
		GetUserId: mock.GetUtilGetUserId,
	})(c)

	assert.Equal(t, res.Code, http.StatusOK)
}
