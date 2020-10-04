package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductModel struct {
	ID               primitive.ObjectID `bson:"_id"`
	Id               string             `bson:"id,omitempty"`
	Name             string             `bson:"name,omitempty"`
	Price            int                `bson:"price,omitempty"`
	Description      string             `bson:"description,omitempty"`
	ImageUrl         string             `bson:"image_url,omitempty"`
	VideoUrl         string             `bson:"video_url,omitempty"`
	PriceMin         int                `bson:"price_min,omitempty"`
	PriceMax         int                `bson:"price_max,omitempty"`
	HasComboDiscount bool               `bson:"has_combo_discount,omitempty"`
	HasFreeShip      bool               `bson:"has_free_ship,omitempty"`
	//Sku      struct{}   ``
}

type ProductCreate struct {
	ID                 primitive.ObjectID `form:"_id" json:"_id" bson:"_id,omitempty"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	Id                 string             `form:"id" json:"id" bson:"id,omitempty"`
	Name               string             `form:"name" json:"name" binding:"required" bson:"name,omitempty"`
	Price              int                `form:"price" json:"price" binding:"required" bson:"price,omitempty"`
	ImageUrl           string             `form:"image_url" json:"image_url" bson:"image_url,omitempty"`
	VideoUrl           string             `form:"video_url" json:"video_url" bson:"video_url,omitempty"`
	Description        string             `form:"description" json:"description" binding:"required" bson:"description,omitempty"`
	HasComboDiscount   bool               `form:"has_combo_discount" json:"has_combo_discount" bson:"has_combo_discount,omitempty"`
	HasFreeShip        bool               `form:"has_free_ship" json:"has_free_ship" bson:"has_free_ship,omitempty"`
	FlashSale          bool               `form:"flash_sale" json:"flash_sale" bson:"flash_sale,omitempty"`
	FlashSaleUnixStart int                `form:"flash_sale_unix_start" json:"flash_sale_unix_start" bson:"flash_sale_unix_start,omitempty"`
	FlashSaleUnixEnd   int                `form:"flash_sale_unix_end" json:"flash_sale_unix_end" bson:"flash_sale_unix_end,omitempty"`
	//Sku      struct{}   ``
}

type ProductSearch struct {
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
}

type ProductFlashSale struct {
	Src     string `form:"src" json:"src"`
	Time    int    `form:"time" json:"time"`
	Limit   int64  `form:"limit" json:"limit"`
	GroupId string `form:"group_id" json:"group_id"`
}

type ProductDetail struct {
	Id string `uri:"id" binding:"required"`
}
