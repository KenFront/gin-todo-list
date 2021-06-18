package controller_test

import (
	"net/http"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
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

	fake := model.AddTodo{
		Title:       "123",
		Description: "456",
	}
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	gormDB := mock.GetMockGorm(t)

	controller.AddTodo(controller.AddTodoProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
		GetNewTodoId:     util.GetNewTodoId,
	})(c)

	var resBody SuccessTodoAPIResponse
	mock.GetResponseBody(res.Body.Bytes(), &resBody)

	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, resBody.Data.Title, fake.Title)
	assert.Equal(t, resBody.Data.Description, fake.Description)
}

func TestAddTodoFailBydMissingPayload(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()

	fake := model.AddTodo{
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
		assert.Equal(t, err.StatusCode, http.StatusBadRequest)
		assert.Equal(t, err.ErrorType, model.ERROR_CREATE_TODO_PAYLOAD_IS_INVALID)
	}()
	controller.AddTodo(controller.AddTodoProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
		GetNewTodoId:     util.GetNewTodoId,
	})(c)
}
