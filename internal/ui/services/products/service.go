package product_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/services/client"
	"github.com/Sri2103/services/internal/ui/views/components"
)

type service struct {
	AllClients *client.ClientAggregator
}

// AddCategory implements ProductService.
func (s *service) AddCategory(category string) error {
	panic("unimplemented")
}

type ProductService interface {
	GetProducts() ([]components.Product, error)
	AddProduct(product components.Product) error
	UpdateProduct(id string, product components.Product) (components.Product, error)
	GetProductsByCategory(category string, pageNumber int, pageSize int, sort string) ([]components.Product, int, error)
	GetProductById(id string) (components.Product, error)
	AddCategory(category string) error
}

func New(cfg *config.AppConfig) ProductService {

	switch cfg.DevConfig.UseApi {
	case true:
		return NewMockService()

	case false:
		return &service{
			AllClients: client.AllClients(client.New(cfg)),
		}
	default:
		return nil
	}
}

// Get all products
func (s *service) GetProducts() ([]components.Product, error) {
	req := s.AllClients.ProductClient.NewRequest()
	res, err := req.Get("/all")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.String())
	}
	var data []Product
	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return nil, fmt.Errorf("cannot marshall the data" + err.Error())
	}
	products := make([]components.Product, len(data))
	for i, v := range data {
		products[i] = convertToComponentModal(v)
	}
	return products, nil
}

func (s *service) AddProduct(p components.Product) error {
	price, err := strconv.Atoi(p.ProductPrice)
	if err != nil {
		return err
	}

	req := s.AllClients.ProductClient.NewRequest().SetBody(Product{
		Name:        p.ProductName,
		Description: p.ProductDescription,
		Price:       float32(price),
		Images:      p.ProductImages,
		Colors:      p.ProductColor,
	})
	res, err := req.Post("")
	if err != nil {
		return fmt.Errorf("could not add product: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("%v", res.StatusCode())
	}

	return nil
}

// Update product
func (s *service) UpdateProduct(id string, p components.Product) (components.Product, error) {
	price, err := strconv.Atoi(p.ProductPrice)
	if err != nil {
		return components.Product{}, err
	}

	product := Product{
		Id:          id,
		Name:        p.ProductName,
		Description: p.ProductPrice,
		Price:       float32(price),
		Images:      p.ProductImages,
		Colors:      p.ProductColor,
	}

	updateReq := s.AllClients.ProductClient.NewRequest().SetBody(product)
	res, err := updateReq.Put("")
	if err != nil {
		return components.Product{}, err
	}
	if res.IsError() {
		return components.Product{}, fmt.Errorf("%v", res.StatusCode())
	}
	var updatedProduct Product
	err = json.Unmarshal(res.Body(), &updatedProduct)
	if err != nil {
		return components.Product{}, err
	}

	return convertToComponentModal(updatedProduct), nil
}

// GetProductsByCategory implements ProductService.
func (s *service) GetProductsByCategory(category string, pageNumber int, pageSize int, sort string) ([]components.Product, int, error) {
	req := s.AllClients.ProductClient.NewRequest()
	req = req.SetQueryParams(map[string]string{
		"page":     strconv.Itoa(pageNumber),
		"pageSize": strconv.Itoa(pageSize),
		"sort":     sort,
	})

	res, err := req.Get("/product-category/" + category)
	if err != nil {
		return nil, 0, err
	}

	if res.IsError() {
		return nil, 0, fmt.Errorf("%v", res.StatusCode())
	}

	var products []Product

	err = json.Unmarshal(res.Body(), &products)
	if err != nil {
		return nil, 0, err
	}
	products_comp := make([]components.Product, len(products))
	for i, v := range products {
		products_comp[i] = convertToComponentModal(v)
	}
	return products_comp, 0, nil
}

// get ProductById
func (s *service) GetProductById(id string) (components.Product, error) {
	panic("unimplemented")
}
