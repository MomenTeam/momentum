package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/momenteam/momentum/database"
	"github.com/momenteam/momentum/models/enums"
	"github.com/momenteam/momentum/utils"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type Need struct {
	ID          string           `bson:"_id" json:"id"`
	Name        string           `bson:"name" json:"name"`
	Description string           `bson:"description" json:"description"`
	LineItems   []LineItem       `bson:"lineItems" json:"lineItems"`
	Status      enums.NeedStatus `bson:"status" json:"status"`
	Priority    int              `bson:"priority" json:"priority"`
	FulfilledAt time.Time        `bson:"fulfilledAt" json:"fulfilledAt"`
	PaidAt      time.Time        `bson:"paidAt" json:"paidAt"`
	PaidBy      string           `bson:"paidBy" json:"paidBy"`
	PayerEmail  string           `bson:"payerEmail" json:"payerEmail"`
	CancelledAt time.Time        `bson:"cancelledAt" json:"cancelledAt"`
	CancelledBy string           `bson:"cancelledBy" json:"cancelledBy"` //TODO: edit
	CreatedAt   time.Time        `bson:"createdAt" json:"createdAt"`
}

type NeedLineItemDetailDto struct {
	Amount int `json:"amount"`
	Id     int `json:"id"`
}

type NeedDetailDto struct {
	ID          string                  `json:"id"`
	FullName    string                  `json:"fullName"`
	FullAddress string                  `json:"fullAddress"`
	Name        string                  `json:"name"`
	LineItems   []NeedLineItemDetailDto `json:"lineItems"`
	PaidBy      string                  `json:"payerName"`
	PayerEmail  string                  `json:"payerEmail"`
	Status      enums.NeedStatus        `json:"status"`
}

func CreateNeed(need Need) (result Need, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("need create error")
		}
	}()

	need.ID = uuid.New().String()
	need.Status = enums.NeedCreated

	_, err = database.NeedCollection.InsertOne(context.Background(), need)

	return need, err
}

func PayNeed(id string, paidBy string, email string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("need create error")
		}
	}()

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"status": enums.NeedPaid, "paidAt": time.Now(), "paidBy": paidBy, "payerEmail": email}}

	_, err = database.NeedCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	utils.SendEmail(paidBy, 2, email)

	return id, err
}

func GetAllNeeds() (result []Need, err error) {
	needs := []Need{}

	cursor, err := database.NeedCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var need Need
		if err = cursor.Decode(&need); err != nil {
			log.Fatal(err)
		}
		if need.Status == enums.NeedPaid || need.Status == enums.NeedFulfilled {
			needs = append(needs, need)
		}
	}

	return needs, err
}

func SetFulfilled(id string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("need create error")
		}
	}()

	need, err := GetNeed(id)

	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"status": enums.NeedFulfilled, "fulfilledAt": time.Now()}}

	_, err = database.NeedCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	utils.SendEmail(need.PaidBy, 1, need.PayerEmail)

	return id, err
}

func GetNeed(id string) (Need, error) {
	need := Need{}
	err := database.NeedCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&need)

	return need, err
}

func CancelNeed(id string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("need cancel error")
		}
	}()

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"status": enums.NeedCancelled, "cancelledAt": time.Now()}}

	_, err = database.NeedCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	return id, err
}

func GetAllNeedDetails() (result []NeedDetailDto, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("get need details error")
		}
	}()

	var needDetails []NeedDetailDto
	needs, err := GetAllNeeds()

	for _, need := range needs {
		needy, err := GetNeedyByNeedId(need.ID)

		if err != nil {
			log.Fatal(err)
		}

		var lineItems []NeedLineItemDetailDto
		for _, lineItem := range need.LineItems {
			lineItems = append(lineItems, NeedLineItemDetailDto{
				Amount: lineItem.Amount,
				Id:     lineItem.Good.GoodId,
			})
		}

		needDetail := NeedDetailDto{
			ID:          need.ID,
			FullName:    fmt.Sprintf("%s %s", needy.FirstName, needy.LastName),
			FullAddress: getFullAddress(needy.Address),
			Name:        need.Name,
			LineItems:   lineItems,
			PaidBy:      need.PaidBy,
			PayerEmail:  need.PayerEmail,
			Status:      need.Status,
		}

		needDetails = append(needDetails, needDetail)
	}

	return needDetails, err
}

func getFullAddress(address Address) string {
	return fmt.Sprintf("%s %s %s %s/%s",
		address.FirstLine, address.SecondLine, address.PostalCode, address.District, address.City)
}
