package models

import "time"

type Need struct {
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description" json:"description"`
	LineItems   []LineItem `bson:"lineItems" json:"lineItem"`
	Status      bool       `bson:"isFulfilled" json:"isFulfilled"`
	FulfilledBy string     `bson:"fulfilledBy" json:"fulfilledBy"` //TODO: change this
	FulfilledAt time.Time  `bson:"fulfilledAt" json:"fulfilledAt"`
	IsCancelled bool       `bson:"isCancelled" json:"isCancelled"`
	CancelledAt time.Time  `bson:"cancelledAt" json:"cancelledAt"`
	CancelledBy string     `bson:"cancelledBy" json:"cancelledBy"`
}
