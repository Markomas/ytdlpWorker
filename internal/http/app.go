package http

import (
	"encoding/json"
	"fmt"
	"github.com/Markomas/ytdlpWorker/internal/service/queue"
	"github.com/adjust/rmq/v5"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

type App struct {
	config        *Config
	rmqConnection *rmq.Connection
	queueClient   *queue.Client
}

type Services struct {
	RmqConnection *rmq.Connection
	QueueClient   *queue.Client
}

func NewApp(config Config, services Services) *App {
	return &App{
		config:        &config,
		rmqConnection: services.RmqConnection,
		queueClient:   services.QueueClient,
	}
}

func (app *App) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/queue/add-to-queue", app.handleQueueAdd)
	mux.HandleFunc("/queue/purge-queue", app.handleQueuePurge)
	mux.Handle("/queue/overview", NewQueueOverviewHandler(*app.queueClient))

	fmt.Printf("Starting server on port %s:%d\n", app.config.Host, app.config.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", app.config.Host, app.config.Port), mux)
	if err != nil {
		return err
	}

	return nil
}

func errorMessage(w http.ResponseWriter, err error) {

	log.Error(err)

	message := JsonMessage{
		Response: "",
		Error:    err.Error(),
	}

	jsonMessage, err := json.Marshal(message)

	if err != nil {
		return
	}

	_, err = w.Write(jsonMessage)
	if err != nil {
		return
	}
}
