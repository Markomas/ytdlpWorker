package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Job struct {
	Url string
}

const DownloadQueueName = "download_queue"

func (app *App) handleQueueAdd(w http.ResponseWriter, r *http.Request) {
	var job Job

	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		errorMessage(w, &UnexpectedError{err: err})
		return
	}

	jobBytes, err := json.Marshal(job)
	if err != nil {
		errorMessage(w, &UnexpectedError{err: err})
		return
	}

	err = app.queueClient.AddToQueue(DownloadQueueName, jobBytes)
	if err != nil {
		errorMessage(w, &UnexpectedError{err: err})
		return
	}

	message := JsonMessage{
		Response: "Added job to queue",
		Error:    "",
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = w.Write(jsonMessage)
	if err != nil {
		log.Fatal(err)
		return
	}
}
