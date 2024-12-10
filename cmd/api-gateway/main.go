package main

import (
	"log"
	"net/http"
	"sth/pkg/config"

	"github.com/gin-gonic/gin"
	"sth/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(logger.RequestLogger())

	// Прокси для Auth Service
	router.POST("/auth/login", proxyRequest(cfg.AuthService+"/login"))
	router.POST("/auth/validate", proxyRequest(cfg.AuthService+"/validate"))
	// Прокси для Test Management Service
	router.POST("/tests", proxyRequest(cfg.TestService+"/tests"))
	log.Fatal(router.Run(":8080"))
}

func proxyRequest(targetURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Proxying request to ", targetURL)
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
