package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"sth/pkg/config"

	"sth/internal/worker-service/queue"
	"syscall"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
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
