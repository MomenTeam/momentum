package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/momenteam/momentum/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Needer type
type Needer struct {
	ID          string    `bson:"_id" json:"id"`
	FirstName   string    `bson:"firstName" json:"firstName"`
	LastName    string    `bson:"lastName" json:"lastName"`
	Address     string    `bson:"address" json:"address"`
	Category    string    `bson:"category" json:"category"`
	PhoneNumber string    `bson:"phoneNumber" json:"phoneNumber"`
	Summary     string    `bson:"summary" json:"summary"`
	Packages    []Package `bson:"packages" json:"packages"`
	CreatedBy   string    `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	IsPublished bool      `bson:"isPublished" json:"isPublished"`
}

// Package type
type Package struct {
	ID          string     `bson:"_id" json:"id"`
	Name        string     `bson:"name" json:"name"`
	TotalPrice  int        `bson:"totalPrice" json:"totalPrice"`
	IsPublished bool       `bson:"isPublished" json:"isPublished"`
	LineItems   []LineItem `bson:"lineItems" json:"lineItems"`
}

// LineItem type
type LineItem struct {
	ID     string `bson:"_id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Price  int    `bson:"price" json:"price"`
	Amount int    `bson:"amount" json:"amount"`
}

type NeedyInformation struct {
	ID         string `json:"id"`
	Category   string `bson:"category" json:"category"`
	Summary    string `json:"summary"`
	ShortName  string `json:"shortName"`
	MaskedName string `json:"maskedName"`
}

type NeedyDetail struct {
	ID         string    `json:"id"`
	Category   string    `bson:"category" json:"category"`
	Summary    string    `json:"summary"`
	ShortName  string    `json:"shortName"`
	MaskedName string    `json:"maskedName"`
	Packages   []Package `bson:"packages" json:"packages"`
}

