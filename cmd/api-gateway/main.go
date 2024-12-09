package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Прокси для Auth Service
	router.POST("/auth/token", proxyRequest("http://localhost:8081/auth/token"))

	// Прокси для Test Management Service
	router.POST("/tests", proxyRequest("http://localhost:8082/tests"))

	router.Run(":8080")
}

func proxyRequest(targetURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := http.Post(targetURL, "application/json", c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
			return
		}
		defer resp.Body.Close()

		c.Status(resp.StatusCode)
		c.Header("Content-Type", resp.Header.Get("Content-Type"))
		c.Writer.WriteHeaderNow()
	}
}
