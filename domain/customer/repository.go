package customer

type Repository interface {
	Create(c *Customer) error
	FindByEmail(email string) (*Customer, error)
}
