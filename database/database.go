package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var ContestantsCollection *mongo.Collection

// kết nối tới MongoDB và tạo collection
func ConnectMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// kiểm tra kết nối
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	MongoClient = client
	db := client.Database("beautyPageant")
	ContestantsCollection = db.Collection("contestants")

	log.Println("Kết nối MongoDB thành công")
	return nil
}

// ngắt kết nối MongoDB
func DisconnectMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return MongoClient.Disconnect(ctx)
}
