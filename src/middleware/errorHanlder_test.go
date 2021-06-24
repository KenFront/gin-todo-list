package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/middleware"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ErrorAPIResponse struct {
	Error model.ErrorType `json:"error"`
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestErrorHandlerSuccessByCutomError(t *testing.T) {
	router := gin.New()
	middleware.UseErrorHandler(router)

	router.GET("/recovery", func(c *gin.Context) {
		controller.ApiOnError(&model.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  model.ERROR_SIGN_IN_FAILED,
			Error:      errors.New("test"),
		})
	})

	res := performRequest(router, "GET", "/recovery")

	var resBody ErrorAPIResponse
	mock.GetResponseBody(res.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, model.ERROR_SIGN_IN_FAILED, resBody.Error)
}

func TestErrorHandlerSuccessByUnexpectedError(t *testing.T) {
	router := gin.New()
	middleware.UseErrorHandler(router)

	router.GET("/recovery", func(c *gin.Context) {
		panic(errors.New("test"))
	})

	res := performRequest(router, "GET", "/recovery")

	var resBody ErrorAPIResponse
	mock.GetResponseBody(res.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, model.ERROR_UNKNOWN, resBody.Error)
}

func TestErrorHandlerSuccessByPanicWithoutError(t *testing.T) {
	router := gin.New()
	middleware.UseErrorHandler(router)

	router.GET("/recovery", func(c *gin.Context) {
		panic("test")
	})

	res := performRequest(router, "GET", "/recovery")

	var resBody ErrorAPIResponse
	mock.GetResponseBody(res.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, model.ERROR_UNKNOWN, resBody.Error)
}
