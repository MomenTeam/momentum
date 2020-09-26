package models

type LineItem struct {
	Amount int `bson: "amount"`
	Price float32 `bson: "price"`
}
