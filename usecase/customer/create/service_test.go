package create_test

import (
	"testing"

	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/santosdavid/go-api-v2/infrastructure/repository"
	"github.com/santosdavid/go-api-v2/usecase/customer/create"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	s := create.New(&repository.InMemory{})
	createCustomerRequest := &create.Request{10, "teste@test.com.br", "123456"}
	_, err := s.Create(*createCustomerRequest)
	assert.Nil(t, err)

	email := "testando@test.com.br"
	cs2, _ := customer.NewCustomer(10, email, "1345453")
	repo2 := &repository.InMemory{Customers: []*customer.Customer{cs2}}
	s2 := create.New(repo2)
	createCustomerRequest2 := &create.Request{10, email, "123456"}
	_, err2 := s2.Create(*createCustomerRequest2)
	assert.Equal(t, customer.ErrorDuplicatedEmail, err2)
}
