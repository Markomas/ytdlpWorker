package config

import (
	"github.com/Markomas/ytdlpWorker/internal/http"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type Config struct {
	Name string      `mapstructure:"name"`
	Http http.Config `mapstructure:"http"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	defaults.SetDefaults(&config)

	return &config, nil
}
