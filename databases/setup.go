package databases

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := "4hub-sop"                                                                                              // configs.GetSettings()["MONGO_URI"]
	mongoDatabase := "mongodb://mongoadmin:secret@localhost:27017/sop-mongo?authSource=admin&authMechanism=SCRAM-SHA-1" // configs.GetSettings()["MONGO_DATABASE"]
	// connectionString := "mongodb://mongoadmin:secret@localhost:27017/?authSource=admin&authMechanism=SCRAM-SHA-1"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDatabase))
	// client, err := mongo.NewClient(options.Client().ApplyURI(mongoDatabase))
	if err != nil {
		log.Printf("Error to connect to MongoDB: %s\n", err.Error())
	}

	// defer func() {
	// 	if err := client.Disconnect(ctx); err != nil {
	// 		log.Printf("Disconnecting MongoDB: %s\n", err.Error())
	// 	}
	// }()

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Printf("Error to connect to MongoDB: %s\n", err.Error())
	// }

	database := client.Database(mongoURI)
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
