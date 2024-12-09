package queue

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Service struct {
	conn *amqp091.Connection
	ch   *amqp091.Channel
}

func NewQueueService(dsn string) (*Service, error) {
	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &Service{conn: conn, ch: ch}, nil
}

func (qs *Service) DeclareQueue(name string) error {
	_, err := qs.ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (qs *Service) Publish(queueName string, message []byte) error {
	return qs.ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}

func (qs *Service) Close() {
	err := qs.ch.Close()
	if err != nil {
		log.Println("Error closing channel:", err)
	}
	err = qs.conn.Close()
	if err != nil {
		log.Println("Error closing connection:", err)
	}
}
func (qs *Service) CreateTestQueue() error {
	return qs.DeclareQueue("tests_queue")
}
