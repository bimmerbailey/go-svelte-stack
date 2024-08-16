package routes

import (
	database "backend/internal/database/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type GetUsers struct {
	Users []*database.User `json:"users"`
	Count int              `json:"count"`
}

func UsersRoutes(router *gin.RouterGroup, db *mongo.Database) {
	collection := database.GetCollection[*database.User](db, "users")
	router.GET("/users", func(c *gin.Context) {
		// FIXME: users should be [] w/o users not null
		users, err := collection.GetDocuments()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, GetUsers{Users: users, Count: len(users)})
	})

	router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"user": name,
		})
	})

	router.POST("/users", func(c *gin.Context) {
		var newUser database.User
		if err := c.BindJSON(&newUser); err != nil {
			return
		}
		user, err := collection.Insert(&newUser)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, gin.H{
			"user": user,
		})
	})
}
