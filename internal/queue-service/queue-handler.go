package queue

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Handler struct {
	channel *amqp091.Channel
}

func NewQueueHandler(conn *amqp091.Connection) (*Handler, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Handler{channel: channel}, nil
}

func (q *Handler) Publish(queueName string, message []byte) error {
	_, err := q.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = q.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	return err
}

func (q *Handler) Close() {
	if q.channel != nil {
		err := q.channel.Close()
		if err != nil {
			log.Printf("Failed to close channel: %v", err)
		}
	}
}
