package controller_todos_test

import (
	"net/http"
	"testing"

	controller_todos "github.com/KenFront/gin-todo-list/src/controller/todos"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/stretchr/testify/assert"
)

type SuccessTodoAPIResponse struct {
	Data model.Todo `json:"data"`
}

func TestAddTodoSuccess(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()
	c.Set("userId", userId)

	fake := model.Add{
		Title:       "123",
		Description: "456",
	}
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	gormDB := mock.GetMockGorm(t)

	controller_todos.Add(controller_todos.AddProps{
		Db:           gormDB,
		GetNewTodoId: util.GetNewTodoId,
	})(c)

	var resBody SuccessTodoAPIResponse
	mock.GetResponseBody(res.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, fake.Title, resBody.Data.Title)
	assert.Equal(t, fake.Description, resBody.Data.Description)
}

func TestAddTodoFailBydMissingPayload(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()
	c.Set("userId", userId)

	fake := model.Add{
		Description: "456",
	}
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Equal(t, model.ERROR_CREATE_TODO_PAYLOAD_IS_INVALID, err.ErrorType)
	}()
	controller_todos.Add(controller_todos.AddProps{
		Db:           gormDB,
		GetNewTodoId: util.GetNewTodoId,
	})(c)
}
