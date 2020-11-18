package free_ship

import (
	"github.com/khanhnguyen234/go-microservices/_common"
	"github.com/khanhnguyen234/go-microservices/_mongo"
)

func InsertToMongo(data FreeshipCreateConsume) {
	db := _mongo.ConnectMongo()
	freeShipCollection := db.Collection("free_ship")

	ctx := _common.GetContext()
	_, err := freeShipCollection.InsertOne(ctx, data)
	if err != nil {
		panic(err)
	}
}
