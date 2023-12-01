package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Redis *redis.Config `mapstructure:"redis"`
}

func LoadConfig(path string) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config Config

	err = viper.Unmarshal(&config)

	fmt.Println(config.Host)
}
