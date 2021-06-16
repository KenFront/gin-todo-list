package mock

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetGinContext(res *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(res)
	return c
}
