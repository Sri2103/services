package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server        *ServerConfig  `mapstructure:"gateway_server"`
	ProductConfig *ProductConfig `mapstructure:"product_config"`
	CartConfig    *CartConfig    `mapstructure:"cart_config"`
	OrderConfig   *OrderConfig   `mapstructure:"order_config"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type ProductConfig struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	Port int    `json:"port" yaml:"port" mapstructure:"port"`
}

type CartConfig struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	Port int    `json:"port" yaml:"port" mapstructure:"port"`
}
type OrderConfig struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	Port int    `json:"port" yaml:"port" mapstructure:"port"`
}

func ReadConfig() *AppConfig {
	viper.SetConfigType("yaml") //
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var conf = &AppConfig{}
	var server ServerConfig
	err = viper.UnmarshalKey("gateway_server", &server)
	if err != nil {
		fmt.Printf("%v not found! err: %s", "server", err.Error())
		panic(err)
	}

	conf.Server = &server

	var product ProductConfig
	err = viper.UnmarshalKey("product_config", &product)
	if err != nil {
		fmt.Printf("%v config not found! err: %s", "product", err.Error())
		panic(err)
	}
	conf.ProductConfig = &product

	var cart CartConfig

	err = viper.UnmarshalKey("cart_config", &cart)

	if err != nil {
		fmt.Printf("%v config not found! err: %s", "cart", err.Error())
		panic(err)
	}

	conf.CartConfig = &cart

	var order OrderConfig
	err = viper.UnmarshalKey("order_config", &order)
	if err != nil {
		fmt.Printf("%v config not found! err: %s", "cart", err.Error())
		panic(err)
	}
	conf.OrderConfig = &order

	return conf

}
