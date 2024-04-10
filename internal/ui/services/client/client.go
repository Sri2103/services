package client

import (
	"fmt"

	"github.com/Sri2103/services/internal/ui/config"
	"github.com/go-resty/resty/v2"
)

type ApiClient struct {
	BaseUrl string
	Port    string
}

type ClientAggregator struct {
	ProductClient *resty.Client
	UserClient    *resty.Client
	CartClient    *resty.Client
}

func AllClients(api *ApiClient) *ClientAggregator {
	return &ClientAggregator{
		ProductClient: ProductClient(api),
		UserClient:    UserClient(api),
		CartClient:    CartClient(api),
	}
}

func New(cfg *config.AppConfig) *ApiClient {
	return &ApiClient{
		BaseUrl: cfg.ApiServer.Url,
		Port:    cfg.ApiServer.Port,
	}
}

func ProductClient(api *ApiClient) *resty.Client {
	client := resty.New().SetBaseURL(fmt.Sprintf("%s:%s/products", api.BaseUrl, api.Port))
	// Add some default headers to all requests
	return client
}

func UserClient(api *ApiClient) *resty.Client {
	client := resty.New().SetBaseURL(fmt.Sprintf("%s:%s/users", api.BaseUrl, api.Port))
	// Add some default headers to all requests
	return client
}

// CartClient
func CartClient(api *ApiClient) *resty.Client {
	client := resty.New().SetBaseURL(fmt.Sprintf("%s:%s/carts", api.BaseUrl, api.Port))
	// Add some default headers to all requests
	return client
}
