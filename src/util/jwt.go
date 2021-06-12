package util

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type authClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

var (
	issuer = os.Getenv("JWT_ISSUER")
	secret = []byte(os.Getenv("JWT_SECRET"))
)

func NewJwtToken(userId string) (string, error) {
	claims := authClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: GetAuthExpiresAt(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ParseJwtToken(clientToken string) (*authClaims, error) {
	token, err := jwt.ParseWithClaims(
		clientToken,
		&authClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
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
