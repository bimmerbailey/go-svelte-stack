package routes

import (
	"backend/internal/database/mongo"
	"backend/internal/handlers"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func UsersRoutes(router chi.Router, db *mongo.Database) {
	collection := database.GetCollection[*database.User](db, "users")
	handler := handlers.NewUserHandler(collection)

	router.Get("/users", handler.GetUsers)
	router.Get("/users/{id}", handler.GetById)
	router.Post("/users", handler.Insert)
	router.Delete("/users/{id}", handler.Delete)
}
