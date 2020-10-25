package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/momenteam/momentum/configs"
)

var Client *mongo.Client
var Context context.Context
var CancelFunc context.CancelFunc
var NeederCollection *mongo.Collection
var MailTemplateCollection *mongo.Collection
var NeedCollection *mongo.Collection

func Setup() {
	Client, Context, CancelFunc = getConnection(configs.GlobalConfig.Database.ConnectionString)

	getCollections()
}

func getConnection(connectionString string) (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	return client, ctx, cancel
}

func getCollections() {
	databaseName := configs.GlobalConfig.Database.DatabaseName
	database := Client.Database(databaseName)

	NeederCollection = database.Collection("Needer")
	MailTemplateCollection = database.Collection("MailTemplates")
	NeedCollection = database.Collection("Needs")
}
