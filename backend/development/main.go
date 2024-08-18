package main

import (
	"backend/internal/database/mongo"
	"context"
	"fmt"
	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func createUsers(db *mongo.Database) {
	err := db.Collection("users").Drop(context.Background())
	if err != nil {
		return
	}
	collection := database.GetCollection[*database.User](db, "users")
	for i := 0; i < 10; i++ {
		firstName := faker.FirstName()
		lastName := faker.LastName()
		newUser := database.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     fmt.Sprintf("%s.%s@example.com", firstName, lastName),
			Password:  faker.Password(),
		}

		_, err := collection.Insert(&newUser)
		if err != nil {
			return
		}
	}
}

func setUpDatabase() (*mongo.Database, error) {
	client, mongoErr := database.ConnectMongoDb()
	if mongoErr != nil {
		log.Fatal(mongoErr)
		return nil, mongoErr
	}
	mongoDB := client.Database("your_app")

	collectionNames := []string{"users", "items"}
	database.InitializeCollections(mongoDB, collectionNames)
	return mongoDB, nil
}

func main() {
	devDb, err := setUpDatabase()

	if err != nil {
		log.Fatal(err)
		return
	}
	createUsers(devDb)
}
