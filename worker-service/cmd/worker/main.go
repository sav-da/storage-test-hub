package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"worker-service/internal/config"
	"worker-service/internal/queue"
)

func main() {
	cfg := config.LoadConfig()
	cfg.Validate()

	consumer, err := queue.NewConsumer(cfg.RabbitMQURL, cfg.QueueName)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ consumer: %v", err)
	}
	defer consumer.Close()

	go func() {
		if err := consumer.Start(); err != nil {
			log.Fatalf("Consumer error: %v", err)
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Worker shutting down")
}
