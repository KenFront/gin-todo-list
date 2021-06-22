package controller_todos_test

import (
	"net/http"
	"testing"

	controller_todos "github.com/KenFront/gin-todo-list/src/controller/todos"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTodoByIdSuccess(t *testing.T) {
	resForAdd := mock.GetResponse()
	cForAdd := mock.GetGinContext(resForAdd)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()
	gormDB := mock.GetMockGorm(t)
	cForAdd.Set("userId", userId)

	fake := model.AddTodo{
		Title:       "123",
		Description: "456",
	}

	cForAdd.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	controller_todos.Add(controller_todos.AddProps{
		Db:           gormDB,
		GetNewTodoId: mock.UtilGetNewTodoId(todoId),
	})(cForAdd)

	assert.Equal(t, http.StatusOK, resForAdd.Code)

	resForGetById := mock.GetResponse()
	cForGetById := mock.GetGinContext(resForGetById)

	cForGetById.Set("userId", userId)

	cForGetById.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	cForGetById.Request = &http.Request{
		Header: make(http.Header),
	}

	controller_todos.GetById(controller_todos.GetByIdProps{
		Db: gormDB,
	})(cForGetById)

	var resBody SuccessTodoAPIResponse
	mock.GetResponseBody(resForGetById.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, resForGetById.Code)
	assert.Equal(t, todoId, resBody.Data.ID)
	assert.Equal(t, fake.Title, resBody.Data.Title)
	assert.Equal(t, fake.Description, resBody.Data.Description)
}

func TestGetTodoByIdFailByIdValidatingError(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()
	c.Set("userId", userId)

	c.Params = []gin.Param{
		{Key: "todoId", Value: "123"},
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
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Equal(t, model.ERROR_DELETE_TODO_PATH_FAILED, err.ErrorType)
	}()

	controller_todos.GetById(controller_todos.GetByIdProps{
		Db: gormDB,
	})(c)
}

func TestGetTodoByIdFailByNotExist(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()
	c.Set("userId", userId)

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

	controller_todos.GetById(controller_todos.GetByIdProps{
		Db: gormDB,
	})(c)
}
