package database

import (
	"context"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDb take mongodb url and related to connections
func ConnectMongoDb() (*mongo.Client, error) {

	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoUri := os.Getenv("MONGO_URI")
	mongoAuth := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	clientOptions := options.Client().SetAuth(mongoAuth).ApplyURI(mongoUri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		slog.Error("Error connecting to MongoDB", err)
		CloseMongoClient(client)
		return nil, err
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		slog.Error("Error pinging to MongoDB", err)
		return nil, err
	}

	slog.Info("MongoClient connected")

	return client, nil
}

// CloseMongoClient disconnect mongo client
func CloseMongoClient(client *mongo.Client) {
	slog.Info("Closing mongo client")
	err := client.Disconnect(context.TODO())
	if err != nil {
		slog.Info("Error disconnecting mongo client", "error", err)
		return
	}
	slog.Info("Mongo client disconnected")
}

func InitializeCollections(db *mongo.Database, collections []string) {
	for i := 0; i < len(collections); i++ {
		slog.Info("Initializing mongo collection", "collection", collections[i])
		err := db.CreateCollection(context.Background(), collections[i])
		if err != nil {
			panic(err)
		}
	}

}
