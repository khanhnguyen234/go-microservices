package freeship

type FreeshipCreateRequest struct {
	ProductId  string `form:"product_id" json:"product_id" xml:"product_id" binding:"required"`
	IsFreeShip bool   `form:"is_free_ship" json:"is_free_ship" xml:"is_free_ship" binding:"required"`
}
