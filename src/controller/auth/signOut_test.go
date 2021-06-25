package controller_auth_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	controller_auth "github.com/KenFront/gin-todo-list/src/controller/auth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignOutSuccess(t *testing.T) {
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)

	controller_auth.SignOut(c)

	cookie := res.Header().Get("Set-Cookie")
	cookieProperties := strings.Split(cookie, "; ")

	assert.Equal(t, http.StatusOK, res.Code)
	for _, v := range cookieProperties {
		switch {
		case strings.HasPrefix(v, "auth"):
			assert.Equal(t, "auth=DELETED", v)
		case strings.HasPrefix(v, "Max-Age"):
			assert.Equal(t, "Max-Age=0", v)
		}
	}
}
