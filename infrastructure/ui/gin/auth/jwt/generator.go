package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/storage"
	"github.com/santosdavid/go-api-v2/usecase/customer/login"
)

func Generate(storage storage.Storage, secrets Secrets, r *login.Response) (*FullCredentials, error) {
	t, err := createFullCredentials(r, secrets)
	if err != nil {
		return &FullCredentials{}, err
	}

	err = storage.Set(t.AcessToken.UUid, r.Id.String(), secondsRemaining(t.AcessToken.Expires))
	if err != nil {
		return &FullCredentials{}, err
	}

	err = storage.Set(t.RefreshToken.UUid, r.Id.String(), secondsRemaining(t.RefreshToken.Expires))
	if err != nil {
		return &FullCredentials{}, err
	}

	return t, err
}

func createFullCredentials(r *login.Response, secrets Secrets) (*FullCredentials, error) {
	acessTokenCredentials, err := createAccessToken(r, secrets.AcessKey)
	if err != nil {
		return &FullCredentials{}, err
	}

	refreshTokenCredentials, err := createRefreshToken(r, secrets.RefreshToken)
	if err != nil {
		return &FullCredentials{}, err
	}

	return &FullCredentials{
		AcessToken:   acessTokenCredentials,
		RefreshToken: refreshTokenCredentials,
	}, nil
}

func createAccessToken(r *login.Response, secret []byte) (credentials, error) {
	accessTokenClaims := createClaims(r, time.Minute*30)
	acessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := acessToken.SignedString(secret)
	if err != nil {
		return credentials{}, err
	}

	return credentials{
		Token:   accessTokenString,
		UUid:    accessTokenClaims.Uuid,
		Expires: accessTokenClaims.ExpiresAt.Unix(),
	}, nil
}

func createRefreshToken(r *login.Response, secret []byte) (credentials, error) {
	refreshTokenClaims := createClaims(r, time.Hour*24*7)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(secret)
	if err != nil {
		return credentials{}, err
	}

	return credentials{
		Token:   refreshTokenString,
		UUid:    refreshTokenClaims.Uuid,
		Expires: refreshTokenClaims.ExpiresAt.Unix(),
	}, nil
}

func createClaims(r *login.Response, expiration time.Duration) customClaims {
	return customClaims{
		r.Id,
		r.Email,
		uuid.NewString() + "++" + r.Id.String(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-api",
		},
	}
}

func secondsRemaining(seconds int64) time.Duration {
	now := time.Now()

	time := time.Unix(seconds, 0)

	return time.Sub(now)
}
