package routes

import (
	"backend/internal/responses"
	"context"
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

		// TODO: Format the sting
		dateNow := time.Now().Local()

		if err := client.Ping(ctx, nil); err != nil {
			responses.InternalServerError(
				"Status unhealthy",
				err,
				map[string]interface{}{
					"Data": "Please check the Client",
					"Time": dateNow.String(),
				},
			)
		}
		responses.JsonResponse(w, http.StatusOK, map[string]interface{}{
			"Data": "The MongoDB client is working successfully",
			"Date": dateNow.String(),
		})
	})
}
