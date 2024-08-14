package main

import (
	database "backend/internal/database/mongo"
	"backend/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	app := gin.Default()

	apiRouter := app.Group("/api")
	V1Router := apiRouter.Group("/v1")
	client, mongoErr := database.ConnectMongoDb()

	if mongoErr != nil {
		log.Fatal(mongoErr)
	}

	collectionNames := []string{"users", "items"}
	database.InitializeCollections(client, "your_app", collectionNames)

	routes.HealthRoutes(apiRouter, client)
	routes.UsersRoutes(V1Router, client)

	err := app.Run(":8080")

	if err != nil {
		return
	}

	//	TODO: Close mongo connection on shutdown
}
