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

func TestGetTodoByidTodoSuccess(t *testing.T) {
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

	assert.Equal(t, http.StatusOK, resForAdd.Code)

	resForGetById := mock.GetResponse()
	cForGetById := mock.GetGinContext(resForGetById)
	cForGetById.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	cForGetById.Request = &http.Request{
		Header: make(http.Header),
	}

	controller.GetTodoById(controller.GetTodoByIdProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(cForGetById)

	var resBody SuccessTodoAPIResponse
	mock.GetResponseBody(resForGetById.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, resForGetById.Code)
	assert.Equal(t, todoId, resBody.Data.ID)
	assert.Equal(t, fake.Title, resBody.Data.Title)
	assert.Equal(t, fake.Description, resBody.Data.Description)
}

func TestGetTodoByidTodoFailByNotExist(t *testing.T) {
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
		assert.Equal(t, http.StatusServiceUnavailable, err.StatusCode)
		assert.Equal(t, model.ERROR_GET_TODO_BY_ID_FAILED, err.ErrorType)
	}()

	controller.GetTodoById(controller.GetTodoByIdProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(c)
}
