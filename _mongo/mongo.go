package _mongo

import (
	"github.com/khanhnguyen234/go-microservices/_common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

var mongodb *mongo.Database

func ConnectMongo() *mongo.Database {
	uri := os.Getenv("MONGO_URI")
	database := os.Getenv("MONGO_DATABASE")
	ctx := _common.GetContext()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	err = client.Ping(ctx, readpref.Primary())

	_common.LogStatus(err, "Connect Mongo")

	mongodb := client.Database(database)
	return mongodb
}

func GetDatabase() *mongo.Database {
	return mongodb
}
