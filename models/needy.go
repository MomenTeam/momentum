package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/momenteam/momentum/database"
	"github.com/momenteam/momentum/models/enums"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type Needy struct {
	ID              string                    `bson:"_id" json:"id"`
	FirstName       string                    `bson:"firstName" json:"firstName"`
	LastName        string                    `bson:"lastName" json:"lastName"`
	PhoneNumber     string                    `bson:"phoneNumber" json:"phoneNumber"`
	Summary         string                    `bson:"summary" json:"summary"`
	Priority        int                       `bson:"priority" json:"priority"`
	Address         Address                   `bson:"address" json:"address"`
	NeedyCategories []enums.NeedyCategoryType `bson:"needyCategories" json:"needyCategories"`
	Needs           []string                  `bson:"category" json:"category"`
	CreatedBy       string                    `bson:"createdBy" json:"createdBy"`
	CreatedAt       time.Time                 `bson:"createdAt" json:"createdAt"`
}

type NeedyInformation struct {
	FullName   string                    `json:"fullName"`
	Address    string                    `json:"address"`
	Categories []enums.NeedyCategoryType `json:"needyCategories"`
	Summary    string                    `json:"summary"`
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

func AddNeed(need Need, needyId string) (result Needy, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("needy create error")
		}
	}()

	needy, err := GetNeedy(needyId)

	if err != nil {
		return result, err
	}

	needResult, err := CreateNeed(need)

	if err != nil {
		return result, err
	}

	needy.Needs = appendIfMissing(needy.Needs, needResult.ID)

	filter := bson.M{"_id": bson.M{"$eq": needyId}}
	update := bson.M{"$set": bson.M{"needs": needy.Needs}}

	_, err = database.NeediesCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {
		return result, err
	}

	return needy, err
}

func GetAllNeediesInformations() ([]NeedyInformation, error) {
	var neediesInformation []NeedyInformation

	cursor, err := database.NeediesCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var needy Needy
		if err = cursor.Decode(&needy); err != nil {
			log.Fatal(err)
		}
		needyInformation := NeedyInformation{
			FullName:   fmt.Sprintf("%s %s", mask(needy.FirstName), mask(needy.LastName)),
			Address:    fmt.Sprintf("%s, %s", needy.Address.District, needy.Address.City),
			Summary:    needy.Summary,
			Categories: needy.NeedyCategories,
		}

		neediesInformation = append(neediesInformation, needyInformation)
	}

	return neediesInformation, err
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

func mask(s string) string {
	runes := []rune(s)
	result := string(runes[0:1])
	return result + "***"
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}