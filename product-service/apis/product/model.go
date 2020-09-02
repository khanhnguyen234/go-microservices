package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductModel struct {
	ID               primitive.ObjectID `bson:"_id"`
	Id               string             `bson:"id,omitempty"`
	Name             string             `bson:"name,omitempty"`
	Price            int                `bson:"price,omitempty"`
	Description      string             `bson:"description,omitempty"`
	PriceMin         int                `bson:"price_min,omitempty"`
	PriceMax         int                `bson:"price_max,omitempty"`
	HasComboDiscount bool               `bson:"has_combo_discount,omitempty"`
	HasFreeShip      bool               `bson:"has_free_ship,omitempty"`
	//Sku      struct{}   ``
}

type ProductCreate struct {
	Id               string `form:"id" json:"id" bson:"id,omitempty"`
	Name             string `form:"name" json:"name" binding:"required" bson:"name,omitempty"`
	Price            int    `form:"price" json:"price" binding:"required" bson:"price,omitempty"`
	Image            bool   `form:"image" json:"image" bson:"image,omitempty"`
	Video            bool   `form:"video" json:"video" bson:"video,omitempty"`
	Description      string `form:"description" json:"description" binding:"required" bson:"description,omitempty"`
	HasComboDiscount bool   `form:"has_combo_discount" json:"has_combo_discount" bson:"has_combo_discount,omitempty"`
	HasFreeShip      bool   `form:"has_free_ship" json:"password" bson:"has_free_ship,omitempty"`
	//Sku      struct{}   ``
}

type ProductSearch struct {
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
}

type ProductDetail struct {
	Id string `uri:"id" binding:"required"`
}
