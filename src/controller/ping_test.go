package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	Ping(c)

	var jsonResponse gin.H
	if err := json.Unmarshal(res.Body.Bytes(), &jsonResponse); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, gin.H{
		"data": "Server is working",
	}, jsonResponse)
}
