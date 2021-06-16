package controller_test

import (
	"net/http"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTodoHanlderSuccess(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	c.Params = []gin.Param{
		{Key: "todoId", Value: uuid.Nil.String()},
	}

	c.Request = &http.Request{
		Header: make(http.Header),
	}

	gormDB := mock.GetMockGorm(t)

	controller.DeleteTodoById(controller.DeleteTodoProps{
		Db:        gormDB,
		GetUserId: mock.UtilGetUserId,
	})(c)

	assert.Equal(t, res.Code, http.StatusOK)
}
