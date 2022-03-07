package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!!!>>**<<",
		})
	})

	// A publicly open (no authorization) endpoint to verify the health of the application
	// Here it always returns 200, but in real cases it should properly indicate if the application is actually healthy or not
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Healthy")
	})

	r.Run(":" + port)
}
