package handlers

import (
	"github.com/Sri2103/services/internal/api-gateway/dependency"
	"github.com/Sri2103/services/pkg/rpc/cart"
)

type handlers struct {
	dep        *dependency.Dependency
	CartClient cart.CartServiceClient
}

func CartHandler(dep *dependency.Dependency) *handlers {
	client := cart.NewCartServiceClient(dep.CartConn)
	return &handlers{
		dep:        dep,
		CartClient: client,
	}
}



