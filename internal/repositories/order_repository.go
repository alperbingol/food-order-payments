package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"food-order-payments/config"
	"food-order-payments/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Delay database initialization until needed
func getOrderCollection() *mongo.Collection {
	fmt.Println("🔍 Fetching MongoDB collection for orders...")
	return config.GetCollection("orders")
}

// CreateOrder saves an order in MongoDB
func CreateOrder(order models.Order) (models.Order, error) {
	order.ID = primitive.NewObjectID()
	order.Status = "pending"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getOrderCollection() // Fetch collection when needed

	_, err := collection.InsertOne(ctx, order)
	if err != nil {
		log.Println("❌ Error inserting order:", err)
		return order, err
	}

	log.Println("✅ Order inserted:", order.ID.Hex())
	return order, nil
}

// GetOrderById fetches an order by ID
func GetOrderById(orderID string) (models.Order, error) {
	var order models.Order
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		log.Println("❌ Invalid order ID format:", err)
		return order, err
	}

	collection := getOrderCollection() // Fetch collection when needed

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		log.Println("❌ Order not found:", err)
		return order, err
	}

	return order, nil
}
