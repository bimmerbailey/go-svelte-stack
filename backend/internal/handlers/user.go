package handlers

import (
	database "backend/internal/database/mongo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler interface {
	Insert(*gin.Context)
	GetUsers(*gin.Context)
	GetById(*gin.Context)
	Delete(*gin.Context)
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

func (handler *userHandler) GetUsers(c *gin.Context) {
	// FIXME: users should be [] w/o users not null, if possible
	// FIXME: Allow params for filtering
	// FIXME: Better way to get count of documents
	users, err := handler.collection.GetDocuments()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, GetUsers{Users: users, Count: len(users)})
}

func (handler *userHandler) Insert(c *gin.Context) {
	var newUser database.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	user, err := handler.collection.Insert(&newUser)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, user)
}

func (handler *userHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	user, err := handler.collection.GetDocumentByID(id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, user)
}

// Delete TODO:
func (handler *userHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, id)
}
