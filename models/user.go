package models

import "github.com/momenteam/momentum/models/enums"

type User struct {
	Email string `bson:"email" json:"email"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Password string `bson:"password" json:"password"`
	Role enums.UserRole `bson:"roles" json:"role"`
}
