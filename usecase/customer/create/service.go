package create

import (
	"github.com/google/uuid"
	"github.com/santosdavid/go-api-v2/domain/customer"
)

type Service struct {
	repo customer.Repository
}

func New(r customer.Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s Service) Create(r Request) (customer.CustomerID, error) {
	c, err := customer.NewCustomer(
		r.PayDay,
		r.Email,
		r.Password,
	)

	if err != nil {
		return uuid.UUID{}, err
	}

	cs, err := s.repo.FindByEmail(c.Email.ToString())

	if cs != nil {
		return uuid.UUID{}, customer.ErrorDuplicatedEmail
	}

	if err != nil && err != customer.ErrorEntityNotFound {
		return uuid.UUID{}, err
	}

	err = s.repo.Create(c)

	if err != nil {
		return uuid.UUID{}, err
	}

	return c.CustomerID, nil
}
