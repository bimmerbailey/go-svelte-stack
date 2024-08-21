package database

type User struct {
	DocumentBase `bson:",inline"`
	FirstName    string `bson:"first_name" json:"first_name"`
	LastName     string `bson:"last_name" json:"last_name"`
	Email        string `bson:"email" json:"email"`
	Password     string `bson:"password" json:"password"`
}

type Item struct {
	DocumentBase `bson:",inline"`
	Cost         float64 `bson:"cost" json:"cost"`
	Name         string  `bson:"name" json:"name"`
	Description  string  `bson:"description" json:"description"`
	Quantity     float64 `bson:"quantity" json:"quantity"`
}
