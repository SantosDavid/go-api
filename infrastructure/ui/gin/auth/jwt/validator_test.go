package jwt

import (
	"errors"
	"testing"

	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/storage"
	"github.com/stretchr/testify/assert"
)

var tokenValidTest = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6ImNlOTAxY2M2LWZlOGEtNDM1YS05NTViLTYzMjQ2MWVmMTA1NSIsIkVtYWlsIjoidGVzdEB0ZXN0LmNvbSIsIlV1aWQiOiIxN2FmYjg1NC02ZDk5LTQ1NDctOTYwNi1lNDBlMjIzODlhZjErK2NlOTAxY2M2LWZlOGEtNDM1YS05NTViLTYzMjQ2MWVmMTA1NSIsImlzcyI6ImdvLWFwaSIsImV4cCI6NTIzODExOTQxOCwiaWF0IjoxNjM4MTE5NDE4fQ.0SAMm3fz_kbnQ9EWScv6UgeRsn5Gx1KNbTvGXN7Tngs"
var tokenValidUUID = "17afb854-6d99-4547-9606-e40e22389af1++ce901cc6-fe8a-435a-955b-632461ef1055"
var userUUID = "ce901cc6-fe8a-435a-955b-632461ef1055"

func TestParseToken(t *testing.T) {
	tests := []struct {
		testName    string
		storageData map[string]string
		token       string
		secret      []byte
		errExpected interface{}
	}{
		{
			"error when signature is invalid",
			map[string]string{},
			tokenValidTest,
			nil,
			errors.New("signature is invalid"),
		},
		{
			"error when signature is none",
			map[string]string{},
			"ewogICJhbGciOiAibm9uZSIsCiAgInR5cCI6ICJKV1QiCn0.ewogICJJZCI6ICJjZTkwMWNjNi1mZThhLTQzNWEtOTU1Yi02MzI0NjFlZjEwNTUiLAogICJFbWFpbCI6ICJ0ZXN0QHRlc3QuY29tIiwKICAiVXVpZCI6ICIxN2FmYjg1NC02ZDk5LTQ1NDctOTYwNi1lNDBlMjIzODlhZjErK2NlOTAxY2M2LWZlOGEtNDM1YS05NTViLTYzMjQ2MWVmMTA1NSIsCiAgImlzcyI6ICJnby1hcGkiLAogICJleHAiOiA1MjM4MTE5NDE4LAogICJpYXQiOiAxNjM4MTE5NDE4Cn0.",
			[]byte(""),
			errors.New("'none' signature type is not allowed"),
		},
		{
			"error when token is expired",
			map[string]string{},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjFmNmIwZDQ1LWY0MmMtNDg1Ny1iNTc4LThjM2ViNDMyOTM3MiIsIkVtYWlsIjoidGVzdEB0ZXN0LmNvbSIsIlV1aWQiOiIwYzg4MGQxNS1jYjFjLTRhZjEtYTgzZi01YjExYjVlZmY4OGMrKzFmNmIwZDQ1LWY0MmMtNDg1Ny1iNTc4LThjM2ViNDMyOTM3MiIsImlzcyI6ImdvLWFwaSIsImV4cCI6MTYzODIzMzIyOCwiaWF0IjoxNjM4MjMzMjI3fQ.O-BFwhgUacf6iCIqCbHBzLq_1hBes1xcjMJGMIs_NzY",
			[]byte("123456"),
			errors.New("signature is invalid"),
		},
		{
			"error when is not found on storage",
			map[string]string{},
			tokenValidTest,
			[]byte("123456"),
			ErrorTokenNotFound,
		},
		{
			"error when is not found on storage",
			map[string]string{},
			tokenValidTest,
			[]byte("123456"),
			ErrorTokenNotFound,
		},
		{
			"error when token belogs to another user",
			map[string]string{tokenValidUUID: "1"},
			tokenValidTest,
			[]byte("123456"),
			ErrorTokenDoesNotBelongsToUser,
		},
		{
			"sucess",
			map[string]string{tokenValidUUID: userUUID},
			tokenValidTest,
			[]byte("123456"),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			storage := storage.Inmemory{Items: tt.storageData}
			secrets := Secrets{tt.secret, nil}

			_, err := ParseToken(&storage, secrets, tt.token)

			assert.Equal(t, tt.errExpected, err)
		})
	}
}
