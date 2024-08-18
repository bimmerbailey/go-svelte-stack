package handlers

import (
	"backend/internal/database/mongo"
	"backend/internal/responses"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type UserHandler interface {
	Insert(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	GetById(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	collection *database.Collection[*database.User]
}

func NewUserHandler(collection *database.Collection[*database.User]) UserHandler {
	return &userHandler{collection: collection}
}

type GetUsers struct {
	Users []*database.User `json:"users"`
	Count int              `json:"count"`
}

func (handler *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// FIXME: users should be [] w/o users not null, if possible
	// FIXME: Allow params for filtering https://gin-gonic.com/docs/examples/bind-query-or-post/
	// FIXME: Better way to get count of documents
	users, err := handler.collection.GetDocuments()
	if err != nil {
		slog.Warn("Error getting users", "error", err)
		responses.StringResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		responses.JsonResponse(w, http.StatusOK, GetUsers{Users: users, Count: len(users)})
	}
}

func (handler *userHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return
	}
	user, err := handler.collection.Insert(&newUser)
	if err != nil {
		slog.Warn("Error inserting user", "error", err)
		responses.StringResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		responses.JsonResponse(w, http.StatusCreated, user)
	}
}

func (handler *userHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := handler.collection.GetDocumentByID(id)
	if err != nil {
		slog.Warn("User not found", "id", id, "err", err)
		errorMessage := fmt.Sprintf("User %s not found", id)
		responses.StringResponse(w, http.StatusNotFound, errorMessage)
	} else {
		responses.JsonResponse(w, http.StatusOK, user)
	}
}

func (handler *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	deleted, err := handler.collection.DeleteDocument(id)
	if err != nil || deleted == 0 {
		slog.Warn("Problem deleting user", "id", id, "err", err)
		errorMessage := fmt.Sprintf("User %s not found", id)
		responses.StringResponse(w, http.StatusNotFound, errorMessage)
	} else {
		slog.Info("User deleted", "id", id, "count", deleted)
		responses.StringResponse(w, http.StatusOK, fmt.Sprintf("User %s deleted", id))
	}
}
