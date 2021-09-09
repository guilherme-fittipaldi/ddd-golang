package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/guilherme-fittipaldi/ddd-golang/aggregate"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) ([]aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}