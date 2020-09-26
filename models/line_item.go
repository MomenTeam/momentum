package models

type LineItem struct {
	Description string `bson:"description" json:"description"`
	Amount      int    `bson:"amount" json:"amount"`
	Good        Good   `bson:"good" json:"good"`
}
