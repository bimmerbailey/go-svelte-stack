package responses

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ResponseResult struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type ResponseError struct {
	Message string                 `json:"message"`
	Status  int                    `json:"status"`
	Error   string                 `json:"error"`
	Data    map[string]interface{} `json:"data"`
}

// BadRequestError return ResponseError with bad_request status and messages
func BadRequestError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Data:    data,
	}
}

// NotFoundRequestError return ResponseError with not_found status and messages
func NotFoundRequestError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
		Data:    data,
	}
}

// InternalServerError message error return ResponseError with internal_server status and messages
func InternalServerError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server",
		Data:    data,
	}
}

// Response BadRequestError return ResponseError with bad_request status and messages
func Response(message string, data map[string]interface{}) *ResponseResult {
	return &ResponseResult{
		Message: message,
		Data:    data,
	}
}

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
