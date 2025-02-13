package db

import (
	"context"
	"log"
	"time"

	"example.com/bbb/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client               *mongo.Client
	Database             *mongo.Database
	RestaurantCollection *mongo.Collection
	ReviewsCollection    *mongo.Collection
}

var globalDB *DB

func InitDB() error {

	mongoURI := utils.GetMongoURI()
	dbName := utils.GetMongoDatabaseName()
	restauntCollectionName := utils.GetMongoDatabaseCollection("RESTAURANT_COLLECTION")
	reviewCollectionName := utils.GetMongoDatabaseCollection("REVIEW_COLLECTION")

	ctx, cancel := initDBContext()
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error Creating MongoDB Client: %v", err)
	}

	database := client.Database((dbName))
	restauntCollection := database.Collection(restauntCollectionName)
	reviewCollection := database.Collection(reviewCollectionName)

	globalDB = &DB{
		Client:               client,
		Database:             database,
		RestaurantCollection: restauntCollection,
		ReviewsCollection:    reviewCollection,
	}
	return nil
}

func initDBContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return ctx, cancel
}

func GetClient() *mongo.Client {
	return globalDB.Client
}
