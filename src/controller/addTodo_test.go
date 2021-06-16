package controller

import (
	"net/http"
	"testing"

	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/stretchr/testify/assert"
)

type addTodoSuccessResponse struct {
	Data model.Todo `json:"data"`
}

func TestAddTodoHanlderSuccess(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)

	fake := model.AddTodo{
		Title:       "123",
		Description: "456",
	}
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	gormDB := mock.GetMockGorm(t)

	AddTodo(model.AddTodoProps{
		Db:        gormDB,
		GetUserId: mock.UtilGetUserId,
	})(c)

	var resBody addTodoSuccessResponse
	mock.GetResponseBody(res.Body.Bytes(), &resBody)

	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, resBody.Data.Title, fake.Title)
	assert.Equal(t, resBody.Data.Description, fake.Description)
}

func TestAddTodoHanlderFailBydMissingPayloa(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)

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
	AddTodo(model.AddTodoProps{
		Db:        gormDB,
		GetUserId: mock.UtilGetUserId,
	})(c)
}
