package db

// @import
import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @dev Creates a MongoDB instance
//
// @return *mongo.Client
func EstablishMongoClient(ctx context.Context) *mongo.Client {
	// get the mongoDB uri
	mongoUri := os.Getenv("MONGODB_URI")
	if mongoUri == "" {log.Fatal("!MONGODB_URI - uri is not defined.")}

	// Establish the connection
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal("Cannot connect to mongoClient")
	}

	// return mongo client
	log.Println("MongoDB connected...")
	return mongoClient
}

// @dev Gets a mongdb collection based on collectionName
// 
// @param mongoClient *mongo.Client
//  
// @param collectionName string
// 
// @return *mongo.Collection
func GetMongoCollection(mongoClient *mongo.Client, collectionName string) *mongo.Collection {
	// get the collection
	collection := mongoClient.Database(os.Getenv("MONGO_DB")).Collection(collectionName)

	// return the collection
	return collection
}
