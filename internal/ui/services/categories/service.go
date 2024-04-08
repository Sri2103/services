package categories_service

import (
	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/services/client"
	products_templ "github.com/Sri2103/services/internal/ui/views/products"
)

type CategoryService interface {
	GetCategories() ([]products_templ.ProductCategory, error)
}

type service struct {
	AllClients *client.ClientAggregator
}

// GetCategories implements CategoryService.
func (s *service) GetCategories() ([]products_templ.ProductCategory, error) {
	panic("unimplemented")
}

func New(cfg *config.AppConfig) CategoryService {
	switch cfg.DevConfig.UseApi {
	case false:
		return NewMockService()

	case true:
		return &service{
			AllClients: client.AllClients(client.New(cfg)),
		}
	default:
		return nil
	}
}
