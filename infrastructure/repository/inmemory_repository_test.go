package repository_test

import (
	"testing"

	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/santosdavid/go-api-v2/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	c1, _ := customer.NewCustomer(10, "test@test.com.br", "123456")
	c2, _ := customer.NewCustomer(10, "test@test.com.br", "123456")
	c3, _ := customer.NewCustomer(10, "test@test.com.br", "123456")

	customers := []*customer.Customer{
		c1, c2, c3,
	}

	repo := repository.InMemory{}

	for _, customer := range customers {
		err := repo.Create(customer)

		assert.Nil(t, err)
	}

	assert.Equal(t, len(customers), len(repo.Customers))
}

func TestFindByEmailNotFound(t *testing.T) {
	tests := []struct {
		email       string
		emailsToAdd []string
		testName    string
	}{
		{
			"test@test.com",
			[]string{},
			"email not found",
		},
		{
			"testtest.com",
			[]string{},
			"email invalid returns not found error",
		},
		{
			"test@test.com",
			[]string{"a@b.com", "test@tes.com", "tes@test.com.br", "test@test.com.br"},
			"email invalid returns not found error when there is multiple customers",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			var customers []*customer.Customer

			for _, email := range tt.emailsToAdd {
				c, _ := customer.NewCustomer(10, email, "123456")

				customers = append(customers, c)
			}

			repo := repository.InMemory{customers}

			c, err := repo.FindByEmail(tt.email)

			assert.Nil(t, c)
			assert.Equal(t, customer.ErrorEntityNotFound, err)
		})
	}
}

func TestFindByEmailFound(t *testing.T) {
	tests := []struct {
		email       string
		emailsToAdd []string
		testName    string
	}{
		{
			"test@test.com",
			[]string{"test@test.com"},
			"email found",
		},
		{
			"test@test.com",
			[]string{"test@tes.com", "tes@test.com.br", "test@test.com.br", "test@test.com"},
			"email returns found error when there is multiple customers",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			var customers []*customer.Customer

			for _, email := range tt.emailsToAdd {
				c, _ := customer.NewCustomer(10, email, "123456")

				customers = append(customers, c)
			}

			repo := repository.InMemory{customers}

			c, err := repo.FindByEmail(tt.email)

			assert.NotNil(t, c)
			assert.Nil(t, err)
		})
	}
}
