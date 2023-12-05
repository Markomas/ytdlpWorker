package http

import (
	"fmt"
	"github.com/adjust/rmq/v5"
	"log"
	"net/http"
)

type QueueOverviewHandler struct {
	connection rmq.Connection
}

func NewQueueOverviewHandler(connection rmq.Connection) *QueueOverviewHandler {
	return &QueueOverviewHandler{connection: connection}
}

func (handler *QueueOverviewHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	layout := request.FormValue("layout")
	refresh := request.FormValue("refresh")

	queues, err := handler.connection.GetOpenQueues()
	if err != nil {
		panic(err)
	}

	stats, err := handler.connection.CollectStats(queues)
	if err != nil {
		panic(err)
	}

	log.Printf("queue stats\n%s", stats)
	_, err = fmt.Fprint(writer, stats.GetHtml(layout, refresh))
	if err != nil {
		return
	}
}
