package responses

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// FIXME: Better error handling?
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Warn("Error marshalling json", "error", err)
		panic(err)
	}
	_, err = w.Write(jsonData)
	if err != nil {
		panic(err)
	}
}

func StringResponse(w http.ResponseWriter, status int, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// FIXME: Better error handling?
	_, err := w.Write([]byte(data))
	if err != nil {
		panic(err)
	}
}

// InternalServerError message error return StringResponse with internal_server status and messages
func InternalServerError(w http.ResponseWriter, err error) {
	StringResponse(w, http.StatusInternalServerError, err.Error())
}
