package models

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/momenteam/momentum/database"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Needer type
type Needer struct {
	ID          string    `bson:"_id" json:"id"`
	FirstName   string    `bson:"firstName" json:"firstName"`
	LastName    string    `bson:"lastName" json:"lastName"`
	Address     string    `bson:"address" json:"address"`
	Category    string    `bson:"category" json:"category"`
	PhoneNumber string    `bson:"phoneNumber" json:"phoneNumber"`
	Summary     string    `bson:"summary" json:"summary"`
	Packages    []Package `bson:"packages" json:"packages"`
	CreatedBy   string    `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
}

// Package type
type Package struct {
	ID         string     `bson:"_id" json:"id"`
	Name       string     `bson:"name" json:"name"`
	TotalPrice int        `bson:"totalPrice" json:"totalPrice"`
	LineItems  []LineItem `bson:"lineItems" json:"lineItems"`
}

// LineItem type
type LineItem struct {
	ID     string `bson:"_id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Price  int    `bson:"price" json:"price"`
	Amount int    `bson:"amount" json:"amount"`
}

// CreateNeeder func
func CreateNeeder(needer Needer) (result Needer, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Needer create error")
		}
	}()

	needer.ID = uuid.New().String()
	_, err = database.NeederCollection.InsertOne(context.Background(), needer)

	return needer, err
}

// GetNeederDetail func
func GetNeederDetail(id string) (Needer, error) {
	needer := Needer{}
	err := database.NeederCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&needer)

	return needer, err
}

// CreatePackage func
func CreatePackage(id string, packageModel Package) (packages Package, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Needer create error")
		}
	}()

	packageModel.ID = uuid.New().String()
	packageModel.LineItems = []LineItem{}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"packages": packageModel}}

	_ = database.NeederCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		&opt,
	)
	return packageModel, err
}

// CreateLineItem func
func CreateLineItem(id string, packageID string, lineItem LineItem) (lineItems LineItem, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Needer create error")
		}
	}()

	lineItem.ID = uuid.New().String()
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	filter := bson.M{"_id": id, "packages._id": packageID}
	update := bson.M{"$push": bson.M{"packages.$.lineItems": lineItem}}

	_ = database.NeederCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		&opt,
	)
	return lineItem, err
}
