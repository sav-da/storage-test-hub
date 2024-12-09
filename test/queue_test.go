package test

import (
	"log"
	"testing"

	"github.com/rabbitmq/amqp091-go"
)

func TestQueuePublish(t *testing.T) {
	conn, err := amqp091.Dial("amqp://user:pass@localhost:5672/")
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"tests_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to declare queue: %v", err)
	}

	message := "Test Message"
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		t.Fatalf("Failed to publish message: %v", err)
	}

	log.Printf("Message published to queue %s: %s", queue.Name, message)
}