type Contact struct {
	ID          string `bson:"_id" json:"id"`
	FirstName   string `bson:"firstName" json:"firstName"`
	LastName    string `bson:"lastName" json:"lastName"`
	PackageId   string `bson:"packageId" json:"packageId"`
	Description string `bson:"description" json:"description"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
	Email       string `bson:"email" json:"email"`
	NeederId    string `bson:"neederId" json:"neederId"`
	Status      string `bson:"status" json:"status"`
}

type ContactForm struct {
	ID          string  `bson:"_id" json:"id"`
	FirstName   string  `bson:"firstName" json:"firstName"`
	LastName    string  `bson:"lastName" json:"lastName"`
	Package     Package `bson:"package" json:"package"`
	Description string  `bson:"description" json:"description"`
	PhoneNumber string  `bson:"phoneNumber" json:"phoneNumber"`
	Email       string  `bson:"email" json:"email"`
	NeederId    string  `bson:"neederId" json:"neederId"`
	Status      string  `bson:"status" json:"status"`
}

// GetAllNeeders func
func GetAllNeeders() ([]Needer, error) {
	needers := []Needer{}

	cursor, err := database.NeederCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var needer Needer
		if err = cursor.Decode(&needer); err != nil {
			log.Fatal(err)
		}
		needers = append(needers, needer)
	}

	return needers, err
}

// CreateNeeder func
func CreateNeeder(needer Needer) (result Needer, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Needer create error")
		}
	}()

	needer.ID = uuid.New().String()
	_, err = database.NeederCollection.InsertOne(context.Background(), needer)

	return needer, err
}

// GetNeederDetail func
func GetNeederDetail(id string) (Needer, error) {
	needer := Needer{}
	err := database.NeederCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&needer)

	return needer, err
}

// CreatePackage func
func CreatePackage(id string, packageModel Package) (packages Package, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Needer create error")
		}
	}()

	packageModel.ID = uuid.New().String()
	packageModel.LineItems = []LineItem{}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"packages": packageModel}}

	_ = database.NeederCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		&opt,
	)
	return packageModel, err
}

// CreateLineItem func
func CreateLineItem(id string, packageID string, lineItem LineItem) (lineItems LineItem, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("LineItem create error")
		}
	}()

	lineItem.ID = uuid.New().String()
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	filter := bson.M{"_id": id, "packages._id": packageID}
	update := bson.M{"$push": bson.M{"packages.$.lineItems": lineItem}}

	_ = database.NeederCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		&opt,
	)
	willAdd := lineItem.Price * lineItem.Amount
	_, err = updateTotalPrice(id, packageID, willAdd)
	return lineItem, err
}

func updateTotalPrice(id string, packageID string, price int) (isTrue bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("updateTotalPrice error")
		}
	}()
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	filter := bson.M{"_id": id, "packages._id": packageID}
	update := bson.M{"$inc": bson.M{"packages.$.totalPrice": price}}

	_ = database.NeederCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		&opt,
	)

	return true, err
}

// UpdatePackageIsPublished func
func UpdatePackageIsPublished(id string, packageID string, isPublished bool) (boolean bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("UpdatePackageIsPublished error")
		}
	}()

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	filter := bson.M{"_id": id, "packages._id": packageID}
	update := bson.M{"$set": bson.M{"packages.$.isPublished": isPublished}}
	err = database.NeederCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		&opt,
	).Err()

	return true, err
}

// UpdateNeederIsPublished func
func UpdateNeederIsPublished(id string, isPublished bool) (boolean bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("UpdateNeederIsPublished error")
		}
	}()

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isPublished": isPublished}}
	err = database.NeederCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		&opt,
	).Err()

	return true, err
}

func GetAllNeediesInformations() ([]NeedyInformation, error) {
	var neediesInformation []NeedyInformation

	cursor, err := database.NeederCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var needer Needer
		if err = cursor.Decode(&needer); err != nil {
			log.Fatal(err)
		}
		needyInformation := NeedyInformation{
			ID:         needer.ID,
			Summary:    needer.Summary,
			Category:   needer.Category,
			ShortName:  fmt.Sprintf("%c%c", needer.FirstName[0], needer.LastName[0]),
			MaskedName: fmt.Sprintf("%s %s", mask(needer.FirstName), mask(needer.LastName)),
		}

		if needer.IsPublished {
			neediesInformation = append(neediesInformation, needyInformation)
		}

	}

	return neediesInformation, err
}

func GetNeederDetailAsUser(id string) (NeedyDetail, error) {
	needer := Needer{}
	err := database.NeederCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&needer)

	var packages []Package
	for _, neederPackage := range needer.Packages {
		if neederPackage.IsPublished {
			packages = append(packages, neederPackage)
		}
	}

	needyDetail := NeedyDetail{
		ID:         needer.ID,
		Category:   needer.Category,
		Summary:    needer.Summary,
		ShortName:  fmt.Sprintf("%c%c", needer.FirstName[0], needer.LastName[0]),
		MaskedName: fmt.Sprintf("%s %s", mask(needer.FirstName), mask(needer.LastName)),
		Packages:   packages,
	}

	return needyDetail, err
}

// CreateContact func
func CreateContact(contact Contact) (result Contact, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Contact create error")
		}
	}()

	contact.ID = uuid.New().String()
	contact.Status = "Pending"
	_, err = database.ContactCollection.InsertOne(context.Background(), contact)

	return contact, err
}

// GetPackage func
func GetPackage(neederId string, packageId string) (Package, error) {
	needer := Needer{}
	resultPackage := Package{}
	err := database.NeederCollection.FindOne(context.TODO(), bson.M{"_id": neederId}).Decode(&needer)

	for _, pkg := range needer.Packages {
		if pkg.ID == packageId {
			resultPackage = pkg
		}
	}

	return resultPackage, err
}

// GetAllContactFormsWithPackage func
func GetAllContactFormsWithPackage(status string) ([]ContactForm, error) {
	var contacts []ContactForm

	cursor, err := database.ContactCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var contact Contact
		if err = cursor.Decode(&contact); err != nil {
			log.Fatal(err)
		}

		pkg, _ := GetPackage(contact.NeederId, contact.PackageId)

		contactForm := ContactForm{
			ID:          contact.ID,
			FirstName:   contact.FirstName,
			LastName:    contact.LastName,
			Package:     pkg,
			Description: contact.Description,
			PhoneNumber: contact.PhoneNumber,
			Email:       contact.Email,
			NeederId:    contact.NeederId,
			Status:      contact.Status,
		}

		if strings.ToLower(status) == "all" {
			contacts = append(contacts, contactForm)
		} else if strings.ToLower(contactForm.Status) == strings.ToLower(status) {
			contacts = append(contacts, contactForm)
		}
	}

	return contacts, err
}

// UpdateContactRequest func
func UpdateContactRequest(id string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("UpdateContactRequest error")
		}
	}()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": "Processing"}}
	err = database.ContactCollection.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
	).Err()

	return id, err
}

func mask(s string) string {
	runes := []rune(s)
	result := string(runes[0:1])
	return result + "***"
}


delete package and lineItems neederId, PackageId, LineItemId