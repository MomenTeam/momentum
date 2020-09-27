package models

type LineItem struct {
	Good        Good   `bson:"good" json:"good"`
	Description string `bson:"description" json:"description"`
	Amount      int    `bson:"amount" json:"amount"`
}
