package jwt

import (
	"errors"

	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/storage"

	"github.com/golang-jwt/jwt/v4"
)

var ErrorInvalidSignature = errors.New("invalid signature")
var ErrorInvalidToken = errors.New("invalid token")
var ErrorTokenNotFound = errors.New("token not found")
var ErrorTokenDoesNotBelongsToUser = errors.New("token does not belogs to user")

func ParseToken(storage storage.Storage, secrets Secrets, tokenString string) (CustomerData, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secrets.AcessKey, nil
	})
	if err != nil {
		return CustomerData{}, errors.New(err.Error())
	}

	claims, ok := token.Claims.(*customClaims)

	if !ok || !token.Valid {
		return CustomerData{}, ErrorInvalidToken
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return CustomerData{}, ErrorInvalidSignature
	}

	v, err := storage.Get(claims.Uuid)
	if err != nil {
		return CustomerData{}, ErrorTokenNotFound
	}

	if v != claims.Id.String() {
		return CustomerData{}, ErrorTokenDoesNotBelongsToUser
	}

	return CustomerData{claims.Id, claims.Email}, nil
}
