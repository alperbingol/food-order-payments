package config

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB       *mongo.Database
	dbClient *mongo.Client
	once     sync.Once
)

// ConnectMongoDB initializes MongoDB connection
func ConnectMongoDB() {
	fmt.Println("🔄 Attempting to connect to MongoDB...") // Debugging Log

	once.Do(func() {
		fmt.Println("🟢 Inside once.Do() - Running MongoDB Connection Logic")

		mongoURI := "mongodb://localhost:27017"

		clientOptions := options.Client().ApplyURI(mongoURI)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("❌ MongoDB Connection Failed:", err)
		}

		// Ping MongoDB
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal("❌ MongoDB Ping Failed:", err)
		}

		// Assign global database instance
		DB = client.Database("food_order_payments")
		dbClient = client

		fmt.Println("✅ MongoDB Connected Successfully!")
	})
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	fmt.Println("🔍 Checking MongoDB initialization...") // Debugging Log

	if DB == nil {
		log.Fatal("❌ ERROR: MongoDB is not initialized. Call ConnectMongoDB() first.")
	}

	fmt.Println("✅ MongoDB is initialized, returning collection:", collectionName)
	return DB.Collection(collectionName)
}
