package models

import "github.com/momenteam/momentum/models/enums"

type Good struct {
	Name         string             `bson:"name" json:"name"`
	Price        float32            `bson:"price" json:"price"`
	PhotoLink    string             `bson:"photoLink" json:"photoLink"`
	IsAvailable  bool               `bson:"isAvailable" json:"isAvailable"`
	GoodCategory enums.GoodCategory `bson:"goodCategory" json:"goodCategory"`
}
