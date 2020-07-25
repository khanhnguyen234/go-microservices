package products

import (
	"github.com/gin-gonic/gin"
)

func ProductNoAuthRegister(router *gin.RouterGroup) {
	router.GET("/", ProductList)
	router.GET("/filter", ProductFilter)
	router.GET("/detail/:id", ProductDetail)
	router.POST("/", ProductCreate)
}
