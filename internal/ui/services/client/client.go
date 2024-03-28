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
}

func AllClients(api *ApiClient) *ClientAggregator {
	return &ClientAggregator{
		ProductClient: ProductClient(api),
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
