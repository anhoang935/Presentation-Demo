package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

// InitMongoDB initializes the MongoDB database connection
func InitMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("MONGODB_DATABASE")

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	// Verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("error pinging MongoDB: %w", err)
	}

	MongoDB = client.Database(database)
	log.Println("✅ MongoDB connected successfully")
	return nil
}

// CloseMongoDB closes the MongoDB connection
func CloseMongoDB() {
	if MongoDB != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := MongoDB.Client().Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		} else {
			log.Println("MongoDB connection closed")
		}
	}
}
