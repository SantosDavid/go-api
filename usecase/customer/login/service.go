package login

import "github.com/santosdavid/go-api-v2/domain/customer"

type Service struct {
	repo customer.Repository
}

func New(repo customer.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Login(email string, password string) (Response, error) {
	c, err := s.repo.FindByEmail(email)
	if err != nil {
		return Response{}, err
	}

	if !c.CheckPassword(password) {
		return Response{}, customer.ErrorInvalidPassword
	}

	return Response{c.CustomerID, c.Email.ToString()}, nil
}
