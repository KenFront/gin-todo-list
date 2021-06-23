package controller_users_test

import (
	"net/http"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
	controller_users "github.com/KenFront/gin-todo-list/src/controller/users"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPatchUserSuccess(t *testing.T) {
	resForAdd := mock.GetResponse()
	userId := util.GetNewTodoId()
	cRorAdd := mock.GetGinContext(resForAdd)

	gormDB := mock.GetMockGorm(t)

	fake := mock.GetMockUser()

	cRorAdd.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	controller_users.Add(controller_users.AddProps{
		Db:           gormDB,
		GetNewUserId: func() uuid.UUID { return userId },
	})(cRorAdd)

	assert.Equal(t, http.StatusOK, resForAdd.Code)

	resForPatch := mock.GetResponse()
	cRorPatch := mock.GetGinContext(resForPatch)
	controller.SetUserId(cRorPatch, userId)

	payload := model.PatchUser{
		Name: "123",
	}

	cRorPatch.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(payload),
	}

	controller_users.Patch(controller_users.PatchProps{
		Db: gormDB,
	})(cRorPatch)

	var resBody SuccessUserAPIResponse
	mock.GetResponseBody(resForPatch.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, resForPatch.Code)
	assert.Equal(t, payload.Name, resBody.Data.Name)
	assert.Equal(t, fake.Email, resBody.Data.Email)
	assert.Equal(t, fake.Account, resBody.Data.Account)
}

func TestPatchUserFailedByNotSignIn(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)

	payload := model.PatchUser{
		Name: "123",
	}

	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(payload),
	}

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Equal(t, model.ERROR_SIGN_IN_FAILED, err.ErrorType)
	}()

	controller_users.Patch(controller_users.PatchProps{
		Db: gormDB,
	})(c)
}

func TestPatchUserFailedByNoNeededPayload(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)

	payload := model.PatchUser{}

	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(payload),
	}

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Equal(t, model.ERROR_NO_VALUE_IN_PATCH_USER_PAYLOAD, err.ErrorType)
	}()

	controller_users.Patch(controller_users.PatchProps{
		Db: gormDB,
	})(c)
}

func TestPatchUserFailedByNoUserId(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)

	payload := model.PatchUser{
		Name: "123",
	}

	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(payload),
	}

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Equal(t, model.ERROR_SIGN_IN_FAILED, err.ErrorType)
	}()

	controller_users.Patch(controller_users.PatchProps{
		Db: gormDB,
	})(c)
}
