package product_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/Sri2103/services/internal/ui/config"
	categories_service "github.com/Sri2103/services/internal/ui/services/categories"
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
	GetProductsByCategory(category string, pageNumber int, pageSize int, sort string) ([]components.Product, int,int, error)
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
	type Body struct {
		Product  Product                     `json:"product"`
		Category categories_service.Category `json:"category"`
	}

	req := s.AllClients.ProductClient.NewRequest().SetBody(Body{
		Product: Product{
			Name:        p.ProductName,
			Description: p.ProductDescription,
			Price:       float32(price),
			Images:      p.ProductImages,
			Colors:      p.ProductColor,
		},
		Category: categories_service.Category{
			Id: p.ProductCategory,
		},
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
func (s *service) GetProductsByCategory(category string, pageNumber int, pageSize int, sort string) ([]components.Product, int, int, error) {
	req := s.AllClients.ProductClient.NewRequest()
	req = req.SetQueryParams(map[string]string{
		"page":     strconv.Itoa(pageNumber),
		"pageSize": strconv.Itoa(pageSize),
		"sort":     sort,
	})

	res, err := req.Get("/products-category/" + category)
	if err != nil {
		return nil, 0, 0, err
	}

	if res.IsError() {
		return nil, 0, 0, fmt.Errorf("%v", res.StatusCode())
	}

	type respData struct {
		Products     []Product `json:"products"`
		TotalPages   int       `json:"total_pages"`
		TotalResults int       `json:"total_results"`
	}

	var d respData

	err = json.Unmarshal(res.Body(), &d)
	if err != nil {
		return nil, 0, 0, err
	}
	products_comp := make([]components.Product, len(d.Products))
	for i, v := range d.Products {
		products_comp[i] = convertToComponentModal(v)
	}
	return products_comp, d.TotalPages, d.TotalResults, nil
}

// get ProductById
func (s *service) GetProductById(id string) (components.Product, error) {
	req := s.AllClients.ProductClient.NewRequest()
	res, err := req.Get("/" + id)
	if err != nil {
		return components.Product{}, err
	}
	product := Product{}
	err = json.Unmarshal(res.Body(), &product)
	if err != nil {
		return components.Product{}, err
	}
	return convertToComponentModal(product), nil

}
