package free_ship

import (
	"khanhnguyen234/api-service-2/_mongo"
	"khanhnguyen234/api-service-2/common"
)

func InsertToMongo(data FreeshipCreateConsume) {
	db := _mongo.ConnectMongo()
	freeShipCollection := db.Collection("free_ship")

	ctx := common.GetContext()
	_, err := freeShipCollection.InsertOne(ctx, data)
	if err != nil {
		panic(err)
	}
}
