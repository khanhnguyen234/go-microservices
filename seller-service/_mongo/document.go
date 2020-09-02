package _mongo

import (
	"khanhnguyen234/seller-service/common"
)

type Document interface {
	InsertOne()
}

func InsertOne(collection string, data interface{}) error {
	db := ConnectMongo()
	c := db.Collection(collection)

	ctx := common.GetContext()
	_, err := c.InsertOne(ctx, data)

	return err
}
