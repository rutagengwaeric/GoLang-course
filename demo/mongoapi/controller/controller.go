package controller

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://rutagengwaeric250:kTJQ3fr0Cr3zahXH@cluster0.xg8vqja.mongodb.net/"

const databaseName = "netflix"
const collectionName = "watchlist"

// MOST IMPORTANT

var collection *mongo.Collection

func init() {

	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	fmt.Println("connected to mongodb")

	// collection reference

	collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("collection reference is ready")

}

func main() {

}
