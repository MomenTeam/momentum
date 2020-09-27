package models

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/momenteam/momentum/database"
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
			err = errors.New("needy create error")
		}
	}()

	need.ID = uuid.New().String()

	_, err = database.NeedCollection.InsertOne(context.Background(), need)

	return need, err
}

