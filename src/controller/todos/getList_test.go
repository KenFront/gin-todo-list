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

type SuccessTodosAPIResponse struct {
	Data []model.Todo `json:"data"`
}

func TestGetTodosTodoSuccess(t *testing.T) {
	resForAdd := mock.GetResponse()
	cForAdd := mock.GetGinContext(resForAdd)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()
	gormDB := mock.GetMockGorm(t)

	cForAdd.Set("userId", userId)

	fake := model.Add{
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

	resForList := mock.GetResponse()
	cForList := mock.GetGinContext(resForList)

	cForList.Set("userId", userId)

	cForList.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	cForList.Request = &http.Request{
		Header: make(http.Header),
	}

	controller_todos.GetList(controller_todos.GetListProps{
		Db: gormDB,
	})(cForList)

	var resBody SuccessTodosAPIResponse
	mock.GetResponseBody(resForList.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, resForList.Code)
	assert.Equal(t, todoId, resBody.Data[0].ID)
	assert.Equal(t, fake.Title, resBody.Data[0].Title)
	assert.Equal(t, fake.Description, resBody.Data[0].Description)
	assert.Equal(t, 1, len(resBody.Data))
}
