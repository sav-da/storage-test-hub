package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	// Dummy implementation
	c.JSON(http.StatusOK, map[string]string{"token": "dummy-jwt-token"})

}

func ValidateHandler(c *gin.Context) {

	// Dummy implementation
	c.JSON(http.StatusOK, map[string]string{"token": "dummy-jwt-token"})

}
