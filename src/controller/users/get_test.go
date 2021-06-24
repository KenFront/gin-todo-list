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

func TestGetUserSuccess(t *testing.T) {
	resForAdd := mock.GetResponse()
	userId := util.GetNewTodoId()
	cForAdd := mock.GetGinContext(resForAdd)

	gormDB := mock.GetMockGorm(t)

	fake := mock.GetMockUser()

	cForAdd.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	controller_users.Add(controller_users.AddProps{
		Db:           gormDB,
		GetNewUserId: func() uuid.UUID { return userId },
	})(cForAdd)

	assert.Equal(t, http.StatusOK, resForAdd.Code)

	resForGet := mock.GetResponse()
	cForGet := mock.GetGinContext(resForGet)

	controller.SetUserId(cForGet, userId)

	controller_users.Get(controller_users.GetProps{
		Db: gormDB,
	})(cForGet)

	var resBody SuccessUserAPIResponse
	mock.GetResponseBody(resForGet.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, resForGet.Code)
	assert.Equal(t, fake.Name, resBody.Data.Name)
	assert.Equal(t, fake.Email, resBody.Data.Email)
	assert.Equal(t, fake.Account, resBody.Data.Account)
}

func TestGetUserFailedByNotSignIn(t *testing.T) {
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
