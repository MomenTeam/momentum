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
	"strings"
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
	Needs           []string                  `bson:"needs" json:"needs"`
	CreatedBy       string                    `bson:"createdBy" json:"createdBy"`
	CreatedAt       time.Time                 `bson:"createdAt" json:"createdAt"`
}

type NeedyInformation struct {
	ID         string                    `json:"id"`
	FullName   string                    `json:"fullName"`
	Address    string                    `json:"address"`
	Categories []enums.NeedyCategoryType `json:"needyCategories"`
	Summary    string                    `json:"summary"`
	ShortName  string                    `json:"shortName"`
}

type NeedyDetail struct {
	ID              string                    `json:"id"`
	FirstName       string                    `json:"firstName"`
	LastName        string                    `json:"lastName"`
	ShortName       string                    `json:"shortName"`
	MaskedName      string                    `json:"maskedName"`
	PhoneNumber     string                    `json:"phoneNumber"`
	Summary         string                    `json:"summary"`
	Priority        int                       `json:"priority"`
	Address         Address                   `json:"address"`
	NeedyCategories []enums.NeedyCategoryType `json:"needyCategories"`
	Needs           []Need                    `json:"needs"`
	CreatedBy       string                    `json:"createdBy"`
	CreatedAt       time.Time                 `json:"createdAt"`
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
			ID:         needy.ID,
			FullName:   fmt.Sprintf("%s %s", mask(needy.FirstName), mask(needy.LastName)),
			Address:    fmt.Sprintf("%s, %s", needy.Address.District, needy.Address.City),
			Summary:    needy.Summary,
			Categories: needy.NeedyCategories,
			ShortName:  fmt.Sprintf("%c%c", needy.FirstName[0], needy.LastName[0]),
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

func GetNeedyByNeedId(id string) (Needy, error) {
	needy := Needy{}
	needies, err := GetAllNeedies()

	for _, needy := range needies {
		if strings.Contains(strings.Join(needy.Needs, ","), id) {
			return needy, err
		}
	}

	return needy, err
}

func GetNeedyDetail(id string) (NeedyDetail, error) {
	needy := Needy{}
	err := database.NeediesCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&needy)

	needFilter := bson.M{"_id": bson.M{"$in": needy.Needs}}

	var needs []Need
	cursor, err := database.NeedCollection.Find(context.Background(), needFilter)

	if cursor != nil {
		for cursor.Next(context.Background()) {
			var need Need
			if err = cursor.Decode(&need); err != nil {
				log.Fatal(err)
			}
			needs = append(needs, need)
		}

		needyDetail := NeedyDetail{
			ID:              needy.ID,
			FirstName:       needy.FirstName,
			LastName:        needy.LastName,
			ShortName:  fmt.Sprintf("%c%c", needy.FirstName[0], needy.LastName[0]),
			MaskedName:        fmt.Sprintf("%s %s", mask(needy.FirstName), mask(needy.LastName)),
			PhoneNumber:     needy.PhoneNumber,
			Summary:         needy.Summary,
			Priority:        needy.Priority,
			Address:         needy.Address,
			NeedyCategories: needy.NeedyCategories,
			Needs:           needs,
			CreatedBy:       needy.CreatedBy,
			CreatedAt:       needy.CreatedAt,
		}

		return needyDetail, err
	}

	return NeedyDetail{}, err
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
