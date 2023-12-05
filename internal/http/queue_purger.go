package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (app *App) handleQueuePurge(w http.ResponseWriter, r *http.Request) {
	queue, err := app.rmqConnection.OpenQueue(DownloadQueueName)
	if err != nil {
		errorMessage(w, err)
		return
	}

	countRejected, err := queue.PurgeRejected()
	if err != nil {
		errorMessage(w, err)
		return
	}

	countReady, err := queue.PurgeReady()
	if err != nil {
		errorMessage(w, err)
		return
	}

	message := JsonMessage{
		Response: fmt.Sprintf("Purged %d messages from queue", countReady+countRejected),
		Error:    "",
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(jsonMessage)
}
