package main

import (
	"log"
	"sth/internal/queue-service/queue"
)

func main() {
	log.Println("Starting Queue Service...")
	queue.StartQueueConsumer()
}
