package controller_auth_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/KenFront/gin-todo-list/src/controller"
	controller_auth "github.com/KenFront/gin-todo-list/src/controller/auth"
	controller_users "github.com/KenFront/gin-todo-list/src/controller/users"
	"github.com/KenFront/gin-todo-list/src/mock"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func getUserIdFromToken(res *httptest.ResponseRecorder) string {

	cookie := res.Header().Get("Set-Cookie")
	cookieProperties := strings.Split(cookie, "; ")
	regexStr := `^auth=(.*)`
	re := regexp.MustCompile(regexStr)
	auth := re.ReplaceAllString(cookieProperties[0], `$1`)

	parsed, _ := util.ParseJwtToken(auth)
	return parsed.UserId
}
func TestSignInSuccess(t *testing.T) {
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

	resForSignIn := mock.GetResponse()
	cForSignIn := mock.GetGinContext(resForSignIn)

	payload := model.SignIn{
		Account:  fake.Account,
		Password: fake.Password,
	}

	cForSignIn.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(payload),
	}

	controller.SetUserId(cForSignIn, userId)

	controller_auth.SignIn(controller_auth.SignInProps{
		Db: gormDB,
	})(cForSignIn)

	userIdInToken := getUserIdFromToken(resForSignIn)

	assert.Equal(t, http.StatusOK, resForSignIn.Code)
	assert.Equal(t, userId.String(), userIdInToken)
}

func TestSignInSuccessFailedByMissingRequiredPayload(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)

	payload := model.SignIn{
		Account:  "123",
		Password: "456",
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
		assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
		assert.Equal(t, model.ERROR_SIGN_IN_FAILED, err.ErrorType)
	}()

	controller_auth.SignIn(controller_auth.SignInProps{
		Db: gormDB,
	})(c)
}

func TestSignInSuccessFailedByNotExistedUser(t *testing.T) {
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
		assert.Equal(t, model.ERROR_SIGN_IN_PAYLOAD_IS_INVALID, err.ErrorType)
	}()

	controller_auth.SignIn(controller_auth.SignInProps{
		Db: gormDB,
	})(c)
}

func TestSignInFailedByWrongPassword(t *testing.T) {
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

	resForSignIn := mock.GetResponse()
	cForSignIn := mock.GetGinContext(resForSignIn)

	payload := model.SignIn{
		Account:  fake.Account,
		Password: "123",
	}

	cForSignIn.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(payload),
	}

	controller.SetUserId(cForSignIn, userId)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
		assert.Equal(t, model.ERROR_SIGN_IN_FAILED, err.ErrorType)
	}()

	controller_auth.SignIn(controller_auth.SignInProps{
		Db: gormDB,
	})(cForSignIn)
}

func TestSignInFailedByIsInactive(t *testing.T) {
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

	resForPatch := mock.GetResponse()
	cForPatch := mock.GetGinContext(resForPatch)
	controller.SetUserId(cForPatch, userId)

	patchPayload := model.PatchUser{
		Status: model.USER_INACTIVE,
	}

	cForPatch.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(patchPayload),
	}

	controller_users.Patch(controller_users.PatchProps{
		Db: gormDB,
	})(cForPatch)

	assert.Equal(t, http.StatusOK, resForAdd.Code)

	resForSignIn := mock.GetResponse()
	cForSignIn := mock.GetGinContext(resForSignIn)

	payload := model.SignIn{
		Account:  fake.Account,
		Password: fake.Password,
	}

	cForSignIn.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(payload),
	}

	controller.SetUserId(cForSignIn, userId)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusServiceUnavailable, err.StatusCode)
		assert.Equal(t, model.ERROR_USER_IS_NOT_ACTIVE, err.ErrorType)
	}()

	controller_auth.SignIn(controller_auth.SignInProps{
		Db: gormDB,
	})(cForSignIn)
}
