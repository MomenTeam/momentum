package models

type Address struct {
	FirstName   string `bson:"firstName" json:"firstName"`
	LastName    string `bson:"lastName" json:"lastName"`
	FirstLine   string `bson:"firstLine" json:"firstLine"`
	SecondLine  string `bson:"secondLine" json:"secondLine"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
	PostalCode  string `bson:"postalCode" json:"postalCode"`
	District    string `bson:"district" json:"district"`
	City        string `bson:"city" json:"city"`
}
