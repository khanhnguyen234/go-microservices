package free_ship

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FreeshipCreateConsume struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Product_Id   string             `json:"product_id," binding:"required"`
	Is_Free_Ship bool               `json:"is_free_ship" binding:"required"`
}
