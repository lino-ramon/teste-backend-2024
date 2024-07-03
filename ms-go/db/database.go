package db

import (
	"context"
	"fmt"
	"ms-go/config/logger"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Collection
	client *mongo.Client
	MONGO_HOST 		= os.Getenv("MONGO_HOST")
	MAIN_COLLECTION = os.Getenv("MAIN_COLLECTION")
	MONGO_DATABASE  = os.Getenv("MONGO_DATABASE")
)

func Connection() *mongo.Collection {
	var err error

	logger.Debug("Variaveis de ambiente: MONGO_HOST", MONGO_HOST)
	logger.Debug("Variaveis de ambiente: MONGO_HOST", MONGO_DATABASE)
	logger.Debug("Variaveis de ambiente: MONGO_HOST", MAIN_COLLECTION)
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + MONGO_HOST +":27017")

	// Connect to MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("MONGO: ", err)
		return nil
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("MONGO: ", err)
		return nil
	}

	// Set the database and collection variables
	db = client.Database(MONGO_DATABASE).Collection(MAIN_COLLECTION)

	return db
}

func Disconnect() {
	client.Disconnect(context.TODO())
}
