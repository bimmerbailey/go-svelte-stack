package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UsersRoutes(router *gin.RouterGroup, client *mongo.Client) {
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
