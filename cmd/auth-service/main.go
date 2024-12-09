package main

import (
	"log"
	"net/http"
	"sth/internal/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/auth/token", func(c *gin.Context) {
		var req struct {
			WorkerID string `json:"worker_id" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		token, err := auth.GenerateToken(req.WorkerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	log.Println("Auth Service started on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to start Auth Service:", err)
	}
}
