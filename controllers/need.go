package controllers

import (
	"github.com/momenteam/momentum/models"
)

type NeedForm struct {
	Name        string         `bson:"name" json:"name"`
	Description string         `bson:"description" json:"description"`
	LineItems   []LineItemForm `bson:"lineItems" json:"lineItems"`
	Priority    int            `bson:"priority" json:"priority"`
}

type LineItemForm struct {
	Description string      `bson:"description" json:"description"`
	Amount      int         `bson:"amount" json:"amount"`
	Good        models.Good `bson:"good" json:"good"`
}

type PaymentForm struct {
	CreditCardNumber string
	Cvv string
	ExpireDate string

}