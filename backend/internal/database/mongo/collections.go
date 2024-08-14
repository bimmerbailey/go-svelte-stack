package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedDate time.Time          `bson:"created_date"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	Email       string
	Password    string
	IsAdmin     bool `bson:"is_admin"`
}

func main() {

}
