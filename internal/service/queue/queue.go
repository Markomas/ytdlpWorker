package queue

import (
	"github.com/adjust/rmq/v5"
)

type Client struct {
	rmqConnection rmq.Connection
}

func NewClient(rmqConnection rmq.Connection) *Client {
	return &Client{rmqConnection: rmqConnection}
}

func (client *Client) AddToQueue(queue string, jobPayload []byte) error {
	jobs, err := client.rmqConnection.OpenQueue(queue)
	if err != nil {
		return err
	}

	err = jobs.PublishBytes(jobPayload)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) PurgeAll() (int64, error) {
	var count int64
	count = 0
	queues, err := client.rmqConnection.GetOpenQueues()
	if err != nil {
		return count, err
	}

	for _, queueName := range queues {
		queue, err := client.rmqConnection.OpenQueue(queueName)
		if err != nil {
			continue
		}

		countReady, _ := queue.PurgeReady()
		count += countReady
		countRejected, _ := queue.PurgeRejected()
		count += countRejected
	}

	return count, nil
}

func (client *Client) GetStats() (rmq.Stats, error) {
	queues, err := client.rmqConnection.GetOpenQueues()
	if err != nil {
		panic(err)
	}

	return client.rmqConnection.CollectStats(queues)
}
