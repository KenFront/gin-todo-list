package controller_test

import (
	"net/http"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTodoHanlderSuccess(t *testing.T) {
	resForAdd := mock.GetResponse()
	cForAdd := mock.GetGinContext(resForAdd)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()
	gormDB := mock.GetMockGorm(t)
	fake := model.AddTodo{
		Title:       "123",
		Description: "456",
	}

	cForAdd.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	controller.AddTodo(controller.AddTodoProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
		GetNewTodoId:     mock.UtilGetNewTodoId(todoId),
	})(cForAdd)

	assert.Equal(t, resForAdd.Code, http.StatusOK)

	resForDelete := mock.GetResponse()
	cForDelete := mock.GetGinContext(resForDelete)
	cForDelete.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	cForDelete.Request = &http.Request{
		Header: make(http.Header),
	}

	controller.DeleteTodoById(controller.DeleteTodoProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(cForDelete)

	var resBody SuccessTodoAPIResponse
	mock.GetResponseBody(resForDelete.Body.Bytes(), &resBody)

	assert.Equal(t, resForDelete.Code, http.StatusOK)
	assert.Equal(t, resBody.Data.ID, todoId)
	assert.Equal(t, resBody.Data.Title, fake.Title)
	assert.Equal(t, resBody.Data.Description, fake.Description)
}

func TestDeleteTodoHanlderFailByNotExist(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()

	c.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
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
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(c)
}
