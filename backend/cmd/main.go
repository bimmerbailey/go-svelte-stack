package main

import (
	"backend/internal/database/mongo"
	"backend/internal/responses"
	"backend/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {

	client, mongoErr := database.ConnectMongoDb()
	if mongoErr != nil {
		log.Fatal(mongoErr)
		return
	}
	mongoDB := client.Database("your_app")

	// FIXME: Not a fan of this
	collectionNames := []string{"users", "items"}
	database.InitializeCollections(mongoDB, collectionNames)

	app := chi.NewRouter()
	// TODO: Look into using slog
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	app.NotFound(func(w http.ResponseWriter, r *http.Request) {
		responses.StringResponse(w, http.StatusNotFound, "route not found")
	})
	app.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		responses.StringResponse(w, http.StatusMethodNotAllowed, "method is not valid")
	})

	app.Route("/api", func(r chi.Router) {
		routes.HealthRoutes(r, client)
	})
	app.Route("/api/v1", func(r chi.Router) {
		routes.UsersRoutes(r, mongoDB)
	})

	err := http.ListenAndServe(":8080", app)
	if err != nil {
		return
	}

	//	TODO: Close mongo connection on shutdown
}
