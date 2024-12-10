package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sth/internal/auth-service/auth"
	"sth/pkg/logger"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(logger.RequestLogger())
	log.Println("Starting Auth Service...")
	r.Use(logger.RequestLogger())
	r.POST("/login", auth.LoginHandler)
	r.POST("/validate", auth.ValidateHandler)

	log.Fatal(r.Run(":8081"))
}
