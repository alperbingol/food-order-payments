package handlers

import (
	"net/http"

	"food-order-payments/internal/models"
	"food-order-payments/internal/repositories"

	"github.com/gin-gonic/gin"
)

// CreateOrderHandler handles POST /order
func CreateOrderHandler(c *gin.Context) {
	var order models.Order

	// Bind JSON request to Order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Save order in MongoDB
	savedOrder, err := repositories.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, savedOrder)
}

// GetOrderHandler handles GET /order/:id
func GetOrderHandler(c *gin.Context) {
	orderID := c.Param("id")

	// Fetch order from MongoDB
	order, err := repositories.GetOrderById(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
