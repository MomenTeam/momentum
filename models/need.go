package models

import (
	"context"
	"errors"
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
		needs = append(needs, need)
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

	utils.SendEmail(need.PaidBy, 2, need.PayerEmail)

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
			err = errors.New("need create error")
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
