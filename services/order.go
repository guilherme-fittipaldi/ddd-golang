package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/guilherme-fittipaldi/ddd-golang/aggregate"
	"github.com/guilherme-fittipaldi/ddd-golang/domain/customer"
	"github.com/guilherme-fittipaldi/ddd-golang/domain/customer/memory"
	"github.com/guilherme-fittipaldi/ddd-golang/domain/product"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithMemoryCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return price, nil
}

// func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
// 	return func(os *OrderService) error {
// 		pr := prodmemory.New()
		
// 		for _, p := range products {
// 			err := pr.Add(p)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		os.products = pr
// 		return nil
// 	}
// }
