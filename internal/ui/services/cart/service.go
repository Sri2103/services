package cart_service

import (
	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/services/client"
)

type Service interface {
	GetUserCart(userId string) (Cart, error)
	GetNonUserCart() (Cart, error)
	AddProductToCart(userId string, productId string, quantity int) (Cart, error)
	AddProductToNonUserCart(productId string, quantity int) (Cart, error)
	RemoveProductFromCart(userId string, productId string) (Cart, error)
	RemoveProductFromNonUserCart(productId string) (Cart, error)
	EmptyCart(userId string) (Cart, error)
	EmptyNonUserCart() (Cart, error)
	UpdateQuantity(userId string, productId string, quantity int) (Cart, error)
	UpdateNonUserQuantity(productId string, quantity int) (Cart, error)
}

type service struct {
	AllClients *client.ClientAggregator
}

// AddProductToCart implements Service.
func (s *service) AddProductToCart(userId string, productId string, quantity int) (Cart, error) {
	panic("unimplemented")
}

// AddProductToNonUserCart implements Service.
func (s *service) AddProductToNonUserCart(productId string, quantity int) (Cart, error) {
	panic("unimplemented")
}

// EmptyCart implements Service.
func (s *service) EmptyCart(userId string) (Cart, error) {
	panic("unimplemented")
}

// EmptyNonUserCart implements Service.
func (s *service) EmptyNonUserCart() (Cart, error) {
	panic("unimplemented")
}

// GetNonUserCart implements Service.
func (s *service) GetNonUserCart() (Cart, error) {
	panic("unimplemented")
}

// GetUserCart implements Service.
func (s *service) GetUserCart(userId string) (Cart, error) {
	panic("unimplemented")
}

// RemoveProductFromCart implements Service.
func (s *service) RemoveProductFromCart(userId string, productId string) (Cart, error) {
	panic("unimplemented")
}

// RemoveProductFromNonUserCart implements Service.
func (s *service) RemoveProductFromNonUserCart(productId string) (Cart, error) {
	panic("unimplemented")
}

// UpdateNonUserQuantity implements Service.
func (s *service) UpdateNonUserQuantity(productId string, quantity int) (Cart, error) {
	panic("unimplemented")
}

// UpdateQuantity implements Service.
func (s *service) UpdateQuantity(userId string, productId string, quantity int) (Cart, error) {
	panic("unimplemented")
}

func New(cfg *config.AppConfig) Service {
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
