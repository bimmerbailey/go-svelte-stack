package routes

import (
	database "backend/internal/database/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetUsers struct {
	Users []*database.User `json:"users"`
	Count int              `json:"count"`
}

func UsersRoutes(router *gin.RouterGroup, db *mongo.Database) {
	collection := database.GetCollection[*database.User](db, "users")
	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": []string{"user1", "user2"},
		})
	})

	router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"user": name,
		})
	})
}
