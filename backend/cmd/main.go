package main

import (
	database "backend/internal/database/mongo"
	"backend/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {

	app := gin.Default()
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
	})

	app.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed"})
	})

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	apiRouter := app.Group("/api")
	V1Router := apiRouter.Group("/v1")
	client, mongoErr := database.ConnectMongoDb()
	if mongoErr != nil {
		log.Fatal(mongoErr)
		return
	}
	mongoDB := client.Database("your_app")

	// FIXME: Not a fan of this
	collectionNames := []string{"users", "items"}
	database.InitializeCollections(mongoDB, collectionNames)

	routes.HealthRoutes(apiRouter, client)
	routes.UsersRoutes(V1Router, mongoDB)

	err := app.Run(":8080")

	if err != nil {
		return
	}

	//	TODO: Close mongo connection on shutdown
}
