package util

import (
	"errors"
	"math"
	"os"

	"github.com/gin-gonic/gin"
)

var AuthKey = "auth"
var hour = int(math.Pow(60, 2))
var authDuration = 1 * hour

func CheckAuth(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(AuthKey)
	if err != nil {
		return cookie, errors.New("not sign in yet")
	}
	return cookie, nil
}

func SetAuth(c *gin.Context) {
	c.SetCookie(AuthKey, "temp", authDuration, "/", os.Getenv("DOMAIN"), os.Getenv("DOMAIN") != "localhost", true)
}

func DeleteAuth(c *gin.Context) {
	c.SetCookie(AuthKey, "temp", -1, "/", os.Getenv("DOMAIN"), os.Getenv("DOMAIN") != "localhost", true)
}
