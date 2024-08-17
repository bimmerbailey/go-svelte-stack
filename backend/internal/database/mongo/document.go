package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Document interface {
	SetID(id string)
	GetID() string
}

// Doc TODO: We might have to convert id to ObjectID to get by id
type Doc struct {
	ID string `bson:"_id,omitempty" json:"id"`
}

func (doc *Doc) SetID(id string) {
	doc.ID = id
}

func (doc *Doc) GetID() string {
	return doc.ID
}

func (doc *Doc) PrepareId() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(doc.ID)
}

type DocumentBase struct {
	Doc       `bson:",inline"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (doc *DocumentBase) BeforeInsert() error {
	now := time.Now().UTC()
	doc.CreatedAt = now
	doc.UpdatedAt = now
	return nil
}

func (doc *DocumentBase) BeforeUpdate() error {
	now := time.Now().UTC()
	doc.UpdatedAt = now
	return nil
}
