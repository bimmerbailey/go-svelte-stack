package routes

import (
	"backend/internal/responses"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"net/http"
	"time"
)

func HealthRoutes(router chi.Router, client *mongo.Client) {
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		responses.StringResponse(w, http.StatusOK, "Healthy")
	})

	router.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, ctxErr := context.WithTimeout(r.Context(), time.Duration(3000)*time.Second)
		defer ctxErr()

		if ctxErr != nil {
			slog.Info("Can't connect to database", "error", ctxErr)
		}

		if err := client.Ping(ctx, nil); err != nil {
			responses.InternalServerError(w, err)
			return
		}
		dateNow := time.Now().UTC()
		responses.StringResponse(w, http.StatusOK, fmt.Sprintf("Successfully pinged database at %s", dateNow))
	})
}
