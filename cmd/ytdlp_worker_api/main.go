package main

import (
	"flag"
	"fmt"
	"github.com/Markomas/ytdlpWorker/internal/config"
	"github.com/Markomas/ytdlpWorker/internal/http"
	"github.com/Markomas/ytdlpWorker/internal/service/queue"
	"github.com/adjust/rmq/v5"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yml", "Path to the configuration file")
	flag.Parse()

	conf, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	RmqConnection, err := rmq.OpenConnection("ytdlp_worker_api", "tcp", fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port), 2, nil)
	if err != nil {
		panic(err)
	}

	queueClient := queue.NewClient(RmqConnection)

	var services http.Services
	services.RmqConnection = &RmqConnection
	services.QueueClient = queueClient

	var app *http.App
	app = http.NewApp(conf.Http, services)

	err = app.Run()
	if err != nil {
		log.Error(err)
	}
}
