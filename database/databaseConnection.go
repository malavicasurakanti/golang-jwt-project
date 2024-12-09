package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Initialize MongoDB instance
func DBinstance() *mongo.Client {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get MongoDB URI from environment variable
	MongoDB := os.Getenv("MONGO_URL")
	if MongoDB == "" {
		log.Fatal("MONGO_URL not found in environment")
	}

	// Create a new MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	// Connect to MongoDB with a context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Verify connection with a Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func init() {
	// Initialize the global MongoDB client when the package is imported
	Client = DBinstance()
}

// OpenCollection opens a collection from the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Use the database name from your environment or hardcode it
	databaseName := os.Getenv("DB_NAME")
	if databaseName == "" {
		databaseName = "default_db" // Replace this with your database name
	}

	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
