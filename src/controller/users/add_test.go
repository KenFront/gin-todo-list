package controller_users_test

import (
	"net/http"
	"testing"

	controller_users "github.com/KenFront/gin-todo-list/src/controller/users"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
)

type SuccessUserAPIResponse struct {
	Data model.User `json:"data"`
}

func TestAddUserSuccess(t *testing.T) {
	res := mock.GetResponse()
	userId := util.GetNewTodoId()
	c := mock.GetGinContext(res)

	gormDB := mock.GetMockGorm(t)

	fake := model.AddUser{
		Name:  "Testing",
		Email: "test@test.com",
		SignIn: model.SignIn{
			Account:  "test",
			Password: "test",
		},
	}

	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	var resBody SuccessUserAPIResponse

	controller_users.Add(controller_users.AddProps{
		Db:           gormDB,
		GetNewUserId: func() uuid.UUID { return userId },
	})(c)
	mock.GetResponseBody(res.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, fake.Name, resBody.Data.Name)
	assert.Equal(t, fake.Email, resBody.Data.Email)
	assert.Equal(t, fake.Account, resBody.Data.Account)
	assert.NotEqual(t, fake.Password, resBody.Data.Password)
}
