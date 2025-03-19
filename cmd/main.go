package main

import (
	"fmt"
	"food-order-payments/config"
	"food-order-payments/internal/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("🔄 Initializing MongoDB Connection...")

	// ✅ Call ConnectMongoDB() at startup
	config.ConnectMongoDB()

	fmt.Println("✅ MongoDB initialization completed.")

	// Initialize Gin router
	router := gin.Default()

	// Health check route
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "API is running"})
	})

	// Order Routes
	router.POST("/order", handlers.CreateOrderHandler)
	router.GET("/order/:id", handlers.GetOrderHandler)

	port := ":8080"
	fmt.Println("🚀 API is running on http://localhost" + port)
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
