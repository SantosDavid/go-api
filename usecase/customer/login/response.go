package login

import "github.com/santosdavid/go-api-v2/domain/customer"

type Response struct {
	Id    customer.CustomerID
	Email string
}
