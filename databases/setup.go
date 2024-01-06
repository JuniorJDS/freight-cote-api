package databases

import (
	"context"
	"freight-cote-api/configs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := configs.GetSettings()["MONGO_URI"]
	mongoDatabase := configs.GetSettings()["MONGO_DATABASE"]

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Printf("Error to connect to MongoDB: %s\n", err.Error())
	}

	database := client.Database(mongoDatabase)
	return database
}

func GetDB() *mongo.Database {
	if db == nil {
		db = ConnectDB()
	}

	return db
}

func GetCollection(collectionName string) *mongo.Collection {
	database := GetDB()
	return database.Collection(collectionName)
}
