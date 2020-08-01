package _mongo

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "khanhnguyen234/api-service-2/common"
)

const (
	uri      = "mongodb://localhost:27017"
	database   = "api_service_2"
)

var mongodb *mongo.Database

func ConnectMongo() *mongo.Database {
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        common.LogErrorService(err, "FAILED: NewClient Mongo")
    }

    ctx := common.GetContext()
    err = client.Connect(ctx)
    if err != nil {
        common.LogErrorService(err, "FAILED: Connect Mongo")
    }
    
    mongodb := client.Database(database)
    return mongodb
}

func GetDatabase() *mongo.Database {
    return mongodb
}
