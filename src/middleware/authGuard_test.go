package middleware_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
	controller_users "github.com/KenFront/gin-todo-list/src/controller/users"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"

	"github.com/KenFront/gin-todo-list/src/middleware"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthGuardSuccess(t *testing.T) {
	res := mock.GetResponse()
	userId := util.GetNewTodoId()
	c := mock.GetGinContext(res)

	gormDB := mock.GetMockGorm(t)

	fake := mock.GetMockUser()

	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	controller_users.Add(controller_users.AddProps{
		Db:           gormDB,
		GetNewUserId: func() uuid.UUID { return userId },
	})(c)

	assert.Equal(t, http.StatusOK, res.Code)

	middleware.AuthGuard(middleware.AuthGuardProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(c)

	id, _ := controller.GetUserId(c)

	assert.Equal(t, userId, id)
}

func TestAuthGuardFailByParseUserIdFailed(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	customError := ""

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Equal(t, model.ErrorType(customError), err.ErrorType)
	}()

	middleware.AuthGuard(middleware.AuthGuardProps{
		Db:               gormDB,
		GetUserIdByToken: func(c *gin.Context) (uuid.UUID, error) { return uuid.Nil, errors.New(customError) },
	})(c)
}

func TestAuthGuardFailByUserNotExist(t *testing.T) {
	res := mock.GetResponse()
	userId := util.GetNewTodoId()
	c := mock.GetGinContext(res)

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusServiceUnavailable, err.StatusCode)
		assert.Equal(t, model.ERROR_SIGN_IN_FAILED, err.ErrorType)
	}()

	middleware.AuthGuard(middleware.AuthGuardProps{
		Db:               gormDB,
		GetUserIdByToken: mock.UtilGetUserIdByToken(userId),
	})(c)
}
