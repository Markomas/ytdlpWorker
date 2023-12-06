package http

import (
	"fmt"
	"github.com/Markomas/ytdlpWorker/internal/service/queue"
	"log"
	"net/http"
)

type QueueOverviewHandler struct {
	client queue.Client
}

func NewQueueOverviewHandler(connection queue.Client) *QueueOverviewHandler {
	return &QueueOverviewHandler{client: connection}
}

func (handler *QueueOverviewHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	layout := request.FormValue("layout")
	refresh := request.FormValue("refresh")

	stats, err := handler.client.GetStats()
	if err != nil {
		errorMessage(writer, err)
		return
	}

	log.Printf("queue stats\n%s", stats)
	_, err = fmt.Fprint(writer, stats.GetHtml(layout, refresh))
	if err != nil {
		return
	}
}
