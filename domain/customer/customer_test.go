package customer_test

import (
	"testing"

	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	var tests = []struct {
		testName string
		payDay   int
		email    string
		password string
		want     error
	}{
		{"email without @", 10, "teste", "123456", customer.ErrorInvalidEmail},
		{"valid email", 10, "test@test.com", "123456", nil},
		{"invalid payday", 1, "test@test.com", "123456", customer.ErrorInvalidPayDay},
		{"valid payday", 25, "test@test.com", "123456", nil},
	}

	for _, tt := range tests {
		_, err := customer.NewCustomer(tt.payDay, tt.email, tt.password)

		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, err, tt.want)
		})
	}
}

func TestNewCustomerPasswordHash(t *testing.T) {
	password := "123456"

	c, err := customer.NewCustomer(10, "test@test.com", password)

	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.NotEqual(t, password, c.Password)
}

func TestCheckPassword(t *testing.T) {
	password := "muitodificil"

	c, err := customer.NewCustomer(10, "test@test.com", password)

	assert.Nil(t, err)
	assert.True(t, c.CheckPassword(password))
}
