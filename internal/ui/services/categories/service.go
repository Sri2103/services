package categories_service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/services/client"
)

type CategoryService interface {
	GetCategories(ctx context.Context) ([]Category, error)
	AddCategory(ctx context.Context, category string) error
}

type service struct {
	AllClients *client.ClientAggregator
}

// AddCategory implements CategoryService.
func (s *service) AddCategory(ctx context.Context, category string) error {
	// route -> /addCategory
	req := s.AllClients.ProductClient.NewRequest().SetBody(Category{
		Name: category,
	})
	res, err := req.Post("/addCategory")
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("%v", res.StatusCode())
	}

	return nil
}

// GetCategories implements CategoryService.
func (s *service) GetCategories(_ context.Context) ([]Category, error) {
	req := s.AllClients.ProductClient.NewRequest()
	res, err := req.Get("/categories/all")
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("%v", res.StatusCode())
	}

	var categories []Category
	err = json.Unmarshal(res.Body(), &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func New(cfg *config.AppConfig) CategoryService {
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
