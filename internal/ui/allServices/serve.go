package handlerServices

import (
	"github.com/Sri2103/services/internal/ui/config"
	product_service "github.com/Sri2103/services/internal/ui/services/products"
	user_service "github.com/Sri2103/services/internal/ui/services/user"
)

type Services struct {
	UserService user_service.UserService
	ProductService product_service.ProductService
}

func New(cfg *config.AppConfig) *Services {

	return &Services{
		UserService: user_service.New(cfg),
		ProductService: product_service.New(cfg),
	}
}
