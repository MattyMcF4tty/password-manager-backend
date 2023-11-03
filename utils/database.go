package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDb() (*mongo.Client, error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// Apply the username and password from the ser
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@go-x-mongodb.pzqixft.mongodb.net/?retryWrites=true&w=majority", username, password)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("Failed to connect:", err)
		return nil, err
	}
	return client, nil
}

func disconnectFromDb(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Println("Failed to disconnect:", err)
	}
}
