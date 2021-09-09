package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/guilherme-fittipaldi/ddd-golang/entity"
)

var (
	ErrMissingValues = errors.New("missing values")
)

type Product struct {
	Item     *entity.Item `json:"item" bson:"item"`
	Price    float64      `json:"price" bson:"price"`
	Quantity int          `json:"quantity" bson:"quantity"`
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		Item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		Price:    price,
		Quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.Item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.Item
}

func (p Product) GetPrice() float64 {
	return p.Price
}
