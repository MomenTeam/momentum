package models

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/momenteam/momentum/database"
	"github.com/momenteam/momentum/models/enums"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type Needy struct {
	ID        string             `bson:"_id" json:"id"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	PhoneNumber       string     `bson:"phoneNumber" json:"phoneNumber"`
	Summary   string             `bson:"summary" json:"summary"`
	Priority  int                `bson:"priority" json:"priority"`
	Address   Address            `bson:"address" json:"address"`
	Category  enums.CategoryType `bson:"category" json:"category"`
	Needs     []Need             `bson:"category" json:"category"`
	CreatedBy   string     `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time  `bson:"createdAt" json:"createdAt"`
}

func CreateNeedy(needy Needy) (result Needy, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("needy create error")
		}
	}()

	needy.ID = uuid.New().String()

	_, err = database.NeediesCollection.InsertOne(context.Background(), needy)

	return needy, err
}

func GetAllNeedies() ([]Needy, error) {
	needies := []Needy{}

	cursor, err := database.NeediesCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var needy Needy
		if err = cursor.Decode(&needy); err != nil {
			log.Fatal(err)
		}
		needies = append(needies, needy)
	}

	return needies, err
}

func GetNeedy(id string) (Needy, error) {
	needy := Needy{}
	err := database.NeediesCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&needy)

	return needy, err
}

func DeleteNeedy(id string, cancelledBy string) error {
	filter := bson.M{"_id": bson.M{"$eq": id}}

	update := bson.M{"$set": bson.M{"isCancelled": true, "cancelledBy": cancelledBy, "cancelledAt": time.Now()}}

	_, err := database.NeediesCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	return err
}
