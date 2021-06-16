package controller_test

import (
	"net/http"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTodoHanlderFailByNotExist(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	c.Params = []gin.Param{
		{Key: "todoId", Value: uuid.Nil.String()},
	}

	c.Request = &http.Request{
		Header: make(http.Header),
	}

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, err.StatusCode, http.StatusServiceUnavailable)
		assert.Equal(t, err.ErrorType, model.ERROR_DELETE_TODO_NOT_EXIST)
	}()

	controller.DeleteTodoById(controller.DeleteTodoProps{
		Db:        gormDB,
		GetUserId: mock.UtilGetUserId,
	})(c)
}
