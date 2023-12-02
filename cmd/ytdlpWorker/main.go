package main

import (
	"flag"
	"github.com/Markomas/ytdlpWorker/internal/config"
	"github.com/Markomas/ytdlpWorker/internal/http"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yml", "Path to the configuration file")
	flag.Parse()

	conf, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	var app *http.App
	app = http.NewApp(conf.Http)
	app.Run()
}
