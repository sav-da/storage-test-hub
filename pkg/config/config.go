package config

import (
	"log"
	"os"
)

type Config struct {
	RabbitMQURL  string
	QueueName    string
	AuthService  string
	TestService  string
	QueueService string
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

	authService := os.Getenv("AUTH_SERVICE")
	if authService == "" {
		authService = "http://localhost:8081"
	}

	testService := os.Getenv("TEST_SERVICE")
	if testService == "" {
		testService = "http://localhost:8082"
	}

	queueService := os.Getenv("QUEUE_SERVICE")
	if queueService == "" {
		queueService = "http://localhost:8083"
	}

	return &Config{
		RabbitMQURL:  rabbitMQURL,
		QueueName:    queueName,
		AuthService:  authService,
		TestService:  testService,
		QueueService: queueService,
	}
}

func (c *Config) Validate() {
	if c.RabbitMQURL == "" || c.QueueName == "" {
		log.Fatalf("Invalid configuration: RabbitMQURL and QueueName must be set")
	}
}
