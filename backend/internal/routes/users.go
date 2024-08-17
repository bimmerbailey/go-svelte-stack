package routes

import (
	database "backend/internal/database/mongo"
	"backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetUsers struct {
	Users []*database.User `json:"users"`
	Count int              `json:"count"`
}

func UsersRoutes(router *gin.RouterGroup, db *mongo.Database) {
	collection := database.GetCollection[*database.User](db, "users")
	handler := handlers.NewUserHandler(collection)

	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetById)
	router.POST("/users", handler.Insert)
}
