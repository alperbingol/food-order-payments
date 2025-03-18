package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Gin router
	router := gin.Default()

	// Health check route
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "API is running"})
	})

	// Placeholder for order & payment routes
	router.POST("/order", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Order created"})
	})

	router.POST("/pay", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Payment processed"})
	})

	port := ":8080"
	fmt.Println("🚀 API is running on http://localhost" + port)
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
