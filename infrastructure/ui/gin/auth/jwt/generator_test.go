package jwt

import (
	"testing"
	"time"

	"github.com/santosdavid/go-api-v2/domain"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/storage"
	"github.com/santosdavid/go-api-v2/usecase/customer/login"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	storage := storage.Inmemory{Items: map[string]string{}}
	secret := Secrets{[]byte("1234"), []byte("12345")}
	r := &login.Response{Id: domain.New(), Email: "test@test.com"}

	credentials, err := Generate(&storage, secret, r)

	item1, _ := storage.Get(credentials.AcessToken.UUid)
	item2, _ := storage.Get(credentials.RefreshToken.UUid)

	assert.Nil(t, err)
	assert.NotNil(t, item1)
	assert.NotNil(t, item2)

}

func TestCreateAcessToken(t *testing.T) {
	uuid := domain.New()
	r := &login.Response{
		Id:    uuid,
		Email: "test@test.com",
	}

	credentials, err := createAccessToken(r, []byte("abcd"))

	assert.Nil(t, err)
	assert.Contains(t, credentials.UUid, uuid.String())
}

func TestCreateRefreshToken(t *testing.T) {
	uuid := domain.New()
	r := &login.Response{
		Id:    uuid,
		Email: "test@test.com",
	}

	credentials, err := createRefreshToken(r, []byte("abcd"))

	assert.Nil(t, err)
	assert.Contains(t, credentials.UUid, uuid.String())
}

func TestCreateClaims(t *testing.T) {
	uuid := domain.New()

	r := &login.Response{
		Id:    uuid,
		Email: "test@test.com",
	}

	claims := createClaims(r, time.Duration(1))

	assert.Equal(t, uuid, claims.Id)
	assert.Equal(t, "test@test.com", claims.Email)
	assert.Equal(t, "go-api", claims.Issuer)
	assert.Contains(t, claims.Uuid, uuid.String())
}
