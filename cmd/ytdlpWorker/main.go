package main

import (
	"flag"
	"ytdlpWorker/internal/config"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yml", "Path to the configuration file")
	flag.Parse()

	config.LoadConfig(configPath)
}
