package handlerServices

import (
	"github.com/Sri2103/services/internal/ui/config"
	categories_service "github.com/Sri2103/services/internal/ui/services/categories"
	product_service "github.com/Sri2103/services/internal/ui/services/products"
	user_service "github.com/Sri2103/services/internal/ui/services/user"
)

type Services struct {
	UserService     user_service.UserService
	ProductService  product_service.ProductService
	CategoryService categories_service.CategoryService
}

func New(cfg *config.AppConfig) *Services {

	return &Services{
		UserService:     user_service.New(cfg),
		ProductService:  product_service.New(cfg),
		CategoryService: categories_service.New(cfg),
	}
}
