package config

import "github.com/spf13/viper"

type AppConfig struct {
	ApiServer *ApiServer `json:"product"`
	DevConfig *DevConfig `json:"dev_config"`
}
type ApiServer struct {
	Port string `mapstructure:"port" json:"port"`
	Url  string `mapstructure:"url" json:"url"`
}

type DevConfig struct {
	UseApi bool `mapstructure:"use_api" json:"use_api"`
}

func GetAppConfig() (*AppConfig, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := &AppConfig{}

	var pCfg = ApiServer{}
	err = viper.UnmarshalKey("api-server", &pCfg)
	if err != nil {
		return nil, err
	}
	conf.ApiServer = &pCfg

	var DevConfig DevConfig = DevConfig{UseApi: false}
	viper.Unmarshal(&DevConfig)
	conf.DevConfig = &DevConfig

	return conf, nil

}
