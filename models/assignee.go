package models

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/momenteam/momentum/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Assignee struct {
	ID         string `bson:"_id" json:"_id"`
	Name       string `json:"name"`
	Department string `json:"department" validate:"department"`
}

// CreateAssignee creates assignee
func CreateAssignee(assignee Assignee) (result Assignee, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("assignee create error")
		}
	}()

	assignee.ID = uuid.New().String()

	_, err = database.AssigneeCollection.InsertOne(context.Background(), assignee)

	return assignee, err
}

// GetAllAssignees returns all candidates
func GetAllAssignees() ([]Assignee, error) {
	assignees := []Assignee{}

	cursor, err := database.AssigneeCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var assignee Assignee
		if err = cursor.Decode(&assignee); err != nil {
			log.Fatal(err)
		}
		assignees = append(assignees, assignee)
	}

	return assignees, err
}

func ReadAssignee(_id string) (Assignee, error) {
	assignee := Assignee{}
	err := database.AssigneeCollection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&assignee)

	return assignee, err
}

func ReadAssigneeByName(name string) (Assignee, error) {
	assignee := Assignee{}
	err := database.AssigneeCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&assignee)

	return assignee, err
}

// DeleteAssignee deletes assignee by id
func DeleteAssignee(_id string) error {
	_, err := database.AssigneeCollection.DeleteOne(context.TODO(), bson.M{"_id": _id})

	return err
}

func FindRandomAssigneeByDepartment(department string) Assignee {
	assignees := []Assignee{}

	cursor, err := database.AssigneeCollection.Find(context.TODO(), bson.M{"department": department})

	for cursor.Next(context.Background()) {
		var assignee Assignee
		if err = cursor.Decode(&assignee); err != nil {
			log.Fatal(err)
		}
		assignees = append(assignees, assignee)
	}

	rand.Seed(time.Now().Unix())
	return assignees[rand.Intn(len(assignees))]
}

func FindAssigneesCandidates(id string) ([]Candidate, error) {
	candidates := []Candidate{}

	cursor, err := database.CandidatesCollection.Find(context.TODO(), bson.M{"assignee": id})

	for cursor.Next(context.Background()) {
		var candidate Candidate
		if err = cursor.Decode(&candidate); err != nil {
			log.Fatal(err)
		}
		candidates = append(candidates, candidate)
	}

	return candidates, err
}
