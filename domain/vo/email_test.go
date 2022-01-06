package vo_test

import (
	"errors"
	"testing"

	"github.com/santosdavid/go-api-v2/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		testName string
		email    string
		expected interface{}
	}{
		{"invalid only numbers", "123123", errors.New("invalid email")},
		{"invalid without domain", "test@", errors.New("invalid email")},
		{"valid", "test@go.com", nil},
		{"valid with br", "test@go.com.br", nil},
		{"valid with numbers", "test123@go.com.br", nil},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := vo.NewEmail(tt.email)

			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestEqualsTo(t *testing.T) {
	tests := []struct {
		baseEmail      string
		emailToCompare string
		expected       bool
	}{
		{"test@test.com", "test@test.com.br", false},
		{"test@test.com.ba", "test@test.com.br", false},
		{"test@test.com.br", "test@test.com.br", true},
	}

	for _, tt := range tests {
		base, _ := vo.NewEmail(tt.baseEmail)
		toCompare, _ := vo.NewEmail(tt.emailToCompare)

		r := base.EqualsTo(toCompare)

		assert.Equal(t, tt.expected, r)
	}
}
