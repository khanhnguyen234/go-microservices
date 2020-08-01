package _mongo

import (
    "os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "khanhnguyen234/api-service-2/common"
)

var mongodb *mongo.Database

func ConnectMongo() *mongo.Database {
    uri := os.Getenv("MONGO_URI")
    database := os.Getenv("MONGO_DATABASE")

    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    common.LogErrorService(err, "NewClient Mongo")


    ctx := common.GetContext()
    err = client.Connect(ctx)
    common.LogErrorService(err, "Connect Mongo")
    
    mongodb := client.Database(database)
    common.LogSuccess("Connect MongoDB")
    return mongodb
}

func GetDatabase() *mongo.Database {
    return mongodb
}
