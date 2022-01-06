package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/santosdavid/go-api-v2/domain/customer"
)

type Secrets struct {
	AcessKey     []byte
	RefreshToken []byte
}

type customClaims struct {
	Id    customer.CustomerID
	Email string
	Uuid  string
	jwt.RegisteredClaims
}

type FullCredentials struct {
	AcessToken   credentials
	RefreshToken credentials
}

type credentials struct {
	Token   string
	UUid    string
	Expires int64
}

type CustomerData struct {
	Id    customer.CustomerID
	Email string
}
