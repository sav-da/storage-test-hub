package config

import (
	"log"
	"os"
)

type Config struct {
	RabbitMQURL string
	QueueName   string
}

func LoadConfig() *Config {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://user:pass@localhost:5672/"
	}

	queueName := os.Getenv("QUEUE_NAME")
	if queueName == "" {
		queueName = "tests_queue"
	}

	return &Config{
		RabbitMQURL: rabbitMQURL,
		QueueName:   queueName,
	}
}

func (c *Config) Validate() {
	if c.RabbitMQURL == "" || c.QueueName == "" {
		log.Fatalf("Invalid configuration: RabbitMQURL and QueueName must be set")
	}
}
