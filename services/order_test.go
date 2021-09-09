package services

import (
	"testing"

	"github.com/guilherme-fittipaldi/ddd-golang/aggregate"
	"github.com/guilherme-fittipaldi/ddd-golang/services"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Good", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("wine", "Good", 0.99)
	if err != nil {
		t.Error(err)
	}

	products := []aggregate.Product{beer, peenuts, wine}

	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
}