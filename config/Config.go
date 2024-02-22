package config

import "github.com/spf13/viper"

type Config struct {
	APIGATEWAYPORT  string `mapstructure:"APIGATEWAYPORT"`
	GRPCUSERPORT    string `mapstructure:"GRPCUSERPORT"`
	GRPCSERVICEPORT string `mapstructure:"GRPCSERVICEPORT"`
	GRPCADMINPORT   string `mapstructure:"GRPCADMINPORT"`
}

func LoadConfig() *Config {
	var config *Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return config
}
