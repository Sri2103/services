package product_service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/services/client"
	"github.com/Sri2103/services/internal/ui/views/components"
)

type ProductService struct {
	AllClients *client.ClientAggregator
}

func New(cfg *config.AppConfig) *ProductService {
	return &ProductService{
		AllClients: client.AllClients(client.New(cfg)),
	}
}

// Get all products
func (s *ProductService) GetProducts() error {
	req := s.AllClients.ProductClient.NewRequest()
	res, err := req.Get("/all")
	if err != nil {
		return err
	}
	if res.IsError() {
		return errors.New(res.String())
	}
	fmt.Println("Products: ", res.String())
	return nil
}

func (s *ProductService) AddProduct(p components.Product) error {
	price, err := strconv.Atoi(p.ProductPrice)
	if err != nil {
		return err
	}

	req := s.AllClients.ProductClient.NewRequest().SetBody(Product{
		Name:        p.ProductName,
		Description: p.ProductPrice,
		Price:       float32(price),
	})
	res, err := req.Post("")
	if err != nil {
		return fmt.Errorf("could not add product: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("%v", res.StatusCode())
	}
	fmt.Println(res.String())

	return nil
}
