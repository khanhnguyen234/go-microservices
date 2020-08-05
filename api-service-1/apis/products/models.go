package products

import (
	"github.com/jinzhu/gorm"
)

type ProductCreateRequest struct {
	Name  string `form:"name" json:"name" xml:"name" binding:"required"`
	Price int    `form:"price" json:"price" xml:"price" binding:"required"`
}

type ProductDetailRequest struct {
	Id string `uri:"id"`
}

type ProductSearchRequest struct {
	Name string `uri:"name"`
}

type ProductFilterRequest struct {
	Name  string `form:"name"`
	Price string `form:"price"`
}

type ProductModel struct {
	gorm.Model
	Name  string
	Price int
}
