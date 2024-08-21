package routes

import (
	"backend/internal/database/mongo"
	"backend/internal/handlers"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func ItemsApiRoutes(router chi.Router, db *mongo.Database) {
	collection := database.GetCollection[*database.Item](db, "items")
	handler := handlers.NewItemHandler(collection)

	router.Get("/items", handler.GetItems)
	router.Get("/items/{id}", handler.GetById)
	router.Post("/items", handler.Insert)
	router.Delete("/items/{id}", handler.Delete)
}
