package redis

type Config struct {
	Host string `default:"localhost"`
	Port int    `default:"63791"`
}
