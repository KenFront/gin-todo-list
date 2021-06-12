package util

import (
	"errors"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/dgrijalva/jwt-go"
)

type authClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func NewJwtToken(userId string) (string, error) {
	env := config.GetEnv()
	claims := authClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: GetAuthExpiresAt(),
			Issuer:    env.JWT_ISSUER,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(env.JWT_SECRET)
}

func ParseJwtToken(clientToken string) (*authClaims, error) {

	token, err := jwt.ParseWithClaims(
		clientToken,
		&authClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return config.GetEnv().JWT_SECRET, nil
		},
	)
	if err != nil {
		return &authClaims{}, err
	}

	claims, ok := token.Claims.(*authClaims)
	if !ok {
		return &authClaims{}, errors.New("couldn't parse")
	}

	if claims.ExpiresAt < GetAuthNow() {
		return &authClaims{}, errors.New("JWT is expired")
	}

	return claims, nil
}
