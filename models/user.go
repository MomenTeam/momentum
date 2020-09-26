package models

type User struct {
	Email string `bson:"email" json:"email"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Password string `bson:"password" json:"password"`
	Role
}
