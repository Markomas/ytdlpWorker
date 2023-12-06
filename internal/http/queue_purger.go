package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (app *App) handleQueuePurge(w http.ResponseWriter, _ *http.Request) {
	count, err := app.queueClient.PurgeAll()
	if err != nil {
		errorMessage(w, err)
		return
	}

	message := JsonMessage{
		Response: fmt.Sprintf("Purged %d messages from queue", count),
		Error:    "",
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = w.Write(jsonMessage)
	if err != nil {
		return
	}
}
