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

type SuccessTodosAPIResponse struct {
	Data []model.Todo `json:"data"`
}

func TestGetTodosTodoHanlderSuccess(t *testing.T) {
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

	resForList := mock.GetResponse()
	cForList := mock.GetGinContext(resForList)
	cForList.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	cForList.Request = &http.Request{
		Header: make(http.Header),
	}

	controller.GetTodos(controller.GetTodosProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(cForList)

	var resBody SuccessTodosAPIResponse
	mock.GetResponseBody(resForList.Body.Bytes(), &resBody)

	assert.Equal(t, resForList.Code, http.StatusOK)
	assert.Equal(t, resBody.Data[0].ID, todoId)
	assert.Equal(t, resBody.Data[0].Title, fake.Title)
	assert.Equal(t, resBody.Data[0].Description, fake.Description)
	assert.Equal(t, len(resBody.Data), 1)
}
