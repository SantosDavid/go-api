package customer

import (
	"errors"

	"github.com/santosdavid/go-api-v2/domain"
	"github.com/santosdavid/go-api-v2/domain/shared"
	"github.com/santosdavid/go-api-v2/domain/vo"
	"golang.org/x/crypto/bcrypt"
)

type CustomerID = domain.ID

var ErrorInvalidEmail = errors.New("invalid email")
var ErrorInvalidPayDay = errors.New("invalid payday")
var ErrorToHashPassword = errors.New("error when trying to hash password")
var ErrorInvalidPassword = errors.New("error password is invalid")
var ErrorDuplicatedEmail = errors.New("this email already exists")
var ErrorEntityNotFound = errors.New("entity not found")

var payDays = []int{5, 10, 15, 25}

type Customer struct {
	CustomerID
	Payday   int
	Email    vo.Email
	Password string
}

func NewCustomer(payday int, email string, password string) (*Customer, error) {
	e, err := vo.NewEmail(email)
	if err != nil {
		return nil, ErrorInvalidEmail
	}

	c := &Customer{
		CustomerID: domain.New(),
		Payday:     payday,
		Email:      e,
		Password:   password,
	}

	if err := c.validatePayDay(); err != nil {
		return nil, err
	}

	if err := c.hashPassword(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Customer) validatePayDay() error {
	if !shared.Contains(payDays, c.Payday) {
		return ErrorInvalidPayDay
	}

	return nil
}

func (c *Customer) hashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(c.Password), 14)

	if err != nil {
		return ErrorToHashPassword
	}

	c.Password = string(bytes)

	return nil
}

func (c *Customer) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))
	return err == nil
}
