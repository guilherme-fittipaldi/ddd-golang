package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/guilherme-fittipaldi/ddd-golang/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add customer not found in the repository")
	ErrUpdateCustomer      = errors.New("failed to update customer not found in the repository")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
