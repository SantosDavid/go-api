package login_test

import (
	"testing"

	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/santosdavid/go-api-v2/infrastructure/repository"
	"github.com/santosdavid/go-api-v2/usecase/customer/login"
	"github.com/stretchr/testify/assert"
)

func newFixtureCustomer(email string, password string) *customer.Customer {
	c, _ := customer.NewCustomer(
		10,
		email,
		password,
	)

	return c
}

func TestLogin(t *testing.T) {
	repo1 := repository.InMemory{}
	service1 := login.New(&repo1)
	_, err1 := service1.Login("david@gmail.com", "123456")
	assert.Equal(t, customer.ErrorEntityNotFound, err1)

	repo2 := repository.InMemory{
		Customers: []*customer.Customer{newFixtureCustomer("david@gmail.com", "1")},
	}
	service2 := login.New(&repo2)
	_, err2 := service2.Login("david@gmail.com", "123456")
	assert.Equal(t, customer.ErrorInvalidPassword, err2)

	c3 := newFixtureCustomer("david@gmail.com", "muitodificil")
	repo3 := repository.InMemory{
		Customers: []*customer.Customer{c3},
	}
	service3 := login.New(&repo3)
	resp, err3 := service3.Login("david@gmail.com", "muitodificil")
	assert.Nil(t, err3)
	assert.Equal(t, c3.CustomerID, resp.Id)
	assert.Equal(t, c3.Email.ToString(), resp.Email)
}
