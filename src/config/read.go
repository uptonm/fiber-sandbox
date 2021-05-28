package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Environment string

var (
	Development Environment = "local"
	Production  Environment = "production"
)

type IngressConfiguration struct {
	Host        string   `mapstructure:"host"`
	Port        string   `mapstructure:"port"`
	CorsOrigins []string `mapstructure:"cors_origins"`
}

type Configuration struct {
	Ingress IngressConfiguration `mapstructure:"ingress"`
}

// ReadConfig utilizes viper to read a .yml config based on environment, and return a pointer to it
func ReadConfig(env Environment) *Configuration {
	var configuration Configuration
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("failed to read config: %s\n", err.Error())
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Panicf("failed to unmarshal config: %s\n", err.Error())
	}

	return &configuration
}
