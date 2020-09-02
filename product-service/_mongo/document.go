package _mongo

import (
	"khanhnguyen234/product-service/common"
)

func InsertOne(collection string, data interface{}) error {
	db := ConnectMongo()
	c := db.Collection(collection)

	ctx := common.GetContext()
	_, err := c.InsertOne(ctx, data)

	return err
}
