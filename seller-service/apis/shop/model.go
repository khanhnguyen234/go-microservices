package shop

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ShopModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	Id          string             `bson:"id,omitempty"`
	Type        string             `bson:"price,omitempty"`
	Level       int                `bson:"description,omitempty"`
	Logo        string             `bson:"has_combo_discount,omitempty"`
	Rating      float32            `bson:"has_free_ship,omitempty"`
	TimePrepare float32            `bson:"has_free_ship,omitempty"`
	RegionId    string             `bson:"has_free_ship,omitempty"`
}

type ShopCreate struct {
	Id          string    `form:"id" json:"id" bson:"id,omitempty"`
	Name        string    `form:"name" json:"name" binding:"required" bson:"name,omitempty"`
	Type        string    `form:"type" json:"type" bson:"type,omitempty"`
	Level       int       `form:"level" json:"level"  bson:"level,omitempty"`
	Logo        string    `form:"logo" json:"logo" bson:"logo,omitempty"`
	Rating      float32   `form:"rating" json:"rating" bson:"rating,omitempty"`
	TimePrepare float32   `form:"time_prepare" json:"time_prepare" bson:"time_prepare,omitempty"`
	RegionId    string    `form:"region_name" json:"region_name" bson:"region_name,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at,omitempty"`
}

type ShopSearch struct {
	Name  string `form:"name" json:"name"`
	Type  string `form:"type" json:"description"`
	Level string `form:"level" json:"description"`
}

type ShopDetail struct {
	Id string `uri:"id" binding:"required"`
}
