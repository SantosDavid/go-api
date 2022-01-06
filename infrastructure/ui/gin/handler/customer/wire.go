//go:build wireinject
// +build wireinject

package customer

import (
	"github.com/google/wire"
	"github.com/santosdavid/go-api-v2/domain/customer"
	"github.com/santosdavid/go-api-v2/infrastructure/repository"
)

func New() (*Handler, error) {
	wire.Build(
		repository.Wired,

		wire.Bind(new(customer.Repository), new(*repository.InMemory)),

		newHandler,
	)

	return nil, nil
}
