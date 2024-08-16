package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"log/slog"
	"time"
)

type Collection[T Document] struct {
	collection *mongo.Collection
}

type BeforeInsertHook interface {
	BeforeInsert() error
}

type BeforeUpdateHook interface {
	BeforeUpdate() error
}

func DefaultContext() (context.Context, context.CancelFunc) {
	ctx, err := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, err
}

func GetCollection[T Document](db *mongo.Database, collectionName string) *Collection[T] {
	return &Collection[T]{db.Collection(collectionName)}
}

func (collection Collection[T]) Insert(doc T) (T, error) {
	ctx, cancel := DefaultContext()
	defer cancel()

	if hook, ok := any(doc).(BeforeInsertHook); ok {
		if err := hook.BeforeInsert(); err != nil {
			return doc, err
		}
	}

	res, err := collection.collection.InsertOne(ctx, doc)
	doc.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	slog.Info("Inserted a new document", "Document", doc)
	return doc, err
}

func (collection Collection[T]) GetDocuments() ([]T, error) {
	ctx, cancel := DefaultContext()
	defer cancel()

	var docs []T

	cur, err := collection.collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = cur.All(context.TODO(), &docs)
	if err != nil {
		log.Fatal(err)
	}

	return docs, err
}

type User struct {
	DocumentBase `bson:",inline"`
	FirstName    string `bson:"first_name" json:"first_name"`
	LastName     string `bson:"last_name" json:"last_name"`
	Email        string `bson:"email" json:"email"`
	Password     string `bson:"password" json:"password"`
}
