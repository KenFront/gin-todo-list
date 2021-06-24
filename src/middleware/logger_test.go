package middleware_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/middleware"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/util"
)

func TestLoggerSuccess(t *testing.T) {
	res := mock.GetResponse()
	userId := util.GetNewTodoId()
	c := mock.GetGinContext(res)
	controller.SetUserId(c, userId)

	c.Request = &http.Request{
		Header: make(http.Header),
		URL: &url.URL{
			Path: "/path",
		},
	}

	middleware.CustomLogger(c)
}

func TestLoggerSuccessByJsonPayload(t *testing.T) {
	res := mock.GetResponse()
	userId := util.GetNewTodoId()
	c := mock.GetGinContext(res)
	controller.SetUserId(c, userId)

	payload := map[string]string{
		"hello": "world",
	}

	c.Request = &http.Request{
		Header: make(http.Header),
		URL: &url.URL{
			Path: "/path",
		},
		Body: mock.GetRequsetBody(payload),
	}

	middleware.CustomLogger(c)
}

func TestLoggerSuccessByForm(t *testing.T) {
	res := mock.GetResponse()
	userId := util.GetNewTodoId()
	c := mock.GetGinContext(res)
	controller.SetUserId(c, userId)

	form := map[string][]string{
		"hello": {"world"},
	}

	c.Request = &http.Request{
		Header: make(http.Header),
		Form:   form,
		URL: &url.URL{
			Path: "/path",
		},
	}

	middleware.CustomLogger(c)
}
