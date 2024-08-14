package routes

import (
	utils "backend/internal"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"net/http"
	"time"
)

func HealthRoutes(router *gin.RouterGroup, client *mongo.Client) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"okay": true,
		})
	})

	router.GET("/readyz", func(c *gin.Context) {
		ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(3000)*time.Second)
		defer ctxErr()

		if ctxErr != nil {
			slog.Info("Can't connect to database", "error", ctxErr)
		}

		// TODO: Format the sting
		dateNow := time.Now().Local()

		if err := client.Ping(ctx, nil); err != nil {
			utils.InternalServerError(
				"Status unhealthy",
				err,
				map[string]interface{}{
					"Data": "Please check the Client",
					"Time": dateNow.String(),
				},
			)
		}

		c.IndentedJSON(
			http.StatusOK,
			utils.Response(
				"Pong",
				map[string]interface{}{
					"Data": "The MongoDB client is working successfully",
					"Date": dateNow.String(),
				},
			),
		)
	})
}
