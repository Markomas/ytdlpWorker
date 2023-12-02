package http

import (
	"fmt"
	"net/http"
)

type App struct {
	Config *Config
}

func NewApp(config Config) *App {
	return &App{
		Config: &config,
	}
}

func (app *App) Run() error {
	http.HandleFunc("/add-to-queue/add-to-queue", app.handleAddToQueue)

	fmt.Printf("Starting server on port %s:%d\n", app.Config.Host, app.Config.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", app.Config.Host, app.Config.Port), nil)
	if err != nil {
		return err
	}

	return nil
}
