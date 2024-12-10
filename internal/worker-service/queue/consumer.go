package queue

import (
	"log"
	"sth/internal/worker-service/task"

	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   string
}

func NewConsumer(rabbitMQURL, queueName string) (*Consumer, error) {
	conn, err := amqp091.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:    conn,
		channel: channel,
		queue:   queueName,
	}, nil
}

func (c *Consumer) Start() error {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", msg.Body)

			// Передаем задачу на обработку
			if err := task.ProcessTask(msg.Body); err != nil {
				log.Printf("Error processing task: %v", err)
			}
		}
	}()

	log.Printf("Consumer started, waiting for messages on queue: %s", c.queue)
	select {} // Keep running
}

func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
