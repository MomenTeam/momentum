package models

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ffurkanhas/otsimo/database"
	"github.com/ffurkanhas/otsimo/validator"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Candidate struct
type Candidate struct {
	ID              string    `bson:"_id" json:"_id"`
	FirstName       string    `bson:"first_name" json:"first_name"`
	LastName        string    `bson:"last_name" json:"last_name"`
	Email           string    `bson:"email" json:"email" validate:"email"`
	Department      string    `bson:"department" json:"department" validate:"department`
	University      string    `bson:"university" json:"university"`
	Experience      bool      `bson:"experience" json:"experience"`
	Status          string    `bson:"status" json:"status"`
	MeetingCount    int       `bson:"meeting_count" json:"meeting_count"`
	NextMeeting     time.Time `bson:"next_meeting" json:"next_meeting"`
	Assignee        string    `bson:"assignee" json:"assignee"`
	ApplicationDate time.Time `bson:"application_date" json:"application_date"`
}

// CreateCandidate creates candidate and sets random assignee
func CreateCandidate(candidate Candidate) (result Candidate, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("candidate create error")
		}
	}()

	validationErrors := validator.Validator.Struct(candidate)
	if validationErrors != nil {
		return candidate, validationErrors
	}

	candidate.ID = uuid.New().String()
	candidate.MeetingCount = 0
	candidate.Status = "Pending"
	if len(candidate.Assignee) <= 0 {
		candidate.Assignee = FindRandomAssigneeByDepartment(candidate.Department).ID
	}
	if candidate.ApplicationDate.IsZero() {
		candidate.ApplicationDate = time.Now()
	}

	_, err = database.CandidatesCollection.InsertOne(context.Background(), candidate)

	return candidate, err
}

// GetAllCandidates returns all candidates
func GetAllCandidates() ([]Candidate, error) {
	candidates := []Candidate{}

	cursor, err := database.CandidatesCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var candidate Candidate
		if err = cursor.Decode(&candidate); err != nil {
			log.Fatal(err)
		}
		candidates = append(candidates, candidate)
	}

	return candidates, err
}

// ReadCandidate find candidate by id
func ReadCandidate(_id string) (Candidate, error) {
	candidate := Candidate{}
	err := database.CandidatesCollection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&candidate)

	return candidate, err
}

// ReadCandidateByName finds candidate by name
func ReadCandidateByName(name string) (Candidate, error) {
	candidate := Candidate{}
	err := database.CandidatesCollection.FindOne(context.TODO(), bson.M{"first_name": name}).Decode(&candidate)

	return candidate, err
}

// DeleteCandidate deleted candidate by id
func DeleteCandidate(_id string) error {
	_, err := database.CandidatesCollection.DeleteOne(context.TODO(), bson.M{"_id": _id})

	return err
}

// ArrangeMeeting I assume that meeting_count is increased only when meeting is completed
func ArrangeMeeting(_id string, nextMeetingTime *time.Time) error {
	candidate, _ := ReadCandidate(_id)

	if candidate.MeetingCount >= 4 {
		return errors.New("meeting_count cannot be greater than 4")
	}

	var newData bson.M

	newData = bson.M{
		"$set": bson.M{
			"next_meeting": nextMeetingTime,
		},
	}

	if candidate.MeetingCount > 3 {
		newData = bson.M{
			"$set": bson.M{
				"next_meeting": nextMeetingTime,
				"assignee":     "5c191acea7948900011168d4", //TODO: change this
			},
		}
	}

	_, err := database.CandidatesCollection.UpdateOne(context.TODO(), bson.M{"_id": _id}, newData)

	return err
}

// CompleteMeeting completes the candidate's last meeting and increases meeting_count
func CompleteMeeting(_id string) error {
	candidate, _ := ReadCandidate(_id)

	if candidate.NextMeeting.IsZero() {
		return errors.New("candidate does not have any meeting")
	}

	newMeetingCount := candidate.MeetingCount + 1

	newData := bson.M{
		"$set": bson.M{
			"meeting_count": newMeetingCount,
			"status":        "In Progress",
			"next_meeting":  time.Time{},
		},
	}

	_, err := database.CandidatesCollection.UpdateOne(context.TODO(), bson.M{"_id": _id}, newData)

	return err
}

// DenyCandidate denies candidate
func DenyCandidate(_id string) error {
	newData := bson.M{
		"$set": bson.M{
			"status": "Denied",
		},
	}

	_, err := database.CandidatesCollection.UpdateOne(context.TODO(), bson.M{"_id": _id}, newData)

	return err
}

// AcceptCandidate acceptes candidate
func AcceptCandidate(_id string) error {
	candidate, _ := ReadCandidate(_id)

	if candidate.MeetingCount < 4 {
		return errors.New("candidates cannot be accepted before the completion of 4 meetings")
	}

	newData := bson.M{
		"$set": bson.M{
			"status": "Accepted",
		},
	}

	_, err := database.CandidatesCollection.UpdateOne(context.TODO(), bson.M{"_id": _id}, newData)

	return err
}
