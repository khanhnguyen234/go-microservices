package _mongo

import (
	"github.com/khanhnguyen234/go-microservices/_common"
)

func InsertOne(collection string, data interface{}) error {
	db := ConnectMongo()
	c := db.Collection(collection)

	ctx := _common.GetContext()
	_, err := c.InsertOne(ctx, data)

	return err
}
