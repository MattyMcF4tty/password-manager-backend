package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name string `bson:"name"`
}
  
func main() {
    // Use the SetServerAPIOptions() method to set the Stable API version to 1
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI("mongodb+srv://matthiaskristensen:Lule1167@go-x-mongodb.pzqixft.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

    // Create a new client and connect to the server
    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
        panic(err)
        }
    }()

    // Send a ping to confirm a successful connection
    if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
        panic(err)
    }
    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

    	// Get a handle to the collection
	db := client.Database("Users")
    filter := bson.D{{"age", 26}}

    collection := db.Collection("User")

    var result User
    err = collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            fmt.Println("No documents found with filter:", filter)
            return
        }
        panic(err)
    }

    fmt.Println("Name of user with age 26 is:", result.Name)
}
  
