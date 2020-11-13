package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DB  *mongo.Database
var DB *mongo.Database

// InitializeDB initializes DB
func InitializeDB() {
	clientOptions := options.Client().ApplyURI(Config.Database.URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Failed to create client: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Config.Database.Timeout)*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to initialize client: ", err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	DB = client.Database(Config.Database.DB)
}
