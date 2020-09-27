package controllers

import "github.com/momenteam/momentum/models"

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
	//GoodIds     []int  `bson:"goodIds" json:"goodIds"`
}

type PaymentForm struct {
	CreditCardNumber string `json:"creditCardNumer"`
	Cvv string `json:"cvv"`
	ExpireDate string `json:"expireDate"`
	FullName string `json:"fullName"`
}