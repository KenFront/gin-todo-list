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

func TestPatchTodoSuccess(t *testing.T) {
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

	resForPatch := mock.GetResponse()
	cForPatch := mock.GetGinContext(resForPatch)

	cForPatch.Set("userId", userId)

	fakePatch := model.PatchTodo{
		Title: "patched",
	}
	cForPatch.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	cForPatch.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fakePatch),
	}

	controller_todos.PatchById(controller_todos.PatchProps{
		Db: gormDB,
	})(cForPatch)

	var resBody SuccessTodoAPIResponse
	mock.GetResponseBody(resForPatch.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, resForPatch.Code)
	assert.Equal(t, todoId, resBody.Data.ID)
	assert.Equal(t, fakePatch.Title, resBody.Data.Title)
	assert.Equal(t, fake.Description, resBody.Data.Description)
}

func TestPatchTodoFailByNotExist(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()

	c.Set("userId", userId)

	c.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	fake := model.PatchTodo{
		Title: "patched",
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
		assert.Equal(t, http.StatusServiceUnavailable, err.StatusCode)
		assert.Equal(t, model.ERROR_GET_PATCHED_TODO_FAILED, err.ErrorType)
	}()

	controller_todos.PatchById(controller_todos.PatchProps{
		Db: gormDB,
	})(c)
}

func TestPatchTodoFailedByNoNeededPayload(t *testing.T) {
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

	resForPatch := mock.GetResponse()
	cForPatch := mock.GetGinContext(resForPatch)

	cForPatch.Set("userId", userId)

	cForPatch.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	var fakeForPatch = map[string]string{
		"123": "456",
	}

	cForPatch.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fakeForPatch),
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Equal(t, model.ERROR_NO_VALUE_IN_PATCH_TODO_PAYLOAD, err.ErrorType)
	}()

	controller_todos.PatchById(controller_todos.PatchProps{
		Db: gormDB,
	})(cForPatch)
}
