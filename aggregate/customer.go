package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/guilherme-fittipaldi/ddd-golang/entity"
	"github.com/guilherme-fittipaldi/ddd-golang/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

type Customer struct {
	Person       *entity.Person            `bson:"person"`
	Products     []*entity.Item            `bson:"products"`
	Transactions []valueobject.Transaction `bson:"transactions"`
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{Name: name, ID: uuid.New()}

	return Customer{
		Person:       person,
		Products:     make([]*entity.Item, 0),
		Transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.Person.ID
}

func (c *Customer) SetName(name string) {
	c.Person.Name = name
}
