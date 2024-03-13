package http_transport

import (
	"fmt"

	"github.com/Sri2103/services/internal/api-gateway/dependency"
	product_handler "github.com/Sri2103/services/internal/api-gateway/services/products/handlers"
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

	
}
