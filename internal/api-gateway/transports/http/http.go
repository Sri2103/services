package http_transport

import (
	"fmt"

	"github.com/Sri2103/services/internal/api-gateway/dependency"
	product_handler "github.com/Sri2103/services/internal/api-gateway/services/products/handlers"
)

func StartHttpServer(dep *dependency.Dependency) error {
	// product handlers
	productHandler := product_handler.New(dep)

	productHandler.SetProductRoutes()
	port := fmt.Sprintf(":%d", dep.Config.Server.Port)
	err := dep.Server.Start(port)

	if err != nil {
		return err
	}

	return nil
}
