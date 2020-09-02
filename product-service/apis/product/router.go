package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProductRouters(router *gin.RouterGroup) {
	router.POST("/search", SearchProductRouter)
	router.GET("/:id", GetProductDetailRouter)
	router.POST("/", CreateProductRouter)
	router.GET("/", GetProductsRouter)
}

func CreateProductRouter(c *gin.Context) {
	var validator ProductCreate

	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := CreateProductController(validator)
	c.JSON(http.StatusCreated, gin.H{"result": product})
}

func GetProductsRouter(c *gin.Context) {
	products := GetProductsController()
	c.JSON(200, gin.H{"result": products})
}

func GetProductDetailRouter(c *gin.Context) {
	var uri ProductDetail

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	product := GetProductDetailController(uri.Id)

	c.JSON(200, gin.H{"result": product})
}

func SearchProductRouter(c *gin.Context) {
	var validator ProductSearch

	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := SearchProductController(validator)
	c.JSON(http.StatusCreated, gin.H{"result": result})
}
