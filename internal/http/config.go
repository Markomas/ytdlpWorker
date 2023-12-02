package http

type Config struct {
	Port int    `mapstructure:"port" default:"8888"`
	Host string `mapstructure:"host" default:"0.0.0.0"`
}
