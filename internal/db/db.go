package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/bbb/internal/models"
	"example.com/bbb/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client             *mongo.Client
	Database           *mongo.Database
	BusinessCollection *mongo.Collection
	ReviewsCollection  *mongo.Collection
}

var globalDB *DB

func InitDB() error {

	mongoURI := utils.GetMongoURI()
	dbName := utils.GetMongoDatabaseName()
	businessCollectionName := utils.GetMongoDatabaseCollection("BUSINESS_COLLECTION")
	reviewCollectionName := utils.GetMongoDatabaseCollection("REVIEW_COLLECTION")

	ctx, cancel := initDBContext()
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error Creating MongoDB Client: %v", err)
	}

	database := client.Database((dbName))
	businessCollection := database.Collection(businessCollectionName)
	reviewCollection := database.Collection(reviewCollectionName)

	globalDB = &DB{
		Client:             client,
		Database:           database,
		BusinessCollection: businessCollection,
		ReviewsCollection:  reviewCollection,
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

func GetBusinessCollection() *mongo.Collection {
	return globalDB.BusinessCollection
}
func CreateBusiness(business *models.Business) (*models.Business, error) {
	ctx, cancel := initDBContext()
	defer cancel()

	collection := globalDB.BusinessCollection

	// Insert order into the collection
	res, err := collection.InsertOne(ctx, business)
	if err != nil {
		return nil, err
	}

	business.ID = res.InsertedID.(primitive.ObjectID)
	return business, nil
}

func UpdateBusiness(business *models.Business) (*models.Business, error) {
	ctx, cancel := initDBContext()
	defer cancel()

	// The filter to find the document we want to update -> using the order ID
	// TODO: Need to define how we want to update x (Either entire document or single attributes or both)
	filter := bson.M{
		"_id": business.ID,
	}

	// the update operation
	update := bson.M{
		"$set": bson.M{
			//			"items":        business.Items,
			//"updated_at": business.UpdatedAt,
		},
	}

	collection := globalDB.BusinessCollection
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// TODO:  Use Log FatalF
	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("no document found to update")
	}

	return business, nil
}

func DeletedBusiness(business *models.Business) (*models.Business, error) {
	ctx, cancel := initDBContext()
	defer cancel()

	filter := bson.M{"_id": business.ID}

	collection := globalDB.BusinessCollection
	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("no document found with the given_id")

	}
	return business, nil
}

// TODO: Discuss the relationship with reviews <-> business
// TODO: Then implement review section
