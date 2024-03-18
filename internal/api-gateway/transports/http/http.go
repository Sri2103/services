package http_transport

import (
	"fmt"

	"github.com/Sri2103/services/internal/api-gateway/dependency"
	cart_handlers "github.com/Sri2103/services/internal/api-gateway/services/cart/handlers"
	product_handler "github.com/Sri2103/services/internal/api-gateway/services/products/handlers"
	user_handler "github.com/Sri2103/services/internal/api-gateway/services/user/handlers"
)

func StartHttpServer(dep *dependency.Dependency) error {
	// product handlers
	setUpRoutes(dep)
	port := fmt.Sprintf(":%d", dep.Config.Server.Port)
	err := dep.Server.Start(port)

	if err != nil {
		return err
	}

	return nil
}

func setUpRoutes(dep *dependency.Dependency) {
	productHandler := product_handler.New(dep)

	productHandler.SetProductRoutes()

	cartHandler := cart_handlers.CartHandler(dep)

	cartHandler.SetCartRoutes()

	userHandler := user_handler.New(dep)

	userHandler.SetUserRoutes()

}
