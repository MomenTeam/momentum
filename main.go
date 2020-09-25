package main

import (
	"context"
	"fmt"
	"github.com/momenteam/momentum/configs"
	"github.com/momenteam/momentum/database"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	configs.Setup()
	database.Setup()
}

func main() {
	client := database.Client

	if client == nil {
		fmt.Println("error")
	}

	result, err := database.CandidatesCollection.InsertOne(
		context.Background(),
		bson.D{
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
		})

	fmt.Println(result)

	if err != nil {
		fmt.Println("error occured")
	}
}