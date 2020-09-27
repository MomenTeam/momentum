package models

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/momenteam/momentum/database"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Need struct {
	ID          string     `bson:"_id" json:"id"`
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description" json:"description"`
	LineItems   []LineItem `bson:"lineItems" json:"lineItems"`
	IsFulfilled bool       `bson:"isFulfilled" json:"isFulfilled"`
	Priority    int        `bson:"priority" json:"priority"`
	FulfilledBy string     `bson:"fulfilledBy" json:"fulfilledBy"` //TODO: change this
	FulfilledAt time.Time  `bson:"fulfilledAt" json:"fulfilledAt"`
	IsCancelled bool       `bson:"isCancelled" json:"isCancelled"`
	CancelledAt time.Time  `bson:"cancelledAt" json:"cancelledAt"`
	CancelledBy string     `bson:"cancelledBy" json:"cancelledBy"`
	CreatedAt   time.Time  `bson:"createdAt" json:"createdAt"`
}

func CreateNeed(need Need) (result Need, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("need create error")
		}
	}()

	need.ID = uuid.New().String()

	_, err = database.NeedCollection.InsertOne(context.Background(), need)

	return need, err
}

func PayNeed(id string, fulfilledBy string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("need create error")
		}
	}()

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"isFulfilled": true, "fulfilledAt": time.Now(), "fulfilledBy": fulfilledBy}}

	_, err = database.NeedCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	return id, err
}
