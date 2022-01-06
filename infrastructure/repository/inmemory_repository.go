package repository

import (
	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/santosdavid/go-api-v2/domain/vo"
)

type InMemory struct {
	Customers []*customer.Customer
}

func New() *InMemory {
	return &InMemory{}
}

func (r *InMemory) Create(c *customer.Customer) error {
	r.Customers = append(r.Customers, c)

	return nil
}

func (r *InMemory) FindByEmail(email string) (*customer.Customer, error) {
	e, _ := vo.NewEmail(email)

	for _, customer := range r.Customers {
		if customer.Email.EqualsTo(e) {
			return customer, nil
		}
	}

	return nil, customer.ErrorEntityNotFound
}
