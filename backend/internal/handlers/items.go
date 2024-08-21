package handlers

import (
	"backend/internal/database/mongo"
	"backend/internal/responses"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type ItemHandler interface {
	Insert(http.ResponseWriter, *http.Request)
	GetItems(http.ResponseWriter, *http.Request)
	GetById(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type itemHandler struct {
	collection *database.Collection[*database.Item]
}

func NewItemHandler(collection *database.Collection[*database.Item]) ItemHandler {
	return &itemHandler{collection: collection}
}

type GetItems struct {
	Items []*database.Item `json:"items"`
	Count int              `json:"count"`
}

func (handler *itemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := handler.collection.GetDocuments()
	if err != nil {
		slog.Warn("Error getting users", "error", err)
		responses.StringResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.JsonResponse(w, http.StatusOK, GetItems{Items: items, Count: len(items)})
}

func (handler *itemHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var newItem database.Item

	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		return
	}
	user, err := handler.collection.Insert(&newItem)
	if err != nil {
		slog.Warn("Error inserting user", "error", err)
		responses.StringResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.JsonResponse(w, http.StatusCreated, user)
}

func (handler *itemHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	item, err := handler.collection.GetDocumentByID(id)
	if err != nil {
		slog.Warn("User not found", "id", id, "err", err)
		errorMessage := fmt.Sprintf("User %s not found", id)
		responses.StringResponse(w, http.StatusNotFound, errorMessage)
		return
	}
	responses.JsonResponse(w, http.StatusOK, item)
}

func (handler *itemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	deleted, err := handler.collection.DeleteDocument(id)
	if err != nil || deleted == 0 {
		slog.Warn("Problem deleting user", "id", id, "err", err)
		errorMessage := fmt.Sprintf("User %s not found", id)
		responses.StringResponse(w, http.StatusNotFound, errorMessage)
		return
	}
	slog.Info("Item deleted", "id", id, "count", deleted)
	responses.StringResponse(w, http.StatusOK, fmt.Sprintf("User %s deleted", id))
}
