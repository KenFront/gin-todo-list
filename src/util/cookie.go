package util

import (
	"errors"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/gin-gonic/gin"
)

var (
	authKey = "auth"
)

func getToken(c *gin.Context) (*authClaims, error) {
	cookie, err := c.Cookie(authKey)
	if err != nil {
		return &authClaims{}, errors.New(string(model.ERROR_NOT_SIGN_IN_YET))
	}

	parsed, err := ParseJwtToken(cookie)
	if err != nil {
		return parsed, errors.New(string(model.ERROR_TOKEN_PARSE_FAILED))
	}

	if parsed.ExpiresAt < GetAuthNow() {
		return &authClaims{}, errors.New(string(model.ERROR_USER_TOKEN_IS_EXPIRED_FAILED))
	}

	return parsed, nil
}

func SetAuth(c *gin.Context, userId string) error {
	token, err := newJwtToken(userId)
	if err != nil {
		return err
	}
	env := config.GetEnv()
	c.SetCookie(authKey, token, getAuthDuration(), "/", env.DOMAIN, env.DOMAIN != "localhost", true)
	return nil
}

func DeleteAuth(c *gin.Context) {
	env := config.GetEnv()

	c.SetCookie(authKey, "DELETED", -1, "/", env.DOMAIN, env.DOMAIN != "localhost", true)
}
