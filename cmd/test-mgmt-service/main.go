package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"sth/internal/queue"
)

func main() {
	queueService, err := queue.NewQueueService("amqp://user:pass@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer queueService.Close()

	err = queueService.CreateTestQueue()
	if err != nil {
		log.Fatal("Failed to create test queue:", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/tests", func(c *gin.Context) {
		var req struct {
			TestName string `json:"test_name" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err := queueService.Publish("tests_queue", []byte(req.TestName))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue test"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Test enqueued successfully"})
	})

	log.Println("Test Management Service started on :8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatal("Failed to start Test Management Service:", err)
	}
}
