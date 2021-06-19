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

func TestPatchTodoSuccess(t *testing.T) {
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

	resForPatch := mock.GetResponse()
	cForPatch := mock.GetGinContext(resForPatch)
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

	controller.PatchTodoById(controller.PatchTodoProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(cForPatch)

	var resBody SuccessTodoAPIResponse
	mock.GetResponseBody(resForPatch.Body.Bytes(), &resBody)

	assert.Equal(t, resForPatch.Code, http.StatusOK)
	assert.Equal(t, resBody.Data.ID, todoId)
	assert.Equal(t, resBody.Data.Title, fakePatch.Title)
	assert.Equal(t, resBody.Data.Description, fake.Description)
}

func TestPatchTodoFailByNotExist(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()

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
		assert.Equal(t, err.StatusCode, http.StatusServiceUnavailable)
		assert.Equal(t, err.ErrorType, model.ERROR_GET_PATCHED_TODO_FAILED)
	}()

	controller.PatchTodoById(controller.PatchTodoProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(c)
}

func TestPatchTodoFailedByNoNeededPayload(t *testing.T) {
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

	resForPatch := mock.GetResponse()
	cForPatch := mock.GetGinContext(resForPatch)

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
		assert.Equal(t, err.StatusCode, http.StatusBadRequest)
		assert.Equal(t, err.ErrorType, model.ERROR_NO_VALUE_IN_PATCH_TODO_PAYLOAD)
	}()

	controller.PatchTodoById(controller.PatchTodoProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(cForPatch)
}
