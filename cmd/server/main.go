package main

import (
	"demo/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	// Set Gin to release mode
	// gin.SetMode(gin.ReleaseMode)
}

func main() {
	// Create a new Gin router with default middleware
	router := gin.Default()

	// Configure trusted proxies
	// For development, we can trust only loopback addresses (127.0.0.1/8)
	// For production, you should set this to your actual proxy IPs
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// User routes
	users := router.Group("/users")
	{
		users.POST("", handlers.CreateUser)
		users.GET("", handlers.ListUsers)
		users.GET("/:id", handlers.GetUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	// Start the server
	fmt.Println("Server starting on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
