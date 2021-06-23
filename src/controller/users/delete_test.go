package controller_users_test

import (
	"net/http"
	"testing"

	controller_users "github.com/KenFront/gin-todo-list/src/controller/users"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUserSuccess(t *testing.T) {
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

	resForDelete := mock.GetResponse()
	cForDelete := mock.GetGinContext(resForDelete)
	cForDelete.Set("userId", userId)

	controller_users.Delete(controller_users.DeleteProps{
		Db: gormDB,
	})(cForDelete)

	var resBody SuccessUserAPIResponse
	mock.GetResponseBody(resForDelete.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, resForDelete.Code)
	assert.Equal(t, fake.Name, resBody.Data.Name)
	assert.Equal(t, fake.Email, resBody.Data.Email)
	assert.Equal(t, fake.Account, resBody.Data.Account)
}

func TestDeleteUserFailedByNotSignIn(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)

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

	controller_users.Delete(controller_users.DeleteProps{
		Db: gormDB,
	})(c)
}

func TestDeleteUserFailedByNotExisted(t *testing.T) {
	res := mock.GetResponse()
	userId := util.GetNewTodoId()
	c := mock.GetGinContext(res)
	c.Set("userId", userId)

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusServiceUnavailable, err.StatusCode)
		assert.Equal(t, model.ERROR_DELETE_USER_NOT_EXIST, err.ErrorType)
	}()

	controller_users.Delete(controller_users.DeleteProps{
		Db: gormDB,
	})(c)
}
