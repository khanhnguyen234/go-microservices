package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShopRouters(router *gin.RouterGroup) {
	router.POST("/search", SearchShopRouter)
	router.GET("/:id", GetShopDetailRouter)
	router.POST("/", CreateShopRouter)
	router.GET("/", GetShopsRouter)
}

func CreateShopRouter(c *gin.Context) {
	var validator ShopCreate

	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shop := CreateShopController(validator)
	c.JSON(http.StatusCreated, gin.H{"result": shop})
}

func GetShopsRouter(c *gin.Context) {
	shops := GetShopsController()
	c.JSON(200, gin.H{"result": shops})
}

func GetShopDetailRouter(c *gin.Context) {
	var uri ShopDetail

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	shop := GetShopDetailController(uri.Id)

	c.JSON(200, gin.H{"result": shop})
}

func SearchShopRouter(c *gin.Context) {
	var validator ShopSearch

	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := SearchShopController(validator)
	c.JSON(http.StatusCreated, gin.H{"result": result})
}
