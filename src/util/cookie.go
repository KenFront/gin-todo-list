package util

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	authKey = "auth"
)

func CheckAuth(c *gin.Context) (*authClaims, error) {
	cookie, err := c.Cookie(authKey)
	if err != nil {
		return &authClaims{}, errors.New("not sign in yet")
	}

	parsed, err := ParseJwtToken(cookie)
	if err != nil {
		return parsed, err
	}

	return parsed, nil
}

func SetAuth(c *gin.Context, userId string) {
	c.SetCookie(authKey, NewJwtToken(userId), GetAuthDuration(), "/", os.Getenv("DOMAIN"), os.Getenv("DOMAIN") != "localhost", true)
}

func DeleteAuth(c *gin.Context) {
	c.SetCookie(authKey, "DELETED", -1, "/", os.Getenv("DOMAIN"), os.Getenv("DOMAIN") != "localhost", true)
}
